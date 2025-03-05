package app

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type Application struct {
	Data map[string]User
}

func NewApplication() *Application {
	return &Application{
		Data: make(map[string]User),
	}
}
