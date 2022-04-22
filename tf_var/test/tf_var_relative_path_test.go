package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestPassingTfVarParentPath(t *testing.T) {

	terraformtfvars := "varfile.tfvars"
	//assert.FileExists(t, terraformtfvars, "File does not exist: ", terraformtfvars)

	terraformOption := &terraform.Options{
		TerraformDir: "../",
		VarFiles:     []string{terraformtfvars},
	}

	out, err := terraform.InitAndPlanE(t, terraformOption)
	require.NoError(t, err)
	log.Printf("%s\n", out)
}
