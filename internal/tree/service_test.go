package tree

import (
	"binary-tree-max-path-sum/internal/api"
	"reflect"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	type args struct {
		t api.Tree
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"WithValidItemsShouldSuccess_1", args{testTree1}, &Node{1, &Node{-10, nil, nil}, &Node{-5, nil, nil}}},
		{"WithValidItemsShouldSuccess_2", args{testTree2}, &Node{1, &Node{2, nil, nil}, &Node{3, nil, nil}}},
		{"WithNilItemsShouldGiveNil", args{testTree3}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateBinaryTree(tt.args.t)
			if got == nil && tt.want == nil {
			} else if !reflect.DeepEqual(got.Value, tt.want.Value) && !reflect.DeepEqual(got.Left.Value, tt.want.Left.Value) && !reflect.DeepEqual(got.Right.Value, tt.want.Right.Value) {
				t.Errorf("CreateBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"WithValidItemsShouldSuccess_1", args{1, 3}, 3},
		{"WithValidItemsShouldSuccess_1", args{5, 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testTree1 = api.Tree{
		testNodes1,
		"1",
	}
	testNodes1  = []*api.Node{&testNode1_1, &testNode1_2, &testNode1_3, &testNode1_4, &testNode1_5, &testNode1_6, &testNode1_7}
	testNode1_1 = api.Node{ID: &testID1_1, Left: "2", Right: "3", Value: &testValue1_1}
	testNode1_2 = api.Node{ID: &testID1_2, Left: "6", Right: "7", Value: &testValue1_2}
	testNode1_3 = api.Node{ID: &testID1_3, Left: "null", Right: "null", Value: &testValue1_3}
	testNode1_4 = api.Node{ID: &testID1_4, Left: "null", Right: "null", Value: &testValue1_4}
	testNode1_5 = api.Node{ID: &testID1_5, Left: "4", Right: "5", Value: &testValue1_5}
	testNode1_6 = api.Node{ID: &testID1_6, Left: "null", Right: "null", Value: &testValue1_6}
	testNode1_7 = api.Node{ID: &testID1_7, Left: "null", Right: "null", Value: &testValue1_7}

	testID1_1 = "1"
	testID1_2 = "3"
	testID1_3 = "7"
	testID1_4 = "6"
	testID1_5 = "2"
	testID1_6 = "5"
	testID1_7 = "4"

	testValue1_1 = int64(1)
	testValue1_2 = int64(3)
	testValue1_3 = int64(7)
	testValue1_4 = int64(6)
	testValue1_5 = int64(2)
	testValue1_6 = int64(5)
	testValue1_7 = int64(4)
)

var (
	testTree2 = api.Tree{
		testNodes2,
		"1",
	}
	testNodes2  = []*api.Node{&testNode2_1, &testNode2_2, &testNode2_3}
	testNode2_1 = api.Node{ID: &testID2_1, Left: "2", Right: "3", Value: &testValue2_1}
	testNode2_2 = api.Node{ID: &testID2_2, Left: "null", Right: "null", Value: &testValue2_2}
	testNode2_3 = api.Node{ID: &testID2_3, Left: "null", Right: "null", Value: &testValue2_3}

	testID2_1 = "1"
	testID2_2 = "3"
	testID2_3 = "2"

	testValue2_1 = int64(1)
	testValue2_2 = int64(3)
	testValue2_3 = int64(2)
)

var (
	testTree3 = api.Tree{
		testNodes3,
		"",
	}
	testNodes3 = []*api.Node{}
)

func TestMaxSum(t *testing.T) {
	type args struct {
		n *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"WithValidRootShouldSuccess_1", args{&root_1}, 18},
		{"WithValidRootShouldSuccess_1", args{&root_2}, 6},
		{"WithNilRootShouldReturn0", args{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSum(tt.args.n); got != tt.want {
				t.Errorf("MaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	root_1            = Node{1, &root_1_Left, &root_1_Right}
	root_1_Right      = Node{3, &root_1_RightLeft, &root_1_RightRight}
	root_1_RightRight = Node{7, nil, nil}
	root_1_RightLeft  = Node{6, nil, nil}
	root_1_Left       = Node{2, &root_1_LeftLeft, &root_1_LeftRight}
	root_1_LeftRight  = Node{5, nil, nil}
	root_1_LeftLeft   = Node{4, nil, nil}
)

var (
	root_2       = Node{1, &root_2_Left, &root_2_Right}
	root_2_Right = Node{3, nil, nil}
	root_2_Left  = Node{2, nil, nil}
)
