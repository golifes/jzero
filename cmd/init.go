/*
Copyright © 2024 jaronnie <jaron@jaronnie.com>

*/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jaronnie/genius"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/jzero-io/jzero/embeded"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "jzero init",
	Long:  `jzero init`,
	Run: func(_ *cobra.Command, _ []string) {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Check $HOME/.jzero/protosets
		err = os.MkdirAll(filepath.Join(home, ".jzero", ".protosets"), 0o755)
		cobra.CheckErr(err)

		file, err := embeded.Config.ReadFile("config.toml")
		cobra.CheckErr(err)

		// write protosets
		dir, err := embeded.Protosets.ReadDir(".protosets")
		cobra.CheckErr(err)
		for _, d := range dir {
			pb, err := embeded.Protosets.ReadFile(filepath.Join(".protosets", d.Name()))
			cobra.CheckErr(err)
			err = os.WriteFile(filepath.Join(home, ".jzero", ".protosets", d.Name()), pb, 0o644)
			cobra.CheckErr(err)
		}

		g, err := genius.NewFromToml(file)
		cobra.CheckErr(err)

		oldProtoSetsI := g.Get("Gateway.Upstreams.0.ProtoSets")
		oldProtoSets := cast.ToStringSlice(oldProtoSetsI)
		var newProtoSets []string

		for _, protoSet := range oldProtoSets {
			newProtoSets = append(newProtoSets, filepath.Join(home, ".jzero", protoSet))
		}
		err = g.Set("Gateway.Upstreams.0.ProtoSets", newProtoSets)
		cobra.CheckErr(err)

		toml, err := g.EncodeToToml()
		cobra.CheckErr(err)
		err = os.WriteFile(filepath.Join(home, ".jzero", "config.toml"), toml, 0o644)
		cobra.CheckErr(err)

		fmt.Println("init success")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
