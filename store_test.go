package wasmtime

import "testing"

func TestStore(t *testing.T) {
	engine := NewEngine()
	NewStore(engine)
}
