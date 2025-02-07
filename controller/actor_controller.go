package controller

import (
	"Todo_APP/model"
	"database/sql"
	"fmt"
	"log"
)

// CreateActor inserts a new actor into the database.
func CreateActor(db *sql.DB, firstName, lastName string) {
	query := "INSERT INTO actor (first_name, last_name) VALUES (?, ?)"
	_, err := db.Exec(query, firstName, lastName)
	if err != nil {
		log.Fatal("Error inserting actor:", err)
	}
	fmt.Println("Actor added successfully!")
}

// GetActors retrieves all actors from the database.
func GetActors(db *sql.DB) ([]model.Actor, error) {
	rows, err := db.Query("SELECT actor_id, CONCAT(first_name, ' ', last_name) FROM actor")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []model.Actor
	for rows.Next() {
		var actor model.Actor
		if err := rows.Scan(&actor.ID, &actor.Name); err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}

	return actors, nil
}

// UpdateActor updates an actor's name based on the ID.
func UpdateActor(db *sql.DB, id int, firstName, lastName string) {
	query := "UPDATE actor SET first_name = ?, last_name = ? WHERE actor_id = ?"
	_, err := db.Exec(query, firstName, lastName, id)
	if err != nil {
		log.Fatal("Error updating actor:", err)
	}
	fmt.Println("Actor updated successfully!")
}

// DeleteActor removes an actor from the database.
func DeleteActor(db *sql.DB, id int) {
	query := "DELETE FROM actor WHERE actor_id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("Error deleting actor:", err)
	}
	fmt.Println("Actor deleted successfully!")
}
