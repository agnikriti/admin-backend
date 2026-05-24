package dto

type ProposalRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
}
