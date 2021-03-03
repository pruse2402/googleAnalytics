package dtos

// BaselineAssessmentResponse Response struct send to client
type BaselineAssessmentResponse struct {
	WelcomeNote        string `json:"welcomeNote"`
	Header             string `json:"header"`
	HeaderNote         string `json:"headerNote"`
	ButtonText         string `json:"buttonText"`
	Logo               string `json:"logo"`
	PatientHealth      string `json:"patientHealth"`
	Audit              string `json:"audit"`
	Habit              string `json:"habit"`
	Goal               string `json:"goal"`
	Strategy           string `json:"strategy"`
	SupportiveContacts string `json:"supportiveContacts"`
	EndNote            string `json:"endNote"`
}
