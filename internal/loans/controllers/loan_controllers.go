package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcelodiasdev/library-mvc/internal/loans/models"
)

type LoanController struct {
	loanService models.LoanService
}

func NewUserController(loanService models.LoanService) *LoanController {
	return &LoanController{loanService: loanService}
}

func (l LoanController) RegisterRoutes(r *gin.Engine) {
	loans := r.Group("/loans")

	{
		loans.POST("", l.CreateLoan)
		loans.GET("/:id", l.GetLoan)
		loans.GET("", l.GetAllLoans)
		loans.PUT("/:id", l.GetUserLoans)
		loans.DELETE("/:id", l.ReturnBook)
	}

	users := r.Group("/loans/users")

	{
		users.GET("/:userId/loans", l.GetUserLoans)

	}
}

func (l LoanController) CreateLoan(ctx *gin.Context) {
	var request struct {
		BookId int64 `json:"book_id"`
		UserId int64 `json:"user_id"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body!"})
		return
	}

	loan, err := l.loanService.CreateLoan(request.BookId, request.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, loan)

}

func (l LoanController) GetLoan(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	loan, err := l.loanService.GetLoan(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (l LoanController) GetAllLoans(ctx *gin.Context) {
	loan, err := l.loanService.GetAllLoans()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loan)
}

func (l LoanController) GetUserLoans(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	loans, err := l.loanService.GetUserLoans(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loans)

}

func (l LoanController) ReturnBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	err = l.loanService.ReturnBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
