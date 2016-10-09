package dotrender

import (
	"os"
	"os/exec"
	"io/ioutil"
)

func writeTempFile(content string) (string, error) {
	var dotfile, err = ioutil.TempFile("", "dotfile")
	defer dotfile.Close()

	if err != nil {
		return "", err
	}

	_, err = dotfile.WriteString(content)
	if err != nil {
		return "", err
	}
	return dotfile.Name(), nil
}

func RenderString(src string, outpath string) error {
	/* var dotfile, err = ioutil.TempFile("", "dotfile")
	defer os.Remove(dotfile.Name())
	defer dotfile.Close()
	if err != nil {
		return err
	}

	_, err = dotfile.WriteString(src)
	if err != nil {
		return err
	} */
	var temppath, err = writeTempFile(src)
	defer os.Remove(temppath)
	if err != nil {
		return err
	}
	return RenderFile(temppath, outpath)
}

func RenderFile(inpath string, outpath string) error {
	var outfile, err = os.Create(outpath)
	if err != nil {
		return err
	}
	var cmd = exec.Command("dot", "-Tpng", inpath)
	cmd.Stdout = outfile
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func RenderStringToString(src string) (string, error) {
	var temppath, err = writeTempFile(src)
	defer os.Remove(temppath)
	if err != nil {
		return "", err
	}

	return RenderFileToString(temppath)
}

func RenderFileToString(path string) (string, error) {
	var r, w, err = os.Pipe()
	if err != nil {
		return "", err
	}
	var cmd = exec.Command("dot", "-Tsvg", path)
	cmd.Stdout = w
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", err
	}
	w.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
