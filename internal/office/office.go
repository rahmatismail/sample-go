package office

// All type for office model
type (
	Data struct {
		ID      int64
		Name    string
		Address string
	}

	CompleteData struct {
		Data     Data
		Employee []string
	}
)

type (
	StaticRepository interface {
		Get(ID int64) Data
	}

	Usecase interface {
		GetData(ID int64) CompleteData
	}
)
