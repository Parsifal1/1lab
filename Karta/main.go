package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fogleman/gg"
	geojson "github.com/paulmach/go.geojson"
)

func main() {

	file, err := os.Open("one.geojson")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	raw := ""
	for {
		r, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		raw = raw + string(data[:r])
	}

	rawGeometryJSON := []byte(raw)

	fc1, err := geojson.UnmarshalFeatureCollection(rawGeometryJSON)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%f", fc1.Features[0].Geometry.Polygon[0][0][0])

	dc := gg.NewContext(1366, 1024)
	dc.SetHexColor("fff")

	dc.InvertY()
	dc.Scale(8, 8)
	dc.MoveTo(fc1.Features[0].Geometry.Polygon[0][0][0], fc1.Features[0].Geometry.Polygon[0][0][1])
	for i := 0; i < len(fc1.Features); i++ {
		for j := 1; j < len(fc1.Features[i].Geometry.Polygon[0]); j++ {
			dc.LineTo(fc1.Features[i].Geometry.Polygon[0][j][0], fc1.Features[i].Geometry.Polygon[0][j][1])
		}
	}

	dc.SetHexColor("f00")
	dc.Fill()
	dc.SavePNG("out.png")

}
