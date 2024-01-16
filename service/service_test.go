package service

import (
	"part2/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberServiceImpl(t *testing.T) {
	nm := &repository.NumberCollection{}
	ns := &NumberServiceImpl{
		NumberCollection: nm,
	}

	t.Run("Insert", func(t *testing.T) {
		// Insert a value
		id := 42
		err := ns.Insert(id)
		assert.NoError(t, err)

		result := ns.GetAll()

		// Check if the value is inserted
		assert.Len(t, result, 1)
		assert.Equal(t, result[0], "Type 1")
	})

	t.Run("GetAll", func(t *testing.T) {
		nm := &repository.NumberCollection{}
		ns := &NumberServiceImpl{
			NumberCollection: nm,
		}

		// Insert some values
		nm.Insert(1)
		nm.Insert(2)

		// Get all values
		result := ns.GetAll()

		// Check if all values are returned
		expected := []string{"1", "2"}
		assert.Equal(t, expected, result)
	})
	
	t.Run("GetByID", func(t *testing.T) {

		nm := &repository.NumberCollection{}
		ns := &NumberServiceImpl{
			NumberCollection: nm,
		}
		
		// Insert some values
		nm.Insert(1)
		nm.Insert(2)
		
		// Get value by ID
		result := ns.GetByID(1)

		// Check if the correct value is returned
		expected := "1"
		
		assert.Equal(t, expected, *result)
	})
}
