package mockos_test

import (
	"testing"

	"github.com/anuvu/disko/mockos"
	"github.com/stretchr/testify/assert"
)

func TestSystem(t *testing.T) {
	sys := mockos.System("model_sys.json")
	assert := assert.New(t)

	assert.NotNil(sys)
}
