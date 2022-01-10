package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AlexTsIvanov/OrderService/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "kitchen-api", log.LstdFlags)

	mh := handlers.NewMenu(l)
	oh := handlers.NewOrder(l)
	uh := handlers.NewUser(l)

	sm := mux.NewRouter()

	getRouterUnAuth := sm.Methods(http.MethodGet).Subrouter()
	getRouterUnAuth.HandleFunc("/api/menu", mh.GetMenu)
	getRouterUnAuth.HandleFunc("/api/menu/{id:[0-9]+}", mh.GetMenuItem)
	getRouterUnAuth.HandleFunc("/api/menu/types", mh.GetMenuTypes)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/orders", oh.GetOrders)
	getRouter.HandleFunc("/api/orders/{id:[0-9]+}", oh.GetOrderByID)
	getRouter.HandleFunc("/api/orders/orderitems", oh.GetOrderItems)
	getRouter.HandleFunc("/api/orders/orderitems/{id:[0-9]+}", oh.GetOrderItemsByOrderID)
	getRouter.HandleFunc("/api/orders/orderitems/user", oh.GetOrderItemsByCustomerID)
	getRouter.Use(uh.IsAuthorized)

	postRouterUnAuth := sm.Methods(http.MethodPost).Subrouter()
	postRouterUnAuth.HandleFunc("/api/users/signup", uh.SignUp)
	postRouterUnAuth.HandleFunc("/api/users/signin", uh.SignIn)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/menu", mh.PostMenu)
	postRouter.HandleFunc("/api/orders", oh.PostOrders)
	postRouter.HandleFunc("/api/orders/orderitems", oh.PostOrderItems)
	postRouter.Use(uh.IsAuthorized)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/api/menu/{id:[0-9]+}", mh.UpdateMenu)
	putRouter.HandleFunc("/api/orders/orderitems/user/{itemId:[0-9]+}", oh.PutOrderItemIDByCustomerID)
	putRouter.Use(uh.IsAuthorized)

	delRouter := sm.Methods(http.MethodDelete).Subrouter()
	delRouter.HandleFunc("/api/menu/{id:[0-9]+}", mh.DeleteMenu)
	delRouter.Use(uh.IsAuthorized)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
