package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB
var tpl *template.Template

func init() {
	db, err := sql.Open("postgres", "postgres://bond:password@localhost/sportBetting?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

}

type sport struct {
	sportid   string
	leagueid  string
	matchid   string
	betid     string
	priceid   float32
	sportname string
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)
}

func adminindex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM sportBetting")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()
	bks := make([]sport, 0)
	for rows.Next() {
		bk := sport{}
		err := row.Scan(&bk.sportid, &bk.leagueid, &bk.matchid, &bk.betid, &bk.priceid)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, %s, %s, %s", bk.sportid, bk.leagueid, bk.matchid, bk.betid, bk.priceid)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

func adminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM sportBetting")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]sport, 0)
	for rows.Next() {
		bk := sport{}
		err := row.Scan(&bk.sportid, &bk.leagueid, &bk.matchid, &bk.betid, &bk.priceid)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl.ExecuteTemplate(w, "in-play.html", bks)
}

func adminShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	sportid := r.FormValue("sportid")
	if sportid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM sportBetting WHERE sportid = $1", sportid)

	bk := sport{}
	err := row.Scan(&bk.sportid, &bk.leagueid, &bk.matchid, &bk.betid, &bk.priceid)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "adminshow.html", bk)
}

func adminCreateForm(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "admincreate.html", nil)
}

func adminCreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	bk := sport{}
	bk.sportid = r.FormValue("sportid")
	bk.leagueid = r.FormValue("leagueid")
	bk.matchid = r.FormValue("matchid")
	bk.betid = r.FormValue("betid")
	p := r.FormValue("")

	// validate form values
	if bk.sportid == "" || bk.leagueid == "" || bk.matchid == "" || bk.betid == "" || bk.priceid == 32 {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.priceid = float32(f64)

	// insert values
	_, err = db.Exec("INSERT INTO sportBeting (sportid, leagueid, matchid, priceid) VALUES ($1, $2, $3, $4)", bk.sportid, bk.leagueid, bk.matchid, bk.priceid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "admincreated.html", bk)
}

func adminUpdateForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	sportid := r.FormValue("sportid")
	if sportid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM sportBetting WHERE ID = $1", sportid)

	bk := sport{}
	err := row.Scan(&bk.sportid, &bk.leagueid, &bk.matchid, &bk.betid, &bk.priceid)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "adminupdate.html", bk)
}

func booksUpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	bk := sport{}
	bk.sportid = r.FormValue("sportid")
	bk.leagueid = r.FormValue("leagueid")
	bk.matchid = r.FormValue("matchid")
	bk.betid = r.FormValue("betid")
	p := r.FormValue("price")

	// validate form values
	if bk.sportid == "" || bk.leagueid == "" || bk.matchid == "" || bk.betid == "" || p == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.priceid = float32(f64)

	// insert values
	_, err = db.Exec("UPDATE sportBetting SET sportid = $1, leagueid=$2, matchid=$3, price=$4 WHERE sportid=$1;", bk.sportid, bk.leagueid, bk.matchid, bk.priceid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "adminupdated.html", bk)
}

func adminDeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	matchid := r.FormValue("matchid")
	if matchid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// delete book
	_, err := db.Exec("DELETE FROM sportBetting WHERE matchid=$1;", matchid)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
