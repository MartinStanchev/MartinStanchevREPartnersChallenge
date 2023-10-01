package main

import (
	"net/http"
	"os"
	"packDistributor/pkg/packDistributor"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type application struct {
	distributor *packDistributor.PackDistributor
	log         *logrus.Logger
}

func main() {
	log := logrus.New()

	packSizes := os.Getenv("PACK_SIZES")
	if packSizes == "" {
		packSizes = "250,500,1000,2000,5000"
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = ":8080"
	}

	sizes, err := parsePackSizes(packSizes)
	if err != nil {
		log.WithError(err).Fatal("Could not parse pack sizes")
	}

	distributor := packDistributor.NewDistributor(sizes)

	app := &application{
		distributor: distributor,
		log:         log,
	}

	servMux := initRoutes(app)

	log.Info("Starting server on port: ", appPort)
	http.ListenAndServe(appPort, servMux)
}

func parsePackSizes(packSizes string) ([]int, error) {
	splitPacks := strings.Split(packSizes, ",")
	sizes := make([]int, len(splitPacks))

	var err error
	for i, s := range splitPacks {
		sizes[i], err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
	}
	return sizes, nil
}

func initRoutes(app *application) *mux.Router {
	servMux := mux.NewRouter()

	servMux.HandleFunc("/order", app.handleOrder).Methods("POST")
	servMux.HandleFunc("/update/sizes", app.updatePackSizes).Methods("PATCH")

	return servMux
}
