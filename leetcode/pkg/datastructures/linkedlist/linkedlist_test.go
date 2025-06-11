package linkedlist

import (
	"reflect"
	"testing"
)

func TestNewListNode(t *testing.T) {
	tests := []struct {
		name string
		val  int
		want *ListNode
	}{
		{
			name: "Create node with value 5",
			val:  5,
			want: &ListNode{Val: 5},
		},
		{
			name: "Create node with value 0",
			val:  0,
			want: &ListNode{Val: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateLinkedListFromArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *ListNode
	}{
		{
			name: "Empty array",
			nums: []int{},
			want: nil,
		},
		{
			name: "Single element array",
			nums: []int{1},
			want: &ListNode{Val: 1},
		},
		{
			name: "Multiple elements array",
			nums: []int{1, 2, 3},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLinkedListFromArray(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLinkedListFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_ToArray(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name:   "Nil list",
			fields: fields{},
			want:   nil, // An actual nil list (not just an empty node) will be tested differently or by CreateLinkedListFromArray
		},
		{
			name:   "Single node list",
			fields: fields{Val: 1},
			want:   []int{1},
		},
		{
			name: "Multiple nodes list",
			fields: fields{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			want:   []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *ListNode
			// Construct the list based on tt.fields. This is a bit simplified.
			// For the nil case, head remains nil.
			if tt.name != "Nil list" { // A bit of a hack for this test structure
				head = &ListNode{
					Val:  tt.fields.Val,
					Next: tt.fields.Next,
				}
			}
			
			if got := head.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNode.ToArray() = %v, want %v", got, tt.want)
			}
		})
	}

	// Test case for an actual nil list passed to ToArray
	t.Run("Actual nil list to ToArray", func(t *testing.T) {
		var head *ListNode = nil
		want := []int(nil) // Or []int{} depending on desired behavior for nil list to array
		if got := head.ToArray(); !reflect.DeepEqual(got, want) {
			t.Errorf("ListNode.ToArray() for nil list = %v, want %v", got, want)
		}
	})
}