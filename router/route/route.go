// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
    //"net/http"

    "devops/core"
    "devops/core/router/middleware"
    //"github.com/gin-gonic/gin"
    //"devops/core/util/gopath"
    //"devops/core/util/gostring"
    reqApi "devops/core/router/route/api"
    "devops/core/router/user"
)

func RegisterRoute() {
    api := core.App.Gin.Group("/api", middleware.ApiPriv())
    {
        api.POST(reqApi.LOGIN, user.Login)
        api.POST(reqApi.LOGOUT, user.Logout)
        api.GET(reqApi.LOGIN_STATUS, user.LoginStatus)

        api.POST(reqApi.MY_USER_SETTING, user.MyUserSetting)
        api.POST(reqApi.MY_USER_PASSWORD, user.MyUserPassword)

        api.GET(reqApi.USER_ROLE_PRIV_LIST, user.RolePrivList)
        api.POST(reqApi.USER_ROLE_ADD, user.RoleAdd)
        api.POST(reqApi.USER_ROLE_UPDATE, user.RoleUpdate)
        api.GET(reqApi.USER_ROLE_LIST, user.RoleList)
        api.GET(reqApi.USER_ROLE_DETAIL, user.RoleDetail)
        api.POST(reqApi.USER_ROLE_DELETE, user.RoleDelete)
        api.POST(reqApi.USER_ADD, user.UserAdd)
        api.POST(reqApi.USER_UPDATE, user.UserUpdate)
        api.GET(reqApi.USER_LIST, user.UserList)
        api.GET(reqApi.USER_EXISTS, user.UserExists)
        api.GET(reqApi.USER_DETAIL, user.UserDetail)
        api.POST(reqApi.USER_DELETE, user.UserDelete)
    }

    //if core.App.FeServeEnable == 1 {
    //    RegisterFeResource()
    //}
}

//func RegisterFeResource() {
//    parentPath, err := gopath.CurrentParentPath()
//    if err != nil {
//        core.App.Logger.Error("get current path failed, err[%s]", err.Error())
//        return
//    }
//    indexFile := gostring.JoinStrings(parentPath, "/public/index.html")
//    staticPath := gostring.JoinStrings(parentPath, "/public")
//
//    core.App.Gin.StaticFile("/", indexFile)
//    core.App.Gin.Static("/static", staticPath)
//
//    core.App.Gin.LoadHTMLFiles(indexFile)
//    core.App.Gin.NoRoute(func(c *gin.Context) {
//        c.HTML(http.StatusOK, "index.html", nil)
//    })
//}