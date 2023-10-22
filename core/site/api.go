package site

import (
	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/loadpoint"
	"github.com/evcc-io/evcc/core/vehicle"
)

// API is the external site API
type API interface {
	Healthy() bool
	Loadpoints() []loadpoint.API

	//
	// battery
	//

	GetBufferSoc() float64
	SetBufferSoc(float64) error
	GetBufferStartSoc() float64
	SetBufferStartSoc(float64) error
	GetPrioritySoc() float64
	SetPrioritySoc(float64) error

	//
	// power and energy
	//

	GetResidualPower() float64
	SetResidualPower(float64) error

	//
	// vehicles
	//

	// GetVehicles returns the list of vehicles
	GetVehicles() []api.Vehicle
	// VehicleSettings returns the list of vehicle setting adapters
	VehicleSettings() []vehicle.API

	//
	// tariffs and costs
	//

	// GetTariff returns the respective tariff
	GetTariff(string) api.Tariff
	GetSmartCostLimit() float64
	SetSmartCostLimit(float64) error
}
