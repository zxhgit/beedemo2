package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beedemo2/controllers:CarController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:CarController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:carId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:CarController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:CarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:CpcAdController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:CpcAdController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:UserController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:UserController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:UserController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:UserController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:UserController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:UserController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:UserController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beedemo2/controllers:UserController"] = append(beego.GlobalControllerRouter["beedemo2/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout1",
			Router: `/logout1`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
