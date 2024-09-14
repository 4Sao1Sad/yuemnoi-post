package db

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/bpremika/post/internal/handler"
	"github.com/bpremika/post/internal/repository"
	"github.com/bpremika/post/proto/post"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes a connection to the PostgreSQL database using GORM
func InitDB() {
	// Set the PostgreSQL connection details
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Format the DSN (Data Source Name) string for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	// Connect to the PostgreSQL database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connection established")
}

func ServerInit(db *gorm.DB) error {
	port := ":8081"

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer func() {
		listen.Close()
	}()

	fmt.Printf("Go gRPC server in port %v!", port)
	grpcServer := grpc.NewServer()
	// register
	lendingPostRepo := repository.NewLendingPostRepository(db)
	lendingPostServer := handler.NewLendingPostGRPC(lendingPostRepo)
	// put register server here
	post.RegisterLendingPostServiceServer(grpcServer, lendingPostServer)
	err = grpcServer.Serve(listen)
	if err != nil {
		return fmt.Errorf("error to serve: %v", err.Error())
	}

	return nil
}
