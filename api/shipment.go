package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Igor-Kreshchenko/shipment-test/services"
	"github.com/gin-gonic/gin"
)

func InjectShipment(gr *gin.RouterGroup, shipmentService services.ShipmentService) {
	handler := gr.Group("shipments")

	handler.GET("", getAllShipments(shipmentService))
	handler.POST("", createNewShipment(shipmentService))
	handler.GET(":id", getShipmentByID(shipmentService))
}

func getAllShipments(shipmentService services.ShipmentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		shipments, err := shipmentService.GetAllShipments()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"shipments": shipments,
		})
	}
}

func createNewShipment(shipmentService services.ShipmentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var shipmentReq *services.ShipmentRequest

		err := c.BindJSON(&shipmentReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})

			return
		}

		res, err := shipmentService.CreateNewShipment(shipmentReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Shipment created",
			"price":   res.Price,
		})
	}
}

func getShipmentByID(shipmentService services.ShipmentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		uid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		shipment, err := shipmentService.GetShipmentByID(uint(uid))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"shipment": shipment,
		})
	}
}
