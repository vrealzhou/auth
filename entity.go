package auth

import (
	"strconv"
	"strings"
)

// LoginType indicate login by email or login by mobile
type LoginType int

const (
	// EmailLogin indicates login by email
	EmailLogin LoginType = 1
	// PhoneLogin login by mobile number
	PhoneLogin LoginType = 2
)

// User struct for login and basic user info
type User struct {
	ID       string `json:"id"`               // user id
	Email    string `json:"email,omitpemty"`  // user email
	Mobile   *Phone `json:"mobile,omitempty"` // user mobile number
	Password string `json:"password"`         // user password
}

// Phone struct is for hold phone number info
type Phone struct {
	Region int `json:"region"`        // country/region code such as AU, US, etc
	Number int `json:"number"`        // phone number
	Ext    int `json:"ext,omitempty"` // extention code
}

// Format function is for format phone number base on country
func (p *Phone) Format() string {
	buf := &strings.Builder{}
	buf.WriteRune('+')
	buf.WriteString(strconv.Itoa(p.Region))
	buf.WriteString(strconv.Itoa(p.Number))
	if p.Ext != 0 {
		buf.WriteRune(' ')
		buf.WriteString(strconv.Itoa(p.Ext))
	}
	return buf.String()
}

// Dial function is for format phone for tel: link on page
func (p *Phone) Dial() string {
	buf := &strings.Builder{}
	buf.WriteRune('+')
	buf.WriteString(strconv.Itoa(p.Region))
	buf.WriteString(strconv.Itoa(p.Number))
	if p.Ext != 0 {
		buf.WriteRune(' ')
		buf.WriteString(strconv.Itoa(p.Ext))
	}
	return buf.String()
}
