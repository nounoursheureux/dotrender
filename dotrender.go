package dotrender

import (
	"os"
	"os/exec"
	"io/ioutil"
)

func RenderString(src string, outpath string) error {
	var dotfile, err = ioutil.TempFile("", "dotfile")
	defer os.Remove(dotfile.Name())
	defer dotfile.Close()
	if err != nil {
		return err
	}

	_, err = dotfile.WriteString(src)
	if err != nil {
		return err
	}
	return RenderFile(dotfile.Name(), outpath)
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
