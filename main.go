package main

import (
	"log"
	"net/http"
)

const JSFILE = `
	<html>
	<body>
		<script>
			const evtSource = new EventSource("sse");
			evtSource.addEventListener("ping", (e) => {
				console.log(e);
			});
		</script>
	</body>
	</html> `

type Game struct {
	ongoing chan
}

func main() {
	log.Print("listening on :8080")
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Context().Done()
		w.Write([]byte(JSFILE))
	})

	router.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		
	})

	http.ListenAndServe(":8080", router)
}
