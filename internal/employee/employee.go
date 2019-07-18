package employee

type (
	Data struct {
		ID   int64
		Name string
	}
)

type (
	StaticRepository interface {
		Get(ID int64) Data
		GetEmployeesNameByOffice(officeID int64) []string
	}
)
