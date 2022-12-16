package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
	"webapp/webui"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	utils.CarregarTemplates()

	r := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	webui.RegisterUIHandlers(r, n)

	r = router.Gerar()
	fmt.Println("Escutando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
