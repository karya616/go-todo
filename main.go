package main

import (
	"final_projek_go/controllers"
	"final_projek_go/models"
	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		panic(err.Error())
	}

	noteController := &controllers.NoteController{}

	router := httprouter.New()

	router.GET("/", noteController.Index)
	router.GET("/create", noteController.Create)
	router.POST("/create", noteController.Store)
	router.GET("/edit/:id", noteController.Edit)
	router.POST("/edit/:id", noteController.Update)
	router.POST("/done/:id", noteController.Done)
	router.POST("/delete/:id", noteController.Delete)

	//port := ":8000"

	//fmt.Println("aman bos")
	//log.Fatal(http.ListenAndServe(port, router))
	http.ListenAndServe(":8000", router)

}
