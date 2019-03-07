package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDAO(t *testing.T) {
	ml := &StubMarkLogic{`<ArticleType><id>lol</id><isLeaf>true</isLeaf></ArticleType>`}

	dao := NewMLDAO(ml)

	got := dao.GetArticleType("whatever")
	want := ArticleType{ID: "lol", IsLeaf: true}

	if !cmp.Equal(got, want) {
		t.Errorf(`got "%v", want "%v"`, got, want)
	}
}
