package service

import "goemt/model"
import "goemt/config"
import "database/sql"
import "log"
import "io/ioutil"

// AllNewsList returns all news in db
func GetAllNewsList() []model.News {
	var newsList []model.News

	db, err := sql.Open("sqlite3", config.NewsDb)
	if err != nil {
		log.Println(err)
		return newsList
	}

	defer db.Close()

	rows, err := db.Query("select id,title,abstract,`source`,`date` from news order by `date` desc")

	if err != nil {
		log.Println(err)
		return newsList
	}

	defer rows.Close()

	for rows.Next() {
		var nw model.News
		rows.Scan(&nw.ID, &nw.Title, &nw.Abstract, &nw.Source, &nw.Date)
		newsList = append(newsList, nw)
	}

	return newsList
}

func GetNewsByID(ID string) model.News {
	var news model.News

	db, err := sql.Open("sqlite3", config.NewsDb)
	if err != nil {
		log.Println(err)
		return news
	}

	defer db.Close()

	rows, err := db.Query("select id,title,abstract,`source`,`date` from news where id=? limit 1", ID)

	if err != nil {
		log.Println(err)
		return news
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&news.ID, &news.Title, &news.Abstract, &news.Source, &news.Date)
		news.Content = readContentByID(ID)
	}
	return news
}

func readContentByID(ID string) string {
	dat, err := ioutil.ReadFile(config.NewsRoot + "/" + ID + ".html")

	check(err)

	return string(dat)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetAllProductList() []model.Product {
	var prodList []model.Product

	db, err := sql.Open("sqlite3", config.ProductDb)
	if err != nil {
		log.Println(err)
		return prodList
	}

	defer db.Close()

	rows, err := db.Query("select id,title,subtitle, wcover, hcover from product order by `id` asc")

	if err != nil {
		log.Println(err)
		return prodList
	}

	defer rows.Close()

	for rows.Next() {
		var nw model.Product
		rows.Scan(&nw.ID, &nw.Title, &nw.Subtitle, &nw.Wcover, &nw.Hcover)
		prodList = append(prodList, nw)
	}

	return prodList
}

func GetProductByID(ID string) model.Product {
	var product model.Product

	db, err := sql.Open("sqlite3", config.ProductDb)
	if err != nil {
		log.Println(err)
		return product
	}

	defer db.Close()

	rows, err := db.Query("select id,title,subtitle,wcover,hcover from product where id=? limit 1", ID)

	if err != nil {
		log.Println(err)
		return product
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&product.ID, &product.Title, &product.Subtitle, &product.Wcover, &product.Hcover)
	}
	return product
}

func GetAllSolutionList() []model.Solution {
	var prodList []model.Solution

	db, err := sql.Open("sqlite3", config.SolutionDb)
	if err != nil {
		log.Println(err)
		return prodList
	}

	defer db.Close()

	rows, err := db.Query("select id,title,subtitle, wcover, hcover from solution order by `id` asc")

	if err != nil {
		log.Println(err)
		return prodList
	}

	defer rows.Close()

	for rows.Next() {
		var nw model.Solution
		rows.Scan(&nw.ID, &nw.Title, &nw.Subtitle, &nw.Wcover, &nw.Hcover)
		prodList = append(prodList, nw)
	}

	return prodList
}

func GetSolutionByID(ID string) model.Solution {
	var solution model.Solution

	db, err := sql.Open("sqlite3", config.SolutionDb)
	if err != nil {
		log.Println(err)
		return solution
	}

	defer db.Close()

	rows, err := db.Query("select id,title,subtitle,wcover,hcover from solution where id=? limit 1", ID)

	if err != nil {
		log.Println(err)
		return solution
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&solution.ID, &solution.Title, &solution.Subtitle, &solution.Wcover, &solution.Hcover)
	}
	return solution
}
