package models

import (
	"time"

	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type Email struct {
	Email string
	CDate time.Time
}

func (dao *Dao) InsertEmail(email *Email) error {
	emailCollection := dao.session.DB(DbName).C(EmailCollection)
	err := emailCollection.Insert(email)
	if err != nil {
		revel.WARN.Printf("Unable to save email: %v", email)
	}
	return err
}

func (dao *Dao) FindEmails() []Email {
	emailCollection := dao.session.DB(DbName).C(EmailCollection)
	emails := []Email{}
	query := emailCollection.Find(bson.M{}).Sort("-cdate")
	query.All(&emails)
	return emails
}

func (dao *Dao) IsEmailExists(email string) bool {
	emailCollection := dao.session.DB(DbName).C(EmailCollection)
	query := emailCollection.Find(bson.M{"email": email})
	cnt, err := query.Count()
	if err != nil {
		revel.WARN.Printf("Unable to Count Email: error %v", err)
	}
	return cnt != 0
}
