package generate

import (
	"fmt"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/gdziwoki/go/gopages/internal/flags"
	"github.com/gdziwoki/go/gopages/internal/pipe"
	"github.com/pkg/errors"
	"golang.org/x/tools/godoc"
	"golang.org/x/tools/godoc/vfs"
)

func addGoPagesFuncs(funcs template.FuncMap, modulePackage string, args flags.Args) {
	funcs["node_html"] = nodeHTML(funcs["node_html"].(node_htmlFunc), args.BaseURL)

	longTitle := fmt.Sprintf("%s | %s", args.SiteTitle, args.SiteDescription)
	if args.SiteTitle == "" || args.SiteDescription == "" {
		longTitle = ""
	}
	values := map[string]interface{}{
		"BaseURL":       args.BaseURL,
		"ModuleURL":     path.Join(args.BaseURL, "/pkg", modulePackage) + "/",
		"SiteTitle":     args.SiteTitle,
		"SiteTitleLong": longTitle,
	}
	funcs["gopages"] = func(defaultValue, firstKey string, keys ...string) (string, error) {
		keys = append([]string{firstKey}, keys...) // require at least one key
		for _, key := range keys {
			value, ok := values[key]
			var valueStr string
			err := pipe.ChainFuncs(
				func() error {
					return pipe.ErrIf(!ok, errors.Errorf("Unknown gopages key: %q", key))
				},
				func() error {
					var isString bool
					valueStr, isString = value.(string)
					return pipe.ErrIf(!isString, errors.Errorf("gopages key %q is not a string", key))
				},
			).Do()
			if err != nil || valueStr != "" {
				return template.HTMLEscapeString(valueStr), err
			}
		}
		return defaultValue, nil
	}
	funcs["gopagesWatchScript"] = func() string {
		script := fmt.Sprintf(`
<script>
const startDate = %q
const timeoutMillis = 2000
const poll = () => {
	fetch(window.location).then(resp => {
		const newDate = resp.headers.get("GoPages-Last-Updated")
		if (newDate != startDate) {
			window.location.reload()
		}
	})
}
window.setInterval(poll, timeoutMillis)
</script>
`, time.Now().Format(time.RFC3339))
		if !args.Watch {
			script = ""
		}
		return script
	}
}

func readTemplates(pres *godoc.Presentation, funcs template.FuncMap, fs vfs.FileSystem) {
	pres.CallGraphHTML = readTemplate(funcs, fs, "callgraph.html")
	pres.DirlistHTML = readTemplate(funcs, fs, "dirlist.html")
	pres.ErrorHTML = readTemplate(funcs, fs, "error.html")
	pres.ExampleHTML = readTemplate(funcs, fs, "example.html")
	pres.GodocHTML = parseTemplate(funcs, "godoc.html", godocHTML)
	pres.ImplementsHTML = readTemplate(funcs, fs, "implements.html")
	pres.MethodSetHTML = readTemplate(funcs, fs, "methodset.html")
	pres.PackageHTML = readTemplate(funcs, fs, "package.html")
	pres.PackageRootHTML = readTemplate(funcs, fs, "packageroot.html")
}

func readTemplate(funcs template.FuncMap, fs vfs.FileSystem, name string) *template.Template {
	// use underlying file system fs to read the template file
	// (cannot use template ParseFile functions directly)
	data, err := vfs.ReadFile(fs, path.Join("lib/godoc", name))
	if err != nil {
		panic(err)
	}
	return parseTemplate(funcs, name, string(data))
}

func parseTemplate(funcs template.FuncMap, name, data string) *template.Template {
	return template.Must(template.New(name).Funcs(funcs).Parse(data))
}

type node_htmlFunc = func(info *godoc.PageInfo, node interface{}, linkify bool) string

// nodeHTML runs the original 'node_html' template func, then rewrites any links inside it
func nodeHTML(original node_htmlFunc, baseURL string) node_htmlFunc {
	pkgURL := path.Join(baseURL, "/pkg")
	return func(info *godoc.PageInfo, node interface{}, linkify bool) string {
		s := original(info, node, linkify)
		return strings.ReplaceAll(s, `<a href="/pkg/`, fmt.Sprintf(`<a href="%s/`, pkgURL))
	}
}
