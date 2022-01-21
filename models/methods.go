package models

import (
	"fmt"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Method struct {
	Model
	FileID             uint
	File               File
	TestCaseID         uint
	TestCase           TestCase
	Name               string `gorm:"not null"`
	CallerID           *uint
	Caller             []Method `gorm:"foreignkey:CallerID"`
	OwnDuration        time.Duration
	CumulativeDuration time.Duration
	OwnCalls           int
	TotalCalls         int
	CallsPercent       float64
	Error              bool
	OwnSize            int
	CumulativeSize     int
	AllocCalls         int
	TotalAllocCalls    int
	AllocCallsPercent  float64
}

func (m *Method) TableName() string {
	return "methods"
}

func CreateMethod(db *gorm.DB, m *Method) (uint, error) {
	err := db.Create(m).Error
	if err != nil {
		fmt.Printf("Error creating Method %s: %s\n", m.Name, err.Error())
		return 0, err
	}
	return m.ID, nil
}

func SaveMethod(db *gorm.DB, m *Method) error {
	return db.Save(m).Error

}

func FindMethodByName(db *gorm.DB, methodname string) (*Method, error) {
	var method Method
	res := db.Where("name like ?", "%"+methodname)
	return &method, res.Error
}

func FindMethodByEndsWithNameAndFileAndTestcase(db *gorm.DB, methodname string, fileID uint, testcaseID uint) (*Method, error) {
	var method Method
	res := db.Where("name like ? and file_id=? AND test_case_id=?", "%"+methodname, fileID, testcaseID).First(&method)
	return &method, res.Error
}