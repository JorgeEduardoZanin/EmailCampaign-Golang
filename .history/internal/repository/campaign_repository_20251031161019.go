package repository

import "email_campaign/internal/domain"

type campaign_repository{
	save(campaign *Campaign)
}