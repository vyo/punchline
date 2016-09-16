package model

import "time"

type AtomicRecord struct {
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"` //optional, will be inferred when constructing daily records
	Key         string `json:"key"`
	Description string `json:"description"`
}

//func NewAtomicRecord(start time.Time, endPtr *time.Time, key string, descriptionPtr *string) *AtomicRecord {
//	var end time.Time
//	if (endPtr == nil) {
//
//	}
//
//	var description string
//	if ( descriptionPtr == nil) {
//		description = ""
//	} else {
//		description = *descriptionPtr
//	}
//
//	return &AtomicRecord{start, end, key, description}
//}

//func (r *AtomicRecord) GetStart() *time.Time {
//	return &r.Start
//}
//
//func (r *AtomicRecord) GetEnd() *time.Time {
//	return &r.End
//}
//
//func (r *AtomicRecord) GetKey() *string {
//	return &r.Key
//}
//
//func (r *AtomicRecord) GetDescription() *string {
//	return &r.Description
//}

type DailyRecord struct {
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"` //optional, will be inferred when constructing daily records
	Key         string `json:"key"`
	Description string `json:"description"`
	OffTime     time.Duration `json:"off_time"`
}

//func NewDailyRecord(Start time.Time, End time.Time, Key string, Description *string, OffTime *time.Duration) *DailyRecord {
//	return &DailyRecord{Start, End, Key, Description, OffTime}
//}

//
//func (r *DailyRecord) GetStart() *time.Time {
//	return &r.Start
//}
//
//func (r *DailyRecord) GetEnd() *time.Time {
//	return &r.End
//}
//
//func (r *DailyRecord) GetKey() *string {
//	return &r.Key
//}
//
//func (r *DailyRecord) GetDescription() *string {
//	return &r.Description
//}
//
//func (r *DailyRecord) GetOffTime() *time.Duration {
//	return &r.OffTime
//}
