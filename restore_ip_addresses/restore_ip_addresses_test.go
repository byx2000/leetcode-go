package restore_ip_addresses

import (
	"github.com/stretchr/testify/assert"
	. "leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, SliceToSet([]string{"255.255.11.135", "255.255.111.35"}), SliceToSet(restoreIpAddresses("25525511135")))
	assert.Equal(t, SliceToSet([]string{"0.0.0.0"}), SliceToSet(restoreIpAddresses("0000")))
	assert.Equal(t, SliceToSet([]string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}), SliceToSet(restoreIpAddresses("101023")))
}
