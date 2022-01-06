package handlers

import (
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ecoprohcm/DMS_BackendServer/models"
	"github.com/ecoprohcm/DMS_BackendServer/utils"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	svc  *models.EmployeeSvc
	mqtt mqtt.Client
}

func NewEmployeeHandler(svc *models.EmployeeSvc, mqtt mqtt.Client) *EmployeeHandler {
	return &EmployeeHandler{
		svc:  svc,
		mqtt: mqtt,
	}
}

// Find all employees info
// @Summary Find All Employee
// @Schemes
// @Description find all employees info
// @Produce json
// @Success 200 {array} []models.Employee
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/employees [get]
func (h *EmployeeHandler) FindAllEmployee(c *gin.Context) {
	eList, err := h.svc.FindAllEmployee(c)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Get all employees failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, eList)
}

// Find employee info by id
// @Summary Find Employee By ID
// @Schemes
// @Description find employee info by employee id
// @Produce json
// @Param        id	path	string	true	"Employee ID"
// @Success 200 {object} models.Employee
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/employee/{id} [get]
func (h *EmployeeHandler) FindEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	s, err := h.svc.FindEmployeeByID(c, id)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Get employee failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, s)
}

// Create employee
// @Summary Create Employee
// @Schemes
// @Description Create employee
// @Accept  json
// @Produce json
// @Param	data	body	models.SwagCreateEmployee	true	"Fields need to create a employee"
// @Success 200 {object} models.Employee
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/employee [post]
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	s := &models.Employee{}
	err := c.ShouldBind(s)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Invalid req body",
			ErrorMsg:   err.Error(),
		})
		return
	}

	_, err = h.svc.CreateEmployee(c.Request.Context(), s)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Create employee failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, s)
}

// Update employee
// @Summary Update Employee By ID
// @Schemes
// @Description Update employee, must have "id" field
// @Accept  json
// @Produce json
// @Param	data	body	models.SwagUpdateEmployee	true	"Fields need to update an employee"
// @Success 200 {boolean} true
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/employee [patch]
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	s := &models.Employee{}
	err := c.ShouldBind(s)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Invalid req body",
			ErrorMsg:   err.Error(),
		})
		return
	}

	isSuccess, err := h.svc.UpdateEmployee(c.Request.Context(), s)
	if err != nil || !isSuccess {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Update employee failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, isSuccess)
}

// Delete employee
// @Summary Delete Employee By ID
// @Schemes
// @Description Delete employee using "id" field
// @Accept  json
// @Produce json
// @Param	data	body	object{id=int}	true	"Employee ID"
// @Success 200 {boolean} true
// @Failure 400 {object} utils.ErrorResponse
// @Router /v1/employee [delete]
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	dId := &models.DeleteID{}
	err := c.ShouldBind(dId)
	if err != nil {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Invalid req body",
			ErrorMsg:   err.Error(),
		})
		return
	}

	isSuccess, err := h.svc.DeleteEmployee(c.Request.Context(), dId.ID)
	if err != nil || !isSuccess {
		utils.ResponseJson(c, http.StatusBadRequest, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Delete employee failed",
			ErrorMsg:   err.Error(),
		})
		return
	}
	utils.ResponseJson(c, http.StatusOK, isSuccess)

}