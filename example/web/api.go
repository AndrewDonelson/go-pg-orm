package web


import (
	goPg "github.com/AndrewDonelson/go-pg-orm"

	"github.com/AndrewDonelson/example/web/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer(db *goPg.Model) {
	router := mux.NewRouter()

	handler := controllers.New(db)

	router.HandleFunc("/", handler.List).Methods(http.MethodGet)
	router.HandleFunc("/view/{id}", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/delete/{id}", handler.Delete).Methods(http.MethodGet)
	router.HandleFunc("/create", handler.Create).Methods(http.MethodPost)
	router.HandleFunc("/update/{id}", handler.Update).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}