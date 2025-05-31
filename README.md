# goldmark-treeblood
goldmark-treeblood is an extension for [goldmark](http://github.com/yuin/goldmark) that renders $LaTeX$ expressions as
MathML, an open web standard for describing mathematical notation. Unlike MathJax or KaTeX, [TreeBlood](https://github.com/Wyatt915/treeblood)
runs server-side, is written in pure Go, and has no dependencies outside the Go standard library.

## Usage

To use goldmark-treeblood, import the `goldmark-treeblood` package
```go
import "github.com/wyatt915/goldmark-treeblood"
```
and include `treeblood.MathML` in the list of extensions passed to goldmark
```go
goldmark.New(
  goldmark.WithExtensions(
    // ...
    treeblood.MathML(),
  ),
  // ...
).Convert(src, out)
```

This extension respects both TeX-style delimiters `$...$` and `$$...$$` as well as their more modern AMS $LaTeX$
counterparts `\(...\)` and `\[...\]`.
