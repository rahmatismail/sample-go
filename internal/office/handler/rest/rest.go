package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rahmatismail/sample-go/internal/office"
)

type (
	officeRest struct {
		d Dependency
	}

	Dependency struct {
		HttpRouter    *httprouter.Router
		OfficeUsecase office.Usecase
	}
)

func RegisterRoute(d Dependency) {
	or := &officeRest{
		d: d,
	}

	d.HttpRouter.GET("/office/test", or.Test)
}

func (or *officeRest) Test(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("test"))
}
