package delete

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/kubectl/pkg/cmd/delete"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func Delete(directory string) error {
	factory := cmdutil.NewFactory(genericclioptions.NewConfigFlags(true))

	cmd := &cobra.Command{}
	cmdutil.AddDryRunFlag(cmd)
	flags := delete.NewDeleteFlags("File containing items to delete")
	flags.AddFlags(cmd)
	cmd.Flags().Set("kustomize", directory)
	cmd.Flags().Set("applyset", filepath.Base(filepath.Dir(directory)))
	cmd.Flags().Set("dry-run", "false")

	o, err := flags.ToOptions(nil, genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err != nil {
		return err
	}

	err = o.Complete(factory, []string{}, cmd)
	if err != nil {
		return err
	}

	err = o.Validate()
	if err != nil {
		return err
	}

	return o.RunDelete(factory)
}
