package main

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

const OutputDir = "output"

func setupOutputDir() error {
	err := os.Mkdir(OutputDir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func plotDistribution(subscribers []Subscriber, dist Distribution, number int) error {
	err := setupOutputDir()

	if err != nil {
		return err
	}

	p := plot.New()
	p.Title.Text = fmt.Sprintf("Subscribers Distribution - %s Distribution - Coverage Radius %d", dist, number)
	p.Title.TextStyle.Font.Size = vg.Points(20)

	v := make(plotter.Values, len(subscribers))
	for i := range v {
		v[i] = subscribers[i].distance
	}

	h, err := plotter.NewHist(v, 16)
	if err != nil {
		return err
	}

	h.Normalize(1)
	p.Add(h)

	if err := p.Save(10*vg.Inch, 10*vg.Inch, fmt.Sprintf("output/distribution_%s_%d.png", dist, number)); err != nil {
		return err
	}

	return nil
}

func plotSubscribers(subscribers []Subscriber, dist Distribution, number int) error {
	err := setupOutputDir()

	if err != nil {
		return err
	}

	p := plot.New()

	pts := make(plotter.XYs, len(subscribers))

	for i := range pts {
		pts[i].X = subscribers[i].location.x
		pts[i].Y = subscribers[i].location.y
	}

	p.Title.Text = fmt.Sprintf("Subscribers Plot - %s Distribution - Coverage Radius %d", dist, number)
	p.Title.TextStyle.Font.Size = vg.Points(20)
	p.X.Label.Text = "X-axis"
	p.Y.Label.Text = "Y-axis"

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}

	s.Color = color.RGBA{R: 255, A: 255}
	s.Shape = draw.CircleGlyph{}
	s.Radius = vg.Points(.8)

	p.Add(s)

	if err := p.Save(10*vg.Inch, 10*vg.Inch, fmt.Sprintf("output/subscribers_%s_%d.png", dist, number)); err != nil {
		return err
	}

	return nil
}
