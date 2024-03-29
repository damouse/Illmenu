package illmenu

// This is the restful portion of the server as well as the model section

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func StartServer() {

	i := Impl{}
	i.InitDB()
	i.InitSchema()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/reminders", i.GetAllReminders),
		rest.Post("/reminders", i.PostReminder),
		rest.Get("/reminders/:id", i.GetReminder),
		rest.Put("/reminders/:id", i.PutReminder),
		rest.Delete("/reminders/:id", i.DeleteReminder),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Reminder struct {
	Id        int64     `json:"id"`
	Message   string    `sql:"size:1024" json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Impl struct {
	DB *gorm.DB
}

func (i *Impl) InitDB() {
	var err error
	i.DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=mydb sslmode=disable password=postgres")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	i.DB.LogMode(true)
}

func (i *Impl) InitSchema() {
	i.DB.AutoMigrate(&Reminder{})
}

func (i *Impl) GetAllReminders(w rest.ResponseWriter, r *rest.Request) {
	reminders := []Reminder{}
	i.DB.Find(&reminders)
	w.WriteJson(&reminders)
}

func (i *Impl) GetReminder(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	reminder := Reminder{}
	if i.DB.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reminder)
}

func (i *Impl) PostReminder(w rest.ResponseWriter, r *rest.Request) {
	reminder := Reminder{}

	if err := r.DecodeJsonPayload(&reminder); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := i.DB.Save(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var count int64

	if err := i.DB.Model(&Reminder{}).Count(&count).Error; err != nil {
		fmt.Errorf(fmt.Sprintf("Count should work, but got err %v", err))
	} else {
		fmt.Println("Count", count)
	}

	w.WriteJson(&reminder)
}

func (i *Impl) PutReminder(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")
	reminder := Reminder{}
	if i.DB.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Reminder{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reminder.Message = updated.Message

	if err := i.DB.Save(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&reminder)
}

func (i *Impl) DeleteReminder(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	reminder := Reminder{}
	if i.DB.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := i.DB.Delete(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
