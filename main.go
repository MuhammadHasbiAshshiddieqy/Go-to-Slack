package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New("xoxb-****") //SLACK TOKEN

  message := "First line \n Second line"
  msgText := slack.NewTextBlockObject("mrkdwn", message, false, false)
  msgSection := slack.NewSectionBlock(msgText, nil, nil)

  msg := slack.MsgOptionBlocks(
       msgSection,
     )

	_, _, _, err := api.SendMessage("#channel-name", msg)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
