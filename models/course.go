package models

// Course represents a course resource
type Course struct {
	ApplicationProvider string             `bson:"application_uk_provider_reference_number,omitempty" json:"application_uk_provider_reference_number,omitempty"`
	Country             *Country           `bson:"country" json:"country"`
	DistanceLearning    *DistanceLearning  `bson:"distance_learning" json:"distance_learning"`
	Foundation          string             `bson:"foundation_year_availability" json:"foundation_year_availability"` // enum
	Honours             bool               `bson:"honours_award_provision" json:"honours_award_provision"`
	ID                  string             `bson:"_id"`
	Institution         *InstitutionObject `bson:"institution" json:"institution"`
	KISCourseID         string             `bson:"kis_course_id" json:"kis_course_id"`
	Length              *LengthObject      `bson:"length_of_course" json:"length_of_course"`
	Links               *LinkList          `bson:"links" json:"links"`
	Location            *Location          `bson:"location" json:"location"`
	Mode                *Mode              `bson:"mode" json:"mode"`
	NHSFunded           *NHSFunded         `bson:"nhs_funded,omitempty" json:"nhs_funded,omitempty"`
	Qualification       *Qualification     `bson:"qualification" json:"qualification"`
	SandwichYear        *Availability      `bson:"sandwich_year" json:"sandwich_year"`
	Statistics          *Statistics        `bson:"statistics" json:"statistics"`
	Title               *Language          `bson:"title" json:"title"`
	UCASCode            string             `bson:"ucas_code_id,omitempty" json:"ucas_code_id,omitempty"`
	YearAbroad          *Availability      `bson:"year_abroad" json:"year_abroad"`
}

// Availability represents an object referring to the availability
type Availability struct {
	Code  string `bson:"code" json:"code"`
	Label string `bson:"label" json:"label"` // enum , 0-2
}

// Country represents a country object
type Country struct {
	Code string `bson:"code" json:"code"`
	Name string `bson:"name" json:"name"`
}

// DistanceLearning represents an object referring
// to the course available through distance learning
type DistanceLearning struct {
	Code  string `bson:"code" json:"code"`
	Label string `bson:"label" json:"label"`
}

// InstitutionObject represents institution data related to course
type InstitutionObject struct {
	PublicUKPRNName string `bson:"public_ukprn_name" json:"public_ukprn_name"`
	PublicUKPRN     string `bson:"public_ukprn" json:"public_ukprn"`
	UKPRN           string `bson:"ukprn" json:"ukprn"`
	UKPRNName       string `bson:"ukprn_name" json:"ukprn_name"`
}

// Language represents an object containing english or welsh strings
type Language struct {
	English string `bson:"english,omitempty" json:"english,omitempty"`
	Welsh   string `bson:"welsh,omitempty" json:"welsh,omitempty"`
}

// LengthObject represents an object referring to the course length
type LengthObject struct {
	Code  string `bson:"code" json:"code"`
	Label string `bson:"label" json:"label"`
}

// LinkList represents a list of links related to resource
type LinkList struct {
	AssessmentMethod    *Language `bson:"assessment_method,omitempty" json:"assessment_method,omitempty"`                 // ASSURL
	CoursePage          *Language `bson:"course_page,omitempty" json:"course_page,omitempty"`                             // CRSEURL
	EmploymentDetails   *Language `bson:"employment_details,omitempty" json:"employment_details,omitempty"`               // EMPLOYURL
	FinancialSupport    *Language `bson:"financial_support_details,omitempty" json:"financial_support_details,omitempty"` // SUPPORTURL
	Institution         string    `bson:"institution" json:"institution"`
	LearningAndTeaching *Language `bson:"learning_and_teaching_methods,omitempty" json:"learning_and_teaching_methods,omitempty"` // LTURL
	Self                string    `bson:"self" json:"self"`
}

// Location represents an object containing fields to enable one to locate institution
type Location struct {
	Changes   bool      `bson:"changes" json:"changes"`
	Latitude  string    `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude string    `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Name      *Language `bson:"name,omitempty" json:"name,omitempty"`
}

// NHSFunded represents an object referring to the course having any NHS funded students
type NHSFunded struct {
	Code  string `bson:"code,omitempty" json:"code,omitempty"`
	Label string `bson:"label,omitempty" json:"label,omitempty"`
}

// Mode represents an object referring to the type of course
type Mode struct {
	Code  string `bson:"code" json:"code"`
	Label string `bson:"label" json:"label"`
}

// Qualification represents an object referring to the qualification received from course
type Qualification struct {
	Code  string `bson:"code" json:"code"`
	Label string `bson:"label" json:"label"`
	Level string `bson:"level" json:"level"`
	Name  string `bson:"name" json:"name"`
}

// LocationIDObject represents a course location object
type LocationIDObject struct {
	ID string `bson:"id" json:"id"`
}
