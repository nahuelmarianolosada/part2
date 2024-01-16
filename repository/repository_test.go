package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNumberCollection(t *testing.T) {
	t.Run("Insert one value", func(t *testing.T) {
		nm := &NumberCollection{}

		// Insert a value
		value := 42
		err := nm.Insert(value)
		if err != nil {
			t.Errorf("Error inserting value: %v", err)
		}

		// Check if the value is inserted
		if len(nm.db) != 1 || nm.db[0] != "42" {
			t.Errorf("Insertion failed. Expected: [42], Got: %v", nm.db)
		}
	})

	t.Run("Insert several values", func(t *testing.T) {
		nm := &NumberCollection{}

		// Insert values
		value := 42
		err := nm.Insert(value)
		assert.NoError(t, err)


		value = 12
		err = nm.Insert(value)
		assert.NoError(t, err)


		value = 15
		err = nm.Insert(value)
		assert.NoError(t, err)

		
		// Check if the value is inserted
		assert.Len(t, nm.db, 3)
		assert.Equal(t, "42", nm.db[0])
	})

	t.Run("GetAll", func(t *testing.T) {
		nm := &NumberCollection{}

		// Insert some values
		nm.Insert(1)
		nm.Insert(2)

		// Get all values
		result, err := nm.GetAll()
		if err != nil {
			t.Errorf("Error getting all values: %v", err)
		}

		// Check if all values are returned
		expected := []string{"1", "2"}
		if len(result) != len(expected) || result[0] != expected[0] || result[1] != expected[1] {
			t.Errorf("GetAll failed. Expected: %v, Got: %v", expected, result)
		}
	})

	t.Run("GetByID", func(t *testing.T) {
		nm := &NumberCollection{}

		// Insert some values
		nm.Insert(1)
		nm.Insert(2)

		// Get value by ID
		result, err := nm.GetByID(1)
		if err != nil {
			t.Errorf("Error getting value by ID: %v", err)
		}

		// Check if the correct value is returned
		expected := "2"
		if result != expected {
			t.Errorf("GetByID failed. Expected: %v, Got: %v", expected, result)
		}
	})
}
