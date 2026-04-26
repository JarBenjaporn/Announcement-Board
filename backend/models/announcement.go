package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// struct ของ database
type Announcement struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	Pinned    bool      `json:"pinned" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate สร้าง ID ใหม่ด้วย uuid ก่อนที่จะสร้าง record ใน database มันเรียกใช้ grom hook อัตโนมัติ
func (a *Announcement) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.NewString()
	return nil
}

type CreateRequest struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	Author string `json:"author" binding:"required"`
	Pinned bool   `json:"pinned"`
}

