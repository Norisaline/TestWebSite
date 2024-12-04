package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.TplName = "index.html"
}

func (c *MainController) Home() {
	c.TplName = "home.html"
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
