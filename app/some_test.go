package app

import "testing"


func TestEmailExist(t *testing.T){
	members = []Member{
		{Name: "Bob", Email: "bob@mail.com", RegistrationDate: "17.01.2019"},
	}

    got := doesEmailExist("bob@mail.com")
    want := true

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }
}

func TestEmailNotExist(t *testing.T){
	members = []Member{
		{Name: "Bob", Email: "bob@mail.com", RegistrationDate: "17.01.2019"},
	}

    got := doesEmailExist("marry@mail.com")
    want := false

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }
}