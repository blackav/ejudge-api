package ejudge

type Submit struct {
	SubmitID          int64  `json:"submit_id"`
	UserID            int32  `json:"user_id"`
	ContestID         int32  `json:"contest_id"`
	ProbID            int32  `json:"prob_id"`
	LangID            *int32 `json:"lang_id,omitempty"`
	Status            int32  `json:"status"`
	StatusStr         string `json:"status_str"`
	StatusDesc        string `json:"status_desc,omitempty"`
	CompilerOutput    string `json:"compiler_output,omitempty"`
	TestCheckerOutput string `json:"test_checker_output,omitempty"`
	Time              *int64 `json:"time,omitempty"`
	RealTime          *int64 `json:"real_time,omitempty"`
	ExitCode          *int32 `json:"exit_code,omitempty"`
	TermSignal        *int32 `json:"term_signal,omitempty"`
	MaxMemoryUsed     *int64 `json:"max_memory_used,omitempty"`
	MaxRSS            *int64 `json:"max_rss,omitempty"`
	Input             string `json:"input,omitempty"`
	Output            string `json:"output,omitempty"`
	Error             string `json:"error,omitempty"`
        ExtUserKind       string `json:"ext_user_kind,omitempty"`
        ExtUser           string `json:"ext_user,omitempty"`
        NotifyDriver      int32  `json:"notify_driver,omitempty"`
        NotifyKind        string `json:"notify_kind,omitempty"`
        NotifyQueue       string `json:"notify_queue,omitempty"`
}

type Run struct {
	RunID            int32  `json:"run_id"`
	ContestID        int32  `json:"contest_id,omitempty"`
	RunUUID          string `json:"run_uuid,omitempty"`
	SerialID         int64  `json:"serial_id,omitempty"`
	Status           int32  `json:"status"`
	StatusStr        string `json:"status_str,omitempty"`
	StatusDesc       string `json:"status_desc,omitempty"`
	RunTime          int64  `json:"run_time,omitempty"`
	NSec             int32  `json:"nsec,omitempty"`
	RunTimeUs        int64  `json:"run_time_us,omitempty"`
	Duration         int32  `json:"duration,omitempty"`
	UserID           int32  `json:"user_id,omitempty"`
	UserLogin        string `json:"user_login,omitempty"`
	UserName         string `json:"user_name,omitempty"`
	ProbID           int32  `json:"prob_id,omitempty"`
	ProbName         string `json:"prob_name,omitempty"`
	ProbInternalName string `json:"prob_internal_name,omitempty"`
	ProbUUID         string `json:"prob_uuid,omitempty"`
	Variant          int32  `json:"variant,omitempty"`
	RawVariant       int32  `json:"raw_variant,omitempty"`
	LangID           int32  `json:"lang_id,omitempty"`
	LangName         string `json:"lang_name,omitempty"`
	IP               string `json:"ip,omitempty"`
	SSLFlag          bool   `json:"ssl_flag,omitempty"`
	IPV6Flag         bool   `json:"ipv6_flag,omitempty"`
	SHA1             string `json:"sha1,omitempty"`
	SHA256           string `json:"sha256,omitempty"`
	LocaleID         int32  `json:"locale_id,omitempty"`
	EOLNType         int32  `json:"eoln_type,omitempty"`
	MIMEType         string `json:"mime_type,omitempty"`
	Size             int64  `json:"size,omitempty"`
	StoreFlags       int32  `json:"store_flags,omitempty"`
	IsImported       bool   `json:"is_imported,omitempty"`
	IsHidden         bool   `json:"is_hidden,omitempty"`
	IsReadOnly       bool   `json:"is_readonly,omitempty"`
	PassedMode       bool   `json:"passed_mode,omitempty"`
	Score            *int32 `json:"score,omitempty"`
	Test             *int32 `json:"test,omitempty"`
	IsMarked         bool   `json:"is_marked,omitempty"`
	ScoreAdj         int32  `json:"score_adj,omitempty"`
	JudgeID          int32  `json:"judge_id,omitempty"`
	JudgeUUID        string `json:"judge_uuid,omitempty"`
	Pages            int32  `json:"pages,omitempty"`
	TokenFlags       int32  `json:"token_flags,omitempty"`
	TokenCount       int32  `json:"token_count,omitempty"`
	IsSaved          bool   `json:"is_saved,omitempty"`
	SavedStatus      *int32 `json:"saved_status,omitempty"`
	SavedStatusStr   string `json:"saved_status_str,omitempty"`
	SavedScore       *int32 `json:"saved_score,omitempty"`
	SavedTest        *int32 `json:"saved_test,omitempty"`
	IsChecked        bool   `json:"is_checked,omitempty"`
	IsVCS            bool   `json:"is_vcs,omitempty"`
	VerdictBits      uint32 `json:"verdict_bits,omitempty"`
	LastChangeUs     int64  `json:"last_change_us,omitempty"`
	TestsPassed      *int32 `json:"tests_passed,omitempty"`
	RawScore         *int32 `json:"raw_score,omitempty"`
	ScoreStr         string `json:"score_str,omitempty"`
        ExtUserKind      string `json:"ext_user_kind,omitempty"`
        ExtUser          string `json:"ext_user,omitempty"`
        NotifyDriver     int32  `json:"notify_driver,omitempty"`
        NotifyKind       string `json:"notify_kind,omitempty"`
        NotifyQueue      string `json:"notify_queue,omitempty"`
}
