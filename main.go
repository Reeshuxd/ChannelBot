//   Channel Bot
//   Copyright (C) 2021 Reeshuxd (@itsreeshu)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.


package main


import (
	"fmt"
	"net/http"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/Reeshuxd/ChannelBot/plugs"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)


func main() {
	b, err := gotgbot.NewBot(TOKEN, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		fmt.Sprintf("Error: %v", err.Error())
	}
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				fmt.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			Panic:       nil,
			ErrorLog:    nil,
			MaxRoutines: 0,
		},
	})
	dispatcher := updater.Dispatcher

	dispatcher.AddHandler(handlers.NewCommand("start", plugs.Start))
	dispatcher.AddHandler(handlers.NewCommand("unban", plugs.UnBan))
        dispatcher.AddHandler(handlers.NewMessage(message.All, plugs.Detector))
	erro := updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if erro != nil {
        fmt.Println("Failed:" + err.Error())
    }
    fmt.Printf("Succesfully Started Bot!")
    updater.Idle()
}
