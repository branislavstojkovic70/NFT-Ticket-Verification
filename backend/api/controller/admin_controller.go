package controller

import (
	"net/http"

	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/branislavstojkovic70/nft-ticket-verification/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AdminController struct {
	adminService service.AdminService
}

func NewAdminController(adminService service.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
}

func (ac *AdminController) GetAllAdmins(c *gin.Context) {
	admins, err := ac.adminService.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch admins"})
		return
	}
	c.JSON(http.StatusOK, admins)
}

func (ac *AdminController) GetAdminByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	admin, err := ac.adminService.GetAdminByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get admin"})
		return
	}
	if admin == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func (ac *AdminController) CreateAdmin(c *gin.Context) {
	var admin users.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	admin.ID = uuid.New()
	if err := ac.adminService.CreateAdmin(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
		return
	}

	c.JSON(http.StatusCreated, admin)
}

func (ac *AdminController) UpdateAdmin(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	var admin users.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	admin.ID = id
	if err := ac.adminService.UpdateAdmin(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin"})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func (ac *AdminController) DeleteAdmin(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	if err := ac.adminService.DeleteAdmin(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete admin"})
		return
	}

	c.Status(http.StatusNoContent)
}