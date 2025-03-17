package models

import (
	"time"
)

type Artist struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `json:"name" gorm:"unique;index"`
	Songs     []Song     `json:"songs,omitempty" gorm:"foreignKey:ArtistID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type Song struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	ArtistID    uint       `json:"artist_id" gorm:"index"`
	Title       string     `json:"song"`
	ReleaseDate time.Time  `json:"release_date"`
	Text        string     `json:"text"`
	Link        string     `json:"link"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type SongDetail struct {
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type AddNewSong struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongUpdate struct {
	Title       *string    `json:"song,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Text        *string    `json:"text,omitempty"`
	Link        *string    `json:"link,omitempty"`
}
