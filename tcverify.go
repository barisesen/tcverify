package tcverify

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/hakanersu/tcvalidate"
)

type Response struct {
	TCKimlikNoDogrulaResult string `xml:"Body>TCKimlikNoDogrulaResponse>TCKimlikNoDogrulaResult"`
}

func Validate(ID string) (bool, error) {
	if !validatetc.Validate(ID) {
		err := errors.New(ID + " tc numarası algoritmik olarak doğrulanamadı.")
		return false, err

	}
	return true, nil
}

func Check(ID string, name string, surname string, birthYear string) (bool, error) {
	rawXml := []byte(`<?xml version="1.0" encoding="utf-8"?>
		<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
			<soap:Body>
				<TCKimlikNoDogrula xmlns="http://tckimlik.nvi.gov.tr/WS">
					<TCKimlikNo>` + ID + `</TCKimlikNo>
					<Ad>` + strings.ToUpperSpecial(unicode.TurkishCase, name) + `</Ad>
					<Soyad>` + strings.ToUpperSpecial(unicode.TurkishCase, surname) + `</Soyad>
					<DogumYili>` + birthYear + `</DogumYili>
				</TCKimlikNoDogrula>
			</soap:Body>
		</soap:Envelope>`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://tckimlik.nvi.gov.tr/Service/KPSPublic.asmx", bytes.NewBuffer(rawXml))
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", "http://tckimlik.nvi.gov.tr/WS/TCKimlikNoDogrula")
	req.Header.Add("Host", "tckimlik.nvi.gov.tr")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	response := Response{}
	if err := xml.NewDecoder(resp.Body).Decode(&response); err != nil {
		return false, err
	}

	status, err := strconv.ParseBool(response.TCKimlikNoDogrulaResult)
	if err != nil {
		return false, err
	}

	if !status {
		error := errors.New("Bu bilgileri ait vatandaşlık doğrulanamadı.")
		return status, error
	}
	return status, nil
}
