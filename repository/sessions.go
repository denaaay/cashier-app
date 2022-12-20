package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type SessionsRepository struct {
	db db.DB
}

func NewSessionsRepository(db db.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) ReadSessions() ([]model.Session, error) {
	records, err := u.db.Load("sessions")
	if err != nil {
		return nil, err
	}

	var listSessions []model.Session
	err = json.Unmarshal([]byte(records), &listSessions)
	if err != nil {
		return nil, err
	}

	return listSessions, nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}

	// Select target token and delete from listSessions
	// TODO: answer here
	var newListSessions []model.Session
	for _, j := range listSessions {
		if tokenTarget == j.Token {
			continue
		} else {
			newListSessions = append(newListSessions, j)
		}
	}

	jsonData, err := json.Marshal(newListSessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	listSesi, err := u.ReadSessions()
	listSesi = append(listSesi, session)

	listSesiJson, err := json.Marshal(listSesi)
	err = ioutil.WriteFile("./data/sessions.json", listSesiJson, 0644)

	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (u *SessionsRepository) CheckExpireToken(token string) (model.Session, error) {
	element, err := u.TokenExist(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(element) == false {
		return element, nil
	}

	err = u.DeleteSessions(element.Token)
	if err != nil {
		return model.Session{}, err
	}

	return model.Session{}, errors.New("Token is Expired!") // TODO: replace this
}

func (u *SessionsRepository) ResetSessions() error {
	err := u.db.Reset("sessions", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) TokenExist(req string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}
	for _, element := range listSessions {
		if element.Token == req {
			return element, nil
		}
	}
	return model.Session{}, fmt.Errorf("Token Not Found!")
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
