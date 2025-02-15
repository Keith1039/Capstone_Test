package parameters

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	regen "github.com/zach-klippenstein/goregen"
)

const DEFAULTEXPR = "[a-zA-Z]+"
const DEFAULVARCHARCODE = REGEX

type VarcharColumnParser struct {
	Column column
}

func (p *VarcharColumnParser) ParseColumn() (string, error) {
	code := p.Column.Code
	if code == 0 {
		code = DEFAULVARCHARCODE
	}
	if code == REGEX {
		return p.handleRegex()
	} else if code == EMAIL {
		return p.handleEmail()
	} else if code == FIRSTNAME {
		return p.handleFirstName()
	} else if code == LASTNAME {
		return p.handleLastName()
	} else if code == FULLNAME {
		return p.handleFullName()
	} else if code == PHONE {
		return p.handlePhone()
	} else if code == COUNTRY {
		return p.handleCountry()
	} else if code == ADDRESS {
		return p.handleAddress()
	} else if code == ZIPCODE {
		return p.handleZipCode()
	} else if code == CITY {
		return p.handleCity()
	} else if code == STATIC {
		return p.handleStatic()
	} else if code == NULL {
		return p.handleNull()
	} else {
		return "", errors.New("invalid code given")
	}
}

func (p *VarcharColumnParser) handleRegex() (string, error) {
	expression := p.Column.Other["Expression"]
	if expression == "" {
		expression = DEFAULTEXPR
	}
	genString, err := regen.Generate(expression)
	return genString, err
}

func (p *VarcharColumnParser) handleEmail() (string, error) {
	return gofakeit.Email(), nil
}

func (p *VarcharColumnParser) handleFirstName() (string, error) {
	return gofakeit.Person().FirstName, nil
}

func (p *VarcharColumnParser) handleLastName() (string, error) {
	return gofakeit.Person().LastName, nil
}

func (p *VarcharColumnParser) handleFullName() (string, error) {
	return gofakeit.Name(), nil
}

func (p *VarcharColumnParser) handlePhone() (string, error) {
	return gofakeit.Phone(), nil
}

func (p *VarcharColumnParser) handleCountry() (string, error) {
	return gofakeit.Country(), nil
}

func (p *VarcharColumnParser) handleAddress() (string, error) {
	return gofakeit.Address().Address, nil
}

func (p *VarcharColumnParser) handleZipCode() (string, error) {
	return gofakeit.Zip(), nil
}

func (p *VarcharColumnParser) handleCity() (string, error) {
	return gofakeit.City(), nil
}

func (p *VarcharColumnParser) handleStatic() (string, error) {
	return p.Column.Other["Value"], nil
}

func (p *VarcharColumnParser) handleNull() (string, error) {
	return "NULL", nil
}
