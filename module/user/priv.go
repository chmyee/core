// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
    "devops/core/util/goslice"
    reqApi "devops/core/router/route/api"
)

var apiToPrivMap = map[string][]int{}

func init() {
    for priv, apiList := range privToApiMap{
        for _, api := range apiList {
            privMap, _ := apiToPrivMap[api]
            apiToPrivMap[api] = append(privMap, priv)
        }
    }
}

func CheckHavePriv(api string, priv []int) bool {
    privMap, exists := apiToPrivMap[api]
    if !exists {
        return false
    }
    return len(goslice.SliceIntersectInt(privMap, priv)) > 0
}

func PrivIn(privCode int, privList []int) bool {
    return goslice.InSliceInt(privCode, privList)
}

const (
    USER_ROLE_VIEW = 3001 // 查看角色
    USER_ROLE_NEW  = 3002 // 新增角色
    USER_ROLE_EDIT = 3003 // 编辑角色
    USER_ROLE_DEL  = 3004 // 删除角色
    USER_VIEW = 3101 // 查看用户
    USER_NEW  = 3102 // 新增用户
    USER_EDIT = 3103 // 编辑用户
    USER_DEL  = 3104 // 删除用户
)

var privToApiMap = map[int][]string{
    USER_ROLE_VIEW: []string{
        reqApi.USER_ROLE_LIST,
    },
    USER_ROLE_NEW: []string{
        reqApi.USER_ROLE_ADD,
        reqApi.USER_ROLE_PRIV_LIST,
    },
    USER_ROLE_EDIT: []string{
        reqApi.USER_ROLE_DETAIL,
        reqApi.USER_ROLE_UPDATE,
        reqApi.USER_ROLE_PRIV_LIST,
    },
    USER_ROLE_DEL: []string{
        reqApi.USER_ROLE_DELETE,
    },
    USER_VIEW: []string{
        reqApi.USER_LIST,
    },
    USER_NEW: []string{
        reqApi.USER_ADD,
        reqApi.USER_ROLE_LIST,
        reqApi.USER_EXISTS,
    },
    USER_EDIT: []string{
        reqApi.USER_ROLE_LIST,
        reqApi.USER_DETAIL,
        reqApi.USER_UPDATE,
        reqApi.USER_EXISTS,
    },
    USER_DEL: []string{
        reqApi.USER_DELETE,
    },
}

type PrivItem struct {
    Label   string  `json:"label"`
    Value   int     `json:"value"`
}

type PrivGroup struct {
    Label   string      `json:"label"`
    Items   []PrivItem  `json:"items"`
}

var PrivList = []PrivGroup {
    privUser,
}

var privUser = PrivGroup {
    Label: "用户",
    Items: []PrivItem {
        PrivItem{
            Label: "角色-查看",
            Value: USER_ROLE_VIEW,
        },
        PrivItem{
            Label: "角色-新增",
            Value: USER_ROLE_NEW,
        },
        PrivItem{
            Label: "角色-删除",
            Value: USER_ROLE_DEL,
        },
        PrivItem{
            Label: "角色-编辑",
            Value: USER_ROLE_EDIT,
        },
        PrivItem{
            Label: "用户-查看",
            Value: USER_VIEW,
        },
        PrivItem{
            Label: "用户-新增",
            Value: USER_NEW,
        },
        PrivItem{
            Label: "用户-编辑",
            Value: USER_EDIT,
        },
        PrivItem{
            Label: "用户-删除",
            Value: USER_DEL,
        },
    },
}