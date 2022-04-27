package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Azzin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Azzin")
		}
	})

	b.Run("Maharil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Maharil")
		}
	})
}

//run benchmark = go test -v -bench=.  //With unit test
// only benchmark = go test -v -run=NotMathUnitTest -bench=.
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Azzin")
	}
}

func BenchmarkHelloWorldMaharil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Maharil")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Azzin",
			request:  "Azzin",
			expected: "Hello Azzin",
		},
		{
			name:     "Maharil",
			request:  "Maharil",
			expected: "Hello Maharil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTes(t *testing.T) {
	t.Run("Azzin", func(t *testing.T) {

		result := HelloWorld("Azzin")
		require.Equal(t, "Hello Azzin", result, "Reslut must be 'Hello Azzin'")

	})
	t.Run("Maharil", func(t *testing.T) {

		result := HelloWorld("Maharil")
		require.Equal(t, "Hello Maharil", result, "Reslut must be 'Hello Maharil'")
	})
}

func TestMain(m *testing.M) {
	//After
	fmt.Println("After Unit Test")

	m.Run()

	fmt.Println("Before Unit Test")
	//before
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Only run in mac os")
	}
	result := HelloWorld("Azzin")
	require.Equal(t, "Hello Azzin", result, "Reslut must be 'Hello Azzin'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Azzin")
	require.Equal(t, "Hello Azzin", result, "Reslut must be 'Hello Azzin'")
	fmt.Println("Test Hello World require")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Azzin")
	assert.Equal(t, "Hello Azzin", result, "Reslut must be 'Hello Azzin'")
	fmt.Println("Test Hello World assert")
}

func TestHelloWorldAzzin(t *testing.T) {
	result := HelloWorld("Azzin")
	if result != "Hello Azzin" {
		//Unit Test
		t.Error("Test Must Be 'Hello Azin'")
	}
	fmt.Println("TestHelloWorldAzzin Done")
}

func TestHelloWorldMaharil(t *testing.T) {
	result := HelloWorld("Maharil")
	if result != "Hello Maharil" {
		//Unit Test
		t.Fatal("Test Must Be'Hello Maharil'")
	}
	fmt.Println("TestHelloWorldMaharil Done")
}
