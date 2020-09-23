# math-vector
This package is a mathematical vector representation written in Go, with some basic operations.

## Operations
This package exports some basic operations as well:
- Dot Product
- Cross Product

## How to use

```
$ go get github.com/tonytcb/math-vector
```

```
import "github.com/tonytcb/math-vector"

func main() {
    vector1 := vector.New(3, 5, 4)
    vector2 := vector.New(2, 7, 5)
    
    dotProduct := vector1.DotProduct(vector2)
    crossProduct := vector1.CrossProduct(vector2)
    
    vector1.Print()
    vector2.Print()
    
    fmt.Println("dotProduct:", dotProduct)
    fmt.Println("crossProduct:", crossProduct.Format())
}
```

## Todo
- ~~Add addition operation~~
- ~~Add subtraction operation~~
- Add more operations 

## References
- https://www.mathsisfun.com/algebra/vectors.html
- https://www.mathsisfun.com/algebra/vectors-dot-product.html
- https://www.mathsisfun.com/algebra/vectors-cross-product.html
- https://en.wikipedia.org/wiki/Cross_product
- https://en.wikipedia.org/wiki/Dot_product