package main

import (
    "testing"
    "image/color"

    qrcode "github.com/skip2/go-qrcode"
)

func TestEqr(t *testing.T) {
    t.Run("Test1", func(t *testing.T) {
        link := "https://"
        output := ""

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        if err != nil {
            t.Errorf("Wrong!")
        }
    })

    t.Run("Test2", func(t *testing.T) {
        link := "https://github.com"
        output := "github"

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        if err != nil {
            t.Errorf("Wrong!")
        }
    })

    /*t.Run("Test3", func(t *testing.T) {
        link := ""
        output := "nothing"

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        if err != nil {
            t.Errorf("Wrong!")
        }
    })*/

    t.Run("Test4", func(t *testing.T) {
        link := "https://github.com"
        output := ""

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        if err != nil {
            t.Errorf("Wrong!")
        }
    })

    /*t.Run("Test5", func(t *testing.T) {
        link := ""
        output := ""

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        if err != nil {
            t.Errorf("Wrong!")
        }
    })*/
}
