package helper

import (
	"fmt"
	"runtime"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Lutfan")

	if result != "Hello Lutfan" {
		// error
		// t.Fail()
		t.Error("Result should be 'Hello Lutfan'")
	}

	fmt.Println("Test Hello World Done")
}

func TestHelloWorldZainul(t *testing.T) {
	result := HelloWorld("Zainul")

	if result != "Hello Zainul" {
		// t.FailNow()
		t.Fatal("Result must be 'Hello Zainul'")
	}

	fmt.Println("TestHelloWorldZainul Done")
}

func TestHeloWorldAssert(t *testing.T) {
	result := HelloWorld("Zainul")

	// t.Assert(result == "Hello Zainul", "Result must be 'Hello Zainul'")
	assert.Equal(t, "Hello Zainul", result, "Result must be 'Hello Zainul'")
	fmt.Println("TestHeloWorldAssert Done")
}

func TestHeloWorldRequire(t *testing.T) {
	result := HelloWorld("Zainul")

	require.Equal(t, "Hello Zainul", result, "Result must be 'Hello Zainul'")
	fmt.Println("TestHeloWorldRequire Done")
}

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")
	}

	result := HelloWorld("Zainul")

	require.Equal(t, "Hello Zainul", result, "Result must be 'Hello Zainul'")
	fmt.Println("TestHeloWorldRequire Done")
}

func TestMain(m *testing.M){
	// before
	fmt.Println("Before UNIT TEST")

	m.Run()
	// after
	fmt.Println("After UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Zainul", func(t *testing.T) {
		result := HelloWorld("Zainul")
		require.Equal(t, "Hello Zainul", result, "Result must be 'Hello Zainul'")
	})

	t.Run("Lutfan", func(t *testing.T) {
		result := HelloWorld("Lutfan")
		require.Equal(t, "Hello Lutfan", result, "Result must be 'Hello Lutfan'")
	})
}


func TestTableHellowWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		want string
	}{
		{
			name: "Zainul",
			request: "Zainul",
			want: "Hello Zainul",
		},
		{
			name: "Lutfan",
			request: "Lutfan",
			want: "Hello Lutfan",
		},
		{
			name: "Budi",
			request: "Budi",
			want: "Hello Budi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.want, result, "Result must be 'Hello Zainul'")
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Zainul")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Zainul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Zainul")
		}
	})

	b.Run("Lutfan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Lutfan")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	tests := []struct {
		name string
		request string
		want string
	}{
		{
			name: "Zainul",
			request: "Zainul",
			want: "Hello Zainul",
		},
		{
			name: "Lutfan",
			request: "Lutfan",
			want: "Hello Lutfan",
		},
		{
			name: "Budi",
			request: "Budi",
			want: "Hello Budi",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test.request)
			}
		})
	}
}