package indirect

import (
	"bytes"
	"testing"
	"unsafe"
)

type t0 int
type t1 string
type t2 []byte
type t3 struct {
	a uint64
	b float32
}

func TestIndirectUnsafePtr(t *testing.T) {
	t.Run("t0", func(t *testing.T) {
		x := t0(15)
		p := uintptr(unsafe.Pointer(&x))
		y := ToUnsafePtr(p)
		z := *(*t0)(y)
		if x != z {
			t.FailNow()
		}
	})
	t.Run("t1", func(t *testing.T) {
		x := t1("foobar")
		p := uintptr(unsafe.Pointer(&x))
		y := ToUnsafePtr(p)
		z := *(*t1)(y)
		if x != z {
			t.FailNow()
		}
	})
	t.Run("t2", func(t *testing.T) {
		x := t2("foobar")
		p := uintptr(unsafe.Pointer(&x))
		y := ToUnsafePtr(p)
		z := *(*t2)(y)
		if !bytes.Equal(x, z) {
			t.FailNow()
		}
	})
	t.Run("t3", func(t *testing.T) {
		x := t3{
			a: 4562344,
			b: 3.1415,
		}
		p := uintptr(unsafe.Pointer(&x))
		y := ToUnsafePtr(p)
		z := *(*t3)(y)
		if x.a != z.a || x.b != z.b {
			t.FailNow()
		}
	})
}

func BenchmarkIndirectUnsafePtr(b *testing.B) {
	b.Run("t0", func(b *testing.B) {
		b.ReportAllocs()
		x := t0(15)
		p := uintptr(unsafe.Pointer(&x))
		for i := 0; i < b.N; i++ {
			y := ToUnsafePtr(p)
			z := *(*t0)(y)
			if x != z {
				b.FailNow()
			}
		}
	})
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		x := t1("foobar")
		p := uintptr(unsafe.Pointer(&x))
		for i := 0; i < b.N; i++ {
			y := ToUnsafePtr(p)
			z := *(*t1)(y)
			if x != z {
				b.FailNow()
			}
		}
	})
	b.Run("t2", func(b *testing.B) {
		b.ReportAllocs()
		x := t2("foobar")
		p := uintptr(unsafe.Pointer(&x))
		for i := 0; i < b.N; i++ {
			y := ToUnsafePtr(p)
			z := *(*t2)(y)
			if !bytes.Equal(x, z) {
				b.FailNow()
			}
		}
	})
	b.Run("t3", func(b *testing.B) {
		b.ReportAllocs()
		x := t3{
			a: 4562344,
			b: 3.1415,
		}
		p := uintptr(unsafe.Pointer(&x))
		y := ToUnsafePtr(p)
		for i := 0; i < b.N; i++ {
			z := *(*t3)(y)
			if x.a != z.a || x.b != z.b {
				b.FailNow()
			}
		}
	})
}
