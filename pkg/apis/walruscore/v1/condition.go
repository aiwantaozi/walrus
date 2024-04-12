package v1

import (
	"reflect"
	"time"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

// ConditionStatus is the value of status.
//
// +enum
type ConditionStatus string

const (
	// ConditionTrue means a resource is in the condition.
	ConditionTrue ConditionStatus = "True"

	// ConditionFalse means a resource is not in the condition.
	ConditionFalse ConditionStatus = "False"

	// ConditionUnknown means a resource is in the condition or not.
	ConditionUnknown ConditionStatus = "Unknown"
)

// Condition describes the state of a condition at a certain point.
type Condition struct {
	// Type of condition name in CamelCase.
	Type ConditionType `json:"type"`

	// Status of the condition, one of True, False, Unknown.
	//
	// +k8s:validation:enum=["True","False","Unknown"]
	Status ConditionStatus `json:"status"`

	// LastTransitionTime is the last time the condition transitioned from one status to another.
	LastTransitionTime meta.Time `json:"lastTransitionTime"`

	// Reason contains a programmatic identifier indicating the reason for the condition's last transition.
	Reason string `json:"reason"`

	// Message is a human-readable message indicating details about the transition.
	Message string `json:"message"`
}

// ConditionType is the type of status.
type ConditionType string

// IsTrue check status value for object,
// object must be a pointer.
func (c ConditionType) IsTrue(obj any) bool {
	return getStatus(obj, string(c)) == string(ConditionTrue)
}

// IsFalse check status value for object,
// object must be a pointer.
func (c ConditionType) IsFalse(obj any) bool {
	return getStatus(obj, string(c)) == string(ConditionFalse)
}

// IsUnknown check status value for object,
// object must be a pointer.
func (c ConditionType) IsUnknown(obj any) bool {
	return getStatus(obj, string(c)) == string(ConditionUnknown)
}

// GetStatus get status from conditionType for object field .Status.Conditions.
func (c ConditionType) GetStatus(obj any) string {
	return getStatus(obj, string(c))
}

// GetMessage get message from conditionType for object field .Status.Conditions.
func (c ConditionType) GetMessage(obj any) string {
	cond := findCond(obj, string(c))
	if cond == nil {
		return ""
	}
	return getFieldValue(*cond, "Message").String()
}

// GetReason get reason from conditionType for object field .Status.Conditions.
func (c ConditionType) GetReason(obj any) string {
	cond := findCond(obj, string(c))
	if cond == nil {
		return ""
	}
	return getFieldValue(*cond, "Reason").String()
}

// GetLastTransitionTime get last transition time for conditionType from object field .Status.Conditions.
func (c ConditionType) GetLastTransitionTime(obj any) string {
	return getTS(obj, string(c))
}

// True set status value to True for object field .Status.Conditions,
// object must be a pointer.
func (c ConditionType) True(obj any, reason, message string) {
	setStatus(obj, string(c), string(meta.ConditionTrue), reason, message)
}

// False set status value to False for object field .Status.Conditions,
// object must be a pointer.
func (c ConditionType) False(obj any, reason, message string) {
	setStatus(obj, string(c), string(meta.ConditionFalse), reason, message)
}

// Unknown set status value to Unknown for object field .Status.Conditions,
// object must be a pointer.
func (c ConditionType) Unknown(obj any, reason, message string) {
	setStatus(obj, string(c), string(meta.ConditionUnknown), reason, message)
}

// Status set status value to custom value for object field .Status.Conditions,
// object must be a pointer.
func (c ConditionType) Status(obj any, status, reason, message string) {
	setStatus(obj, string(c), status, reason, message)
}

// Message set message to conditionType for object field .Status.Conditions,
// object must be a pointer.
func (c ConditionType) Message(obj any, message string) {
	cond := upsertCond(obj, string(c))
	setValue(cond, "Message", message)
}

// Reason set reason to conditionType for object field .Status.Conditions,
// object must be a pointer.
func (c ConditionType) Reason(obj any, reason string) {
	cond := upsertCond(obj, string(c))
	getFieldValue(cond, "Reason").SetString(reason)
}

// LastTransitionTime set last transition time to conditionType for object field .Status.Conditions,
// object must be a pointer.
func (c ConditionType) LastTransitionTime(obj any, ts string) {
	upsertTS(obj, string(c), ts)
}

// SetError set error to conditionType for object field .Status.Conditions,
// object must be a pointer.
func (c ConditionType) SetError(obj any, reason string, err error) {
	if reason == "" {
		reason = "Error"
	}
	c.False(obj, reason, err.Error())
}

// SetMessageIfBlank set message to conditionType for object field .Status.Conditions if message is blank,
// object must be a pointer.
func (c ConditionType) SetMessageIfBlank(obj any, message string) {
	if c.GetMessage(obj) == "" {
		c.Message(obj, message)
	}
}

// Reset clean the object field .Status.Conditions,
// and set the status as Unknown type into the object field .Status.Conditions,
// object must be a pointer.
func (c ConditionType) Reset(obj any, message string) {
	resetCond(obj, string(c), string(meta.ConditionUnknown), message)
}

// upsertTS create to update condition and set last transition time.
func upsertTS(obj any, condName, ts string) {
	cond := upsertCond(obj, condName)
	getFieldValue(cond, "LastTransitionTime").SetString(ts)
}

// setTS set last transition time to condition.
func setTS(value reflect.Value) {
	now := meta.Time{
		Time: time.Now().UTC(),
	}

	getFieldValue(value, "LastTransitionTime").Set(reflect.ValueOf(now))
}

// getTS get last transition time from condition.
func getTS(obj any, condName string) string {
	cond := findCond(obj, condName)
	if cond == nil {
		return ""
	}
	return getFieldValue(*cond, "LastTransitionTime").String()
}

// getStatus get status from condition.
func getStatus(obj any, condName string) string {
	cond := findCond(obj, condName)
	if cond == nil {
		return ""
	}
	return getFieldValue(*cond, "Status").String()
}

// setStatus set status and message to condition.
func setStatus(obj any, condName, status, reason, message string) {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic("obj passed must be a pointer")
	}

	cond := upsertCond(obj, condName)
	setValue(cond, "Status", status)
	setValue(cond, "Reason", reason)
	setValue(cond, "Message", message)

	originalStatus := getValue(cond, "Status").String()
	if originalStatus != status {
		setTS(cond)
	}
}

