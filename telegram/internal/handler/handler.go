package handler

import (
	"context"
	"log"

	bt "github.com/SakoDroid/telego"
	objs "github.com/SakoDroid/telego/objects"
	"github.com/escalopa/gopray/telegram/internal/application"
)

type Handler struct {
	b  *bt.Bot
	ac *application.UseCase
	c  context.Context
}

func New(b *bt.Bot, ac *application.UseCase, ctx context.Context) *Handler {
	return &Handler{
		b:  b,
		ac: ac,
		c:  ctx,
	}
}

func (h *Handler) Register() {
	// h.b.AddHandler("/start", h.Start, "all")
	h.b.AddHandler("/help", h.Help, "all")
	h.b.AddHandler("/subscribe", h.Subscribe, "all")
	h.b.AddHandler("/unsubscribe", h.Unsubscribe, "all")
	h.b.AddHandler("/prayers", h.GetPrayers, "all")
	h.b.AddHandler("/prayersdate", h.Getprayersdate, "all")
	h.b.AddHandler("/lang", h.SetLang, "all")
	h.b.AddHandler("/feedback", h.Feedback, "all")
	h.b.AddHandler("/bug", h.Bug, "all")
}

func (h *Handler) Help(u *objs.Update) {
	h.b.SendMessage(u.Message.Chat.Id, `
	Asalamu alaykum, I am a kazan prayers time, I can help you know prayers time anytime to always pray on time 🙏.
	
	Available commands are below: 👇	

	<b>Prayers</b>
	/prayers - Get prayers for today ⏰
	/prayersdate - Get prayers for a specific date 📅
	/subscribe - Subscribe to daily prayers notification 🔔
	/unsubscribe - Unsubscribe from daily prayers notification 🔕

	<b>Support</b>
	/help - Show this message 📖   
	/lang - Set bot language  🌐
	/feedback - Send feedback or idea to the bot developers 📩
	/bug - Report a bug to the bot developers 🐞

	`, "HTML", 0, false, false)
}

// SimpleSend sends a simple message
func (bh *Handler) simpleSend(chatID int, text string, replyTo int) {
	_, err := bh.b.SendMessage(chatID, text, "", replyTo, false, false)
	if err != nil {
		log.Println(err)
	}
}
