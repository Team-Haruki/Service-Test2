package handler

import (
	"context"
	"fmt"
	"slices"
	"strings"
)

var commandHandlers []CommandHandler = []CommandHandler{}

func GetAllCommandHandlers() []CommandHandler {
	return commandHandlers
}

// TODO: 返回的结果不论是错误还是别的什么，统一封装，具体怎么样之后再说
func Dispatch(ctx context.Context, event Event) interface{} {
	HandlerContext := &HandlerContext{
		Context:     ctx,
		messageType: event.MessageType,
		message:     event.Message,
		event:       event,
		messageId:   event.MessageId,
		userId:      event.UserId,
		senderName:  event.SenderName,
		groupId:     event.GroupId,
	}
	for _, handler := range commandHandlers {
		for _, cmd := range handler.GetCommands() {
			if strings.HasPrefix(event.Message, cmd) {
				HandlerContext.triggerCmd = cmd
				HandlerContext.argText = strings.TrimSpace(event.Message[len(cmd):])
				result, err := handler.Handle(HandlerContext)
				if err != nil {
					return fmt.Errorf("处理命令时发生错误: %v", err)
				}
				if result == nil {
					continue
				}
				return result
			}
		}
	}
	return nil
}

type CommandHandler interface {
	GetCommands() []string               // 获取该处理器负责的命令列表
	GetPriority() int                    // 处理器的优先级
	GetHelper() string                   // 帮助文本
	Handle(Context) (interface{}, error) // 处理方法，TODO: 返回的结构暂未定义
}

type BaseCommandHandler struct {
	commands   []string
	priority   int
	helper     string
	handleFunc func(Context) (interface{}, error)
}

func (b *BaseCommandHandler) GetCommands() []string {
	return b.commands
}
func (b *BaseCommandHandler) GetPriority() int {
	return b.priority
}
func (b *BaseCommandHandler) GetHelper() string {
	return b.helper
}
func (b *BaseCommandHandler) Handle(ctx Context) (interface{}, error) {
	if b.handleFunc != nil {
		return b.handleFunc(ctx)
	}
	cmdName := "未定义"
	if len(b.commands) > 0 {
		cmdName = b.commands[0]
	}
	return nil, fmt.Errorf("命令处理器 %s 没有处理方法", cmdName)
}

func RegisterCommandHandler(handler CommandHandler) {
	commandHandlers = append(commandHandlers, handler)
	sortCommandHandlers()
}

func sortCommandHandlers() {
	slices.SortFunc(commandHandlers,
		func(a, b CommandHandler) int {
			return a.GetPriority() - b.GetPriority()
		},
	)
}
