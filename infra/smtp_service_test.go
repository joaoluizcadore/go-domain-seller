package infra_test

import (
	"testing"

	"github.com/joaoluizcadore/domain-seller/infra"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}

	err := infra.SendNotification("Test message")

	assert.Nil(t, err)
}
