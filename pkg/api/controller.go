package api

import "github.com/IaC/go-kcloutie/pkg/database"

type Controller struct {
	DBInterface database.DBInterface
}

func NewController(db database.DBInterface) Controller {
	return Controller{db}
}
