package spdy

import (
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckNodes(t *testing.T) {

	options := k8s.NewKubectlOptions("", "", "default")
	nodes := k8s.GetNodes(t, options)
	require.Equal(t, len(nodes), 1)
}
