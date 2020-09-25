package skr

import (
	"fmt"

	"github.com/spf13/cobra"
)

// UpgradeKymaCommand represents an execution of the skr upgrade kyma command. Inherits fields and methods of UpgradeCommand
type UpgradeKymaCommand struct {
	UpgradeCommand
}

// NewUpgradeKymaCmd constructs a new instance of UpgradeKymaCommand and configures it in terms of a cobra.Command
func NewUpgradeKymaCmd() *cobra.Command {
	cmd := UpgradeKymaCommand{}
	cobraCmd := &cobra.Command{
		Use:   "kyma --target <TARGET SPEC> ... [--target-exclude <TARGET SPEC> ...]",
		Short: "Upgrade or reconfigure Kyma on one or more Kyma runtimes",
		Long: `Upgrades or reconfigures Kyma on targets of SKRs.
The upgrade is performed by the Kyma Control plane within a new orchestration asynchronously, the id of which is returned by the command upon success.
The targets of SKRs are specified via the --target and --target-exclude options. At lease one --target must be specified.
The Kyma version and configurations to use for the upgrade is taken from the Kyma Control Plane during processing of the orchestration.`,
		PreRunE: func(_ *cobra.Command, _ []string) error { return cmd.Validate() },
		Example: `  skr upgrade kyma --target all --schedule maintenancewindow     Upgrade Kyma on all runtimes in their next respective maintenance window hours
  skr upgrade kyma --target "account=CA.*"                       Upgrade Kyma on runtimes of all Global Accounts starting with CA
  skr upgrade kyma --target all --target-exclude "account=CA.*"  Upgrade Kyma on runtimes of all Global Accounts not starting with CA
  skr upgrade kyma --target "region=europe|eu|uk"                Upgrade Kyma on runtimes whose region belongs to Europe`,
		RunE: func(_ *cobra.Command, _ []string) error { return cmd.Run() },
	}

	cmd.SetUpgradeOpts(cobraCmd)
	return cobraCmd
}

// Run executes the upgrade kyma command
func (cmd *UpgradeKymaCommand) Run() error {
	fmt.Println("Not implemented yet.")
	return nil
}

// Validate checks the input parameters of the upgrade kyma command
func (cmd *UpgradeKymaCommand) Validate() error {
	err := cmd.ValidateTransformUpgradeOpts()
	if err != nil {
		return err
	}
	return nil
}
