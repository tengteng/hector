package eval

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"
)

func ROC(predictions0 []*LabelPrediction) {
	result := []float64{}

	predictions := []*LabelPrediction{}
	for _, pred := range predictions0 {
		predictions = append(predictions, pred)
	}
	prediction := func(p1, p2 *LabelPrediction) bool {
		return p1.Prediction > p2.Prediction
	}

	// predictions in descending order.
	By(prediction).Sort(predictions)

	total_p := 0.0
	for _, lp := range predictions {
		if lp.Label > 0 {
			total_p += 1.0
		}
	}

	pn := 0.0
	for _, lp := range predictions {
		if lp.Label > 0 {
			pn += 1.0
		}
		result = append(result, pn/total_p)
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "ROC"
	p.X.Label.Text = "FP"
	p.Y.Label.Text = "TP"

	pts := make(plotter.XYs, len(result))
	for i := range pts {
		pts[i].X = float64(i+1) / float64(len(result))
		pts[i].Y = result[i]
	}

	err = plotutil.AddLinePoints(p, "ROC", pts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(64, 64, "ROC.png"); err != nil {
		panic(err)
	}
}
