package db

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/bpremika/post/internal/config"
	grpcHandler "github.com/bpremika/post/internal/handler/grpc"
	restHandler "github.com/bpremika/post/internal/handler/rest"
	"github.com/bpremika/post/internal/repository"
	"github.com/bpremika/post/internal/route"
	post "github.com/bpremika/post/proto/post"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes a connection to the PostgreSQL database using GORM
func InitDB(cfg *config.Config) {
	// Format the DSN (Data Source Name) string for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.Db.Host, cfg.Db.Username, cfg.Db.Password, cfg.Db.Database, cfg.Db.Port)

	log.Println(dsn)
	// Connect to the PostgreSQL database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connection established")
}

func ServerInit(cfg *config.Config, db *gorm.DB) {
	// Configure gRPC
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listen.Close()

	log.Printf("Go gRPC server is running on port %v!", cfg.GRPCPort)
	grpcServer := grpc.NewServer()

	// Fiber App
	app := fiber.New()
	app.Use(requestid.New())

	// Repository and Handlers
	lendingPostRepo := repository.NewLendingPostRepository(db)
	borrowingPostRepo := repository.NewBorrowingPostRepository(db)

	// gRPC Handlers
	lendingPostGRPCServer := grpcHandler.NewLendingPostGRPC(lendingPostRepo)
	borrowingPostGRPCServer := grpcHandler.NewBorrowingPostGRPC(borrowingPostRepo)

	// REST Handlers
	lendingPostRestHandler := restHandler.NewLendingPostRest(lendingPostRepo)
	borrowingPostRestHandler := restHandler.NewBorrowingPostRest(borrowingPostRepo)

	// Register gRPC services
	post.RegisterLendingPostServiceServer(grpcServer, lendingPostGRPCServer)
	post.RegisterBorrowingPostServiceServer(grpcServer, borrowingPostGRPCServer)

	// Register REST routes
	r := route.NewHandler(borrowingPostRestHandler, lendingPostRestHandler)
	r.RegisterRouter(app, cfg)

	// Run gRPC server in a separate goroutine
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve gRPC server: %v", err)
		}
	}()

	// Run Fiber server
	if err := app.Listen(":" + strconv.Itoa(int(cfg.Port))); err != nil {
		log.Fatalf("failed to serve Fiber server: %v", err)
	}

}
