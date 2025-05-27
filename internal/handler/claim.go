package handler

import (
	"net/http"
	"strconv"

	Kursach_UD "github.com/Ocas17/Kursach_UD"
	"github.com/gin-gonic/gin"
)

// @Summary Create Claim
// @Tags claims
// @Description create claim for policy
// @ID create-claim
// @Accept  json
// @Produce  json
// @Param policy_id path int true "Policy ID"
// @Param input body Kursach_UD.Claim true "claim info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/policies/{policy_id}/claims [post]
func (h *Handler) createClaim(c *gin.Context) {
	policyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid policy id parameter")
		return
	}

	var input Kursach_UD.Claim
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.PolicyId = policyId

	id, err := h.services.Claim.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Claims
// @Tags claims
// @Description get all claims for policy
// @ID get-all-claims
// @Accept  json
// @Produce  json
// @Param policy_id path int true "Policy ID"
// @Success 200 {array} Kursach_UD.Claim
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/policies/{policy_id}/claims [get]
func (h *Handler) getAllClaims(c *gin.Context) {
	policyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid policy id parameter")
		return
	}

	claims, err := h.services.Claim.GetAll(policyId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, claims)
}

// @Summary Get Claim By Id
// @Tags claims
// @Description get claim by id
// @ID get-claim-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Claim ID"
// @Success 200 {object} Kursach_UD.Claim
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/claims/{id} [get]
func (h *Handler) getClaimById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	claim, err := h.services.Claim.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, claim)
}

// @Summary Update Claim
// @Tags claims
// @Description update claim
// @ID update-claim
// @Accept  json
// @Produce  json
// @Param id path int true "Claim ID"
// @Param input body Kursach_UD.UpdateClaimInput true "claim update info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/claims/{id} [put]
func (h *Handler) updateClaim(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input Kursach_UD.UpdateClaimInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Claim.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete Claim
// @Tags claims
// @Description delete claim
// @ID delete-claim
// @Accept  json
// @Produce  json
// @Param id path int true "Claim ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/claims/{id} [delete]
func (h *Handler) deleteClaim(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	if err := h.services.Claim.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
