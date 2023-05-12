// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ej/master/contest-status-json": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the contest status (privileged)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "privileged contest"
                ],
                "summary": "Get the contest status (privileged)",
                "operationId": "master-get-contest-status-json",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "contest_id",
                        "name": "contest_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.MasterGetContestStatusJSONReply"
                        }
                    }
                }
            }
        },
        "/ej/master/download-run": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the source of the run (privileged)",
                "tags": [
                    "privileged run"
                ],
                "summary": "Get the source of the run (privileged)",
                "operationId": "master-get-download-run",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "contest_id",
                        "name": "contest_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "run_id",
                        "name": "run_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "do not generate content-disposition header",
                        "name": "no_disp",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/ej/master/get-submit": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get information about a submit (privileged)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "privileged submit"
                ],
                "summary": "Get information about a submit (privileged)",
                "operationId": "master-get-submit",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Contest ID",
                        "name": "contest_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Submit ID",
                        "name": "submit_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Submit ID",
                        "name": "json",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.MasterGetGetSubmitReply"
                        }
                    }
                }
            }
        },
        "/ej/master/list-runs-json": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List the runs (privileged)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "privileged run"
                ],
                "summary": "List the runs (privileged)",
                "operationId": "master-get-list-runs-json",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "contest_id",
                        "name": "contest_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter expression",
                        "name": "filter_expr",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "First run to list",
                        "name": "first_run",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Last run to list",
                        "name": "last_run",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Fields to display",
                        "name": "field_mask",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.MasterGetContestStatusJSONReply"
                        }
                    }
                }
            }
        },
        "/ej/master/raw-audit-log": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the raw audit log of the run (privileged)",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "privileged run"
                ],
                "summary": "Get the raw audit log of the run (privileged)",
                "operationId": "master-get-raw-audit-log",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "contest_id",
                        "name": "contest_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "run_id",
                        "name": "run_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "run_uuid",
                        "name": "run_uuid",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/ej/master/raw-report": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the raw testing report of the run (privileged)",
                "tags": [
                    "privileged run"
                ],
                "summary": "Get the raw testing report of the run (privileged)",
                "operationId": "master-get-raw-report",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "contest_id",
                        "name": "contest_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "run_id",
                        "name": "run_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "run_uuid",
                        "name": "run_uuid",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/ej/master/run-status-json": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the info for the run (privileged)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "privileged run"
                ],
                "summary": "Get the info for the run (privileged)",
                "operationId": "master-get-run-status-json",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "contest_id",
                        "name": "contest_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "run_id",
                        "name": "run_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "run_uuid",
                        "name": "run_uuid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.MasterGetRunStatusJSONReply"
                        }
                    }
                }
            }
        },
        "/ej/master/submit-run": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "submit a program for testing and scoring (privileged)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "privileged run"
                ],
                "summary": "Submit a program for testing and scoring (privileged)",
                "operationId": "master-submit-run",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Impersonated user login",
                        "name": "sender_user_login",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Impersonated user ID",
                        "name": "sender_user_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Fake sender IP",
                        "name": "sender_ip",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "HTTPS flag",
                        "name": "sender_ssl_flag",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "contest ID",
                        "name": "contest_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "UUID of the problem",
                        "name": "problem_uuid",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "short_name or internal_name of the problem",
                        "name": "problem_name",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "prob_id of the problem",
                        "name": "problem",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "variant of the problem",
                        "name": "variant",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "short_name of the programming language",
                        "name": "language_name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "lang_id or short_name",
                        "name": "lang_id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "End-of-Line translation type",
                        "name": "eoln_type",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "if value \u003e 0 the run will be visible",
                        "name": "is_visible",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "if value \u003e 0 any non OK verdict is Check failed",
                        "name": "not_ok_is_cf",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "if value \u003e 0 the run is tested with rejudge priority",
                        "name": "rejudge_flag",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Source code",
                        "name": "file",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Source code",
                        "name": "text_form",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.MasterPostSubmitRunReply"
                        }
                    }
                }
            }
        },
        "/ej/master/submit-run-input": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "submit a program and its input for testing (privileged)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "privileged submit"
                ],
                "summary": "Submit a program and its input for testing (privileged)",
                "operationId": "master-submit-run-input",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Impersonated user login",
                        "name": "sender_user_login",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Impersonated user ID",
                        "name": "sender_user_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Fake sender IP",
                        "name": "sender_ip",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "HTTPS flag",
                        "name": "sender_ssl_flag",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "contest ID",
                        "name": "contest_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "problem ID or short_name",
                        "name": "prob_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "language ID or short_name",
                        "name": "lang_id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "End-of-Line translation type",
                        "name": "eoln_type",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Source code",
                        "name": "file",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "Source code",
                        "name": "text_form",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Input data",
                        "name": "file_input",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Input data",
                        "name": "text_form_input",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.MasterPostSubmitRunInputReply"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.EjudgeRun": {
            "type": "object",
            "properties": {
                "contest_id": {
                    "type": "integer"
                },
                "duration": {
                    "type": "integer"
                },
                "eoln_type": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string"
                },
                "ipv6_flag": {
                    "type": "boolean"
                },
                "is_checked": {
                    "type": "boolean"
                },
                "is_hidden": {
                    "type": "boolean"
                },
                "is_imported": {
                    "type": "boolean"
                },
                "is_marked": {
                    "type": "boolean"
                },
                "is_readonly": {
                    "type": "boolean"
                },
                "is_saved": {
                    "type": "boolean"
                },
                "is_vcs": {
                    "type": "boolean"
                },
                "judge_id": {
                    "type": "integer"
                },
                "judge_uuid": {
                    "type": "string"
                },
                "lang_id": {
                    "type": "integer"
                },
                "lang_name": {
                    "type": "string"
                },
                "last_change_us": {
                    "type": "integer"
                },
                "locale_id": {
                    "type": "integer"
                },
                "mime_type": {
                    "type": "string"
                },
                "nsec": {
                    "type": "integer"
                },
                "pages": {
                    "type": "integer"
                },
                "passed_mode": {
                    "type": "boolean"
                },
                "prob_id": {
                    "type": "integer"
                },
                "prob_internal_name": {
                    "type": "string"
                },
                "prob_name": {
                    "type": "string"
                },
                "prob_uuid": {
                    "type": "string"
                },
                "raw_score": {
                    "type": "integer"
                },
                "raw_variant": {
                    "type": "integer"
                },
                "run_id": {
                    "type": "integer"
                },
                "run_time": {
                    "type": "integer"
                },
                "run_time_us": {
                    "type": "integer"
                },
                "run_uuid": {
                    "type": "string"
                },
                "saved_score": {
                    "type": "integer"
                },
                "saved_status": {
                    "type": "integer"
                },
                "saved_status_str": {
                    "type": "string"
                },
                "saved_test": {
                    "type": "integer"
                },
                "score": {
                    "type": "integer"
                },
                "score_adj": {
                    "type": "integer"
                },
                "score_str": {
                    "type": "string"
                },
                "serial_id": {
                    "type": "integer"
                },
                "sha1": {
                    "type": "string"
                },
                "sha256": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "ssl_flag": {
                    "type": "boolean"
                },
                "status": {
                    "type": "integer"
                },
                "status_str": {
                    "type": "string"
                },
                "store_flags": {
                    "type": "integer"
                },
                "test": {
                    "type": "integer"
                },
                "tests_passed": {
                    "type": "integer"
                },
                "token_count": {
                    "type": "integer"
                },
                "token_flags": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_login": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "variant": {
                    "type": "integer"
                },
                "verdict_bits": {
                    "type": "integer"
                }
            }
        },
        "main.EjudgeSubmit": {
            "type": "object",
            "properties": {
                "compiler_output": {
                    "type": "string"
                },
                "contest_id": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                },
                "exit_code": {
                    "type": "integer"
                },
                "input": {
                    "type": "string"
                },
                "lang_id": {
                    "type": "integer"
                },
                "max_memory_used": {
                    "type": "integer"
                },
                "max_rss": {
                    "type": "integer"
                },
                "output": {
                    "type": "string"
                },
                "prob_id": {
                    "type": "integer"
                },
                "real_time": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "status_str": {
                    "type": "string"
                },
                "submit_id": {
                    "type": "integer"
                },
                "term_signal": {
                    "type": "integer"
                },
                "test_checker_output": {
                    "type": "string"
                },
                "time": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "main.MasterGetContestStatusJSONReply": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "error": {
                    "$ref": "#/definitions/main.ReplyError"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/main.MasterGetContestStatusJSONReplyResult"
                },
                "server_time": {
                    "type": "integer"
                }
            }
        },
        "main.MasterGetContestStatusJSONReplyResult": {
            "type": "object",
            "properties": {
                "contest": {
                    "$ref": "#/definitions/main.MasterGetContestStatusJSONReplyResultContest"
                },
                "online": {
                    "$ref": "#/definitions/main.MasterGetContestStatusJSONReplyResultOnline"
                }
            }
        },
        "main.MasterGetContestStatusJSONReplyResultContest": {
            "type": "object",
            "properties": {
                "board_fog_duration": {
                    "type": "integer"
                },
                "board_unfog_duration": {
                    "type": "integer"
                },
                "close_time": {
                    "type": "integer"
                },
                "contest_load_time": {
                    "type": "integer"
                },
                "duration": {
                    "type": "integer"
                },
                "expected_stop_time": {
                    "type": "integer"
                },
                "freeze_time": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_clients_suspended": {
                    "type": "boolean"
                },
                "is_continue_enabled": {
                    "type": "boolean"
                },
                "is_final_visibility": {
                    "type": "boolean"
                },
                "is_freezable": {
                    "type": "boolean"
                },
                "is_frozen": {
                    "type": "boolean"
                },
                "is_judge_score": {
                    "type": "boolean"
                },
                "is_olympiad_accepting_mode": {
                    "type": "boolean"
                },
                "is_printing_suspended": {
                    "type": "boolean"
                },
                "is_report_closed": {
                    "type": "boolean"
                },
                "is_report_open": {
                    "type": "boolean"
                },
                "is_restartable": {
                    "type": "boolean"
                },
                "is_source_closed": {
                    "type": "boolean"
                },
                "is_source_open": {
                    "type": "boolean"
                },
                "is_standings_updated": {
                    "type": "boolean"
                },
                "is_started": {
                    "type": "boolean"
                },
                "is_stopped": {
                    "type": "boolean"
                },
                "is_testing_finished": {
                    "type": "boolean"
                },
                "is_testing_suspended": {
                    "type": "boolean"
                },
                "is_unlimited": {
                    "type": "boolean"
                },
                "is_upsolving": {
                    "type": "boolean"
                },
                "is_valuer_judge_comments": {
                    "type": "boolean"
                },
                "is_virtual": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "name_en": {
                    "type": "string"
                },
                "open_time": {
                    "type": "integer"
                },
                "scheduled_finish_time": {
                    "type": "integer"
                },
                "scheduled_start_time": {
                    "type": "integer"
                },
                "score_system": {
                    "type": "integer"
                },
                "server_start_time": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "stop_time": {
                    "type": "integer"
                },
                "unfreeze_time": {
                    "type": "integer"
                }
            }
        },
        "main.MasterGetContestStatusJSONReplyResultOnline": {
            "type": "object",
            "properties": {
                "max_time": {
                    "type": "integer"
                },
                "max_user_count": {
                    "type": "integer"
                },
                "user_count": {
                    "type": "integer"
                }
            }
        },
        "main.MasterGetGetSubmitReply": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "error": {
                    "$ref": "#/definitions/main.ReplyError"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/main.EjudgeSubmit"
                },
                "server_time": {
                    "type": "integer"
                }
            }
        },
        "main.MasterGetRunStatusJSONReply": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "error": {
                    "$ref": "#/definitions/main.ReplyError"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/main.MasterGetRunStatusJSONReplyResult"
                },
                "server_time": {
                    "type": "integer"
                }
            }
        },
        "main.MasterGetRunStatusJSONReplyResult": {
            "type": "object",
            "properties": {
                "accepting_mode": {
                    "type": "boolean"
                },
                "run": {
                    "$ref": "#/definitions/main.EjudgeRun"
                }
            }
        },
        "main.MasterPostSubmitRunInputReply": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "error": {
                    "$ref": "#/definitions/main.ReplyError"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/main.MasterPostSubmitRunInputReplyResult"
                },
                "server_time": {
                    "type": "integer"
                }
            }
        },
        "main.MasterPostSubmitRunInputReplyResult": {
            "type": "object",
            "properties": {
                "submit_id": {
                    "type": "integer"
                }
            }
        },
        "main.MasterPostSubmitRunReply": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "error": {
                    "$ref": "#/definitions/main.ReplyError"
                },
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/main.MasterPostSubmitRunReplyResult"
                },
                "server_time": {
                    "type": "integer"
                }
            }
        },
        "main.MasterPostSubmitRunReplyResult": {
            "type": "object",
            "properties": {
                "run_id": {
                    "type": "integer"
                },
                "run_uuid": {
                    "type": "string"
                }
            }
        },
        "main.ReplyError": {
            "type": "object",
            "properties": {
                "log_id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "num": {
                    "type": "integer"
                },
                "symbol": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "API Key Authorization",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.10.3",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Ejudge API",
	Description:      "The ejudge API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
