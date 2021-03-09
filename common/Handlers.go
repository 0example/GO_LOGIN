package common

import (
	"github.com/gorilla/securecookie" //Package gorilla/securecookie encodes and decodes authenticated and optionally encrypted cookie values
	"net/http" //Package http provides HTTP client and server implementations.
	"fmt" //Package fmt implements formatted I/O with functions
	"../helpers" //
	"../repos" //
)

 //New returns a new SecureCookie.
var cookieHandler = securecookie.New(	
securecookie.GenerateRandomKey(64), //GenerateRandomKey creates a random key with the given strength.
	securecookie.GenerateRandomKey(32))

// Handlers

// for GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Print("wada")

	var body, _ = helpers.LoadFile("lk.wixis360.login/templates/login.html")
	fmt.Fprintf(response, body)
}

// for POST
//LoginHandler reads the name and the password from the submitted form
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")

	fmt.Println(pass)
	redirectTarget := "/"
	if !helpers.IsEmpty(name) && !helpers.IsEmpty(pass) {
		// Database check for user data!
		_userIsValid := repos.UserIsValid(name, pass)

		if _userIsValid {
			SetCookie(name, response)
			redirectTarget = "/index"
		} else {
			redirectTarget = "/register"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// for GET
func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("lk.wixis360.login/templates/register.html")
	fmt.Fprintf(response, body)
}

// for POST
 //RegisterHandler is the handler for the "Add Account" page
func RegisterHandler(w http.ResponseWriter, r *http.Request) {	

r.ParseForm()
	uName := r.FormValue("username")
	email := r.FormValue("email")
	pwd := r.FormValue("password")
	confirmPwd := r.FormValue("confirmPassword")

	_uName, _email, _pwd, _confirmPwd := false, false, false, false
	_uName = !helpers.IsEmpty(uName)
	_email = !helpers.IsEmpty(email)
	_pwd = !helpers.IsEmpty(pwd)
	_confirmPwd = !helpers.IsEmpty(confirmPwd)

	if _uName && _email && _pwd && _confirmPwd {
		fmt.Fprintln(w, "Username for Register : ", uName)
		fmt.Fprintln(w, "Email for Register : ", email)
		fmt.Fprintln(w, "Password for Register : ", pwd)
		fmt.Fprintln(w, "ConfirmPassword for Register : ", confirmPwd)
	} else {
		fmt.Fprintln(w, "This fields can not be blank!")
	}
}

// for GET
//IndexPageHandler returns the defined template
func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)
	if !helpers.IsEmpty(userName) {
		var indexBody, _ = helpers.LoadFile("templates/index.html")
		fmt.Fprintf(response, indexBody, userName)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}

// for POST
//LogoutHandler clears any existing session and redirects to the login page.
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	ClearCookie(response)
	http.Redirect(response, request, "/", 302)
}

// Cookie
//SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers.
func SetCookie(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearCookie(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
