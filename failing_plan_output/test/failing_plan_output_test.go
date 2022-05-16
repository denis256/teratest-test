package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPassingTfVarParentPath(t *testing.T) {
	terraformOption := &terraform.Options{
		TerraformDir: "../",
	}
	_, err := terraform.InitE(t, terraformOption)
	require.NoError(t, err)

	output, err := terraform.RunTerraformCommandE(t, terraformOption, "plan")
	if err != nil {
		fmt.Printf("%s", output)
	}

}
