package v1

import (
	"reflect"
	"time"
)

// ConditionStatus is the value of status.
type ConditionStatus string

// These are valid condition statuses.
// "ConditionStatusTrue" means a resource is in the condition.
// "ConditionStatusFalse" means a resource is not in the condition.
// "ConditionStatusUnknown" means a resource is in the condition or not.
const (
	ConditionStatusTrue    ConditionStatus = "True"
	ConditionStatusFalse   ConditionStatus = "False"
	ConditionStatusUnknown ConditionStatus = "Unknown"
)

// ConditionType is the type of status.
type ConditionType string

func (ct ConditionType) String() string {
	return string(ct)
}

// True set status value to True for object field .Status.Conditions,
// object must be a pointer.
func (ct ConditionType) True(obj any, message string) {
	setCondStatusAndMessage(obj, ct, ConditionStatusTrue, message)
}

// False set status value to False for object field .Status.Conditions,
// object must be a pointer.
func (ct ConditionType) False(obj any, message string) {
	setCondStatusAndMessage(obj, ct, ConditionStatusFalse, message)
}

// Unknown set status value to Unknown for object field .Status.Conditions,
// object must be a pointer.
func (ct ConditionType) Unknown(obj any, message string) {
	setCondStatusAndMessage(obj, ct, ConditionStatusUnknown, message)
}

// Status set status value to custom value for object field .Status.Conditions,
// object must be a pointer.
func (ct ConditionType) Status(obj any, status ConditionStatus) {
	setCondStatusAndMessage(obj, ct, status, "")
}

// Message set message to conditionType for object field .Status.Conditions,
// object must be a pointer.
func (ct ConditionType) Message(obj any, message string) {
	setCondMessage(obj, ct, message)
}

// IsTrue check status value for object,
// object must be a pointer.
func (ct ConditionType) IsTrue(obj any) bool {
	s, _ := getCondStatus(obj, ct)
	return s == ConditionStatusTrue
}

// IsFalse check status value for object,
// object must be a pointer.
func (ct ConditionType) IsFalse(obj any) bool {
	s, _ := getCondStatus(obj, ct)
	return s == ConditionStatusFalse
}

// IsUnknown check status value for object,
// object must be a pointer.
func (ct ConditionType) IsUnknown(obj any) bool {
	s, _ := getCondStatus(obj, ct)
	return s == ConditionStatusUnknown
}

// Exist returns true if the status is existed,
// object must be a pointer.
func (ct ConditionType) Exist(obj any) bool {
	_, exist := getCondStatus(obj, ct)
	return exist
}

// GetMessage get message from conditionType for object field .Status.Conditions.
func (ct ConditionType) GetMessage(obj any) string {
	return getCondMessage(obj, ct)
}

func setCondMessage(obj any, condType ConditionType, message string) {
	if obj == nil || reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return
	}

	conds := getConditions(obj)

	// Condition existed.
	cond := findCond(conds, condType.String())
	if cond != nil {
		setValue(*cond, "Message", message)
		return
	}

	// Condition not existed.
	newCond := reflect.New(conds.Type().Elem()).Elem()
	newCond.FieldByName("Type").SetString(condType.String())
	newCond.FieldByName("Message").SetString(string(message))

	conds.Set(reflect.Append(conds, newCond))
	cond = findCond(conds, condType.String())
}

func setCondStatusAndMessage(obj any, condType ConditionType, status ConditionStatus, message string) {
	if obj == nil || reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return
	}

	conds := getConditions(obj)

	// Condition existed.
	cond := findCond(conds, condType.String())
	if cond != nil {
		setValue(*cond, "Status", string(status))
		setValue(*cond, "Message", message)
		return
	}

	// Condition not existed.
	newCond := reflect.New(conds.Type().Elem()).Elem()
	newCond.FieldByName("Type").SetString(condType.String())
	newCond.FieldByName("Status").SetString(string(status))
	newCond.FieldByName("Message").SetString(string(message))

	conds.Set(reflect.Append(conds, newCond))
	cond = findCond(conds, condType.String())
}

func setValue(cond reflect.Value, fieldName, newValue string) {
	value := getFieldValue(cond, fieldName)
	if value.String() != newValue {
		value.SetString(newValue)
		touchTS(cond)
	}
}

func touchTS(value reflect.Value) {
	now := time.Now().Format(time.RFC3339)
	getFieldValue(value, "LastUpdateTime").SetString(now)
}

func getCondStatus(obj any, condType ConditionType) (ConditionStatus, bool) {
	if obj == nil || reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return "", false
	}

	conds := getConditions(obj)
	cond := findCond(conds, condType.String())
	if cond == nil {
		return "", false
	}

	val := getValue(cond, "Status")

	return ConditionStatus(val.String()), true
}

func getCondMessage(obj any, condType ConditionType) string {
	if obj == nil || reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return ""
	}

	conds := getConditions(obj)
	cond := findCond(conds, condType.String())
	if cond == nil {
		return ""
	}

	return getFieldValue(*cond, "Message").String()
}

func getConditions(obj any) reflect.Value {
	v := reflect.ValueOf(obj)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return reflect.Value{}
	}

	val := getValue("Status", "conditions")
	return val
}

func findCond(val reflect.Value, name string) *reflect.Value {
	for i := 0; i < val.Len(); i++ {
		cond := val.Index(i)
		typeVal := getFieldValue(cond, "Type")
		if typeVal.String() == name {
			return &cond
		}
	}

	return nil
}

func getValue(obj interface{}, name ...string) reflect.Value {
	if obj == nil {
		return reflect.Value{}
	}
	v := reflect.ValueOf(obj)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return getFieldValue(v, name...)
}

func getFieldValue(v reflect.Value, name ...string) reflect.Value {
	field := v.FieldByName(name[0])
	if len(name) == 1 {
		return field
	}
	return getFieldValue(field, name[1:]...)
}
