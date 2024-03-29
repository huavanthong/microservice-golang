package gapi

import (
	"context"
	"time"

	pb "github.com/huavanthong/microservice-golang/email-grpc/proto/email"
	"github.com/huavanthong/microservice-golang/user-api-v3/utils"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.GenericResponse, error) {
	code := req.GetVerificationCode()

	verificationCode := utils.Encode(code)

	query := bson.D{{Key: "verificationCode", Value: verificationCode}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "verified", Value: true}, {Key: "updated_at", Value: time.Now()}}}, {Key: "$unset", Value: bson.D{{Key: "verificationCode", Value: ""}}}}
	result, err := server.userCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if result.MatchedCount == 0 {
		return nil, status.Errorf(codes.PermissionDenied, "Could not verify email address")
	}

	res := &pb.GenericResponse{
		Status:  "success",
		Message: "Email verified successfully",
	}
	return res, nil
}
