package models

import (
	"time"

	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type Message struct {
	Email   string
	Url     string
	QQ      string
	CDate   time.Time
	Content string
}

func (message *Message) Validate(v *revel.Validation) {
	v.Check(message.Email,
		revel.Required{},
		revel.MaxSize{50},
	)
	v.Email(message.Email)
	v.Check(message.Url,
		revel.MaxSize{200},
	)
	v.Check(message.Content,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{1000},
	)
}

func (dao *Dao) InsertMessage(message *Message) error {
	messageCollection := dao.session.DB(DbName).C(MessageCollection)
	message.CDate = time.Now()
	err := messageCollection.Insert(message)
	if err != nil {
		revel.WARN.Printf("Unable to save Message: %v error %v", message, err)
	}
	return err
}

func (dao *Dao) FindAllMessages() []Message {
	messageCollection := dao.session.DB(DbName).C(MessageCollection)
	messages := []Message{}
	query := messageCollection.Find(bson.M{}).Sort("-cdate")
	query.All(&messages)
	return messages
}
