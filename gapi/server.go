package gapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "simple_bank/db/sqlc"
	"simple_bank/pb"
	"simple_bank/token"
	"simple_bank/utils"
	"simple_bank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          utils.Config
	store           db.Store
	tokenMaker      token.Maker
	router          *gin.Engine
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server
func NewServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMarker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMarker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
