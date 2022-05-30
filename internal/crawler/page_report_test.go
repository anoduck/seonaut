package crawler_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/stjudewashere/seonaut/internal/crawler"
)

const (
	testURL = "https://example.com/test-page/"
)

func TestNewPageReport(t *testing.T) {
	u, err := url.Parse(testURL)
	if err != nil {
		fmt.Println(err)
	}

	contentType := "text/html"
	statusCode := 200
	body := []byte("<html>")

	headers := http.Header{
		"Content-Type": []string{contentType},
	}

	pageReport := crawler.NewPageReport(u, statusCode, &headers, body)

	if pageReport.URL != testURL {
		t.Error("NewPageReport URL != testURL")
	}

	if pageReport.ParsedURL != u {
		t.Error("NewPageReport ParsedURL != u")
	}

	if pageReport.StatusCode != statusCode {
		t.Error("NewPageReport StatusCode != statusCode")
	}

	if pageReport.ContentType != "text/html" {
		t.Error("NewPageReport ContentType != contentType")
	}

	if string(pageReport.Body) != string(body) {
		t.Error("NewPageReport Body != body")
	}
}

func TestNewRedirectPageReport(t *testing.T) {
	u, err := url.Parse(testURL)
	if err != nil {
		fmt.Println(err)
	}

	body := []byte("<html>")
	statusCode := 301
	redirectURL := "https://example.com/redirect"

	headers := http.Header{
		"Location":     []string{redirectURL},
		"Content-Type": []string{"text/html"},
	}

	pageReport := crawler.NewPageReport(u, statusCode, &headers, body)

	if pageReport.RedirectURL != redirectURL {
		t.Errorf("NewPageReport RedirectURL != %s", pageReport.RedirectURL)
	}

	if pageReport.StatusCode != statusCode {
		t.Error("NewPageReport StatusCode != statusCode")
	}
}

func TestPageReportHTML(t *testing.T) {
	u, err := url.Parse(testURL)
	if err != nil {
		fmt.Println(err)
	}

	contentType := "text/html"
	statusCode := 200

	body, err := ioutil.ReadFile("./testdata/test.html")
	if err != nil {
		log.Fatal(err)
	}

	headers := &http.Header{
		"Content-Type": []string{contentType},
	}

	pageReport := crawler.NewPageReport(u, statusCode, headers, body)

	if pageReport.Lang != "en" {
		t.Error("Lang != en")
	}

	if pageReport.Title != "Test Page Title" {
		t.Error("Title != Test Page Title")
	}

	if pageReport.Description != "Test Page Description" {
		t.Error("Description != Test Page Description")
	}

	if len(pageReport.Links) != 6 {
		t.Error("len(Links) != 6")
	}

	if len(pageReport.Links) > 0 {
		if pageReport.Links[0].URL != "https://example.com/link1" {
			t.Error("pageReport.Links[0].URL != https://example.com/link1")
		}
		if pageReport.Links[1].URL != "https://example.com/test-page/link2" {
			t.Error(pageReport.URL)
			t.Errorf("%s != https://example.com/test-page/link2", pageReport.Links[1].URL)
		}
		if pageReport.Links[0].Text != "link1" {
			t.Error("pageReport.Links[0].Text != link1")
		}
		if pageReport.Links[0].Rel != "nofollow" {
			t.Error("pageReport.Links[0].Rel != nofollow")
		}
		if pageReport.Links[0].External != false {
			t.Error("pageReport.Links[0].External != false")
		}
		if pageReport.Links[3].Text != "" {
			t.Error("pageReport.Links[3].Text != \"\"")
		}

		if pageReport.Links[4].URL != "https://example.com/" {
			t.Errorf("%s != \"https://example.com/\"", pageReport.Links[4].URL)
		}

		if pageReport.Links[5].URL != "https://example.com/test-page/" {
			t.Errorf("%s != \"https://example.com/test-page\"", pageReport.Links[5].URL)
		}
	}

	if len(pageReport.ExternalLinks) != 1 {
		t.Error("len(pageReport.ExternalLinks) != 1")
	}

	if pageReport.Refresh != "0;URL='/'" {
		t.Errorf("Refresh != \"0;URL='%s'\"", pageReport.Refresh)
	}

	if pageReport.RedirectURL != "https://example.com/" {
		t.Error("RedirectURL != https://example.com/")
	}

	if pageReport.Robots != "noindex, nofollow" {
		t.Error("Robots != noindex, nofollow")
	}

	if pageReport.Canonical != "https://example.com/canonical/" {
		t.Error("Canonical != https://example.com/canonical/")
	}

	if pageReport.H1 != "H1 Title" {
		t.Error("H1 != H1 Title")
	}

	if pageReport.H2 != "H2 Title" {
		t.Error("H2 != H2 Title")
	}

	if pageReport.Words != 10 {
		t.Error("Words != 10")
	}

	if len(pageReport.Hreflangs) != 1 {
		t.Error("Hreflang != 1")
	}

	if len(pageReport.Hreflangs) == 1 && pageReport.Hreflangs[0].URL != "https://example.com/fr" {
		t.Error("Hreglangs[0].URL != https://example.com/fr")
	}

	if len(pageReport.Hreflangs) == 1 && pageReport.Hreflangs[0].Lang != "fr" {
		t.Error("Hreglangs[0].URL != fr")
	}

	if len(pageReport.Images) != 7 {
		t.Error("Images != 7")
	}

	if pageReport.Images[0].URL != "https://example.com/img/logo.png" {
		t.Error("pageReport.Images[0].URL != https://example.com/img/logo.png")
	}

	if len(pageReport.Scripts) != 1 {
		t.Error("Scripts != 1")
	}

	if len(pageReport.Scripts) == 1 && pageReport.Scripts[0] != "https://example.com/js/app.js" {
		t.Error("Scripts[0] != https://example.com/js/app.js")
	}

	if len(pageReport.Styles) != 1 {
		t.Error("Styles != 1")
	}

	if len(pageReport.Styles) == 1 && pageReport.Styles[0] != "https://example.com/css/style.css" {
		t.Error("Styles[0] != https://example.com/css/style.css")
	}

	if pageReport.ValidHeadings == true {
		t.Error("pageReport.validHeadings == true")
	}

	if pageReport.Noindex == false {
		t.Error("pageReport.Noindex == false")
	}
}

