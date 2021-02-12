// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to auto generate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"domio-api/controllers"

	"github.com/astaxie/beego"
)

// Beego will support swagger if we use NSRouter

func init() {
	ns := beego.NewNamespace("/v1",
		// Get all and Post Onboarding Image
		beego.NSNamespace("/on-boarding",
			beego.NSInclude(
				&controllers.OnBoardingController{},
			),
		),
		// Only for the Admin to get single, update and delete the Onboarding Image
		beego.NSNamespace("/admin/on-boarding/:id",
			beego.NSInclude(
				&controllers.AdminOnBoardingController{},
			),
		),
		// beego.NSRouter("/on-boarding", &controllers.OnBoardingController{}, "post:Post"),
		beego.NSNamespace("/signup",
			beego.NSInclude(
				&controllers.RegisterController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserDetailController{},
			),
		),
	)
	beego.Router("/v1/preferences/?:id", &controllers.PreferenceController{},)
	beego.Router("/v1/user/preferences/", &controllers.UserPreferenceController{},)
	beego.Router("/v1/user/story/?:id", &controllers.UserStoryController{},)
	beego.AddNamespace(ns)
}
