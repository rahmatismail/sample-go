package usecase

import (
	"github.com/rahmatismail/sample-go/internal/employee"
	"github.com/rahmatismail/sample-go/internal/office"
)

type (
	usecase struct {
		dep Dependency
	}

	Dependency struct {
		OfficeStaticRepo   office.StaticRepository
		EmployeeStaticRepo employee.StaticRepository
	}
)

func New(d Dependency) office.Usecase {
	return &usecase{
		dep: d,
	}
}

func (u *usecase) GetData(ID int64) (cd office.CompleteData) {
	return
}
