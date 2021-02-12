package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["domio-api/controllers:AdminOnBoardingController"] = append(beego.GlobalControllerRouter["domio-api/controllers:AdminOnBoardingController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:AdminOnBoardingController"] = append(beego.GlobalControllerRouter["domio-api/controllers:AdminOnBoardingController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:LoginController"] = append(beego.GlobalControllerRouter["domio-api/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["domio-api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["domio-api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["domio-api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["domio-api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["domio-api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:OnBoardingController"] = append(beego.GlobalControllerRouter["domio-api/controllers:OnBoardingController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:OnBoardingController"] = append(beego.GlobalControllerRouter["domio-api/controllers:OnBoardingController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:PreferenceController"] = append(beego.GlobalControllerRouter["domio-api/controllers:PreferenceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:PreferenceController"] = append(beego.GlobalControllerRouter["domio-api/controllers:PreferenceController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:PreferenceController"] = append(beego.GlobalControllerRouter["domio-api/controllers:PreferenceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:PreferenceController"] = append(beego.GlobalControllerRouter["domio-api/controllers:PreferenceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:RegisterController"] = append(beego.GlobalControllerRouter["domio-api/controllers:RegisterController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:UserDetailController"] = append(beego.GlobalControllerRouter["domio-api/controllers:UserDetailController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:UserDetailController"] = append(beego.GlobalControllerRouter["domio-api/controllers:UserDetailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:UserDetailController"] = append(beego.GlobalControllerRouter["domio-api/controllers:UserDetailController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:UserPreferenceController"] = append(beego.GlobalControllerRouter["domio-api/controllers:UserPreferenceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:UserPreferenceController"] = append(beego.GlobalControllerRouter["domio-api/controllers:UserPreferenceController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:UserPreferenceController"] = append(beego.GlobalControllerRouter["domio-api/controllers:UserPreferenceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:UserStoryController"] = append(beego.GlobalControllerRouter["domio-api/controllers:UserStoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:UserStoryController"] = append(beego.GlobalControllerRouter["domio-api/controllers:UserStoryController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["domio-api/controllers:UserStoryController"] = append(beego.GlobalControllerRouter["domio-api/controllers:UserStoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
