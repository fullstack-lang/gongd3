// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Bar:
		if stage.OnAfterBarCreateCallback != nil {
			stage.OnAfterBarCreateCallback.OnAfterCreate(stage, target)
		}
	case *Key:
		if stage.OnAfterKeyCreateCallback != nil {
			stage.OnAfterKeyCreateCallback.OnAfterCreate(stage, target)
		}
	case *Pie:
		if stage.OnAfterPieCreateCallback != nil {
			stage.OnAfterPieCreateCallback.OnAfterCreate(stage, target)
		}
	case *Scatter:
		if stage.OnAfterScatterCreateCallback != nil {
			stage.OnAfterScatterCreateCallback.OnAfterCreate(stage, target)
		}
	case *Serie:
		if stage.OnAfterSerieCreateCallback != nil {
			stage.OnAfterSerieCreateCallback.OnAfterCreate(stage, target)
		}
	case *Value:
		if stage.OnAfterValueCreateCallback != nil {
			stage.OnAfterValueCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Bar:
		newTarget := any(new).(*Bar)
		if stage.OnAfterBarUpdateCallback != nil {
			stage.OnAfterBarUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Key:
		newTarget := any(new).(*Key)
		if stage.OnAfterKeyUpdateCallback != nil {
			stage.OnAfterKeyUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Pie:
		newTarget := any(new).(*Pie)
		if stage.OnAfterPieUpdateCallback != nil {
			stage.OnAfterPieUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Scatter:
		newTarget := any(new).(*Scatter)
		if stage.OnAfterScatterUpdateCallback != nil {
			stage.OnAfterScatterUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Serie:
		newTarget := any(new).(*Serie)
		if stage.OnAfterSerieUpdateCallback != nil {
			stage.OnAfterSerieUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Value:
		newTarget := any(new).(*Value)
		if stage.OnAfterValueUpdateCallback != nil {
			stage.OnAfterValueUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Bar:
		if stage.OnAfterBarDeleteCallback != nil {
			staged := any(staged).(*Bar)
			stage.OnAfterBarDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Key:
		if stage.OnAfterKeyDeleteCallback != nil {
			staged := any(staged).(*Key)
			stage.OnAfterKeyDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Pie:
		if stage.OnAfterPieDeleteCallback != nil {
			staged := any(staged).(*Pie)
			stage.OnAfterPieDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Scatter:
		if stage.OnAfterScatterDeleteCallback != nil {
			staged := any(staged).(*Scatter)
			stage.OnAfterScatterDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Serie:
		if stage.OnAfterSerieDeleteCallback != nil {
			staged := any(staged).(*Serie)
			stage.OnAfterSerieDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Value:
		if stage.OnAfterValueDeleteCallback != nil {
			staged := any(staged).(*Value)
			stage.OnAfterValueDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Bar:
		if stage.OnAfterBarReadCallback != nil {
			stage.OnAfterBarReadCallback.OnAfterRead(stage, target)
		}
	case *Key:
		if stage.OnAfterKeyReadCallback != nil {
			stage.OnAfterKeyReadCallback.OnAfterRead(stage, target)
		}
	case *Pie:
		if stage.OnAfterPieReadCallback != nil {
			stage.OnAfterPieReadCallback.OnAfterRead(stage, target)
		}
	case *Scatter:
		if stage.OnAfterScatterReadCallback != nil {
			stage.OnAfterScatterReadCallback.OnAfterRead(stage, target)
		}
	case *Serie:
		if stage.OnAfterSerieReadCallback != nil {
			stage.OnAfterSerieReadCallback.OnAfterRead(stage, target)
		}
	case *Value:
		if stage.OnAfterValueReadCallback != nil {
			stage.OnAfterValueReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bar:
		stage.OnAfterBarUpdateCallback = any(callback).(OnAfterUpdateInterface[Bar])
	
	case *Key:
		stage.OnAfterKeyUpdateCallback = any(callback).(OnAfterUpdateInterface[Key])
	
	case *Pie:
		stage.OnAfterPieUpdateCallback = any(callback).(OnAfterUpdateInterface[Pie])
	
	case *Scatter:
		stage.OnAfterScatterUpdateCallback = any(callback).(OnAfterUpdateInterface[Scatter])
	
	case *Serie:
		stage.OnAfterSerieUpdateCallback = any(callback).(OnAfterUpdateInterface[Serie])
	
	case *Value:
		stage.OnAfterValueUpdateCallback = any(callback).(OnAfterUpdateInterface[Value])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bar:
		stage.OnAfterBarCreateCallback = any(callback).(OnAfterCreateInterface[Bar])
	
	case *Key:
		stage.OnAfterKeyCreateCallback = any(callback).(OnAfterCreateInterface[Key])
	
	case *Pie:
		stage.OnAfterPieCreateCallback = any(callback).(OnAfterCreateInterface[Pie])
	
	case *Scatter:
		stage.OnAfterScatterCreateCallback = any(callback).(OnAfterCreateInterface[Scatter])
	
	case *Serie:
		stage.OnAfterSerieCreateCallback = any(callback).(OnAfterCreateInterface[Serie])
	
	case *Value:
		stage.OnAfterValueCreateCallback = any(callback).(OnAfterCreateInterface[Value])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bar:
		stage.OnAfterBarDeleteCallback = any(callback).(OnAfterDeleteInterface[Bar])
	
	case *Key:
		stage.OnAfterKeyDeleteCallback = any(callback).(OnAfterDeleteInterface[Key])
	
	case *Pie:
		stage.OnAfterPieDeleteCallback = any(callback).(OnAfterDeleteInterface[Pie])
	
	case *Scatter:
		stage.OnAfterScatterDeleteCallback = any(callback).(OnAfterDeleteInterface[Scatter])
	
	case *Serie:
		stage.OnAfterSerieDeleteCallback = any(callback).(OnAfterDeleteInterface[Serie])
	
	case *Value:
		stage.OnAfterValueDeleteCallback = any(callback).(OnAfterDeleteInterface[Value])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bar:
		stage.OnAfterBarReadCallback = any(callback).(OnAfterReadInterface[Bar])
	
	case *Key:
		stage.OnAfterKeyReadCallback = any(callback).(OnAfterReadInterface[Key])
	
	case *Pie:
		stage.OnAfterPieReadCallback = any(callback).(OnAfterReadInterface[Pie])
	
	case *Scatter:
		stage.OnAfterScatterReadCallback = any(callback).(OnAfterReadInterface[Scatter])
	
	case *Serie:
		stage.OnAfterSerieReadCallback = any(callback).(OnAfterReadInterface[Serie])
	
	case *Value:
		stage.OnAfterValueReadCallback = any(callback).(OnAfterReadInterface[Value])
	
	}
}
