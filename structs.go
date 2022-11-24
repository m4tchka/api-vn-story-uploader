package main

type E struct {
	Key   string
	Value interface{}
}
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

type SceneObj struct {
	Id    int           `json:"id"`
	Scene []DialogueObj `json:"scene"`
}
type DialogueObj struct {
	Name       string      `json:"Name" bson:"Name"`
	Dialogue   string      `json:"Dialogue" bson:"Dialogue"`
	Background string      `json:"Background,omitempty" bson:"Background,omitempty"`
	Question   string      `json:"Question,omitempty" bson:"Question,omitempty"`
	Options    []OptionObj `json:"Options,omitempty" bson:"Options,omitempty"`
}
type OptionObj struct {
	Text       string `json:"Text" bson:"Text,omitempty"`
	Next       int    `json:"Next" bson:"Next,omitempty"`
	LuckChange int    `json:"LuckChange" bson:"LuckChange,omitempty"`
	MinLuck    int    `json:"MinLuck" bson:"MinLuck,omitempty"`
}
