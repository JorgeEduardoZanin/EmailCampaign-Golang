package domain_test

import (
	"email_campaign/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func createNewCampaignConstructor() (*domain.Campaign, error) {

	name := "New year campaign"
	content := "New year campaign this 2026"
	contacts := []string{"jorge@gmail.com", "marcos@gmail.com"}

	return domain.Constructor(name, content, contacts)

}

func createNewCampaignConstructorError() (*domain.Campaign, error) {

	name := ""
	content := ""
	contacts := make([]string, 0)

	return domain.Constructor(name, content, contacts)

}

func Test_CampaignConstructor(t *testing.T) {

	assert := assert.New(t)

	content := "New year campaign this 2026"

	newCampaign, _ := createNewCampaignConstructor()

	assert.NotEqual(newCampaign.Contacts, 0)
	assert.Equal(newCampaign.Content, content)
}

func Test_CampaignConstructor_CreatedAtIsNotNil(t *testing.T) {

	assert := assert.New(t)

	newCampaign, _ := createNewCampaignConstructor()

	nowTime := time.Now()
	assert.Equal(newCampaign.CreatedAt.Day(), nowTime.Day())

}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {

	assert := assert.New(t)

	newCampaign, _ := createNewCampaignConstructor()

	assert.NotNil(newCampaign.Id)
	assert.NotEmpty(newCampaign.Id)

}

func Test_NewCampaign_Error(t *testing.T) {
	assert := assert.New(t)

	_, error := createNewCampaignConstructorError()

	assert.Equal(error.Error(), "name and content is empty")
	assert.NotEmpty(error.Error())

}
