package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"agnikriti_admin_backend/internet_services/model"
)

type ProposalRepository struct {
	DB *pgx.Conn
}

func NewProposalRepository(
	db *pgx.Conn,
) *ProposalRepository {

	return &ProposalRepository{
		DB: db,
	}
}

func (r *ProposalRepository) SaveProposal(
	ctx context.Context,
	proposal model.Proposal,
) error {
	query := `
		INSERT INTO email_requests (
			title,
			description,
			email,
			mobile
		)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		proposal.Title,
		proposal.Description,
		proposal.Email,
		proposal.Mobile,
	)

	return err
}
