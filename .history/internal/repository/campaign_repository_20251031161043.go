package repository

import "email_campaign/internal/domain"

type campaign_repository interface {
	save(campaign *domain.Campaign)
}
