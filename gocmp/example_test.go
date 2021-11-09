package gocmp

import (
	"errors"
	"math"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCmp(t *testing.T) {
	type ComplexObject struct {
		ID          int64
		Name        string
		Score       float64
		Public      bool
		RelationIds []int64
	}
	type ComplexObjects []ComplexObject

	type Compare struct {
		Name    string
		Value   int
		Objects ComplexObjects
	}

	v1 := &Compare{
		Name:  "Tom",
		Value: 100,
		Objects: ComplexObjects{
			{
				ID:          1,
				Name:        "xyz",
				Score:       1.0,
				Public:      true,
				RelationIds: []int64{1, 2, 3, 4, 5},
			},
			{
				ID:          2,
				Name:        "xxxxx",
				Score:       99.9,
				Public:      true,
				RelationIds: []int64{10, 20, 30, 40, 50},
			},
		},
	}
	v2 := &Compare{
		Name:  "Andrew",
		Value: 50,
		Objects: ComplexObjects{
			{
				ID:          1,
				Name:        "xyz",
				Score:       1.0,
				Public:      true,
				RelationIds: []int64{1, 2, 3, 4, 5},
			},
			{
				ID:          3,
				Name:        "abc",
				Score:       11.1,
				Public:      false,
				RelationIds: []int64{6, 7, 8, 9, 10},
			},
		},
	}
	if diff := cmp.Diff(v1, v2); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestAllowUnexported(t *testing.T) {
	type Compare struct {
		Exported   int
		unexported int
	}

	v1 := &Compare{
		Exported:   100,
		unexported: 1,
	}
	v2 := &Compare{
		Exported:   100,
		unexported: 2,
	}
	opt := cmp.AllowUnexported(Compare{})
	if diff := cmp.Diff(v1, v2, opt); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestIgnoreUnexported(t *testing.T) {
	type Compare struct {
		Exported   int
		unexported int
	}

	v1 := &Compare{
		Exported:   100,
		unexported: 1,
	}
	v2 := &Compare{
		Exported:   100,
		unexported: 2,
	}
	// func IgnoreUnexported(typs ...interface{}) cmp.Option
	opt := cmpopts.IgnoreUnexported(Compare{})
	if diff := cmp.Diff(v1, v2, opt); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestIgnoreFields(t *testing.T) {
	type Compare struct {
		Exported  int
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	v1 := &Compare{
		Exported:  100,
		CreatedAt: time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	v2 := &Compare{
		Exported:  100,
		CreatedAt: time.Date(3100, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(3100, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	// func IgnoreFields(typ interface{}, names ...string) cmp.Option
	opts := []cmp.Option{
		cmpopts.IgnoreFields(Compare{}, "CreatedAt", "UpdatedAt"),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestIgnoreTypes(t *testing.T) {
	type Compare struct {
		Exported  int
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	v1 := &Compare{
		Exported:  100,
		CreatedAt: time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	v2 := &Compare{
		Exported:  100,
		CreatedAt: time.Date(3100, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(3100, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	// func IgnoreTypes(typs ...interface{}) cmp.Option
	opts := []cmp.Option{
		cmpopts.IgnoreTypes(time.Time{}),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestIgnoreSliceElements(t *testing.T) {
	type Compare struct {
		Exported int
		Slice    []int
	}

	v1 := &Compare{
		Exported: 100,
		Slice:    []int{1, 2, 3, 4, 5},
	}
	v2 := &Compare{
		Exported: 100,
		Slice:    []int{1, 9, 2, 9, 3, 9, 4, 9, 5},
	}
	// func IgnoreSliceElements(discardFunc interface{}) cmp.Option
	opts := []cmp.Option{
		cmpopts.IgnoreSliceElements(func(elem int) bool {
			return elem == 9
		}),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestIgnoreMapEntries(t *testing.T) {
	type Compare struct {
		Exported int
		Map      map[string]int
	}

	v1 := &Compare{
		Exported: 100,
		Map: map[string]int{
			"A": 1,
			"B": 1,
		},
	}
	v2 := &Compare{
		Exported: 100,
		Map: map[string]int{
			"B": 1,
			"A": 1,
			"X": 9999,
		},
	}
	// func IgnoreMapEntries(discardFunc interface{}) cmp.Option
	opts := []cmp.Option{
		cmpopts.IgnoreMapEntries(func(key string, val int) bool {
			return key == "X" && val == 9999
		}),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

type InterfaceA interface {
	Get() int
}

type InterfaceAImpl struct {
	X int
}

func (i *InterfaceAImpl) Get() int {
	return i.X
}

type InterfaceB interface {
	Set(x int)
}

type InterfaceBImpl struct {
	X int
}

func (i *InterfaceBImpl) Set(x int) {
	i.X = x
}

func TestIgnoreInterfaces(t *testing.T) {

	type Compare struct {
		Exported int
		IA       InterfaceA
		IB       InterfaceB
	}

	v1 := &Compare{
		Exported: 100,
		IA:       &InterfaceAImpl{X: 1},
		IB:       &InterfaceBImpl{X: 1},
	}
	v2 := &Compare{
		Exported: 100,
		IA:       &InterfaceAImpl{X: 100},
		IB:       &InterfaceBImpl{X: 1},
	}
	// func IgnoreInterfaces(ifaces interface{}) cmp.Option
	opts := []cmp.Option{
		cmpopts.IgnoreInterfaces(struct{ InterfaceA }{}),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}
func TestEquoteEmpty(t *testing.T) {

	type Compare struct {
		Exported int
		Slice    []int
		Map      map[string]int
	}

	v1 := &Compare{
		Exported: 100,
		Slice:    []int{},
		Map:      map[string]int{},
	}
	v2 := &Compare{
		Exported: 100,
		Slice:    nil,
		Map:      nil,
	}
	// func EquateEmpty() cmp.Option
	opts := []cmp.Option{
		cmpopts.EquateEmpty(),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestEquoteNans(t *testing.T) {

	type Compare struct {
		Exported float64
	}

	v1 := &Compare{
		Exported: math.NaN(),
	}
	v2 := &Compare{
		Exported: math.NaN(),
	}
	// func EquateNaNs() cmp.Option
	opts := []cmp.Option{
		cmpopts.EquateNaNs(),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestEquoteErrors(t *testing.T) {

	type Compare struct {
		Error error
	}
	err := errors.New("error occured")
	v1 := &Compare{
		Error: err,
	}
	v2 := &Compare{
		Error: err,
	}
	// func EquateNaNs() cmp.Option
	opts := []cmp.Option{
		cmpopts.EquateErrors(),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestEquoteApprox(t *testing.T) {

	type Compare struct {
		Exported float64
	}
	v1 := &Compare{
		Exported: 0.1,
	}
	v2 := &Compare{
		Exported: 0.01,
	}
	// func EquateApprox(fraction, margin float64) cmp.Option
	opts := []cmp.Option{
		cmpopts.EquateApprox(0, 0.091),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestEquoteApproxTime(t *testing.T) {

	type Compare struct {
		Exported time.Time
	}
	v1 := &Compare{
		Exported: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	v2 := &Compare{
		Exported: time.Date(2021, 1, 1, 0, 0, 0, 10, time.UTC),
	}
	// func EquateApproxTime(margin time.Duration) cmp.Option
	opts := []cmp.Option{
		cmpopts.EquateApproxTime(time.Second * 10),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}
func TestSortSlices(t *testing.T) {

	type Compare struct {
		NumSlice []int
	}
	v1 := &Compare{
		NumSlice: []int{1, 2, 3, 4, 5},
	}
	v2 := &Compare{
		NumSlice: []int{5, 4, 3, 2, 1},
	}
	// func SortSlices(lessFunc interface{}) cmp.Option
	opts := []cmp.Option{
		cmpopts.SortSlices(func(i, j int) bool {
			return i < j
		}),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestSortMaps(t *testing.T) {
	v1 := map[string]int{
		"AAA": 111,
		"BBB": 222,
		"CCC": 333,
	}
	v2 := map[string]int{
		"BBB": 222,
		"CCC": 333,
		"AAA": 111,
	}

	// func SortMaps(lessFunc interface{}) cmp.Option
	opts := []cmp.Option{
		cmpopts.SortMaps(func(i, j string) bool {
			return i < j
		}),
	}
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestAcyclicTransformer(t *testing.T) {
	type Compare struct {
		SplitLines interface{}
	}
	v1 := &Compare{
		SplitLines: "1,2,3",
	}
	v2 := &Compare{
		SplitLines: "1,10,3",
	}

	// func AcyclicTransformer(name string, xformFunc interface{}) cmp.Option
	opts := []cmp.Option{
		// cmpopts.AcyclicTransformer("SplitLines", func(s string) []string {
		// 	return strings.Split(s, ",")
		// }),
	}
	// cmp.FilterValues()
	// cmp.FilterPath()
	// cmp.Equal()
	// cmp.Exporter()
	// cmp.Ignore()
	// cmp.Comparer()
	if diff := cmp.Diff(v1, v2, opts...); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}
