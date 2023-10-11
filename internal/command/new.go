package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wfrodriguez/console"
	"github.com/wfrodriguez/gomposer/cfg"
	"github.com/wfrodriguez/gomposer/internal/ui"
	"github.com/wfrodriguez/mimir/text"
)

const (
	BasePerm = 0764
)

var empty = map[string]any{}

var skeletonDir = []string{
	filepath.Join(cfg.DistDir, "post"),
	filepath.Join(cfg.DistDir, "tag"),
	filepath.Join(cfg.PostDir),
	filepath.Join(cfg.StaticDir, "img", "ui"),
	filepath.Join(cfg.StaticDir, "js"),
	filepath.Join(cfg.StaticDir, "css"),
	filepath.Join(cfg.StaticDir, "fonts"),
	filepath.Join(cfg.StaticDir, "html"),
	filepath.Join(cfg.TemplateDir),
	filepath.Join(cfg.BuildDir, cfg.TagDir),
}

func skelMkdirs(baseDir string) error {
	err := os.MkdirAll(baseDir, BasePerm)
	if err != nil {
		return err
	}
	for _, dir := range skeletonDir {
		err = os.MkdirAll(filepath.Join(baseDir, dir), BasePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func fjoin(path ...string) string {
	return filepath.Join(path...)
}

func writeSkeletonFile(baseDir, name, out string, render bool, data map[string]any) error {
	var s []byte
	var err error
	var tpl []byte
	s, err = ui.Skeleton.ReadFile(name)
	if err != nil {
		return err
	}

	if render {
		tpl, err = ui.Render(string(s), data)
		if err != nil {
			return err
		}
	} else {
		tpl = s
	}

	console.Info(fmt.Sprintf("󰁕 Creando archivo %s", out))

	err = os.WriteFile(filepath.Join(baseDir, out), tpl, BasePerm)
	if err != nil {
		return err
	}
	return nil
}

func NewProject(baseDir, name string) {
	var err error

	slug := text.Slugify(name)
	baseDir = filepath.Join(baseDir, slug)
	err = skelMkdirs(baseDir)
	if err != nil {
		panic(err) // TODO: handle error
	}

	// Crear archivos base
	err = writeSkeletonFile(baseDir, "skeleton/gitignore", ".gitignore", false, empty)
	if err != nil {
		panic(err) // TODO: handle error
	}
	err = writeSkeletonFile(baseDir, "skeleton/Makefile", "Makefile", true, map[string]any{"Name": name})
	if err != nil {
		panic(err) // TODO: handle error
	}
	err = writeSkeletonFile(baseDir, "skeleton/index.html", fjoin("template", "index.html"), false, empty)
	if err != nil {
		panic(err) // TODO: handle error
	}
	err = writeSkeletonFile(baseDir, "skeleton/post.html", fjoin("template", "post.html"), false, empty)
	if err != nil {
		panic(err) // TODO: handle error
	}
	err = writeSkeletonFile(baseDir, "skeleton/tag.html", fjoin("template", "tag.html"), false, empty)
	if err != nil {
		panic(err) // TODO: handle error
	}
	err = writeSkeletonFile(baseDir, "skeleton/skel.css", fjoin("static", "css", "skel.css"), false, empty)
	if err != nil {
		panic(err) // TODO: handle error
	}
	err = writeSkeletonFile(baseDir, "skeleton/style.css", fjoin("static", "css", "style.css"), false, empty)
	if err != nil {
		panic(err) // TODO: handle error
	}
	err = writeSkeletonFile(baseDir, "skeleton/markdown-pandoc-reference.md", fjoin("post", "markdown-pandoc-reference.md"), true, empty)
	if err != nil {
		panic(err) // TODO: handle error
	}
	err = writeSkeletonFile(baseDir, "skeleton/index.md", "index.md", true, empty)
	if err != nil {
		panic(err) // TODO: handle error
	}
}
