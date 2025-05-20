package controller

import (
	"net/http"

	users "github.com/branislavstojkovic70/nft-ticket-verification/domain/users"
	"github.com/branislavstojkovic70/nft-ticket-verification/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrganizerController struct {
	organizerService service.OrganizerService
}

func NewOrganizerController(organizerService service.OrganizerService) *OrganizerController {
	return &OrganizerController{organizerService: organizerService}
}

func (oc *OrganizerController) GetAllOrganizers(c *gin.Context) {
	organizers, err := oc.organizerService.GetAllOrganizers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch organizers"})
		return
	}
	c.JSON(http.StatusOK, organizers)
}

func (oc *OrganizerController) GetOrganizerByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organizer ID"})
		return
	}

	organizer, err := oc.organizerService.GetOrganizerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get organizer"})
		return
	}
	if organizer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organizer not found"})
		return
	}

	c.JSON(http.StatusOK, organizer)
}

func (oc *OrganizerController) CreateOrganizer(c *gin.Context) {
	var organizer users.Organizer
	if err := c.ShouldBindJSON(&organizer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	organizer.ID = uuid.New()
	if err := oc.organizerService.CreateOrganizer(&organizer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organizer"})
		return
	}

	c.JSON(http.StatusCreated, organizer)
}

func (oc *OrganizerController) UpdateOrganizer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organizer ID"})
		return
	}

	var organizer users.Organizer
	if err := c.ShouldBindJSON(&organizer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	organizer.ID = id
	if err := oc.organizerService.UpdateOrganizer(&organizer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update organizer"})
		return
	}

	c.JSON(http.StatusOK, organizer)
}

func (oc *OrganizerController) DeleteOrganizer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organizer ID"})
		return
	}

	if err := oc.organizerService.DeleteOrganizer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete organizer"})
		return
	}

	c.Status(http.StatusNoContent)
}