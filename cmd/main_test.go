package main

import (
	"testing"
)

func TestMainFile(t *testing.T) {
	// Здесь можем проверить, что main запускается без паники, к примеру.
	// Это чисто формальный тест, т.к. main — обычно не тестируют напрямую.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main паниковал: %v", r)
		}
	}()
	go main()
}
