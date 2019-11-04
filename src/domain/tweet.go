package domain

import (
	"fmt"
	"time"
)

type Tweet struct {
	Id   int
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet {
	t := time.Now()
	return &Tweet{-1, user, text, &t}
}

func (t *Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t *Tweet) String() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}
