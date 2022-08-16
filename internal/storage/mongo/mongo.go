package mongo

import (
	"context"
	"github.com/dingowd/RB/internal/config"
	"github.com/dingowd/RB/internal/logger"
	"github.com/dingowd/RB/internal/storage"
	"github.com/dingowd/RB/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type StoreM struct {
	Log        logger.Logger
	Session    *mgo.Session
	Collection *mgo.Collection
	Conf       *config.Config
}

func New(log logger.Logger, conf *config.Config) *StoreM {
	return &StoreM{Log: log, Conf: conf}
}

func (s *StoreM) Connect(ctx context.Context, dsn string) error {
	var err error
	s.Session, err = mgo.Dial(s.Conf.DB.DSN)
	s.Collection = s.Session.DB(s.Conf.DB.DB).C(s.Conf.DB.Collection)
	return err
}

func (s *StoreM) Close() {
	s.Session.Close()
}

func (s *StoreM) GetAll() (*model.Students, error) {
	var d model.Students
	query := bson.M{}
	err := s.Collection.Find(query).All(&d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (s *StoreM) Update(m model.CacheStudent) error {
	tm, _ := time.Parse("02.01.2006", m.BirthDate)
	query1 := bson.M{"_id": bson.ObjectIdHex(m.Id)}
	query2 := bson.M{"$set": bson.M{"first_name": m.FirstName, "second_name": m.SecondName, "faculty": m.Faculty, "birth_date": tm.Unix()}}
	err := s.Collection.Update(query1, query2)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreM) Delete(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := s.Collection.Remove(query)
	return err
}

func (s *StoreM) Insert(j model.ForJson) error {
	var m model.Student
	m.Id = bson.NewObjectId()
	m.FirstName = j.FirstName
	m.SecondName = j.SecondName
	m.Faculty = j.Faculty
	t, err := time.Parse("02.01.2006", j.BirthDate)
	if err != nil {
		return err
	}
	m.BirthDate = t.Unix()
	if s.IsDocumentExist(m) {
		return storage.ErrorDocumentExist
	}
	if err := s.Collection.Insert(m); err != nil {
		return err
	}
	return nil
}

func (s *StoreM) IsDocumentExist(m model.Student) bool {
	var d model.Students
	query := bson.M{"first_name": m.FirstName, "second_name": m.SecondName, "faculty": m.Faculty, "birth_date": m.BirthDate}
	err := s.Collection.Find(query).All(&d)
	if err != nil {
		return false
	}
	if len(d) > 0 {
		return true
	}
	return false
}
