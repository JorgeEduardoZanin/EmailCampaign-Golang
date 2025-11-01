package service

import (
	"email_campaign/internal/dto/request"
	"email_campaign/internal/repository"
)

type Campaign_service struct {
	campaign_repository repository.Campaign_repository
}

func (service *campaign_service) create(campaign_dto request.Campaign_request_dto)
