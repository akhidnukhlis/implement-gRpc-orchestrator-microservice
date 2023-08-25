package routers

import (
	"fmt"
	"github.com/akhidnukhlis/implement-gRpc-orchestrator-microservice/internal/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Service struct {
	User *usecase.AuthorHandler
}

type Serve struct {
	DB      *gorm.DB
	Router  *mux.Router
	Service Service
}

// Initialize is used to initialize db driver connection
func (s *Serve) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	//initiate database driver
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		s.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s author=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbName, DbPassword)
		s.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	s.Router = mux.NewRouter()

	s.initializeRoutes()
}

// Run is used to execute connection and run our providers
func (s *Serve) Run() {
	port := os.Getenv("APP_PORT")
	fmt.Println("Listening to port", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Router))
}
