package routes

import "time"

type Interval struct {
	Distance float64 `json:"distance"`
	Time     float64 `json:"time"`
	RestTime float64 `json:"restTime"`
	Pace     float64 `json:"pace"`
	Type     string  `json:"type"`
}

type Workout struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Athlete   string     `json:"athlete"`
	Distance  float64    `json:"distance"`
	Time      float64    `json:"time"`
	Date      time.Time  `json:"date"`
	Type      string     `json:"type"`
	Intervals []Interval `json:"intervals"`
}
