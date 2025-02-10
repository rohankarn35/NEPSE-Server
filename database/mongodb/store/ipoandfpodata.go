package store

import (
	"context"
	"log"
	dbmodels "nepseserver/database/models"
	"nepseserver/server"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StoreIpoandFpoData(collection *mongo.Collection) {

	ipoData, err := server.GetIPOAlert("IPO")
	if err != nil {
		log.Fatalf("Failed to get IPO data: %v", err)
	}

	fpoData, err := server.GetIPOAlert("FPO")
	if err != nil {
		log.Fatalf("Failed to get FPO data: %v", err)
	}
	var ipoAlerts []dbmodels.IPOAlert
	for _, ipo := range ipoData {
		ipoAlerts = append(ipoAlerts, dbmodels.IPOAlert{
			IssuerName:             ipo.CompanyName,
			StockSymbol:            ipo.StockSymbol,
			ShareRegistrar:         ipo.ShareRegistrar,
			IndustrySector:         ipo.SectorName,
			ShareType:              ipo.ShareType,
			UnitPrice:              ipo.PricePerUnit,
			Rating:                 ipo.Rating,
			NumberOfShares:         ipo.Units,
			MinimumUnits:           ipo.MinUnits,
			MaximumUnits:           ipo.MaxUnits,
			TotalApplicationAmount: ipo.TotalAmount,
			ApplicationStartDateAD: ipo.OpeningDateAD,
			ApplicationStartDateBS: ipo.OpeningDateBS,
			ApplicationEndDateAD:   ipo.ClosingDateAD,
			ApplicationEndDateBS:   ipo.ClosingDateBS,
			ApplicationClosingTime: ipo.ClosingDateClosingTime,
			Status:                 ipo.Status,
		})
	}

	var fpoAlerts []dbmodels.IPOAlert
	for _, fpo := range fpoData {
		fpoAlerts = append(fpoAlerts, dbmodels.IPOAlert{
			IssuerName:             fpo.CompanyName,
			StockSymbol:            fpo.StockSymbol,
			ShareRegistrar:         fpo.ShareRegistrar,
			IndustrySector:         fpo.SectorName,
			ShareType:              fpo.ShareType,
			UnitPrice:              fpo.PricePerUnit,
			Rating:                 fpo.Rating,
			NumberOfShares:         fpo.Units,
			MinimumUnits:           fpo.MinUnits,
			MaximumUnits:           fpo.MaxUnits,
			TotalApplicationAmount: fpo.TotalAmount,
			ApplicationStartDateAD: fpo.OpeningDateAD,
			ApplicationStartDateBS: fpo.OpeningDateBS,
			ApplicationEndDateAD:   fpo.ClosingDateAD,
			ApplicationEndDateBS:   fpo.ClosingDateBS,
			ApplicationClosingTime: fpo.ClosingDateClosingTime,
			Status:                 fpo.Status,
		})
	}

	totalData := dbmodels.IPOAndFpoAlert{
		Ipo: ipoAlerts,
		Fpo: fpoAlerts,
	}

	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalf("Failed to count documents: %v", err)
	}

	if count == 0 {
		_, err := collection.InsertOne(context.TODO(), totalData)
		if err != nil {
			log.Printf("Failed to insert data: %v", err)
		}
	} else {
		_, err := collection.ReplaceOne(context.TODO(), bson.M{}, totalData)
		if err != nil {
			log.Printf("Failed to update data: %v", err)
		}
	}
	log.Print("fpo data updated sucessfully")
}
