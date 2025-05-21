package model

type User struct {
	ID   int
	Name string
}

func GetUser(id int) (User, error) {

	if id == 1 {
		return User{ID: 1, Name: "John Doe"}, nil
	}
	return User{}, nil
}
