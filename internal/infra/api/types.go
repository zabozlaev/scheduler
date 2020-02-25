package api

import "github.com/graphql-go/graphql"

var MetaType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Meta",
		Description: "Meta of the university class",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
				Description: "Title of the class",
			},
			"building": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Description: "Building the class room is in",
			},
			"room": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Description: "Room the class is given",
			},
			"lecturer": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Description: "Lecturer that will be on the class",
			},
		},
	},
)

var ClassType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Class",
		Fields:      graphql.Fields{
			"time": &graphql.Field{
				Type:              graphql.NewNonNull(graphql.String),
				Description:       "Time the class starts and ends",
			},
			"meta": &graphql.Field{
				Type:              graphql.NewList(graphql.NewNonNull(MetaType)),
				Description:       "Meta list of the class",
			},
		},
		Description: "University class",
	})

var DayEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Day",
	Description: "University work days",
	Values: graphql.EnumValueConfigMap{
		"MON": &graphql.EnumValueConfig{
			Value:             0,
			Description:       "Monday",
		},
		"TUE": &graphql.EnumValueConfig{
			Value:             1,
			Description:       "Tuesday",
		},
		"WED": &graphql.EnumValueConfig{
			Value:             2,
			Description:       "Wednesday",
		},
		"THU": &graphql.EnumValueConfig{
			Value:             3,
			Description:       "Thursday",
		},
		"FRI": &graphql.EnumValueConfig{
			Value:             4,
			Description:       "Friday",
		},
		"SAT": &graphql.EnumValueConfig{
			Value:             5,
			Description:       "Saturday",
		},
	},
})

var ScheduleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Schedule",
		Fields:      graphql.Fields{
			"monday": &graphql.Field{
				Type:              graphql.NewList(ClassType),
				Description:       "Classes for monday",
			},
			"tuesday": &graphql.Field{
				Type:              graphql.NewList(ClassType),
				Description:       "Classes for tuesday",
			},
			"wednesday": &graphql.Field{
				Type:              graphql.NewList(ClassType),
				Description:       "Classes for wednesday",
			},
			"thursday": &graphql.Field{
				Type:              graphql.NewList(ClassType),
				Description:       "Classes for thursday",
			},
			"friday": &graphql.Field{
				Type:              graphql.NewList(ClassType),
				Description:       "Classes for friday",
			},
			"saturday": &graphql.Field{
				Type:              graphql.NewList(ClassType),
				Description:       "Classes for saturday",
			},
		},
		Description: "Weekly schedule",
	})

