package domain_test

import (
	"email_campaign/internal/domain"
	"testing"
)

func TestCampaignConstructor(t *testing.T) {
	name := "New year campaign"
	content := "New year campaign this 2026"
	contacts := []string{"jorge@gmail.com", "marcos@gmail.com"}

	newCampaign := domain.Constructor(name, content, contacts)

	if len(newCampaign.Contacts) == 0 {
		t.Errorf("Contact list is empty")
	}
}
