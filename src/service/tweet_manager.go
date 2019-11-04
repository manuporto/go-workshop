package service

import (
	"fmt"

	"github.com/manuporto/go-workshop/src/domain"
)

type TweetManager struct {
	tweets       []*domain.Tweet
	tweetsByUser map[string][]*domain.Tweet
}

func NewTweetManager() *TweetManager {
	return &TweetManager{make([]*domain.Tweet, 0), make(map[string][]*domain.Tweet)}
}

// PublishTweet is a function to publish a tweet
func (tm *TweetManager) PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return -1, fmt.Errorf("user can not be empty")
	} else if tweet.Text == "" {
		return -1, fmt.Errorf("text can not be empty")
	}
	tweet.Id = len(tm.tweets)
	tm.tweets = append(tm.tweets, tweet)
	tm.tweetsByUser[tweet.User] = append(tm.tweetsByUser[tweet.User], tweet)
	return tweet.Id, nil
}

// GetTweet returns the last tweet
func (tm *TweetManager) GetTweet() *domain.Tweet {
	if len(tm.tweets) == 0 {
		return nil
	}
	return tm.tweets[len(tm.tweets)-1]
}

// GetTweets returns all tweets
func (tm *TweetManager) GetTweets() []*domain.Tweet {
	return tm.tweets
}

func (tm TweetManager) GetTweetById(id int) *domain.Tweet {
	return tm.tweets[id]
}

func (tm *TweetManager) CountTweetsByUser(user string) int {
	if tweetsByUser, exists := tm.tweetsByUser[user]; exists {
		return len(tweetsByUser)
	}
	return 0
}

func (tm *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	return tm.tweetsByUser[user]
}
