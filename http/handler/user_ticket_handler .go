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
func NewTicketHandler(T *template.Template, tSrv bet.TicketService) *TicketHandler {
	return &TicketHandler{tmpl: T, ticketSrv: tSrv}
}

func (th *TicketHandler) Ticket(w http.ResponseWriter, r *http.Request) {
	tickets, errs := th.ticketSrv.Tickets()
	if errs != nil {
		panic(errs)
	}
	th.tmpl.ExecuteTemplate(w, "user.ticket.layout", tickets)
}
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

		th.tmpl.ExecuteTemplate(w, "user.ticket.update.layout", tkt)

	} else if r.Method == http.MethodPost {

		tkt := &entity.match{}
		id, _ := strconv.Atoi(r.FormValue("mid"))
		fId, _ := strconv.Atoi(r.FormValue("bet_id"))
		uId, _ := strconv.Atoi(r.FormValue("user_id"))
		status, _ := strconv.ParseBool(r.FormValue("status"))

		tkt.user_id = uint(id)
		tkt.bet_id = uint(id)
		tkt.sport_id = uint(id)
		tkt.matchStats = status

		_, errs := th.ticketSrv.UpdateTicket(tkt)
		if len(errs) > 0 {
			panic(errs)
		}

		http.Redirect(w, r, "/user/ticket", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/user/ticket", http.StatusSeeOther)
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

	http.Redirect(w, r, "/user/ticket", http.StatusSeeOther)
	}
