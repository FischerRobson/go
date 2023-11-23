package main

import (
	"encoding/json"
	"imersao-fullcycle/internal/entity"
	"net/http"

	_ "github.com/go-chi/chi/middleware"
	_ "github.com/go-chi/chi/v5"
	"github.com/labstack/echo/v4"
)

func main() {
	// Vanilla
	// http.ListenAndServe(":8888", nil)
	// http.HandleFunc("/order", OrderHandler)

	// Using router
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", OrderHandler)
	// http.ListenAndServe(":8888", r)

	// Using Echo framework
	e := echo.New()
	e.GET("/order", OrderHandlerEcho)
	e.Logger.Fatal(e.Start(":8888"))
}

func OrderHandlerEcho(e echo.Context) error {
	order := entity.Order{
		ID:    "1",
		Price: 10,
		Tax:   10,
	}
	err := order.CalculateFinalPrice()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusCreated, order)
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	order := entity.Order{
		ID:    "1",
		Price: 10,
		Tax:   10,
	}
	err := order.CalculateFinalPrice()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(order)
}
