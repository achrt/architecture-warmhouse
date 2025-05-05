package services

import (
	"errors"
	"fmt"
	"time"

	"math/rand"
)

type Temperature struct{}

type TemperatureAPI interface {
	Get(sensorID, location string) (int, error)
}

func NewTemperatureApi() *Temperature {
	return &Temperature{}
}

const (
	maxTemp = 60
	minTemp = -60
)

const (
	sensor1 = "1"
	sensor2 = "2"
	sensor3 = "3"
)

const (
	living  = "Living Room"
	bedroom = "Bedroom"
	kitchen = "Kitchen"
)

var ErrSensorUndefined = errors.New("sensor is undefined")

func (s *Temperature) Get(sensorID, location string) (int, error) {
	var err error

	if sensorID == "" {
		sensorID, err = s.findDevice(location)
		if err != nil {
			return 0, err
		}
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	t := rand.Intn(maxTemp-minTemp+1) + minTemp
	fmt.Println("sensor: "+sensorID+"; temperature: ", t)

	return t, nil
}

func (s *Temperature) findDevice(locationID string) (string, error) {
	if locationID == living {
		return sensor1, nil
	}

	if locationID == bedroom {
		return sensor2, nil
	}

	if locationID == kitchen {
		return sensor3, nil
	}

	return "", ErrSensorUndefined
}
