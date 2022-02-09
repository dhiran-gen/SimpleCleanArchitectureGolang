package store

type Details struct{
	Id int
	Name string
	Type string
}
type Fstore interface {
	GetById(id int) (*Details,error)
}