package dto

type RegisterUserDTO struct {
	Name      string
	Email     string
	Gender    string
	Documents []DocumentDTO
	Password  string
}

type DocumentDTO struct {
	Number  string
	DocType string
}

type RegisterCarDTO struct {
	OwnerId     string
	Model       string
	Brand       string
	Year        int
	Color       string
	Plate       string
	Description *string
}
