package handlers

import (
	"net/http"

	"github.com/ecoprohcm/DMS_BackendServer/models"
	"github.com/ecoprohcm/DMS_BackendServer/utils"
	"github.com/gin-gonic/gin"
)

type AreaHandler struct {
	svc *models.AreaSvc
}

func NewAreaHandler(svc *models.AreaSvc) *AreaHandler {
	return &AreaHandler{
		svc: svc,
	}
}

// Find all areas and doorlocks info
// @Summary Find All Area
// @Schemes
// @Description find all areas info
// @Produce json
// @Success 200 {array} []models.Area
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/areas [get]
func (h *AreaHandler) FindAllArea(c *gin.Context) {
	aList, err := h.svc.FindAllArea(c)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Get all areas failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, aList)
}

// Find area and doorlock info by id
// @Summary Find Area By ID
// @Schemes
// @Description find area and doorlock info by area id
// @Produce json
// @Param        id	path	string	true	"Area ID"
// @Success 200 {object} models.Area
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/area/{id} [get]
func (h *AreaHandler) FindAreaByID(c *gin.Context) {
	id := c.Param("id")

	a, err := h.svc.FindAreaByID(c, id)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Get area failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, a)
}

// Create area
// @Summary Create Area
// @Schemes
// @Description Create area
// @Accept  json
// @Produce json
// @Param	data	body	models.SwagCreateArea	true	"Fields need to create a area"
// @Success 200 {object} models.Area
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/area [post]
func (h *AreaHandler) CreateArea(c *gin.Context) {
	a := &models.Area{}
	err := c.ShouldBind(a)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Invalid req body",
			ErrorMsg:   err.Error(),
		})
		return
	}
	a, err = h.svc.CreateArea(a, c.Request.Context())
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Create area failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, a)
}

// Update area
// @Summary Update Area By ID
// @Schemes
// @Description Update area, must have "id" field
// @Accept  json
// @Produce json
// @Param	data	body	models.SwagUpdateArea	true	"Fields need to update a area"
// @Success 200 {boolean} true
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/area [patch]
func (h *AreaHandler) UpdateArea(c *gin.Context) {
	a := &models.Area{}
	err := c.ShouldBind(a)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Invalid req body",
			ErrorMsg:   err.Error(),
		})
		return
	}
	isSuccess, err := h.svc.UpdateArea(c.Request.Context(), a)
	if err != nil || !isSuccess {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Update area failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, isSuccess)
}

// Delete area
// @Summary Delete Area By ID
// @Schemes
// @Description Delete area using "id" field
// @Accept  json
// @Produce json
// @Param	data	body	object{id=int}	true	"Area ID"
// @Success 200 {boolean} true
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/area [delete]
func (h *AreaHandler) DeleteArea(c *gin.Context) {
	a := &models.Area{}
	err := c.ShouldBind(a)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Invalid req body",
			ErrorMsg:   err.Error(),
		})
		return
	}

	isSuccess, err := h.svc.DeleteArea(c.Request.Context(), a)
	if err != nil || !isSuccess {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Delete area failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, isSuccess)

}
