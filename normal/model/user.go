package model

type User struct {

}

func (*User)Output() (r string, err error) {
	r = "Hi";
	return;
}
