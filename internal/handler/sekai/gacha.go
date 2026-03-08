package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"fmt"
	"strconv"
	"strings"
)

func (sekaiHandlers) GachaHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk gacha", "/卡池列表", "/卡池一览", "/卡池", "/查卡池",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			// TODO: 迁移 parse_search_gacha_args(ctx, args) 单卡池分支
			// TODO: 迁移 parse_search_multiple_gacha_args(ctx, args) 多卡池分支
			// TODO: 迁移 compose_gacha_detail_image / compose_gacha_list_image 回图逻辑
			return nil, fmt.Errorf("TODO: 卡池查询未实现，query=%q", args)
		},
	}
}

func (sekaiHandlers) GachaRecordHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk gacha record", "/抽卡记录", "/抽卡历史",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			specGIDs := make([]int, 0)
			if args != "" {
				for _, part := range strings.Fields(args) {
					gid, err := strconv.Atoi(part)
					if err != nil {
						return nil, fmt.Errorf("卡池ID参数错误: %s", part)
					}
					specGIDs = append(specGIDs, gid)
				}
			}

			// TODO: 校验 spec_gids 是否存在（ctx.md.gachas.find_by_id）
			// TODO: 迁移 compose_gacha_record_image(ctx, ctx.user_id, spec_gids) 回图逻辑
			return nil, fmt.Errorf("TODO: 抽卡记录未实现，spec_gids=%v", specGIDs)
		},
	}
}
