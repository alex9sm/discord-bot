package main

import(
	"fmt"
	"github.com/alex9sm/discord-bot/bot"
	"github.com/alex9sm/discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}