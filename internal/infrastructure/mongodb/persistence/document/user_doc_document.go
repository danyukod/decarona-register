package document

type IUserDocDocument interface {
	GetDocType() string
	GetDocNumber() string
}

type UserDocDocument struct {
	DocType string `bson:"doc_type"`
	Number  string `bson:"doc_number"`
}

func (d *UserDocDocument) GetDocType() string {
	return d.DocType
}

func (d *UserDocDocument) GetDocNumber() string {
	return d.Number
}
