package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type Part struct {
	StarIndex int `json:"star_index"`
	GetStarTs int `json:"get_star_ts"`
}
type Day struct {
	Part1 Part `json:"1"`
	Part2 Part `json:"2"`
}
type Member struct {
	LocalScore         int            `json:"local_score"`
	Stars              int            `json:"stars"`
	CompletionDayLevel map[string]Day `json:"completion_day_level"`
	Id                 int            `json:"id"`
	GlobalScore        int            `json:"global_score"`
	Name               string         `json:"name"`
	LastStarTs         int            `json:"last_star_ts"`
}
type Event struct {
	Members map[string]Member `json:"members"`
	OwnerID int               `json:"owner_id"`
	Event   string            `json:"event"`
}

func GetAocMessages(lastTs int64) []string {
	//Make a request to the API
	messages := make([]string, 0)

	httpClient := http.Client{
		Timeout: time.Second * 20,
	}

	req, err := http.NewRequest(http.MethodGet, os.Getenv("LEADERBOARD_URL"), nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return messages
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: os.Getenv("AOC_SESSION_COOKIE")})
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Failed to get response")
		return messages
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Failed to get 200 response")
		return messages
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read body")
		return messages
	}
	var members Event
	err = json.Unmarshal(body, &members)
	if err != nil {
		fmt.Println("Failed to unmarshal body")
		return messages
	}
	for _, member := range members.Members {
		for dayN, day := range member.CompletionDayLevel {

			if int64(day.Part1.GetStarTs) > lastTs {
				message := fmt.Sprintf("%v completed D%v P1 at %v", member.Name, dayN, time.Unix(int64(day.Part1.GetStarTs), 0).Format("02/01 15:04:05"))
				messages = append(messages, message)
			}
			if int64(day.Part2.GetStarTs) > lastTs {
				message := fmt.Sprintf("%v completed D%v P2 at %v", member.Name, dayN, time.Unix(int64(day.Part2.GetStarTs), 0).Format("02/01 15:04:05"))
				messages = append(messages, message)
			}
		}
	}
	return messages
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	GROUP_ID := os.Getenv("GROUP_ID")
	fmt.Println("Starting notifier")

	store.DeviceProps.Os = proto.String("AOC")

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:wha.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "INFO", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	fmt.Println("Starting ticker")
	ticker := time.NewTicker(15 * time.Minute)
	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	group := types.NewJID(GROUP_ID, types.GroupServer)
	me := types.NewJID(os.Getenv("ME"), types.GroupServer)
	lastTs := time.Now().Unix()

	client.SendMessage(context.Background(), me, &waProto.Message{
		Conversation: proto.String(fmt.Sprintf("(%v) Im up!", time.Unix(lastTs, 0).Format("02/01 15:04:05"))),
	})
	
	for {
		select {
		case t := <-ticker.C:
			{
				fmt.Printf("(%v) Checking for new stars since %v\n", t.Format("02/01 15:04:05"), time.Unix(lastTs, 0).Format("02/01 15:04:05"))
				messages := GetAocMessages(lastTs)
				for _, message := range messages {
					fmt.Println(message)
					client.SendMessage(context.Background(), group, &waProto.Message{
						Conversation: proto.String(message),
					})
				}
				lastTs = time.Now().Unix()
			}

		case <-quit:
			client.Disconnect()
			ticker.Stop()
			return
		}
	}

}
