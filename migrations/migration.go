package migrations

import (
	"gin-boilerplate/infra/database"
	"gin-boilerplate/models"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var migrationModels = []interface{}{&models.Example{}, &models.Category{},
		&models.Provider{}, &models.ProductDeal{},
		&models.ProviderProduction{}, &models.ProviderRule{}, &models.DealRule{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
