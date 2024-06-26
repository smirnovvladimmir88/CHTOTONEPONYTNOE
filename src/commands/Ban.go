/*
 * Ban.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"errors"
	"feedbackBot/src/constants"
	"feedbackBot/src/helpers"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Ban(b *gotgbot.Bot, ctx *ext.Context) error {
	// Resolve user
	user, err := helpers.ResolveUser(ctx, b)

	if err != nil || user == nil {
		return err
	}

	err = helpers.BanUser(user)

	if errors.Is(err, constants.UserAlreadyBanned) {
		_, err = ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
	} else if err != nil {
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(
		b,
		fmt.Sprintf("#u%d Был заблокирован.", user.UserId),
		&gotgbot.SendMessageOpts{},
	)

	return err
}
