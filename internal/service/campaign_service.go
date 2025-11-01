package service

import (
	"email_campaign/internal/domain"
	"email_campaign/internal/dto/request"
	"email_campaign/internal/repository"
)

type CampaignService struct {
	CampaignRepository repository.Campaign_repository
}

func (service *CampaignService) Create(campaign_dto request.Campaign_request_dto) (string, error) {

	campaign, err := domain.Constructor(campaign_dto.Name, campaign_dto.Content, campaign_dto.Contacts)

	if err != nil {
		return "", err
	}

	return campaign.Id, nil
}
