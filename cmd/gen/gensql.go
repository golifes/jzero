package gen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/core/color"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
)

type JzeroSql struct {
	Wd    string
	Style string
}

func (js *JzeroSql) Gen() error {
	sqlDir := filepath.Join(js.Wd, "app", "desc", "sql")
	if f, err := os.Stat(sqlDir); err == nil && f.IsDir() {
		fs, err := os.ReadDir(sqlDir)
		if err != nil {
			return err
		}
		fmt.Printf("%s to generate model code.\n", color.WithColor("Start", color.FgGreen))
		for _, f := range fs {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".sql") {
				sqlFilePath := filepath.Join(sqlDir, f.Name())
				fmt.Printf("%s sql file %s\n", color.WithColor("Using", color.FgGreen), sqlFilePath)
				command := fmt.Sprintf("goctl model mysql ddl --src app/desc/sql/%s --dir ./app/internal/model/%s --home %s --style %s", f.Name(), f.Name()[0:len(f.Name())-len(path.Ext(f.Name()))], filepath.Join(js.Wd, ".template", "go-zero"), js.Style)
				_, err = execx.Run(command, js.Wd)
				if err != nil {
					return err
				}
			}
		}
		fmt.Println(color.WithColor("Done", color.FgGreen))
	}
	return nil
}
