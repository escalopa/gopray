package handler

import (
	"context"
	"log"

	"github.com/escalopa/gopray/pkg/language"

	"github.com/SakoDroid/telego"
	objs "github.com/SakoDroid/telego/objects"

	"github.com/escalopa/gopray/telegram/internal/application"
)

type Handler struct {
	c context.Context
	b *telego.Bot
	u *application.UseCase

	botOwner int                 // Bot owner's ID.
	userCtx  map[int]userContext // userID => latest user context

	userScript map[int]*language.Script // userID => scripts for the user.
}

func New(ctx context.Context, b *telego.Bot, ownerID int, u *application.UseCase) *Handler {
	return &Handler{
		b: b,
		u: u,
		c: ctx,

		botOwner: ownerID,

		userCtx:    make(map[int]userContext),
		userScript: make(map[int]*language.Script),
	}
}

func (h *Handler) Run() error {
	err := h.register()
	if err != nil {
		return err
	}
	go h.notifySubscribers() // Notify subscriber about the prayer times.
	return nil
}

func (h *Handler) register() error {
	var err error
	err = h.b.AddHandler("/start", h.defaultWrapper(h.Start), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/help", h.defaultWrapper(h.Help), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/subscribe", h.defaultWrapper(h.Subscribe), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/unsubscribe", h.defaultWrapper(h.Unsubscribe), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/today", h.defaultWrapper(h.GetPrayers), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/date", h.defaultWrapper(h.GetPrayersByDate), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/lang", h.defaultWrapper(h.SetLang), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/feedback", h.defaultWrapper(h.Feedback), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/bug", h.defaultWrapper(h.Bug), "all")
	if err != nil {
		return err
	}

	//////////////////////////
	///// Admin Commands /////
	//////////////////////////

	err = h.b.AddHandler("/respond", h.defaultWrapper(h.Respond, h.admin), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/subs", h.defaultWrapper(h.GetSubscribers, h.admin), "all")
	if err != nil {
		return err
	}
	err = h.b.AddHandler("/sall", h.defaultWrapper(h.SendAll, h.admin), "all")
	if err != nil {
		return err
	}
	return nil
}

// simpleSend sends a simple message to the chat with the given chatID & text and replyTo.
func (h *Handler) simpleSend(chatID int, text string, replyTo int) (messageID int) {
	r, err := h.b.SendMessage(chatID, text, "", replyTo, false, false)
	if err != nil {
		log.Printf("failed to send message on simpleSend: %s", err)
		return 0
	}
	return r.Result.MessageId
}

// cancelOperation checks if the message is /cancel and sends a response.
// Returns true if the message is /cancel.
func (h *Handler) cancelOperation(message, response string, chatID int) bool {
	if message == "/cancel" {
		h.simpleSend(chatID, response, 0)
		return true
	}
	return false
}

// deleteMessage deletes the message with the given chatID & messageID.
// If error occurs, it will be logged.
func (h *Handler) deleteMessage(chatID, messageID int) {
	editor := h.b.GetMsgEditor(chatID)
	_, err := editor.DeleteMessage(messageID)
	if err != nil {
		log.Printf("failed to delete message: %s", err)
		return
	}
}

type wrapperFunc func(func(*objs.Update)) func(*objs.Update)

func (h *Handler) defaultWrapper(command func(u *objs.Update), extra ...wrapperFunc) func(u *objs.Update) {
	wrappers := []wrapperFunc{
		h.contextWrapper,
		h.userWrapper,
		h.scriptWrapper,
	}
	wrappers = append(wrappers, extra...)

	// Wrap the command with the wrappers.
	finalWrapper := func(u *objs.Update) {
		command(u)
	}
	for i := len(wrappers) - 1; i >= 0; i-- {
		prevWrapper := finalWrapper
		finalWrapper = wrappers[i](func(uu *objs.Update) { prevWrapper(uu) })
	}
	return finalWrapper
}
