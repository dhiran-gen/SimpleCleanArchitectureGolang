package service

import "github.com/Prodget/store"

type servProd struct {
	s store.Fstore
}

func New(s store.Fstore) servProd {
	return servProd{s}
}

func (srv servProd) GetById(id int) (*store.Details,error){
	return srv.s.GetById(id)
}

