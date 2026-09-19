package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"HelloWorld/internal/models"
)

// FetchCourses queries the Canvas API and filters for active enrolled courses
func FetchCourses() ([]models.Course, error) {
	token := os.Getenv("CANVAS_API_TOKEN")
	baseURL := os.Getenv("CANVAS_BASE_URL")

	// Append query parameters to only fetch active courses where the user is enrolled
	url := fmt.Sprintf("%s/api/v1/courses?enrollment_state=active&state[]=available", baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("canvas API returned status: %d", resp.StatusCode)
	}

	var rawCourses []models.Course
	if err := json.NewDecoder(resp.Body).Decode(&rawCourses); err != nil {
		return nil, err
	}

	// Filter out any courses that don't have a valid Name (prevents blank cards)
	var activeCourses []models.Course
	for _, course := range rawCourses {
		if course.Name != "" {
			// Fallback course code if Canvas omits it
			if course.Code == "" {
				course.Code = "No Course Code"
			}
			activeCourses = append(activeCourses, course)
		}
	}

	return activeCourses, nil
}

// FetchAssignments queries Canvas for a specific course's assignments
func FetchAssignments(courseID string) ([]models.Assignment, error) {
	token := os.Getenv("CANVAS_API_TOKEN")
	baseURL := os.Getenv("CANVAS_BASE_URL")

	url := fmt.Sprintf("%s/api/v1/courses/%s/assignments", baseURL, courseID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("canvas API returned status: %d", resp.StatusCode)
	}

	var assignments []models.Assignment
	if err := json.NewDecoder(resp.Body).Decode(&assignments); err != nil {
		return nil, err
	}

	return assignments, nil
}

func SubmitAssignment(courseID string, assignmentID int, submissionURL string) error {
	token := os.Getenv("CANVAS_API_TOKEN")
	baseURL := os.Getenv("CANVAS_BASE_URL")
	devMode := os.Getenv("DEV_MODE") == "true"

	// If dev flag is active, mock the submission for testing safely
	if devMode {
		log.Printf("[DEV_MODE] Mock submitting assignment %d for course %s with URL: %s", assignmentID, courseID, submissionURL)
		return nil
	}

	apiURL := fmt.Sprintf("%s/api/v1/courses/%s/assignments/%d/submissions", baseURL, courseID, assignmentID)

	// Prepare form data for Canvas submission
	data := url.Values{}
	data.Set("submission[submission_type]", "online_url")
	data.Set("submission[url]", submissionURL)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("canvas submission failed with status: %d", resp.StatusCode)
	}

	return nil
}
