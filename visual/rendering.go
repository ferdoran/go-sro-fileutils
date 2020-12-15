package visual

import (
	"fmt"
	"github.com/ferdoran/go-sro-fileutils/navmesh"
	"github.com/fogleman/gg"
	"math"
	"math/rand"
)

func NewRendering(data map[string]navmesh.NavMeshData) {
	ctx := gg.NewContext(1920, 1920)
	//counter := 0
	//max := len(data)
	ctx.SetHexColor("FFFFFF")
	ctx.SetLineWidth(2.0)
	ctx.DrawRectangle(0, 0, 1920, 1920)
	ctx.Fill()
	ctx.SetHexColor("FF0000")
	// draw tiles
	for x := 0; x < 96; x++ {
		for y := 0; y < 96; y++ {
			ctx.Push()
			ctx.DrawRectangle(float64(x*20), float64(y*20), 20, 20)
			ctx.Stroke()
			ctx.Pop()
		}
	}

	ctx.SetLineWidth(5.0)
	d := data["nv_62a8.nvm"]
	//for k, d := range data {
	//	ctx.Clear()
	//	//ctx.SetRGB(0,255,0)
	//	ctx.SetLineWidth(5.0)
	//	ctx.SetHexColor("FFFFFF")
	//	ctx.DrawRectangle(0, 0, 1920, 1920)
	//	ctx.Fill()
	//
	//	ctx.SetHexColor("0000FF")
	//
	ctx.SetHexColor("0000FF")
	for i, c := range d.Cells {
		//randomColor(ctx)
		ctx.Push()
		dx := math.Abs(float64(c.Min.X - c.Max.X))
		dy := math.Abs(float64(c.Min.Y - c.Max.Y))
		ctx.DrawRectangle(float64(c.Min.X), float64(c.Min.Y), dx, dy)
		ctx.DrawString(fmt.Sprintf("C[%d]", i), float64(c.Min.X+(0.5*(c.Max.X-c.Min.X))), float64(c.Min.Y+(0.5*(c.Max.Y-c.Min.Y))))
		//ctx.DrawString(fmt.Sprintf("(%d|%d)", int(c.Min.X), int(c.Min.Y)), float64(c.Min.X), float64(c.Min.Y))
		//ctx.DrawString(fmt.Sprintf("(%d|%d)", int(c.Min.X), int(c.Max.Y)), float64(c.Min.X), float64(c.Max.Y))
		//ctx.DrawString(fmt.Sprintf("(%d|%d)", int(c.Max.X), int(c.Min.Y)), float64(c.Max.X), float64(c.Min.Y))
		//ctx.DrawString(fmt.Sprintf("(%d|%d)", int(c.Max.X), int(c.Max.Y)), float64(c.Max.X), float64(c.Max.Y))
		ctx.Stroke()
		ctx.Pop()
	}
	ctx.SavePNG("image2.png")
	//	k = strings.ReplaceAll(k, ".nvm", ".jpg")
	//	ctx.SaveJPG("E:\\SRO\\images\\" + k, 80)
	//	counter++
	//	fmt.Printf("\rFinished drawing image [%d / %d]", counter, max)
	//}
}

func randomColor(ctx *gg.Context) {
	r := rand.Float64()
	g := rand.Float64()
	b := rand.Float64()
	ctx.SetRGB(r, g, b)
}
