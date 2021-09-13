package gocmp

import (
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

// https://pkg.go.dev/github.com/google/go-cmp@v0.5.4/cmp/cmpopts
// func AcyclicTransformer(name string, xformFunc interface{}) cmp.Option
// func EquateApprox(fraction, margin float64) cmp.Option
// func EquateApproxTime(margin time.Duration) cmp.Option
// func EquateEmpty() cmp.Option
// func EquateErrors() cmp.Option
// func EquateNaNs() cmp.Option
// func IgnoreInterfaces(ifaces interface{}) cmp.Option
// func IgnoreMapEntries(discardFunc interface{}) cmp.Option
// func IgnoreSliceElements(discardFunc interface{}) cmp.Option
// func IgnoreTypes(typs ...interface{}) cmp.Option
// func SortMaps(lessFunc interface{}) cmp.Option
// func SortSlices(lessFunc interface{}) cmp.Option
