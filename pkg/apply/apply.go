package apply

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/kubectl/pkg/cmd/apply"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func Apply(directory string) error {
	f := cmdutil.NewFactory(genericclioptions.NewConfigFlags(true))

	cmd := &cobra.Command{}
	flags := apply.NewApplyFlags(genericiooptions.NewTestIOStreamsDiscard())
	flags.AddFlags(cmd)
	cmd.Flags().Set("kustomize", directory)
	cmd.Flags().Set("applyset", filepath.Base(filepath.Dir(directory)))

	o, err := flags.ToOptions(f, cmd, "kubectl", []string{})
	if err != nil {
		return err
	}

	err = o.Validate()
	if err != nil {
		return err
	}

	return o.Run()
}
