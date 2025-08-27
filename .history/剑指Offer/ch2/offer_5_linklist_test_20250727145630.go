package ch2

import (
	"reflect"
	"testing"
)

func TestList_Append(t *testing.T) {
	type fields struct {
		head *ListNode
		tail *ListNode
		Len  int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				tail: tt.fields.tail,
				Len:  tt.fields.Len,
			}
			if got := l.Append(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}
