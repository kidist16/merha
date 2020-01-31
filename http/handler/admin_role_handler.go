package handler

import (
	"github.com/kidist16/sport-betting/entity"
	"github.com/kidist16/sport-betting/user"
	"html/template"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	tmpl        *template.Template
	roleService user.RoleService
}

// NewRoleHandler initializes and returns new RoleHandler
func NewRoleHandler(T *template.Template, rSrv user.RoleService) *RoleHandler {
	return &RoleHandler{tmpl: T, roleService: rSrv}
}

// Role displays all the roles in the database
func (rh *RoleHandler) Role(w http.ResponseWriter, r *http.Request) {
	roles, errs := rh.roleService.Roles()
	if errs != nil {
		panic(errs)
	}
	rh.tmpl.ExecuteTemplate(w, "admin.role.layout", roles)
}

// StoreRole creates new role in the database
func (rh *RoleHandler) StoreRole(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		rl := &entity.Role{}

		rl.Name = r.FormValue("roleName")

		_, errs := rh.roleService.StoreRole(rl)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w, r, "/admin/role", http.StatusSeeOther)

	} else {

		rh.tmpl.ExecuteTemplate(w, "admin.role.create.layout", nil)

	}
}

// UpdateRole handle requests on /admin/role/update
func (rh *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}

		rl, errs := rh.roleService.Role(uint(id))
		if len(errs) > 0 {
			panic(errs)
		}

		rh.tmpl.ExecuteTemplate(w, "admin.role.update.layout", rl)

	} else if r.Method == http.MethodPost {

		rl := &entity.Role{}
		id, _ := strconv.Atoi(r.FormValue("id"))

		rl.ID = uint(id)
		rl.Name = r.FormValue("roleName")

		_, errs := rh.roleService.UpdateRole(rl)
		if len(errs) > 0 {
			panic(errs)
		}

		http.Redirect(w, r, "/admin/role", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/role", http.StatusSeeOther)
	}
}

// DeleteRole handles requests on route /admin/role/delete
func (rh *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}

		_, errs := rh.roleService.DeleteRole(uint(id))
		if len(errs) > 0 {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/role", http.StatusSeeOther)
}
