package controller

import (
	"caixa-eletronico/dto"
	"caixa-eletronico/errors"
	"caixa-eletronico/model"
	"caixa-eletronico/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexUsers(c *gin.Context) {
	var user model.Usuario

	users, err := user.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: "Falied to find user"})
		return
	}

	c.JSON(http.StatusOK, users)
	return
}

func ShowUser(c *gin.Context) {
	var user model.Usuario
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: "Invalid param"})
		return
	}

	if err := user.FindByID(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: "Falied to find user"})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func CreateUser(c *gin.Context) {
	var user model.Usuario
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: "Invalid payload"})
		return
	}

	if err := util.Validate(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: err.Error()})
		return
	}

	if err := user.Create(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: "Falied to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func UpdateUser(c *gin.Context) {
	var request model.Usuario
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: "Invalid payload"})
		return
	}

	if err := util.Validate(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: "Invalid param"})
		return
	}

	var user model.Usuario

	if err := user.FindByID(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: "Falied to find user"})
		return
	}

	if err := user.Update(request); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: "Falied to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: "Invalid param"})
		return
	}

	var user model.Usuario

	if err := user.Delete(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: "Falied to delete user"})
		return
	}

	c.JSON(http.StatusOK, dto.MessageReply{Message: "Successfully deleted"})
	return
}
