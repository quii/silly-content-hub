package main

import "encoding/xml"

type MarkLogic interface {
	GetArticleTypeXML() []byte
}

type DAO interface {
	GetArticleType(id string) ArticleType
}

type MLDAO struct {
	ml MarkLogic
}

func NewMLDAO(ml MarkLogic) *MLDAO {
	return &MLDAO{ml: ml}
}

type articleTypeXML struct {
	XMLName xml.Name `xml:"ArticleType"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id"`
	IsLeaf  bool     `xml:"isLeaf"`
}

func (m *MLDAO) GetArticleType(id string) ArticleType {
	artXML := m.ml.GetArticleTypeXML()
	var art articleTypeXML
	xml.Unmarshal(artXML, &art)

	return ArticleType{
		art.ID,
		art.IsLeaf,
	}
}
