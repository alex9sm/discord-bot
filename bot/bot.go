package bot

import (
	"discord-bot/config"
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session
const prefix string = "!"

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

	u, err := goBot.User("@me")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	
	BotID = u.ID

	goBot.AddHandler(messageHandler)
	goBot.AddHandler(Ready)

	err = goBot.Open()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

	fmt.Println("Bot is running")
}

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateListeningStatus("alexc.00")
	fmt.Println(session.State.User.Username, "is online")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID{
		return
	}

	args := strings.Split(m.Content, " ")

	if args[0] != prefix {
		return
	}

	if args[1] == "commands"{
		s.ChannelMessageSend(m.ChannelID, "! cat")
	}
	if args[1] == "Commands"{
		s.ChannelMessageSend(m.ChannelID, "! cat")
	}
	if args[1] == "cat"{
		cat := []string{
			"https://tenor.com/view/madison-r-cute-kitty-cute-cat-gif-25737639",
			"https://tenor.com/view/kitty-kitten-lazy-pet-me-gif-7245759",
			"https://tenor.com/view/cats-cat-cute-cat-cat-kissing-cat-hugging-gif-15786228053719683335",
			"https://tenor.com/view/cats-kittens-kitten-cat-cute-cat-gif-7309452916863172056",
			"https://tenor.com/view/good-night-cute-kitty-kitty-cat-cat-gif-23989694",
			"https://tenor.com/view/hello-hi-cute-kitten-cat-gif-19842520",
			"https://tenor.com/view/meow-kitty-happy-cat-kitty-cat-gif-4532088786446233986",
			"https://tenor.com/view/cute-kitty-best-kitty-alex-cute-pp-kitty-omg-yay-cute-kitty-munchkin-kitten-gif-15917800",
			"https://tenor.com/view/kitty-cat-gif-25023317",
			"https://tenor.com/view/kitty-kiss-kitty-kitty-kisses-camera-cat-cat-kisses-gif-5171242563612607703",
			"https://tenor.com/view/kitty-gif-25023523",
			"https://tenor.com/view/helpurseif-helpurself-vereena-reena-cat-gif-17056035502126816704",
			"https://tenor.com/view/kitty-gif-25023528",
			"https://tenor.com/view/on-my-way-cat-run-cat-on-my-way-cat-cat-on-my-way-gif-26471384",
			"https://tenor.com/view/hollyweencandy-cat-cat-drink-milk-cat-eating-cat-drinking-gif-16978701194275077898",
			"https://tenor.com/view/cat-kitty-gif-23521077",
		}
			
		selection := rand.Intn(len(cat))

		s.ChannelMessageSend(m.ChannelID, cat[selection])
	}
}