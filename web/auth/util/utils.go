package util

func ConvertUserstoHTML(users []string) string {
	var data string
	for _, user := range users {
		data += "<form action=\"follow\" method=\"post\"> "
		data += "<input type=\"submit\" value=\"" + user + " \" readonly=\"readonly\" "
		data += "</form>\n"
	}
	return data
}
