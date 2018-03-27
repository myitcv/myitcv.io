package myitcv_io

import (
	"html/template"
	"net/http"
	"path"
	"strings"
)

type pkg struct {
	ImportPath string
	RelPath    string
	GithubPath string
	WikiUrl    string
	Desc       string
}

var ps = []pkg{
	pkg{
		RelPath:    "gopherize.me",
		GithubPath: "github.com/myitcv/gopherize.me",
		Desc:       "GopherJS React version of Gopherize.me",
	},
	pkg{
		RelPath:    "react",
		GithubPath: "github.com/myitcv/react",
		WikiUrl:    "https://github.com/myitcv/react/wiki",
		Desc:       "Package react is a set of GopherJS bindings for Facebook's React, a Javascript library for building user interfaces.",
	},
	pkg{
		RelPath:    "remarkable",
		GithubPath: "github.com/myitcv/remarkable",
		Desc:       "Package remarkable provides an incomplete wrapper for remarkable (https://github.com/jonschlinkert/remarkable), a pure Javascript markdown parser",
	},
	pkg{
		RelPath:    "highlightjs",
		GithubPath: "github.com/myitcv/highlightjs",
		Desc:       "Package highlightjs provides an incomplete wrapper for Highlight.js (https://github.com/isagalaev/highlight.js), a Javascript syntax highlighter",
	},
	pkg{
		RelPath:    "immutable",
		GithubPath: "github.com/myitcv/immutable",
		Desc:       "Package immutable is a helper package for the immutable data structures generated by myitcv.io/immutable/cmd/immutableGen.",
	},
	pkg{
		RelPath:    "sorter",
		GithubPath: "github.com/myitcv/sorter",
	},
	pkg{
		RelPath:    "gogenerate",
		GithubPath: "github.com/myitcv/gogenerate",
		Desc:       "Package gogenerate exposes some of the unexported internals of the go generate command as a convenience for the authors of go generate generators.  See https://github.com/myitcv/gogenerate/wiki/Go-Generate-Notes for further notes on such generators. It also exposes some convenience functions that might be useful to authors of generators",
	},
	pkg{
		RelPath:    "go",
		GithubPath: "github.com/myitcv/go",
		Desc:       "go is a wrapper around the go tool that automatically sets the GOPATH env variable based on the process' current directory.",
	},
	pkg{
		RelPath:    "gg",
		GithubPath: "github.com/myitcv/gg",
	},
	pkg{
		RelPath:    "gjbt",
		GithubPath: "github.com/myitcv/gjbt",
		Desc:       "gjbt is a simple (temporary) wrapper for GopherJS to run tests in Chrome as opposed to NodeJS. It should be considered to be a direct replacement for gopherjs test",
	},
	pkg{
		RelPath:    "gai",
		GithubPath: "github.com/myitcv/gai",
	},
	pkg{
		RelPath:    "gitgodoc",
		GithubPath: "github.com/myitcv/gitgodoc",
		Desc:       "gitgodoc allows you to view godoc documentation for different branches of a Git repository",
	},
	pkg{
		RelPath:    "g",
		GithubPath: "github.com/myitcv/g",
	},
}

func init() {
	pkgs := make(map[string]pkg, len(ps))

	for _, v := range ps {
		v.ImportPath = path.Join("myitcv.io", v.RelPath)
		pkgs[v.RelPath] = v
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		p := strings.TrimPrefix(req.URL.EscapedPath(), "/")
		ps := strings.Split(p, "/")

		r, ok := pkgs[ps[0]]
		if ok {
			tmpls.ExecuteTemplate(w, "pkg", r)
		} else {
			tmpls.ExecuteTemplate(w, "root", pkgs)
		}
	})
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
{{define "pkg"}}
<!DOCTYPE html>
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.ImportPath}}</title>
    <meta name="go-import" content="{{.ImportPath}} git https://{{.GithubPath}}">
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
    <a href="https://{{.GithubPath}}">Source</a>
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
  </head>
  <body>
  <h3><code>myitcv.io</code></h3>
  <p>
  <a href="http://blog.myitcv.org.uk">Blog</a><br/>
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
