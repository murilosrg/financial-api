package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/murilosrg/financial-api/internal/model/accounts"
)

type IAccountController interface {
	Post(c *gin.Context)
	Find(c *gin.Context)
}

type AccountController struct {
	service accounts.IAccountService
}

func NewAccountController(service accounts.IAccountService) IAccountController {
	return &AccountController{
		service: service,
	}
}

func (a *AccountController) Post(c *gin.Context) {
	account := &accounts.Account{}
	c.BindJSON(&account)

	account, err := a.service.Create(account)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (a *AccountController) Find(c *gin.Context) {
	id := c.Params.ByName("accountId")
	account := &accounts.Account{}

	idParam, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	account, err = a.service.Find(idParam)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}
