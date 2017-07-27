package model

type User struct {
	Id   int;
	Name string;
}

func (this *User)SayHello() string {
	return "Hi " + this.Name;
}
