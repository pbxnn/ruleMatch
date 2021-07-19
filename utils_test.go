package rulematch

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestInStringSlice(t *testing.T) {

    collection := []string{"1", "2", "3"}

    r1 := InStringSlice(collection, "1")
    r2 := InStringSlice(collection, "0")

    t.Log(r1, r2)
    assert.True(t, r1)
    assert.False(t, r2)
}

func TestVersionCompare(t *testing.T) {
    v1 := "1.2.0"
    v2 := "1.2.1"
    v3 := "1.2.1.1"

    r11 := VersionCompare(v1, v1)
    r12 := VersionCompare(v1, v2)
    r13 := VersionCompare(v1, v3)
    r23 := VersionCompare(v2, v3)

    t.Log(r11, r12, r13, r23)

    assert.Equal(t, r11, 0)
    assert.Less(t, r12, 0)
    assert.Less(t, r13, 0)
    assert.Less(t, r23, 0)
}
