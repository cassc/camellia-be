package main

import (
	"database/sql"
	"fmt"
	"goemt/config"
	"goemt/model"
	"goemt/service"
	"log"
	"strconv"

	"github.com/kataras/iris"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.Django(config.TemplateRoot, ".tmpl").Reload(true))

	app.StaticWeb("/js", "../fe/js")
	app.StaticWeb("/css", "../fe/css")
	app.StaticWeb("/img", "../fe/img")
	app.StaticWeb("/assets", "../fe/assets")

	app.Get("/", handleGetIndex)
	app.Post("/emt-feedback", handleEmtFeedback)
	app.Get("/page/news-list", handleGetPageNewsList)
	app.Get("/page/news", handleGetPageNews)
	app.Get("/page/products", handleGetPageProducts)
	app.Get("/page/product", handleGetPageProduct)
	app.Get("/page/solutions", handleGetPageSolutions)
	app.Get("/page/solution", handleGetPageSolution)
	app.Get("/page/aboutus", handleGetPageAboutus)
	app.Get("/page/contactus", handleGetPageContactus)

	app.Run(iris.Addr("0.0.0.0:" + config.Port))
}

func attachBaseInfo(ctx iris.Context) {
	solList := service.GetAllSolutionList()
	prodList := service.GetAllProductList()
	ctx.ViewData("SolutionList", solList)
	ctx.ViewData("ProductList", prodList)
}

func handleGetPageProduct(ctx iris.Context) {
	attachBaseInfo(ctx)
	id := ctx.URLParams()["id"]
	product := service.GetProductByID(id)
	ctx.ViewData("product", product)
	ctx.ViewData("Page", "products")
	ctx.ViewData("Title", product.Title+" - __title.simple__")
	ctx.View(fmt.Sprintf("products/%s.tmpl", id))
}

func handleGetPageProducts(ctx iris.Context) {
	attachBaseInfo(ctx)
	ctx.ViewData("Title", "产品中心 - __title.simple__")
	ctx.ViewData("Page", "products")
	ctx.View("products.tmpl")
}

func handleGetPageSolution(ctx iris.Context) {
	attachBaseInfo(ctx)
	id := ctx.URLParams()["id"]
	solution := service.GetSolutionByID(id)
	ctx.ViewData("solution", solution)
	ctx.ViewData("Page", "solutions")
	ctx.ViewData("Title", solution.Title+" - __title.simple__")
	ctx.View(fmt.Sprintf("solutions/%s.tmpl", id))
}

func handleGetPageSolutions(ctx iris.Context) {
	attachBaseInfo(ctx)
	ctx.ViewData("Page", "solutions")
	ctx.ViewData("Title", "解决方案 - __title.simple__")
	ctx.View("solutions.tmpl")
}

func handleGetPageAboutus(ctx iris.Context) {
	attachBaseInfo(ctx)
	ctx.ViewData("Page", "aboutus")
	ctx.ViewData("Title", "关于我们 - __title.simple__")
	ctx.View("aboutus.tmpl")
}

func handleGetPageContactus(ctx iris.Context) {
	attachBaseInfo(ctx)
	ctx.ViewData("Page", "contactus")
	ctx.ViewData("Title", "联系我们 - __title.simple__")
	ctx.View("contactus.tmpl")
}

func handleGetIndex(ctx iris.Context) {
	attachBaseInfo(ctx)
	newsList := service.GetAllNewsList()

	stop := len(newsList)
	if stop > 4 {
		stop = 4
	}
	ctx.ViewData("Page", "index")
	ctx.ViewData("NewsList", newsList[0:stop])
	ctx.View("index.tmpl")
}

func handleGetPageNewsList(ctx iris.Context) {
	attachBaseInfo(ctx)
	newsList := service.GetAllNewsList()

	ctx.ViewData("NewsList", newsList)
	ctx.ViewData("Page", "news")
	ctx.ViewData("Title", "行业新闻 - __title.simple__")
	ctx.View("news.tmpl")
}

func handleGetPageNews(ctx iris.Context) {
	attachBaseInfo(ctx)
	id := ctx.URLParams()["id"]
	news := service.GetNewsByID(id)
	newsList := service.GetAllNewsList()
	ctx.ViewData("NewsList", newsList)
	ctx.ViewData("news", news)
	ctx.ViewData("Page", "news")
	ctx.ViewData("id", id)
	idInt, err := strconv.Atoi(news.ID)
	if err == nil && idInt > 0 {
		ctx.ViewData("Title", news.Title+" - __title.simple__")
		ctx.View("single-news.tmpl")
	} else {
		ctx.ViewData("Title", "行业新闻 - __title.simple__")
		ctx.View("news.tmpl")
	}

}

func handleEmtFeedback(ctx iris.Context) {
	var feedback model.UserFeedback
	ctx.ReadJSON(&feedback)

	if len(feedback.Message) == 0 {
		ctx.JSON(model.Response{101, "请输入您要反馈的消息!"})
		return
	} else if len(feedback.Email) == 0 && len(feedback.Phone) == 0 {
		ctx.JSON(model.Response{102, "请输入Email或电话号码以便我们能及时联系您!"})
		return
	}

	addFeedback(feedback)
	ctx.JSON(model.Response{0, "success"})
}

func addFeedback(feedback model.UserFeedback) {
	db, err := sql.Open("sqlite3", config.Sdb)
	if err != nil {
		log.Println("open db error")
		return
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := tx.Prepare("insert into feedback(username, message, phone, email) values(?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(feedback.Username, feedback.Message, feedback.Phone, feedback.Email)
	if err != nil {
		log.Println(err)
		return
	}

	tx.Commit()

}
