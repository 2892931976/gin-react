package controllers

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/tinchi/gin-react/db"
  "github.com/tinchi/gin-react/models"
  "net/http"
  "github.com/tinchi/gin-react/forms"
)

type DepositController struct{}

func getCurrentUser(c *gin.Context) models.User {
  current_user, _ := c.Get("current_user")

  fmt.Println("current_user_id", current_user.(models.User))

  return current_user.(models.User)
}

func (ctrl DepositController) IndexEndpoint(c *gin.Context) {
  var deposits []models.Deposit

  current_user := getCurrentUser(c)

  err := db.Engine.Where("user_id = ?", current_user.Id).Find(&deposits)

  if err != nil {
    fmt.Println(err)
  }

  if len(deposits) > 0 {
    c.JSON(http.StatusOK, gin.H{"deposits": deposits, "count": len(deposits)})
  } else {
    c.JSON(http.StatusOK, gin.H{"deposits": []models.Deposit{}, "count": len(deposits)})
  }
}

func (ctrl DepositController) CreateEndpoint(c *gin.Context) {
  var form forms.DepositForm

  current_user := getCurrentUser(c)

  err := c.BindJSON(&form)

  if err == nil {
    deposit := models.Deposit{
      BankName:      form.BankName,
      AccountNumber: form.AccountNumber,
      Ammount:       form.Ammount,
      StartDate:     form.StartDate,
      EndDate:       form.EndDate,
      Interest:      form.Interest,
      Taxes:         form.Taxes,
      UserId:        current_user.Id,
    }

    _, err = db.Engine.Insert(&deposit)

    if err != nil {
      panic(err)
    }

    c.JSON(http.StatusCreated, gin.H{"deposit": deposit})
  } else {
    fmt.Println(err)

    c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
  }
}
func (ctrl DepositController) ShowEndpoint(c *gin.Context) {
  var deposit models.Deposit

  id := c.Param("id")

  _, err := db.Engine.Where("deposits.id = ?", id).
    Get(&deposit)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, gin.H{"deposit": deposit})
}

func (ctrl DepositController) UpdateEndpoint(c *gin.Context) {
  var form forms.DepositForm

  // current_user := getCurrentUser(c)
  id := c.Param("id")
  err := c.BindJSON(&form)

  if err == nil {
    deposit := models.Deposit{
      BankName:      form.BankName,
      AccountNumber: form.AccountNumber,
      Ammount:       form.Ammount,
      StartDate:     form.StartDate,
      EndDate:       form.EndDate,
      Interest:      form.Interest,
      Taxes:         form.Taxes,
    }
    _, err = db.Engine.Where("deposits.id = ?", id).
      Update(&deposit)

    if err != nil {
      panic(err)
    }

    c.JSON(http.StatusOK, gin.H{"deposit": deposit})
  } else {
    fmt.Println(err)

    c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
  }
}

func (ctrl DepositController) DeleteEndpoint(c *gin.Context) {
  var deposit models.Deposit

  id := c.Param("id")

  _, err := db.Engine.Where("deposits.id = ?", id).
    Delete(&deposit)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, gin.H{})
}