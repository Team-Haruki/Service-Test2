package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"errors"
	"fmt"
	"strings"
)

func (sekaiHandlers) ChallengeInfoHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk challenge info", "/pjsk_challenge_info",
				"/挑战信息", "/挑战详情", "/挑战进度", "/挑战一览", "/每日挑战",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			// TODO: 迁移 compose_challenge_live_detail_image(ctx, ctx.user_id)
			// TODO: 迁移 get_image_cq(..., low_quality=true) 回复流程
			return nil, fmt.Errorf("TODO: 挑战信息未实现，user_id=%s", ctx.GetUserId())
		},
	}
}

func (sekaiHandlers) PowerBonusInfoHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk power bonus info", "/pjsk_power_bonus_info",
				"/加成信息", "/加成详情", "/加成进度", "/加成一览", "/角色加成",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			// TODO: 迁移 compose_power_bonus_detail_image(ctx, ctx.user_id)
			// TODO: 迁移 get_image_cq(..., low_quality=true) 回复流程
			return nil, fmt.Errorf("TODO: 加成信息未实现，user_id=%s", ctx.GetUserId())
		},
	}
}

func (sekaiHandlers) AreaItemHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk area item", "/area item",
				"/区域道具", "/区域道具升级", "/区域道具升级材料",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())

			helpText := fmt.Sprintf(
				"可用参数: 团名/角色名/属性/树/花\n"+
					"加上\"all\"可以查询所有级别材料，不加则查询你的账号的升级情况，示例：\n"+
					"\"%s 树\" 所有树\n"+
					"\"%s miku\" miku的道具\n"+
					"\"%s 25h\" 25的SEKAI里的所有区域道具\n"+
					"\"%s miku all\" miku的道具所有等级",
				ctx.OriginalTriggerCmd,
				ctx.OriginalTriggerCmd,
				ctx.OriginalTriggerCmd,
				ctx.OriginalTriggerCmd,
			)

			useAll := false
			for _, keyword := range []string{"all", "full"} {
				if strings.Contains(args, keyword) {
					useAll = true
					args = strings.TrimSpace(strings.ReplaceAll(args, keyword, ""))
					break
				}
			}

			tree := false
			if strings.Contains(args, "树") {
				tree = true
				args = strings.TrimSpace(strings.ReplaceAll(args, "树", ""))
			}

			flower := false
			if strings.Contains(args, "花") {
				flower = true
				args = strings.TrimSpace(strings.ReplaceAll(args, "花", ""))
			}

			// TODO: 迁移 extract_unit(args) / extract_card_attr(args) / get_cid_by_nickname(args)
			// TODO: 迁移 AreaItemFilter + compose_area_item_upgrade_materials_image(ctx, qid, filter)
			// TODO: 与 refer 一致：当 unit/attr/cid/tree/flower 全为空时返回 HELP_TEXT
			if !tree && !flower && args == "" {
				return nil, errors.New(helpText)
			}

			return nil, fmt.Errorf(
				"TODO: 区域道具查询未实现，args=%q, use_all=%t, tree=%t, flower=%t",
				args, useAll, tree, flower,
			)
		},
	}
}

func (sekaiHandlers) BondsHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk bonds", "/pjsk bond",
				"/羁绊", "/羁绊等级", "/角色羁绊", "/羁绊信息",
				"/牵绊等级", "/牵绊", "/角色牵绊", "/牵绊信息",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			if args != "" {
				// TODO: 迁移 get_cid_by_nickname(args)，并在无效角色时返回“请指定其中一个角色名称”
			}

			// TODO: 迁移 compose_bonds_image(ctx, ctx.user_id, cid)
			// TODO: 迁移 get_image_cq(..., low_quality=true) 回复流程
			return nil, fmt.Errorf("TODO: 羁绊查询未实现，query=%q, user_id=%s", args, ctx.GetUserId())
		},
	}
}

func (sekaiHandlers) LeaderCountHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk leader count",
				"/队长次数", "/角色次数", "/队长游玩次数", "/角色游玩次数",
			},
		},
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			// TODO: 迁移 compose_leader_count_image(ctx, ctx.user_id)
			// TODO: 迁移 get_image_cq(..., low_quality=true) 回复流程
			return nil, fmt.Errorf("TODO: 队长次数查询未实现，user_id=%s", ctx.GetUserId())
		},
	}
}
