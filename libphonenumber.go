package libphonenumber

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/ttacon/libphonenumber"
	"log"
	"strings"
)

const channelName = "codeheadlabs.com/libphonenumber"

type LibPhoneNumber struct{}

var _ flutter.Plugin = &LibPhoneNumber{} // compile-time type check

func (p *LibPhoneNumber) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("isValidPhoneNumber", handleisValidPhoneNumber)
	channel.HandleFunc("normalizePhoneNumber", handleisNormalizePhoneNumber)
	return nil // no error
}

func handleisValidPhoneNumber(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	phoneNumber := argsMap["phone_number"].(string)
	isoCode := argsMap["iso_code"].(string)

	println(phoneNumber)
	println(isoCode)

	p, err := libphonenumber.Parse(phoneNumber, strings.ToUpper(isoCode))
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	isValid := libphonenumber.IsValidNumber(p)

	println(isValid)

	return isValid, err
}

func handleisNormalizePhoneNumber(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	phone_number := argsMap["phone_number"].(string)
	iso_code := argsMap["iso_code"].(string)

	p, err := libphonenumber.Parse(phone_number, strings.ToUpper(iso_code))
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return libphonenumber.Format(p, libphonenumber.E164), err
}
