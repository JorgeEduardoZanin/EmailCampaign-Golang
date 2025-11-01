package repository

import "email_campaign/internal/domain"

type Campaign_repository interface {
	save(campaign *domain.Campaign)
}
