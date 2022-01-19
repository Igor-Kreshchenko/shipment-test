package repositories

import (
	"github.com/Igor-Kreshchenko/shipment-test/models"
	"gorm.io/gorm"
)

type ShipmentRepository interface {
	GetAllShipments() ([]models.Shipment, error)
	CreateNewShipment(shipment *models.Shipment) (*models.Shipment, error)
	GetShipmentByID(id uint) (*models.Shipment, error)
}

type shipmentRepository struct {
	db *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) ShipmentRepository {
	return &shipmentRepository{db: db}
}

func (r *shipmentRepository) GetAllShipments() ([]models.Shipment, error) {
	var shipments []models.Shipment
	res := r.db.Find(&shipments)

	return shipments, res.Error
}

func (r *shipmentRepository) CreateNewShipment(shipment *models.Shipment) (*models.Shipment, error) {
	res := r.db.Create(&shipment)

	return shipment, res.Error
}

func (r *shipmentRepository) GetShipmentByID(id uint) (*models.Shipment, error) {
	var shipment *models.Shipment
	res := r.db.First(&shipment, id)

	return shipment, res.Error
}