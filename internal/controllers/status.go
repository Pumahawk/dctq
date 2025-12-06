package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pumahawk/dctq/internal/dto"
	"github.com/Pumahawk/dctq/internal/mappers"
	"github.com/Pumahawk/dctq/internal/services"
)

type StatusController struct {
	statusService services.StatusService
}

func NewStatusController(statusService services.StatusService) *StatusController {
	return &StatusController{
		statusService,
	}
}

func (c *StatusController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := c.statusService.GetAll()
		if err != nil {
			log.Printf("StatusController - error. %s", err)
			return
		}
		statusDto := mappers.ToGetAllStatusResponseDto(status)
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(statusDto)
		if err != nil {
			log.Printf("StatusController - error. %s", err)
			return
		}
	}
}

func (c *StatusController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createStatusInfoRequestDto dto.CreateStatusInfoRequestDto
		json.NewDecoder(r.Body).Decode(&createStatusInfoRequestDto)
		status, err := c.statusService.Create(*mappers.ToSimplStatusCreateInfoModel(&createStatusInfoRequestDto))
		if err != nil {
			log.Printf("CreateStatusController - error. %s", err)
			return
		}
		responseDto := mappers.ToCreateStatusInfoResponseDto(status)
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(responseDto)
		if err != nil {
			log.Printf("CreateStatusController - error. %s", err)
			return
		}
	}
}

func (c *StatusController) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		statusModel, err := c.statusService.GetById(id)
		if err != nil {
			log.Printf("GetStatusByIdController - error. %s", err)
			return
		}
		responseDto := mappers.ToGetStatusByIdResponseDto(statusModel)
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(responseDto)
		if err != nil {
			log.Printf("GetStatusByIdController - error. %s", err)
			return
		}
	}
}

func (c *StatusController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var gdto dto.UpdateStatusInfoRequestDto
		json.NewDecoder(r.Body).Decode(&gdto)
		gum := *mappers.ToStatusUpdateModel(&gdto)
		err := c.statusService.UpdateById(id, gum)
		if err != nil {
			log.Printf("UpdateStatusController - error on update. %s", err)
			return
		}
		status, err := c.statusService.GetById(id)
		if err != nil {
			log.Printf("UpdateStatusController - error on update. %s", err)
			return
		}
		response := mappers.ToUpdateStatusResponse(status)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
}
