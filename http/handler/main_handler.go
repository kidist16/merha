package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/kidist16/sport-betting/entity"
	"github.com/kidist16/sport-betting/user"
)

type MainHandler struct {
	tmpl           *template.Template
	ticketSrv      book.TicketService
	userSrv        user.User
}s

// NewMainHandler initializes and returns new MainHandler
func NewMainHandler(T *template.Template, dSrv flight.DestinationService, fSrv flight.FlightService, tSrv book.TicketService, uSrv user.UserService) *MainHandler {
	return &MainHandler{tmpl: T, destinationSrv: dSrv, flightSrv: fSrv, ticketSrv: tSrv, userSrv: uSrv}
}

func (mh *MainHandler) Index(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "index.html", index)
}
func (mh *MainHandler) Admin(w http.ResponseWriter, r *http.Request) {
	flights, errs := mh.flightSrv.FlightInfo()
	if errs != nil {
		panic(errs)
	}
	mh.tmpl.ExecuteTemplate(w, "adminindex.html", flights)
}

// Book saves ticket info on post or shows flight destination on get
func (mh *MainHandler) index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tkt := &entity.Ticket{}
		usr := &entity.User{}
		getUsr := &entity.User{}
		mch, _ := strconv.Atoi(r.FormValue("matchid"))

		mch.home = r.FormValue("home")
		mch.away = r.FormValue("away")
		mch.tie = r.FormValue("tie")
		getUsr, errs := mh.userSrv.StoreUser(usr)
		if len(errs) > 0 {
			panic(errs)
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)

	} else {
		fltDstn, errs := mh.flightSrv.FlightDestination()
		if errs != nil {
			panic(errs)
		}
		mh.tmpl.ExecuteTemplate(w, "index.layout", fltDstn)
	}
}

func (mh *MainHandler) result(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		status := true
		matchid := r.FormValue("matchid")
		matchstat := r.FormValue("matchstat")
		usr, errs := mh.userSrv.UserByFullNameAndPassport(matchid matchstat)
		if errs != nil {
			panic(errs)
		}
		tkt, errs := mh.ticketSrv.TicketByUserId(usr.ID)
		if errs != nil {
			panic(errs)
		}
	}
	mh.tmpl.ExecuteTemplate(w, "result.html", nil)
}

func (mh *MainHandler) statistics(w http.ResponseWriter, r *http.Request) {
	
	myFlts, errs := mh.flightSrv.MyFlight(1)
	if errs != nil {
		panic(errs)
	}
	mh.tmpl.ExecuteTemplate(w, "statistics.html", sta)
}
