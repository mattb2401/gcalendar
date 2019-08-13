package src

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	calendar "google.golang.org/api/calendar/v3"
)

//ListAllEvents lists all upcoming events on your
func ListAllEvents() error {
	tokFile := "token.json"
	_, err := tokenFromFile(tokFile)
	if err != nil {
		return err
	}
	client, err := GetClientToken()
	if err != nil {
		return err
	}
	srv, err := calendar.New(client)
	if err != nil {
		return fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		return fmt.Errorf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("These are your upcoming events on your calendar")
	if len(events.Items) == 0 {
		fmt.Println("You have no upcoming events on your calendar.")
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "NAME", "DATE"})
		dt := [][]string{}
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			dt = append(dt, []string{item.Id, item.Summary, date})
		}
		for _, t := range dt {
			table.Append(t)
		}
		table.Render()
	}
	return nil
}

//ListEventsByDate lists today's upcoming events
func ListEventsByDate(date string) error {
	tokFile := "token.json"
	_, err := tokenFromFile(tokFile)
	if err != nil {
		return err
	}
	client, err := GetClientToken()
	if err != nil {
		return err
	}
	srv, err := calendar.New(client)
	if err != nil {
		return fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}
	layout := "2006-01-02"
	day, err := time.Parse(layout, date)
	if err != nil {
		return err
	}
	dayAfter := day.AddDate(0, 0, 1).Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(day.Format(time.RFC3339)).TimeMax(dayAfter).MaxResults(1).OrderBy("startTime").Do()
	if err != nil {
		return fmt.Errorf("Unable to retrieve today's user's events: %v", err)
	}
	if len(events.Items) == 0 {
		fmt.Println("You have no upcoming events on your calendar for this day.")
	} else {
		fmt.Println("These are your upcoming events on your calendar for this day")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"EVENTID", "NAME", "DATE"})
		dt := [][]string{}
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			dt = append(dt, []string{item.Id, item.Summary, date})
		}
		for _, t := range dt {
			table.Append(t)
		}
		table.Render()
	}
	return nil
}

//AddEventToCalendar adds new events to your calendar.
func AddEventToCalendar(reoccurance string, attendies []string) error {
	var (
		summary     string
		location    string
		description string
		startTime   string
		endTime     string
	)
	tokFile := "token.json"
	_, err := tokenFromFile(tokFile)
	if err != nil {
		return err
	}
	client, err := GetClientToken()
	if err != nil {
		return err
	}
	fmt.Println("gCalendar")
	fmt.Println("Creating a new event.... \n")
	fmt.Println("What is the summary of the event?")
	reader := bufio.NewReader(os.Stdin)
	summary, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println("Where is the event going to take place?")
	location, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println("Please enter the description of the event")
	description, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println("What time does the event start? eg. 2006-01-02 15:00:00")
	startTime, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println("What time does the event end? eg. 2006-01-02 15:00:00")
	endTime, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	srv, err := calendar.New(client)
	if err != nil {
		return fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}
	layout := "2006-01-02 15:04:05"
	startDate, err := time.Parse(layout, startTime[:len(startTime)-1])
	if err != nil {
		return err
	}
	endDate, err := time.Parse(layout, endTime[:len(endTime)-1])
	if err != nil {
		return err
	}
	event := &calendar.Event{
		Summary:     summary,
		Location:    location,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: startDate.Format(time.RFC3339),
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: endDate.Format(time.RFC3339),
			TimeZone: "America/Los_Angeles",
		},
	}
	if reoccurance != "" {
		event.Recurrence = []string{reoccurance}
	}
	if len(attendies) > 0 {
		var emails []*calendar.EventAttendee
		for _, attendie := range attendies {
			emails = append(emails, &calendar.EventAttendee{
				Email: attendie,
			})
		}
		event.Attendees = emails
	}
	_, err = srv.Events.Insert("primary", event).Do()
	if err != nil {
		return err
	}
	fmt.Print("\nEvent is being added.. \n")
	return nil
}

//UpdateEventOnCalendar adds new events to your calendar.
func UpdateEventOnCalendar(reoccurance string, attendies []string, eventID string) error {
	var (
		summary     string
		location    string
		description string
		startTime   string
		endTime     string
	)
	tokFile := "token.json"
	_, err := tokenFromFile(tokFile)
	if err != nil {
		return err
	}
	client, err := GetClientToken()
	if err != nil {
		return err
	}
	fmt.Println("gCalendar")
	fmt.Println("Updating a event.... \n")
	fmt.Println("What is the summary of the event?")
	reader := bufio.NewReader(os.Stdin)
	summary, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println("Where is the event going to take place?")
	location, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println("Please enter the description of the event")
	description, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println("What time does the event start? eg. 2006-01-02 15:00:00")
	startTime, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println("What time does the event end? eg. 2006-01-02 15:00:00")
	endTime, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	srv, err := calendar.New(client)
	if err != nil {
		return fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}
	layout := "2006-01-02 15:04:05"
	startDate, err := time.Parse(layout, startTime[:len(startTime)-1])
	if err != nil {
		return err
	}
	endDate, err := time.Parse(layout, endTime[:len(endTime)-1])
	if err != nil {
		return err
	}
	event := &calendar.Event{
		Summary:     summary,
		Location:    location,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: startDate.Format(time.RFC3339),
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: endDate.Format(time.RFC3339),
			TimeZone: "America/Los_Angeles",
		},
	}
	if reoccurance != "" {
		event.Recurrence = []string{reoccurance}
	}
	if len(attendies) > 0 {
		var emails []*calendar.EventAttendee
		for _, attendie := range attendies {
			emails = append(emails, &calendar.EventAttendee{
				Email: attendie,
			})
		}
		event.Attendees = emails
	}
	_, err = srv.Events.Update("primary", eventID, event).Do()
	if err != nil {
		return err
	}
	fmt.Print("\nEvent is being added.. \n")
	return nil
}

//DeleteEventToCalendar deletes events from calendar
func DeleteEventToCalendar(eventID string) error {
	tokFile := "token.json"
	_, err := tokenFromFile(tokFile)
	if err != nil {
		return err
	}
	client, err := GetClientToken()
	if err != nil {
		return err
	}
	srv, err := calendar.New(client)
	if err != nil {
		return fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}
	err = srv.Events.Delete("primary", eventID).Do()
	if err != nil {
		return fmt.Errorf("Unable to retrieve today's user's events: %v", err)
	}
	fmt.Print("\nEvent is being removed from your calendar.. \n")
	return nil
}
