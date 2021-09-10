package gocomp

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestGoCompService_GetResponse(t *testing.T) {
	tests := []struct {
		name string
		want *GoCompResponse
	}{
		{
			name: "test 1",
			want: &GoCompResponse{
				Name: "result",
				Data: map[string][]string{
					"data1": {"1", "2", "3"},
					"data2": {"4", "5", "6"},
					"data3": {"7", "8", "9"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GoCompService{}
			got := g.GetResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GoCompService.GetResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

type User struct {
	UserID    string
	UserName  string
	Languages []string
}

func TestTome(t *testing.T) {
	tom1 := User{
		UserID:    "0001",
		UserName:  "Tom",
		Languages: []string{"Java", "Go"},
	}

	tom2 := User{
		UserID:    "0001",
		UserName:  "Tom",
		Languages: []string{"Ruby", "Go"},
	}

	if diff := cmp.Diff(tom1, tom2); diff != "" {
		t.Errorf("User value is mismatch (-tom +tom2):\n%s", diff)
	}
}

func TestX(t *testing.T) {
	type X struct {
		numUnExport int
		NumExport   int
	}

	num1 := X{100, -1}
	num2 := X{999, -1}

	//opt := cmp.AllowUnexported(X{})
	opt := cmpopts.IgnoreUnexported(X{})

	if diff := cmp.Diff(num1, num2, opt); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}

func TestX2(t *testing.T) {
	type X struct {
		NumExport int
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	num1 := X{-1, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}
	num2 := X{-1, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}

	opt := cmpopts.IgnoreFields(X{}, "CreatedAt", "UpdatedAt")
	if diff := cmp.Diff(num1, num2, opt); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}

func TestX3(t *testing.T) {
	type Item struct {
		HashKey   string
		CreatedAt time.Time
	}

	type X struct {
		Items []Item
	}

	x1 := X{[]Item{
		{HashKey: "aaa", CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
		{HashKey: "bbb", CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
	}}
	x2 := X{[]Item{
		{HashKey: "aaa", CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		{HashKey: "bbb", CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
	}}

	opt := cmpopts.IgnoreFields(Item{}, "CreatedAt")

	if diff := cmp.Diff(x1, x2, opt); diff != "" {
		t.Errorf("X value is mismatch (-x1 +x2):%s\n", diff)
	}
}

func TestX4(t *testing.T) {
	type X struct {
		NumberExport int
		numUnExport  int
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	num1 := X{100, -1, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}
	num2 := X{999, -111, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()}

	opts := []cmp.Option{
		cmpopts.IgnoreUnexported(X{}),
		cmpopts.IgnoreFields(X{}, "CreatedAt", "UpdatedAt"),
	}

	if diff := cmp.Diff(num1, num2, opts...); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}

func TestX5(t *testing.T) {
	type X struct {
		Numbers []int
	}

	num1 := X{[]int{1, 2, 3, 4, 5}}
	num2 := X{[]int{5, 4, 3, 2, 1}}

	opt := cmpopts.SortSlices(func(i, j int) bool {
		return i < j
	})

	if diff := cmp.Diff(num1, num2, opt); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}

func TestX6(t *testing.T) {
	type X struct {
		HashKey string
		SortKey string
	}

	x1 := []X{
		{HashKey: "AAA", SortKey: "AAA"},
		{HashKey: "BBB", SortKey: "BBB"},
		{HashKey: "CCC", SortKey: "CCC"},
	}
	x2 := []X{
		{HashKey: "CCC", SortKey: "CCC"},
		{HashKey: "AAA", SortKey: "AAA"},
		{HashKey: "BBB", SortKey: "BBB"},
	}

	opt := cmpopts.SortSlices(func(i, j X) bool {
		// ハッシュキーとソートキーの昇順でソート
		return i.HashKey < j.HashKey ||
			(i.HashKey == j.HashKey && i.SortKey < j.SortKey)
	})

	if diff := cmp.Diff(x1, x2, opt); diff != "" {
		t.Errorf("X value is mismatch (-x1 +x2):%s\n", diff)
	}
}

func TestX7(t *testing.T) {
	// 空スライスとnilスライスの違いを無視する
	type X struct {
		Numbers []int
	}

	num1 := X{[]int{}}
	num2 := X{nil}

	opt := cmpopts.EquateEmpty()

	if diff := cmp.Diff(num1, num2, opt); diff != "" {
		t.Errorf("X value is mismatch (-num1 +num2):%s\n", diff)
	}
}
