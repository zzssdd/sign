// Code generated by hertz generator.

package api

import (
	"github.com/cloudwego/hertz/pkg/app"
	"sign/biz/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _addactivityMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JwtMW()}
}

func _chooseMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JwtMW()}
}

func _choosesubmitMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JwtMW()}
}

func _creategroupMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JwtMW()}
}

func _joinMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JwtMW()}
}

func _signMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JwtMW()}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _addprizeMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JwtMW()}
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _signmonthMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JwtMW()}
}
