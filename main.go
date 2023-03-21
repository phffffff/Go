package main

import (
	"RestAPI/component/appContext"
	"RestAPI/middleware"
	"RestAPI/module/restaurant/transport/ginRestaurant"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Note struct {
	Id     int    `json:"id" gorm:"column:id;"`
	Title  string `json:"title" gorm:"column:title;"`
	Status int    `json:"status" gorm:"column:status;"`
}

type NoteUpdate struct {
	Title  *string `json:"title" gorm:"title"`
	Status *int    `json:"status" gorm:"status"`
}

type Page struct {
	//Page it
}

func main() {
	dns := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("HOST"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.Debug()

	appCtx := appContext.NewAppCtx(db)

	router := gin.Default()
	router.Use(middleware.Recover(appCtx))

	v1 := router.Group("/RestAPI")

	v1.GET("/listRestaurants", ginRestaurant.ListRestaurant(appCtx))
	v1.GET("/getNotes", func(c *gin.Context) {
		var noteArr []Note

		if err := db.Find(&noteArr).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": noteArr,
			})
		}
	})

	v1.POST("/createRestaurant", ginRestaurant.CreateRestaurant(appCtx))

	v1.POST("/createNote", func(c *gin.Context) {
		var data Note

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		}
	})

	v1.GET("/getNoteById/:id", func(c *gin.Context) {
		id := c.Param("id")
		var note Note

		if err := db.Where("id = ?", id).First(&note).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": note,
			})
		}
	})

	v1.PUT("/updateNoteById/:id", func(c *gin.Context) {
		id := c.Param("id")

		var data NoteUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Table("notes").Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	v1.PATCH("/updateElementInNote/:id", func(c *gin.Context) {
		id := c.Param("id")

		var data NoteUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Table("notes").Where("id = ?", id).Updates(&data).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	v1.PATCH("/deleteRestaurant/:id", ginRestaurant.DeleteRestaurant(appCtx))

	v1.DELETE("/deleteNoteById/:id", func(c *gin.Context) {
		id := c.Param("id")

		if err := db.Table("notes").Where("id = ?", id).Delete(nil).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
		})
	})
	router.Run()

}
