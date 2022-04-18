package confs

import (
	"api/models"
	"gorm.io/gorm"
)

func MigrateUsersOrgs(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Organization{}, &models.OrgInvit{})
}