func TestNoindex(t *testing.T) {
	u, err := url.Parse(testURL)
	if err != nil {
		fmt.Println(err)
	}

	body := []byte("<html>")
	statusCode := 200
	headers := http.Header{
		"X-Robots-Tag": []string{"noindex, nofollow"},
		"Content-Type": []string{"text/html"},
	}

	pageReport := crawler.NewPageReport(u, statusCode, &headers, body)

	if pageReport.Nofollow == false {
		t.Error("Nofollow == false")
	}

	if pageReport.Noindex == false {
		t.Error("Noindex == false")
	}
}

func TestContentLanguage(t *testing.T) {
	u, err := url.Parse(testURL)
	if err != nil {
		fmt.Println(err)
	}

	body := []byte("<html>")
	statusCode := 200
	contentLanguage := "en-us"
	headers := http.Header{
		"Content-Language": []string{contentLanguage},
		"Content-Type":     []string{"text/html"},
	}

	pageReport := crawler.NewPageReport(u, statusCode, &headers, body)

	if pageReport.Lang != contentLanguage {
		t.Errorf("ContentLanguage: %s != %s", pageReport.Lang, contentLanguage)
	}
}

func TestHreflangHeaders(t *testing.T) {
	u, err := url.Parse(testURL)
	if err != nil {
		fmt.Println(err)
	}

	linkHeader := `
		<https://example.com/file.pdf>; rel="alternate"; hreflang="en",
		<https://de-ch.example.com/file.pdf>; rel="alternate"; hreflang="de-ch",
		<https://de.example.com/file.pdf>; rel="alternate"; hreflang="de"
	`

	body := []byte("<html>")
	statusCode := 200
	headers := http.Header{
		"Link":         []string{linkHeader},
		"Content-Type": []string{"text/html"},
	}

	pageReport := crawler.NewPageReport(u, statusCode, &headers, body)

	if len(pageReport.Hreflangs) != 3 {
		t.Errorf("HreflangHeader: %d != 3", len(pageReport.Hreflangs))
	}

	if pageReport.Hreflangs[0].URL != "https://example.com/file.pdf" || pageReport.Hreflangs[0].Lang != "en" {
		t.Errorf("HreflangHeader: Hreflangs[0]: %v ", pageReport.Hreflangs[0])
	}

	if pageReport.Hreflangs[1].URL != "https://de-ch.example.com/file.pdf" || pageReport.Hreflangs[1].Lang != "de-ch" {
		t.Errorf("HreflangHeader: Hreflangs[1]: %v ", pageReport.Hreflangs[1])
	}

	if pageReport.Hreflangs[2].URL != "https://de.example.com/file.pdf" || pageReport.Hreflangs[2].Lang != "de" {
		t.Errorf("HreflangHeader: Hreflangs[2]: %v ", pageReport.Hreflangs[2])
	}
}

func TestCanonicalHeaders(t *testing.T) {
	u, err := url.Parse(testURL)
	if err != nil {
		fmt.Println(err)
	}

	linkHeader := `
		<https://example.com/canonical>; rel="canonical",
		<https://de-ch.example.com/file.pdf>; rel="alternate"; hreflang="de-ch",
		<https://de.example.com/file.pdf>; rel="alternate"; hreflang="de"
	`

	body := []byte("<html>")
	statusCode := 200
	headers := http.Header{
		"Link":         []string{linkHeader},
		"Content-Type": []string{"text/html"},
	}

	pageReport := crawler.NewPageReport(u, statusCode, &headers, body)

	if pageReport.Canonical != "https://example.com/canonical" {
		t.Errorf("Canonical headers: %s != https://example.com/canonical", pageReport.Canonical)
	}
}