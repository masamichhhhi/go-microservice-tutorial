package transport

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/masamichhhhi/watermark-service/api/v1/pb/watermark"
	"github.com/masamichhhhi/watermark-service/pkg/watermark/endpoints"
)

type grpcServer struct {
	get           grpctransport.Handler
	status        grpctransport.Handler
	addDocument   grpctransport.Handler
	watermark     grpctransport.Handler
	serviceStatus grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) watermark.WatermarkServer {

}
