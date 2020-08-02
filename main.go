package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/crisgarner/gas-bot/oracle"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
var (
	log = logrus.WithField("prefix", "main")
)

// Message handler implements several functions which will
// be in charge of responding to discord messages the bot
// observes.
type messageHandler struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// The init function runs on package initialization, helping us setup
// some useful globals such as a logging formatter.
func init() {
	formatter := new(prefixed.TextFormatter)
	formatter.TimestampFormat = "2020-01-01 07:12:23"
	formatter.FullTimestamp = true
	logrus.SetFormatter(formatter)
}

func main() {



	token := getEnv("DISCORD_TOKEN")
	if token == "" {
		log.Fatalf("Expected DISCORD_TOKEN env var, provided none")
	}
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		log.Fatalf("Could not initialize discord session: %v", err)
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	// We initialize a new context with a cancelation function, useful
	// for cleanup of every possible goroutine on SIGTERM.
	ctx, cancel := context.WithCancel(context.Background())
	handler := &messageHandler{
		ctx:    ctx,
		cancel: cancel,
	}

	// Go gas bot habdler
	dg.AddHandler(handler.gasHandler)

	// Wait here until SIGTERM or another interruption signal is received.
	log.Println("Bot is now running, press ctrl-c to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session and cancel the global
	// context gracefully.
	cancel()
	if err := dg.Close(); err != nil {
		log.Fatalf("Could not gracefully stop discord session: %v", err)
	}

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func (mh *messageHandler) gasHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}
	commandPrefix := "!gas"
	if !strings.Contains(m.Content, commandPrefix) {
		return
	}

	client, err := oracle.New(context.Background(), &oracle.Config{
		URL:     getEnv("CLIENT_URL"),
		Address: common.HexToAddress(getEnv("ORACLE_ADDRESS")),
	})

	if err != nil {
		log.Fatalf("Could not initialize client: %v", err)
	}
	//Gas from Chainlink
	gas, err := client.GetGas()
	if err != nil {
		log.Errorf("Could not reach Chainlink Oracle: %v", err)
		return
	}
	toGwei(gas)
	//Gas from GETH
	//gasPrice, err := client.Client().SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	res := fmt.Sprintf("⛽️ Gas Price is %d Gwei", gas)
	if _, err := s.ChannelMessageSend(m.ChannelID, res); err != nil {
		log.Errorf("Could not send message over channel: %v", err)
	}
}

func toGwei(val *big.Int) *big.Int{
	return val.Div(val , big.NewInt(1000000000))
}