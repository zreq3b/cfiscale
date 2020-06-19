package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Person declare data to be passed
type Person struct {
	Name      string
	Surname   string
	BirthCity string
	BirthDate string
	Gender    string
	EndPoint  string
}

// URL defines the service provider domain 
const URL = "http://webservices.dotnethell.it"

func main() {
	p := NewPerson("silvio", "berlusconi", "milano", "29/09/1936", "M")
	p.BuildString()
	XML, err := p.Get()

	if err != nil {
		log.Println(err)
	}

	result, err := p.FormatData(XML)
	if err != nil {
		log.Println(err)
	}

	log.Println("Result:", result)
}

// Get retrieve endpoint data
func (p *Person) Get() (string, error) {
	var retVal = ""
	resp, err := http.Get(p.EndPoint)

	if err != nil {
		log.Println(err)
		return retVal, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
		return retVal, err
	}


	return string(body), nil
}

//FormatData return string from input xml string
func (p *Person) FormatData(inXML string) (string, error) {
	var fc string
	err := xml.Unmarshal([]byte(inXML), &fc)
	if err != nil {
		log.Println(err)
		return fc, err
	}

	return fc, nil
}

// BuildString return the full endpoint
func (p *Person) BuildString() {
	p.EndPoint = fmt.Sprintf(
		"%v/codicefiscale.asmx/CalcolaCodiceFiscale?Nome=%v&Cognome=%v&ComuneNascita=%v&DataNascita=%v&Sesso=%v",
		URL,
		p.Name,
		p.Surname,
		p.BirthCity,
		p.BirthDate,
		p.Gender,
	)
}

// NewPerson return Person object
func NewPerson(name string, surname string, birthCity string, birthDate string, gender string) Person {
	return Person{
		Name:      strings.ReplaceAll(name, " ", ""),
		Surname:   strings.ReplaceAll(surname, " ", ""),
		BirthCity: strings.ReplaceAll(birthCity, " ", ""),
		BirthDate: strings.ReplaceAll(birthDate, " ", ""),
		Gender:    strings.ReplaceAll(gender, " ", ""),
	}
}
