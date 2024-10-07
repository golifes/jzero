package desc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/parser"
	rpcparser "github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type (
	ImportLines []string

	RegisterLines []string
)

func (l ImportLines) String() string {
	return "\n\n\t" + strings.Join(l, "\n\t")
}

func (l RegisterLines) String() string {
	return "\n\t\t" + strings.Join(l, "\n\t\t")
}

func getProtoDir(protoDirPath string) ([]os.DirEntry, error) {
	protoDir, err := os.ReadDir(protoDirPath)
	if err != nil {
		return nil, nil
	}
	return protoDir, nil
}

func GetProtoFilepath(protoDirPath string) ([]string, error) {
	var protoFilenames []string

	protoDir, err := getProtoDir(protoDirPath)
	if err != nil {
		return nil, err
	}

	for _, protoFile := range protoDir {
		if protoFile.IsDir() {
			filenames, err := GetProtoFilepath(filepath.Join(protoDirPath, protoFile.Name()))
			if err != nil {
				return nil, err
			}
			protoFilenames = append(protoFilenames, filenames...)
		} else {
			if strings.HasSuffix(protoFile.Name(), ".proto") {
				if b, err := protoHasService(filepath.Join(protoDirPath, protoFile.Name())); err == nil && b {
					protoFilenames = append(protoFilenames, filepath.Join(protoDirPath, protoFile.Name()))
				} else if err != nil {
					return nil, err
				}
			}
		}
	}
	return protoFilenames, nil
}

func protoHasService(fp string) (bool, error) {
	r := rpcparser.DefaultProtoParser{}

	parse, err := r.Parse(fp, true)
	if err != nil {
		if strings.Contains(err.Error(), "rpc service not found") {
			return false, nil
		}
		return false, err
	}
	return len(parse.Service) > 0, nil
}

func GetMainApiFilePath(apiDirName string) (string, bool, error) {
	apiDir, err := os.ReadDir(apiDirName)
	if err != nil {
		return "", false, err
	}

	var mainApiFilePath string
	var isDelete bool

	for _, file := range apiDir {
		if file.Name() == "main.api" {
			mainApiFilePath = filepath.Join(apiDirName, file.Name())
			isDelete = false
			break
		}
	}

	if mainApiFilePath == "" {
		apiFilePath, err := getApiFileRelPath(apiDirName)
		if err != nil {
			return "", false, err
		}
		sb := strings.Builder{}
		sb.WriteString("syntax = \"v1\"")
		sb.WriteString("\n")

		for _, api := range apiFilePath {
			sb.WriteString(fmt.Sprintf("import \"%s\"\n", api))
		}

		f, err := os.CreateTemp(apiDirName, "*.api")
		if err != nil {
			return "", false, err
		}

		_, err = f.WriteString(sb.String())
		if err != nil {
			return f.Name(), true, err
		}
		mainApiFilePath = f.Name()
		isDelete = true
		f.Close()
	}
	return mainApiFilePath, isDelete, nil
}

func GetApiServiceName(apiDirName string) string {
	fs, err := getApiFileRelPath(apiDirName)
	if err != nil {
		return ""
	}
	for _, file := range fs {
		apiSpec, err := parser.Parse(filepath.Join(apiDirName, file), "")
		if err != nil {
			return ""
		}
		if apiSpec.Service.Name != "" {
			return apiSpec.Service.Name
		}
	}
	return ""
}

func GetRpcMethodUrl(method *descriptorpb.MethodDescriptorProto) string {
	ext := proto.GetExtension(method.GetOptions(), annotations.E_Http)
	switch rule := ext.(type) {
	case *annotations.HttpRule:
		switch httpRule := rule.GetPattern().(type) {
		case *annotations.HttpRule_Get:
			return "GET:" + httpRule.Get
		case *annotations.HttpRule_Post:
			return "POST:" + httpRule.Post
		case *annotations.HttpRule_Put:
			return "PUT:" + httpRule.Put
		case *annotations.HttpRule_Delete:
			return "DELETE:" + httpRule.Delete
		case *annotations.HttpRule_Patch:
			return "PATCH:" + httpRule.Patch
		}
	}
	return ""
}

func getApiFileRelPath(apiDirName string) ([]string, error) {
	var apiFiles []string

	allApiFiles, err := FindApiFiles(apiDirName)
	if err != nil {
		return nil, err
	}
	for _, file := range allApiFiles {
		rel, err := filepath.Rel(apiDirName, file)
		if err != nil {
			return nil, err
		}
		apiFiles = append(apiFiles, filepath.ToSlash(rel))
	}

	return apiFiles, nil
}

func FindApiFiles(dir string) ([]string, error) {
	var apiFiles []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subFiles, err := FindApiFiles(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			apiFiles = append(apiFiles, subFiles...)
		} else if filepath.Ext(file.Name()) == ".api" {
			apiFiles = append(apiFiles, filepath.Join(dir, file.Name()))
		}
	}

	return apiFiles, nil
}

func FindRouteApiFiles(dir string) ([]string, error) {
	var apiFiles []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subFiles, err := FindRouteApiFiles(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			apiFiles = append(apiFiles, subFiles...)
		} else if filepath.Ext(file.Name()) == ".api" {
			parse, err := parser.Parse(filepath.Join(dir, file.Name()), "")
			if err != nil {
				return nil, errors.Wrapf(err, "parse api: %s meet error", filepath.Join(dir, file.Name()))
			}
			if len(parse.Service.Routes()) > 0 {
				apiFiles = append(apiFiles, filepath.Join(dir, file.Name()))
			}
		}
	}

	return apiFiles, nil
}

// GetApiFrameMainGoFilename goctl/api/gogen/genmain.go
func GetApiFrameMainGoFilename(wd, style string) string {
	serviceName := GetApiServiceName(filepath.Join(wd, "desc", "api"))
	serviceName = strings.ToLower(serviceName)
	filename, err := format.FileNamingFormat(style, serviceName)
	if err != nil {
		return ""
	}

	if strings.HasSuffix(filename, "-api") {
		filename = strings.ReplaceAll(filename, "-api", "")
	}
	return filename + ".go"
}

// GetApiFrameEtcFilename goctl/api/gogen/genetc.go
func GetApiFrameEtcFilename(wd, style string) string {
	serviceName := GetApiServiceName(filepath.Join(wd, "desc", "api"))
	filename, err := format.FileNamingFormat(style, serviceName)
	if err != nil {
		return ""
	}
	return filename + ".yaml"
}

// GetProtoFrameMainGoFilename goctl/rpc/generator/genmain.go
func GetProtoFrameMainGoFilename(source, style string) string {
	filename, err := format.FileNamingFormat(style, source)
	if err != nil {
		return ""
	}
	return filename + ".go"
}

// GetProtoFrameEtcFilename goctl/rpc/generator/genetc.go
func GetProtoFrameEtcFilename(source, style string) string {
	filename, err := format.FileNamingFormat(style, source)
	if err != nil {
		return ""
	}
	return filename + ".yaml"
}
