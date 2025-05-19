package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

var zones = []string{
	"UTC",
	"America/New_York",
	"Europe/London",
	"Asia/Tokyo",
	"Australia/Sydney",
	"Europe/Berlin",
	"America/Los_Angeles",
	"Asia/Kolkata",
	"Africa/Cairo",
	"America/Sao_Paulo",
	"Pacific/Auckland",
	"Europe/Moscow",
}

var zoneImages = map[string]string{
	"UTC":                 "https://source.unsplash.com/featured/?universe",
	"America/New_York":    "https://source.unsplash.com/featured/?newyork",
	"Europe/London":       "https://source.unsplash.com/featured/?london",
	"Asia/Tokyo":          "https://source.unsplash.com/featured/?tokyo",
	"Australia/Sydney":    "https://source.unsplash.com/featured/?sydney",
	"Europe/Berlin":       "https://source.unsplash.com/featured/?berlin",
	"America/Los_Angeles": "https://source.unsplash.com/featured/?losangeles",
	"Asia/Kolkata":        "https://source.unsplash.com/featured/?kolkata",
	"Africa/Cairo":        "https://source.unsplash.com/featured/?cairo",
	"America/Sao_Paulo":   "https://source.unsplash.com/featured/?saopaulo",
	"Pacific/Auckland":    "https://source.unsplash.com/featured/?auckland",
	"Europe/Moscow":       "https://source.unsplash.com/featured/?moscow",
}

type Clock struct {
	Zone     string
	Time     string
	ImageURL string
}

type PageData struct {
	Clocks []Clock
	Arch   string
	OS     string
}

var tmpl = template.Must(template.New("page").Parse(`<!DOCTYPE html>
<html>
<head>
    <title>World Clocks</title>
    <style>
        .city {
            display: inline-block;
            margin: 10px;
            text-align: center;
        }
        .city img {
            width: 150px;
            height: 100px;
            object-fit: cover;
            display: block;
        }
    </style>
</head>
<body>
    <h1>World Clocks</h1>
    {{range .Clocks}}
    <div class="city">
        <img src="{{.ImageURL}}" alt="{{.Zone}}">
        <p>{{.Zone}}: {{.Time}}</p>
    </div>
    {{end}}
    <footer>
        <p>ARCH: {{.Arch}}</p>
        <p>OS: {{.OS}}</p>
    </footer>
</body>
</html>`))

func greet(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	order := rand.Perm(len(zones))
	clocks := make([]Clock, 0, 10)
	for i := 0; i < 10 && i < len(order); i++ {
		zone := zones[order[i]]
		loc, err := time.LoadLocation(zone)
		if err != nil {
			continue
		}
		clocks = append(clocks, Clock{
			Zone:     zone,
			Time:     time.Now().In(loc).Format(time.RFC1123),
			ImageURL: zoneImages[zone],
		})
	}
	data := PageData{
		Clocks: clocks,
		Arch:   runtime.GOARCH,
		OS:     runtime.GOOS,
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":80", nil)
}
