package main

import (
    "testing"
    "image/color"

    "github.com/i0Ek3/asrt"
    qrcode "github.com/skip2/go-qrcode"
)

func TestEqr(t *testing.T) {
    t.Run("Test1", func(t *testing.T) {
        link := "https://"
        output := ""

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        asrt.Asrt(t, err, nil)
    })

    t.Run("Test2", func(t *testing.T) {
        link := "https://github.com"
        output := "github"

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        asrt.Asrt(t, err, nil)
    })

    /*t.Run("Test3", func(t *testing.T) {
        link := ""
        output := "nothing"

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        asrt.Asrt(t, err, nil)
    })*/

    t.Run("Test4", func(t *testing.T) {
        link := "https://github.com"
        output := ""

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        asrt.Asrt(t, err, nil)
    })

    /*t.Run("Test5", func(t *testing.T) {
        link := ""
        output := ""

	    err := qrcode.WriteColorFile(link, qrcode.Medium, 256, color.Black, color.White, output)

        asrt.Asrt(t, err, nil)
    })*/
}
