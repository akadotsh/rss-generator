package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func  main()  {

	godotenv.Load(".env")
	portString:= os.Getenv("PORT");
	if portString ==""{
		log.Fatal("PORT must be set")
	}

     router:= chi.NewRouter();
    
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}));

	v1Router := chi.NewRouter();

	v1Router.Get("/ready",handlerRediness)
    v1Router.Get("/err",handlerErr)
	
	router.Mount("/v1",v1Router)

	 srv:= &http.Server{
		Handler: router,
		Addr: ":" + portString,
	 }
	 fmt.Println("Port:",portString);
	 
	 err:= srv.ListenAndServe()
	 fmt.Printf("Server starting on port %v",portString);
	if err!= nil {
		log.Fatal(err)
	}

}