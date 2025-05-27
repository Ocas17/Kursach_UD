package handler

import (
	"net/http"
	"strconv"

	Kursach_UD "github.com/Ocas17/Kursach_UD"
	"github.com/gin-gonic/gin"
)

// @Summary Create Policy
// @Tags policies
// @Description create policy for client
// @ID create-policy
// @Accept  json
// @Produce  json
// @Param client_id path int true "Client ID"
// @Param input body Kursach_UD.Policy true "policy info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/clients/{client_id}/policies [post]
func (h *Handler) createPolicy(c *gin.Context) {
	clientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid client id parameter")
		return
	}

	var input Kursach_UD.Policy
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.ClientId = clientId

	id, err := h.services.Policy.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Policies
// @Tags policies
// @Description get all policies for client
// @ID get-all-policies
// @Accept  json
// @Produce  json
// @Param client_id path int true "Client ID"
// @Success 200 {array} Kursach_UD.Policy
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/clients/{client_id}/policies [get]
func (h *Handler) getAllPolicies(c *gin.Context) {
	clientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid client id parameter")
		return
	}

	policies, err := h.services.Policy.GetAll(clientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, policies)
}

// @Summary Get Policy By Id
// @Tags policies
// @Description get policy by id
// @ID get-policy-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Policy ID"
// @Success 200 {object} Kursach_UD.Policy
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/policies/{id} [get]
func (h *Handler) getPolicyById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	policy, err := h.services.Policy.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, policy)
}

// @Summary Update Policy
// @Tags policies
// @Description update policy
// @ID update-policy
// @Accept  json
// @Produce  json
// @Param id path int true "Policy ID"
// @Param input body Kursach_UD.UpdatePolicyInput true "policy update info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/policies/{id} [put]
func (h *Handler) updatePolicy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input Kursach_UD.UpdatePolicyInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Policy.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete Policy
// @Tags policies
// @Description delete policy
// @ID delete-policy
// @Accept  json
// @Produce  json
// @Param id path int true "Policy ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/policies/{id} [delete]
func (h *Handler) deletePolicy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	if err := h.services.Policy.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
