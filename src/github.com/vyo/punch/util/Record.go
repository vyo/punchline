package util

import (
	"github.com/vyo/punch/model"
	"sort"
	"time"
)

type ByStartTime []model.AtomicRecord

func (t ByStartTime) Len() int {
	return len(t)
}
func (t ByStartTime) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t ByStartTime) Less(i, j int) bool {
	return t[i].Start.Before(t[j].Start)
}

func GetUniqueKeys(atomics []model.AtomicRecord) []string {
	keys := make([]string, 0)

	for _, atomic := range atomics {
		foundAlready := false
		for _, key := range keys {
			if (atomic.Key == key) {
				foundAlready = true
				break
			}
		}
		if (!foundAlready) {
			keys = append(keys, atomic.Key)
		}
	}

	return keys
}

func FilterByKey(atomics []model.AtomicRecord, key string) []model.AtomicRecord {
	filteredAtomics := make([]model.AtomicRecord, 0)

	for _, atomic := range atomics {
		if (atomic.Key == key) {
			filteredAtomics = append(filteredAtomics, atomic)
		}
	}

	return filteredAtomics
}

func atomicsToDaily(atomics []model.AtomicRecord) model.DailyRecord {
	sort.Sort(ByStartTime(atomics))

	offTimeNanos := 0

	daily := model.DailyRecord{}
	if (len(atomics) == 0) {
		return model.DailyRecord{}
	}
	key := atomics[0].Key
	daily.Key = key
	daily.Start = atomics[0].Start

	for index, atomic := range atomics {
		if (atomic.Key != key) {
			panic("non-uniform atomic records collection: expected %v but got %v")//, key, atomic.Key)
		}
		daily.End = atomic.End
		if (index > 0) {
			gapTime := atomics[index].Start.Sub(atomics[index - 1].End)
			if (gapTime.Hours() > 0 || gapTime.Minutes() > 0) {
				offTimeNanos = offTimeNanos + int(gapTime.Hours()) * 1000000000 * 3600
				offTimeNanos = offTimeNanos + int(gapTime.Minutes()) * 1000000000 * 60
			}
		}
	}

	daily.OffTime = time.Duration(offTimeNanos)

	return daily
}

func AtomicsToDailies(atomics []model.AtomicRecord) []model.DailyRecord {
	dailies := make([]model.DailyRecord, 0)

	for _, key := range GetUniqueKeys(atomics) {
		dailies = append(dailies, atomicsToDaily(FilterByKey(atomics, key)))
	}

	return dailies
}

func InferEndTimes(atomics *[]model.AtomicRecord) {
	sort.Sort(ByStartTime(*atomics))
	for index, atomic := range *atomics {
		if (index == len(*atomics) - 1) {
			(*atomics)[index].End = atomic.Start
			(*atomics)[index - 1].End = atomic.Start
		} else if (index > 0) {
			(*atomics)[index - 1].End = atomic.Start //(*atomics)[index].Start
		}
	}
}