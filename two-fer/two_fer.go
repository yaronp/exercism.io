package twofer

// ShareWith is a software defined twofar device
func ShareWith(name string) string {

	var forWho = name
	if name == "" {
		forWho = "you"
	}
	return "One for " + forWho + ", one for me."
}
