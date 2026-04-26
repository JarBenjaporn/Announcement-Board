package services

import (
	"announcement-board/models"
	"announcement-board/repositories"
	"errors"
)

// business logic ไว้คุยกับ repository เพื่อดึงข้อมูลจาก database มาแปรรูปก่อนส่งไปให้ handler
type AnnouncementService struct {
	repo *repositories.AnnouncementRepository
}

// สร้าง instance ของ AnnouncementService
func NewAnnouncementService(repo *repositories.AnnouncementRepository) *AnnouncementService {
	return &AnnouncementService{repo: repo}
}

// GetAll ดึงประกาศทั้งหมด
func (s *AnnouncementService) GetAll() ([]models.Announcement, error) {
	return s.repo.FindAll()
}

// สร้างประกาศใหม่
func (s *AnnouncementService) CreateAnnouncementService(req *models.CreateRequest) (*models.Announcement, error) {
	announcement := &models.Announcement{
		Title:  req.Title,
		Body:   req.Body,
		Author: req.Author,
		Pinned: req.Pinned,
	}
	err := s.repo.CreateNewAnnouncement(announcement)
	return announcement, err
}

// แก้ไขประกาศ
func (s *AnnouncementService) UpdateAnnouncementService(id string, req *models.CreateRequest) (*models.Announcement, error) {
	announcement, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("announcement not found")
	}
	err = s.repo.UpdateAnnouncement(announcement, req)
	return announcement, err
}	

// ลบประกาศ
func (s *AnnouncementService) DeleteAnnouncementService(id string) error {
	found, err := s.repo.DeleteAnnouncement(id)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("announcement not found")
	}
	return nil
}