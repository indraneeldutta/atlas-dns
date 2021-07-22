package dnsservice

import (
	"errors"
	"math"

	"github.com/atlas-dns/common"
	"github.com/atlas-dns/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DnsService ..
type DnsService struct {
	collection *mongo.Collection
}

// IDnsService interface
type IDnsService interface {
	GetDroneLoc(*common.Context, models.DroneRequest) (models.DroneResponse, error)
}

// NewDnsService initialises new service with required values
func NewDnsService(client *mongo.Client) *DnsService {
	return &DnsService{
		collection: client.Database(viper.GetString("MONGO.DB_NAME")).Collection("drones"),
	}
}

// GetDroneLoc calculates the drone loc to be sent back as response
func (s *DnsService) GetDroneLoc(ctx *common.Context, request models.DroneRequest) (models.DroneResponse, error) {
	var droneDetails models.DroneDetails
	var response models.DroneResponse
	err := s.collection.FindOne(ctx.Ctx, bson.M{"_id": request.ID}).Decode(&droneDetails)
	if errors.Is(err, nil) {
		x, err := common.StringToFloat64(ctx, request.X)
		if !errors.Is(err, nil) {
			ctx.Logger.Error(err)
			return response, err
		}
		y, err := common.StringToFloat64(ctx, request.Y)
		if !errors.Is(err, nil) {
			ctx.Logger.Error(err)
			return response, err
		}
		z, err := common.StringToFloat64(ctx, request.Z)
		if !errors.Is(err, nil) {
			ctx.Logger.Error(err)
			return response, err
		}
		vel, err := common.StringToFloat64(ctx, request.Vel)
		if !errors.Is(err, nil) {
			ctx.Logger.Error(err)
			return response, err
		}

		loc := math.Round(((x*float64(droneDetails.SectorID))+(y*float64(droneDetails.SectorID))+(z*float64(droneDetails.SectorID))+vel)*100) / 100

		response.Loc = loc
		return response, nil
	}

	ctx.Logger.Error(err)
	return response, err
}
