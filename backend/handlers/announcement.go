package handlers

import (
	"announcement-board/models"
	"announcement-board/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// รับ request จาก client และส่งต่อไปยัง service

type AnnouncementHandler struct {
	service *services.AnnouncementService
}

// NewAnnouncementHandler สร้าง instance ของ AnnouncementHandler
func NewAnnouncementHandler(service *services.AnnouncementService) *AnnouncementHandler {
	return &AnnouncementHandler{service: service}
}

// getAll
func ( h *AnnouncementHandler) GetAll(c *gin.Context){
	announcements, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดึงข้อมูลได้"})
		return
	}
	c.JSON(http.StatusOK, announcements)
}

// create ประกาศใหม่ เส้น post โดยรับข้อมูลจาก body เป็น json และแปลงเป็น struct CreateRequest ก่อนส่งไปให้ service เพื่อสร้างประกาศใหม่ใน database
func (h *AnnouncementHandler) CreateAnnouncementRequest(c *gin.Context) {
	var req models.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	announcement, err := h.service.CreateAnnouncementService(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "สร้างประกาศไม่สำเร็จ"})
		return
	}
	c.JSON(http.StatusCreated, announcement)
}

// update ประกาศ เส้น put โดยรับ id จาก url และข้อมูลใหม่จาก body
// announcement/:id
func (h *AnnouncementHandler) UpdateAnnouncementRequest(c *gin.Context){
	id := c.Param("id")
	var req models.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	announcement, err := h.service.UpdateAnnouncementService(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "แก้ไขประกาศไม่สำเร็จ"})
		return
	}
	c.JSON(http.StatusOK, announcement)
}

// delete ประกาศ เส้น delete โดยรับ id จาก url
// announcement/:id
func (h *AnnouncementHandler) DeleteAnnouncementRequest(c *gin.Context){
	id := c.Param("id")
	if err := h.service.DeleteAnnouncementService(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ลบประกาศไม่สำเร็จ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ลบประกาศสำเร็จ"})
}



