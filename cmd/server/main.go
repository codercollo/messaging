package main

import (
	"context"
	"fmt"

	"github.com/codercollo/messaging/internal/comment"
	"github.com/codercollo/messaging/internal/db"
	transportHttp "github.com/codercollo/messaging/internal/transport/http"
)

// Run - is going to be responsibe for the instantiation and startup of the application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the databse")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil

}

func main() {
	fmt.Println("Go Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
