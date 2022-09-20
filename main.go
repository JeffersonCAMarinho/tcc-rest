package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Movies struct {
	Movieid   int
	Title     string
	Genres    string
	Userid    int
	Rating    float64
	Timestamp int64
	Tag       string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "movies"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func ListarFilmes(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	rows, err := db.Query(`
							select distinct m.movieid, m.title, m.genres, r.userid, r.rating, r.timestamp, t.tag from movies m 
							inner join ratings r on r.movieid = m.movieid 
							inner join tags t on t.movieid = m.movieid
						`)
	if err != nil {
		log.Fatal(err)
	}

	var films []Movies

	for rows.Next() {
		var movies Movies
		rows.Scan(&movies.Movieid, &movies.Title, &movies.Genres, &movies.Userid, &movies.Rating, &movies.Timestamp, &movies.Tag)
		films = append(films, movies)
	}

	filmsBytes, _ := json.MarshalIndent(films, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(filmsBytes)

	defer rows.Close()
	defer db.Close()
}

func main() {
	http.HandleFunc("/listarFilmes", ListarFilmes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
