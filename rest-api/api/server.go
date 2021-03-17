package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/routers"
)

func listen(p int) {
	fmt.Printf("\n\nListening port %d\n\n<-", p)
	port := fmt.Sprintf(":%d", p)
	r := routers.MyHandlers()
	log.Fatal(http.ListenAndServe(port, r))
}
