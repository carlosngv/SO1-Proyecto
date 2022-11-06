package models

type Response struct {
	Team1	string			`json:"team1" bson:"team1"`
	Team2	string			`json:"team2" bson:"team2"`
	Score	string			`json:"score" bson:"score"`
	Phase	string			`json:"phase" bson:"phase"`
}

type Responses struct {
	Responses []Response `json:"data" bson:"data"`
}
