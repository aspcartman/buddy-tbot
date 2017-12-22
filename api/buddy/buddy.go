package buddy

import "time"

type WebHookEvent struct {
	Workspace struct {
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Domain  string `json:"domain"`
	} `json:"workspace"`
	Invoker struct {
		URL       string      `json:"url"`
		HTMLURL   string      `json:"html_url"`
		ID        int         `json:"id"`
		Name      string      `json:"name"`
		AvatarURL string      `json:"avatar_url"`
		Title     interface{} `json:"title"`
		Email     string      `json:"email"`
	} `json:"invoker"`
	Project struct {
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Status      string `json:"status"`
	} `json:"project"`
	CurrentDate time.Time `json:"current_date"`
	Execution struct {
		URL        string    `json:"url"`
		HTMLURL    string    `json:"html_url"`
		ID         int       `json:"id"`
		StartDate  time.Time `json:"start_date"`
		FinishDate time.Time `json:"finish_date"`
		Mode       string    `json:"mode"`
		Refresh    bool      `json:"refresh"`
		Status     string    `json:"status"`
		Comment    string    `json:"comment"`
		Branch struct {
			URL     string `json:"url"`
			HTMLURL string `json:"html_url"`
			Name    string `json:"name"`
			Default bool   `json:"default"`
		} `json:"branch"`
		FromRevision struct {
			URL           string `json:"url"`
			HTMLURL       string `json:"html_url"`
			Revision      string `json:"revision"`
			ShortRevision string `json:"short_revision"`
			Message       string `json:"message"`
			Committer struct {
				Name  string      `json:"name"`
				Title interface{} `json:"title"`
				Email string      `json:"email"`
			} `json:"committer"`
			Author struct {
				Title interface{} `json:"title"`
				Email string      `json:"email"`
			} `json:"author"`
			Tags []interface{} `json:"tags"`
		} `json:"from_revision"`
		ToRevision struct {
			URL           string `json:"url"`
			HTMLURL       string `json:"html_url"`
			Revision      string `json:"revision"`
			ShortRevision string `json:"short_revision"`
			Message       string `json:"message"`
			Committer struct {
				Name  string      `json:"name"`
				Title interface{} `json:"title"`
				Email string      `json:"email"`
			} `json:"committer"`
			Author struct {
				Title interface{} `json:"title"`
				Email string      `json:"email"`
			} `json:"author"`
			Tags []interface{} `json:"tags"`
		} `json:"to_revision"`
		Creator struct {
			URL       string      `json:"url"`
			HTMLURL   string      `json:"html_url"`
			ID        int         `json:"id"`
			Name      string      `json:"name"`
			AvatarURL string      `json:"avatar_url"`
			Title     interface{} `json:"title"`
			Email     string      `json:"email"`
		} `json:"creator"`
		Pipeline struct {
			URL                   string `json:"url"`
			HTMLURL               string `json:"html_url"`
			ID                    int    `json:"id"`
			Name                  string `json:"name"`
			TriggerMode           string `json:"trigger_mode"`
			RefName               string `json:"ref_name"`
			LastExecutionStatus   string `json:"last_execution_status"`
			LastExecutionRevision string `json:"last_execution_revision"`
			AlwaysFromScratch     bool   `json:"always_from_scratch"`
			NoSkipToMostRecent    bool   `json:"no_skip_to_most_recent"`
			AutoClearCache        bool   `json:"auto_clear_cache"`
		} `json:"pipeline"`
		ActionExecutions []struct {
			Status   string  `json:"status"`
			Progress float32 `json:"progress"`
			Action struct {
				URL                 string `json:"url"`
				HTMLURL             string `json:"html_url"`
				ID                  int    `json:"id"`
				Name                string `json:"name"`
				Type                string `json:"type"`
				TriggerTime         string `json:"trigger_time"`
				LastExecutionStatus string `json:"last_execution_status"`
				Disabled            bool   `json:"disabled"`
				TriggerCondition    string `json:"trigger_condition"`
			} `json:"action"`
			Log    []string `json:"log"`
			LogURL string   `json:"log_url"`
		} `json:"action_executions"`
	} `json:"execution"`
	RepositoryURL string    `json:"repository_url"`
	Ref           string    `json:"ref"`
	PushID        int       `json:"push_id"`
	Operation     string    `json:"operation"`
	OldHead       string    `json:"old_head"`
	NewHead       string    `json:"new_head"`
	CommitsCount  int       `json:"commits_count"`
	Commits       []struct {
		URL        string    `json:"url"`
		HTMLURL    string    `json:"html_url"`
		Revision   string    `json:"revision"`
		AuthorDate time.Time `json:"author_date"`
		CommitDate time.Time `json:"commit_date"`
		Message    string    `json:"message"`
		Committer  struct {
			URL       string      `json:"url"`
			HTMLURL   string      `json:"html_url"`
			ID        int         `json:"id"`
			Name      string      `json:"name"`
			AvatarURL string      `json:"avatar_url"`
			Title     interface{} `json:"title"`
			Email     string      `json:"email"`
		} `json:"committer"`
		Author struct {
			Title interface{} `json:"title"`
			Email string      `json:"email"`
		} `json:"author"`
	} `json:"commits"`
	PushURL   string `json:"push_url"`
	BranchURL string `json:"branch_url"`
}
