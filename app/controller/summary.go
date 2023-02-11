package controller

import (
	"summary/app/helpers"
	"summary/app/models"
	"summary/config"
	"sync"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	DDMMYYYY = "2006-01-02"
)

func CreateOrUpdate(startDate time.Time, providerCode string) ([]models.SummaryMembersJan, error) {
	payload := config.DB1
	db := config.ConnectionDB1(&payload)
	channel := make(chan *[]models.SummaryMembersJan)
	defer close(channel)
	var mutex = &sync.Mutex{}
	getTableName := helpers.GetTablename(providerCode, startDate.Month())

	go func() {
		data, err := GetSum(db, startDate, getTableName)

		if err != nil {
			panic(err.Error())
		}

		mutex.Lock()
		channel <- &data
		mutex.Unlock()

	}()

	responseChanel := <-channel

	db = db.Table(getTableName)
	// create then check if any conflict and update
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}, {Name: "merchant_code"}, {Name: "currency"}},
		DoUpdates: clause.AssignmentColumns([]string{"product", "amount", "category", "description"}), // column needed to be updated
	}).Create(&responseChanel).Error

	if err != nil {
		return nil, err
	}

	return *responseChanel, nil
}

func GetSum(db *gorm.DB, startDate time.Time, tableName string) ([]models.SummaryMembersJan, error) {
	channel := make(chan *[]models.SummaryMembersJan)
	defer close(channel)
	var mutex = &sync.Mutex{}

	go func() {
		var summary []models.SummaryMembersJan

		err := db.Table(tableName).Where("id = ?", startDate.Format(DDMMYYYY)).Find(&summary).Error

		if err != nil {
			panic(err.Error())
		}

		mutex.Lock()
		channel <- &summary
		mutex.Unlock()

	}()

	return *<-channel, nil
}
