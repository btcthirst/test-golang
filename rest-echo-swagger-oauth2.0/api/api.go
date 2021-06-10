package api

<<<<<<< Updated upstream
import "github.com/btcthirst/practical-tasks-nix/rest-echo-swagger-oauth2.0/api/models"

func RunApi() {
=======
import (
	"github.com/btcthirst/practical-tasks-nix/rest-echo-swagger-oauth2.0/api/models"
)

func RunApi() {

>>>>>>> Stashed changes
	// migration
	models.Automigrator()
	// run server
	StartServ()
}
