package sekai

import (
	"Haruki-Command-Parser/cmd/handler"
	"log"
	"reflect"
)

type sekaiHandlers struct{}

func RegisterSekaiCommandHandler() {
	handlersVal := reflect.ValueOf(sekaiHandlers{})
	handlersTyp := handlersVal.Type()
	configTyp := reflect.TypeOf(handler.SekaiCommandHandlerConfig{})
	for i := 0; i < handlersVal.NumMethod(); i++ {
		methodVal := handlersVal.Method(i)
		methodTyp := methodVal.Type()
		methodName := handlersTyp.Method(i).Name
		//
		if methodTyp.NumIn() == 0 &&
			methodTyp.NumOut() == 1 &&
			methodTyp.Out(0) == configTyp {
			log.Printf("注册指令解析器：%s\n", methodName)
			results := methodVal.Call(nil)
			conf := results[0].Interface().(handler.SekaiCommandHandlerConfig)
			handler.RegisterCommandHandler(conf.ToSekaiCommandHandler())
		}
	}
}
