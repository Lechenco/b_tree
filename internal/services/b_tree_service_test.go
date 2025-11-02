package services_test

import (
	"testing"

	"github.com/Lechenco/b_tree/internal/models"
	"github.com/Lechenco/b_tree/internal/services"
)

func TestServiceImplementsDataStore(t *testing.T) {
	var _ models.DataStore[int] = (*services.BTreeService[int])(nil)
}
