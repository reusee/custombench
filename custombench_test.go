package custombench

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

var runCustomBench = flag.Bool("cb", false, "run custom benchmarks")

type bench struct{}

func TestMain(m *testing.M) {
	flag.Parse()

	if *runCustomBench {
		v := reflect.ValueOf(new(bench))
		typ := v.Type()
		for i := 0; i < v.NumMethod(); i++ {
			m, ok := v.Method(i).Interface().(func(*testing.B))
			if !ok {
				continue
			}
			fmt.Printf("run %s\n", typ.Method(i).Name)
			result := testing.Benchmark(m)
			fmt.Printf("%#v\n", result)
		}
	}

	os.Exit(m.Run())
}

func (_ bench) Foo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func (_ bench) Bar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", time.Now())
	}
}
