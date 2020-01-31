package handler

import (
	"html/template"
	"net/http"
	"strconv"
)

type TicketHandler struct {
	tmpl      *template.Template
	ticketSrv bet.TicketService
}

// NewTicketHandler initializes and returns new TicketHandler
func NewTicketHandler(T *template.Template, tSrv book.TicketService) *TicketHandler {
	return &TicketHandler{tmpl: T, ticketSrv: tSrv}
}

// Ticket displays all the tickets in the database
func (th *TicketHandler) Ticket(w http.ResponseWriter, r *http.Request) {
	tickets, errs := th.ticketSrv.Tickets()
	if errs != nil {
		panic(errs)
	}
	th.tmpl.ExecuteTemplate(w, "admin.ticket.layout", tickets)
}

// UpdateTicket handle requests on /user/ticket/update
func (th *TicketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}

		tkt, errs := th.ticketSrv.Ticket(uint(id))
		if len(errs) > 0 {
			panic(errs)
		}

		th.tmpl.ExecuteTemplate(w, "admin.ticket.update.layout", tkt)

	} else if r.Method == http.MethodPost {

		tkt := &entity.Ticket{}
		id, _ := strconv.Atoi(r.FormValue("id"))
		fId, _ := strconv.Atoi(r.FormValue("bet_id"))
		uId, _ := strconv.Atoi(r.FormValue("user_id"))
		status, _ := strconv.ParseBool(r.FormValue("status"))

		tkt.ID = uint(id)
		tkt.UserID = uint(uId)
		tkt.FlightID = uint(fId)
		tkt.Status = status

		_, errs := th.ticketSrv.UpdateTicket(tkt)
		if len(errs) > 0 {
			panic(errs)
		}

		http.Redirect(w, r, "/admin/ticket", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/ticket", http.StatusSeeOther)
	}
}


// DeleteTicket handles requests on route /admin/ticket/delete
func (th *TicketHandler) DeleteTicket(w http.ResponseWriter, r*http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
			if err != nil {
				panic(err)
			}

		_, errs := th.ticketSrv.DeleteTicket(uint(id))
		if len(errs) > 0 {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/ticket", http.StatusSeeOther)
	}
