package internet_services

import (
	"github.com/gin-gonic/gin"

	"agnikriti_admin_backend/database"
	"agnikriti_admin_backend/internet_services/controller"
	"agnikriti_admin_backend/internet_services/repository"
	"agnikriti_admin_backend/internet_services/service"
)

func RegisterRoutes(router *gin.Engine) {
	proposalRepository := repository.NewProposalRepository(database.DB)
	proposalService := service.NewProposalService(proposalRepository)
	proposalConroller := controller.NewProposalController(proposalService)

	internetServices := router.Group("/api/internet_services/v1")

	{
		internetServices.POST(
			"/sendProposal",
			proposalConroller.SendProposal,
		)
	}
}
