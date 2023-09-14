package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
problem solver
*/

func solve(N, M int, H, A, B []int) interface{} {
	debug.Println(N, M, H, A, B)

	var (
		graph = NewGraph()
		stack = NewSliceStack()
		ans   int
	)

	for id, height := range H {
		node := NewNode()
		node.height = height
		graph.AddNode(id+1, node)
		stack.Push(id + 1)
	}

	for i := 0; i < M; i++ {
		succ := NewEdge()
		pred := NewEdge()
		succ.successors = append(succ.successors, B[i])
		pred.successors = append(pred.successors, A[i])
		graph.AddEdge(A[i], succ)
		graph.AddEdge(B[i], pred)
	}

	for stack.Length() != 0 {
		head := stack.Pop()
		check := true

		if graph.edges[head] == nil {
			ans++
			continue
		}

		for _, succ := range graph.edges[head].successors {
			check = check && graph.nodes[head].height > graph.nodes[succ].height
		}

		debug.Println(check, head)
		if check {
			ans++
		}
	}

	return ans
}

func main() {
	var (
		N, M    int
		H, A, B []int
	)
	fmt.Scan(&N, &M)
	input.Scan()
	H, _ = InputListInt()
	for input.Scan() {
		line, _ := InputListInt()
		A = append(A, line[0])
		B = append(B, line[1])
	}
	ans := solve(N, M, H, A, B)
	fmt.Println(ans)
}

type SliceStack struct {
	stack []interface{}
}

func NewSliceStack(vars ...interface{}) *SliceStack {
	stack := new(SliceStack)
	stack.stack = append(make([]interface{}, 0), vars...)
	return stack
}

func (stack *SliceStack) Push(val ...interface{}) {
	stack.stack = append(stack.stack, val...)
}

func (stack *SliceStack) Pop() interface{} {
	res := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return res
}

func (stack *SliceStack) Length() int {
	return len(stack.stack)
}

type SliceQueue struct {
	queue []interface{}
}

func NewSliceQueue(vars ...interface{}) *SliceQueue {
	queue := new(SliceQueue)
	queue.queue = append(make([]interface{}, 0), vars...)
	return queue
}

func (queue *SliceQueue) Push(val ...interface{}) {
	queue.queue = append(queue.queue, val...)
}

func (queue *SliceQueue) Pop() interface{} {
	res := queue.queue[0]
	queue.queue = queue.queue[1:]
	return res
}

func (queue *SliceQueue) Length() int {
	return len(queue.queue)
}

type Graph struct {
	nodes map[interface{}]*Node
	edges map[interface{}]*Edge
}

func NewGraph() Graph {
	newGraph := new(Graph)
	newGraph.nodes = make(map[interface{}]*Node)
	newGraph.edges = make(map[interface{}]*Edge)
	return *newGraph
}

func (g *Graph) AddNode(id interface{}, node *Node) {
	g.nodes[id] = node
}

func (g *Graph) AddEdge(root interface{}, edge *Edge) {
	if g.edges[root] == nil {
		g.edges[root] = NewEdge()
	}
	g.edges[root].successors = append(g.edges[root].successors, edge.successors...)
}

type Node struct {
	// if you want to add fields, you can do below.
	height int
}

func NewNode() *Node {
	return new(Node)
}

type Edge struct {
	successors []interface{}
	// if you want to add fields, you can do below.
}

func NewEdge() *Edge {
	return new(Edge)
}

/*
initialize common variables
*/

var (
	input             = bufio.NewScanner(os.Stdin)
	output, debug     = log.New(os.Stdout, "", 0), log.New(os.Stderr, "DEBUG : ", log.Lshortfile)
	YES, Yes, yes     = "YES", "Yes", "yes"
	NO, No, no        = "NO", "No", "no"
	Even              = "Even"
	Odd               = "Odd"
	mod, mod998244353 = 1e9 + 7, 998244353
)

