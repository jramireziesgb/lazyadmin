package lazyadmin

import (
	"fmt"
	"math/rand"
	s "strings"
	"time"
)

// User Registro para almacenar los alumnos en el formato Lazy Admin Tools
type User struct {
	user       string
	first      string
	last       string
	password   string
	department string
	company    string
	street     string
	city       string
	tel        string
	forward    string
	email      string
	uid        string
	group      string
}

func init() {
	rand.Seed(time.Now().Unix())
}

func createUserName(firstName string, lastName string) string {

	r := s.NewReplacer("á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u", "ü", "u", "ñ", "n")
	first := r.Replace(s.ToLower(s.TrimSpace(firstName)))
	last := r.Replace(s.ToLower(s.TrimSpace(lastName)))

	firsts := s.Split(first, " ")
	lasts := s.Split(last, " ")

	username := ""

	for _, v := range firsts {
		username = username + string(v[0])
	}

	return username + lasts[0]
}

func randomPassword() string {
	passwd := capitales[rand.Intn(len(capitales))]

	return passwd + "." + string(rand.Intn(100))
}

func createPassword(password string) string {
	if password == "" || password == "RANDOM" {
		return randomPassword()
	}

	return password
}

// NewUser Crea un usuario nuevo
func (a *User) NewUser(prefix string, firstName string, lastName string, password string, department string, company string, street string, city string, tel string, forward string, email string, uid string, group string) {

	a.user = prefix + createUserName(firstName, lastName)
	a.first = firstName
	a.last = lastName
	a.password = createPassword(password)
	a.department = department
	a.company = company
	a.street = street
	a.city = city
	a.tel = tel
	a.forward = forward
	a.email = email
	a.uid = uid
	a.group = group
}

func (a *User) String() string {
	return fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s", a.user, a.first, a.last, a.password, a.department, a.company, a.street, a.city, a.tel, a.forward, a.email, a.uid, a.group)
}
