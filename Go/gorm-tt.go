package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PolicyObjectSet struct {
	gorm.Model
	// ID 는 자동 증가되는 PolicyObjectSet 의 식별자이다.
	ID uint64 `gorm:"primary_key;autoIncrement" json:"id,omitempty"`

	// Name 은 PolicyObjectSet 의 이름이다.
	Name string `gorm:"not null;unique" json:"name,omitempty"`

	// Remarks 는 PolicyObjectSet 의 비고이다.
	Remarks *string `json:"remarks,omitempty"`

	// CreatedAt 는 PolicyObjectSet 이 생성된 시간(UnixTime)이다.
	CreatedAt int64 `gorm:"autoCreateTime;not null" json:"createdAt,omitempty"`

	// UpdatedAt 는 PolicyObjectSet 정보가 수정된 최종시간(UnixTime)이다.
	UpdatedAt int64 `gorm:"autoUpdateTime;not null" json:"updatedAt,omitempty"`
}

func (PolicyObjectSet) TableName() string {
	return "policy_objectset"
}

func main() {
	dsn := "host=192.168.33.160 user=postgres password=dkssud12 dbname=customer port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	// 생성
	db.Create(&PolicyObjectSet{ID: 1, Name: "Jaden"})

	// 읽기
	var PolicyObjectSet PolicyObjectSet
	db.First(&PolicyObjectSet, 1) // primary key기준으로 product 찾기
	db.First(&PolicyObjectSet, "name = ?", "Jaden")

	// 수정 - product의 price를 200으로
	db.Model(&PolicyObjectSet).Update("Name", "cdm-replicator")
}
