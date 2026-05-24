package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"agnikriti_admin_backend/internet_services/dto"
	"agnikriti_admin_backend/internet_services/service"
	"agnikriti_admin_backend/utils"
)

type ProposalController struct {
	Service *service.ProposalService
}

func NewProposalController(
	service *service.ProposalService,
) *ProposalController {
	return &ProposalController{
		Service: service,
	}
}

func (pc *ProposalController) SendProposal(c *gin.Context) {
	request := dto.ProposalRequest{}
	ctx := c.Request.Context()

	err := c.BindJSON(&request)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			err.Error(),
		)

		return
	}

	err = pc.Service.CreateProposal(ctx, request)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"Failed to save proposal",
			err.Error(),
		)

		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Proposal saved successfully",
		nil,
	)
}
