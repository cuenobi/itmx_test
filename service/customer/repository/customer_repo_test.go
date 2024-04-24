package repository

import (
	"testing"

	"itmx_test/service/entity"
	"itmx_test/util"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	gormDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error creating SQLite in-memory database: %s", err)
	}
	gormDB.AutoMigrate(&entity.Customer{})

	// สร้าง customerRepo โดยใช้ฐานข้อมูล SQLite ที่จำลอง
	cr := &customerRepo{gormDB}

	// เรียกใช้ Create และตรวจสอบว่าไม่มีข้อผิดพลาด
	err = cr.Create(&entity.Customer{ID: util.GenerateUuid(), Name: "test", Age: 30})
	assert.NoError(t, err)

	// ตรวจสอบว่าข้อมูลถูกสร้างขึ้นในฐานข้อมูลหรือไม่
	var count int64
	if err := gormDB.Model(&entity.Customer{}).Count(&count).Error; err != nil {
		t.Fatalf("error querying customers: %s", err)
	}
	assert.Equal(t, int64(1), count)
}
