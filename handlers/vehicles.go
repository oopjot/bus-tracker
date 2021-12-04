import (
	"net/http"
	"json"
	"net/url"
	"log"
)

func VehiclesHandler(w http.ResponseWriter, req *http.Request) {
	lines = req.URL.Query().Get("lines")

	fmt.Println(lines)

}
