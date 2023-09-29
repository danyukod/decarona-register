package commands

type RegisterUserDTO struct {
	Name      string
	Email     string
	Documents []DocumentDTO
	Password  string
}

type DocumentDTO struct {
	Number  string
	DocType string
}

type RegisterCarDTO struct {
	Model       string
	Brand       string
	Year        int
	Color       string
	Document    string
	Description *string
}
