package twilio

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/terraform-provider-twilio/twilio/resources"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ProviderName = "twilio"

var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		ProviderName: func() (*schema.Provider, error) { return Provider(), nil },
	}
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestTwilioResourcesMap(t *testing.T) {
	twilioResources := resources.NewTwilioResources()
	assert.NotNil(t, twilioResources.Map["twilio_api_accounts_messages_v2010"])
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TWILIO_ACCOUNT_SID"); v == "" {
		t.Fatal("TWILIO_ACCOUNT_SID must be set for acceptance tests")
	}
	if v := os.Getenv("TWILIO_AUTH_TOKEN"); v == "" {
		t.Fatal("TWILIO_AUTH_TOKEN must be set for acceptance tests")
	}
}
