package model

import (
	"foreignKey/pkg/logger"
	"time"

	"gorm.io/gorm"
)

type Foreign2 struct {
	gorm.Model
	Timestamp  time.Time `gorm:"type:timestamp"`
	Device_PID uint      `gorm:"column:device_pid"`
	Device_UID Devices   `gorm:"foreignKey:Device_PID;references:ID"`
	Data       uint      `gorm:"column:data"`
}

func (Foreign2) TableName() string {
	return "foreign2"
}
func (m Foreign2) Create(db *gorm.DB) error {
	return db.Omit("ID").Create(&m).Error
}
func (m Foreign2) Update(db *gorm.DB) error {
	return db.Select("*").Where("id = ?", m.ID).Updates(&m).Error
}

func (x Foreign2) Find(db *gorm.DB, device_pid uint) (Foreign2, error) {
	m := Foreign2{}
	err := db.Table(m.TableName()).Where("device_pid = ?", device_pid).First(&m)
	if err.RowsAffected == 0 || err.Error == gorm.ErrRecordNotFound {
		return m, logger.ErrorNotFound
	}
	if err.Error != nil {
		return m, err.Error
	}

	return m, nil
}

func (m Foreign2) Delete(db *gorm.DB) error { //soft
	return db.Table(m.TableName()).Where("device_pid = ?", m.Device_PID).Delete(&m).Error
}

func (m Foreign2) Remove(db *gorm.DB) error { //permanent
	return db.Unscoped().Delete(&m).Error
}
