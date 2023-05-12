package main

import (
	"net/http"
	"time"

	"github.com/blackav/ejudge-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type MasterGetGetSubmitQueryParams struct {
	ContestID int32 `form:"contest_id"`
	Json      int32 `form:"json"`
	SubmitID  int64 `form:"submit_id"`
}

type ReplyError struct {
	Num     int64  `json:"num"`
	Symbol  string `json:"symbol,omitempty"`
	Message string `json:"message,omitempty"`
	LogID   string `json:"log_id,omitempty"`
}

type EjudgeSubmit struct {
	SubmitID          int64  `json:"submit_id"`
	UserID            int32  `json:"user_id"`
	ContestID         int32  `json:"contest_id"`
	ProbID            int32  `json:"prob_id"`
	LangID            *int32 `json:"lang_id,omitempty"`
	Status            int32  `json:"status"`
	StatusStr         string `json:"status_str"`
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
}

type MasterGetGetSubmitReply struct {
	Ok         bool          `json:"ok"`
	ServerTime int64         `json:"server_time"`
	Action     string        `json:"action"`
	Error      *ReplyError   `json:"error,omitempty"`
	Result     *EjudgeSubmit `json:"result,omitempty"`
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
// @Success		200	{object}	MasterGetGetSubmitReply
// @Router		/ej/master/get-submit [get]
// @Security	ApiKeyAuth
func MasterGetGetSubmit(c *gin.Context) {
	var q MasterGetGetSubmitQueryParams
	if err := c.Bind(&q); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	var langID int32 = 3
	res := MasterGetGetSubmitReply{
		Ok:         true,
		ServerTime: time.Now().Unix(),
		Action:     "get-submit",
		Result: &EjudgeSubmit{
			SubmitID:  q.SubmitID,
			ContestID: q.ContestID,
			UserID:    1000,
			ProbID:    25,
			LangID:    &langID,
		},
	}

	c.JSON(http.StatusOK, res)
}

type MasterPostSubmitRunInputReplyResult struct {
	SubmitID int64 `json:"submit_id"`
}

type MasterPostSubmitRunInputReply struct {
	Ok         bool                                 `json:"ok"`
	ServerTime int64                                `json:"server_time"`
	Action     string                               `json:"action"`
	Error      *ReplyError                          `json:"error,omitempty"`
	Result     *MasterPostSubmitRunInputReplyResult `json:"result,omitempty"`
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
// @Success		200	{object}	MasterPostSubmitRunInputReply
// @Router		/ej/master/submit-run-input [post]
// @Security	ApiKeyAuth
func MasterPostSubmitRunInput(c *gin.Context) {
}

type MasterPostSubmitRunReplyResult struct {
	RunID   int32  `json:"run_id"`
	RunUUID string `json:"run_uuid,omitempty"`
}

type MasterPostSubmitRunReply struct {
	Ok         bool                            `json:"ok"`
	ServerTime int64                           `json:"server_time"`
	Action     string                          `json:"action"`
	Error      *ReplyError                     `json:"error,omitempty"`
	Result     *MasterPostSubmitRunReplyResult `json:"result,omitempty"`
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
// @Success		200	{object}	MasterPostSubmitRunReply
// @Router		/ej/master/submit-run [post]
// @Security	ApiKeyAuth
func MasterPostSubmitRun(c *gin.Context) {
}

type EjudgeRun struct {
	RunID            int32  `json:"run_id"`
	ContestID        int32  `json:"contest_id,omitempty"`
	RunUUID          string `json:"run_uuid,omitempty"`
	SerialID         int64  `json:"serial_id,omitempty"`
	Status           int32  `json:"status"`
	StatusStr        string `json:"status_str,omitempty"`
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
}

type MasterGetRunStatusJSONReplyResult struct {
	AcceptingMode bool      `json:"accepting_mode,omitempty"`
	Run           EjudgeRun `json:"run,omitempty"`
}

type MasterGetRunStatusJSONReply struct {
	Ok         bool                               `json:"ok"`
	ServerTime int64                              `json:"server_time"`
	Action     string                             `json:"action"`
	Error      *ReplyError                        `json:"error,omitempty"`
	Result     *MasterGetRunStatusJSONReplyResult `json:"result,omitempty"`
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
// @Success		200	{object}	MasterGetRunStatusJSONReply
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

type MasterGetContestStatusJSONReplyResultContest struct {
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

type MasterGetContestStatusJSONReplyResultOnline struct {
	UserCount    int32 `json:"user_count"`
	MaxUserCount int32 `json:"max_user_count"`
	MaxTime      int64 `json:"max_time,omitempty"`
}

type MasterGetContestStatusJSONReplyResult struct {
	Contest MasterGetContestStatusJSONReplyResultContest `json:"contest,omitempty"`
	Online  MasterGetContestStatusJSONReplyResultOnline  `json:"online,omitempty"`
}

type MasterGetContestStatusJSONReply struct {
	Ok         bool                                   `json:"ok"`
	ServerTime int64                                  `json:"server_time"`
	Action     string                                 `json:"action"`
	Error      *ReplyError                            `json:"error,omitempty"`
	Result     *MasterGetContestStatusJSONReplyResult `json:"result,omitempty"`
}

// MasterGetContestStatusJSON godoc
//
// @Summary		Get the contest status (privileged)
// @Description	Get the contest status (privileged)
// @ID			master-get-contest-status-json
// @Tags		privileged contest
// @Produce		json
// @Param		contest_id query int false "contest_id"
// @Success		200	{object}	MasterGetContestStatusJSONReply
// @Router		/ej/master/contest-status-json [get]
// @Security	ApiKeyAuth
func MasterGetContestStatusJSON(c *gin.Context) {
}

type MasterGetListRunsJSONReplyResult struct {
	TotalRuns     int32       `json:"total_runs,omitempty"`
	FilteredRuns  int32       `json:"filtered_runs,omitempty"`
	ListedRuns    int32       `json:"listed_runs,omitempty"`
	TransientRuns int32       `json:"transient_runs,omitempty"`
	FilterExpr    string      `json:"filter_expr,omitempty"`
	FirstRun      int32       `json:"first_run,omitempty"`
	LastRun       int32       `json:"last_run,omitempty"`
	FieldMask     int32       `json:"field_mask,omitempty"`
	DisplayedSize int32       `json:"displayed_size,omitempty"`
	DisplayedMask string      `json:"displayed_mask,omitempty"`
	Runs          []EjudgeRun `json:"runs,omitempty"`
}

type MasterGetListRunsJSONReply struct {
	Ok         bool                              `json:"ok"`
	ServerTime int64                             `json:"server_time"`
	Action     string                            `json:"action"`
	Error      *ReplyError                       `json:"error,omitempty"`
	Result     *MasterGetListRunsJSONReplyResult `json:"result,omitempty"`
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
// @Success		200	{object}	MasterGetContestStatusJSONReply
// @Router		/ej/master/list-runs-json [get]
// @Security	ApiKeyAuth
func MasterGetListRunsJSON(c *gin.Context) {
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
			master.GET("/get-submit", MasterGetGetSubmit)
			master.POST("/submit-run-input", MasterPostSubmitRunInput)
			master.POST("/submit-run", MasterPostSubmitRun)
			master.GET("/run-status-json", MasterGetRunStatusJSON)
			master.GET("/raw-audit-log", MasterGetRawAuditLog)
			master.GET("/raw-report", MasterGetRawReport)
			master.GET("/contest-status-json", MasterGetContestStatusJSON)
			master.GET("/list-runs-json", MasterGetListRunsJSON)
			master.GET("/download-run", MasterGetDownloadRun)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8889")
}
