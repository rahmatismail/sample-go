package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	employeeRepoPkg "github.com/rahmatismail/sample-go/internal/employee/repository"
	officeHandlerRestPkg "github.com/rahmatismail/sample-go/internal/office/handler/rest"
	officeRepoPkg "github.com/rahmatismail/sample-go/internal/office/repository"
	officeUsecasePkg "github.com/rahmatismail/sample-go/internal/office/usecase"
)

type (
	CoreDependency struct {
		HttpRouter *httprouter.Router
	}
)

func main() {
	// core
	router := httprouter.New()

	cd := CoreDependency{
		HttpRouter: router,
	}
	// end init core dependency
	InitSampleGo(cd)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func InitSampleGo(dep CoreDependency) {
	// init repository
	// employee
	employeeStaticRepo := employeeRepoPkg.New()

	// office
	officeStaticRepo := officeRepoPkg.New()

	// all usecase
	officeUsecaseDependency := officeUsecasePkg.Dependency{
		OfficeStaticRepo:   officeStaticRepo,
		EmployeeStaticRepo: employeeStaticRepo,
	}
	officeUsecase := officeUsecasePkg.New(officeUsecaseDependency)

	// init handler
	// office
	officeHandlerRestDependency := officeHandlerRestPkg.Dependency{
		OfficeUsecase: officeUsecase,
		HttpRouter:    dep.HttpRouter,
	}
	officeHandlerRestPkg.RegisterRoute(officeHandlerRestDependency)

}
