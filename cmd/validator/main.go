package main

import (
	"context"
	"fmt"
	"net/http"

	sendpulse "github.com/dimuska139/sendpulse-sdk-go/v7"
)

func main() {
	config := &sendpulse.Config{
		UserID: "",
		Secret: "",
	}
	client := sendpulse.NewClient(http.DefaultClient, config)

	emails := []*sendpulse.EmailToAdd{
		{
			Email:     "test@test.com",
			Variables: map[string]interface{}{"age": 21, "weight": 99},
		},
	}

	ctx := context.Background()
	if err := client.Emails.MailingLists.SingleOptIn(ctx, 1266208, emails); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*emails[0])
}
