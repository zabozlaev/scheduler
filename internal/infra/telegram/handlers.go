package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"scheduler/internal/domain"
	"strings"
)

func (a *adapter) startHandler(u *tgbotapi.Update) {

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, startResponse)
	_, err := a.bot.Send(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *adapter) groupScheduleHandler(u *tgbotapi.Update) {
	group := strings.Split(u.Message.Text, " ")[1]
	a.service.SetGroupByChatId(u.Message.Chat.ID, group)
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выбери день")
	addMenuToMessage(&msg)
	_, err := a.bot.Send(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func addMenuToMessage(msg *tgbotapi.MessageConfig) {
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Понедельник", "1"),
			tgbotapi.NewInlineKeyboardButtonData("Вторник", "2"),
			tgbotapi.NewInlineKeyboardButtonData("Среда", "3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Четверг", "4"),
			tgbotapi.NewInlineKeyboardButtonData("Пятница", "5"),
			tgbotapi.NewInlineKeyboardButtonData("Суббота", "6"),
		),
	)
}

func (a *adapter) groupDayScheduleHandler(u *tgbotapi.Update, group string, day int) {
	result, err := a.service.GetDayScheduleForGroup(group, day)
	if err != nil {
		fmt.Println(err)
	}

	resp := formatScheduleDayResponse(result)
	msg := tgbotapi.NewMessage(u.CallbackQuery.Message.Chat.ID, resp)
	addMenuToMessage(&msg)

	_, err = a.bot.Send(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func formatScheduleDayResponse(classes []*domain.Class) string {
	str := ""

	for _, class := range classes {
		str += "\n\n" + class.Time + "\n"
		for _, classMeta := range class.ClassMeta {
			str += "\tПредмет: " + classMeta.Title + "\n"
			str += "\tПреподаватель: " + classMeta.Lecturer + "\n"
			str += "\tАудитория: " + classMeta.ClassRoom + "\n"
			str += "\tЗдание: " + classMeta.Building + "\n"
			str += "\n"
		}
	}

	return str
}


