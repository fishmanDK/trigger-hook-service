package trigger_hook

import (
	"context"
	trigger_hook_v1 "github.com/fishmanDK/proto_avito_test_task/protos/gen/go/trigger_hook"
	"google.golang.org/grpc"
)

type serverAPI struct {
	trigger_hook_v1.UnimplementedTriggerHookManagerServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	trigger_hook_v1.RegisterTriggerHookManagerServer(gRPC, &serverAPI{})
}

func (s *serverAPI) ScheduleFullDeletion(ctx context.Context, request *trigger_hook_v1.CreateFullDeletionRequest) (*trigger_hook_v1.NewDeletionResponse, error) {
	panic("implemented me")
}

func (s *serverAPI) SchedulePartialDeletion(ctx context.Context, request *trigger_hook_v1.CreatePartialDeletionRequest) (*trigger_hook_v1.NewDeletionResponse, error) {
	panic("implemented me")
}
