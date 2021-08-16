package cmd

import (
	"fmt"
	"github.com/Narven/igniter-cli/templates"
	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		box := packr.NewBox("./templates")

		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			log.Fatalf("name is mandatory")
		}

		path, _ := cmd.Flags().GetString("path")
		if path == "" {
			currWd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			path = currWd
		}

		moduleName, _ := cmd.Flags().GetString("module")
		if moduleName == "" {
			moduleName = name
		}

		dbDriver, _ := cmd.Flags().GetString("driver")

		generateBoilerplate(path, name, moduleName, dbDriver, box)
	},
}

func generateBoilerplate(path, projectName, moduleName, dbDriver string, box packr.Box) {
	p := filepath.Join(path, projectName)

	// curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	type OperationLogic struct {
		Path    *string
		Content *string
	}

	operations := []OperationLogic{
		{
			Path:    StrPtr("README.md"),
			Content: StrPtr(templates.GenREADME(projectName)),
		},
		{
			Path:    StrPtr(".gitignore"),
			Content: StrPtr(templates.GenGitignore()),
		},
		{
			Path:    StrPtr(".gitattributes"),
			Content: StrPtr(templates.GenGitAttributes()),
		},
		{
			Path:    StrPtr("go.mod"),
			Content: StrPtr(templates.GenGoMod(moduleName)),
		},
		{
			Path:    StrPtr(".env"),
			Content: StrPtr(templates.GenEnv(projectName, dbDriver)),
		},
		{
			Path:    StrPtr("Makefile"),
			Content: StrPtr(templates.GenMakeFile(projectName)),
		},
		{
			Path:    StrPtr(".editorconfig"),
			Content: StrPtr(templates.GenEditorConfig()),
		},
		{
			Path:    StrPtr(".dockerignore"),
			Content: StrPtr(templates.GenDockerIgnore()),
		},
		{
			Path:    StrPtr("docker-compose.yml"),
			Content: StrPtr(templates.GenDockerCompose(projectName, dbDriver)),
		},
		{
			Path:    StrPtr("Dockerfile"),
			Content: StrPtr(templates.GenDockerfile()),
		},
		{
			Path:    StrPtr("LICENSE"),
			Content: StrPtr(templates.GenMITLicense(time.Now().Year())),
		},
		{
			Path:    StrPtr("AUTHORS"),
			Content: StrPtr(templates.GenAuthors()),
		},
		{
			Path:    StrPtr(".air.toml"),
			Content: StrPtr(templates.GenAirToml()),
		},
		{
			Path: StrPtr("migrations"),
		},
		{
			Path: StrPtr(".github"),
		},
		{
			Path: StrPtr("pkg"),
		},
		{
			Path: StrPtr("entity"),
		},
		{
			Path: StrPtr("config"),
		},
		{
			Path: StrPtr("api"),
		},
		{
			Path: StrPtr("usecase"),
		},
		{
			Path:    StrPtr(filepath.Join(".github", "ISSUE_TEMPLATE.md")),
			Content: StrPtr(templates.GenGithubIssueTemplate()),
		},
		{
			Path:    StrPtr(filepath.Join(".github", "PULL_REQUEST_TEMPLATE.md")),
			Content: StrPtr(templates.GenPullRequestTemplate()),
		},
		{
			Path: StrPtr(filepath.Join(".github", "workflows")),
		},
		{
			Path:    StrPtr(filepath.Join(".github", "workflows", "ci.yml")),
			Content: StrPtr(templates.GenGithubCI()),
		},
		{
			Path: StrPtr(filepath.Join("infrastructure", "repository")),
		},
		{
			Path:    StrPtr(filepath.Join("entity", "entity.go")),
			Content: StrPtr(templates.GenBaseEntity()),
		},
		{
			Path:    StrPtr(filepath.Join("config", "config.go")),
			Content: StrPtr(templates.GenProjectConfig()),
		},
		{
			Path: StrPtr(filepath.Join("api", "handler")),
		},
		{
			Path: StrPtr(filepath.Join("api", "middleware")),
		},
		{
			Path: StrPtr(filepath.Join("api", "presenter")),
		},
		{
			Path:    StrPtr(filepath.Join("api", "main.go")),
			Content: StrPtr(templates.GenerateApiMain(moduleName, projectName, dbDriver)),
		},
		{
			Content: StrPtr(templates.GenWelcomeMsg(projectName)),
		},
	}

	for _, o := range operations {
		if o.Path == nil {
			fmt.Println(*o.Content)
			continue
		}

		if o.Content != nil { // its a file
			f, err := os.Create(filepath.Join(p, *o.Path))
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
			if err := os.MkdirAll(filepath.Join(p, *o.Path), os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}
	}

	args := []string{"mod", "tidy"}
	cmd := exec.Command("go", args...)
	cmd.Dir = p
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("name", "n", "", "Project name")
	newCmd.Flags().StringP("path", "p", "", "Project path")
	newCmd.Flags().StringP("module", "m", "", "Module name")
	newCmd.Flags().StringP("driver", "d", "mysql", "Database driver name: mysql|postgres, default: mysql")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func StrPtr(v string) *string {
	return &v
}