func init() {
	input.Buffer(make([]byte, 1<<30), 1<<30)
}

/*
Input / Output
*/

func dump(variable ...interface{}) {
	for _, v := range variable {
		debug.Printf("%v : %v\n", reflect.TypeOf(v), v)
	}
}

func InputListInt() ([]int, error) {
	var (
		res []int
	)

	for _, s := range strings.Split(input.Text(), " ") {
		n, err := strconv.ParseInt(s, 0, 0)

		if err != nil {
			return []int{}, err
		}

		res = append(res, int(n))
	}

	return res, nil
}

/*
Data Structures
*/

type Set map[interface{}]bool

func (st *Set) NewSet(vars ...interface{}) *Set {
	var s = make(Set)
	for _, v := range vars {
		s[v] = true
	}
	return &s
}

func (st *Set) SetKeys() []interface{} {
	var keys []interface{}

	for k, _ := range *st {
		keys = append(keys, k)
	}

	return keys
}

func (st *Set) Add(v interface{}) { (*st)[v] = true }

func (st *Set) Exist(v interface{}) bool { return (*st)[v] }

func (st *Set) Empty() bool { return len(*st) == 0 }

func (st *Set) Union(y Set) Set {
	var res = make(Set)

	for k, _ := range *st {
		res[k] = true
	}

	for k, _ := range y {
		res[k] = true
	}

	return res
}

func (st *Set) Difference(y Set) Set {
	var res = make(Set)

	for k, _ := range *st {
		if !y[k] {
			res[k] = true
		}
	}

	return res
}

func (st *Set) Intersection(y Set) Set {
	var res = make(Set)

	for k, _ := range *st {
		if y[k] {
			res[k] = true
		}
	}

	return res
}

func (st *Set) IsEqual(y Set) bool {
	var (
		x_keys            = st.SetKeys()
		y_keys            = y.SetKeys()
		intersection_xy   = st.Intersection(y)
		intersection_keys = intersection_xy.SetKeys()
	)

	return len(x_keys) == len(intersection_keys) && len(y_keys) == len(intersection_keys)
}

func (st *Set) IsSubset(x, y Set) bool { return len(y.Difference(x)) == 0 }

/*
Character / String
- String Reversal
*/

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

/*
Convert value
- Absolute value
- Power
*/

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func AbsFloat64(x float64) float64 {
	return math.Abs(x)
}

func ByteToInt(r byte) int {
	return int(r - '0')
}

func PowInt(a, n int) int {
	return int(math.Pow(float64(a), float64(n)))
}

func PowInt64(a, n int) int64 {
	return int64(math.Pow(float64(a), float64(n)))
}

/*
Number theoretic algorithm
- total
- Cumulative Sum
- Arithmetic Progressions
- Geometic Progressions
- Greatest Common Divisor
- Least Common Multipul
*/

func SumInt(vars ...int) int {
	var (
		sum int
	)

	for _, v := range vars {
		sum += v
	}

	return sum

}

func SumInt64(vars ...int64) int64 {
	var (
		sum int64
	)

	for _, v := range vars {
		sum += v
	}

	return sum
}

func CumulativeSumInt(vars []int) []int {
	var (
		cumsum = make([]int, len(vars))
	)

	if len(vars) != 0 {
		copy(cumsum, vars)

		for i, v := range vars[1:] {
			index := i + 1
			cumsum[index] = cumsum[index-1] + v
		}

	}

	return cumsum
}

func ArithmeticProgressionsSumInt(a0, d, n int) (sum int64, err error) {
	if n < 0 {
		sum, err = 0, errors.New("number of terms must not be negative.")
	} else {
		var (
			a0_i64 = int64(a0)
			d_i64  = int64(d)
			n_i64  = int64(n)
		)
		sum, err = (a0_i64*n_i64)+n_i64*(n_i64-1)/2*d_i64, nil
	}

	return sum, err
}

