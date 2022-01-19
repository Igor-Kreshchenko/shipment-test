package main

import (
	"github.com/Igor-Kreshchenko/shipment-test/api"
	"github.com/Igor-Kreshchenko/shipment-test/repositories"
	"github.com/Igor-Kreshchenko/shipment-test/services"
	"github.com/Igor-Kreshchenko/shipment-test/setup"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := setup.ConnectDataBase()
	if err != nil {
		panic(err)
	}

	gr := router.Group("v1/api")

	shipmentRepository := repositories.NewShipmentRepository(db)
	shipmentService := services.NewShipmentService(shipmentRepository)

	api.InjectShipment(gr, shipmentService)

	router.Run(":8080")
}
