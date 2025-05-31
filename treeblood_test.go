package treeblood

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark"
)

func TestTreeBlood(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(MathML()),
	)
	var buffer bytes.Buffer
	if err := markdown.Convert([]byte(`
Math Test
=========
$$\int_0^1 x^{-x} dx = \sum_{n=1}^\infty n^{-n}$$
	`), &buffer); err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(buffer.String()) != strings.TrimSpace(`<h1>Math Test</h1>
<p>
<math class="math-displaystyle" display="block" displaystyle="true" style="font-feature-settings: 'dtls' off;" xmlns="http://www.w3.org/1998/Math/MathML">
  <semantics>
    <mrow>
      <msubsup>
        <mo largeop="true" movablelimits="true">∫</mo>
        <mn>0</mn>
        <mn>1</mn>
      </msubsup>
      <msup>
        <mi>x</mi>
        <mrow>
          <mo>−</mo>
          <mi>x</mi>
        </mrow>
      </msup>
      <mi>d</mi>
      <mi>x</mi>
      <mo>=</mo>
      <munderover>
        <mo largeop="true" movablelimits="true">∑</mo>
        <mrow>
          <mi>n</mi>
          <mo>=</mo>
          <mn>1</mn>
        </mrow>
        <mi>∞</mi>
      </munderover>
      <msup>
        <mi>n</mi>
        <mrow>
          <mo>−</mo>
          <mi>n</mi>
        </mrow>
      </msup>
    </mrow>
    <annotation encoding="application/x-tex">\int_0^1 x^{-x} dx = \sum_{n=1}^\infty n^{-n}</annotation>
  </semantics>
</math>
</p>`) {
		t.Error("failed to render MathML")
	}
}
