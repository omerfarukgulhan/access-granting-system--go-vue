package postgresql

import (
	"access-granting/common/security"
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

	err = db.Exec(`
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'fk_user') THEN
                ALTER TABLE user_roles ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
            END IF;
        END$$;
    `).Error
	if err != nil {
		log.Fatalf("Failed to add foreign key constraint for user_id: %v", err)
	}

	err = db.Exec(`
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'fk_role') THEN
                ALTER TABLE user_roles ADD CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE;
            END IF;
        END$$;
    `).Error
	if err != nil {
		log.Fatalf("Failed to add foreign key constraint for role_id: %v", err)
	}

	log.Println("Tables migrated and foreign key constraints added successfully.")
	SeedData(db)
}

func SeedData(db *gorm.DB) {
	role := entities.Role{Name: "Admin"}
	if err := db.Where("name = ?", role.Name).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			role.Id = 1
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

	hashedPassword, _ := security.HashPassword("P4ssword", 10)
	user := entities.User{
		Id:           1,
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
