package subnet

import (
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextAvailableSubnet(t *testing.T) {
	assert := assert.New(t)
	p, err := netip.ParsePrefix("fd18:fad4:bce5:4400::/56")
	assert.NoError(err)
	calc := NewCalculator()
	err = calc.AddPool(p)
	assert.NoError(err)
	next, err := calc.NextAvailableSubnet(64)
	if assert.NoError(err) {
		assert.Equal("fd18:fad4:bce5:4400::/64", next.String())
	}
}
