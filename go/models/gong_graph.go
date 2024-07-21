// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Bar:
		ok = stage.IsStagedBar(target)

	case *Key:
		ok = stage.IsStagedKey(target)

	case *Pie:
		ok = stage.IsStagedPie(target)

	case *Scatter:
		ok = stage.IsStagedScatter(target)

	case *Serie:
		ok = stage.IsStagedSerie(target)

	case *Value:
		ok = stage.IsStagedValue(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedBar(bar *Bar) (ok bool) {

	_, ok = stage.Bars[bar]

	return
}

func (stage *StageStruct) IsStagedKey(key *Key) (ok bool) {

	_, ok = stage.Keys[key]

	return
}

func (stage *StageStruct) IsStagedPie(pie *Pie) (ok bool) {

	_, ok = stage.Pies[pie]

	return
}

func (stage *StageStruct) IsStagedScatter(scatter *Scatter) (ok bool) {

	_, ok = stage.Scatters[scatter]

	return
}

func (stage *StageStruct) IsStagedSerie(serie *Serie) (ok bool) {

	_, ok = stage.Series[serie]

	return
}

func (stage *StageStruct) IsStagedValue(value *Value) (ok bool) {

	_, ok = stage.Values[value]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Bar:
		stage.StageBranchBar(target)

	case *Key:
		stage.StageBranchKey(target)

	case *Pie:
		stage.StageBranchPie(target)

	case *Scatter:
		stage.StageBranchScatter(target)

	case *Serie:
		stage.StageBranchSerie(target)

	case *Value:
		stage.StageBranchValue(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchBar(bar *Bar) {

	// check if instance is already staged
	if IsStaged(stage, bar) {
		return
	}

	bar.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bar.X != nil {
		StageBranch(stage, bar.X)
	}
	if bar.Y != nil {
		StageBranch(stage, bar.Y)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _serie := range bar.Set {
		StageBranch(stage, _serie)
	}

}

func (stage *StageStruct) StageBranchKey(key *Key) {

	// check if instance is already staged
	if IsStaged(stage, key) {
		return
	}

	key.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPie(pie *Pie) {

	// check if instance is already staged
	if IsStaged(stage, pie) {
		return
	}

	pie.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pie.X != nil {
		StageBranch(stage, pie.X)
	}
	if pie.Y != nil {
		StageBranch(stage, pie.Y)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _serie := range pie.Set {
		StageBranch(stage, _serie)
	}

}

func (stage *StageStruct) StageBranchScatter(scatter *Scatter) {

	// check if instance is already staged
	if IsStaged(stage, scatter) {
		return
	}

	scatter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scatter.X != nil {
		StageBranch(stage, scatter.X)
	}
	if scatter.Y != nil {
		StageBranch(stage, scatter.Y)
	}
	if scatter.Text != nil {
		StageBranch(stage, scatter.Text)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _serie := range scatter.Set {
		StageBranch(stage, _serie)
	}

}

func (stage *StageStruct) StageBranchSerie(serie *Serie) {

	// check if instance is already staged
	if IsStaged(stage, serie) {
		return
	}

	serie.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if serie.Key != nil {
		StageBranch(stage, serie.Key)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _value := range serie.Values {
		StageBranch(stage, _value)
	}

}

func (stage *StageStruct) StageBranchValue(value *Value) {

	// check if instance is already staged
	if IsStaged(stage, value) {
		return
	}

	value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Bar:
		toT := CopyBranchBar(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key:
		toT := CopyBranchKey(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pie:
		toT := CopyBranchPie(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scatter:
		toT := CopyBranchScatter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Serie:
		toT := CopyBranchSerie(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Value:
		toT := CopyBranchValue(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchBar(mapOrigCopy map[any]any, barFrom *Bar) (barTo *Bar) {

	// barFrom has already been copied
	if _barTo, ok := mapOrigCopy[barFrom]; ok {
		barTo = _barTo.(*Bar)
		return
	}

	barTo = new(Bar)
	mapOrigCopy[barFrom] = barTo
	barFrom.CopyBasicFields(barTo)

	//insertion point for the staging of instances referenced by pointers
	if barFrom.X != nil {
		barTo.X = CopyBranchKey(mapOrigCopy, barFrom.X)
	}
	if barFrom.Y != nil {
		barTo.Y = CopyBranchKey(mapOrigCopy, barFrom.Y)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _serie := range barFrom.Set {
		barTo.Set = append(barTo.Set, CopyBranchSerie(mapOrigCopy, _serie))
	}

	return
}

func CopyBranchKey(mapOrigCopy map[any]any, keyFrom *Key) (keyTo *Key) {

	// keyFrom has already been copied
	if _keyTo, ok := mapOrigCopy[keyFrom]; ok {
		keyTo = _keyTo.(*Key)
		return
	}

	keyTo = new(Key)
	mapOrigCopy[keyFrom] = keyTo
	keyFrom.CopyBasicFields(keyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPie(mapOrigCopy map[any]any, pieFrom *Pie) (pieTo *Pie) {

	// pieFrom has already been copied
	if _pieTo, ok := mapOrigCopy[pieFrom]; ok {
		pieTo = _pieTo.(*Pie)
		return
	}

	pieTo = new(Pie)
	mapOrigCopy[pieFrom] = pieTo
	pieFrom.CopyBasicFields(pieTo)

	//insertion point for the staging of instances referenced by pointers
	if pieFrom.X != nil {
		pieTo.X = CopyBranchKey(mapOrigCopy, pieFrom.X)
	}
	if pieFrom.Y != nil {
		pieTo.Y = CopyBranchKey(mapOrigCopy, pieFrom.Y)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _serie := range pieFrom.Set {
		pieTo.Set = append(pieTo.Set, CopyBranchSerie(mapOrigCopy, _serie))
	}

	return
}

func CopyBranchScatter(mapOrigCopy map[any]any, scatterFrom *Scatter) (scatterTo *Scatter) {

	// scatterFrom has already been copied
	if _scatterTo, ok := mapOrigCopy[scatterFrom]; ok {
		scatterTo = _scatterTo.(*Scatter)
		return
	}

	scatterTo = new(Scatter)
	mapOrigCopy[scatterFrom] = scatterTo
	scatterFrom.CopyBasicFields(scatterTo)

	//insertion point for the staging of instances referenced by pointers
	if scatterFrom.X != nil {
		scatterTo.X = CopyBranchKey(mapOrigCopy, scatterFrom.X)
	}
	if scatterFrom.Y != nil {
		scatterTo.Y = CopyBranchKey(mapOrigCopy, scatterFrom.Y)
	}
	if scatterFrom.Text != nil {
		scatterTo.Text = CopyBranchKey(mapOrigCopy, scatterFrom.Text)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _serie := range scatterFrom.Set {
		scatterTo.Set = append(scatterTo.Set, CopyBranchSerie(mapOrigCopy, _serie))
	}

	return
}

func CopyBranchSerie(mapOrigCopy map[any]any, serieFrom *Serie) (serieTo *Serie) {

	// serieFrom has already been copied
	if _serieTo, ok := mapOrigCopy[serieFrom]; ok {
		serieTo = _serieTo.(*Serie)
		return
	}

	serieTo = new(Serie)
	mapOrigCopy[serieFrom] = serieTo
	serieFrom.CopyBasicFields(serieTo)

	//insertion point for the staging of instances referenced by pointers
	if serieFrom.Key != nil {
		serieTo.Key = CopyBranchKey(mapOrigCopy, serieFrom.Key)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _value := range serieFrom.Values {
		serieTo.Values = append(serieTo.Values, CopyBranchValue(mapOrigCopy, _value))
	}

	return
}

func CopyBranchValue(mapOrigCopy map[any]any, valueFrom *Value) (valueTo *Value) {

	// valueFrom has already been copied
	if _valueTo, ok := mapOrigCopy[valueFrom]; ok {
		valueTo = _valueTo.(*Value)
		return
	}

	valueTo = new(Value)
	mapOrigCopy[valueFrom] = valueTo
	valueFrom.CopyBasicFields(valueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Bar:
		stage.UnstageBranchBar(target)

	case *Key:
		stage.UnstageBranchKey(target)

	case *Pie:
		stage.UnstageBranchPie(target)

	case *Scatter:
		stage.UnstageBranchScatter(target)

	case *Serie:
		stage.UnstageBranchSerie(target)

	case *Value:
		stage.UnstageBranchValue(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchBar(bar *Bar) {

	// check if instance is already staged
	if !IsStaged(stage, bar) {
		return
	}

	bar.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bar.X != nil {
		UnstageBranch(stage, bar.X)
	}
	if bar.Y != nil {
		UnstageBranch(stage, bar.Y)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _serie := range bar.Set {
		UnstageBranch(stage, _serie)
	}

}

func (stage *StageStruct) UnstageBranchKey(key *Key) {

	// check if instance is already staged
	if !IsStaged(stage, key) {
		return
	}

	key.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPie(pie *Pie) {

	// check if instance is already staged
	if !IsStaged(stage, pie) {
		return
	}

	pie.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pie.X != nil {
		UnstageBranch(stage, pie.X)
	}
	if pie.Y != nil {
		UnstageBranch(stage, pie.Y)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _serie := range pie.Set {
		UnstageBranch(stage, _serie)
	}

}

func (stage *StageStruct) UnstageBranchScatter(scatter *Scatter) {

	// check if instance is already staged
	if !IsStaged(stage, scatter) {
		return
	}

	scatter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scatter.X != nil {
		UnstageBranch(stage, scatter.X)
	}
	if scatter.Y != nil {
		UnstageBranch(stage, scatter.Y)
	}
	if scatter.Text != nil {
		UnstageBranch(stage, scatter.Text)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _serie := range scatter.Set {
		UnstageBranch(stage, _serie)
	}

}

func (stage *StageStruct) UnstageBranchSerie(serie *Serie) {

	// check if instance is already staged
	if !IsStaged(stage, serie) {
		return
	}

	serie.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if serie.Key != nil {
		UnstageBranch(stage, serie.Key)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _value := range serie.Values {
		UnstageBranch(stage, _value)
	}

}

func (stage *StageStruct) UnstageBranchValue(value *Value) {

	// check if instance is already staged
	if !IsStaged(stage, value) {
		return
	}

	value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
