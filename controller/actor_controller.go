package controller

import (
	"Todo_APP/model"
	"database/sql"
	"fmt"
	"strings"
)

// CreateActor inserts a new actor into the database with SQL injection protection.
func CreateActor(db *sql.DB, firstName, lastName string) error {
	firstName = strings.ToUpper(firstName)
	lastName = strings.ToUpper(lastName)
	query := "INSERT INTO actor (first_name, last_name) VALUES (?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(firstName, lastName)
	if err != nil {
		return fmt.Errorf("error inserting actor: %w", err)
	}
	return nil
}

// GetActors retrieves all actors from the database with SQL injection protection.
func GetActors(db *sql.DB) ([]model.Actor, error) {
	query := "SELECT actor_id, first_name,  last_name FROM actor LIMIT 10"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []model.Actor
	for rows.Next() {
		var actor model.Actor
		if err := rows.Scan(&actor.ID, &actor.First_Name, &actor.Last_Name); err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}
	return actors, nil
}

// UpdateActor updates an actor's name based on the ID with SQL injection protection.
func UpdateActor(db *sql.DB, id int, firstName, lastName string) error {
	firstName = strings.ToUpper(firstName)
	lastName = strings.ToUpper(lastName)
	query := "UPDATE actor SET first_name = ?, last_name = ? WHERE actor_id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(firstName, lastName, id)
	if err != nil {
		return fmt.Errorf("error updating actor: %w", err)
	}
	fmt.Println("Actor updated successfully!")
	return nil
}

// DeleteActor removes an actor from the database with SQL injection protection.
func DeleteActor(db *sql.DB, id int) error {
	query := "DELETE FROM actor WHERE actor_id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error deleting actor: %w", err)
	}
	fmt.Println("Actor deleted successfully!")
	return nil
}
