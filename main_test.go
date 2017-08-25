package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadingVariablesFromEnv(t *testing.T) {
	os.Setenv("FOO_BAR_BAZ", "x123a")
	os.Setenv("FOO_QUUX", "z014a")
	os.Setenv("BAR_FOO_BAZ", "d34dm34t")
	variables := readVariablesFromEnv("FOO")
	assert.Equal(t, 2, len(variables), "wrong number of Environment variables converted")
	assert.Equal(t, "x123a", variables["FOO_BAR_BAZ"])
	assert.Equal(t, "z014a", variables["FOO_QUUX"])
}

func TestApplyTemplate(t *testing.T) {
	variables := map[string]string{
		"FOO_BAR_BAZ": "x123a",
		"FOO_QUUX":    "z014a",
	}
	tmpl := "first={{.FOO_BAR_BAZ}} second={{.FOO_QUUX}}"
	result := applyTemplate(tmpl, variables)
	assert.Equal(t, "first=x123a second=z014a", result)
}

func TestApplyDefaultsInTemplate(t *testing.T) {
	variables := map[string]string{"FOO_XXX": "pqrs", "FOO_QUUX": ""}
	tmpl := `first={{ .FOO_BAR_BAZ | default "x123a" }} second={{ .FOO_QUUX | default "z014a" }} third={{ .FOO_XXX | default "abcd" }}`
	result := applyTemplate(tmpl, variables)
	assert.Equal(t, "first=x123a second=z014a third=pqrs", result)
}

func TestFailBadArgsInTemplate(t *testing.T) {
	variables := map[string]string{}
	tmpl := `{{ false | default "z014a" }}`

	assert.Panics(t, func() { _ = applyTemplate(tmpl, variables) }, "should panic")
}

func TestFailExecutionOfTemplate(t *testing.T) {
	variables := map[string]string{}
	tmpl := `{{ .XXX | default }}`
	assert.Panics(t, func() { _ = applyTemplate(tmpl, variables) }, "should panic")
}
