package service

import "github.com/Prodget/store"

type Fservice interface {
	GetById(id int) (*store.Details, error)
}