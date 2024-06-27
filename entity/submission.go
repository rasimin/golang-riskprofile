package entity

import (
	"encoding/json"
	"time"
)

type Submission struct {
	ID           int             `gorm:"primaryKey" json:"id"`                       // ID submission sebagai primary key
	UserID       int             `gorm:"not null" json:"user_id" binding:"required"` // ID pengguna yang mengirimkan submission
	Answers      json.RawMessage `gorm:"type:jsonb;not null" json:"answers" binding:"required"`
	RiskScore    int             `json:"risk_score"`                            // Skor risiko
	RiskCategory string          `gorm:"type:varchar(50)" json:"risk_category"` // Kategori risiko
	CreatedAt    time.Time       `json:"created_at"`                            // Waktu pembuatan submission
	UpdatedAt    time.Time       `json:"updated_at"`                            // Waktu pembaruan terakhir submission
	// User         User      `gorm:"foreignKey:UserID"`                                     // Relasi ke entitas User
}

type Post_Submission struct {
	ID           int       `gorm:"primaryKey" json:"id"`                       // ID submission sebagai primary key
	UserID       int       `gorm:"not null" json:"user_id" binding:"required"` // ID pengguna yang mengirimkan submission
	Answers      []Answer  `gorm:"type:jsonb;not null" json:"answers" binding:"required"`
	RiskScore    int       `json:"risk_score"`                            // Skor risiko
	RiskCategory string    `gorm:"type:varchar(50)" json:"risk_category"` // Kategori risiko
	CreatedAt    time.Time `json:"created_at"`                            // Waktu pembuatan submission
	UpdatedAt    time.Time `json:"updated_at"`                            // Waktu pembaruan terakhir submission
	// User         User      `gorm:"foreignKey:UserID"`                                     // Relasi ke entitas User
}

type Answer struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
}
