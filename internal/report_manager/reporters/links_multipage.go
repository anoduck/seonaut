package reporters

import (
	"github.com/stjudewashere/seonaut/internal/models"
)

// Creates a MultipageIssueReporter object that contains the SQL query to check for indexable pages
// that are internally linked using the nofollow attribute.
func NoFollowIndexableReporter(c *models.Crawl) *MultipageIssueReporter {
	query := `
		SELECT
			pagereports.id,
			pagereports.url,
			pagereports.title
		FROM pagereports
		INNER JOIN (
			SELECT
				a.pagereport_id
			FROM (
				SELECT
					DISTINCT links.pagereport_id,
					links.url_hash
				FROM links
				WHERE links.crawl_id = ? AND links.nofollow = 1
			) AS a INNER JOIN pagereports ON a.url_hash = pagereports.url_hash
			WHERE pagereports.noindex = 0 AND pagereports.crawled = 1 AND pagereports.crawl_id = ?
		) AS b
		ON b.pagereport_id = pagereports.id`

	return &MultipageIssueReporter{
		Query:      query,
		Parameters: []interface{}{c.Id, c.Id},
		ErrorType:  ErrorInternalNoFollowIndexable,
	}
}

// Creates a MultipageIssueReporter object that contains the SQL query to check for pages
// that are internally linked with and without the nofollow attribute.
func FollowNoFollowReporter(c *models.Crawl) *MultipageIssueReporter {
	query := `
		SELECT
			pagereports.id,
			pagereports.url,
			pagereports.title
		FROM pagereports WHERE crawl_id = ? and url_hash in (
			SELECT
				url_hash
			FROM links
			WHERE crawl_id = ?
			GROUP BY url_hash
			HAVING COUNT(DISTINCT nofollow) > 1
			)`

	return &MultipageIssueReporter{
		Query:      query,
		Parameters: []interface{}{c.Id, c.Id},
		ErrorType:  IncomingFollowNofollow,
	}
}