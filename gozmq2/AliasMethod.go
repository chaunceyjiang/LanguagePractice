package main

type AliasTable struct {
	Prab  []float64
	Alias []float64
}

func NewAliasTable(p []float64) *AliasTable {
	l := len(p)
	a := &AliasTable{
		Prab:  make([]float64, l),
		Alias: make([]float64, l),
	}
	pp := make([]float64, l)
	copy(pp, p)
	for i, v := range pp {
		pp[i] = v * float64(l)
	}

	return a
}

func main() {
	p := []float64{1 / 2.0, 1 / 4.0, 1 / 12.0, 1 / 12.0}

}
