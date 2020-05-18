package main

const (
	godocHTML = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">
{{with .Tabtitle}}
  <title>{{html .}} - Go Documentation Server</title>
{{else}}
  <title>Go Documentation Server</title>
{{end}}
<link type="text/css" rel="stylesheet" href="{{baseURL}}/lib/godoc/style.css">
{{if .TreeView}}
<link rel="stylesheet" href="{{baseURL}}/lib/godoc/jquery.treeview.css">
{{end}}
<script>window.initFuncs = [];</script>
<script src="{{baseURL}}/lib/godoc/jquery.js" defer></script>
{{if .TreeView}}
<script src="{{baseURL}}/lib/godoc/jquery.treeview.js" defer></script>
<script src="{{baseURL}}/lib/godoc/jquery.treeview.edit.js" defer></script>
{{end}}

{{if .Playground}}
<script src="{{baseURL}}/lib/godoc/playground.js" defer></script>
{{end}}
{{with .Version}}<script>var goVersion = {{printf "%q" .}};</script>{{end}}
<script src="{{baseURL}}/lib/godoc/godocs.js" defer></script>
</head>
<body>

<div id='lowframe' style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
...
</div><!-- #lowframe -->

<div id="topbar"{{if .Title}} class="wide"{{end}}><div class="container">
<div class="top-heading" id="heading-wide"><a href="{{baseURL}}/pkg/">Go Documentation Server</a></div>
<div class="top-heading" id="heading-narrow"><a href="{{baseURL}}/pkg/">GoDoc</a></div>
<a href="#" id="menu-button"><span id="menu-button-arrow">&#9661;</span></a>
<form method="GET" action="/search">
<div id="menu">
{{if (and .Playground .Title)}}
<a id="playgroundButton" href="http://play.golang.org/" title="Show Go Playground">Play</a>
{{end}}
<span class="search-box"><input type="search" id="search" name="q" placeholder="Search" aria-label="Search" required><button type="submit"><span><!-- magnifying glass: --><svg width="24" height="24" viewBox="0 0 24 24"><title>submit search</title><path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/><path d="M0 0h24v24H0z" fill="none"/></svg></span></button></span>
</div>
</form>

</div></div>

{{if .Playground}}
<div id="playground" class="play">
	<div class="input"><textarea class="code" spellcheck="false">package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}</textarea></div>
	<div class="output"></div>
	<div class="buttons">
		<a class="run" title="Run this code [shift-enter]">Run</a>
		<a class="fmt" title="Format this code">Format</a>
		{{if not $.GoogleCN}}
		<a class="share" title="Share this code">Share</a>
		{{end}}
	</div>
</div>
{{end}}

<div id="page"{{if .Title}} class="wide"{{end}}>
<div class="container">

{{if or .Title .SrcPath}}
  <h1>
    {{html .Title}}
    {{html .SrcPath | srcBreadcrumb}}
  </h1>
{{end}}

{{with .Subtitle}}
  <h2>{{html .}}</h2>
{{end}}

{{with .SrcPath}}
  <h2>
    Documentation: {{html . | srcToPkgLink}}
  </h2>
{{end}}

{{/* The Table of Contents is automatically inserted in this <div>.
     Do not delete this <div>. */}}
<div id="nav"></div>

{{/* Body is HTML-escaped elsewhere */}}
{{printf "%s" .Body}}

<div id="footer">
Build version {{html .Version}}.<br>
</div>

</div><!-- .container -->
</div><!-- #page -->
</body>
</html>
`
)
