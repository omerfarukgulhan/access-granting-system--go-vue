package postgresql

import (
	"access-granting/common/security"
	"access-granting/common/util/id"
	"access-granting/domain/entities"
	"errors"
	"log"

	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(&entities.User{}, &entities.Role{}, &entities.UserRole{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	log.Println("Tables migrated successfully.")
	SeedData(db)
}

func SeedData(db *gorm.DB) {
	role := entities.Role{Name: "Admin"}
	if err := db.Where("name = ?", role.Name).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			roleId, _ := id.GetUniqueId()
			role.Id = roleId
			if err := db.Create(&role).Error; err != nil {
				log.Fatalf("Failed to seed roles: %v", err)
			}
			log.Println("Roles seeded successfully.")
		} else {
			log.Fatalf("Failed to check existing role: %v", err)
		}
	} else {
		log.Println("Role already exists. No need to seed.")
	}

	userId, _ := id.GetUniqueId()
	hashedPassword, _ := security.HashPassword("P4ssword", 10)
	user := entities.User{
		Id:           userId,
		Username:     "username1",
		Email:        "omer@omer.com",
		Password:     hashedPassword,
		ProfileImage: "default.png",
		IsActive:     true,
	}
	var existingUser entities.User
	if err := db.Where("email = ? OR username = ?", user.Email, user.Username).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&user).Error; err != nil {
				log.Fatalf("Failed to seed admin user: %v", err)
			}
			log.Println("Admin user seeded successfully.")
		} else {
			log.Fatalf("Failed to check existing user: %v", err)
		}
	} else {
		log.Println("User already exists. No need to seed.")
		return
	}

	adminRole := entities.UserRole{
		UserId: user.Id,
		RoleId: role.Id,
	}
	if err := db.Where("user_id = ? AND role_id = ?", adminRole.UserId, adminRole.RoleId).First(&adminRole).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&adminRole).Error; err != nil {
				log.Fatalf("Failed to assign admin role to user: %v", err)
			}
			log.Println("Admin role assigned to user successfully.")
		} else {
			log.Fatalf("Failed to check existing user-role association: %v", err)
		}
	} else {
		log.Println("User role already exists. No need to assign.")
	}
}
