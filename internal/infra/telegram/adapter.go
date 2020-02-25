package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
	"scheduler/internal/configs"
	"scheduler/internal/domain"
)

type Adapter interface {
	Start() error
}

type adapter struct {
	logger *zap.Logger
	config *configs.Config
	service domain.Service
	bot *tgbotapi.BotAPI
}

func (a *adapter) Start() error {
	bot, err := tgbotapi.NewBotAPI(a.config.TgToken)
	if err != nil {
		return err
	}


	a.bot = bot
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		// check message
		if update.CallbackQuery != nil {
			go a.routeForCallback(&update)
		}
		if update.Message != nil {
			go a.route(&update)
		}

	}

	return nil
}

func NewAdapter(logger *zap.Logger, config *configs.Config, service domain.Service) Adapter {
	return &adapter{
		logger: logger,
		config: config,
		service: service,
	}
}



