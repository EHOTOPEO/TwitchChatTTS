package main

import (
	"github.com/asticode/go-texttospeech/texttospeech"
	twitch "github.com/gempir/go-twitch-irc/v3"
)

func main() {
	client := twitch.NewAnonymousClient()
	tts := texttospeech.NewTextToSpeech()
	client.Join("Channel name")
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		tts.Say(message.Message)
	})

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
