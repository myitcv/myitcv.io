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
}

var ps = []pkg{
	pkg{
		RelPath:    "react",
		GithubPath: "github.com/myitcv/react",
	},
	pkg{
		RelPath:    "remarkable",
		GithubPath: "github.com/myitcv/remarkable",
	},
	pkg{
		RelPath:    "immutable",
		GithubPath: "github.com/myitcv/immutable",
	},
	pkg{
		RelPath:    "sorter",
		GithubPath: "github.com/myitcv/sorter",
	},
	pkg{
		RelPath:    "gogenerate",
		GithubPath: "github.com/myitcv/gogenerate",
	},
	pkg{
		RelPath:    "gg",
		GithubPath: "github.com/myitcv/gg",
	},
	pkg{
		RelPath:    "gai",
		GithubPath: "github.com/myitcv/gai",
	},
	pkg{
		RelPath:    "gitgodoc",
		GithubPath: "github.com/myitcv/gitgodoc",
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

		r, ok := pkgs[p]
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
    <title>{{.ImportPath}}</title>
    <meta name="go-import" content="{{.ImportPath}} git https://{{.GithubPath}}">
  </head>
  <body>
  <p>
	 <b><code>go get -u {{.ImportPath}}</code></b><br/><br/>
    <a href="https://godoc.org/{{.ImportPath}}">Documentation</a><br/>
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
  <title>myitcv.io</title>
  </head>
  <body>
  <h3><code>myitcv.io</code></h3>
  <p><a href="http://blog.myitcv.org.uk">Blog</a></p>
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
