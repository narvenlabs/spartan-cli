package cmd

import (
	"fmt"
	"github.com/narvenlabs/spartan-cli/templates"
	"golang.org/x/mod/modfile"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// resourcesCmd represents the resources command
var resourcesCmd = &cobra.Command{
	Use:   "resource",
	Short: "Generate a resource",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			log.Fatal("Name is required")
		}
		path, _ := cmd.Flags().GetString("path")
		if path == "" {
			currWd, _ := os.Getwd()
			path = currWd
		}

		// -f "username:string, email:string:unique"
		fieldList, _ := cmd.Flags().GetString("fields")
		fieldRow := strings.Split(fieldList, ",")

		generateResource(path, name, fieldRow)
		generateMigration(path, name)
	},
}

func genFieldLine(field string, withJ bool) string {
	parts := strings.Split(field, ":")
	fName := parts[0]
	fType := parts[1]
	fExtra := parts[2]

	line := fmt.Sprintf("%s	%s	%s", fName, fType, fExtra)

	if withJ {
		return fmt.Sprintf("%s	`json:\"%s\"`", line, strings.ToLower(fName))
	}
	return line
}

func getModuleName(path string) (string, error) {
	b, err := ioutil.ReadFile(filepath.Join(path, "go.mod"))
	if err != nil {
		return "", err
	}
	return modfile.ModulePath(b), nil
}

func generateResource(path, name string, fields []string) {
	fmt.Println(path, name)
	type OperationLogic struct {
		Path    *string
		Content *string
	}

	entityPath := filepath.Join(path, "entity")
	usecasePath := filepath.Join(path, "usecase")
	repositoryPath := filepath.Join(path, "infrastructure", "repository")
	handlerPath := filepath.Join(path, "api/handler")
	presenterPath := filepath.Join(path, "api/presenter")
	dbDriver := "mysql"
	moduleName, err := getModuleName(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	operations := []OperationLogic{
		{
			Path:    StrPtr(filepath.Join(entityPath, "errors.go")),
			Content: StrPtr(templates.GenEntityErrors()),
		},
		{
			Path:    StrPtr(filepath.Join(entityPath, fmt.Sprintf("%s.go", strings.ToLower(name)))),
			Content: StrPtr(templates.GenCustomEntity(name)),
		},
		{
			Path:    StrPtr(filepath.Join(entityPath, fmt.Sprintf("%s_test.go", strings.ToLower(name)))),
			Content: StrPtr(templates.GenTestCustomEntity(moduleName, name)),
		},
		{
			Path:    StrPtr(filepath.Join(usecasePath, strings.ToLower(name), "interface.go")),
			Content: StrPtr(templates.GenEntityUsecaseInterface(moduleName, name)),
		},
		{
			Path:    StrPtr(filepath.Join(usecasePath, strings.ToLower(name), "service.go")),
			Content: StrPtr(templates.GenEntityUsecaseService(moduleName, name)),
		},
		{
			Path:    StrPtr(filepath.Join(repositoryPath, fmt.Sprintf("%s_%s%s", strings.ToLower(name), dbDriver, ".go"))),
			Content: StrPtr(templates.GenEntityRepository(moduleName, name)),
		},
		{
			Path:    StrPtr(filepath.Join(handlerPath, fmt.Sprintf("%s%s", strings.ToLower(name), ".go"))),
			Content: StrPtr(templates.GenResourceHandler(moduleName, name)),
		},
		{
			Path:    StrPtr(filepath.Join(handlerPath, fmt.Sprintf("%s%s", strings.ToLower(name), ".go"))),
			Content: StrPtr(templates.GenResourceHandler(moduleName, name)),
		},
		{
			Path:    StrPtr(filepath.Join(presenterPath, fmt.Sprintf("%s%s", strings.ToLower(name), ".go"))),
			Content: StrPtr(templates.GenResourcePresenter(moduleName, name)),
		},
	}

	for _, o := range operations {
		if o.Path == nil {
			fmt.Println(*o.Content)
			continue
		}

		if o.Content != nil { // its a file
			f, err := os.Create(*o.Path)
			if err != nil {
				log.Fatal(err)
			}
			_, err = f.WriteString(*o.Content)
			if err != nil {
				log.Fatal(err)
			}

			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		} else { // its a folder
			if err := os.MkdirAll(*o.Path, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}
	}

	args := []string{"mod", "tidy"}
	cmd := exec.Command("go", args...)
	cmd.Dir = path
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}

func generateMigration(path, resourceName string) {
	p := filepath.Join(
		path,
		"migrations",
		fmt.Sprintf("%d_add_%s_table.sql", time.Now().UnixNano(), strings.ToLower(resourceName)),
	)
	content := StrPtr(templates.GenMigration(resourceName))

	f, err := os.Create(p)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(*content)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	generateCmd.AddCommand(resourcesCmd)

	resourcesCmd.Flags().StringP("name", "n", "", "Resource name")
	resourcesCmd.Flags().StringP("path", "p", "", "Project path")
	resourcesCmd.Flags().StringP("fields", "f", "", "Fields")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resourcesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resourcesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
