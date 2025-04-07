package mongo

import (
	"GoNews/pkg/storage"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		connString string
	}
	tests := []struct {
		name    string
		args    args
		want    *Store
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.connString)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Posts(t *testing.T) {
	tests := []struct {
		name    string
		s       *Store
		want    []storage.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Posts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.Posts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.Posts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_AddPost(t *testing.T) {
	type args struct {
		post storage.Post
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.AddPost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Store.AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_UpdatePost(t *testing.T) {
	type args struct {
		post storage.Post
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UpdatePost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Store.UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_DeletePost(t *testing.T) {
	type args struct {
		post storage.Post
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeletePost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Store.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
