package controller

import (
	"net/http"
	"time"

	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

type LoginRequest struct {
	Email    string     `json:"email" binding:"required,email"`
	Password string     `json:"password" binding:"required"`
	Role     users.Role `json:"role" binding:"required"`
}

func (a *AuthController) Login(db *gorm.DB, c *gin.Context, jwtSecret string) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nevalidan unos"})
		return
	}

	var (
		userID         uuid.UUID
		hashedPassword string
		role           users.Role
	)

	switch req.Role {
	case users.RoleUser:
		var user users.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Neispravan email ili lozinka"})
			return
		}
		userID = user.ID
		hashedPassword = user.Password
		role = user.Role

	case users.RoleOrganizer:
		var organizer users.Organizer
		if err := db.Where("email = ?", req.Email).First(&organizer).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Neispravan email ili lozinka"})
			return
		}
		userID = organizer.ID
		hashedPassword = organizer.Password
		role = organizer.Role

	case users.RoleAdmin:
		var admin users.Admin
		if err := db.Where("email = ?", req.Email).First(&admin).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Neispravan email ili lozinka"})
			return
		}
		userID = admin.ID
		hashedPassword = admin.Password
		role = admin.Role

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nepoznata uloga"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Neispravan email ili lozinka"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID.String(),
		"role":    string(role),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Greška pri generisanju tokena"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"role":  role,
		"id":    userID,
	})
}
