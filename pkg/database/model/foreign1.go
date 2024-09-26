package model

import (
	"foreignKey/pkg/logger"
	"time"

	"gorm.io/gorm"
)

type Foreign1 struct {
	gorm.Model
	Timestamp  time.Time `gorm:"type:timestamp(0)"`
	Device_PID uint      `gorm:"column:device_pid"`
	Device_UID Devices   `gorm:"foreignKey:Device_PID"`
	Data       uint      `gorm:"column:data"`
}

func (Foreign1) TableName() string {
	return "foreign1"
}
func (m Foreign1) Create(db *gorm.DB) error {
	return db.Omit("ID").Create(&m).Error
}
func (m Foreign1) Update(db *gorm.DB) error {
	return db.Select("*").Where("id = ?", m.ID).Updates(&m).Error
}

func (x Foreign1) Find(db *gorm.DB, device_pid uint) (Foreign1, error) {
	m := Foreign1{}
	err := db.Table(m.TableName()).Where("device_pid = ?", device_pid).First(&m)
	if err.RowsAffected == 0 || err.Error == gorm.ErrRecordNotFound {
		return m, logger.ErrorNotFound
	}
	if err.Error != nil {
		return m, err.Error
	}

	return m, nil
}

func (m Foreign1) Delete(db *gorm.DB) error { //soft
	return db.Table(m.TableName()).Where("device_pid = ?", m.Device_PID).Delete(&m).Error
}

func (m Foreign1) Remove(db *gorm.DB) error { //permanent
	return db.Unscoped().Delete(&m).Error
}
