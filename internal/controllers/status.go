package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pumahawk/dctq/internal/dto"
	"github.com/Pumahawk/dctq/internal/mappers"
	"github.com/Pumahawk/dctq/internal/model"
	"github.com/Pumahawk/dctq/internal/services"
)

type StatusController struct {
	statusService  services.StatusService
	messageService services.MessageService
}

func NewStatusController(statusService services.StatusService, messageService services.MessageService) *StatusController {
	return &StatusController{
		statusService:  statusService,
		messageService: messageService,
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
		err := json.NewDecoder(r.Body).Decode(&createStatusInfoRequestDto)
		if err != nil {
			log.Printf("CreateStatusController - Unable to read body request. %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
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
		err := json.NewDecoder(r.Body).Decode(&gdto)
		if err != nil {
			log.Printf("CreateStatusController - Unable to read body request. %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		gum := *mappers.ToStatusUpdateModel(&gdto)
		err = c.statusService.UpdateById(id, gum)
		if err != nil {
			log.Printf("UpdateStatusController - error on update. %s", err)
			return
		}
		status, err := c.statusService.GetById(id)
		if err != nil {
			log.Printf("UpdateStatusController - error on update. %s", err)
			return
		}
		c.messageService.Send(&model.CreateMessageModel{
			ProjectId: status.Id,
			Type:      "status-update",
			Message:   status.Data,
		})
		response := mappers.ToUpdateStatusResponse(status)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
}
