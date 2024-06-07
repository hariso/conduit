package cli

import (
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func cmdInit() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a conduit pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			copyFile("./examples/pipelines/generator-to-log.yaml", "./pipelines/generator-to-log.yaml")
		},
	}
}
