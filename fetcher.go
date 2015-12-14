package cufinals

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

var DefaultScheduleURL = "https://registrar.cornell.edu/Sched/EXFA.html"

func FetchSchedule(url string) ([]Entry, error) {
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	expr1 := regexp.MustCompile("<pre>((.|\\n|\\r)*?)</pre>")
	preformatted := expr1.FindAllStringSubmatch(string(body), 1)
	if len(preformatted) != 1 {
		return nil, errors.New("no <pre> tag")
	}

	expr2 := regexp.MustCompile("(?m)^([A-Z]*) ([0-9]*) ([0-9]*)\\s*" +
		"(Mon|Wed|Tue|Thu|Fri|Sat|Sun), (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) " +
		"([0-9]*)\\s*([0-9]*):([0-9]*) (AM|PM)\\s*" +
		"(.*?): (.*?)\\r?$")
	matches := expr2.FindAllStringSubmatch(preformatted[0][1], -1)
	res := make([]Entry, 0, len(matches))
	for _, match := range matches {
		var e Entry
		e.Course.Department = match[1]

		numStrs := []string{match[2], match[3], match[6], match[7], match[8]}
		nums := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			nums[i], err = strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
		}
		e.Course.Number = nums[0]
		e.Course.Section = nums[1]

		monthNames := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
			"Oct", "Nov", "Dec"}
		for i, x := range monthNames {
			if x == match[5] {
				e.Date.Month = i + 1
				break
			}
		}

		e.Date.Day = nums[2]
		e.Time.Hour = nums[3]
		e.Time.Minute = nums[4]
		e.Time.AM = (match[9] == "AM")

		e.Room.ShortName = match[10]
		e.Room.LongName = match[11]

		res = append(res, e)
	}
	return res, nil
}
