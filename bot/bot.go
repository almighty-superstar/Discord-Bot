package bot

import(
	"fmt"
	"github.com/akhil/discord-ping/config"
	"github.com/bwmarrin/discordgo"
)

var botID string
var goBot *discordgo.Session

func StartBot(){
	goBot , err := discordgo.New("Bot "+config.Token)
	if err != nil{
		fmt.Println(err.Error())
		return 
	}
	u, err := goBot.User("@me")
	if err != nil{
		fmt.Println(err.Error())
		return 
	}
	botID = u.ID
	
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	
	if err != nil{
		fmt.Println(err.Error())
		return 
	}
	fmt.Println("The bot is currently running smoothly!")
}


func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID == botID{
		return 
	}
	if m.Content=="ping"{
		_,_ =s.ChannelMessageSend(m.ChannelID,"pong")
	}

}