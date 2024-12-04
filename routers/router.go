package routers

import (
	"WebSitie/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/index")
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/home", &controllers.MainController{}, "get:Home") // "get:Home"
	beego.Router("/about", &controllers.MainController{}, "get:About")
	beego.Router("/bloghome", &controllers.MainController{}, "get:Bloghome")
	beego.Router("/blogpost", &controllers.MainController{}, "get:BlogPost")
	beego.Router("/contact", &controllers.MainController{}, "get:Contact")
	beego.Router("/faq", &controllers.MainController{}, "get:Faq")
	beego.Router("/portfolioitem", &controllers.MainController{}, "get:PortfolioItem")
	beego.Router("/portfoliooverview", &controllers.MainController{}, "get:PortfolioOverview")
	beego.Router("/pricing", &controllers.MainController{}, "get:Pricing")
}
