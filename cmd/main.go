package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/peteryurkovich/cluid/pkg/apply"
)

func getPath(templateName string) string {
	switch templateName {
	case "Deployment":
		return "/templates/hack/openshift/config/deployment"
	case "Operators":
		return "/templates/hack/openshift/config/operators"
	case "Resources":
		return "/templates/hack/openshift/config/resources"
	default:
		return ""
	}
}

type State struct {
	Templates []string
}

func main() {
	var state = State{Templates: []string{}}

	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("Cluid").
			Description("Welcome to _Cluid_.\n\nWhat would you like to deploy today?\n\n").
			Next(true).
			NextLabel("Let's go!"),
		),

		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Deployments").
				Description("Select your desired deployments.").
				Options(
					huh.NewOption("Deployment", "Deployment").Selected(true),
					huh.NewOption("Operators", "Operators"),
					huh.NewOption("Resources", "Resources"),
				).
				Validate(func(templates []string) error {
					if len(templates) <= 0 {
						return fmt.Errorf("at least one topping is required")
					}
					return nil
				}).
				Value(&state.Templates).
				Filterable(true).
				Limit(4),
		),
	).WithAccessible(accessible)

	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	action := func() {
		for _, template := range state.Templates {
			apply.Apply(getPath(template))
		}
	}
	spinner.New().Title("Applying Deployments...").Action(action).Run()

	{
		formatDeployments := func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		readableDeployments := []string{}
		for _, t := range state.Templates {
			readableDeployments = append(readableDeployments, formatDeployments(t))
		}
		var sb strings.Builder

		fmt.Fprintf(&sb,
			"%s\n\n%s",
			lipgloss.NewStyle().Bold(true).Render("Deployment Results"),
			strings.Join(readableDeployments, "\n"),
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
}
