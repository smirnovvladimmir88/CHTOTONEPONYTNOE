// В файле commands/GetAdmins.go

package commands

import (
	"feedbackBot/src/db"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func GetAdmins(b *gotgbot.Bot, ctx *ext.Context) error {
	// Получаем всех администраторов из базы данных
	admins, err := db.GetAllAdmins()
	if err != nil {
		return err
	}

	// Формируем сообщение с администраторами
	message := "Список администраторов:\n"
	for _, admin := range admins {
		message += fmt.Sprintf("- %s (@%s)\n", admin.FirstName, admin.Username)
	}

	// Отправляем сообщение с администраторами
	_, err = ctx.EffectiveMessage.Reply(
		b,
		message,
		&gotgbot.SendMessageOpts{
			ParseMode: "HTML",
		},
	)
	return err
}
