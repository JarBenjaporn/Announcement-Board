package test

import (
    "announcement-board/models"
	"announcement-board/validators"
    "testing"
)

func TestValidateCreateRequest(t *testing.T) {
    tests := []struct {
        name    string
        req     *models.CreateRequest
        wantErr bool
    }{
        {
            name:    "happy path",
            req:     &models.CreateRequest{Title: "ประกาศ", Body: "เนื้อหา", Author: "คนเขียน"},
            wantErr: false,
        },
        {
            name:    "title empty",
            req:     &models.CreateRequest{Title: "", Body: "เนื้อหา", Author: "คนเขียน"},
            wantErr: true,
        },
        {
            name:    "title whitespace only",
            req:     &models.CreateRequest{Title: "   ", Body: "เนื้อหา", Author: "คนเขียน"},
            wantErr: true,
        },
        {
            name:    "body empty",
            req:     &models.CreateRequest{Title: "ประกาศ", Body: "", Author: "คนเขียน"},
            wantErr: true,
        },
        {
            name:    "author empty",
            req:     &models.CreateRequest{Title: "ประกาศ", Body: "เนื้อหา", Author: ""},
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validators.ValidateCreateRequest(tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("got err = %v, wantErr = %v", err, tt.wantErr)
            }
        })
    }
}
