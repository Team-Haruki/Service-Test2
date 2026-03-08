package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	sekairegion "Haruki-Command-Parser/internal/sekai_region"
	"fmt"
	"strconv"
	"strings"
)

func (sekaiHandlers) ScoreControlHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk score",
				"/控分",
			},
		},
		Regions:    []*sekairegion.SekaiRegion{sekairegion.GetRegionById("jp")},
		PrefixArgs: []string{"wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			parts := strings.SplitN(args, " ", 2)
			if len(parts) == 0 {
				return nil, fmt.Errorf("使用方式:\n%s 活动pt 歌曲名(可选)", ctx.OriginalTriggerCmd)
			}
			targetPT, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil || targetPT <= 0 {
				return nil, fmt.Errorf("使用方式:\n%s 活动pt 歌曲名(可选)", ctx.OriginalTriggerCmd)
			}
			query := ""
			if len(parts) > 1 {
				query = strings.TrimSpace(parts[1])
			}
			// TODO: 迁移 search_music(query) + DEFAULT_MID 逻辑
			// TODO: 迁移 compose_score_control_image(ctx, targetPT, mid, ctx.PrefixArg == "wl") 回图逻辑
			return nil, fmt.Errorf("TODO: 控分未实现，target_pt=%d, query=%q, wl=%t", targetPT, query, ctx.PrefixArg == "wl")
		},
	}
}

func (sekaiHandlers) CustomRoomScoreControlHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk custom room score", "/custom room score",
				"/自定义房间控分", "/自定义房控分", "/自定义控分",
			},
		},
		Regions: []*sekairegion.SekaiRegion{sekairegion.GetRegionById("jp")},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			targetPT, err := strconv.Atoi(args)
			if err != nil || targetPT <= 0 {
				return nil, fmt.Errorf("使用方式: %s 目标PT", ctx.OriginalTriggerCmd)
			}
			// TODO: 迁移 compose_custom_room_score_control_image(ctx, targetPT) 回图逻辑
			return nil, fmt.Errorf("TODO: 自定义房间控分未实现，target_pt=%d", targetPT)
		},
	}
}

func (sekaiHandlers) MusicMetaHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk music meta", "/music meta",
				"/歌曲meta",
			},
			Priority: 1,
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			segments := strings.Split(strings.ReplaceAll(args, "/", "|"), "|")
			clean := make([]string, 0, len(segments))
			for _, seg := range segments {
				seg = strings.TrimSpace(seg)
				if seg != "" {
					clean = append(clean, seg)
				}
			}
			if len(clean) == 0 {
				return nil, fmt.Errorf("请至少提供一个歌曲ID或名称")
			}
			if len(clean) > 3 {
				return nil, fmt.Errorf("一次最多进行3首歌曲的比较")
			}
			// TODO: 迁移 search_music(use_emb=false) + compose_music_meta_image 回图逻辑
			return nil, fmt.Errorf("TODO: 歌曲meta未实现，segments=%v", clean)
		},
	}
}

func (sekaiHandlers) MusicBoardHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk music board", "/music board",
				"/歌曲排行", "/歌曲比较", "/歌曲排名",
			},
			Priority: 1,
		},
		Regions: []*sekairegion.SekaiRegion{sekairegion.GetRegionById("jp")},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.ToLower(strings.TrimSpace(ctx.GetArgs()))
			// TODO: 迁移分页/类型/排序/策略参数解析 + compose_music_board_image 回图逻辑
			return nil, fmt.Errorf("TODO: 歌曲排行未实现，query=%q", args)
		},
	}
}
