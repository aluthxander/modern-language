package test

import (
	"belajar-12-resful-api/simple"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simplaService, err := simple.InitializeService(true)
	fmt.Println(err)
	fmt.Println(simplaService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simplaService, err := simple.InitializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simplaService)
}