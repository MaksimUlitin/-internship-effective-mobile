package models

import (
	"time"
)

type Group struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `json:"name" gorm:"unique;index"`
	Songs     []Song     `json:"songs,omitempty" gorm:"foreignKey:GroupId"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type Song struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	GroupId     uint       `json:"group_id" gorm:"index"`
	GroupName   string     `json:"group_name" gorm:"index"`
	Title       string     `json:"song"`
	ReleaseDate time.Time  `json:"release_date"`
	Text        string     `json:"text"`
	Link        string     `json:"link"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type SongDetail struct {
	GroupName   string `json:"group_name"`
	SongName    string `json:"song_name"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type AddNewSong struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongUpdate struct {
	GroupName   *string `json:"group_name"`
	Song        *string `json:"song,omitempty"`
	ReleaseDate *string `json:"release_date,omitempty"`
	Text        *string `json:"text,omitempty"`
	Link        *string `json:"link,omitempty"`
}
