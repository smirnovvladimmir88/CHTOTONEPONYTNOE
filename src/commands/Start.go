package commands

import (
	"feedbackBot/src/config"
	"feedbackBot/src/rates"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"feedbackBot/src/db"
	"feedbackBot/src/models"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	var err error

	// Check if user is not rate-limited and welcome message is enabled
	if rates.Check(ctx.EffectiveChat.Id, 10) && config.CurrentConfig.Welcome.Enabled {
		// Send welcome message
		_, err = ctx.EffectiveMessage.Reply(
			b,
			config.CurrentConfig.Welcome.Message,
			&gotgbot.SendMessageOpts{ParseMode: "HTML"},
		)
	}

	// Reset MessageSent to false when user sends /start
	var user models.User
	res := db.Connection.Where("user_id = ?", ctx.EffectiveUser.Id).First(&user)
	if res.Error != nil {
		return err
	}
	user.MessageSent = false
	db.Connection.Save(&user)

	return err
}
