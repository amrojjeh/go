package main

import (
	"log"
	"net/http"
	"time"
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

// type Game struct {
// 	ongoing chan
// }

func main() {
	// buffer:=make(map int[Game])
	log.Print("listening on :8080")
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Context().Done()
		w.Write([]byte(JSFILE))
	})

	router.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Accel-Buffering", "no")
		w.Header().Add("Content-Type", "text/event-stream")
		w.Header().Add("Cache-Control", "no-cache")
		for range 5 {
			w.Write([]byte(`event: ping
data: {hi:"ji"}

`))
			w.(http.Flusher).Flush()
			time.Sleep(time.Second * 5)
		}
	})

	http.ListenAndServe(":8080", router)
}
