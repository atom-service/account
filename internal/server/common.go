package server

import (
	"context"

	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
)

type PermissionRule struct {
	Action string
	Key    string
	Value  string
}

func (p *PermissionRule) ExactMatch(action string,  key string, value string) bool {
	return p.Action == action && p.Key == key && p.Value == value
}

func ResolvePermissionFormIncomeContext(ctx context.Context, resourceName string, handler func(rule PermissionRule) bool) bool {
	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil || authData.Secret == nil {
		return false
	}

	// 查询当前用户所有的权限信息并交给 ruleHandler 进行判断处理
	userRoleResult, err := model.Permission.QueryUserResourceSummary(ctx, model.UserResourceSummarySelector{UserID: authData.User.ID})
	if err != nil || len(userRoleResult) == 0 {
		return false
	}

	for _, role := range userRoleResult {
		// 特殊的万能之神，拥有一切权限
		if role.Name == "all" {
			return true
		}

		permission := PermissionRule{}
		permission.Action = role.Action

		if role.Name == resourceName {
			for _, rule := range role.Rules {
				permission.Key = rule.Key
				permission.Value = rule.Value
				if check := handler(permission); check == true {
					return true
				}
			}
		}
	}

	return false
}
