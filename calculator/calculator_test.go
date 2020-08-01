package calculator

import (
	"testing"
)

func TestSum(t *testing.T) {
	a := 40.0
	b := 2.0
	want := 42.0
	got := Sum(a, b)
	if got != want {
		t.Errorf("want: %f\ngot: %f\n", want, got)
	}
}

func TestSub(t *testing.T) {
	a := 44.0
	b := 2.0
	want := 42.0
	got := Sub(a, b)
	if got != want {
		t.Errorf("want: %f\ngot: %f\n", want, got)
	}
}

func TestMul(t *testing.T) {
	a := 7.0
	b := 6.0
	want := 42.0
	got := Mul(a, b)
	if got != want {
		t.Errorf("want: %f\ngot: %f\n", want, got)
	}
}

func TestDiv(t *testing.T) {
	cases := []struct {
		name string
		a    float64
		b    float64
		want float64
		err  bool
	}{
		{
			name: "normal",
			a:    84.0,
			b:    2.0,
			want: 42.0,
			err:  false,
		},
		{
			name: "zero",
			a:    42.0,
			b:    0.0,
			want: 0.0,
			err:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Div(c.a, c.b)
			if c.err {
				if err == nil {
					t.Errorf("want: error\ngot: nil\n")
				}
			} else {
				if got != c.want {
					t.Errorf("want: %f\ngot: %f\n", c.want, got)
				}
			}
		})
	}
}
