package db

import (
	"fmt"
	"log"
	"net"

	"github.com/bpremika/post/internal/config"
	"github.com/bpremika/post/internal/handler"
	"github.com/bpremika/post/internal/repository"
	post "github.com/bpremika/post/proto/post"
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

	fmt.Println(dsn)
	// Connect to the PostgreSQL database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connection established")
}

func ServerInit(cfg *config.Config, db *gorm.DB) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer func() {
		listen.Close()
	}()

	fmt.Printf("Go gRPC server in port %v!", cfg.Port)
	grpcServer := grpc.NewServer()
	// register
	lendingPostRepo := repository.NewLendingPostRepository(db)
	lendingPostServer := handler.NewLendingPostGRPC(lendingPostRepo)
	borrowingPostRepo := repository.NewBorrowingPostRepository(db)
	borrowingPostServer := handler.NewBorrowingPostGRPC(borrowingPostRepo)
	// put register server here
	post.RegisterLendingPostServiceServer(grpcServer, lendingPostServer)
	post.RegisterBorrowingPostServiceServer(grpcServer, borrowingPostServer)
	err = grpcServer.Serve(listen)
	if err != nil {
		return fmt.Errorf("error to serve: %v", err.Error())
	}

	return nil
}
