package model

import (
	"foreignKey/pkg/logger"
	"time"

	"gorm.io/gorm"
)

type Devices struct {
	gorm.Model
	Timestamp time.Time `gorm:"type:timestamp(0)"`
	Device_ID string    `gorm:"column:device_id"`
}

func (Devices) TableName() string {
	return "devices"
}
func (m Devices) Create(db *gorm.DB) error {
	return db.Omit("ID").Create(&m).Error
}
func (m Devices) Update(db *gorm.DB) error {
	return db.Select("*").Where("id = ?", m.ID).Updates(&m).Error
}

func (x Devices) Find(db *gorm.DB, device_id string) (Devices, error) {
	m := Devices{}
	err := db.Table(m.TableName()).Where("device_id = ?", device_id).First(&m)
	if err.RowsAffected == 0 || err.Error == gorm.ErrRecordNotFound {
		return m, logger.ErrorNotFound
	}
	if err.Error != nil {
		return m, err.Error
	}

	return m, nil
}

func (m Devices) Delete(db *gorm.DB) error { //soft
	return db.Table(m.TableName()).Where("device_id = ?", m.Device_ID).Delete(&m).Error
}

func (m Devices) Remove(db *gorm.DB) error { //permanent
	return db.Unscoped().Delete(&m).Error
}
