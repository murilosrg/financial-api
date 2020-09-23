package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/murilosrg/financial-api/internal/model"
)

func PostAccount(c *gin.Context) {
	account := model.Account{}
	c.BindJSON(&account)

	if err := account.Create(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, account)
}

func FindAccount(c *gin.Context) {
	id := c.Params.ByName("accountId")
	account := model.Account{}

	idParam, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	if err := account.Find(idParam); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, account)
}
