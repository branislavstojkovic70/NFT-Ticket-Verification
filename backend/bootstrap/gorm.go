package bootstrap

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	events "github.com/branislavstojkovic70/nft-ticket-verification/domain/events"
	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection(hostDb, userDb, passDb, nameDb, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Belgrade",
		hostDb, userDb, passDb, nameDb, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&users.User{}, &users.Admin{}, &users.Organizer{}, &events.Ticket{}, &events.Event{}, &events.Tag{})
	if err != nil {
		return err
	}
	return nil
}

func InitDB(hostDb string, userDb string, passDb string, nameDb string, port string) (*gorm.DB, error) {
	db, err := CreateConnection(hostDb, userDb, passDb, nameDb, port)
	if err != nil {
		return nil, err
	}
	err = CreateMigration(db)
	if err != nil {
		return nil, err
	}
	err = SeedTestData(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SeedTestData(db *gorm.DB) error {
	// Tags
	tags := []string{"technology", "music", "festival"}
	tagsJSON, _ := json.Marshal(tags)

	// Interests
	interests := []string{"blockchain", "music", "sports"}
	interestsJSON, _ := json.Marshal(interests)

	// Hash passwords
	hash := func(password string) string {
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}
		return string(hashed)
	}

	// Users
	user1 := users.User{
		ID:        uuid.New(),
		Email:     "user1@example.com",
		Password:  hash("password1"),
		Wallet:    "wallet_1",
		Age:       25,
		Location:  "Belgrade",
		Interests: interestsJSON,
		Gender:    users.Male,
		Name:      "Petar",
		Surname:   "Petrović",
		Role:      users.RoleUser,
	}
	if err := db.Create(&user1).Error; err != nil {
		log.Println("Failed to create user1:", err)
		return err
	}

	user2 := users.User{
		ID:        uuid.New(),
		Email:     "user2@example.com",
		Password:  hash("password2"),
		Wallet:    "wallet_2",
		Age:       30,
		Location:  "Novi Sad",
		Interests: interestsJSON,
		Gender:    users.Female,
		Name:      "Milica",
		Surname:   "Milić",
		Role:      users.RoleUser,
	}
	if err := db.Create(&user2).Error; err != nil {
		log.Println("Failed to create user2:", err)
		return err
	}

	// Organizer
	organizer := users.Organizer{
		ID:       uuid.New(),
		Email:    "organizer@example.com",
		Password: hash("organizer123"),
		Wallet:   "wallet_org",
		Name:     "Ognjen",
		Surname:  "Organizator",
		Gender:   users.Male,
		Role:     users.RoleOrganizer,
	}
	if err := db.Create(&organizer).Error; err != nil {
		log.Println("Failed to create organizer:", err)
		return err
	}

	// Events
	event1 := events.Event{
		ID:          uuid.New(),
		Location:    "Belgrade",
		Type:        events.Music,
		DateStart:   time.Now().AddDate(0, 1, 0),
		DateEnd:     time.Now().AddDate(0, 1, 1),
		Description: "Summer music festival",
		Title:       "Belgrade Beats",
		Tags:        tagsJSON,
		OrganizerID: organizer.ID,
		NumberOfTickets: 13,
	}
	if err := db.Create(&event1).Error; err != nil {
		log.Println("Failed to create event1:", err)
		return err
	}

	event2 := events.Event{
		ID:          uuid.New(),
		Location:    "Novi Sad",
		Type:        events.Conference,
		DateStart:   time.Now().AddDate(0, 2, 0),
		DateEnd:     time.Now().AddDate(0, 2, 2),
		Description: "Tech conference",
		Title:       "NS Tech 2025",
		Tags:        tagsJSON,
		OrganizerID: organizer.ID,
		NumberOfTickets: 15,
	}
	if err := db.Create(&event2).Error; err != nil {
		log.Println("Failed to create event2:", err)
		return err
	}

	// Admin
	admin := users.Admin{
		ID:       uuid.New(),
		Email:    "admin@example.com",
		Password: hash("admin123"),
		Wallet:   "wallet_admin",
		Name:     "Ana",
		Surname:  "Adminović",
		Gender:   users.Female,
		Role:     users.RoleAdmin,
	}
	if err := db.Create(&admin).Error; err != nil {
		log.Println("Failed to create admin:", err)
		return err
	}

	// Tickets
	ticket1 := events.Ticket{
		ID:        uuid.New(),
		UserID:    user1.ID,
		EventID:   event1.ID,
		Price:     100.0,
		DateStart: event1.DateStart,
		DateEnd:   event1.DateEnd,
		IsUsed:    false,
	}
	if err := db.Create(&ticket1).Error; err != nil {
		log.Println("Failed to create ticket1:", err)
		return err
	}

	ticket2 := events.Ticket{
		ID:        uuid.New(),
		UserID:    user2.ID,
		EventID:   event2.ID,
		Price:     150.0,
		DateStart: event2.DateStart,
		DateEnd:   event2.DateEnd,
		IsUsed:    true,
	}
	if err := db.Create(&ticket2).Error; err != nil {
		log.Println("Failed to create ticket2:", err)
		return err
	}

	log.Println("Test data seeded successfully.")
	return nil
}
