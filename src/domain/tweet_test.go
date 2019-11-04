package domain_test

import (
	"testing"

	"github.com/manuporto/go-workshop/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestCanGetAPrintableTweet(t *testing.T) {
	tweet := domain.NewTweet("grupoesfera", "This is my tweet")

	text := tweet.String()

	expectedText := "@grupoesfera: This is my tweet"

	assert.Equal(t, expectedText, text)
}
