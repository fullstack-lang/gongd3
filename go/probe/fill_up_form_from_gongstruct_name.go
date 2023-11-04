// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongd3/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Bar":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Bar Form",
			OnSave: __gong__New__BarFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		bar := new(models.Bar)
		FillUpForm(bar, formGroup, probe)
	case "Key":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Key Form",
			OnSave: __gong__New__KeyFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		key := new(models.Key)
		FillUpForm(key, formGroup, probe)
	case "Pie":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Pie Form",
			OnSave: __gong__New__PieFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		pie := new(models.Pie)
		FillUpForm(pie, formGroup, probe)
	case "Scatter":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Scatter Form",
			OnSave: __gong__New__ScatterFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		scatter := new(models.Scatter)
		FillUpForm(scatter, formGroup, probe)
	case "Serie":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Serie Form",
			OnSave: __gong__New__SerieFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		serie := new(models.Serie)
		FillUpForm(serie, formGroup, probe)
	case "Value":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Value Form",
			OnSave: __gong__New__ValueFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		value := new(models.Value)
		FillUpForm(value, formGroup, probe)
	}
	formStage.Commit()
}
