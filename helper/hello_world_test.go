package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Gita", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Gita")
		}
	})
	b.Run("Adit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Adit")
		}
	})
}
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{{name: "adit", request: "adit"}, {name: "Gita", request: "Gita"}, {name: "joko", request: "joko"}, {name: "Beni", request: "Beni"}, {name: "Prigi", request: "Prigi"}}
	for _, bencmark := range benchmarks {
		b.Run(bencmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bencmark.request)
			}
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("======Before======")
	m.Run()
	fmt.Println("==========After============")
}
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Gita")
	require.Equal(t, "Hi Gita", result)

	fmt.Println("sukses test Gita")

}

func TestHelloAdit(t *testing.T) {
	result := HelloWorld("Adit")
	assert.Equal(t, "Hi Adit", result)

}

func TestSubHeello(t *testing.T) {
	t.Run("Putra", func(t *testing.T) {
		result := HelloWorld("Putra")
		assert.Equal(t, "Hi Putra", result)
	})
	t.Run("Prigi", func(t *testing.T) {
		result := HelloWorld("Prigi")
		assert.Equal(t, "Hi Prigi", result)
	})

}
