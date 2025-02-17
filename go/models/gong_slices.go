// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Bar:
		// insertion point per field
		if fieldName == "Set" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Bar) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Bar)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Set).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Set = _inferedTypeInstance.Set[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Set =
								append(_inferedTypeInstance.Set, any(fieldInstance).(*Serie))
						}
					}
				}
			}
		}

	case *Key:
		// insertion point per field

	case *Pie:
		// insertion point per field
		if fieldName == "Set" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Pie) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Pie)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Set).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Set = _inferedTypeInstance.Set[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Set =
								append(_inferedTypeInstance.Set, any(fieldInstance).(*Serie))
						}
					}
				}
			}
		}

	case *Scatter:
		// insertion point per field
		if fieldName == "Set" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Scatter) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Scatter)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Set).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Set = _inferedTypeInstance.Set[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Set =
								append(_inferedTypeInstance.Set, any(fieldInstance).(*Serie))
						}
					}
				}
			}
		}

	case *Serie:
		// insertion point per field
		if fieldName == "Values" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Serie) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Serie)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Values).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Values = _inferedTypeInstance.Values[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Values =
								append(_inferedTypeInstance.Values, any(fieldInstance).(*Value))
						}
					}
				}
			}
		}

	case *Value:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Bar
	// insertion point per field
	clear(stage.Bar_Set_reverseMap)
	stage.Bar_Set_reverseMap = make(map[*Serie]*Bar)
	for bar := range stage.Bars {
		_ = bar
		for _, _serie := range bar.Set {
			stage.Bar_Set_reverseMap[_serie] = bar
		}
	}

	// Compute reverse map for named struct Key
	// insertion point per field

	// Compute reverse map for named struct Pie
	// insertion point per field
	clear(stage.Pie_Set_reverseMap)
	stage.Pie_Set_reverseMap = make(map[*Serie]*Pie)
	for pie := range stage.Pies {
		_ = pie
		for _, _serie := range pie.Set {
			stage.Pie_Set_reverseMap[_serie] = pie
		}
	}

	// Compute reverse map for named struct Scatter
	// insertion point per field
	clear(stage.Scatter_Set_reverseMap)
	stage.Scatter_Set_reverseMap = make(map[*Serie]*Scatter)
	for scatter := range stage.Scatters {
		_ = scatter
		for _, _serie := range scatter.Set {
			stage.Scatter_Set_reverseMap[_serie] = scatter
		}
	}

	// Compute reverse map for named struct Serie
	// insertion point per field
	clear(stage.Serie_Values_reverseMap)
	stage.Serie_Values_reverseMap = make(map[*Value]*Serie)
	for serie := range stage.Series {
		_ = serie
		for _, _value := range serie.Values {
			stage.Serie_Values_reverseMap[_value] = serie
		}
	}

	// Compute reverse map for named struct Value
	// insertion point per field

}
