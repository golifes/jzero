/*
Copyright © 2024 jaronnie <jaron@jaronnie.com>

*/

package cmd

import (
	"os"
	"path/filepath"

	"github.com/jzero-io/jzero/cmd/gen"
	"github.com/jzero-io/jzero/cmd/genswagger"
	"github.com/jzero-io/jzero/embeded"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "jzero gen code",
	Long:  `jzero gen code`,
	PreRun: func(_ *cobra.Command, _ []string) {
		gen.Version = Version

		// check go-zero api template
		home, _ := os.UserHomeDir()
		if !pathx.FileExists(filepath.Join(home, ".jzero", Version, "go-zero")) {
			err := embeded.WriteTemplateDir(filepath.Join("go-zero"), filepath.Join(home, ".jzero", Version, "go-zero"))
			cobra.CheckErr(err)
		}
	},
	RunE:         gen.Gen,
	SilenceUsage: false,
}

// genSwaggerCmd represents the genSwagger command
var genSwaggerCmd = &cobra.Command{
	Use:   "swagger",
	Short: "jzero gen swagger",
	Long:  `jzero gen swagger`,
	RunE:  genswagger.Gen,
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.AddCommand(genSdkCmd)
	genCmd.AddCommand(genSwaggerCmd)

	genCmd.Flags().StringVarP(&gen.WorkingDir, "working-dir", "w", "", "set working dir")

	dir, _ := os.UserHomeDir()
	genCmd.Flags().StringVarP(&embeded.Home, "home", "", filepath.Join(dir, ".jzero"), "set template home")

	genSwaggerCmd.Flags().StringVarP(&genswagger.Dir, "dir", "d", filepath.Join("app", "desc", "swagger"), "set swagger output dir")
	genCmd.Flags().StringVarP(&gen.Style, "style", "", "gozero", "The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md]")

}
