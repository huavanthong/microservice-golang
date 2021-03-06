/*
 * @File: controllers.profile.go
 * @Description: Implements User API logic functions
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/daos"
	"github.com/huavanthong/microservice-golang/user-api/models"
	"github.com/huavanthong/microservice-golang/user-api/payload"
	"github.com/huavanthong/microservice-golang/user-api/utils"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

// Define profile manages
type Profile struct {
	utils      utils.Utils
	profileDAO daos.Profile
}

// ListProfiles godoc
// @Summary List all existing user profiles
// @Description List all existing user profiles
// @Tags profile
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 500 {object} payload.Error
// @Success 200 {array} models.Profile
// @Router /profile/list [get]
// ListUsers get all users exist in DB
func (p *Profile) ListProfiles(ctx *gin.Context) {

	// array of profiles
	var profiles []models.Profile

	// get all user profiles
	profiles, err := p.profileDAO.GetAll()

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, profiles)
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}

}

// GetProfileByUserId godoc
// @Summary Get a profile user by userid
// @Description Profile user
// @Tags profile
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param userid path string true "User ID"
// @Failure 500 {object} payload.Error
// @Success 200 {object} models.Profile
// @Router /profile/{userid} [get]
// GetProfileByUserId get a profile by user id in DB
func (u *Profile) GetProfileByUserId(ctx *gin.Context) {

	// filter parameter id from context
	userId := ctx.Params.ByName("userid")

	// find user by id
	user, err := u.profileDAO.GetProfileByUserId(userId)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}
}

// UpdateProfileByUserId godoc
// @Summary Update an existing profile user
// @Description Update an existing profile user
// @Tags profile
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param profile body models.Profile true "User ID"
// @Failure 500 {object} payload.Error
// @Success 200 {object} payload.Message
// @Router /profile [patch]
func (p *Profile) UpdateProfileByUserId(ctx *gin.Context) {

	// bind user data to json
	var profile models.Profile
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, payload.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	// update user by user
	err := p.profileDAO.Update(profile)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, payload.Message{"Successfully"})
		fmt.Errorf("Update profile name = " + profile.ProfileName + ", firtname = " + profile.FirstName)
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}
}

// AddProfile godoc
// @Summary Add a profile for user
// @Description Update profile user
// @Tags profile
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body models.Profile true "Add profile user"
// @Failure 500 {object} payload.Error
// @Failure 400 {object} payload.Error
// @Success 200 {object} payload.Message
// @Router /profile [post]
func (p *Profile) AddProfile(ctx *gin.Context) {
	// bind profile info to json getting context
	var mp models.Profile
	if err := ctx.ShouldBindJSON(&mp); err != nil {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	// validate data on profile of user
	v := utils.NewValidation()

	verr := v.Validate(mp)
	if verr != nil {
		err1 := errors.New("Todo: Fix this issue")
		ctx.JSON(http.StatusBadRequest, payload.Error{common.StatusCodeUnknown, err1.Error()})
		return
	}

	address := &models.Address{
		Street:   "Kinh Duong Vuong",
		Ward:     "12",
		District: "6",
		City:     "Ho Chi Minh",
		Country:  "Viet Nam",
	}

	// create profile from models
	profile := models.Profile{
		ID:          bson.NewObjectId(),
		FirstName:   mp.FirstName,
		LastName:    mp.LastName,
		UserID:      bson.NewObjectId(), // define a new object bson
		PhoneNumber: mp.PhoneNumber,
		Addresses:   []*models.Address{address},
	}

	// insert user to DB
	err := p.profileDAO.Insert(profile)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, payload.Message{"Successfully"})
		log.Debug("Create profile name = " + profile.ProfileName + ", firstname = " + profile.FirstName)
	} else {
		err1 := errors.New("Todo: Fix this issue")
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err1.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// DeteleProfileByUserId godoc
// @Summary Delete a user profile by UserID
// @Description Delete a user profile by UserID
// @Tags profile
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param userid path string true "User ID"
// @Failure 500 {object} payload.Error
// @Success 200 {object} payload.Message
// @Router /profile/{userid} [delete]
func (p *Profile) DeteleProfileByUserId(ctx *gin.Context) {
	// filter parameter id context
	id := ctx.Params.ByName("userid")

	// delete user by id
	err := p.profileDAO.DeleteByUserID(id)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, payload.Message{"Successfully"})
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}
}
