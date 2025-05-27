package handler

import (
	"net/http"
	"strconv"

	Kursach_UD "github.com/Ocas17/Kursach_UD"
	"github.com/gin-gonic/gin"
)


func (h *Handler) createClient(c *gin.Context) {
	var input Kursach_UD.Client
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Client.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Clients
// @Tags clients
// @Description get all clients
// @ID get-all-clients
// @Accept  json
// @Produce  json
// @Success 200 {array} Kursach_UD.Client
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/clients [get]
func (h *Handler) getAllClients(c *gin.Context) {
	clients, err := h.services.Client.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, clients)
}

// @Summary Get Client By Id
// @Tags clients
// @Description get client by id
// @ID get-client-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Client ID"
// @Success 200 {object} Kursach_UD.Client
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/clients/{id} [get]
func (h *Handler) getClientById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	client, err := h.services.Client.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, client)
}

// @Summary Update Client
// @Tags clients
// @Description update client
// @ID update-client
// @Accept  json
// @Produce  json
// @Param id path int true "Client ID"
// @Param input body Kursach_UD.UpdateClientInput true "client update info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/clients/{id} [put]
func (h *Handler) updateClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input Kursach_UD.UpdateClientInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Client.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete Client
// @Tags clients
// @Description delete client
// @ID delete-client
// @Accept  json
// @Produce  json
// @Param id path int true "Client ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/clients/{id} [delete]
func (h *Handler) deleteClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	if err := h.services.Client.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
