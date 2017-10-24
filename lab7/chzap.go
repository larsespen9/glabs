// +build !solution

// Leave an empty line above this comment.
package lab7

import (
	"errors"
	"strings"
	"time"
	"unicode"
)

const timeFormat = "2006/01/02, 15:04:05"
const dateFormat = "2006/01/02"
const timeOnly = "15:04:05"
const timeLen = len(timeFormat)

type StatusChange struct {
	Time       time.Time
	ipadr      string
	statusType string
	value      string
	//TODO finish this struct (1p)
}

type ChZap struct {
	Time     time.Time
	ipadr    string
	toChan   string
	fromChan string
	//TODO finish this struct (1p)
}

func NewSTBEvent(event string) (*ChZap, *StatusChange, error) {
	heleStringen := fjernMellomrom(event)
	stringListe := strings.Split(heleStringen, ",")
	tidogdato := stringListe[0] + ", " + stringListe[1]
	time, err := time.Parse(timeFormat, tidogdato)
	if err != nil {
		return nil, nil, errors.New("NewSTBEvent: failed to parse timestamp")
	}
	if len(stringListe) == 4 {
		s := new(StatusChange)
		s.Time = time
		if s.Time.Format(timeOnly) == "00:00:00" {
			errorString := "NewSTBEvent: too short event string: " + s.Time.Format(timeFormat)
			err := errors.New(errorString)
			return nil, s, err
		}
		s.ipadr = stringListe[2]
		status := strings.Split(stringListe[3], ":")
		s.statusType = status[0]
		s.value = status[1]
		return nil, s, nil
	} else if len(stringListe) == 5 {
		c := new(ChZap)
		c.Time = time
		c.ipadr = stringListe[2]
		c.toChan = stringListe[3]
		c.fromChan = stringListe[4]
		return c, nil, nil
	} else if len(stringListe) > 5 {
		return nil, nil, errors.New("NewSTBEvent: event with too many fields")
	}
	return nil, nil, errors.New("NewSTBEvent: too short event string:%q")

}

func (zap ChZap) String() string {
	return "Time:" + zap.Time.Format(timeFormat) + ", IP:" + zap.ipadr + ", ToChan:" + zap.toChan + ", FromChan:" + zap.fromChan
}

func (schg StatusChange) String() string {
	return "Time:" + schg.Time.Format(timeFormat) + ", IP:" + schg.ipadr + ", StatusType:" + schg.statusType + ", Value:" + schg.value
}

// The duration... between receiving (this) zap event and the provided event
func (zap ChZap) Duration(provided ChZap) time.Duration {
	return zap.Time.Sub(provided.Time)
}

func (zap ChZap) Date() string {
	return zap.Time.Format(dateFormat)
}

func fjernMellomrom(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
