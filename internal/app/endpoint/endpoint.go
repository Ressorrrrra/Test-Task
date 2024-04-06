package endpoint

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Ressorrrrra/Test-Task/internal/app/data/order"
	service "github.com/Ressorrrrra/Test-Task/internal/app/services"
)

type Endpoint struct {
	S *service.Service
}

func New(s *service.Service) (e *Endpoint) {
	e = &Endpoint{S: s}
	return
}

func (e *Endpoint) GetAll(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		response, err := e.S.Get()
		if err != nil {
			log.Fatal(err)
		}

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (e *Endpoint) GetById(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		head := request.Header.Get("Content-Type")
		if !strings.Contains(head, "application/json") {
			writer.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		type IdObject struct {
			Id int
		}
		var id IdObject

		err := json.NewDecoder(request.Body).Decode(&id)
		if err != nil {
			http.Error(writer, fmt.Sprintln(err), http.StatusBadRequest)
			return
		}

		order, err := e.S.GetById(id.Id)
		if err != nil {
			http.Error(writer, fmt.Sprintln(err), http.StatusNotFound)
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(order)

	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		log.Println((request.Method))
	}
}

func (e *Endpoint) Create(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		head := request.Header.Get("Content-Type")
		if !strings.Contains(head, "application/json") {
			writer.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		var doc order.Order

		err := json.NewDecoder(request.Body).Decode(&doc)
		if err != nil {
			http.Error(writer, fmt.Sprintln(err), http.StatusBadRequest)
			return
		}
		err = e.S.Create(doc)
		if err != nil {
			http.Error(writer, fmt.Sprintln(err), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		log.Println((request.Method))
	}
}

func (e *Endpoint) Update(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPut {
		head := request.Header.Get("Content-Type")
		if !strings.Contains(head, "application/json") {
			writer.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		var doc order.Order

		err := json.NewDecoder(request.Body).Decode(&doc)
		if err != nil {
			http.Error(writer, fmt.Sprintln(err), http.StatusBadRequest)
			return
		}
		err = e.S.Update(doc)
		if err != nil {
			http.Error(writer, fmt.Sprintln(err), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(doc)

	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		log.Println((request.Method))
	}
}

func (e *Endpoint) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodDelete {
		head := request.Header.Get("Content-Type")
		if !strings.Contains(head, "application/json") {
			writer.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		type IdObject struct {
			Id int
		}
		var id IdObject

		err := json.NewDecoder(request.Body).Decode(&id)
		if err != nil {
			http.Error(writer, fmt.Sprintln(err), http.StatusBadRequest)
			return
		}
		err = e.S.Delete(id.Id)
		if err != nil {
			http.Error(writer, fmt.Sprintln(err), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusNoContent)

	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		log.Println((request.Method))
	}
}
