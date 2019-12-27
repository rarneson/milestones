package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rarneson/milestones/github"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "close",
				Usage: "close a milestone across your repositories",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "add a milestone to your repositories",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "token",
						Aliases: []string{"t"},
						Usage:   "Your personal access token for GitHub",
					},
					&cli.StringFlag{
						Name:    "server",
						Aliases: []string{"s"},
						Value:   "github.com",
						EnvVars: []string{"MILESTONES_GIT_SERVER"},
						Usage:   "GitHub server",
					},
					&cli.StringFlag{
						Name:    "repos",
						Usage:   "comma separated list of repositories in the format org/repo",
						EnvVars: []string{"MILESTONES_REPOS"},
					},
				},
				Action: addAction,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func addAction(c *cli.Context) error {
	gitOptions := github.Options{
		Token:  c.String("token"),
		Server: c.String("server"),
	}

	g, err := github.New(gitOptions)
	if err != nil {
		return err
	}

	// fmt.Printf("%+v\n", gitOptions)

	var questions = []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "What is the name of this milestone?"},
			Validate: survey.Required,
		},
		{
			Name:     "dueDate",
			Prompt:   &survey.Input{Message: "What is the date this milestone will end (yyyy-mm-dd)?"},
			Validate: survey.ComposeValidators(survey.Required, validateDateFormat),
		},
	}

	answers := struct {
		Name    string
		DueDate string
	}{}

	// ask the common questions
	err = survey.Ask(questions, &answers)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", g) // eventually, g.AddMilestones

	return nil
}

func validateDateFormat(val interface{}) error {
	date := val.(string)

	match, err := regexp.MatchString(`^[\d]{4}-[\d]{2}-[\d]{2}$`, date)
	if err != nil {
		return err
	}

	if match == false {
		return errors.New("Due date needs to be in the format yyyy-mm-dd")
	}

	return nil
}
