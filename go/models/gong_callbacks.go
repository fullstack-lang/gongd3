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
	case *Serie:
		if stage.OnAfterSerieCreateCallback != nil {
			stage.OnAfterSerieCreateCallback.OnAfterCreate(stage, target)
		}
	case *Value:
		if stage.OnAfterValueCreateCallback != nil {
			stage.OnAfterValueCreateCallback.OnAfterCreate(stage, target)
		}
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
	case *Serie:
		if stage.OnAfterSerieReadCallback != nil {
			stage.OnAfterSerieReadCallback.OnAfterRead(stage, target)
		}
	case *Value:
		if stage.OnAfterValueReadCallback != nil {
			stage.OnAfterValueReadCallback.OnAfterRead(stage, target)
		}
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
	
	case *Serie:
		stage.OnAfterSerieReadCallback = any(callback).(OnAfterReadInterface[Serie])
	
	case *Value:
		stage.OnAfterValueReadCallback = any(callback).(OnAfterReadInterface[Value])
	
	}
}
