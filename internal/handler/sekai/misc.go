package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	sekairegion "Haruki-Command-Parser/internal/sekai_region"
	"errors"
	"fmt"
	"strings"
)

func (sekaiHandlers) UpdateHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk update", "/pjsk refresh", "/pjsk更新",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			// TODO: 迁移 RegionMasterDbManager.get(ctx.region).update() 逻辑
			return nil, fmt.Errorf("TODO: 更新查询未实现，region=%v", ctx.Region)
		},
	}
}

func (sekaiHandlers) NgWordHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk ng", "/pjsk ngword", "/pjsk ng word",
				"/pjsk屏蔽词", "/pjsk屏蔽", "/pjsk敏感", "/pjsk敏感词",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			text := strings.TrimSpace(ctx.GetArgs())
			if text == "" {
				return nil, errors.New("请输入要查询的文本")
			}
			// TODO: 迁移 ctx.md.ng_words.get() + 屏蔽词检测逻辑
			return nil, fmt.Errorf("TODO: 屏蔽词检测未实现，text=%q", text)
		},
	}
}

func (sekaiHandlers) UploadHelpHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/抓包帮助", "/抓包", "/pjsk upload help",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			// TODO: 迁移读取 upload_help.txt 并回复逻辑
			return nil, errors.New("TODO: 抓包帮助未实现")
		},
	}
}

func (sekaiHandlers) ExtractCardHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/提取卡牌",
			},
		},
		Regions: []*sekairegion.SekaiRegion{sekairegion.GetRegionById("jp")},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			// TODO: 迁移 CardExtractor 初始化、回复图片提取、网格渲染与回图逻辑
			return nil, errors.New("TODO: 提取卡牌未实现")
		},
	}
}

func (sekaiHandlers) CharaBirthdayHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk chara birthday", "/角色生日", "/生日",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			// TODO: 迁移角色生日信息汇总、参数解析、绘图与回图逻辑
			return nil, fmt.Errorf("TODO: 角色生日未实现，query=%q", args)
		},
	}
}

func (sekaiHandlers) HeyiweiHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjskb30", "/pjskdetail", "/b30", "/b39", "/pjskb39", "/pjsk b30", "/pjsk b39", "/pjsk detail",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			return "何意味", nil
		},
	}
}
