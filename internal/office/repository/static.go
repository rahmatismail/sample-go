package repository

import "github.com/rahmatismail/sample-go/internal/office"

type (
	staticRepository struct{}
)

func New() office.StaticRepository {
	return &staticRepository{}
}

func (sr *staticRepository) Get(ID int64) (od office.Data) {
	return
}
