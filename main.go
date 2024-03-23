package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/f-velka/sqlboiler-test/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	// open connection
	ctx := context.Background()
	db, err := sql.Open("sqlite3", "sqlboiler-test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// set debug mode
	boil.DebugMode = true
	boil.DebugWriter = os.Stdout

	// fetch all tasks with relations
	tasks, err := models.Tasks(
		qm.Load(models.TaskRels.DependsOnTasks),
	).All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tasks {
		fmt.Printf("task: %v,%v\n", t.ID, t.Name)
		for _, d := range t.R.DependsOnTasks {
			fmt.Printf("\tdepends on: %v,%v\n", d.ID, d.Name)
		}
	}

	// start transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	// add new record
	newTask := &models.Task{
		ID:   5,
		Name: "newTask",
	}
	err = newTask.Insert(ctx, tx, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	// remove a record
	models.Tasks(
		qm.Where("id=?", newTask.ID),
	).DeleteAll(ctx, tx)

	tx.Commit()
}
