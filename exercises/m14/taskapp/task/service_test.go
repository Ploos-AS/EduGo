package task

import "testing"

type fakeStore struct {
	created string
}

func (f *fakeStore) Create(text string) Task {
	f.created = text
	return Task{ID: 7, Text: text}
}

func TestCreate(t *testing.T) {
	store := &fakeStore{}
	service := NewService(store)

	got, err := service.Create("  learn architecture  ")
	if err != nil {
		t.Fatal(err)
	}
	if store.created != "learn architecture" {
		t.Fatalf("stored %q", store.created)
	}
	if got.ID != 7 || got.Text != "learn architecture" {
		t.Fatalf("got %#v", got)
	}
}

func TestCreateRejectsBlank(t *testing.T) {
	service := NewService(&fakeStore{})
	if _, err := service.Create("   "); err == nil {
		t.Fatal("expected error")
	}
}
