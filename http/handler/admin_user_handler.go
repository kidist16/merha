package handler

import (
	"html/template"
	"net/http"
	"net/url"
	"strconv"


	"golang.org/x/crypto/bcrypt"
	"github.com//kidist16/sport-betting/rtoken"
	"github.com//kidist16/sport-betting/entity"
	"github.com/kidist16/sport-betting/user"
)

// AdminUserHandler handler handles user related requests
type AdminUserHandler struct {
	tmpl        *template.Template
	userService user.UserService
	loginService user.LoginService
	roleService    user.RoleServices
	csrfSignKey    []byte
}

// NewUserHandler returns new AdminUserHandler object
func NewAdminUserHandler(t *template.Template, usrServ user.UserService, lgnServ user.LoginService, uRole user.RoleService, csKey []byte) *AdminUserHander {
	return &AdminUserHandler{tmpl: t, userService: usrServ, loginService: lgnServ, roleService: uRole, csrfSignKey: csKey}
}

// AdminUsers handles Get /admin/users request
func (auh *AdminUserHandler) AdminUser(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(auh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	usrs, errs := auh.userService.Users()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	lgn, errs := auh.loginService.Logins()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		User   []entity.User
		Login	[]entity.Login
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		User:   usrs,
		Login:   lgn,
		CSRF:    token,
	}
	auh.tmpl.ExecuteTemplate(w, "adminuser.html", tmplData)
}

// AdminUsersNew handles GET/POST /admin/users/create request
func (auh *AdminUserHandler) AdminStoreUser(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(auh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		
		roles, errs := auh.roleService.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		accountForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Roles   []entity.Role
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			Roles:   roles,
			CSRF:    token,
		}
		auh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate the form contents
		accountForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		accountForm.Required("FName", "LName","email", "password", "confirmPassword")
		accountForm.MatchesPattern("email", form.EmailRX)
		accountForm.MatchesPattern("username", form.UsernameRX)
		accountForm.MinLength("password", 8)
		accountForm.PasswordMatches("password", "confirmpassword")
		accountForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !accountForm.Valid() {
			auh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
	
		uExists := uh.loginService.UsernameExists(r.FormValue("username")) 
		if uExists {
			upAccForm.VErrors.Add("usename", "Usename Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		eExists := uh.userService.EmailExists(r.FormValue("email"))
		if eExists {
			upAccForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			accountForm.VErrors.Add("password", "Password Could not be stored")
			auh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		roleId, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			accountForm.VErrors.Add("role", "could not retrieve role id")
			auh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		// ready to store
		usr := &entity.User{
			FName = r.FormValue("FName"),
			LName = r.FormValue("LName"),
			Email = r.FormValue("email"),
			Passport = r.FormValue("passport"),
			Registered = true,
			RoleID:   role.ID,
		}
		// store user info
		getUsr, errs := auh.userSrv.StoreUser(usr)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		lgn := &entity.Login{
			Username = r.FormValue("username"),
			Password: string(hashedPassword),
			UserID = getUsr.ID,
		}
		// store login info
		_, errs := uh.userSrv.StoreLogin(lgn)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)		
	}
}

// AdminUsersUpdate handles GET/POST /admin/users/update?id={id} request
func (auh *AdminUserHandler) AdminUpdateUser(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(auh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		usr, errs := auh.userService.User(uint(UserID))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		lgn, errs := auh.loginService.LoginByUserId(usr.UserID)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		rls, errs := auh.roleService.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		rl, errs := auh.roleService.Role(user.RoleID)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		values := url.Values{}
		values.Add("userId", idRaw)
		values.Add("FName", usr.FName)
		values.Add("LName", usr.LName)
		values.Add("email", usr.Email)
		values.Add("role", string(usr.RoleID))
		Values.Add("username", lgn.Username)
		Values.Add("password", lgn.Password)
		values.Add("roleName", rl.Name)

		upAccForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Roles   []entity.Role
			User    *entity.User
			CSRF    string
		}{
			Values:  values,
			VErrors: form.ValidationErrors{},
			Roles:   rls,
			User:    usr,
			CSRF:    token,
		}
		auh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		upAccForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		upAccForm.Required("FName", "LName","email", "password", "confirmPassword")
		upAccForm.MatchesPattern("email", form.EmailRX)
		upAccForm.MatchesPattern("username", form.UsernameRX)
		upAccForm.MinLength("password", 8)
		upAccForm.PasswordMatches("password", "confirmpassword")
		upAccForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !upAccForm.Valid() {
			auh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}
		usrId := r.FormValue("userId")
		uid, err := strconv.Atoi(usrId)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		getUsr, errs := auh.userService.User(uint(uid))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		getLgn, errs := auh.loginService.LoginByUserId(getUsr.ID))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}



		eExists := auh.userService.EmailExists(r.FormValue("email"))
		if (getUsr.Email != r.FormValue("email")) && eExists {
			upAccForm.VErrors.Add("email", "Email Already Exists")
			auh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}
		uExists := uh.loginService.UsernameExists(r.FormValue("username")) 
		if (getLgn.Username != r.FormValue("username")) && uExists {
			upAccForm.VErrors.Add("username", "Username Already Exists")
			auh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}

		roleId, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			upAccForm.VErrors.Add("role", "could not retrieve role id")
			auh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}
		// ready to store
		usr := &entity.User{
			ID:       getUsr.ID,
			FName = r.FormValue("FName"),
			LName = r.FormValue("LName"),
			Email = r.FormValue("email"),
			Password = r.FormValue("password"),
			Registered = true,
			RoleID:   uint(roleId),
		}
		// store user info
		getUsr, errs := auh.userService.UpdateUser(usr)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		lgn := &entity.Login{
			ID:       getLgn.ID,
			Username = r.FormValue("username"),
			Password: string(hashedPassword),
			UserID = getUsr.ID,
		}
		// store login info
		_, errs := auh.userSrv.UpdateLogin(lgn)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}
}

// AdminUsersDelete handles Delete /admin/users/delete?id={id} request
func (auh *AdminUserHandler) AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := auh.userService.DeleteUser(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
