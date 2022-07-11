// MIT License

// Copyright (c) 2022 Tree Xie

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package charts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieChart(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		render func(*Painter) ([]byte, error)
		result string
	}{
		{
			render: func(p *Painter) ([]byte, error) {
				values := []float64{
					1048,
					735,
					580,
					484,
					300,
				}
				_, err := NewPieChart(p, PieChartOption{
					SeriesList: NewPieSeriesList(values, PieSeriesOption{
						Label: SeriesLabel{
							Show: true,
						},
					}),
					Title: TitleOption{
						Text:    "Rainfall vs Evaporation",
						Subtext: "Fake Data",
						Left:    PositionCenter,
					},
					Padding: Box{
						Top:    20,
						Right:  20,
						Bottom: 20,
						Left:   20,
					},
					Legend: LegendOption{
						Orient: OrientVertical,
						Data: []string{
							"Search Engine",
							"Direct",
							"Email",
							"Union Ads",
							"Video Ads",
						},
						Left: PositionLeft,
					},
				}).Render()
				if err != nil {
					return nil, err
				}
				return p.Bytes()
			},
			result: "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"600\" height=\"400\">\\n<path  d=\"M 0 0\nL 560 0\nL 560 360\nL 0 360\nL 0 0\" style=\"stroke-width:0;stroke:none;fill:rgba(255,255,255,1.0)\"/><path  d=\"M 40 49\nL 70 49\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><circle cx=\"55\" cy=\"49\" r=\"5\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"72\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Search Engine</text><path  d=\"M 40 69\nL 70 69\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><circle cx=\"55\" cy=\"69\" r=\"5\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"72\" y=\"75\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Direct</text><path  d=\"M 40 89\nL 70 89\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><circle cx=\"55\" cy=\"89\" r=\"5\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"72\" y=\"95\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Email</text><path  d=\"M 40 109\nL 70 109\" style=\"stroke-width:3;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><circle cx=\"55\" cy=\"109\" r=\"5\" style=\"stroke-width:3;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"72\" y=\"115\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Union Ads</text><path  d=\"M 40 129\nL 70 129\" style=\"stroke-width:3;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><circle cx=\"55\" cy=\"129\" r=\"5\" style=\"stroke-width:3;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"\" style=\"stroke-width:3;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"72\" y=\"135\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Video Ads</text><text x=\"222\" y=\"55\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Rainfall vs Evaporation</text><text x=\"266\" y=\"70\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:15.3px;font-family:'Roboto Medium',sans-serif\">Fake Data</text><path  d=\"M 300 250\nL 300 186\nA 64 64 119.89 0 1 355 281\nL 300 250\nZ\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><path  d=\"M 355 218\nL 368 211\nM 368 211\nL 383 211\" style=\"stroke-width:1;stroke:rgba(84,112,198,1.0);fill:rgba(84,112,198,1.0)\"/><text x=\"386\" y=\"216\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Search Engine: 33.3%</text><path  d=\"M 300 250\nL 355 281\nA 64 64 84.08 0 1 275 308\nL 300 250\nZ\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><path  d=\"M 319 310\nL 324 325\nM 324 325\nL 339 325\" style=\"stroke-width:1;stroke:rgba(145,204,117,1.0);fill:rgba(145,204,117,1.0)\"/><text x=\"342\" y=\"330\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Direct: 23.35%</text><path  d=\"M 300 250\nL 275 308\nA 64 64 66.35 0 1 237 250\nL 300 250\nZ\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><path  d=\"M 247 284\nL 234 292\nM 234 292\nL 219 292\" style=\"stroke-width:1;stroke:rgba(250,200,88,1.0);fill:rgba(250,200,88,1.0)\"/><text x=\"135\" y=\"297\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Email: 18.43%</text><path  d=\"M 300 250\nL 237 250\nA 64 64 55.37 0 1 264 198\nL 300 250\nZ\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><path  d=\"M 244 220\nL 231 213\nM 231 213\nL 216 213\" style=\"stroke-width:1;stroke:rgba(238,102,102,1.0);fill:rgba(238,102,102,1.0)\"/><text x=\"105\" y=\"218\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Union Ads: 15.37%</text><path  d=\"M 300 250\nL 264 198\nA 64 64 34.32 0 1 300 186\nL 300 250\nZ\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><path  d=\"M 282 189\nL 277 175\nM 277 175\nL 262 175\" style=\"stroke-width:1;stroke:rgba(115,192,222,1.0);fill:rgba(115,192,222,1.0)\"/><text x=\"159\" y=\"180\" style=\"stroke-width:0;stroke:none;fill:rgba(70,70,70,1.0);font-size:12.8px;font-family:'Roboto Medium',sans-serif\">Video Ads: 9.53%</text></svg>",
		},
	}
	for _, tt := range tests {
		p, err := NewPainter(PainterOptions{
			Type:   ChartOutputSVG,
			Width:  600,
			Height: 400,
		}, PainterThemeOption(defaultTheme))
		assert.Nil(err)
		data, err := tt.render(p.Child(PainterPaddingOption(Box{
			Left:   20,
			Top:    20,
			Right:  20,
			Bottom: 20,
		})))
		assert.Nil(err)
		assert.Equal(tt.result, string(data))
	}
}
