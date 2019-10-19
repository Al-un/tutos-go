package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// // Call the HTML method of the Context to render a template
	// c.HTML(
	// 	// Set the HTTP status to 200 (OK)
	// 	http.StatusOK,
	// 	// Use the index.html template
	// 	"index.html",
	// 	// Pass the data that the page uses
	// 	gin.H{
	// 		"title":   "Home Page",
	// 		"payload": articles,
	// 	},
	// )
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")

}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if art, err := getArticleByID(articleID); err == nil {
			// // Call the HTML method of the Context to render a template
			// c.HTML(
			// 	// Set the HTTP status to 200 (OK)
			// 	http.StatusOK,
			// 	// Use the index.html template
			// 	"article.html",
			// 	// Pass the data that the page uses
			// 	gin.H{
			// 		"title":   art.Title,
			// 		"payload": art,
			// 	},
			// )

			render(c, gin.H{
				"title":   "Home Page",
				"payload": art}, "article.html")

		} else {

			render(c, gin.H{
				"title": "Not found",
				"payload": article{
					Title:   "???",
					Content: fmt.Sprintf("Could not find %d", articleID),
				}}, "article.html")

			// c.HTML(
			// 	// Set the HTTP status to 200 (OK)
			// 	http.StatusOK,
			// 	// Use the index.html template
			// 	"article.html",
			// 	// Pass the data that the page uses
			// 	gin.H{
			// 		"title": "Not found",
			// 		"payload": article{
			// 			Title:   "???",
			// 			Content: fmt.Sprintf("Could not find %d", articleID),
			// 		},
			// 	},
			// )

			// // If the article is not found, abort with an error
			// c.AbortWithError(http.StatusNotFound, err)

		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
