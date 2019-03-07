package main

import (
	"encoding/json"
	"net/http"
)

type ArticleTypeHandler struct {
	dao DAO
}

func (a ArticleTypeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	art := a.dao.GetArticleType("whatever")
	encoder := json.NewEncoder(w)
	encoder.Encode(art)
}

func NewArticleTypeHandler(dao DAO) *ArticleTypeHandler {
	return &ArticleTypeHandler{dao: dao}
}