// getValue get value from object with field names.
func getValue(obj any, name ...string) reflect.Value {
	if obj == nil {
		return reflect.Value{}
	}
	v := reflect.ValueOf(obj)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	field := v.FieldByName(name[0])
	if len(name) == 1 {
		return field
	}
	return getFieldValue(field, name[1:]...)
}

// setValue set value to condition.
func setValue(cond reflect.Value, fieldName, newValue string) {
	value := getFieldValue(cond, fieldName)
	if value.String() != newValue {
		value.SetString(newValue)
	}
}

// findCond find condition from object.
func findCond(obj any, condName string) *reflect.Value {
	condSlice := getValue(obj, "Status", "Conditions")
	if !condSlice.IsValid() {
		condSlice = getValue(obj, "Conditions")
	}
	return queryCondsByName(obj, condSlice, condName)
}

// upsertCond create or update condition.
func upsertCond(obj any, condName string) reflect.Value {
	conds := getValue(obj, "Status", "Conditions")
	cond := queryCondsByName(obj, conds, condName)
	if cond != nil {
		return *cond
	}

	newCond := reflect.New(conds.Type().Elem()).Elem()
	newCond.FieldByName("Type").SetString(condName)
	newCond.FieldByName("Status").SetString("Unknown")
	setTS(newCond)

	conds.Set(reflect.Append(conds, newCond))
	return *queryCondsByName(obj, conds, condName)
}

// resetCond clean the object field .Status.Conditions, and set the status to Unknown.
func resetCond(obj any, condName, status, message string) {
	conds := getValue(obj, "Status", "Conditions")

	newCond := reflect.New(conds.Type().Elem()).Elem()
	newCond.FieldByName("Type").SetString(condName)
	newCond.FieldByName("Status").SetString(status)
	newCond.FieldByName("Message").SetString(message)
	setTS(newCond)

	slice := reflect.MakeSlice(reflect.SliceOf(newCond.Type()), 0, 0)
	conds.Set(reflect.Append(slice, newCond))
}

// queryCondsByName query condition by name.
func queryCondsByName(obj any, val reflect.Value, condName string) *reflect.Value {
	defer func() {
		if recover() != nil {
			klog.Fatalf("failed to find .Status.Conditions field on %v", reflect.TypeOf(obj))
		}
	}()

	for i := 0; i < val.Len(); i++ {
		cond := val.Index(i)
		typeVal := getFieldValue(cond, "Type")
		if typeVal.String() == condName {
			return &cond
		}
	}

	return nil
}

// getFieldValue get value from field names.
func getFieldValue(v reflect.Value, name ...string) reflect.Value {
	if !v.IsValid() {
		return v
	}
	field := v.FieldByName(name[0])
	if len(name) == 1 {
		return field
	}
	return getFieldValue(field, name[1:]...)
}
