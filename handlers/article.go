package handlers

import (
	"fmt"
	"github.com/david-sorm/goblog/article"
	"github.com/david-sorm/goblog/globals"
	"net/http"
	"strings"
)

type ArticleView struct {
	BlogName string
	Article  article.Article
	RootURL  string
}

func HandleArticle(rw http.ResponseWriter, req *http.Request) {
	// split the uri (example: /articles/1 )
	split := strings.Split(req.RequestURI, "/")

	// make sure there are 3 splits
	if len(split) != 3 {
		Handle404(rw, req)
		return
	}

	// make sure article with the ID exists
	article, exists := globals.Cfg.ArticleStore.GetArticleByID(split[2])
	if !exists {
		Handle404(rw, req)
		return
	}

	// respond
	articleView := ArticleView{
		BlogName: globals.Cfg.BlogName,
		Article:  article,
		RootURL:  "//" + req.Host + "/",
	}
	if err := globals.Templates[globals.TemplateArticle].Execute(rw, articleView); err != nil {
		fmt.Println("Error while parsing template:", err.Error())
	}
}
