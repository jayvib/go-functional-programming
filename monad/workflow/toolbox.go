package workflow

import (
	"encoding/base64"
	"github.com/go-functional-programming/monad/car"
	"encoding/json"
	"time"
	"strconv"
)

// Base64toBytes will return a Monad type function that will
// decode the base 64 encoded data into bytes.
func Base64ToBytes(d Data) Monad {
	dString := d.(string)
	return func(e error) (Data, error) {
		return base64.StdEncoding.DecodeString(dString)
	}
}

func BytesToData(d Data) Monad {
	b := d.([]byte)
	return func(e error) (Data, error) {
		return Data(b), e
	}
}

func TimestampData(d Data) Monad {
	b := d.([]byte)
	return func(e error) (Data, error) {
		_car := new(car.Car)
		err := json.Unmarshal(b, _car)
		if err != nil {
			return nil, err
		}
		_car.Timestamp = StringTimestamp()
		b, err = json.Marshal(_car)
		if err != nil {
			return nil, err
		}
		return Data(b), nil
	}
}

func DataToJson(d Data) Monad {
	return func(e error) (Data, error) {
		b, err := json.Marshal(d)
		if err != nil {
			return nil, err
		}
		return Data(b), nil
	}
}

func StringTimestamp() string {
	now := time.Now().Unix()
	return strconv.Itoa(int(now))
}

func Get(d Data) Monad {
	return func(e error) (Data, error) {
		return d, e
	}
}

func Next(m Monad, f func(Data) Monad) Monad {
	return func(e error) (Data, error) {
		newData, newError := m(e)
		if newError != nil {
			return nil, newError
		}
		return f(newData)(newError)
	}
}