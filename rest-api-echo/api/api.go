package api

import "github.com/btcthirst/practical-tasks-nix/rest-api-echo/api/models"

func RunServ() {
	models.AutoMigrations()
	listenServ()
}
