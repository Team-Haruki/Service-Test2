package sekai

import (
	"Haruki-Command-Parser/cmd/handler"
	"context"
	"log"
	"testing"
)

func TestRegisterCommandHandler(t *testing.T) {

	RegisterSekaiCommandHandler()
	log.Println(handler.GetAllCommandHandlers())
	v := handler.Dispatch(context.Background(), handler.Event{
		Message: "/cn查谱 虾",
	})
	log.Println(v)
}
