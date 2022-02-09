package products

import (
	//"Prodget/models"
	
	"database/sql"
	//"errors"

	"github.com/Prodget/store"
	

)
type prodStore struct {
	db *sql.DB
}
func New(db *sql.DB) store.Fstore{
  return &prodStore{db}
}


func (p *prodStore) GetById(id int) (*store.Details,error){
	res:=p.db.QueryRow("select * from products where id = ?",id)

    var d store.Details
	err := res.Scan(&d.Id,&d.Name,&d.Type)
    if err != nil {
	     return nil,err
	}

 return &d, nil
}