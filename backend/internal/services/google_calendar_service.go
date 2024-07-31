package services

import (
	"context"
	"faladev/internal/auth"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendarService struct{}

func NewGoogleCalendarService() CalendarService {
	return &GoogleCalendarService{}
}

func (googleCalendarService *GoogleCalendarService) InitializeService(ctx context.Context, config *oauth2.Config, token *oauth2.Token) (CalendarAPI, error) {

	client, err := auth.CreateOAuthClient(ctx, config, token)

	if err != nil {
		return nil, fmt.Errorf("error creating OAuth client: %v", err)
	}

	service, err := calendar.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {
		return nil, fmt.Errorf("error creating calendar service: %v", err)
	}
	return &RealCalendarService{CalendarService: service}, nil
}

func (googleCalendarService *GoogleCalendarService) FindEventByKey(ctx context.Context, api CalendarAPI, eventKey string) (*calendar.Event, error) {

	event, err := api.GetEvent("primary", eventKey).Do()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching event with eventKey %s", eventKey))
	}

	return event, nil
}

func (googleCalendarService *GoogleCalendarService) AddGuestToEvent(ctx context.Context, api CalendarAPI, hangoutLink, email string) (*calendar.Event, error) {

	eventDetails, err := googleCalendarService.FindEventByKey(ctx, api, hangoutLink)

	if err != nil {
		return nil, err
	}

	updatedEvent, err := api.GetEvent("primary", eventDetails.Id).Do()

	if err != nil {
		return nil, errors.Wrap(err, "error getting event details")
	}

	for _, attendee := range updatedEvent.Attendees {
		if attendee.Email == email {
			return updatedEvent, nil
		}
	}

	attendee := &calendar.EventAttendee{Email: email}

	updatedEvent.Attendees = append(updatedEvent.Attendees, attendee)

	_, err = api.UpdateEvent("primary", updatedEvent.Id, updatedEvent).Do()

	if err != nil {
		return nil, errors.Wrap(err, "error adding guest to event")
	}

	return updatedEvent, nil
}
