package model

import (
	"foreignKey/pkg/logger"
	"time"

	"gorm.io/gorm"
)

type Foreign3 struct {
	gorm.Model
	Timestamp  time.Time `gorm:"type:timestamp(0)"`
	Device_PID *uint     `gorm:"column:device_pid"`                                                                  // Foreign key
	Device_UID Devices   `gorm:"foreignKey:Device_PID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // relation to Devices with foreign key
	Data       uint      `gorm:"column:data"`
}

func (Foreign3) TableName() string {
	return "foreign3"
}
func (m Foreign3) Create(db *gorm.DB) error {
	return db.Omit("ID").Create(&m).Error
}
func (m Foreign3) Update(db *gorm.DB) error {
	return db.Select("*").Where("id = ?", m.ID).Updates(&m).Error
}

func (x Foreign3) Find(db *gorm.DB, device_pid uint) (Foreign3, error) {
	m := Foreign3{}
	err := db.Table(m.TableName()).Where("device_pid = ?", device_pid).First(&m)
	if err.RowsAffected == 0 || err.Error == gorm.ErrRecordNotFound {
		return m, logger.ErrorNotFound
	}
	if err.Error != nil {
		return m, err.Error
	}

	return m, nil
}

func (m Foreign3) Delete(db *gorm.DB) error { //soft
	return db.Table(m.TableName()).Where("device_pid = ?", m.Device_PID).Delete(&m).Error
}

func (m Foreign3) Remove(db *gorm.DB) error { //permanent
	return db.Unscoped().Delete(&m).Error
}
