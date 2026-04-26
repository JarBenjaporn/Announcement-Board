package repositories

import (
	"announcement-board/models"

	"gorm.io/gorm"
)

// คุยกับ database ไดด้โดยตรง

type AnnouncementRepository struct {
	db *gorm.DB
}

type AnnouncementRepositoryInterface interface {
	CreateNewAnnouncement(announcement *models.Announcement) error
	FindAll() ([]models.Announcement, error)
	FindByID(id string) (*models.Announcement, error)
	UpdateAnnouncement(announcement *models.Announcement, req *models.CreateRequest) error
	DeleteAnnouncement(id string) (bool, error)
}

// NewAnnouncementRepository สร้าง instance ของ AnnouncementRepository
func NewAnnouncementRepository(db *gorm.DB) AnnouncementRepositoryInterface {
	return &AnnouncementRepository{db: db}
}

// create ประกาศใหม่
func (r *AnnouncementRepository) CreateNewAnnouncement(announcement *models.Announcement) error {
	return r.db.Create(announcement).Error
}

// ดึงข้อมูลทั้งหมด และ card pinned ขึ้นก่อน เรียงจากใหม่ไปเก่า
func (r *AnnouncementRepository) FindAll() ([]models.Announcement, error) {
	var announcements []models.Announcement
	result := r.db.Order("pinned DESC, created_at DESC").Find(&announcements)
	return announcements, result.Error
}

// FindByID หาจาก ID
func (r *AnnouncementRepository) FindByID(id string) (*models.Announcement, error) {
	var announcement models.Announcement
	result := r.db.First(&announcement, "id = ?", id)
	return &announcement, result.Error
}

// update แก้ไขประกาศ เช่น การ pin ประกาศ
func (r *AnnouncementRepository) UpdateAnnouncement(announcement *models.Announcement, req *models.CreateRequest) error {
	return r.db.Model(announcement).Updates(map[string]interface{}{
		"title":  req.Title,
		"body":   req.Body,
		"author": req.Author,
		"pinned": req.Pinned,
	}).Error
}

// Delete ลบประกาศออก โดยลบจาก ID
func (r *AnnouncementRepository) DeleteAnnouncement(id string) (bool, error) {
	result := r.db.Where("id = ?", id).Delete(&models.Announcement{})
	return result.RowsAffected > 0, result.Error
}

