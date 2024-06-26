package new

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jzero-io/jzero/app/pkg/templatex"
	"github.com/jzero-io/jzero/embeded"
	"github.com/rinchsan/gosimports"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var (
	Module string
	Dir    string
	APP    string
	// ConfigType config type
	ConfigType string
	// Remote templates repo
	Remote string
	Branch string

	Version string
)

func NewProject(_ *cobra.Command, _ []string) error {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	if embeded.Home == "" {
		embeded.Home = filepath.Join(homeDir, ".jzero", Version)
	}

	// mkdir output
	err = os.MkdirAll(Dir, 0o755)
	cobra.CheckErr(err)
	// go mod init
	_, err = execx.Run(fmt.Sprintf("go mod init %s", Module), Dir)
	cobra.CheckErr(err)

	templateData := map[string]interface{}{
		"Module":     Module,
		"APP":        APP,
		"ConfigType": ConfigType,
	}

	// touch main.go
	mainFile, err := templatex.ParseTemplate(templateData, embeded.ReadTemplateFile(filepath.Join("jzero", "main.go.tpl")))
	cobra.CheckErr(err)
	err = checkWrite(filepath.Join(Dir, "main.go"), mainFile)
	cobra.CheckErr(err)

	// write cmd dir
	jzeroCmd := JzeroCmd{TemplateData: templateData}
	err = jzeroCmd.New()
	cobra.CheckErr(err)

	// write app/*.go
	jzeroApp := JzeroApp{TemplateData: templateData}
	err = jzeroApp.New()
	cobra.CheckErr(err)

	// write proto dir
	jzeroProto := JzeroProto{TemplateData: templateData}
	err = jzeroProto.New()
	cobra.CheckErr(err)

	// write app/desc/api
	jzeroApi := JzeroApi{TemplateData: templateData}
	err = jzeroApi.New()
	cobra.CheckErr(err)

	// write config.yaml
	jzeroConfig := JzeroConfig{TemplateData: templateData}
	err = jzeroConfig.New()
	cobra.CheckErr(err)

	// write Dockerfile
	dockerFile, err := templatex.ParseTemplate(templateData, embeded.ReadTemplateFile(filepath.Join("jzero", "Dockerfile.tpl")))
	cobra.CheckErr(err)
	err = checkWrite(filepath.Join(Dir, "Dockerfile"), dockerFile)
	cobra.CheckErr(err)

	// write Taskfile.yml
	err = checkWrite(filepath.Join(Dir, "Taskfile.yml"), embeded.ReadTemplateFile(filepath.Join("jzero", "Taskfile.yml.tpl")))
	cobra.CheckErr(err)

	err = embeded.WriteTemplateDir(filepath.Join("go-zero"), filepath.Join(embeded.Home, "go-zero"))
	cobra.CheckErr(err)

	// write .gitignore
	gitignoreFile, err := templatex.ParseTemplate(templateData, embeded.ReadTemplateFile(filepath.Join("jzero", "gitignore.tpl")))
	cobra.CheckErr(err)
	err = checkWrite(filepath.Join(Dir, ".gitignore"), gitignoreFile)
	cobra.CheckErr(err)

	return nil
}

func checkWrite(path string, bytes []byte) error {
	var err error
	if len(bytes) == 0 {
		return nil
	}
	if !pathx.FileExists(filepath.Dir(path)) {
		err = os.MkdirAll(filepath.Dir(path), 0o755)
		if err != nil {
			return err
		}
	}

	bytesFormat := bytes
	// if is go file. format it
	if filepath.Ext(path) == ".go" {
		bytesFormat, err = gosimports.Process("", bytes, &gosimports.Options{FormatOnly: true})
		if err != nil {
			return err
		}
	}

	return os.WriteFile(path, bytesFormat, 0o644)
}
