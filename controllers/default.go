package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "kakoito-site.com"
	c.Data["Email"] = "Norisaline@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) index() {
	c.TplName = "index.html"
}

func (c *MainController) Home() {
	c.TplName = "index.html"
}

func (c *MainController) About() {
	c.TplName = "about.html"
}

func (c *MainController) Bloghome() {
	c.TplName = "bloghome.html"
}

func (c *MainController) BlogPost() {
	c.TplName = "blogpost.html"
}

func (c *MainController) Contact() {
	c.TplName = "contact.html"
}

func (c *MainController) Faq() {
	c.TplName = "faq.html"
}

func (c *MainController) PortfolioItem() {
	c.TplName = "portfolioitem.html"
}

func (c *MainController) PortfolioOverview() {
	c.TplName = "portfoliooverview.html"
}

func (c *MainController) Pricing() {
	c.TplName = "pricing.html"
}
