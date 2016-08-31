package planout

import "testing"

func TestUtilities(t *testing.T) {
	testObj := make(map[string]interface{})

	testObj["first"] = make(map[string]interface{})

	testObj["first"].(map[string]interface{})["second"] = make(map[string]interface{})

	testObj["first"].(map[string]interface{})["second"].(map[string]interface{})["third"] = 1

	result, ok := delve(testObj, "first.second.third")
	if compare(result, 1) != 0 {
		t.Errorf("Variable 'result'. Expected value 1. Actual %v\n", result)
	}

	resultNil, okNil := delve(testObj, "first.second.nothing")
	if okNil {
		t.Errorf("Variable 'resultNil'. Expected to be nil. Actual %v\n", resultNil)
	}

	resultReallyNil, okReallyNil := delve(testObj, "first.do.nothing")
	if okReallyNil {
		t.Errorf("Variable 'resultReallyNil'. Expected to be nil. Actual %v\n", resultReallyNil)
	}

	// Craeat

	ok = delveCreate(testObj, "first.do.no", "harm")
	if !ok {
		t.Error("Expected delveCreate to not fail")
	}

	result, ok = delve(testObj, "first.do.no")
	if compare(result, "harm") != 0 {
		t.Errorf("Variable 'result'. Expected value \"harm\". Actual %v\n", result)
	}

	ok = delveCreate(testObj, "something.else.is", "here")
	if !ok {
		t.Error("Expected delveCreate to not fail")
	}

	result, ok = delve(testObj, "something.else.is")
	if compare(result, "here") != 0 {
		t.Errorf("Variable 'result'. Expected value \"here\". Actual %v\n", result)
	}
}
