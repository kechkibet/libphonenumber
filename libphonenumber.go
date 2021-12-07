package libphonenumber

import (
	"fmt"
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
	channel.HandleFunc("getRegionInfo", handleGetRegionInfo)
	return nil // no error
}

func handleGetRegionInfo(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	phoneNumber := argsMap["phone_number"].(string)
	isoCode := argsMap["iso_code"].(string)

	p, err := libphonenumber.Parse(phoneNumber, strings.ToUpper(isoCode))

	regionCode := libphonenumber.GetRegionCodeForNumber(p)
	countryCode := p.CountryCode
	formattedNumber := libphonenumber.Format(p, libphonenumber.NATIONAL)

	ret := make(map[string]string)
	ret["regionCode"] = regionCode
	ret["countryCode"] = fmt.Sprintf("%d", countryCode)
	ret["formattedNumber"] = formattedNumber

	return ret, nil

}

func handleisValidPhoneNumber(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	phoneNumber := argsMap["phone_number"].(string)
	isoCode := argsMap["iso_code"].(string)

	println(phoneNumber)
	println(isoCode)

	p, err := libphonenumber.Parse(phoneNumber, strings.ToUpper(isoCode))
	if err != nil {
		log.Println(err)
		return false, nil
	}
	isValid := libphonenumber.IsValidNumber(p)
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
