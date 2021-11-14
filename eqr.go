package main

import (
	"fmt"
	"image/color"
	"os"
	"os/exec"

    "github.com/i0Ek3/noerr"
	qrcode "github.com/skip2/go-qrcode"
)

func Process(link, output string) {
	err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)
    //msg := "Qrcode create failed"
    noerr.Xerr(err)

	qr, err := os.Open(output)
    //msg = "Cannot open the file"
    noerr.Xerr(err)
	defer qr.Close()

	cmd := exec.Command("imgcat", output)
	stdout, errCmd := cmd.Output()
    //msg = "Cannot run command"
    noerr.Xerr(errCmd/*, msg*/)
	fmt.Println(string(stdout))
}

func runIO(input *string, output string) {
	fmt.Printf(output)
	fmt.Scan(input)
}

const (
	MSG1 = "Please input the link you want to convert: "
	MSG2 = "Please input the name you want to save: "
)

func main() {
	var link, output string

	runIO(&link, MSG1)
	runIO(&output, MSG2)

	Process(link, output)
}
