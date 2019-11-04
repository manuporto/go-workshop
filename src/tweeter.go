package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/manuporto/go-workshop/src/domain"
	"github.com/manuporto/go-workshop/src/service"
)

func main() {
	tweeterManager := service.NewTweetManager()
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")
			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			id, _ := tweeterManager.PublishTweet(domain.NewTweet(user, text))

			c.Println("Tweet with id:", id, "sent")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tweeterManager.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweeterManager.GetTweets()

			c.Println(tweets)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows tweet with :id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the tweet id: ")
			id, _ := strconv.Atoi(c.ReadLine())
			tweets := tweeterManager.GetTweetById(id)

			c.Println(tweets)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsOfUser",
		Help: "Counts the number of tweets of :user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user: ")
			user := c.ReadLine()
			count := tweeterManager.CountTweetsByUser(user)

			c.Println(count)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsOfUser",
		Help: "Shows the tweets of :user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user: ")
			user := c.ReadLine()
			tweets := tweeterManager.GetTweetsByUser(user)

			c.Println(tweets)

			return
		},
	})

	shell.Run()

}
