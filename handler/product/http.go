package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/Prodget/service"
	"github.com/gorilla/mux"
)

type HttpH struct {
	Srv service.Fservice
}

func (h HttpH) GetById(wr http.ResponseWriter, r *http.Request) {
	//mux.NewRouter()

	Param := mux.Vars(r)

	v := Param["id"]

	id, _ := strconv.Atoi(v)

	res, err := h.Srv.GetById(id)

	if err != nil {
		b, _ := json.Marshal(err)
		wr.Write(b)
		return
	}
	b, _ := json.Marshal(res)
	wr.Write(b)
}

func (h HttpH) PutById(wr http.ResponseWriter, r *http.Request) {
	//mux.NewRouter()

	Param := mux.Vars(r)

	v := Param["id"]

	id, _ := strconv.Atoi(v)

	res, err := h.Srv.PutById(id)

	if err != nil {
		b, _ := json.Marshal(err)
		wr.Write(b)
		return
	}
	b, _ := json.Marshal(res)
	wr.Write(b)
}
