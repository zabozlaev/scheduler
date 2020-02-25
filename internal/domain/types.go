package domain


type Meta struct {
	Title string `json:"title"`
	Building string `json:"building"`
	ClassRoom string `json:"room"`
	Lecturer string `json:"lecturer"`
}

type Class struct {
	Time string `json:"time"`
	ClassMeta []*Meta `json:"meta"`
}

type Schedule struct {
	Mon []*Class `json:"monday"`
	Tue []*Class `json:"tuesday"`
	Wed []*Class `json:"wednesday"`
	Thu []*Class `json:"thursday"`
	Fri []*Class `json:"friday"`
	Sat []*Class `json:"saturday"`
}

type ScheduleResponse struct {
	Sch *Schedule `json:"schedule"`
	Group string `json:"group"`
}

type ScheduleDayResponse struct {
	Group string `json:"group"`
	Day int `json:"day"`
	Classes []*Class `json:"classes"`
}
