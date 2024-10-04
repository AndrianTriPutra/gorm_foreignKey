package model

import (
	"foreignKey/pkg/logger"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID         uint      `gorm:"primarykey"`
	Created_at time.Time `gorm:"type:timestamp(0);default:null"`
	Updated_at time.Time `gorm:"type:timestamp(0);default:null"`
	Deleted_at time.Time `gorm:"type:timestamp(0);default:null"`
}

type Foreign3 struct {
	Model
	Timestamp  time.Time `gorm:"type:timestamp(0)"`
	Device_PID *uint     `gorm:"column:device_pid"`                                                                  // Foreign key
	Device_UID Devices   `gorm:"foreignKey:Device_PID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // relation to Devices with foreign key
	Data       uint      `gorm:"column:data"`
}

func (Foreign3) TableName() string {
	return "foreign3"
}
func (m Foreign3) Create(db *gorm.DB) error {
	ts := time.Now().UTC()
	m.Model.Created_at = ts
	m.Model.Updated_at = ts
	return db.Omit("ID").Create(&m).Error
}
func (m Foreign3) Update(db *gorm.DB) error {
	ts := time.Now().UTC()
	m.Model.Updated_at = ts
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
	ts := time.Now().UTC()
	m.Model.Deleted_at = ts
	return db.Table(m.TableName()).Where("device_pid = ?", m.Device_PID).Delete(&m).Error
}

func (m Foreign3) Remove(db *gorm.DB) error { //permanent
	return db.Unscoped().Delete(&m).Error
}
