package routers

import (
	"github.com/akhidnukhlis/implement-gRpc-orchestrator-microservice/config/providers/grpc/client"
	"github.com/akhidnukhlis/implement-gRpc-orchestrator-microservice/config/providers/grpc/servicecontract"
	"github.com/akhidnukhlis/implement-gRpc-orchestrator-microservice/internal/service/author"
	"github.com/akhidnukhlis/implement-gRpc-orchestrator-microservice/internal/usecase"
)

func (se *Serve) initializeRoutes() {
	//======================== REPOSITORIES ========================
	//initiate repository
	atrgrpc := client.ServiceUser()
	r := servicecontract.NewGrpcService(atrgrpc)

	//======================== ROUTER ========================
	//Setting Services
	//Setting User Service

	//=== AUTHOR ===
	s := author.NewService(r)
	h := usecase.NewAuthorHandler(s)
	//=========================================================

	//======================== ENDPOINT ========================
	//Initialize endpoint route

	//=== AUTHOR ===
	se.Router.HandleFunc("/author/create", h.RegisterNewAuthor).Methods("POST")
	se.Router.HandleFunc("/author/detail/{id}", h.FindUserByAuthorID).Methods("GET")
	//==========================================================
}
