package app

func doesEmailExist(email string) bool {
	doesExist := false
	for _, m := range members {
		if m.Email == email {
			doesExist = true
			break
		}
	}
	return doesExist
}

func getMembers() []MemberView {
	membersView := make([]MemberView, 0)
	for idx, m := range members {
		mv := MemberView{
			Id:               idx + 1,
			Name:             m.Name,
			Email:            m.Email,
			RegistrationDate: m.RegistrationDate,
		}

		membersView = append(membersView, mv)
	}
	return membersView
}