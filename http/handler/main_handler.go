package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Nahom7wos/Airlines-Booking-System/book"
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
	"github.com/Nahom7wos/Airlines-Booking-System/user"
)

type MainHandler struct {
	tmpl           *template.Template
	destinationSrv flight.DestinationService
	flightSrv      flight.FlightService
	ticketSrv      book.TicketService
	userSrv        user.UserService
}

// NewMainHandler initializes and returns new MainHandler
func NewMainHandler(T *template.Template, dSrv flight.DestinationService, fSrv flight.FlightService, tSrv book.TicketService, uSrv user.UserService) *MainHandler {
	return &MainHandler{tmpl: T, destinationSrv: dSrv, flightSrv: fSrv, ticketSrv: tSrv, userSrv: uSrv}
}

func (mh *MainHandler) Index(w http.ResponseWriter, r *http.Request) {
	destinations, errs := mh.destinationSrv.Destinations()
	if errs != nil {
		panic(errs)
	}
	mh.tmpl.ExecuteTemplate(w, "index.layout", destinations)
}
func (mh *MainHandler) Admin(w http.ResponseWriter, r *http.Request) {
	flights, errs := mh.flightSrv.FlightInfo()
	if errs != nil {
		panic(errs)
	}
	mh.tmpl.ExecuteTemplate(w, "admin.index.layout", flights)
}

// Book saves ticket info on post or shows flight destination on get
func (mh *MainHandler) Book(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tkt := &entity.Ticket{}
		lyl := &entity.Loyalty{}
		usr := &entity.User{}
		getUsr := &entity.User{}
		fltID, _ := strconv.Atoi(r.FormValue("flightDestination"))

		usr.FullName = r.FormValue("fullName")
		usr.Email = r.FormValue("email")
		usr.Passport = r.FormValue("passport")
		getUsr, errs := mh.userSrv.StoreUser(usr)
		if len(errs) > 0 {
			panic(errs)
		}

		tkt.FlightID = uint(fltID)
		tkt.UserID = getUsr.ID

		_, errs := mh.ticketSrv.StoreTicket(tkt)
		if len(errs) > 0 {
			panic(errs)
		}
		// if loyalty by getUsr.ID exists
			// update
		// else store

		_, errs := mh.loyaltySrv.StoreLoyalty(lyt)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w, r, "/book", http.StatusSeeOther)

	} else {
		fltDstn, errs := mh.flightSrv.FlightDestination()
		if errs != nil {
			panic(errs)
		}
		mh.tmpl.ExecuteTemplate(w, "book.layout", fltDstn)
	}
}

func (mh *MainHandler) Checkin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		status := true
		fullName := r.FormValue("fullName")
		passport := r.FormValue("passport")
		// get user by passport and name
		usr, errs := mh.userSrv.UserByFullNameAndPassport(fullName, passport)
		if errs != nil {
			panic(errs)
		}
		// if found then get ticket by userid
		tkt, errs := mh.ticketSrv.TicketByUserId(usr.ID)
		if errs != nil {
			panic(errs)
		}
		// update TicketStatus to false
		_, errs := mh.ticketSrv.UpdateTicketStatus(tkt, status)
		if len(errs) > 0 {
			panic(errs)
		}
		// redirect
	}
	mh.tmpl.ExecuteTemplate(w, "checkin.layout", nil)
}

func (mh *MainaHandler) MyFlight(w http.ResponseWriter, r *http.Request) {
	//userid
	myFlts, errs := mh.flightSrv.MyFlight(1)
	if errs != nil {
		panic(errs)
	}
	mh.tmpl.ExecuteTemplate(w, "myflight.layout", myFlts)
}

func (mh *MainHandler) Loyalty(w http.ResponseWriter, r *http.Request) {
	// get username
	// UserID and join with loyalty
	// get loyalty then send

	mh.tmpl.ExecuteTemplate(w, "loyalty.layout", nil)
}
