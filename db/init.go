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

func ServerInit(cfg *config.Config, db *gorm.DB) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer func() {
		listen.Close()
	}()

	log.Printf("Go gRPC server in port %v!", cfg.GRPCPort)

	grpcServer := grpc.NewServer()
	app := fiber.New()
	app.Use(requestid.New())
	// register
	//repo
	lendingPostRepo := repository.NewLendingPostRepository(db)
	borrowingPostRepo := repository.NewBorrowingPostRepository(db)

	//grpc
	lendingPostGPRCServer := grpcHandler.NewLendingPostGRPC(lendingPostRepo)
	borrowingPostGPRCServer := grpcHandler.NewBorrowingPostGRPC(borrowingPostRepo)

	//rest
	lendingPostRestHandler := restHandler.NewLendingPostRest(lendingPostRepo)
	borrowingPostRestHandler := restHandler.NewBorrowingPostRest(borrowingPostRepo)

	// put register server here
	post.RegisterLendingPostServiceServer(grpcServer, lendingPostGPRCServer)
	post.RegisterBorrowingPostServiceServer(grpcServer, borrowingPostGPRCServer)

	r := route.NewHandler(borrowingPostRestHandler, lendingPostRestHandler)
	r.RegisterRouter(app, cfg)

	app.Listen(":" + strconv.Itoa(int(cfg.Port)))
	err = grpcServer.Serve(listen)
	if err != nil {
		return fmt.Errorf("error to serve: %v", err.Error())
	}

	return nil
}
