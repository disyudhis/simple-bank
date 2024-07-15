package gapi

import (
	"fmt"

	db "github.com/disyudhis/simplebank/db/sqlc"
	"github.com/disyudhis/simplebank/pb"
	"github.com/disyudhis/simplebank/token"
	"github.com/disyudhis/simplebank/util"
	"github.com/disyudhis/simplebank/worker"
)

// Server serves gRPC request for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}
	return server, nil
}
