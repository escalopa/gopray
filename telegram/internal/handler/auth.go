package handler

import (
	objs "github.com/SakoDroid/telego/objects"
)

// admin is a wrapper for admin commands to check if the user is the bot owner
func (h *Handler) admin(command func(*objs.Update)) func(*objs.Update) {
	return func(u *objs.Update) {
		user, err := h.u.GetUser(h.userCtx[u.Message.Chat.Id].ctx, u.Message.Chat.Id)
		if err == nil && user.IsAdmin {
			command(u)
		} else {
			h.Help(u)
		}
	}
}
