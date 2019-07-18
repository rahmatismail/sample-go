package repository

import "github.com/rahmatismail/sample-go/internal/employee"

type (
	staticRepository struct{}
)

func New() employee.StaticRepository {
	return &staticRepository{}
}

func (sr *staticRepository) Get(ID int64) (d employee.Data) {
	return
}

func (sr *staticRepository) GetEmployeesNameByOffice(officeID int64) (names []string) {
	return
}
