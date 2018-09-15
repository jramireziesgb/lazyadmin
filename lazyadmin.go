package lazyadmin

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

func createUserName(firstName string, lastName string) string {

	username := firstName + lastName
	return username
}

// NewUser Crea un usuario nuevo
func (a *User) NewUser(
	firstName string,
	lastName string,
	password string,
	department string,
	company string,
	street string,
	city string,
	tel string,
	forward string,
	email string,
	uid string,
	group string) {
	a.user = createUserName(firstName, lastName)
	a.first = firstName
	a.last = lastName
	a.password = password
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
