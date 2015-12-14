package cufinals

import (
	"fmt"
)

type Date struct {
	Month int
	Day   int
}

func (d Date) String() string {
	return fmt.Sprintf("%d/%02d", d.Month, d.Day)
}

type Time struct {
	Hour   int
	Minute int
	AM     bool
}

func (t Time) String() string {
	amStr := "PM"
	if t.AM {
		amStr = "AM"
	}
	return fmt.Sprintf("%d:%02d%s", t.Hour, t.Minute, amStr)
}

type Course struct {
	Department string
	Number     int
	Section    int
}

func (c Course) String() string {
	return fmt.Sprintf("%s %04d %03d", c.Department, c.Number, c.Section)
}

type Room struct {
	ShortName string
	LongName  string
}

func (r Room) String() string {
	return fmt.Sprintf("%s (%s)", r.ShortName, r.LongName)
}

type Entry struct {
	Course Course
	Date   Date
	Time   Time
	Room   Room
}

func (e Entry) String() string {
	return fmt.Sprintf("%s - %s %s - %s", e.Course, e.Date, e.Time, e.Room)
}
