package models

type Bug struct{
	ID uint `json:"id" gorm:"primarykey"`
	Title string `json:"title"`
	Description string `json:"description"`
	Serverity string `json:"Severity"`
	Status string  `json:"status"`
	Assignee string `json:"assignee"`

}