package main

import (
	"log"
	"nepseserver/constants"
	"nepseserver/database/mongodb"
	"nepseserver/database/mongodb/cronjobs"
	"nepseserver/graph"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	constants.InitConstant()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	loc := time.FixedZone("NPT", 5*60*60+45*60) // NPT is UTC+5:45
	time.Local = loc

	c := cron.New(cron.WithLocation(loc))

	cron := cronjobs.NewCronJob(c)
	go cron.InitScheduler()
	mongoClient := mongodb.Init()
	if mongoClient == nil {
		log.Fatal("Failed to initialize MongoDB client")
	}

	// ✅ Pass MongoDB client to Resolver
	resolver := &graph.Resolver{
		MongoClient: mongoClient.Database("nepsedata"),
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))

}
