package api

import (
	"fmt"
	"github.com/mspangler/go-flickr/oauth"
)

func Authenticate() {

	consumerKey := ""
	consumerSecret := ""

	// TODO: validate key & secret

	// TODO: determine if we have an accessToken

	c := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://www.flickr.com/services/oauth/request_token",
			AuthorizeTokenUrl: "http://www.flickr.com/services/oauth/authorize",
			AccessTokenUrl:    "http://www.flickr.com/services/oauth/access_token",
		})

	requestToken, url, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		panic(err)
	}

	fmt.Println("(1) Go to: " + url)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		panic(err)
	}

	fmt.Println("Received access token: " + accessToken.Token)
}
