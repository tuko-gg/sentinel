package db

import (
	"time"
)

type Label struct {
	LabelId      string `gorm:"primaryKey"`
	UserId       string `gorm:"index"`
	Name         string
	Description  string
	Action       string
	CreatedTime  time.Time
	ModifiedTime time.Time
}

type LabelAction struct {
	LabelId     string `gorm:"primaryKey"`
	ItemId      string `gorm:"primaryKey;index"`
	UserId      string `gorm:"index"`
	CreatedTime time.Time

	Label Label
}
