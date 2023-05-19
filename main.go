package main

import (
	"net/http"
	"time"

	"github.com/blackav/ejudge-api/docs"
	"github.com/blackav/ejudge-api/ejudge"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type MasterGetGetSubmitQueryParams struct {
	ContestID int32 `form:"contest_id"`
	Json      int32 `form:"json"`
	SubmitID  int64 `form:"submit_id"`
}

// MasterGetGetSubmit godoc
//
// @Summary		Get information about a submit (privileged)
// @Description	get information about a submit (privileged)
// @ID			master-get-submit
// @Tags		privileged submit
// @Produce		json
// @Param		contest_id query int false "Contest ID"
// @Param		submit_id query int true "Submit ID"
// @Param		json query int false "Submit ID" default(1)
// @Success		200	{object}	ejudge.Reply[ejudge.Submit]
// @Router		/ej/master/get-submit [get]
// @Security	ApiKeyAuth
func MasterGetGetSubmit(c *gin.Context) {
	var q MasterGetGetSubmitQueryParams
	if err := c.Bind(&q); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	var langID int32 = 3
	res := ejudge.Reply[ejudge.Submit]{
		Ok:         true,
		ServerTime: time.Now().Unix(),
		Action:     "get-submit",
		Result: &ejudge.Submit{
			SubmitID:  q.SubmitID,
			ContestID: q.ContestID,
			UserID:    1000,
			ProbID:    25,
			LangID:    &langID,
		},
	}

	c.JSON(http.StatusOK, res)
}

type MasterPostSubmitRunInputResult struct {
	SubmitID int64 `json:"submit_id"`
}

// MasterPostSubmitRunInput godoc
//
// @Summary		Submit a program and its input for testing (privileged)
// @Description	submit a program and its input for testing (privileged)
// @ID			master-submit-run-input
// @Tags		privileged submit
// @Consume     mpfd
// @Produce		json
// @Param		sender_user_login formData string false "Impersonated user login"
// @Param		sender_user_id formData int false "Impersonated user ID"
// @Param		sender_ip formData string false "Fake sender IP"
// @Param		sender_ssl_flag formData int false "HTTPS flag"
// @Param		contest_id formData int false "contest ID"
// @Param		prob_id formData string true "problem ID or short_name"
// @Param		lang_id formData string false "language ID or short_name"
// @Param		eoln_type formData int false "End-of-Line translation type"
// @Param		file formData file false "Source code"
// @Param		text_form formData string false "Source code" format(text)
// @Param		file_input formData file false "Input data"
// @Param		text_form_input formData string false "Input data"
// @Success		200	{object}	ejudge.Reply[MasterPostSubmitRunInputResult]
// @Router		/ej/master/submit-run-input [post]
// @Security	ApiKeyAuth
func MasterPostSubmitRunInput(c *gin.Context) {
}

type PostSubmitRunResult struct {
	RunID   int32  `json:"run_id"`
	RunUUID string `json:"run_uuid,omitempty"`
}

// MasterPostSubmitRun godoc
//
// @Summary		Submit a program for testing and scoring (privileged)
// @Description	submit a program for testing and scoring (privileged)
// @ID			master-submit-run
// @Tags		privileged run
// @Consume     mpfd
// @Produce		json
// @Param		sender_user_login formData string false "Impersonated user login"
// @Param		sender_user_id formData int false "Impersonated user ID"
// @Param		sender_ip formData string false "Fake sender IP"
// @Param		sender_ssl_flag formData int false "HTTPS flag"
// @Param		contest_id formData int false "contest ID"
// @Param		problem_uuid formData string false "UUID of the problem"
// @Param		problem_name formData string false "short_name or internal_name of the problem"
// @Param		problem formData int false "prob_id of the problem"
// @Param		variant formData int false "variant of the problem"
// @Param		language_name formData string false "short_name of the programming language"
// @Param		lang_id formData string false "lang_id or short_name"
// @Param		eoln_type formData int false "End-of-Line translation type"
// @Param		is_visible formData int false "if value > 0 the run will be visible"
// @Param		not_ok_is_cf formData int false "if value > 0 any non OK verdict is Check failed"
// @Param		rejudge_flag formData int false "if value > 0 the run is tested with rejudge priority"
// @Param		file formData file false "Source code"
// @Param		text_form formData string false "Source code" format(text)
// @Success		200	{object}	ejudge.Reply[PostSubmitRunResult]
// @Router		/ej/master/submit-run [post]
// @Security	ApiKeyAuth
func MasterPostSubmitRun(c *gin.Context) {
}

type MasterGetRunStatusJSONResult struct {
	AcceptingMode bool       `json:"accepting_mode,omitempty"`
	Run           ejudge.Run `json:"run,omitempty"`
}

// MasterGetRunStatusJSON godoc
//
// @Summary		Get the info for the run (privileged)
// @Description	Get the info for the run (privileged)
// @ID			master-get-run-status-json
// @Tags		privileged run
// @Produce		json
// @Param		contest_id query int false "contest_id"
// @Param		run_id query int false "run_id"
// @Param		run_uuid query string false "run_uuid"
// @Success		200	{object}	ejudge.Reply[MasterGetRunStatusJSONResult]
// @Router		/ej/master/run-status-json [get]
// @Security	ApiKeyAuth
func MasterGetRunStatusJSON(c *gin.Context) {
}

// MasterGetRawAuditLog godoc
//
// @Summary		Get the raw audit log of the run (privileged)
// @Description	Get the raw audit log of the run (privileged)
// @ID			master-get-raw-audit-log
// @Tags		privileged run
// @Produce		plain
// @Param		contest_id query int false "contest_id"
// @Param		run_id query int false "run_id"
// @Param		run_uuid query string false "run_uuid"
// @Router		/ej/master/raw-audit-log [get]
// @Security	ApiKeyAuth
func MasterGetRawAuditLog(c *gin.Context) {
}

// MasterGetRawReport godoc
//
// @Summary		Get the raw testing report of the run (privileged)
// @Description	Get the raw testing report of the run (privileged)
// @ID			master-get-raw-report
// @Tags		privileged run
// @Param		contest_id query int false "contest_id"
// @Param		run_id query int false "run_id"
// @Param		run_uuid query string false "run_uuid"
// @Router		/ej/master/raw-report [get]
// @Security	ApiKeyAuth
func MasterGetRawReport(c *gin.Context) {
}

// MasterGetDownloadRun godoc
//
// @Summary		Get the source of the run (privileged)
// @Description	Get the source of the run (privileged)
// @ID			master-get-download-run
// @Tags		privileged run
// @Param		contest_id query int false "contest_id"
// @Param		run_id query int false "run_id"
// @Param       no_disp query int false "do not generate content-disposition header"
// @Router		/ej/master/download-run [get]
// @Security	ApiKeyAuth
func MasterGetDownloadRun(c *gin.Context) {
}

type EjudgeContestStatus struct {
	ID                      int32  `json:"id"`
	Name                    string `json:"name,omitempty"`
	NameEn                  string `json:"name_en,omitempty"`
	ServerStartTime         int64  `json:"server_start_time,omitempty"`
	ContestLoadTime         int64  `json:"contest_load_time,omitempty"`
	ScoreSystem             int32  `json:"score_system"`
	IsContinueEnabled       bool   `json:"is_continue_enabled,omitempty"`
	IsVirtual               bool   `json:"is_virtual,omitempty"`
	IsUnlimited             bool   `json:"is_unlimited,omitempty"`
	Duration                int64  `json:"duration,omitempty"`
	BoardFogDuration        int64  `json:"board_fog_duration,omitempty"`
	BoardUnfogDuration      int64  `json:"board_unfog_duration,omitempty"`
	IsRestartable           bool   `json:"is_restartable,omitempty"`
	IsUpsolving             bool   `json:"is_upsolving,omitempty"`
	IsStarted               bool   `json:"is_started,omitempty"`
	IsClientsSuspended      bool   `json:"is_clients_suspended,omitempty"`
	IsTestingSuspended      bool   `json:"is_testing_suspended,omitempty"`
	IsPrintingSuspended     bool   `json:"is_printing_suspended,omitempty"`
	IsSourceClosed          bool   `json:"is_source_closed,omitempty"`
	IsSourceOpen            bool   `json:"is_source_open,omitempty"`
	IsReportClosed          bool   `json:"is_report_closed,omitempty"`
	IsReportOpen            bool   `json:"is_report_open,omitempty"`
	IsJudgeScore            bool   `json:"is_judge_score,omitempty"`
	IsFinalVisibility       bool   `json:"is_final_visibility,omitempty"`
	IsValuerJudgeComments   bool   `json:"is_valuer_judge_comments,omitempty"`
	StartTime               int64  `json:"start_time,omitempty"`
	IsOlympiadAcceptingMode bool   `json:"is_olympiad_accepting_mode,omitempty"`
	IsTestingFinished       bool   `json:"is_testing_finished,omitempty"`
	IsStopped               bool   `json:"is_stopped,omitempty"`
	StopTime                int64  `json:"stop_time,omitempty"`
	IsFreezable             bool   `json:"is_freezable,omitempty"`
	UnfreezeTime            int64  `json:"unfreeze_time,omitempty"`
	IsStandingsUpdated      bool   `json:"is_standings_updated,omitempty"`
	IsFrozen                bool   `json:"is_frozen,omitempty"`
	FreezeTime              int64  `json:"freeze_time,omitempty"`
	ExpectedStopTime        int64  `json:"expected_stop_time,omitempty"`
	SheduledFinishTime      int64  `json:"scheduled_finish_time,omitempty"`
	OpenTime                int64  `json:"open_time,omitempty"`
	CloseTime               int64  `json:"close_time,omitempty"`
	ScheduledStartTime      int64  `json:"scheduled_start_time,omitempty"`
}

type EjudgeContestOnline struct {
	UserCount    int32 `json:"user_count"`
	MaxUserCount int32 `json:"max_user_count"`
	MaxTime      int64 `json:"max_time,omitempty"`
}

type MasterGetContestStatusJSONResult struct {
	Contest EjudgeContestStatus `json:"contest,omitempty"`
	Online  EjudgeContestOnline `json:"online,omitempty"`
}

// MasterGetContestStatusJSON godoc
//
// @Summary		Get the contest status (privileged)
// @Description	Get the contest status (privileged)
// @ID			master-get-contest-status-json
// @Tags		privileged contest
// @Produce		json
// @Param		contest_id query int false "contest_id"
// @Success		200	{object}	ejudge.Reply[MasterGetContestStatusJSONResult]
// @Router		/ej/master/contest-status-json [get]
// @Security	ApiKeyAuth
func MasterGetContestStatusJSON(c *gin.Context) {
}

type MasterGetListRunsJSONResult struct {
	TotalRuns     int32        `json:"total_runs,omitempty"`
	FilteredRuns  int32        `json:"filtered_runs,omitempty"`
	ListedRuns    int32        `json:"listed_runs,omitempty"`
	TransientRuns int32        `json:"transient_runs,omitempty"`
	FilterExpr    string       `json:"filter_expr,omitempty"`
	FirstRun      int32        `json:"first_run,omitempty"`
	LastRun       int32        `json:"last_run,omitempty"`
	FieldMask     int32        `json:"field_mask,omitempty"`
	DisplayedSize int32        `json:"displayed_size,omitempty"`
	DisplayedMask string       `json:"displayed_mask,omitempty"`
	Runs          []ejudge.Run `json:"runs,omitempty"`
}

// MasterGetListRunsJSON godoc
//
// @Summary		List the runs (privileged)
// @Description	List the runs (privileged)
// @ID			master-get-list-runs-json
// @Tags		privileged run
// @Produce		json
// @Param		contest_id query int false "contest_id"
// @Param		filter_expr query string false "Filter expression"
// @Param		first_run query int false "First run to list"
// @Param		last_run query int false "Last run to list"
// @Param		field_mask query int false "Fields to display"
// @Success		200	{object}	ejudge.Reply[MasterGetListRunsJSONResult]
// @Router		/ej/master/list-runs-json [get]
// @Security	ApiKeyAuth
func MasterGetListRunsJSON(c *gin.Context) {
}

type ClientGetContestStatusJSONCompiler struct {
	Id        int32  `json:"id"`
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name,omitempty"`
	SrcSfx    string `json:"src_sfx,omitempty"`
}

type ClientGetContestStatusJSONProblem struct {
	Id        int32  `json:"id"`
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name,omitempty"`
}

type ClientGetContestStatusJSONResult struct {
	Contest   EjudgeContestStatus                  `json:"contest"`
	Online    *EjudgeContestOnline                 `json:"online,omitempty"`
	Compilers []ClientGetContestStatusJSONCompiler `json:"compilers,omitempty"`
	Problems  []ClientGetContestStatusJSONProblem  `json:"problems,omitempty"`
}

// ClientGetContestStatusJSON godoc
//
// @Summary		Get the contest status (unprivileged)
// @Description	Get the contest status (unprivileged)
// @ID			client-get-contest-status-json
// @Tags		unprivileged contest
// @Produce		json
// @Success		200	{object}	ejudge.Reply[ClientGetContestStatusJSONResult]
// @Router		/ej/client/contest-status-json [get]
// @Security	ApiKeyAuth
func ClientGetContestStatusJSON(c *gin.Context) {
}

type ClientGetProblemStatusJSONProblem struct {
	Id                      int32   `json:"id"`
	ShortName               string  `json:"short_name,omitempty"`
	LongName                string  `json:"long_name,omitempty"`
	Type                    int32   `json:"type"`
	FullScore               *int32  `json:"full_score,omitempty"`
	FullUserScore           *int32  `json:"full_user_score,omitempty"`
	MinScore1               *int32  `json:"min_score_1,omitempty"`
	MinScore2               *int32  `json:"min_score_2,omitempty"`
	UseStdin                bool    `json:"use_stdin,omitempty"`
	UseStdout               bool    `json:"use_stdout,omitempty"`
	CombinedStdin           bool    `json:"combined_stdin,omitempty"`
	CombinedStdout          bool    `json:"combined_stdout,omitempty"`
	UseACNotOk              bool    `json:"use_ac_not_ok,omitempty"`
	IgnorePrevAC            bool    `json:"ignore_prev_ac,omitempty"`
	TeamEnableRepView       bool    `json:"team_enable_rep_view,omitempty"`
	TeamEnableCEView        bool    `json:"team_enable_ce_view,omitempty"`
	IgnoreCompileErrors     bool    `json:"ignore_compile_errors,omitempty"`
	DisableUserSubmit       bool    `json:"disable_user_submit,omitempty"`
	DisableTab              bool    `json:"disable_tab,omitempty"`
	EnableSubmitAfterReject bool    `json:"enable_submit_after_reject,omitempty"`
	EnableTokens            bool    `json:"enable_tokens,omitempty"`
	TokensForUserAC         bool    `json:"tokens_for_user_ac,omitempty"`
	DisableSubmitAfterOk    bool    `json:"disable_submit_after_ok,omitempty"`
	DisableTesting          bool    `json:"disable_testing,omitempty"`
	EnableCompilation       bool    `json:"enable_compilation,omitempty"`
	Hidden                  bool    `json:"hidden,omitempty"`
	StandHideTime           bool    `json:"stand_hide_time,omitempty"`
	StandIgnoreScore        bool    `json:"stand_ignore_score,omitempty"`
	StandLastColumn         bool    `json:"stand_last_column,omitempty"`
	DisableStderr           bool    `json:"disable_stderr,omitempty"`
	RealTimeLimitMs         int32   `json:"real_time_limit_ms"`
	TimeLimitMs             int32   `json:"time_limit_ms"`
	ACMRunPenalty           *int32  `json:"acm_run_penalty,omitempty"`
	TestScore               *int32  `json:"test_score,omitempty"`
	RunPenalty              *int32  `json:"run_penalty,omitempty"`
	DisqualifiedPenalty     *int32  `json:"disqualified_penalty,omitempty"`
	CompileErrorPenalty     *int32  `json:"compile_error_penalty,omitempty"`
	TestsToAccept           *int32  `json:"tests_to_accept,omitempty"`
	MinTestsToAccept        *int32  `json:"min_tests_to_accept,omitempty"`
	ScoreMultiplier         int32   `json:"score_multiplier,omitempty"`
	MaxUserRunCount         int32   `json:"max_user_run_count,omitempty"`
	StandName               string  `json:"stand_name,omitempty"`
	StandColumn             string  `json:"stand_column,omitempty"`
	InputFile               string  `json:"input_file,omitempty"`
	OutputFile              string  `json:"output_file,omitempty"`
	OkStatus                string  `json:"ok_status,omitempty"`
	StartDate               int64   `json:"start_date,omitempty"`
	Compilers               []int32 `json:"compilers,omitempty"`
	EnableMaxStackSize      bool    `json:"enable_max_stack_size,omitempty"`
	MaxVMSize               int64   `json:"max_vm_size,omitempty"`
	MaxStackSize            int64   `json:"max_stack_size,omitempty"`
	MaxRSSSize              int64   `json:"max_rss_size,omitempty"`
	IsStatementAvailable    bool    `json:"is_statement_avaiable,omitempty"`
	EstStmtSize             int64   `json:"est_stmt_size,omitempty"`
}

type ClientGetProblemStatusJSONProblemStatus struct {
	IsViewable        bool   `json:"is_viewable,omitempty"`
	IsSubmittable     bool   `json:"is_submittable,omitempty"`
	IsTabable         bool   `json:"is_tabable,omitempty"`
	IsSolved          bool   `json:"is_solved,omitempty"`
	IsAccepted        bool   `json:"is_accepted,omitempty"`
	IsPending         bool   `json:"is_pending,omitempty"`
	IsPendingReview   bool   `json:"is_pending_review,omitempty"`
	IsTransient       bool   `json:"is_transient,omitempty"`
	IsLastUntokenized bool   `json:"is_last_untokenized,omitempty"`
	IsMarked          bool   `json:"is_marked,omitempty"`
	IsAutoOk          bool   `json:"is_autook,omitempty"`
	IsRejected        bool   `json:"is_rejected,omitempty"`
	IsEffTimeNeeded   bool   `json:"is_eff_time_needed,omitempty"`
	BestRun           *int32 `json:"best_run,omitempty"`
	Attempts          int32  `json:"attempts,omitempty"`
	Disqualified      int32  `json:"disqualified,omitempty"`
	CEAttempts        int32  `json:"ce_attempts,omitempty"`
	BestScore         int32  `json:"best_score,omitempty"`
	PrevSuccesses     int32  `json:"prev_successes,omitempty"`
	AllAttempts       int32  `json:"all_attempts,omitempty"`
	EffAttempts       int32  `json:"eff_attempts,omitempty"`
	TokenCount        int32  `json:"token_count,omitempty"`
	Deadline          int64  `json:"deadline,omitempty"`
	EffectiveTime     int64  `json:"effective_time,omitempty"`
}

type ClientGetProblemStatusJSONResult struct {
	Problem       ClientGetProblemStatusJSONProblem       `json:"problem"`
	ProblemStatus ClientGetProblemStatusJSONProblemStatus `json:"problem_status"`
}

// ClientGetProblemStatusJSON godoc
//
// @Summary		Get the problem status (unprivileged)
// @Description	Get the problem status (unprivileged)
// @ID			client-get-problem-status-json
// @Tags		unprivileged problem
// @Produce		json
// @Param		problem query int false "Problem ID"
// @Success		200	{object}	ejudge.Reply[ClientGetProblemStatusJSONResult]
// @Router		/ej/client/problem-status-json [get]
// @Security	ApiKeyAuth
func ClientGetProblemStatusJSON(c *gin.Context) {
}

// ClientGetProblemStatementJSON godoc
//
// @Summary		Get the problem statement (unprivileged)
// @Description	Get the problem statement (unprivileged)
// @ID			client-get-problem-statement-json
// @Tags		unprivileged problem
// @Produce		html
// @Param		problem query int false "Problem ID"
// @Router		/ej/client/problem-statement-json [get]
// @Security	ApiKeyAuth
func ClientGetProblemStatementJSON(c *gin.Context) {
}

// ClientPostSubmitRun godoc
//
// @Summary		Submit a program for testing and scoring (unprivileged)
// @Description	submit a program for testing and scoring (unprivileged)
// @ID			client-submit-run
// @Tags		unprivileged run
// @Consume     mpfd
// @Produce		json
// @Param		prob_id formData string false "prob_id or the short_name of the problem"
// @Param		lang_id formData string false "lang_id or short_name"
// @Param		eoln_type formData int false "End-of-Line translation type"
// @Param		file formData file false "Source code"
// @Param		text_form formData string false "Source code" format(text)
// @Success		200	{object}	ejudge.Reply[PostSubmitRunResult]
// @Router		/ej/client/submit-run [post]
// @Security	ApiKeyAuth
func ClientPostSubmitRun(c *gin.Context) {
}

type ClientGetListRunsJSONRun struct {
	RunId       int32  `json:"run_id"`
	ProbID      int32  `json:"prob_id"`
	RunTime     int64  `json:"run_time"`
	Status      int32  `json:"status"`
	FailedTest  *int32 `json:"failed_test,omitempty"`
	PassedTests *int32 `json:"passed_tests,omitempty"`
	Score       *int32 `json:"score,omitempty"`
}

type ClientGetListRunsJSONResult struct {
	Runs []ClientGetListRunsJSONRun `json:"runs,omitempty"`
}

// ClientGetListRunsJSON godoc
//
// @Summary		Get the list of the runs (unprivileged)
// @Description	Get the list of the runs (unprivileged)
// @ID			client-get-list-runs-json
// @Tags		unprivileged run
// @Produce		json
// @Param		problem query int false "Problem ID"
// @Success		200	{object}	ejudge.Reply[ClientGetListRunsJSONResult]
// @Router		/ej/client/list-runs-json [get]
// @Security	ApiKeyAuth
func ClientGetListRunsJSON(c *gin.Context) {
}

type ClientGetRunStatusJSONRun struct {
	RunID                     int32  `json:"run_id"`
	ProbID                    int32  `json:"prob_id"`
	RunTimeUs                 int64  `json:"run_time_us"`
	RunTime                   int64  `json:"run_time"`
	Duration                  int64  `json:"duration"`
	LangID                    int32  `json:"lang_id,omitempty"`
	UserID                    int32  `json:"user_id,omitempty"`
	Size                      int32  `json:"size"`
	Status                    int32  `json:"status"`
	IsImported                bool   `json:"is_imported,omitempty"`
	IsHidden                  bool   `json:"is_hidden,omitempty"`
	IsWithDuration            bool   `json:"is_with_duration,omitempty"`
	IsStandardProblem         bool   `json:"is_standard_problem,omitempty"`
	IsMinimalReport           bool   `json:"is_minimal_report,omitempty"`
	IsWithEffectiveTime       bool   `json:"is_with_effective_time,omitempty"`
	EffectiveTime             *int64 `json:"effective_time,omitempty"`
	IsSrcEnabled              bool   `json:"is_src_enabled,omitempty"`
	SrcSfx                    string `json:"src_sfx,omitempty"`
	IsReportEnabled           bool   `json:"is_report_enabled,omitempty"`
	IsFailedTestAvailable     bool   `json:"is_failed_test_available,omitempty"`
	FailedTest                *int32 `json:"failed_test,omitempty"`
	IsPassedTestsAvailable    bool   `json:"is_passed_tests_available,omitempty"`
	PassedTests               *int32 `json:"passed_tests,omitempty"`
	IsScoreAvailable          bool   `json:"is_score_available,omitempty"`
	Score                     *int32 `json:"score,omitempty"`
	ScoreStr                  string `json:"score_str,omitempty"`
	MessageCount              int32  `json:"message_count,omitempty"`
	IsCompilerOutputAvailable bool   `json:"is_compiler_output_available,omitempty"`
	IsReportAvailable         bool   `json:"is_report_available,omitempty"`
}

type EjudgeEncodedContent struct {
	Method int32  `json:"method"`
	Size   int32  `json:"size"`
	Data   string `json:"data"`
}

type ClientGetRunStatusJSONCompilerOutput struct {
	Content EjudgeEncodedContent `json:"content"`
}

type ClientGetRunStatusJSONValuerComment struct {
	Content EjudgeEncodedContent `json:"content"`
}

type ClientGetRunStatusJSONValuerJudgeComment struct {
	Content EjudgeEncodedContent `json:"content"`
}

type ClientGetRunStatusJSONTestArgs struct {
	IsTooBig bool   `json:"is_too_big,omitempty"`
	Size     *int32 `json:"size,omitempty"`
}

type ClientGetRunStatusJSONTestBrief struct {
	IsTooBig bool   `json:"is_too_big,omitempty"`
	Size     *int32 `json:"size,omitempty"`
	IsBinary bool   `json:"is_binary,omitempty"`
}

type ClientGetRunStatusJSONTest struct {
	Num                int32                            `json:"num"`
	IsVisibilityExists bool                             `json:"is_visibility_exists,omitempty"`
	Status             int32                            `json:"status,omitempty"`
	TimeMs             *int64                           `json:"time_ms,omitempty"`
	ExitComment        string                           `json:"exit_comment,omitempty"`
	TermSignal         *int32                           `json:"term_signal,omitempty"`
	ExitCode           *int32                           `json:"exit_code,omitempty"`
	Score              *int32                           `json:"score,omitempty"`
	MaxScore           *int32                           `json:"max_score,omitempty"`
	TeamComment        string                           `json:"team_comment,omitempty"`
	IsVisibilityFull   bool                             `json:"is_visibility_full,omitempty"`
	Args               *ClientGetRunStatusJSONTestArgs  `json:"args,omitempty"`
	Input              *ClientGetRunStatusJSONTestBrief `json:"input,omitempty"`
	Output             *ClientGetRunStatusJSONTestBrief `json:"output,omitempty"`
	Correct            *ClientGetRunStatusJSONTestBrief `json:"correct,omitempty"`
	Error              *ClientGetRunStatusJSONTestBrief `json:"error,omitempty"`
	Checker            *ClientGetRunStatusJSONTestBrief `json:"checker,omitempty"`
}

type ClientGetRunStatusJSONTestingReport struct {
	ValuerComment      *ClientGetRunStatusJSONValuerComment      `json:"valuer_comment,omitempty"`
	ValuerJudgeComment *ClientGetRunStatusJSONValuerJudgeComment `json:"valuer_judge_comment,omitempty"`
	Tests              []ClientGetRunStatusJSONTest              `json:"tests"`
}

type ClientGetRunStatusJSONResult struct {
	Run            ClientGetRunStatusJSONRun             `json:"run"`
	CompilerOutput *ClientGetRunStatusJSONCompilerOutput `json:"compiler_output,omitempty"`
	TestingReport  *ClientGetRunStatusJSONTestingReport  `json:"testing_report,omitempty"`
}

// ClientGetRunStatusJSON godoc
//
// @Summary		Get the status of the run (unprivileged)
// @Description	Get the status of the run (unprivileged)
// @ID			client-get-run-status-json
// @Tags		unprivileged run
// @Produce		json
// @Param		run_id query int false "Run ID"
// @Success		200	{object}	ejudge.Reply[ClientGetRunStatusJSONResult]
// @Router		/ej/client/run-status-json [get]
// @Security	ApiKeyAuth
func ClientGetRunStatusJSON(c *gin.Context) {
}

type ClientGetRunMessagesJSONMessage struct {
	ClarID  int32                `json:"clar_id"`
	Size    int32                `json:"size"`
	TimeUs  int64                `json:"time_us"`
	From    int32                `json:"from"`
	To      int32                `json:"to"`
	Subject string               `json:"subject,omitempty"`
	Content EjudgeEncodedContent `json:"content"`
}

type ClientGetRunMessagesJSONResult struct {
	Messages []ClientGetRunMessagesJSONMessage `json:"messages,omitempty"`
}

// ClientGetRunMessagesJSON godoc
//
// @Summary		Get the messages associated with the run (unprivileged)
// @Description	Get the messages associated with the run (unprivileged)
// @ID			client-get-run-messages-json
// @Tags		unprivileged run
// @Produce		json
// @Param		run_id query int false "Run ID"
// @Success		200	{object}	ejudge.Reply[ClientGetRunMessagesJSONResult]
// @Router		/ej/client/run-messages-json [get]
// @Security	ApiKeyAuth
func ClientGetRunMessagesJSON(c *gin.Context) {
}

// ClientGetDownloadRun godoc
//
// @Summary		Get the source of the run (unprivileged)
// @Description	Get the source of the run (unprivileged)
// @ID			client-get-download-run
// @Tags		unprivileged run
// @Param		run_id query int false "run_id"
// @Param       no_disp query int false "do not generate content-disposition header"
// @Router		/ej/client/download-run [get]
// @Security	ApiKeyAuth
func ClientGetDownloadRun(c *gin.Context) {
}

// ClientGetDownloadRunFile godoc
//
// @Summary		Get the input/output/correct file of the run (unprivileged)
// @Description	Get the input/output/correct of the run (unprivileged)
// @ID			client-get-download-run-file
// @Tags		unprivileged run
// @Param		run_id query int true "run_id"
// @Param		num query int true "test number"
// @Param		index query int true "file index"
// @Param		offset query int false "initial offset of the file fragment"
// @Param		length query int false "length of the file fragment"
// @Router		/ej/client/download-run-file [get]
// @Security	ApiKeyAuth
func ClientGetDownloadRunFile(c *gin.Context) {
}

// ClientGetGetSubmit godoc
//
// @Summary		Get information about a submit (unprivileged)
// @Description	get information about a submit (unprivileged)
// @ID			client-get-submit
// @Tags		unprivileged submit
// @Produce		json
// @Param		submit_id query int true "Submit ID"
// @Param		json query int false "Submit ID" default(1)
// @Success		200	{object}	ejudge.Reply[ejudge.Submit]
// @Router		/ej/client/get-submit [get]
// @Security	ApiKeyAuth
func ClientGetGetSubmit(c *gin.Context) {
}

type ClientPostSubmitRunInputResult struct {
	SubmitID int64 `json:"submit_id"`
}

// ClientPostSubmitRunInput godoc
//
// @Summary		Submit a program and its input for testing (unprivileged)
// @Description	submit a program and its input for testing (unprivileged)
// @ID			client-submit-run-input
// @Tags		unprivileged submit
// @Consume     mpfd
// @Produce		json
// @Param		prob_id formData string true "problem ID or short_name"
// @Param		lang_id formData string true "language ID or short_name"
// @Param		eoln_type formData int false "End-of-Line translation type"
// @Param		file formData file false "Source code"
// @Param		text_form formData string false "Source code" format(text)
// @Param		file_input formData file false "Input data"
// @Param		text_form_input formData string false "Input data"
// @Success		200	{object}	ejudge.Reply[ClientPostSubmitRunInputResult]
// @Router		/ej/client/submit-run-input [post]
// @Security	ApiKeyAuth
func ClientPostSubmitRunInput(c *gin.Context) {
}

type EjudgeUserProb struct {
	SerialID         int64  `json:"serial_id"`
	CreateTimeMs     int64  `json:"create_time_ms"`
	LastChangeTimeMs int64  `json:"last_change_time_ms,omitempty"`
	ContestID        int32  `json:"contest_id"`
	UserID           int32  `json:"user_id"`
	ProbID           int32  `json:"prob_id"`
	LangName         string `json:"lang_name,omitempty"`
	HookID           string `json:"hook_id,omitempty"`
	GitlabToken      string `json:"gitlab_token,omitempty"`
	VCSType          string `json:"vcs_type,omitempty"`
	VCSURL           string `json:"vcs_url,omitempty"`
	VCSSubdir        string `json:"vcs_subdir,omitempty"`
	VCSBranchSpec    string `json:"vcs_branch_spec,omitempty"`
	SshPrivateKey    string `json:"ssh_private_key,omitempty"`
}

// ClientGetGetUserprob godoc
//
// @Summary		Get VCS problem settings (unprivileged)
// @Description	get VCS problem settings (unprivileged)
// @ID			client-get-get-userprob
// @Tags		unprivileged vcs
// @Produce		json
// @Param		prob_id query int true "Problem ID"
// @Param		json query int false "Submit ID" default(1)
// @Success		200	{object}	ejudge.Reply[EjudgeUserProb]
// @Router		/ej/client/get-userprob [get]
// @Security	ApiKeyAuth
func ClientGetGetUserprob(c *gin.Context) {
}

// ClientPostCreateUserprob godoc
//
// @Summary		Create VCS problem settings (unprivileged)
// @Description	create VCS problem settings (unprivileged)
// @ID			client-post-create-userprob
// @Tags		unprivileged vcs
// @Produce		json
// @Param		prob_id query int true "Problem ID"
// @Param		json query int false "Submit ID" default(1)
// @Success		200	{object}	ejudge.Reply[EjudgeUserProb]
// @Router		/ej/client/create-userprob [post]
// @Security	ApiKeyAuth
func ClientPostCreateUserprob(c *gin.Context) {
}

// ClientPostSaveUserprob godoc
//
// @Summary		Save VCS problem settings (unprivileged)
// @Description	save VCS problem settings (unprivileged)
// @ID			client-post-save-userprob
// @Tags		unprivileged vcs
// @Produce		json
// @Param		serial_id query int true "Serial ID of the properties"
// @Param		vcs_type query string true "VCS type: gitlab | github"
// @Param		lang_name query string true "Language name"
// @Param		vcs_url query string true "VCS URL"
// @Param		vcs_subdir query string true "VCS subdir"
// @Param		vcs_branch_spec query string true "VCS branch"
// @Param		ssh_private_key query string true "ssh private key"
// @Param		json query int false "Submit ID" default(1)
// @Success		200	{object}	ejudge.BaseReply
// @Router		/ej/client/save-userprob [post]
// @Security	ApiKeyAuth
func ClientPostSaveUserprob(c *gin.Context) {
}

// ClientPostRemoveUserprob godoc
//
// @Summary		Remove VCS problem settings (unprivileged)
// @Description	remove VCS problem settings (unprivileged)
// @ID			client-post-remove-userprob
// @Tags		unprivileged vcs
// @Produce		json
// @Param		serial_id query int true "Serial ID of the properties"
// @Param		json query int false "Submit ID" default(1)
// @Success		200	{object}	ejudge.BaseReply
// @Router		/ej/client/remove-userprob [post]
// @Security	ApiKeyAuth
func ClientPostRemoveUserprob(c *gin.Context) {
}

// @title       Ejudge API
// @version		3.10.3
// @description	The ejudge API
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description					API Key Authorization
func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	ej := router.Group("/ej")
	{
		master := ej.Group("/master")
		{
			master.GET("/contest-status-json", MasterGetContestStatusJSON)
			master.GET("/download-run", MasterGetDownloadRun)
			master.GET("/get-submit", MasterGetGetSubmit)
			master.GET("/list-runs-json", MasterGetListRunsJSON)
			master.GET("/raw-audit-log", MasterGetRawAuditLog)
			master.GET("/raw-report", MasterGetRawReport)
			master.GET("/run-status-json", MasterGetRunStatusJSON)
			master.POST("/submit-run", MasterPostSubmitRun)
			master.POST("/submit-run-input", MasterPostSubmitRunInput)
		}
		client := ej.Group("/client")
		{
			client.GET("/contest-status-json", ClientGetContestStatusJSON)
			client.GET("/download-run", ClientGetDownloadRun)
			client.GET("/download-run-file", ClientGetDownloadRunFile)
			client.GET("/get-submit", ClientGetGetSubmit)
			client.GET("/get-userprob", ClientGetGetUserprob)
			client.GET("/list-runs-json", ClientGetListRunsJSON)
			client.GET("/problem-statement-json", ClientGetProblemStatementJSON)
			client.GET("/problem-status-json", ClientGetProblemStatusJSON)
			client.GET("/run-messages-json", ClientGetRunMessagesJSON)
			client.GET("/run-status-json", ClientGetRunStatusJSON)
			client.POST("/create-userprob", ClientPostCreateUserprob)
			client.POST("/remove-userprob", ClientPostRemoveUserprob)
			client.POST("/save-userprob", ClientPostSaveUserprob)
			client.POST("/submit-run", ClientPostSubmitRun)
			client.POST("/submit-run-input", ClientPostSubmitRunInput)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8889")
}
