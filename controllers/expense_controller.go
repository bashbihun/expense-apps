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

func GetExpenseById(ctx *gin.Context) {
	id := ctx.Param("id")
	var expense models.Expense

	if err := config.DB.First(&expense, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "pengeluaran tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": expense})
}

func UpdateExpense(ctx *gin.Context) {
	id := ctx.Param("id")
	var expense models.Expense

	// Cek apakah data ada
	if err := config.DB.First(&expense, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Pengeluaran tidak ditemukan"})
		return
	}

	var updatedExpense models.Expense
	if err := ctx.ShouldBindJSON(&updatedExpense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Format input tidak valid"})
		return
	}

	// Validasi input
	if err := validate.Struct(updatedExpense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update data
	expense.Title = updatedExpense.Title
	expense.Amount = updatedExpense.Amount
	expense.Description = updatedExpense.Description

	if err := config.DB.Save(&expense).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate", "data": expense})
}
