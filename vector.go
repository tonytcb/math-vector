package vector

import "fmt"

// Vector represents a 3D vector
type Vector struct {
	i int
	j int
	k int
}

// New build a valid Vector struct
func New(i int, j int, k int) Vector {
	return Vector{i: i, j: j, k: k}
}

// I returns the value of i
func (v Vector) I() int {
	return v.i
}

// J returns the value of j
func (v Vector) J() int {
	return v.j
}

// K returns the value of k
func (v Vector) K() int {
	return v.k
}

// Format return the vector formatted as {i, j, k}
func (v Vector) Format() string {
	return fmt.Sprintf("{%d, %d, %d}", v.I(), v.J(), v.K())
}

// Print writes the formatted vector in the output
func (v Vector) Print() {
	fmt.Println(v.Format())
}

// DotProduct reproduces the algebraic operation dot product, aka scalar product, resulting in a integer
func (v Vector) DotProduct(v2 Vector) int {
	return (v.i * v2.i) + (v.j * v2.j) + (v.k * v2.k)
}

// CrossProduct reproduces the algebraic operation cross product, resulting in an new vector
func (v Vector) CrossProduct(v2 Vector) Vector {
	i := (v.j * v2.k) - (v.k * v2.j)
	j := (v.k * v2.i) - (v.i * v2.k)
	k := (v.i * v2.j) - (v.j * v2.i)

	return New(i, j, k)
}
