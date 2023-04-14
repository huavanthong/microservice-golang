package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/services"
)

type DiscountController struct {
	discountService services.DiscountService
}

func NewDiscountController(discountService services.DiscountService) DiscountController {
	return DiscountController{
		discountService: discountService,
	}
}

// GetAllDiscounts godoc
// @Summary Get all discounts in repository.
// @Description Get all discounts in repository.
// @Tags discount
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Router /discounts [get]
func (c *DiscountController) GetAllDiscounts(ctx *gin.Context) {

	discounts, err := c.discountService.GetDiscounts()
	if err != nil {
		// Not found
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				models.GenericResponse{
					Success: false,
					Code:    http.StatusNotFound,
					Message: "Discount not found",
					Data:    nil,
					Errors:  []string{err.Error()},
				})
			return
		}
		// Internal error
		ctx.JSON(http.StatusInternalServerError,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusInternalServerError,
				Message: "Internal error",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	// Success
	ctx.JSON(http.StatusOK,
		models.GenericResponse{
			Success: true,
			Code:    http.StatusOK,
			Message: "Get discount success",
			Data:    discounts,
			Errors:  nil,
		})
	return
}

// GetDiscount godoc
// @Summary Get discount for product name
// @Description Get discount for product name
// @Tags discount
// @Accept  json
// @Produce  json
// @Param id path string true "Discount ID"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Router /discounts/{id} [get]
func (c *DiscountController) GetDiscount(ctx *gin.Context) {

	// get user ID from URL path
	id := ctx.Param("id")

	discount, err := c.discountService.GetDiscount(id)
	if err != nil {
		// Not found
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				models.GenericResponse{
					Success: false,
					Code:    http.StatusNotFound,
					Message: "Discount not found",
					Data:    nil,
					Errors:  []string{err.Error()},
				})
			return
		}

		// Internal error
		ctx.JSON(http.StatusInternalServerError,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusInternalServerError,
				Message: "Internal error",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	// Success
	ctx.JSON(http.StatusOK,
		models.GenericResponse{
			Success: true,
			Code:    http.StatusOK,
			Message: "Get discount success",
			Data:    discount,
			Errors:  nil,
		})
	return
}

// CreateDiscount godoc
// @Summary Create discount for product
// @Description Create discount for product
// @Tags discount
// @Accept  json
// @Produce  json
// @Param discount body models.CreateDiscountRequest true "New discount"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Router /discounts [post]
func (c *DiscountController) CreateDiscount(ctx *gin.Context) {

	var discountReq models.CreateDiscountRequest
	if err := ctx.ShouldBindJSON(&discountReq); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "Invalid data request to create discount",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	discount, err := c.discountService.CreateDiscount(&discountReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusInternalServerError,
				Message: "failed to create discount",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.GenericResponse{
			Success: true,
			Code:    http.StatusOK,
			Message: "discount created successfully",
			Data:    discount,
			Errors:  nil,
		})
	return
}

// UpdateDiscount godoc
// @Summary Update coupon
// @Description Update coupon
// @Tags discount
// @Accept  json
// @Produce  json
// @Param discount body models.UpdateDiscountRequest true "Update discount"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Router /discounts [patch]
func (c *DiscountController) UpdateDiscount(ctx *gin.Context) {

	var discountReq models.UpdateDiscountRequest
	if err := ctx.ShouldBindJSON(&discountReq); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "Invalid data request to update discount",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	discount, err := c.discountService.UpdateDiscount(&discountReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusInternalServerError,
				Message: "Internal server service error",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.GenericResponse{
			Success: true,
			Code:    http.StatusOK,
			Message: "discount updated successfully",
			Data:    discount,
			Errors:  nil,
		})
	return
}

// DeleteDiscount godoc
// @Summary Delete coupon by product name
// @Description Delete coupon by product name
// @Tags discount
// @Accept  json
// @Produce  json
// @Param id path string true "Discount ID"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Router /discounts/{id}} [delete]
func (c *DiscountController) DeleteDiscount(ctx *gin.Context) {

	// get user ID from URL path
	id := ctx.Param("id")

	// call discount service to find discount by ID
	err := c.discountService.DeleteDiscount(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			models.GenericResponse{
				Success: false,
				Code:    http.StatusInternalServerError,
				Message: "Internal server service error",
				Data:    nil,
				Errors:  []string{err.Error()},
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.GenericResponse{
			Success: true,
			Code:    http.StatusOK,
			Message: "discount deleted successfully",
			Data:    nil,
			Errors:  nil,
		})
	return
}
