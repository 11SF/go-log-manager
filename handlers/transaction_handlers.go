package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Member_id    string  `json:"member_id" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	Family_id    string  `json:"family_id" validate:"required"`
	Family_name  string  `json:"family_name" validate:"required"`
	Price        float32 `json:"price" validate:"required"`
	Month        int32   `json:"month" validate:"required"`
	Date_overdue int32   `json:"date_overdue" validate:"required"`
	Status       string  `json:"status" validate:"required"`
}

type Message struct {
	Message string `json:"message"`
}

type TransactionHandler struct {
	DB *gorm.DB
}

func messageResponce(msg string) (res Message) {
	res.Message = msg
	return res
}

func (h *TransactionHandler) Initialize() {
	db, err := gorm.Open(sqlite.Open("log.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Transaction{})
	h.DB = db
}

func (h *TransactionHandler) GetAllTransaction(c echo.Context) error {
	transaction := []Transaction{}
	h.DB.Find(&transaction)
	return c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) GetAllTransactionByFamilyId(c echo.Context) error {
	familyId := c.Param("familyId")

	transaction := []Transaction{}
	h.DB.Where("family_id = ?", familyId).Find(&transaction)

	if len(transaction) == 0 {
		return c.JSON(http.StatusOK, messageResponce("Not Found"))
	}
	return c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) SaveTransaction(c echo.Context) (err error) {
	u := new(Transaction)
	if err = c.Bind(u); err != nil {
		return
	}
	data := Transaction{
		Member_id:    u.Member_id,
		Name:         u.Name,
		Family_id:    u.Family_id,
		Family_name:  u.Family_name,
		Price:        u.Price,
		Month:        u.Month,
		Date_overdue: u.Date_overdue,
	}
	if data.Member_id != "" && data.Name != "" && data.Family_id != "" && data.Family_name != "" {
		h.DB.Create(&data)
		return c.NoContent(http.StatusCreated)
	}
	return c.NoContent(http.StatusBadRequest)
}

func (h *TransactionHandler) DeleteTransaction(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	h.DB.Where("id = ?", id).Delete(&Transaction{})
	return c.NoContent(http.StatusOK)
}

func (h *TransactionHandler) UpdateTransaction(c echo.Context) (err error) {
	familyId := c.Param("familyId")

	u := new(Transaction)
	if err = c.Bind(u); err != nil {
		return
	}
	data := Transaction{
		Member_id:    u.Member_id,
		Name:         u.Name,
		Family_id:    u.Family_id,
		Family_name:  u.Family_name,
		Price:        u.Price,
		Month:        u.Month,
		Date_overdue: u.Date_overdue,
	}

	h.DB.Where("family_id = ?", familyId).Updates(Transaction{Price: data.Price, Month: data.Month, Date_overdue: data.Date_overdue})

	// h.DB.Where("id = ?", id).Delete(&Transaction{})
	return c.NoContent(http.StatusOK)
}
