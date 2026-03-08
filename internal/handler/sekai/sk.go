package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	sekairegion "Haruki-Command-Parser/internal/sekai_region"
	"fmt"
	"strconv"
	"strings"
)

func (sekaiHandlers) SKPredictHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/pjsk sk predict", "/pjsk board predict", "/sk预测", "/榜线预测", "/skp",
		}},
		PrefixArgs: []string{"", "wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs()) + ctx.PrefixArg
			// TODO: 迁移 extract_wl_event + wl 单榜拦截 + compose_skp_image 回图逻辑
			return nil, fmt.Errorf("TODO: 榜线预测未实现，query=%q", args)
		},
	}
}

func (sekaiHandlers) SKLineHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/pjsk sk line", "/pjsk board line", "/sk线", "/skl", "/榜线",
		}},
		PrefixArgs: []string{"", "wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs()) + ctx.PrefixArg
			full := false
			if strings.Contains(args, "full") || strings.Contains(args, "all") || strings.Contains(args, "全部") {
				full = true
			}
			// TODO: 迁移 extract_wl_event + full 参数 + compose_skl_image 回图逻辑
			return nil, fmt.Errorf("TODO: 整体榜线未实现，query=%q, full=%t", args, full)
		},
	}
}

func (sekaiHandlers) SKSpeedHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/pjsk sk speed", "/pjsk board speed", "/时速", "/sks", "/skv", "/sk时速",
		}},
		PrefixArgs: []string{"", "wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs()) + ctx.PrefixArg
			minutes := 60
			if v, err := strconv.Atoi(strings.TrimSpace(args)); err == nil {
				minutes = v
			}
			// TODO: 迁移 extract_wl_event + compose_sks_image(unit='h') 回图逻辑
			return nil, fmt.Errorf("TODO: 时速查询未实现，minutes=%d, query=%q", minutes, args)
		},
	}
}

func (sekaiHandlers) SKDailySpeedHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/pjsk sk daily speed", "/pjsk board daily speed", "/日速", "/skds", "/skdv", "/sk日速",
		}},
		PrefixArgs: []string{"", "wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs()) + ctx.PrefixArg
			days := 1
			if v, err := strconv.Atoi(strings.TrimSpace(args)); err == nil {
				days = v
			}
			// TODO: 迁移 extract_wl_event + compose_sks_image(unit='d') 回图逻辑
			return nil, fmt.Errorf("TODO: 日速查询未实现，days=%d, query=%q", days, args)
		},
	}
}

func (sekaiHandlers) SKBoardHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/pjsk sk board", "/pjsk board", "/sk",
		}},
		PrefixArgs: []string{"", "wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs()) + ctx.PrefixArg
			// TODO: 迁移 extract_wl_event + parse_sk_query_params + compose_sk_image 回图逻辑
			return nil, fmt.Errorf("TODO: 指定榜线查询未实现，query=%q", args)
		},
	}
}

func (sekaiHandlers) CFHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/cf", "/查房", "/pjsk查房",
		}},
		PrefixArgs: []string{"", "wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs()) + ctx.PrefixArg
			// TODO: 迁移 extract_wl_event + parse_sk_query_params + compose_cf_image 回图逻辑
			return nil, fmt.Errorf("TODO: 查房未实现，query=%q", args)
		},
	}
}

func (sekaiHandlers) CSBHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/csb", "/查水表", "/pjsk查水表", "/停车时间",
		}},
		PrefixArgs: []string{"", "wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs()) + ctx.PrefixArg
			// TODO: 迁移 extract_wl_event + parse_sk_query_params + compose_csb_image 回图逻辑
			return nil, fmt.Errorf("TODO: 查水表未实现，query=%q", args)
		},
	}
}

func (sekaiHandlers) PlayerTraceHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/ptr", "/玩家追踪", "/pjsk玩家追踪",
		}},
		PrefixArgs: []string{"", "wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs()) + ctx.PrefixArg
			// TODO: 迁移 extract_wl_event + parse_sk_query_params + compose_player_trace_image 回图逻辑
			return nil, fmt.Errorf("TODO: 玩家追踪未实现，query=%q", args)
		},
	}
}

func (sekaiHandlers) RankTraceHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/rtr", "/skt", "/追踪", "/pjsk追踪", "/sklt", "/sktl", "/分数线追踪", "/pjsk分数线追踪",
		}},
		PrefixArgs: []string{"", "wl"},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs()) + ctx.PrefixArg
			// TODO: 迁移 extract_wl_event + get_rank_from_text + rank 校验 + compose_rank_trace_image 回图逻辑
			return nil, fmt.Errorf("TODO: 分数线追踪未实现，query=%q", args)
		},
	}
}

func (sekaiHandlers) WinratePredictHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{Commands: []string{
			"/pjsk winrate predict", "/胜率预测", "/5v5预测", "/胜率", "/5v5胜率", "/预测胜率", "/预测5v5",
		}},
		Regions: []*sekairegion.SekaiRegion{sekairegion.GetRegionById("jp")},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			// TODO: 迁移 compose_winrate_predict_image 回图逻辑
			return nil, fmt.Errorf("TODO: 5v5胜率预测未实现")
		},
	}
}
