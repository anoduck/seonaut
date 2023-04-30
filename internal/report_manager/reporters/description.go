package reporters

import (
	"github.com/stjudewashere/seonaut/internal/models"
)

// Returns a PageIssueReporter with a callback function that checks if a page has
// an empty or missing description. It returns true if the status code is between
// 200 and 299, the media type is text/html and the description is not set.
func NewEmptyDescriptionReporter() *PageIssueReporter {
	c := func(pageReport *models.PageReport) bool {
		if pageReport.Crawled == false {
			return false
		}

		if pageReport.MediaType != "text/html" {
			return false
		}

		if pageReport.StatusCode < 200 || pageReport.StatusCode >= 300 {
			return false
		}

		return pageReport.Description == ""
	}

	return &PageIssueReporter{
		ErrorType: ErrorEmptyDescription,
		Callback:  c,
	}
}

// Returns a PageIssueReporter with a callback function that checks if a page has a short description.
// The callback function returns true if the page is text/html, has a status code between 200 and 299,
// and has a description of less than an specified amount of letters.
func NewShortDescriptionReporter() *PageIssueReporter {
	c := func(pageReport *models.PageReport) bool {
		if pageReport.Crawled == false {
			return false
		}

		if pageReport.MediaType != "text/html" {
			return false
		}

		if pageReport.StatusCode < 200 || pageReport.StatusCode >= 300 {
			return false
		}

		return len(pageReport.Description) > 0 && len(pageReport.Description) < 80
	}

	return &PageIssueReporter{
		ErrorType: ErrorShortDescription,
		Callback:  c,
	}
}

// Returns a PageIssueReporter with a callback function that checks if a page has a short description.
// The callback function returns true if the page is text/html, has a status code between 200 and 299,
// and has a description of more than an specified amount of letters.
func NewLongDescriptionReporter() *PageIssueReporter {
	c := func(pageReport *models.PageReport) bool {
		if pageReport.Crawled == false {
			return false
		}

		if pageReport.MediaType != "text/html" {
			return false
		}

		if pageReport.StatusCode < 200 || pageReport.StatusCode >= 300 {
			return false
		}

		return len(pageReport.Description) > 160
	}

	return &PageIssueReporter{
		ErrorType: ErrorLongDescription,
		Callback:  c,
	}
}
