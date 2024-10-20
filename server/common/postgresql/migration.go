package postgresql

import (
	"access-granting/common/security"
	"access-granting/domain/entities"
	"log"

	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) {
	err := db.Migrator().DropTable(&entities.User{}, &entities.Role{}, &entities.UserRole{})
	if err != nil {
		log.Fatalf("Failed to drop existing tables: %v", err)
	}
	log.Println("Existing tables dropped successfully.")

	err = db.AutoMigrate(&entities.User{}, &entities.Role{}, &entities.UserRole{})
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
	roles := []entities.Role{
		{Id: 1, Name: "Admin"},
		{Id: 2, Name: "Editor"},
		{Id: 3, Name: "Author"},
		{Id: 4, Name: "Moderator"},
		{Id: 5, Name: "Viewer"},
	}
	if err := db.Create(&roles).Error; err != nil {
		log.Fatalf("Failed to seed roles: %v", err)
	}
	log.Println("Roles seeded successfully.")

	hashedPassword, _ := security.HashPassword("P4ssword", 10)
	users := []entities.User{
		{
			Id:           1,
			Username:     "Omer Faruk Gulhan",
			Email:        "omer@gulhan.com",
			Password:     hashedPassword,
			ProfileImage: "default.png",
			IsActive:     true,
		},
		{
			Id:           2,
			Username:     "John Doe",
			Email:        "john@doe.com",
			Password:     hashedPassword,
			ProfileImage: "default.png",
			IsActive:     true,
		},
		{
			Id:           3,
			Username:     "Jane Doe",
			Email:        "jane@doe.com",
			Password:     hashedPassword,
			ProfileImage: "default.png",
			IsActive:     true,
		},
		{
			Id:           4,
			Username:     "Max Mustermann",
			Email:        "max@mustermann.com",
			Password:     hashedPassword,
			ProfileImage: "default.png",
			IsActive:     true,
		},
	}
	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}
	log.Println("Users seeded successfully.")

	adminRole := entities.UserRole{
		UserId: users[0].Id,
		RoleId: roles[0].Id,
	}
	if err := db.Create(&adminRole).Error; err != nil {
		log.Fatalf("Failed to assign admin role to user: %v", err)
	}
	log.Println("Admin role assigned to user successfully.")
}
