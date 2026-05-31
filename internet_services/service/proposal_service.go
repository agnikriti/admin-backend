package service

import (
	"context"
	"errors"

	"agnikriti_admin_backend/internet_services/dto"
	"agnikriti_admin_backend/internet_services/model"
	"agnikriti_admin_backend/internet_services/repository"
	"agnikriti_admin_backend/utils"
)

type ProposalService struct {
	Repo *repository.ProposalRepository
}

func NewProposalService(
	repo *repository.ProposalRepository,
) *ProposalService {
	return &ProposalService{
		Repo: repo,
	}
}

func (s *ProposalService) CreateProposal(
	ctx context.Context,
	request dto.ProposalRequest,
) error {
	if request.Email == "" && request.Mobile == "" {
		return errors.New("either email or mobile is required")
	}

	proposal := model.Proposal{
		Title:       request.Title,
		Description: request.Description,
		Email:       request.Email,
		Mobile:      request.Mobile,
		Quote:       request.Quote,
	}

	err := s.Repo.SaveProposal(ctx, proposal)

	if err != nil {
		return err
	}

	err = utils.SendProposalEmail(
		proposal.Title,
		proposal.Description,
		proposal.Email,
		proposal.Mobile,
		proposal.Quote,
	)

	if err != nil {
		return err
	}

	return nil
}
