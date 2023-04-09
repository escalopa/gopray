package handler

import (
	"log"

	objs "github.com/SakoDroid/telego/objects"
	"github.com/escalopa/gopray/pkg/core"
	"github.com/escalopa/gopray/pkg/language"
)

func (h *Handler) userWrapper(command func(u *objs.Update)) func(u *objs.Update) {
	return func(u *objs.Update) {
		if _, err := h.u.GetUser(h.userCtx[u.Message.Chat.Id].ctx, u.Message.Chat.Id); err != nil {
			// Check that the user language is valid & supported.
			userLang := u.Message.From.LanguageCode
			if !language.IsValidLang(userLang) {
				userLang = language.DefaultLang().Short
			}
			// Create a new user.
			err = h.u.StoreUser(h.userCtx[u.Message.Chat.Id].ctx, core.CreateUserParams{
				TelegramID: u.Message.Chat.Id,
				FirstName:  u.Message.Chat.FirstName,
				LastName:   u.Message.Chat.LastName,
				Username:   u.Message.Chat.Username,
				LangCode:   userLang,
			})
			if err != nil {
				log.Printf("Error creating user: %v", err)
				return
			}
			command(u)
		} else {
			command(u)
		}
	}
}