func GeometicProgressionsSumInt(a0, r, n int) (sum int64, err error) {
	if n < 0 {
		sum, err = 0, errors.New("number of terms must not be negative.")
	} else {
		var (
			a0_f64 = float64(a0)
			r_f64  = float64(r)
			n_f64  = float64(n)
		)
		sum, err = int64(a0_f64*(math.Pow(r_f64, n_f64)-1)/(r_f64-1)), nil
	}

	return sum, err
}

func GreatestCommonDivisorInt(x, y int) int {
	if y == 0 {
		return x
	} else {
		return GreatestCommonDivisorInt(y, x%y)
	}
}

func GreatestCommonDivisorIntegers(vars ...int) (gcd int, err error) {
	if len(vars) == 0 {
		return -1, errors.New("vars length is 0")
	} else {
		gcd, err = vars[0], nil

		for i := 1; i < len(vars); i++ {
			gcd = GreatestCommonDivisorInt(gcd, vars[i])
		}

		return gcd, err
	}
}

func LeastCommonMultipulInt(x, y int) int {
	return x / GreatestCommonDivisorInt(x, y) * y
}

func LeastCommonMultipulIntegers(vars ...int) (lcm int, err error) {
	if len(vars) == 0 {
		return -1, errors.New("vars length is 0")
	} else {
		lcm, err = vars[0], nil

		for i := 1; i < len(vars); i++ {
			lcm = LeastCommonMultipulInt(lcm, vars[i])
		}

		return lcm, nil
	}
}

/*
Comparison of Values
- larger / smaller
- largest / smallest
*/

// Get larger values

func MaxInt(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MaxInt64(x, y int64) int64 {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MaxFloat64(x, y float64) float64 {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MaxIntegers(vars ...int) (int, error) {
	var (
		res int
		err error
	)

	if len(vars) == 0 {
		res, err = math.MaxInt32, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MaxInt(res, v)
		}
	}

	return res, err
}

func MaxInteger64s(vars ...int64) (int64, error) {
	var (
		res int64
		err error
	)

	if len(vars) == 0 {
		res, err = math.MaxInt64, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MaxInt64(res, v)
		}
	}

	return res, err
}

func MaxFloat64s(vars ...float64) (float64, error) {
	var (
		res float64
		err error
	)

	if len(vars) == 0 {
		res, err = math.MaxFloat64, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MaxFloat64(res, v)
		}
	}

	return res, err
}

// Get smaller values

func MinInt(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func MinInt64(x, y int64) int64 {
	if x <= y {
		return x
	} else {
		return y
	}
}

func MinFloat64(x, y float64) float64 {
	if x <= y {
		return x
	} else {
		return y
	}
}

func MinIntegers(vars ...int) (int, error) {
	var (
		res int
		err error
	)

	if len(vars) == 0 {
		res, err = math.MinInt32, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MinInt(res, v)
		}
	}

	return res, err
}

func MinInteger64s(vars ...int64) (int64, error) {
	var (
		res int64
		err error
	)

	if len(vars) == 0 {
		res, err = math.MinInt64, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MinInt64(res, v)
		}
	}

	return res, err
}

func MinFloat64s(vars ...float64) (float64, error) {
	var (
		res float64
		err error
	)

	if len(vars) == 0 {
		res, err = math.MaxFloat64, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MinFloat64(res, v)
		}
	}

	return res, err
}

/*
Search Algorithm
- Binary Search
*/

func BinarySearch(negative, positive, dist interface{},
	IsContinue func(negative, positive, dist interface{}) bool,
	HowToMiddle func(negative, positive interface{}) interface{},
	IsPositive func(mid interface{}) bool) (interface{}, interface{}) {

	for IsContinue(negative, positive, dist) {
		mid := HowToMiddle(negative, positive)

		if IsPositive(mid) {
			positive = mid
		} else {
			negative = mid
		}
	}

	return negative, positive
}
