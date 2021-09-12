package gocmp

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCmp(t *testing.T) {
	type Compare struct {
		Name  string
		Value int
	}

	v1 := &Compare{
		Name:  "Tom",
		Value: 100,
	}
	v2 := &Compare{
		Name:  "Andrew",
		Value: 50,
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
		Exported:   200,
		unexported: 1,
	}
	opt := cmpopts.IgnoreUnexported(Compare{})
	if diff := cmp.Diff(v1, v2, opt); diff != "" {
		t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
	}
}

func TestIgnoreFields(t *testing.T) {
	type Compare struct {
		Exported   int
		unexported int
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}

	v1 := &Compare{
		Exported:   100,
		unexported: 1,
		CreatedAt:  time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:  time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	v2 := &Compare{
		Exported:   100,
		unexported: 1,
		CreatedAt:  time.Date(3100, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:  time.Date(3100, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	opts := []cmp.Option{
		cmpopts.IgnoreUnexported(Compare{}),
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
// func IgnoreFields(typ interface{}, names ...string) cmp.Option
// func IgnoreInterfaces(ifaces interface{}) cmp.Option
// func IgnoreMapEntries(discardFunc interface{}) cmp.Option
// func IgnoreSliceElements(discardFunc interface{}) cmp.Option
// func IgnoreTypes(typs ...interface{}) cmp.Option
// func IgnoreUnexported(typs ...interface{}) cmp.Option
// func SortMaps(lessFunc interface{}) cmp.Option
// func SortSlices(lessFunc interface{}) cmp.Option
