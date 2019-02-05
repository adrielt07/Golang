package pack

import (
	"bytes"
	"testing"
	"text/template"
)

// A benchmark test example
func BenchmarkExample(b *testing.B) {
	temp, _ := template.New("").Parse("Hello, Go")

	/*
		b.ResetTimer()
			This resets the time so the creation of the template
			is not factored in to the benchmark
	*/
	b.ResetTimer()
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		temp.Execute(&buf, nil)
		buf.Reset()
	}
}
