package keboola

// api_endpoints.go contains definitions of all API endpoints used in Keboola Go Client.
// Centralizing endpoints allows easier maintenance and updates if API addresses change.

const (
	// Endpoints for encryption API.
	EncryptionAPIEncrypt = "encrypt"

	// Endpoints for notification API.
	NotificationAPIProjectSubscriptions = "project-subscriptions"
	NotificationAPIProjectSubscription  = "project-subscriptions/{subscriptionId}"

	// Endpoints for queue API.
	QueueAPIJobs       = "jobs"
	QueueAPIJob        = "jobs/{jobId}"
	QueueAPISearchJobs = "search/jobs"

	// Endpoints for scheduler API.
	SchedulerAPISchedules       = "schedules"
	SchedulerAPISchedule        = "schedules/{scheduleId}"
	SchedulerAPIConfigSchedules = "configurations/{configurationId}"
	SchedulerAPIRefreshToken    = "schedules/{scheduleId}/refreshToken"

	// Endpoints for data science apps (sandboxes service /apps).
	DataScienceAPIApps = "apps"
	DataScienceAPIApp  = "apps/{appId}"

	// Endpoints for editor API.
	EditorAPISessions                = "sql/sessions"
	EditorAPISession                 = "sql/sessions/{sessionId}"
	EditorAPISessionResetCredentials = "sql/sessions/{sessionId}/reset-credentials" // nolint: gosec
)
