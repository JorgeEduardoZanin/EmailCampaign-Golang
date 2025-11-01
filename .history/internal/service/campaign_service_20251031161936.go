package service

import "email_campaign/internal/repository"

type campaign_service struct {
	campaign_repository repository.Campaign_repository
}
