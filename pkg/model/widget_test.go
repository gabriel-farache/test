package model

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetRandomFromList(t *testing.T) {
	colorList := getColor()
	for i := 0; i < 50000; i++ {
		assert.Assert(t, getRandomFromList(colorList) != "")
	}
}

func TestNewRandomWidget(t *testing.T) {
	widget := NewRandomWidget()
	assert.Assert(t, widget.Creator == "Admin")
	assert.Assert(t, widget.Name != "")
	assert.Assert(t, widget.Count > 0 && widget.Count < 1000)
}

func TestGetSeedData(t *testing.T) {
	widgets := GetSeedData(1000)
	assert.Assert(t, len(widgets) == 1000)
	assert.Assert(t, widgets[4].Creator == "Admin")
	assert.Assert(t, widgets[4].Name != "")
	assert.Assert(t, widgets[4].Count > 0 && widgets[4].Count < 1000)
}
