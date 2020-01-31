package main

import (
	"html/template"

	 "kidist16/sport-betting/user/repository"
	 "github.com/kidist16/sport-betting/user/service"

	"github.com/kidist16/sport-betting/entity"
	"github.com/kidist16/sport-betting/user"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.User{}, &entity.Role{}, &entity.Session{}, &entity.match{}, &entity.sport{}, &entity.stat{}, &entity.legue{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:Postgres@localhost/sportbetdb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	createTables(dbconn)
	defer dbconn.Close()
	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	ticketRepo := brepim.NewTicketGormRepo(dbconn)
	ticketServ := bsrvim.NewTicketService(ticketRepo)
	userRepo := urepim.NewUserGormRepo(dbconn)
	userServ := usrvim.NewUserService(userRepo)
	sessionRepo := urepim.NewSessionGormRepo(dbconn)
	sessionServ := usrvim.NewSessionService(sessionRepo)
	roleRepo := urepim.NewRoleGormRepo(dbconn)
	roleServ := usrvim.NewRoleService(roleRepo)
	loginRepo := urepim.NewLoginGormRepo(dbconn)
	loginServ := usrvim.NewLoginService(loginRepo)
	fs := http.FileServer(http.Dir("../../ui/assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", mainHandler.Index)
	mux.HandleFunc("/result", mainHandler.result)
	mux.HandleFunc("/statistics", mainHandler.statistics)
	mux.HandleFunc("/admin", uh.Authenticated(uh.Authorized(http.HandlerFunc(mainHandler.Admin))))
	http.HandleFunc("/login", uh.Login)
	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
	http.HandleFunc("/signup", uh.Signup)
    mux.Handle("/admin/role", uh.Authenticated(uh.Authorized(http.HandlerFunc(roleHandler.role))))
	mux.Handle("/admin/role/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(roleHandler.UpdateRole))))
	mux.Handle("/admin/role/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(roleHandler.DeleteRole))))
	http.ListenAndServe(":8080", mux)

}
