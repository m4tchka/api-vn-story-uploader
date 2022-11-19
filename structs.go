package main

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
	Name       string      `json:"Name"`
	Dialogue   string      `json:"Dialogue"`
	Background string      `json:"Background,omitempty"`
	Question   string      `json:"Question,omitempty"`
	Options    []OptionObj `json:"Options,omitempty"`
}
type OptionObj struct {
	Text       string `json:"Text"`
	Next       int    `json:"Next"`
	LuckChange int    `json:"LuckChange"`
	MinLuck    int    `json:"MinLuck"`
}
