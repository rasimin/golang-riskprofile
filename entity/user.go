package entity

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`                                                    // ID pengguna sebagai primary key
	Name      string    `gorm:"type:varchar;not null" json:"name" binding:"required"`                    // Nama pengguna (wajib diisi)
	Email     string    `gorm:"type:varchar;uniqueIndex;not null" json:"email" binding:"required,email"` // Email pengguna (wajib diisi, harus unik)
	CreatedAt time.Time `json:"created_at"`                                                              // Waktu pembuatan pengguna
	UpdatedAt time.Time `json:"updated_at"`                                                              // Waktu pembaruan terakhir pengguna
}

type SubmissionUserView struct {
	SubmissionID      int             `json:"submission_id"`
	UserID            int             `json:"user_id"`
	UserName          string          `json:"user_name"`
	UserEmail         string          `json:"user_email"`
	Answers           json.RawMessage `json:"answers,omitempty"`
	RiskScore         int             `json:"risk_score,omitempty"`
	RiskCategory      string          `json:"risk_category,omitempty"`
	Definition        string          `json:"Definition"`
	SubmissionCreated time.Time       `json:"submission_created_at"`
	SubmissionUpdated time.Time       `json:"submission_updated_at"`
}
