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
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Bar":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bar Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarFormCallback(
			nil,
			probe,
			formGroup,
		)
		bar := new(models.Bar)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bar, formGroup, probe)
	case "Key":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Key Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KeyFormCallback(
			nil,
			probe,
			formGroup,
		)
		key := new(models.Key)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(key, formGroup, probe)
	case "Pie":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Pie Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PieFormCallback(
			nil,
			probe,
			formGroup,
		)
		pie := new(models.Pie)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pie, formGroup, probe)
	case "Scatter":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Scatter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScatterFormCallback(
			nil,
			probe,
			formGroup,
		)
		scatter := new(models.Scatter)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(scatter, formGroup, probe)
	case "Serie":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Serie Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SerieFormCallback(
			nil,
			probe,
			formGroup,
		)
		serie := new(models.Serie)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(serie, formGroup, probe)
	case "Value":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Value Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ValueFormCallback(
			nil,
			probe,
			formGroup,
		)
		value := new(models.Value)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(value, formGroup, probe)
	}
	formStage.Commit()
}
