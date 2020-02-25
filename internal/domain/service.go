package domain

import (
"github.com/PuerkitoBio/goquery"
"go.uber.org/zap"
"golang.org/x/net/html/charset"
"net/http"
"regexp"
"strings"
)

var chatStorage = make(map[int64]string)

type Service interface {
	GetScheduleForGroup(group string) (*Schedule, error)
	GetDayScheduleForGroup(group string, day int) ([]*Class, error)
	GetGroupByChatId(id int64) *string
	SetGroupByChatId(id int64, group string)
}

type service struct {
	logger       *zap.Logger
}

func (s *service) GetScheduleForGroup(group string) (*Schedule, error) {
	data, err := fetchSchedule(group)
	if err != nil {
		return nil, err
	}

	schedule := new (Schedule)

	classData := *data

	schedule.Mon = classData[0]
	schedule.Tue = classData[1]
	schedule.Wed = classData[2]
	schedule.Thu = classData[3]
	schedule.Fri = classData[4]
	schedule.Sat = classData[5]

	return schedule, nil
}

func (s *service) GetGroupByChatId(id int64) *string {
	group, ok := chatStorage[id]
	if !ok {
		return nil
	}

	return &group
}

func (s *service) SetGroupByChatId(id int64, group string) {
	chatStorage[id] = group
}

func (s *service) GetDayScheduleForGroup(group string, day int) ([]*Class, error) {
	data, err := fetchSchedule(group)
	if err != nil {
		return nil, err
	}

	val := (*data)[day]


	return val, nil
}

type Selected struct {
	WithFontTag bool
	Text string
	Html string
}

func fetchSchedule(group string) (*[][]*Class, error) {
	req, err := http.NewRequest(http.MethodGet, "https://kpfu.ru/week_sheadule_print?p_group_name=" + group, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	decoder, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}


	doc, err := goquery.NewDocumentFromReader(decoder)
	if err != nil {
		return nil, err
	}


	var root [][]Selected
	days := make(map[int]string)
	doc.Find("table").Find("tr").Each(func(i int, selection *goquery.Selection) {
		root = append(root, make([]Selected, 0))

		selection.Find("td").Each(func(j int, selection *goquery.Selection) {
			text := selection.Text()
			html, _ := selection.Html()
			if i == 0 && len(strings.TrimSpace(text)) != 0 {
				days[j] = text
			}

			selected := Selected{
				WithFontTag: strings.Contains(html, "font"),
				Text:        text,
				Html: html,
			}

			root[i] = append(root[i], selected)
		})
	})


	result := make([][]*Class, 7)
	for key, _ := range days {
		classes := make([]*Class, 0)

		for _, data := range root[1:] {
			if len(strings.TrimSpace(data[key].Text)) != 0 {
				str := data[key]
				class := &Class{
					Time:  data[0].Text,
				}

				txt, _ := goquery.NewDocumentFromReader(strings.NewReader(str.Text))
				splitter := strings.Split(txt.Text(), ", ")


				if !str.WithFontTag {

					var re = regexp.MustCompile(`\d+`)

					room := re.FindAllString(txt.Text(), -1)[0]
					meta := &Meta{}
					meta.Building = strings.TrimSpace(splitter[1])
					meta.Lecturer = strings.TrimSpace(splitter[2])
					meta.ClassRoom = room
					meta.Title = strings.Split(str.Text, meta.ClassRoom)[0]
					classMeta := make([]*Meta, 0)
					classMeta = append(classMeta, meta)
					class.ClassMeta = classMeta
				} else {
					title := strings.Split(str.Html, "<br/>")
					classMeta := make([]*Meta, 0)

					for idx, l := range title {
						if strings.Contains(l, "font") {
							el, _ := goquery.NewDocumentFromReader(strings.NewReader(l))

							text := el.Text()
							meta := Meta{}
							splt := strings.Split(text, ", ")
							meta.Title = title[idx - 1]

							if len(splt) == 3 {
								meta.ClassRoom = strings.TrimSpace(splt[0])
								meta.Building = strings.TrimSpace(splt[1])
								meta.Lecturer = strings.TrimSpace(splt[2])
								classMeta = append(classMeta, &meta)
							}
						}

					}
					class.ClassMeta = classMeta
				}

				classes = append(classes, class)
			}
		}

		result[key - 1] = classes
	}

	return &result, nil
}

func NewService(logger *zap.Logger) Service {
	return &service{logger:logger}
}


