package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
	w.WriteHeader(http.StatusForbidden)
        fmt.Fprint(w, "Sorry! Only <b>POST</b> requests are allowed!")
        return
    }
    r.ParseForm()
    continueUrl := r.Form.Get("continue")
    fmt.Fprint(w,
`
<head>
<title>Redirect on Post</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1, user-scalable=0">
</head>
<body>
This is indeed a <b>POST</b> request, I'm happy to accept it! :)<br><br>
I will take you back to where you came from in <span id="secs">3</span> seconds.
<script>
  let secsLeft = document.getElementById("secs");
  setTimeout(() => {secsLeft.textContent = "2"}, 1000);
  setTimeout(() => {secsLeft.textContent = "1"}, 2000);
  setTimeout(() => {window.top.location = "` + continueUrl + `#cameBack"}, 3000);
</script>
</body>`)
}
