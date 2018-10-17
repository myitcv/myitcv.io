package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type pkg struct {
	ImportPath   string
	RelPath      string
	GithubPath   string
	WikiUrl      string
	Desc         string
	MonoRepo     bool
	TestMonoRepo bool
}

var ps = []pkg{
	pkg{
		RelPath:    "go-modules-by-example",
		GithubPath: "github.com/go-modules-by-example/index",
	},
	pkg{
		RelPath:    "go",
		GithubPath: "github.com/myitcv/go",
		Desc:       "go is a wrapper around the go tool that automatically sets the GOPATH env variable based on the process' current directory.",
	},
	pkg{
		RelPath:    "vgoimporter",
		GithubPath: "github.com/myitcv/vgoimporter",
		Desc:       "Package vgoimporter is an implementation of go/types.ImporterFrom that uses non-stale package dependency targets where they exist, else falls back to a source-file based importer.",
	},
}

func main() {
	pkgs := make(map[string]pkg, len(ps))

	for _, v := range ps {
		v.ImportPath = path.Join("myitcv.io", v.RelPath)
		pkgs[v.RelPath] = v
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		p := strings.TrimPrefix(req.URL.EscapedPath(), "/")
		ps := strings.Split(p, "/")

		if p == "" {
			// root
			tmpls.ExecuteTemplate(w, "root", pkgs)
		} else if r, ok := pkgs[ps[0]]; ok {
			tmpls.ExecuteTemplate(w, "pkg", r)
		} else {
			// mono repo
			tmpls.ExecuteTemplate(w, "mono", pkg{
				ImportPath: path.Join("myitcv.io", p),
			})
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// copied from Go 1.8.1 std lib
func stripPort(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return hostport
	}
	if i := strings.IndexByte(hostport, ']'); i != -1 {
		return strings.TrimPrefix(hostport[:i], "[")
	}
	return hostport[:colon]
}

var tmpls = template.Must(template.New("tmpls").Parse(`
{{define "mono"}}
<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<meta name="go-import" content="myitcv.io git https://github.com/myitcv/x">
	<!--<meta name="go-import" content="{{.ImportPath}} mod https://raw.githubusercontent.com/myitcv/pubx/master">-->
	<meta name="go-source" content="myitcv.io https://github.com/myitcv/x/wiki https://github.com/myitcv/x/tree/master{/dir} https://github.com/myitcv/x/blob/master{/dir}/{file}#L{line}">
	<meta http-equiv="refresh" content="0; url=https://godoc.org/{{.ImportPath}}">
</head>
<body>
Redirecting to docs at <a href="https://godoc.org/{{.ImportPath}}">godoc.org/{{.ImportPath}}</a>...
</body>
</html>
{{end}}

{{define "pkg"}}
<!DOCTYPE html>
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.ImportPath}}</title>
	 {{if .MonoRepo -}}
    <meta name="go-import" content="myitcv.io git https://github.com/myitcv/x">
    <!--<meta name="go-import" content="{{.ImportPath}} mod https://raw.githubusercontent.com/myitcv/pubx/master">-->
	 {{else if .TestMonoRepo -}}
    <meta name="go-import" content="myitcv.io git https://github.com/myitcv/y">
	 {{else -}}
    <meta name="go-import" content="{{.ImportPath}} git https://{{.GithubPath}}">
	 {{end}}
  </head>
  <body>
  <p>
    <h3><code>{{.ImportPath}}</code></h3>
	 <p>{{.Desc}}</p>
	 <code>go get -u {{.ImportPath}}</code><br/><br/>
	 {{if .WikiUrl}}
    <a href="{{.WikiUrl}}">Wiki</a><br/>
	 {{end}}
    <a href="https://godoc.org/{{.ImportPath}}"><code>godoc</code></a><br/>
	 {{if .MonoRepo -}}
    <a href="https://github.com/myitcv/x/tree/master/{{.RelPath}}">Source</a>
	 {{else if .TestMonoRepo -}}
    <a href="https://github.com/myitcv/y/tree/master/{{.RelPath}}">Source</a>
	 {{else -}}
    <a href="https://{{.GithubPath}}">Source</a>
	 {{end}}
  </p>
	 {{template "footer"}}
  </body>
</html>
{{end}}

{{define "root"}}
<!DOCTYPE html>
<html>
  <head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>myitcv.io</title>
  <meta name="go-import" content="myitcv.io git https://github.com/myitcv/x">
  </head>
  <body>
  <h3><code>myitcv.io</code></h3>
  <p>
  <a href="https://blog.myitcv.io">Blog</a><br/>
  <a href="https://twitter.com/_myitcv">Twitter</a>
  </p>
  <ul style="list-style: none; padding-left:0;">
  <li>Packages: <ul>
  {{range .}}
  <li><a href="/{{.RelPath}}">{{.ImportPath}}</a></li>
  {{end}}
  </ul>
  </li></ul>
	{{template "footer"}}
  </body>
</html>
{{end}}

{{define "footer"}}
<br/>
<hr style="border-width: 1px 1px 0;
           border-style: solid;
           border-color: darkgray;"/>
<span style="font-size:smaller">
<em>
<a href="mailto:paul@myitcv.io">paul@myitcv.io</a><br/>
<a href="/">Home</a>
</em>
</span>
{{end}}
`))
