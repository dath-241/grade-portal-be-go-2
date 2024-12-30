package controller_client


type InterfaceAccount struct {
	IDToken string `json:"idToken"`
}

type InterfaceScoreController struct {
	BT  []float32 `json:"BT"`
	TN  []float32 `json:"TN"`
	BTL []float32 `json:"BTL"`
	GK  float32   `json:"GK"`
	CK  float32   `json:"CK"`
}

type InterfaceResultScoreController struct {
	SCORE []struct {
		MSSV string                   `json:"MMSV"`
		Data InterfaceScoreController `json:"Data"`
	} `json:"score"`
	ClassID string `json:"class_id"`
}

type InterfaceHallOfFame struct {
	Semester string          `json:"semester"`
	Tier     []InterfaceTier `json:"tier"`
}

type InterfaceTier struct {
	CourseID any                    `json:"course_id" bson:"course_id"`
	Data     []InterfaceStudentData `json:"data"`
}

type InterfaceStudentData struct {
	MSSV string  `json:"mssv"`
	DTB  float32 `json:"dtb"`
}

type LoginInterface struct {
	Ms       string `json:"ms"`
	Password string `json:"password"`
}
type Class4Teacher struct {
	Semester      string   `json:"semester,omitempty"`
	Name          string   `json:"name,omitempty"`
	CourseId      string   `json:"course_id,omitempty"`
	ListStudentMs []string `json:"listStudent_ms,omitempty"`
	CsvURL        string   `json:"csv_url,omitempty"`
	LastModified  string   `json:"last_mod,omitempty"`
}
type StudentRecord struct {
	MSSV string                   `json:"MMSV"`
	Data InterfaceScoreController `json:"Data"`
}

type InterfaceChangeClassController struct {
	Semester      string   `json:"semester,omitempty"`
	Name          string   `json:"name,omitempty"`
	CourseId      any  `json:"course_id,omitempty"`
	CsvURL        string   `json:"csv_url,omitempty" bson:"csv_url,omitempty"`
}