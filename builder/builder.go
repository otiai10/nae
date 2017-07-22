package builder

import (
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Builder ...
type Builder struct {
	Name        string
	GoPath      string
	ProjectPath string
	PackagePath string
}

const naepath = "github.com/otiai10/nae"

// SetName ...
func (builder *Builder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("App name is required: usage: nae new {appname}")
	}
	builder.Name = name
	return nil
}

// SetGoPath ...
func (builder *Builder) SetGoPath(p string) error {
	if p == "" {
		return fmt.Errorf("Env $GOPATH is not set")
	}
	builder.GoPath = p
	return nil
}

// SetProjectPath ...
func (builder *Builder) SetProjectPath(p string) error {
	if p == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		if strings.HasPrefix(wd, builder.GoPath) {
			p = filepath.Join(wd, builder.Name)
		} else {
			p = filepath.Join(builder.GoPath, builder.Name)
		}
	}
	builder.ProjectPath = p
	return nil
}

// CopySkeleton ...
func (builder *Builder) CopySkeleton() error {
	pkg, err := build.Import(naepath, "", build.FindOnly)
	if err != nil {
		return err
	}
	if err = copy(filepath.Join(pkg.Dir, "skel"), builder.ProjectPath); err != nil {
		return err
	}
	return nil
}

// EditSource ...
func (builder *Builder) EditSource() error {
	builder.PackagePath = strings.Replace(builder.ProjectPath, builder.GoPath+"/src/", "", 1)
	targets := []string{
		"app/init.go",
		"server/controllers/index.go",
	}
	for _, target := range targets {
		if err := builder.edit(target); err != nil {
			return err
		}
	}
	return nil
}

func (builder *Builder) edit(target string) error {
	target = filepath.Join(builder.ProjectPath, target)
	f, err := os.Open(target)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	if err = os.Remove(target); err != nil {
		return err
	}
	f, err = os.Create(target)
	if err != nil {
		return err
	}
	tpl, err := template.New(target).Parse(string(b))
	if err != nil {
		return err
	}
	return tpl.Execute(f, builder)
}

// Revert ...
func (builder *Builder) Revert() error {
	return os.RemoveAll(builder.ProjectPath)
}

// SuccessMessage ...
func (builder *Builder) SuccessMessage(wr io.Writer) error {
	tpl, err := template.New("success").Parse(success)
	if err != nil {
		return err
	}
	return tpl.Execute(wr, builder)
}
