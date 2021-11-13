package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"

	qrcode "github.com/skip2/go-qrcode"
)

func process(link, output string) {
	err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)
	if err != nil {
		log.Fatalf("Qrcode create failed!")
	}

	qr, err := os.Open(output)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err.Error())
	}
	defer qr.Close()

	cmd := exec.Command("imgcat", output)
	stdout, errCmd := cmd.Output()
	if errCmd != nil {
		log.Fatalf("Cannot run command: %v", errCmd.Error())
	}
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

	process(link, output)
}
