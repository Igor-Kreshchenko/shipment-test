package services

import (
	"github.com/Igor-Kreshchenko/shipment-test/models"
	"github.com/Igor-Kreshchenko/shipment-test/repositories"
)

type ShipmentRequest struct {
	From string `json:"from" binding:"required"`
	To   string `json:"to" binding:"required"`
}

type ShipmentService interface {
	GetAllShipments() ([]models.Shipment, error)
	CreateNewShipment(shipment *ShipmentRequest) (*models.Shipment, error)
	GetShipmentByID(id uint) (*models.Shipment, error)
}

type shipmentService struct {
	shipmentRepository repositories.ShipmentRepository
}

func NewShipmentService(shipmentRepo repositories.ShipmentRepository) ShipmentService {
	return &shipmentService{shipmentRepository: shipmentRepo}
}

func (s *shipmentService) GetAllShipments() ([]models.Shipment, error) {
	res, err := s.shipmentRepository.GetAllShipments()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *shipmentService) CreateNewShipment(shipmentReq *ShipmentRequest) (*models.Shipment, error) {
	shipment := models.Shipment{
		From: shipmentReq.From,
		To:   shipmentReq.To,
	}

	res, err := s.shipmentRepository.CreateNewShipment(&shipment)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *shipmentService) GetShipmentByID(id uint) (*models.Shipment, error) {
	shipment, err := s.shipmentRepository.GetShipmentByID(id)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}
