package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rndd/discord-bot/gitlab"
	"github.com/rndd/discord-bot/tracker"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	fmt.Println("Token: " + Token)
	if Token == "" {
		fmt.Println("use -t to set auth token")
		return
	}
	dg, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	response := findSomethingAndPrepareResponse(m.Content)
	if len(response) > 0 {
		fmt.Printf("C Id: %s. M id: %s. Msg: %s\n", m.ChannelID, m.Message.ID, m.Content)
		_, err := s.ChannelMessageSend(m.ChannelID, response)

		if err != nil {
			fmt.Println("Error in sending: ", err)
			return
		}
	}

}

func findSomethingAndPrepareResponse(msg string) string {
	tasksIds := tracker.FindTasksIds(msg)
	mrIds := gitlab.FindMRIds(msg)
	response := ""
	if len(tasksIds) > 0 || len(mrIds) > 0 {
		response += gitlab.GetTextForMr(mrIds) + "\n" + tracker.GetTextForTasks(tasksIds)
	}

	return response
}
