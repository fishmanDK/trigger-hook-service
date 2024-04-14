package trigger_hook

import (
	"context"
	"github.com/fishmanDK/proto_avito_test_task/protos/gen/go/trigger_hook"
	"github.com/fishmanDK/trigger_service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	trg_hk.UnimplementedTriggerHookManagerServer
	service *service.Service
}

func RegisterServerAPI(gRPC *grpc.Server, service *service.Service) {
	trg_hk.RegisterTriggerHookManagerServer(gRPC, &serverAPI{service: service})
}

func (s *serverAPI) ScheduleDeletion(ctx context.Context, request *trg_hk.CreateDeletionRequest) (*trg_hk.NewDeletionResponse, error) {
	err := s.service.ScheduleDeletion(ctx, request.GetBannerID(), request.GetTagID(), request.GetFeatureID())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to make a new application")
	}

	return &trg_hk.NewDeletionResponse{Success: true}, nil
}
