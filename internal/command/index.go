package command

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/wfrodriguez/gomposer/internal/data"
	"github.com/wfrodriguez/gomposer/internal/tpl"
	"github.com/wfrodriguez/gomposer/internal/ui"
	"github.com/wfrodriguez/mimir/fs"
	"github.com/wfrodriguez/mimir/text"
	"gopkg.in/yaml.v3"
)

type ymlMap = map[string]any
type dtag struct {
	Tag  string
	Slug string
}

var monthMap = map[string]string{
	"Enero":      "01",
	"Febrero":    "02",
	"Marzo":      "03",
	"Abril":      "04",
	"Mayo":       "05",
	"Junio":      "06",
	"Julio":      "07",
	"Agosto":     "08",
	"Septiembre": "09",
	"Octubre":    "10",
	"Noviembre":  "11",
	"Diciembre":  "12",
}
var reDateValid *regexp.Regexp

func init() {
	var err error
	reDateValid, err = regexp.Compile(`(?i)^[0-9]{2}\s+(enero|febrero|marzo|abril|mayo|junio|julio|agosto|septiembre|octubre|noviembre|diciembre)\s+\d{4}$`)
	if err != nil {
		panic(err) // TODO: handle error
	}
}

func convertDate(date string) int64 {
	if reDateValid.MatchString(date) {
		// fmt.Print(" ", date, " ")
		for k, v := range monthMap {
			date = strings.ReplaceAll(date, k, v)
		}
		t, err := time.Parse("02 01 2006", date)
		if err != nil {
			panic(err) // TODO: handle error
		}
		// fmt.Println("->", t.Format("2006-01-02"), "")

		return t.Unix()
	}

	return 0

}

func IndexProject(projectDir string) {
	db, err := data.NewMemDB()
	if err != nil {
		panic(err) // TODO: handle error
	}
	defer db.Close()

	if err = data.CreatePostsTable(db); err != nil {
		panic(err) // TODO: handle error
	}

	// fmt.Println("Project:", projectDir)

	// Generar Tags
	buildTags(projectDir, db)
	// Generar listado de posts
	buildPosts(projectDir, db)
}

func buildTags(projectDir string, db *sql.DB) {
	files, err := fs.FindFileByExt(filepath.Join(projectDir, "post"), fs.Markdown)
	if err != nil {
		panic(err) // TODO: handle error
	}

	for _, file := range files {
		fmt.Println(" ", file)
		fileStat, err := os.Stat(file)
		if err != nil {
			panic(err) // TODO: handle error
		}
		name := fileStat.Name()
		slug, _ := strings.CutSuffix(name, filepath.Ext(name))
		var b []byte
		b, err = extractHeader(file)
		if err != nil {
			panic(err) // TODO: handle error
		}
		var ymap ymlMap
		err = yaml.Unmarshal(b, &ymap)
		if err != nil {
			panic(err) // TODO: handle error
		}
		title, titleOk := ymap["title"]
		date, dateOk := ymap["date"]
		tags, tagsOk := ymap["tags"]
		desc, descOk := ymap["description"]

		if tagsOk {
			atags := tags.([]interface{})
			for _, tag := range atags {
				err := data.SaveTag(db, tag.(string))
				if err != nil {
					panic(err) // TODO: handle error
				}
			}
		}
		var tit, dat, des string = "", "", ""
		var tgs = []string{}
		if titleOk {
			tit = title.(string)
		}
		if dateOk {
			dat = date.(string)
		}
		if tagsOk {
			t := tags.([]interface{})
			for _, tag := range t {
				tgs = append(tgs, tag.(string))
			}
		}
		if descOk {
			des = desc.(string)
		}

		err = data.SavePost(db, tit, slug, des, convertDate(dat), tgs)
		if err != nil {
			panic(err) // TODO: handle error
		}
	}
	generateTags(projectDir, db)
}

func buildPosts(projectDir string, db *sql.DB) {
	projectDir = filepath.Join(projectDir, "build")
	posts, err := data.GetPosts(db)
	var tmp []byte
	if err != nil {
		panic(err) // TODO: handle error
	}
	data := map[string]any{
		"posts": posts,
	}
	tmp, err = ui.Render(tpl.AllPosts, data)
	file := filepath.Join(projectDir, "posts.md")
	err = os.WriteFile(file, tmp, 0666)
	if err != nil {
		panic(err) // TODO: handle error
	}
	fmt.Println(" ", file)
}

func generateTags(projectDir string, db *sql.DB) {
	dataTags := make([]dtag, 0)
	tags, err := data.GetTags(db)
	if err != nil {
		panic(err) // TODO: handle error
	}
	tagDir := filepath.Join(projectDir, "tag")
	for _, tag := range tags {
		fmt.Print(" ", tag, " ")
		posts, err := data.GetPostsByTag(db, tag)
		if err != nil {
			panic(err) // TODO: handle error
		}
		datos := map[string]any{
			"tag":   tag,
			"posts": posts,
		}
		tmp, err := ui.Render(tpl.TplTag, datos)
		if err != nil {
			panic(err) // TODO: handle error
		}

		slug := text.Slugify(tag)
		dataTags = append(dataTags, dtag{Tag: tag, Slug: slug})
		err = os.WriteFile(filepath.Join(tagDir, fmt.Sprintf("%s.md", slug)), tmp, 0666)
		if err != nil {
			panic(err) // TODO: handle error
		}
	}
	fmt.Println()

	datos := map[string]any{
		"tags": dataTags,
	}
	tmp, err := ui.Render(tpl.AllTags, datos)
	if err != nil {
		panic(err) // TODO: handle error
	}

	projectDir = filepath.Join(projectDir, "build")
	file := filepath.Join(projectDir, "tags.md")
	err = os.WriteFile(file, tmp, 0666)
	if err != nil {
		panic(err) // TODO: handle error
	}
	fmt.Println(" ", file)
}

func extractHeader(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}

	lines := strings.Split(string(file), "\n")
	yml := make([]string, 0)
	for i, line := range lines {
		if strings.HasPrefix(line, "---") && i > 0 {
			break
		} else if strings.HasPrefix(line, "---") {
			continue
		}

		yml = append(yml, line)
	}

	return []byte(strings.Join(yml, "\n")), nil
}
