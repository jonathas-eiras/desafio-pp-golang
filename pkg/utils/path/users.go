package path

func PathUsers() string {
	return "/users/"
}

func PathUserById(id string) string {
	return "/users/" + id
}
