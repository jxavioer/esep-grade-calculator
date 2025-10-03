package esepunittests

type GradeCalculator struct {
	/* ORIGINAL CODE
	assignments []Grade
	exams       []Grade
	essays      []Grade
	*/
	grades     []Grade
	isPassFail bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(isPassFail bool) *GradeCalculator {
	return &GradeCalculator{
		/* ORIGINAL CODE
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
		*/
		grades:     make([]Grade, 0),
		isPassFail: isPassFail,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.isPassFail {
		if numericalGrade >= 70 {
			return "Pass"
		}
		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	/* ORIGINAL CODE
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{
			Name:  name,
			Grade: grade,
			Type:  Assignment,
		})
	case Exam:
		gc.exams = append(gc.exams, Grade{
			Name:  name,
			Grade: grade,
			Type:  Exam,
		})
	case Essay:
		gc.essays = append(gc.essays, Grade{
			Name:  name,
			Grade: grade,
			Type:  Essay,
		})
	}
	*/
	gc.grades = append(gc.grades, Grade{name, grade, gradeType})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.grades, Assignment)
	exam_average := computeAverage(gc.grades, Exam)
	essay_average := computeAverage(gc.grades, Essay)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade, gt GradeType) int {
	sum := 0
	catCounter := 0

	for _, grade := range grades {
		if grade.Type == gt {
			sum += grade.Grade
			catCounter += 1
		}
	}

	return sum / catCounter
}
