package test

import (
	"announcement-board/models"
	"announcement-board/repositories"
	"announcement-board/services"
	"errors"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=announcement_test port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("Cannot connect to test database:", err)
	}
	db.AutoMigrate(&models.Announcement{})
	return db
}

func TestCreateAnnouncement_HappyPath(t *testing.T) {
	db := setupTestDB(t)
	svc := services.NewAnnouncementService(repositories.NewAnnouncementRepository(db))

	req := &models.CreateRequest{
		Title:  "Test Announcement",
		Body:   "This is a test announcement",
		Author: "Tester",
		Pinned: false,
	}

	result, err := svc.CreateAnnouncement(req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer db.Delete(result)

	if result.ID == "" {
		t.Error("expected ID to be generated")
	}
	if result.Title != req.Title {
		t.Errorf("Title: expected %q, got %q", req.Title, result.Title)
	}
	if result.Body != req.Body {
		t.Errorf("Body: expected %q, got %q", req.Body, result.Body)
	}
	if result.Author != req.Author {
		t.Errorf("Author: expected %q, got %q", req.Author, result.Author)
	}
	if result.Pinned != req.Pinned {
		t.Errorf("Pinned: expected %v, got %v", req.Pinned, result.Pinned)
	}
}


func TestGetAll_PinnedFirst(t *testing.T) {
	db := setupTestDB(t)
	svc := services.NewAnnouncementService(repositories.NewAnnouncementRepository(db))

	unpinned, err := svc.CreateAnnouncement(&models.CreateRequest{
		Title: "ไม่ได้ pin", Body: "body", Author: "คน",
	})
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}
	defer db.Delete(unpinned)

	pinned, err := svc.CreateAnnouncement(&models.CreateRequest{
		Title: "pin แล้ว", Body: "body", Author: "คน", Pinned: true,
	})
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}
	defer db.Delete(pinned)

	results, err := svc.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// หา pinned ที่สร้างใน results แล้วเช็คว่าทุกอันก่อนหน้าต้อง pinned
	var found bool
	for i, r := range results {
		if r.ID == pinned.ID {
			found = true
			for _, prev := range results[:i] {
				if !prev.Pinned {
					t.Error("พบ unpinned announcement อยู่ก่อน pinned")
				}
			}
			break
		}
	}
	if !found {
		t.Error("ไม่พบ pinned announcement ใน results")
	}
}

func TestUpdateAnnouncement_HappyPath(t *testing.T) {
	db := setupTestDB(t)
	svc := services.NewAnnouncementService(repositories.NewAnnouncementRepository(db))

	created, err := svc.CreateAnnouncement(&models.CreateRequest{
		Title: "titleOld", Body: "bodyOld", Author: "tester1",
	})
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}
	defer db.Delete(created)

	updated, err := svc.UpdateAnnouncement(created.ID, &models.CreateRequest{
		Title: "titleNew", Body: "bodyNew", Author: "tester1", Pinned: true,
	})

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if updated.Title != "titleNew" {
		t.Errorf("Title: expected %q, got %q", "titleNew", updated.Title)
	}
	if updated.Body != "bodyNew" {
		t.Errorf("Body: expected %q, got %q", "bodyNew", updated.Body)
	}
	if !updated.Pinned {
		t.Error("expected pinned to be true")
	}
}

func TestUpdateAnnouncement_NotFound(t *testing.T) {
	db := setupTestDB(t)
	svc := services.NewAnnouncementService(repositories.NewAnnouncementRepository(db))

	_, err := svc.UpdateAnnouncement("00000000-0000-0000-0000-000000000000", &models.CreateRequest{
		Title: "title", Body: "body", Author: "คน",
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var notFound *models.NotFoundError
	if !errors.As(err, &notFound) {
		t.Errorf("expected NotFoundError, got %T", err)
	}
}

func TestDeleteAnnouncement_HappyPath(t *testing.T) {
	db := setupTestDB(t)
	svc := services.NewAnnouncementService(repositories.NewAnnouncementRepository(db))

	created, err := svc.CreateAnnouncement(&models.CreateRequest{
		Title: "จะลบ", Body: "body", Author: "คน",
	})
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	err = svc.DeleteAnnouncement(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestDeleteAnnouncement_NotFound(t *testing.T) {
	db := setupTestDB(t)
	svc := services.NewAnnouncementService(repositories.NewAnnouncementRepository(db))

	err := svc.DeleteAnnouncement("00000000-0000-0000-0000-000000000000")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var notFound *models.NotFoundError
	if !errors.As(err, &notFound) {
		t.Errorf("expected NotFoundError, got %T", err)
	}
}
