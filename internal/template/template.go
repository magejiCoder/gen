package template

import (
	"embed"
	"fmt"
)

var (
	//go:embed set/*
	setTemplate embed.FS
)

func LoadSetTemplate() ([]byte, error) {
	tpl, err := setTemplate.ReadFile("set/set.go.tpl")
	if err != nil {
		return nil, fmt.Errorf("load template: %w", err)
	}
	return tpl, nil
}

func LoadSetTestTemplate() ([]byte, error) {
	tpl, err := setTemplate.ReadFile("set/set.go.test.tpl")
	if err != nil {
		return nil, fmt.Errorf("load template: %w", err)
	}
	return tpl, nil
}
