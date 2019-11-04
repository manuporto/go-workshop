package service

import (
	"github.com/manuporto/go-workshop/src/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func validateTweets(t *testing.T,  expectedTweet *domain.Tweet, actualTweet *domain.Tweet) {
	assert.Equal(t, expectedTweet.User, actualTweet.User, "Tweets users should be equal")
	assert.Equal(t, expectedTweet.Text, actualTweet.Text, "Tweets texts should be equal")
	assert.NotNil(t, actualTweet.Date, "Tweet should have a date")
}

func TestGetTweetWithNoPreviousTweetsReturnsNil(t *testing.T) {
	tweetManager := NewTweetManager()
	publishedTweet := tweetManager.GetTweet()

	assert.Nil(t, publishedTweet, "Published tweet should be nil since there are not any published tweets")
}

func TestPublishedTweetIsSaved(t *testing.T) {
	tweetManager := NewTweetManager()
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)

	_, _ = tweetManager.PublishTweet(tweet)
	publishedTweet := tweetManager.GetTweet()

	validateTweets(t, tweet, publishedTweet)
}

func TestWhenTryingToPublishTweetWithoutUserAnErrorIsReturned(t *testing.T) {
	tweetManager := NewTweetManager()
	text := "Tweet without user"
	tweet := domain.NewTweet("", text)

	_, err := tweetManager.PublishTweet(tweet)

	assert.Error(t, err, "Should throw an error")
	assert.EqualError(t, err, "user can not be empty", "Error messages should match")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	tweetManager := NewTweetManager()
	user := "user"
	tweet := domain.NewTweet(user, "")

	_, err := tweetManager.PublishTweet(tweet)

	assert.Error(t, err, "Should throw an error")
	assert.EqualError(t, err, "text can not be empty", "Error messages should match")
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	tweetManager := NewTweetManager()
	firstTweet := domain.NewTweet("firstUser", "firstTweetText")
	secondTweet := domain.NewTweet("secondUser", "secondTweetText")

	_, _ = tweetManager.PublishTweet(firstTweet)
	_, _ = tweetManager.PublishTweet(secondTweet)

	publishedTweets := tweetManager.GetTweets()

	assert.Len(t, publishedTweets, 2, "There should be two published tweets")
	validateTweets(t, firstTweet, publishedTweets[0])
	validateTweets(t, secondTweet, publishedTweets[1])
}

func TestCanRetrieveTweetById(t *testing.T) {
	tweetManager := NewTweetManager()
	tweet := domain.NewTweet("grupoesfera", "Tweet text")

	id, _ := tweetManager.PublishTweet(tweet)

	publishedTweet := tweetManager.GetTweetById(id)

	assert.NotNil(t, id, "Id should not be nil")
	validateTweets(t, tweet, publishedTweet)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	tweetManager := NewTweetManager()
	user := "grupoesfera"
	firstTweet := domain.NewTweet(user, "First tweet")
	secondTweet := domain.NewTweet(user, "Second tweet")
	thirdTweet := domain.NewTweet(user, "Third tweet")

	_, _ = tweetManager.PublishTweet(firstTweet)
	_, _ = tweetManager.PublishTweet(secondTweet)
	_, _ = tweetManager.PublishTweet(thirdTweet)

	count := tweetManager.CountTweetsByUser(user)

	assert.Equal(t, 3, count)
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	tweetManager := NewTweetManager()
	user := "grupoesfera"
	anotherUser := "meli"
	text := "Tweet text"
	secondText := "Second tweet text"
	tweet := domain.NewTweet(user, text)
	secondTweet := domain.NewTweet(user, secondText)
	thirdTweet := domain.NewTweet(anotherUser, secondText)

	_, _ = tweetManager.PublishTweet(tweet)
	_, _ = tweetManager.PublishTweet(secondTweet)
	_, _ = tweetManager.PublishTweet(thirdTweet)

	tweets := tweetManager.GetTweetsByUser(user)

	assert.Len(t, tweets, 2, "User should have two tweets published")
	validateTweets(t, tweet, tweets[0])
	validateTweets(t, secondTweet, tweets[1])
}