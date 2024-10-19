package util

import (
	"context"
	"log"

	activitypb "github.com/bpremika/post/proto/activity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CallActivityLogService(userID uint64, logDetail string) error {
	conn, err := grpc.NewClient("localhost:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to ActivityLogService: %v", err)
		return err
	}
	defer conn.Close()

	client := activitypb.NewActivityLogServiceClient(conn)

	req := &activitypb.CreateActivityLogRequest{
		LogDetail: logDetail,
		UserId:    userID,
	}

	_, err = client.CreateActivityLog(context.Background(), req)
	if err != nil {
		log.Printf("Error calling CreateActivityLog: %v", err)
		return err
	}
	log.Printf("Activity log created for user %v: %s", userID, logDetail)
	return nil
}
