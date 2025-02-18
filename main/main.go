package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/wcharczuk/go-chart/v2"

	charts "github.com/vicanso/go-charts/v2"
)

func writeFile(buf []byte) error {
	tmpPath := "./main"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, "time-line-chart.png")
	err = os.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	xAxisValue := []string{}
	values := []float64{}
	now := time.Now()
	firstAxis := 0
	for i := 0; i < 300; i++ {
		if firstAxis == 0 && now.Minute() == 0 {
			firstAxis = i
		}
		xAxisValue = append(xAxisValue, now.Format("15:04"))
		now = now.Add(time.Minute)
		value, _ := rand.Int(rand.Reader, big.NewInt(100))
		if (i%50)/10 < 3 {
			values = append(values, math.NaN())
		} else {
			values = append(values, float64(value.Int64()))
		}
	}
	p, err := charts.LineRender(
		[][]float64{
			values,
		},
		charts.TitleTextOptionFunc("Line"),
		charts.XAxisDataOptionFunc(xAxisValue, charts.FalseFlag()),
		charts.LegendLabelsOptionFunc([]string{
			"Demo",
		}, "50"),
		func(opt *charts.ChartOption) {
			opt.SeriesList[0].MarkLine = charts.SeriesMarkLine{
				Data: []charts.SeriesMarkData{
					{
						Type:       charts.SeriesMarkDataTypeCustom,
						CustomYVal: -1,
						FillColor: &charts.Color{
							R: 240,
							G: 0,
							B: 0,
							A: 255,
						},
						StrokeColor: &charts.Color{
							R: 240,
							G: 0,
							B: 0,
							A: 150,
						},
						HideValue:          true,
						IgnoreStrokeDashed: true,
						IgnoreArrow:        true,
						StrokeWidth:        4,
						XAxisIndex:         200,
						XAxisEndIndex:      250,
					},
					{
						Type:       charts.SeriesMarkDataTypeCustom,
						CustomYVal: 80,
						FillColor: &charts.Color{
							R: 240,
							G: 0,
							B: 0,
							A: 255,
						},
						StrokeColor: &charts.Color{
							R: 240,
							G: 0,
							B: 0,
							A: 255,
						},
						AboveColor: &charts.Color{
							R: 240,
							G: 0,
							B: 0,
							A: 20,
						},
					},
					{
						Type:       charts.SeriesMarkDataTypeCustom,
						CustomYVal: 50,
						FillColor: &charts.Color{
							R: 255,
							G: 165,
							B: 0,
							A: 255,
						},
						StrokeColor: &charts.Color{
							R: 255,
							G: 165,
							B: 0,
							A: 255,
						},
						BelowColor: &charts.Color{
							R: 255,
							G: 165,
							B: 0,
							A: 20,
						},
					},
				},
			}
			opt.XAxis.FirstAxis = firstAxis
			opt.XAxis.SplitNumber = 60
			opt.Legend.Padding = charts.Box{
				Top:    5,
				Bottom: 10,
			}
			opt.SymbolShow = charts.FalseFlag()
			opt.LineStrokeWidth = 1
			opt.ValueFormatter = func(f float64) string {
				return fmt.Sprintf("%.0f", f)
			}
		},
		charts.PaddingOptionFunc(chart.NewBox(10, 10, 30, 10)),
	)

	if err != nil {
		panic(err)
	}

	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}
	err = writeFile(buf)
	if err != nil {
		panic(err)
	}
}
