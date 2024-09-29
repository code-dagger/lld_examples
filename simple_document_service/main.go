package main

import (
	"fmt"

	"github.com/code-dagger/simple_doc_service/docservice"
)

func main() {
	srv := docservice.NewService()
	srv.CreateDocument("user1", "Doc1", "This is the content of Doc1")
	srv.CreateDocument("user2", "Doc2", "This is the content of Doc2")
	srv.GrantAccess("user1", "Doc1", "user2", docservice.AccessRead)
	// Reading document
	content, err := srv.ReadDocument("user2", "Doc1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Read Doc1 as user2:", content)
	}

	// Trying to edit without permission
	err = srv.EditDocument("user2", "Doc1", "Updated content")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Editing with owner permission
	err = srv.EditDocument("user1", "Doc1", "Updated content by user1")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Deleting document
	err = srv.DeleteDocument("user1", "Doc1")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
