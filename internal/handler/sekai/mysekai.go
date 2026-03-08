package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"fmt"
	"strconv"
	"strings"
)

func (sekaiHandlers) MysekaiResHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk mysekai res", "/msr", "/msmap", "/msa",
				"/mysekai 资源",
			},
			Priority: 1,
		},
		// TODO: refer 中限制 regions=get_regions(RegionAttributes.MYSEKAI)
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			showHarvested := strings.Contains(args, "all")
			checkTime := !strings.Contains(args, "force")
			// TODO: 迁移 bd_msr_sub 群限制校验 + compose_mysekai_res_image 回图逻辑
			return nil, fmt.Errorf("TODO: mysekai资源查询未实现，show_harvested=%t, check_time=%t", showHarvested, checkTime)
		},
	}
}

func (sekaiHandlers) MysekaiBlueprintHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk mysekai blueprint", "/mysekai blueprint",
				"/msb", "/mysekai 蓝图",
			},
		},
		// TODO: refer 中限制 regions=get_regions(RegionAttributes.MYSEKAI)
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			showID := strings.Contains(args, "id")
			showAllTalks := strings.Contains(args, "all")
			// TODO: 迁移 unit/cid 解析 + compose_mysekai_fixture_list_image/compose_mysekai_talk_list_image 回图逻辑
			return nil, fmt.Errorf("TODO: mysekai蓝图查询未实现，args=%q, show_id=%t, show_all_talks=%t", args, showID, showAllTalks)
		},
	}
}

func (sekaiHandlers) MysekaiFurnitureHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk mysekai furniture", "/pjsk mysekai fixture",
				"/msf", "/mysekai 家具", "/家具列表",
			},
		},
		// TODO: refer 中限制 regions=get_regions(RegionAttributes.MYSEKAI)
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			parts := strings.Fields(args)
			fids := make([]int, 0, len(parts))
			allInts := len(parts) > 0
			for _, p := range parts {
				v, err := strconv.Atoi(p)
				if err != nil {
					allInts = false
					break
				}
				fids = append(fids, v)
			}
			// TODO: 迁移按 fids 查询详情分支 + 列表分支（含 unit/cid/all）逻辑
			return nil, fmt.Errorf("TODO: mysekai家具查询未实现，all_ints=%t, fids=%v, args=%q", allInts, fids, args)
		},
	}
}

func (sekaiHandlers) MysekaiPhotoHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk mysekai photo", "/pjsk mysekai picture",
				"/msp", "/mysekai 照片",
			},
		},
		// TODO: refer 中限制 regions=get_regions(RegionAttributes.MYSEKAI)
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			seq, err := strconv.Atoi(args)
			if err != nil {
				return nil, fmt.Errorf("请输入正确的照片编号（从1或-1开始）")
			}
			// TODO: 迁移群限制校验 + get_mysekai_photo_and_time + 回图逻辑
			return nil, fmt.Errorf("TODO: mysekai照片下载未实现，seq=%d", seq)
		},
	}
}

func (sekaiHandlers) CheckMysekaiDataHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk check mysekai data",
				"/pjsk烤森抓包数据", "/pjsk烤森抓包", "/烤森抓包", "/烤森抓包数据",
				"/msd",
			},
		},
		// TODO: refer 中限制 regions=get_regions(RegionAttributes.MYSEKAI)
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			// TODO: 迁移 at用户解析 + get_player_bind_id + get_mysekai_upload_time + 文本组装逻辑
			return nil, fmt.Errorf("TODO: 烤森抓包状态未实现，user_id=%s", ctx.GetUserId())
		},
	}
}

func (sekaiHandlers) MysekaiGateHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk mysekai gate", "/msg", "/msgate",
			},
		},
		// TODO: refer 中限制 regions=get_regions(RegionAttributes.MYSEKAI)
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			qidAll := strings.Contains(args, "all")
			// TODO: 迁移 unit->gate_id 解析 + compose_mysekai_door_upgrade_image 回图逻辑
			return nil, fmt.Errorf("TODO: 烤森门升级查询未实现，args=%q, all=%t", args, qidAll)
		},
	}
}

func (sekaiHandlers) MysekaiMusicRecordHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/pjsk mysekai musicrecord", "/msm", "/mss", "/mssong",
			},
		},
		// TODO: refer 中限制 regions=get_regions(RegionAttributes.MYSEKAI)
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			showID := strings.Contains(args, "id")
			// TODO: 迁移群限制校验 + compose_mysekai_musicrecord_image 回图逻辑
			return nil, fmt.Errorf("TODO: 烤森唱片数据查询未实现，show_id=%t", showID)
		},
	}
}

func (sekaiHandlers) MSRChangeBindHandle() SekaiCommandHandler {
	return SekaiCommandHandler{
		CommandHandlerBase: handler.CommandHandlerBase{
			Commands: []string{
				"/msr换绑",
			},
		},
		// TODO: refer 中限制 regions=get_regions(RegionAttributes.BD_MYSEKAI)
		handleFunc: func(ctx SekaiHandlerContext) (interface{}, error) {
			args := strings.TrimSpace(ctx.GetArgs())
			force := strings.Contains(args, "force")
			// TODO: force 需要 superuser 权限
			// TODO: 迁移参数校验 + update_bd_msr_limit_uid 调用
			return nil, fmt.Errorf("TODO: msr换绑未实现，force=%t, args=%q", force, args)
		},
	}
}
