package main

func initialiseRoutes() {
	// handle index route
	router.GET("/", showIndexPage)

	// Single article
	router.GET("/article/view/:article_id", getArticle)
}
