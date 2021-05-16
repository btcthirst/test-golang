package api

import "github.com/btcthirst/practical-tasks-nix/rest-echo-swagger-oauth2.0/api/models"

func RunApi() {
	// migration
	models.Automigrator()
	// run server
	StartServ()
}
