package banner

type BannerTerm struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type BannerResponse struct {
	Success               bool        `json:"success"`
	TotalCount            int64       `json:"totalCount"`
	Data                  []Datum     `json:"data"`
	PageOffset            int64       `json:"pageOffset"`
	PageMaxSize           int64       `json:"pageMaxSize"`
	SectionsFetchedCount  int64       `json:"sectionsFetchedCount"`
	PathMode              string      `json:"pathMode"`
	SearchResultsConfigs  interface{} `json:"searchResultsConfigs"`
	ZtcEncodedImage       interface{} `json:"ztcEncodedImage"`
	AllowHoldRegistration interface{} `json:"allowHoldRegistration"`
}

type Datum struct {
	ID                             int64             `json:"id"`
	Term                           string            `json:"term"`
	TermDesc                       string            `json:"termDesc"`
	CourseReferenceNumber          string            `json:"courseReferenceNumber"`
	PartOfTerm                     string            `json:"partOfTerm"`
	CourseNumber                   string            `json:"courseNumber"`
	Subject                        string            `json:"subject"`
	SubjectDescription             string            `json:"subjectDescription"`
	SequenceNumber                 string            `json:"sequenceNumber"`
	CampusDescription              string            `json:"campusDescription"`
	ScheduleTypeDescription        string            `json:"scheduleTypeDescription"`
	CourseTitle                    string            `json:"courseTitle"`
	CreditHours                    float64           `json:"creditHours"`
	MaximumEnrollment              int64             `json:"maximumEnrollment"`
	Enrollment                     int64             `json:"enrollment"`
	SeatsAvailable                 int64             `json:"seatsAvailable"`
	WaitCapacity                   int64             `json:"waitCapacity"`
	WaitCount                      int64             `json:"waitCount"`
	WaitAvailable                  int64             `json:"waitAvailable"`
	CrossList                      interface{}       `json:"crossList"`
	CrossListCapacity              interface{}       `json:"crossListCapacity"`
	CrossListCount                 interface{}       `json:"crossListCount"`
	CrossListAvailable             interface{}       `json:"crossListAvailable"`
	CreditHourHigh                 float64           `json:"creditHourHigh"`
	CreditHourLow                  float32           `json:"creditHourLow"`
	CreditHourIndicator            string            `json:"creditHourIndicator"`
	OpenSection                    bool              `json:"openSection"`
	LinkIdentifier                 string            `json:"linkIdentifier"`
	IsSectionLinked                bool              `json:"isSectionLinked"`
	SubjectCourse                  string            `json:"subjectCourse"`
	Faculty                        []Faculty         `json:"faculty"`
	MeetingsFaculty                []MeetingsFaculty `json:"meetingsFaculty"`
	ReservedSeatSummary            interface{}       `json:"reservedSeatSummary"`
	SectionAttributes              interface{}       `json:"sectionAttributes"`
	InstructionalMethod            string            `json:"instructionalMethod"`
	InstructionalMethodDescription string            `json:"instructionalMethodDescription"`
}

type Faculty struct {
	BannerID              string      `json:"bannerId"`
	Category              interface{} `json:"category"`
	Class                 string      `json:"class"`
	CourseReferenceNumber string      `json:"courseReferenceNumber"`
	DisplayName           string      `json:"displayName"`
	EmailAddress          string      `json:"emailAddress"`
	PrimaryIndicator      bool        `json:"primaryIndicator"`
	Term                  string      `json:"term"`
}

type MeetingsFaculty struct {
	Category              string        `json:"category"`
	Class                 string        `json:"class"`
	CourseReferenceNumber string        `json:"courseReferenceNumber"`
	Faculty               []interface{} `json:"faculty"`
	MeetingTime           MeetingTime   `json:"meetingTime"`
	Term                  string        `json:"term"`
}

type MeetingTime struct {
	BeginTime              string  `json:"beginTime"`
	Building               string  `json:"building"`
	BuildingDescription    string  `json:"buildingDescription"`
	Campus                 string  `json:"campus"`
	CampusDescription      string  `json:"campusDescription"`
	Category               string  `json:"category"`
	Class                  string  `json:"class"`
	CourseReferenceNumber  string  `json:"courseReferenceNumber"`
	CreditHourSession      float64 `json:"creditHourSession"`
	EndDate                string  `json:"endDate"`
	EndTime                string  `json:"endTime"`
	Friday                 bool    `json:"friday"`
	HoursWeek              float64 `json:"hoursWeek"`
	MeetingScheduleType    string  `json:"meetingScheduleType"`
	MeetingType            string  `json:"meetingType"`
	MeetingTypeDescription string  `json:"meetingTypeDescription"`
	Monday                 bool    `json:"monday"`
	Room                   string  `json:"room"`
	Saturday               bool    `json:"saturday"`
	StartDate              string  `json:"startDate"`
	Sunday                 bool    `json:"sunday"`
	Term                   string  `json:"term"`
	Thursday               bool    `json:"thursday"`
	Tuesday                bool    `json:"tuesday"`
	Wednesday              bool    `json:"wednesday"`
}
