package main

import (
	"log"
	"net/http"
	"os"
    "database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/guilhermeNatan/full_cycle/13-GraphQL/graph"
	"github.com/guilhermeNatan/full_cycle/13-GraphQL/internal/database"
    _ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
    db, err := sql.Open("sqlite3", "./data.db")
    if err != nil {
        log.Fatalf("falha ao abrir db %v ", err)
    }
    // fecha a conexao do banco quando toda a execuação do metodo main acabar
    defer db.Close()

    categoryDb := database.NewCategory(db)
     courseDb := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
	    CategoryDB: categoryDb,
	    CourseDB: courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
