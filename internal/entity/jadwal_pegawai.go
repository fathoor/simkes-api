package entity

import (
	"time"
)

type JadwalPegawai struct {
	NIP       string    `gorm:"column:nip;primaryKey"`
	Pegawai   Pegawai   `gorm:"foreignKey:nip;references:nip"`
	Tahun     int16     `gorm:"column:tahun;primaryKey"`
	Bulan     int16     `gorm:"column:bulan;primaryKey"`
	Hari      int16     `gorm:"column:hari;primaryKey"`
	ShiftNama string    `gorm:"column:shift_nama;not null"`
	Shift     Shift     `gorm:"foreignKey:shift_nama;references:nama"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (JadwalPegawai) TableName() string {
	return "jadwal_pegawai"
}
