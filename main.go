package main

import (
	"net/http"
)

func main() {
	ml := &StubMarkLogic{`<ArticleType><id>lol</id><isLeaf>true</isLeaf></ArticleType>`}
	dao := NewMLDAO(ml)
	handler := NewArticleTypeHandler(dao)

	http.Handle("/", handler)

	http.ListenAndServe(":5000", nil)
}
