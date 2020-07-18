package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"

	//"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var (
	log    = logrus.New()
	dg     *discordgo.Session
	prefix = "!"
	//db     *pgx.Conn
)

func init() {
	log.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading dotenv file")
	}

	pref := os.Getenv("PREFIX")
	if pref != "" {
		prefix = pref
		log.Info("Using prefix \"" + pref + "\"")
	} else {
		log.Warn("No PREFIX env variable declared! Using default prefix \"!\"")
	}

	//db, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Failed to connect to database")
	// }
}

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageCreateEvent)

	dg.AddHandler(readyEvent)
	dg.AddHandler(guildJoinEvent)

	if err := dg.Open(); err != nil {
		log.Fatal("Error opening connection")
		return
	}

	//Prevent it from closing
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-sc

	defer dg.Close()
}
