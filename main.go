package main

import (
	"fmt"
	"os"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	xstrings "github.com/charmbracelet/x/exp/strings"
	"strings"
)

var Theme = huh.ThemeCatppuccin()
func main() {
	err := huh.NewNote().
		Title("Welcome to the Exam Revision Helper").
		Description("This is a simple tool to help you revise for your exams ✌️✌️✌️").
		WithTheme(Theme).
		Run()
	
	HelpRevise()
	
	HandleError(err)
}

func GetSubject() string {
	var subject string
	err:= huh.NewSelect[string]().
		Title("Select a subject").
		Options(
			huh.NewOption("DAA","DAA"),
			huh.NewOption("OS","OS"),
		).
		Value(&subject).
		WithTheme(Theme).
		Run()
			
	HandleError(err)
	return subject
}

func HelpRevise() string{
	var subject, unit string
	subject = GetSubject()
	unit = GetUnit()

	DisplaySelection(subject, unit)

	switch subject{
		case "DAA":
			DAA(unit)
		case "OS":
			return "OS"
	}

	return ""
}

func GetUnit() string{
	var unit string
	err := huh.NewSelect[string]().
		Title("Select a unit").
		Options(
			huh.NewOption("Unit 1","Unit 1"),
			huh.NewOption("Unit 2","Unit 2"),
			huh.NewOption("Unit 3","Unit 3"),
			huh.NewOption("Unit 4","Unit 4"),
		).
		Value(&unit).
		WithTheme(Theme).
		Run()
	
	HandleError(err)

	return unit
}

func DAA(unit string) string{
	var subtopic string
	switch unit {
	case "Unit 1":
		return "Unit 1"		
	case "Unit 2":
		return "Unit 2"
	case "Unit 3":
		err:= huh.NewSelect[string]().
		Title("Select a sub-topic").
		Description("Press anykey to exit").
		Options(
			huh.NewOption("Intro Presorting","Intro Presorting"),
			huh.NewOption("Heap Sort","Heap Sort"),
			huh.NewOption("Red-Black Trees","Red-Black Trees"),
			huh.NewOption("2-3 Trees","2-3 Trees"),
		).
		Value(&subtopic).
		WithTheme(Theme).
		Run()
		
		material := DAA_Helper(subtopic, unit)
		DisplayRevision(material)

		HandleError(err)
	}
	
	return unit
}


func DAA_Helper(subtopic string, unit string) string{
	var revision string
	switch unit{
		case "Unit 1":
			switch subtopic{
				
			} 
		case "Unit 2":
			switch subtopic{
				
			} 
		
		case "Unit 3":
			switch subtopic{
				case "Intro Presorting":
					revision = `ALGORITHM PresortElementUniqueness(A[0..n − 1])
						//Solves the element uniqueness problem by sorting the array first
						//Input: An array A[0..n − 1] of orderable elements
						//Output: Returns “true” if A has no equal elements, “false” otherwise
						sort the array A
						for i ← 0 to n − 2 do
							if A[i] = A[i + 1]return false
						return true`
			}
	}

	return revision
}

func DisplaySelection(subject string, unit string) string{
	
	huh.NewNote().
		Title("You have selected: " + subject +" " + unit).
		WithTheme(Theme).
		Run()
	
	
	return subject

}
func Display(input string) string{
	huh.NewNote().
	Title(input).
	WithTheme(Theme).
	Run()

	return ""
}


func DisplayRevision(revision string, subtopic string){
	var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		fmt.Fprintf(&sb,
			"Topic: %s\n Important Points: %s\n"
			lipgloss.NewStyle().Bold(true).Render("REVISION "),
			keyword(subtopic),
			keyword(revision),
		)
		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(sb.String()),
		)
}



func HandleError(err error) {
	if err == huh.ErrUserAborted {
		os.Exit(0)
	} else if err != nil {
		fmt.Println("An error occurred:", err)
		os.Exit(1)
	}
}