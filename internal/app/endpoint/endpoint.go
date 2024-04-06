package endpoint

import (
	"encoding/json"
	"log"
	"net/http"

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
			log.Println(err)
		}

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	}
}
