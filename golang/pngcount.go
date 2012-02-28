package main

import (
    "fmt"
    "io"
//    "strings"
    /* and more */
    "os"
    pngimage "image/png"
)

func CountColor(png io.Reader) int {
    cnt := 0
    var color [256*256*4]uint64
    img, err := pngimage.Decode(png)
    if err != nil {
        return 0
    }
    border := img.Bounds()
    for n := border.Min.Y; n < border.Max.Y; n++ {
        for m := border.Min.X; m < border.Max.X; m++ {
            colr := img.At(m,n)
            r0, g0, b0, _ := colr.RGBA()
            index := (((r0 >> 8) << 16) | ((g0 >> 8) << 8) | (b0 >> 8)) >> 6
            var flag uint64 = 1 << ((b0 >> 8) & 0x3f)
            if (color[index] & flag == 0) {
                color[index] |= flag
                cnt += 1
            }
        }
    }
    return cnt
}

func main() {
    png := GetPngBinary()
    cnt := CountColor(png)
    fmt.Println(cnt)
}

func GetPngBinary() io.Reader {
    filename := "sample_google_logo.png"
    file, _ := os.Open(filename)
    return file
} 
