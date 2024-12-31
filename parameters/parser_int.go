package parameters

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

type IntColumnParser struct {
	Column column
	latest int
}

const DEFAULTRANGE = "0,100"
const DEFAULTSTATIC = "0"

func (p *IntColumnParser) ParseColumn() (string, error) {
	code := p.Column.Code
	if code == RANDOM {
		return p.handleRandomCode()
	} else if code == STATIC {
		return p.handleStatic()
	} else if code == SEQ {
		l := p.latest // get the latest
		p.latest++    // increment
		return strconv.Itoa(l), nil
	} else if code == NULL {
		return "NULL", nil
	} else {
		err := errors.New("invalid code given")
		return "", err
	}
}

func (p *IntColumnParser) handleRandomCode() (string, error) {
	var value string
	var err error
	r, ok := p.Column.Other["Range"]
	if !ok {
		r = DEFAULTRANGE
	}
	ranges := strings.Split(r, ",")
	if len(ranges) != 2 {
		err = errors.New("malformed range")
		return "", err
	}
	lowerBound, boundErr := strconv.Atoi(ranges[0])
	if boundErr != nil {
		return "", boundErr
	}
	upperBound, boundErr2 := strconv.Atoi(ranges[1])
	if boundErr2 != nil {
		return "", boundErr
	}
	if lowerBound > upperBound {
		err = errors.New("lower bound is greater than upper bound")
		return "", err
	}
	value = strconv.Itoa(rand.Intn(upperBound-lowerBound) + lowerBound)
	return value, err
}

func (p *IntColumnParser) handleStatic() (string, error) {
	var value string
	var err error
	r, ok := p.Column.Other["Value"]
	if !ok {
		r = DEFAULTSTATIC
	}
	_, err = strconv.Atoi(r)
	if err != nil {
		return "", err
	} else {
		value = r
	}
	return value, err
}
