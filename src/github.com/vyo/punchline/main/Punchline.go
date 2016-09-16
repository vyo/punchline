package main

import (
	"fmt"
	"time"
	"github.com/vyo/punch/model"
	"gopkg.in/alecthomas/kingpin.v2"
	"encoding/json"
	"github.com/vyo/punch/util"
)

const ADD = "add"
const ENTRY = "entry"
const ADD_ENTRY = ADD + " " + ENTRY

func main() {

	entry1 := model.AtomicRecord{Start: time.Now().Add(time.Duration(-60 * time.Minute)), Key: "oauth"}
	entry2 := model.AtomicRecord{Start: time.Now().Add(time.Duration(-50 * time.Minute)), Key: "paket"}
	entry3 := model.AtomicRecord{Start: time.Now().Add(time.Duration(-40 * time.Minute)), Key: "paka"}
	entry4 := model.AtomicRecord{Start: time.Now().Add(time.Duration(-30 * time.Minute)), Key: "oauth"}
	entry5 := model.AtomicRecord{Start: time.Now().Add(time.Duration(-20 * time.Minute)), Key: "bonus"}
	entry6 := model.AtomicRecord{Start: time.Now().Add(time.Duration(-10 * time.Minute)), Key: "oauth"}
	entry7 := model.AtomicRecord{Start: time.Now().Add(time.Duration(-10 * time.Minute)), Key: "end"}

	entries := []model.AtomicRecord{entry1, entry2, entry3, entry4, entry5, entry6, entry7}
	fmt.Println("raw entries:")
	for _, entry := range entries {
		fmt.Printf("%+v\t%+v\n", entry.Key, entry.Start)
		fmt.Printf("%+v\t%+v\n\n", entry.Key, entry.End)
	}

	fmt.Println("sorted/inferred entries:")
	util.InferEndTimes(&entries)
	for _, entry := range entries {
		fmt.Printf("%+v\t%+v\n", entry.Key, entry.Start)
		fmt.Printf("%+v\t%+v\n\n", entry.Key, entry.End)
	}

	dailies := util.AtomicsToDailies(entries)
	fmt.Println("dailies:")
	for _, entry := range dailies {
		fmt.Printf("%+v\t%+v\n", entry.Key, entry.Start.Format(time.Kitchen))
		fmt.Printf("%+v\t%+v\n", entry.Key, entry.End.Format(time.Kitchen))
		fmt.Printf("%+v\t%+vm total\t\n", entry.Key, int(entry.End.Sub(entry.Start).Minutes()))
		fmt.Printf("%+v\t%+vm break\t\n\n", entry.Key, int(entry.OffTime.Minutes()))
	}
	var config = model.NewConfig(nil, 0)

	*config.GetChunkSize() = 5

	var (
		add = kingpin.Command(ADD, "add a new punch, project w/e").Default()
		addEntry = add.Command(ENTRY, "add a new punchcard entry").Default()
		addPunchKey = addEntry.Arg("key", "the key of the project you want to punch time for goes here").Required().String()
		//fuzzy = addPunch.Flag("fuzzy", "rounds all times punched up/down to this minimum timespan " +
		//	"measured in minutes; " +
		//	"tries to minimize overall deviation from exact time punched").Short('f').Default(strconv.Itoa(config.ChunkSize)).Int()
	)

	kingpin.CommandLine.Help = "A simple time tracking tool"
	kingpin.UsageTemplate(kingpin.DefaultUsageTemplate).Version("0.1").Author("Manuel Weidmann")
	kingpin.CommandLine.HelpFlag.Short('h')

	switch kingpin.Parse() {
	case ADD_ENTRY:
		entry := model.AtomicRecord{Start: time.Now(), End: time.Now(), Key: *addPunchKey, Description: "description"}
		content, err := json.Marshal(entry)

		if (err == nil) {
			fmt.Printf("json entry %s\n", content)
			fmt.Printf("entry      %+v\n", entry)
			//fmt.Printf("fuzzy      %v\n", *fuzzy)
		}
	}
}
