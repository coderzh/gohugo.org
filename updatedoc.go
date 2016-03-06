package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/hugo/hugolib"
	"github.com/spf13/hugo/parser"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const OriginHugoDir string = "../originhugo"
const OriginDocDir string = OriginHugoDir + "/docs/content"
const TargetDocDir string = "content/doc"
const OriginImgDir string = OriginHugoDir + "/docs/static/img"
const TargetImgDir string = "static/img"
const TargetShowcaseDir string = "content/showcase"

func executeCommands(cmdList []string) error {
	for _, cmdString := range cmdList {
		fmt.Println("===>", cmdString)
		cmdArgs := strings.Split(cmdString, " ")
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			log.Fatal(cmdString, out.String(), err)
			return err
		}
		fmt.Println(out.String())
	}
	return nil
}

func updateOriginHugoRepo() error {
	dir, _ := os.Getwd()

	// clone if not exist
	if _, err := os.Stat(OriginHugoDir); os.IsNotExist(err) {
		if cmdErr := executeCommands([]string{
			"git clone https://github.com/spf13/hugo.git " + OriginHugoDir,
		}); cmdErr != nil {
			return err
		}
		fmt.Println("Cloned!")
	}

	os.Chdir(OriginHugoDir)

	if err := executeCommands([]string{
		"git checkout v0.15.docs",
		"git pull --rebase origin v0.15.docs",
	}); err != nil {
		return err
	}

	os.Chdir(dir)

	return nil
}

func convertContentText(docSubDirs []string, content string) string {
	type ReplaceItem struct {
		search  string
		replace string
	}

	replaceList := []ReplaceItem{}

	for _, docSubDir := range docSubDirs {
		replaceList = append(replaceList, ReplaceItem{`(["\(\s])(/?)(` + docSubDir + `)/(\w+)`, "${1}${2}doc/${3}/${4}"})
	}

	for _, replaceItem := range replaceList {
		re := regexp.MustCompile(replaceItem.search)
		content = re.ReplaceAllString(content, replaceItem.replace)
	}

	return content
}

func isFileContentConverted(path string) bool {
	if _, err := os.Stat(path); err == nil {
		if contentCN, err := ioutil.ReadFile(path); err == nil {
			return strings.Contains(string(contentCN), "translated: true")
		}
	}
	return false
}

func saveMarkdownFile(metadata interface{}, content string, path string) error {
	dir, filename := filepath.Split(path)
	page, err := hugolib.NewPage(filename)
	if err != nil {
		return err
	}

	page.SetDir(dir)
	page.SetSourceContent([]byte(content))
	page.SetSourceMetaData(metadata, parser.FormatToLeadRune("yaml"))
	page.SaveSourceAs(path)

	return nil
}

func convertMetadata(path string, metadata map[string]interface{}) (map[string]interface{}, error) {
	keys := []string{"url", "next", "prev"}
	for _, key := range keys {
		if url, ok := metadata[key]; ok {
			metadata[key] = filepath.Join("/doc", url.(string))
		}
	}

	return metadata, nil
}

func getMetadata(path string) (map[string]interface{}, string, error) {
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, "", err
	}

	psr, err := parser.ReadFrom(bytes.NewReader(contentBytes))
	if err != nil {
		return nil, "", err
	}

	m, err := psr.Metadata()
	if err != nil {
		return nil, "", err
	}

	metadata, err := cast.ToStringMapE(m)
	if err != nil {
		return nil, "", err
	}

	return metadata, string(psr.Content()), nil
}

func convertFileContent(docSubDirs []string, path string, pathEN string, pathCN string) error {
	metadata, content, err := getMetadata(path)
	if err != nil {
		return err
	}

	dirname := filepath.Base(filepath.Dir(path))

	newMetadata, err := convertMetadata(path, metadata)
	if err != nil {
		return err
	}

	convertedContent := convertContentText(docSubDirs, content)

	if dirname == "commands" {
		newMetadata["doc"] = []string{"commands_en"}
	}
	if err := saveMarkdownFile(newMetadata, convertedContent, pathEN); err != nil {
		return err
	}

	if !isFileContentConverted(pathCN) {
		if dirname == "commands" {
			newMetadata["doc"] = []string{"commands"}
		}
		if err := saveMarkdownFile(newMetadata, convertedContent, pathCN); err != nil {
			return err
		}
	}

	return nil
}

func copyFile(source string, dest string) error {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}

	}
	return nil
}

func copyDir(source string, dest string) error {
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New(source + " is not a directory")
	}
	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}
	entries, err := ioutil.ReadDir(source)
	for _, entry := range entries {
		sfp := filepath.Join(source, entry.Name())
		dfp := filepath.Join(dest, entry.Name())
		if entry.IsDir() {
			err = copyDir(sfp, dfp)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copyFile(sfp, dfp)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return nil
}

func onDocContentDirWalk(docSubDirs []string, path string, fi os.FileInfo, err error) error {
	if fi.IsDir() {
		return nil
	}

	rel, err := filepath.Rel(OriginDocDir, path)
	if err != nil {
		return err
	}

	dir, file := filepath.Split(rel)
	dirname := filepath.Base(dir)
	ext := filepath.Ext(file)
	fileEn := file[:len(file)-len(ext)] + "_en"

	if dirname == "showcase" {
		if err := copyDir(OriginImgDir, TargetImgDir); err != nil {
			return err
		}
		return copyFile(path, filepath.Join(TargetShowcaseDir, file))
	} else {
		pathEN := filepath.Join(TargetDocDir, dir, fileEn+ext)
		pathCN := filepath.Join(TargetDocDir, rel)
		fmt.Println(pathEN)

		parentDir := filepath.Dir(pathCN)
		if _, err := os.Stat(parentDir); os.IsNotExist(err) {
			os.MkdirAll(parentDir, 0777)
		}

		return convertFileContent(docSubDirs, path, pathEN, pathCN)
	}
}

func getDocSubDirs() ([]string, error) {
	dirs, err := ioutil.ReadDir(OriginDocDir)
	subDirs := make([]string, len(dirs))
	if err != nil {
		return subDirs, err
	}

	for i, dirInfo := range dirs {
		subDirs[i] = dirInfo.Name()
	}

	return subDirs, nil
}

func convertAllDocs() error {
	docSubDirs, err := getDocSubDirs()

	if err != nil {
		return err
	}

	callback := func(path string, fi os.FileInfo, err error) error {
		return onDocContentDirWalk(docSubDirs, path, fi, err)
	}
	return filepath.Walk(filepath.Join(OriginDocDir), callback)
}

func main() {
	if err := updateOriginHugoRepo(); err != nil {
		fmt.Println(err)
	}

	if err := convertAllDocs(); err != nil {
		fmt.Println(err)
	}
}
