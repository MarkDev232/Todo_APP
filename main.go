package main
import (
	"Todo_APP/controller"
	"Todo_APP/database"
	"fmt"
	"log"
	"os"
)
func askifcontinue(){
	fmt.Println("Do you want to continue? (y/n)")
	var choice string
	fmt.Scan(&choice)
	if choice == "n" {
		os.Exit(0)
	} else if choice == "y" {
		main()
	} else {
		fmt.Println("Invalid choice!")
		askifcontinue()
	}

}
func main() {
	// Establish database connection
	db := database.Connect()
	defer db.Close()

	fmt.Println("Choose an operation:")
	fmt.Println("1. Create Actor")
	fmt.Println("2. Read Actors")
	fmt.Println("3. Update Actor")
	fmt.Println("4. Delete Actor")
	fmt.Println("5. Exit")

	var choice int
	fmt.Print("Enter choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		log.Fatal("Invalid input:", err)
		askifcontinue()
	}

	switch choice {
	case 1:
		var firstName, lastName string
		fmt.Print("Enter first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter last name: ")
		fmt.Scan(&lastName)
		controller.CreateActor(db, firstName, lastName)


	case 2:
		actors, err := controller.GetActors(db)
		if err != nil {
			log.Fatal("Error fetching actors:", err)
		}
		fmt.Println("ID\tName")
		fmt.Println("-----------------")
		for _, actor := range actors {
			fmt.Printf("%d\t%s\n", actor.ID, actor.Name)
		}

	case 3:
		var id int
		var firstName, lastName string
		fmt.Print("Enter actor ID to update: ")
		fmt.Scan(&id)
		fmt.Print("Enter new first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter new last name: ")
		fmt.Scan(&lastName)
		controller.UpdateActor(db, id, firstName, lastName)

	case 4:
		var id int
		fmt.Print("Enter actor ID to delete: ")
		fmt.Scan(&id)
		controller.DeleteActor(db, id)
	case 5:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice!")
	}
	askifcontinue()
}
