package controllers

import (
	"expense-tracker/config"
	"expense-tracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateExpense(ctx *gin.Context) {
	var expense models.Expense

	if err := ctx.ShouldBindJSON(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "format input tidak valid"})
		return
	}

	err := validate.Struct(expense)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&expense).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Pengeluaran berhasil ditambahkan", "data": expense})
}

func GetAllExpense(ctx *gin.Context) {
	var expense []models.Expense

	if err := config.DB.Find(&expense).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": expense})
}
