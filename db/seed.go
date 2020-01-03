package database

import (
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/oshalygin/gqlgen-pg-todo-example/models"
	"github.com/theckman/yacspin"
)

func Seed(db *pg.DB) error {

	if err := createSchemas(db); err != nil {
		return err
	}

	if err := seedUsers(db); err != nil {
		return err
	}

	if err := seedTodos(db); err != nil {
		return err
	}

	return nil
}

func createSchemas(db *pg.DB) error {
	spinner, _ := yacspin.New(yacspin.Config{
		Delay:         100 * time.Millisecond,
		CharSet:       yacspin.CharSets[59],
		Suffix:        " Hydrating Schema",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()
	for _, model := range []interface{}{
		(*models.User)(nil),
		(*models.Todo)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			Temp:          false,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	spinner.Stop()
	return nil
}

func seedUsers(db *pg.DB) error {
	spinner, _ := yacspin.New(yacspin.Config{
		Delay:         100 * time.Millisecond,
		CharSet:       yacspin.CharSets[59],
		Suffix:        " Hydrating Users ",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()

	users := []models.User{
		{
			Email:     "oshalygin@gmail.com",
			FirstName: "Oleg",
			LastName:  "Shalygin",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Email:     "john.snow@gmail.com",
			FirstName: "John",
			LastName:  "Snow",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Email:     "baby.yoda@gmail.com",
			FirstName: "Baby",
			LastName:  "Yoda",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, model := range users {
		var user models.User
		if err := db.Model(&user).Where("email = ?", model.Email).Select(); err != nil {
			if err := db.Insert(&model); err != nil {
				return err
			}
		}
	}
	spinner.Stop()
	return nil
}

func seedTodos(db *pg.DB) error {
	spinner, _ := yacspin.New(yacspin.Config{
		Delay:         100 * time.Millisecond,
		CharSet:       yacspin.CharSets[59],
		Suffix:        " Hydrating Todos ",
		StopMessage:   "Complete",
		Message:       "",
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	})
	spinner.Start()

	todos := []models.Todo{
		{
			Name:       "kubectl all the things",
			IsComplete: false,
			IsDeleted:  false,
			CreatedBy:  1,
			UpdatedBy:  1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Name:       "install a k8s cluster inside of another k8s cluster",
			IsComplete: false,
			IsDeleted:  false,
			CreatedBy:  2,
			UpdatedBy:  2,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Name:       "inception",
			IsComplete: false,
			IsDeleted:  false,
			CreatedBy:  3,
			UpdatedBy:  3,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	for _, model := range todos {
		if err := db.Model(&model).Where("name = ?", model.Name).Select(); err != nil {
			if err := db.Insert(&model); err != nil {
				return err
			}
		}
	}
	spinner.Stop()
	return nil
}
