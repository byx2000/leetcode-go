package satisfiability_of_equality_equations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.False(t, equationsPossible([]string{"a==b", "b!=a"}))
	assert.True(t, equationsPossible([]string{"b==a", "a==b"}))
	assert.True(t, equationsPossible([]string{"a==b", "b==c", "a==c"}))
	assert.False(t, equationsPossible([]string{"a==b", "b!=c", "c==a"}))
	assert.True(t, equationsPossible([]string{"c==c", "b==d", "x!=z"}))
	assert.True(t, equationsPossible([]string{"a!=b", "b!=c", "c!=a"}))
}
