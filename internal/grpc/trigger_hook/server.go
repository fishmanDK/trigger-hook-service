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

func (s *serverAPI) ScheduleFullDeletion(ctx context.Context, request *trg_hk.CreateFullDeletionRequest) (*trg_hk.NewDeletionResponse, error) {
	if request.GetBannerID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "bannerID is required")
	}

	err := s.service.ScheduleFullDeletion(ctx, request.GetBannerID())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to make a new application")
	}

	return &trg_hk.NewDeletionResponse{Success: true}, nil
}

func (s *serverAPI) SchedulePartialDeletion(ctx context.Context, request *trg_hk.CreatePartialDeletionRequest) (*trg_hk.NewDeletionResponse, error) {
	panic("implemented me")
}
