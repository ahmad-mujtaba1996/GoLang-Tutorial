package auth

import "fmt"

// Starting function with small letter makes it unexported (private) to the package. Can only be used in current folder(auth) files
// Starting function with capital letter makes it exported (public) to the package. Can now be used in other folder files as well using import
func LoginWithCredentials(username, password string) {
	fmt.Println("Logging user using:", username, password)
}
