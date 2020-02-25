package telegram
import (
tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
"strconv"
)

func (a *adapter) route(u *tgbotapi.Update) {

	switch u.Message.Command() {
	case "start":
		a.startHandler(u)
	case "gr":
		a.groupScheduleHandler(u)
	}
}

func (a *adapter) routeForCallback(u *tgbotapi.Update) {
	groupId := a.service.GetGroupByChatId(u.CallbackQuery.Message.Chat.ID)
	if groupId != nil {

		day, _ := strconv.Atoi(u.CallbackQuery.Data)
		a.groupDayScheduleHandler(u, *groupId, day)
	}
}
