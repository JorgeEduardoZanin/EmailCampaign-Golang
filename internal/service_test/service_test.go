package service_test

import (
	"email_campaign/internal/domain"
	"email_campaign/internal/dto/request"
	"email_campaign/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateNewCampaignConstructor() *request.Campaign_request_dto {

	name := "New year campaign"
	content := "New year campaign this 2026"
	contacts := []string{"jorge@gmail.com", "marcos@gmail.com"}

	return &request.Campaign_request_dto{
		Name:     name,
		Content:  content,
		Contacts: contacts,
	}

}

func CreateNewCampaignConstructorError() (*domain.Campaign, error) {

	name := ""
	content := ""
	contacts := make([]string, 0)

	return domain.Constructor(name, content, contacts)

}

func Test_ServiceCampaign(t *testing.T) {
	assert := assert.New(t)
	service := service.CampaignService{}

	newCampaign := CreateNewCampaignConstructor()

	id, err := service.Create(*newCampaign)

	assert.Nil(err)
	assert.NotEmpty(id)

}
