package routers

import (
	"bee_project/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.AggregateController{}, "get:Index")
	beego.Router("/vote", &controllers.VoteController{}, "post:Vote")
	beego.Router("/favourite", &controllers.MakeFavouriteController{}, "post:Favourite")

}
