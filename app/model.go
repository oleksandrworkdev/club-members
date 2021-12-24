package app

type IndexPageData struct {
	Members []MemberView
	ErrorText string
}

type Member struct {
	Name             string
	Email            string
	RegistrationDate string
}

type MemberView struct {
	Id               int
	Name             string
	Email            string
	RegistrationDate string
}