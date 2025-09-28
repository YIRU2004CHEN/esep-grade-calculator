package esepunittests

type GradeCalculator struct {
    grades []Grade
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

func NewGradeCalculator() *GradeCalculator {
    return &GradeCalculator{
        grades: make([]Grade, 0),
    }
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

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

//new
//append the grade(including 3 types)
func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
    gc.grades = append(gc.grades, Grade{
        Name:  name,
        Grade: grade,
        Type:  gradeType,
    })
}

//new
func (gc *GradeCalculator) calculateNumericalGrade() int {
	//	initialize the grades
    assignmentSum, assignmentCount := 0, 0
    examSum, examCount := 0, 0
    essaySum, essayCount := 0, 0

    for _, g := range gc.grades {
        switch g.Type {
        case Assignment:
            assignmentSum += g.Grade
            assignmentCount++
        case Exam:
            examSum += g.Grade
            examCount++
        case Essay:
            essaySum += g.Grade
            essayCount++
        }
    }

	//get each grade's average, count it 0 if no grade
    assignmentAvg := 0
    if assignmentCount > 0 {
        assignmentAvg = assignmentSum / assignmentCount
    }

    examAvg := 0
    if examCount > 0 {
        examAvg = examSum / examCount
    }

    essayAvg := 0
    if essayCount > 0 {
        essayAvg = essaySum / essayCount
    }

    weighted := float64(assignmentAvg)*0.5 + float64(examAvg)*0.35 + float64(essayAvg)*0.15
    return int(weighted)
}


func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
        return 0
    	}

	sum := 0
	for _, g := range grades {
		sum += g.Grade
	}

	return sum / len(grades)
}

//for bonus, pass and fail requirement
//assuming 60 is the passing grade
func (gc *GradeCalculator) GetPassFail() string {
    if gc.calculateNumericalGrade() >= 60 {
        return "Pass"
    }
    return "Fail"
}