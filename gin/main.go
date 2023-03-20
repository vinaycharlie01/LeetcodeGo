package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Emp struct {
	EMPID  string `json:"empid"`
	Name   string `json:"name"`
	SALARY int    `json:"salary"`
}

const (
	address = "localhost:8080"
)

func NewRoters(r *gin.Engine) {
	b := Emp{
		Name:   "vinay",
		EMPID:  "1",
		SALARY: 12000,
	}
	a = append(a, b)
	api := r.Group("/get")
	api.GET("/", GetEmp)
	api.POST("/", ADDemp)
	// api.DELETE("", DeleteEMP)
	// api.PUT("", UdateEMP)
}

var a []Emp

func GetEmp(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, &a)
}

func GetEmp2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode()
	// ctx.Header("Content-Type", "application/json")
	// ctx.JSON(http.StatusOK, &a)
}

func ADDemp(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var body Emp
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	a = append(a, body)
	ctx.JSON(http.StatusOK, &a)
}

// func DeleteEMP(wr *gin.Context) {

// }

// func UdateEMP(wr *gin.Context) {

// }

func main() {
	router := gin.Default()
	NewRoters(router)
	// router.GET("/")
	router.Run(address)
}
