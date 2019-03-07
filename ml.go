package main

type StubMarkLogic struct {
	Payload string
}

func (s *StubMarkLogic) GetArticleTypeXML() []byte {
	return []byte(s.Payload)
}
