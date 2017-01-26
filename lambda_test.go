package λ_test

import (
	"."
	"testing"
)

var λExamples = map[λ.Λ][]λ.Λ{
	λ.NewΛ("x",
		func(μ λ.Μ) λ.Λ {
			return μ.NewΛ("y",
				func(μ λ.Μ) λ.Λ {
					return μ.Read("x")
				},
			)
		},
	): {
		λ.I,
		λ.NewΛ("x",
			func(μ λ.Μ) λ.Λ {
				return μ.NewΛ("y",
					func(μ λ.Μ) λ.Λ {
						return μ.Read("y")
					},
				)
			},
		),
	},
	λ.NewΛ("x",
		func(μ λ.Μ) λ.Λ {
			return μ.NewΛ("y",
				func(μ λ.Μ) λ.Λ {
					return μ.Read("y")
				},
			)
		},
	): {
		λ.NewΛ("x",
			func(μ λ.Μ) λ.Λ {
				return μ.NewΛ("y",
					func(μ λ.Μ) λ.Λ {
						return μ.Read("y")
					},
				)
			},
		),
		λ.I,
	},
}

func TestΛCalculus(t *testing.T) {
	for example, inputs := range λExamples {
		fn := example

		for _, arg := range inputs {
			fn = fn.Call(arg)
		}

		if fn != λ.I {
			t.Errorf("expected equivalent lambdas")
		}
	}
}
