package routers

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/handler"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/internal/src/author"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/service/grpc/client"
	"github.com/akhidnukhlis/implement-gRpc-microservice-orchestrator/service/grpc/servicecontract"
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
	h := handler.NewAuthorHandler(s)
	//=========================================================

	//======================== ENDPOINT ========================
	//Initialize endpoint route

	//=== AUTHOR ===
	se.Router.HandleFunc("/author/create", h.RegisterNewAuthor).Methods("POST")
	se.Router.HandleFunc("/author/{id}/find", h.FindUserByAuthorID).Methods("GET")
	//==========================================================

}
