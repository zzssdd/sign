// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "sign/biz/handler/sign/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/activity", append(_addactivityMw(), api.AddActivity)...)
	root.GET("/choose", append(_chooseMw(), api.Choose)...)
	root.POST("/choose", append(_choosesubmitMw(), api.ChooseSubmit)...)
	root.POST("/group", append(_creategroupMw(), api.CreateGroup)...)
	root.POST("/join", append(_joinMw(), api.Join)...)
	root.POST("/login", append(_loginMw(), api.Login)...)
	root.POST("/prize", append(_addprizeMw(), api.AddPrize)...)
	root.POST("/register", append(_registerMw(), api.Register)...)
	root.POST("/sign", append(_signMw(), api.Sign)...)
	root.GET("/signMonth", append(_signmonthMw(), api.SignMonth)...)
}
