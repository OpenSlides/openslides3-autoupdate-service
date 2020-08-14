// Code generated with export.json DO NOT EDIT.

package test

import "encoding/json"

var exampleData = map[string]json.RawMessage{
	"agenda/item:1": []byte(`{
          "id": 1,
          "item_number": "",
          "title_information": {
            "title": "Topic"
          },
          "comment": null,
          "closed": false,
          "type": 1,
          "is_internal": false,
          "is_hidden": false,
          "duration": null,
          "content_object": {
            "collection": "topics/topic",
            "id": 1
          },
          "weight": 2,
          "parent_id": null,
          "level": 0,
          "tags_id": []
        }`),
	"agenda/item:2": []byte(`{
          "id": 2,
          "item_number": "",
          "title_information": {
            "title": "Hidden"
          },
          "comment": null,
          "closed": false,
          "type": 3,
          "is_internal": false,
          "is_hidden": true,
          "duration": null,
          "content_object": {
            "collection": "topics/topic",
            "id": 2
          },
          "weight": 4,
          "parent_id": null,
          "level": 0,
          "tags_id": []
        }`),
	"agenda/item:3": []byte(`{
          "id": 3,
          "item_number": "",
          "title_information": {
            "title": "Internal"
          },
          "comment": null,
          "closed": false,
          "type": 2,
          "is_internal": true,
          "is_hidden": false,
          "duration": null,
          "content_object": {
            "collection": "topics/topic",
            "id": 3
          },
          "weight": 8,
          "parent_id": null,
          "level": 0,
          "tags_id": []
        }`),
	"agenda/item:4": []byte(`{
          "id": 4,
          "item_number": "",
          "title_information": {
            "title": "Assignment"
          },
          "comment": null,
          "closed": false,
          "type": 1,
          "is_internal": false,
          "is_hidden": true,
          "duration": null,
          "content_object": {
            "collection": "assignments/assignment",
            "id": 1
          },
          "weight": 6,
          "parent_id": 2,
          "level": 1,
          "tags_id": []
        }`),
	"agenda/item:5": []byte(`{
          "id": 5,
          "item_number": "",
          "title_information": {
            "title": "Leadmotion1",
            "identifier": null
          },
          "comment": null,
          "closed": false,
          "type": 1,
          "is_internal": true,
          "is_hidden": false,
          "duration": null,
          "content_object": {
            "collection": "motions/motion",
            "id": 1
          },
          "weight": 14,
          "parent_id": 3,
          "level": 1,
          "tags_id": []
        }`),
	"agenda/item:6": []byte(`{
          "id": 6,
          "item_number": "",
          "title_information": {
            "title": "Änderungsantrag zu Leadmotion1",
            "identifier": "Ä-1"
          },
          "comment": null,
          "closed": false,
          "type": 1,
          "is_internal": true,
          "is_hidden": false,
          "duration": 0,
          "content_object": {
            "collection": "motions/motion",
            "id": 2
          },
          "weight": 16,
          "parent_id": 3,
          "level": 1,
          "tags_id": []
        }`),
	"agenda/item:7": []byte(`{
          "id": 7,
          "item_number": "",
          "title_information": {
            "title": "Public",
            "identifier": "2"
          },
          "comment": null,
          "closed": false,
          "type": 2,
          "is_internal": true,
          "is_hidden": false,
          "duration": null,
          "content_object": {
            "collection": "motions/motion",
            "id": 3
          },
          "weight": 18,
          "parent_id": 6,
          "level": 2,
          "tags_id": []
        }`),
	"agenda/item:8": []byte(`{
          "id": 8,
          "item_number": "",
          "title_information": {
            "title": "block"
          },
          "comment": null,
          "closed": false,
          "type": 3,
          "is_internal": true,
          "is_hidden": true,
          "duration": null,
          "content_object": {
            "collection": "motions/motion-block",
            "id": 1
          },
          "weight": 10,
          "parent_id": 3,
          "level": 1,
          "tags_id": []
        }`),
	"agenda/item:9": []byte(`{
          "id": 9,
          "item_number": "",
          "title_information": {
            "title": "Another public topic"
          },
          "comment": null,
          "closed": false,
          "type": 1,
          "is_internal": true,
          "is_hidden": true,
          "duration": null,
          "content_object": {
            "collection": "topics/topic",
            "id": 4
          },
          "weight": 12,
          "parent_id": 8,
          "level": 2,
          "tags_id": []
        }`),
	"agenda/list-of-speakers:1": []byte(`{
          "id": 1,
          "title_information": {
            "title": "Topic"
          },
          "speakers": [
            {
              "id": 3,
              "user_id": 6,
              "begin_time": "2020-08-11T12:28:30.894574+02:00",
              "end_time": null,
              "weight": null,
              "marked": false
            }
          ],
          "closed": false,
          "content_object": {
            "collection": "topics/topic",
            "id": 1
          }
        }`),
	"agenda/list-of-speakers:10": []byte(`{
          "id": 10,
          "title_information": {
            "title": "Public",
            "identifier": "2"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "motions/motion",
            "id": 3
          }
        }`),
	"agenda/list-of-speakers:11": []byte(`{
          "id": 11,
          "title_information": {
            "title": "A.txt"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "mediafiles/mediafile",
            "id": 4
          }
        }`),
	"agenda/list-of-speakers:12": []byte(`{
          "id": 12,
          "title_information": {
            "title": "block"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "motions/motion-block",
            "id": 1
          }
        }`),
	"agenda/list-of-speakers:13": []byte(`{
          "id": 13,
          "title_information": {
            "title": "block internal"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "motions/motion-block",
            "id": 2
          }
        }`),
	"agenda/list-of-speakers:14": []byte(`{
          "id": 14,
          "title_information": {
            "title": "Another public topic"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "topics/topic",
            "id": 4
          }
        }`),
	"agenda/list-of-speakers:2": []byte(`{
          "id": 2,
          "title_information": {
            "title": "Hidden"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "topics/topic",
            "id": 2
          }
        }`),
	"agenda/list-of-speakers:3": []byte(`{
          "id": 3,
          "title_information": {
            "title": "Internal"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "topics/topic",
            "id": 3
          }
        }`),
	"agenda/list-of-speakers:4": []byte(`{
          "id": 4,
          "title_information": {
            "title": "folder"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "mediafiles/mediafile",
            "id": 1
          }
        }`),
	"agenda/list-of-speakers:5": []byte(`{
          "id": 5,
          "title_information": {
            "title": "A.txt"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "mediafiles/mediafile",
            "id": 2
          }
        }`),
	"agenda/list-of-speakers:6": []byte(`{
          "id": 6,
          "title_information": {
            "title": "in.jpg"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "mediafiles/mediafile",
            "id": 3
          }
        }`),
	"agenda/list-of-speakers:7": []byte(`{
          "id": 7,
          "title_information": {
            "title": "Assignment"
          },
          "speakers": [
            {
              "id": 4,
              "user_id": 6,
              "begin_time": "2020-08-11T12:29:55.054553+02:00",
              "end_time": null,
              "weight": null,
              "marked": false
            },
            {
              "id": 6,
              "user_id": 7,
              "begin_time": null,
              "end_time": null,
              "weight": 2,
              "marked": false
            }
          ],
          "closed": false,
          "content_object": {
            "collection": "assignments/assignment",
            "id": 1
          }
        }`),
	"agenda/list-of-speakers:8": []byte(`{
          "id": 8,
          "title_information": {
            "title": "Leadmotion1",
            "identifier": null
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "motions/motion",
            "id": 1
          }
        }`),
	"agenda/list-of-speakers:9": []byte(`{
          "id": 9,
          "title_information": {
            "title": "Änderungsantrag zu Leadmotion1",
            "identifier": "Ä-1"
          },
          "speakers": [],
          "closed": false,
          "content_object": {
            "collection": "motions/motion",
            "id": 2
          }
        }`),
	"assignments/assignment-option:1": []byte(`{
          "user_id": 2,
          "weight": 1,
          "id": 1,
          "yes": "1.000000",
          "no": "0.000000",
          "abstain": "0.000000",
          "poll_id": 1,
          "pollstate": 2
        }`),
	"assignments/assignment-option:2": []byte(`{
          "user_id": 3,
          "weight": 3,
          "id": 2,
          "yes": "0.000000",
          "no": "0.000000",
          "abstain": "0.000000",
          "poll_id": 1,
          "pollstate": 2
        }`),
	"assignments/assignment-poll:1": []byte(`{
          "assignment_id": 1,
          "description": "",
          "pollmethod": "votes",
          "votes_amount": 1,
          "allow_multiple_votes_per_candidate": false,
          "global_no": true,
          "amount_global_no": "0.000000",
          "global_abstain": true,
          "amount_global_abstain": "0.000000",
          "state": 2,
          "type": "named",
          "title": "Wahlgang",
          "groups_id": [
            3
          ],
          "votesvalid": "1.000000",
          "votesinvalid": "0.000000",
          "votescast": "1.000000",
          "options_id": [
            1,
            2
          ],
          "id": 1,
          "onehundred_percent_base": "valid",
          "majority_method": "simple",
          "voted_id": [
            4
          ]
        }`),
	"assignments/assignment-vote:1": []byte(`{
          "id": 1,
          "weight": "1.000000",
          "value": "Y",
          "user_id": 4,
          "option_id": 1,
          "pollstate": 2
        }`),
	"assignments/assignment:1": []byte(`{
          "id": 1,
          "title": "Assignment",
          "description": "",
          "open_posts": 1,
          "phase": 0,
          "assignment_related_users": [
            {
              "id": 1,
              "user_id": 2,
              "weight": 1
            },
            {
              "id": 3,
              "user_id": 3,
              "weight": 3
            }
          ],
          "default_poll_description": "",
          "agenda_item_id": 4,
          "list_of_speakers_id": 7,
          "tags_id": [
            2
          ],
          "attachments_id": [],
          "number_poll_candidates": false,
          "polls_id": [
            1
          ]
        }`),
	"core/config:10": []byte(`{
          "id": 10,
          "key": "general_system_conference_show",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Show live conference window",
            "helpText": "Server settings required to activate Jitsi Meet integration.",
            "choices": null,
            "weight": 140,
            "group": "General",
            "subgroup": "Live conference"
          }
        }`),
	"core/config:100": []byte(`{
          "id": 100,
          "key": "motion_poll_default_100_percent_base",
          "value": "YNA",
          "data": {
            "defaultValue": "YNA",
            "inputType": "choice",
            "label": "Default 100 % base of a voting result",
            "helpText": "",
            "choices": [
              {
                "value": "YN",
                "display_name": "Yes/No"
              },
              {
                "value": "YNA",
                "display_name": "Yes/No/Abstain"
              },
              {
                "value": "valid",
                "display_name": "All valid ballots"
              },
              {
                "value": "cast",
                "display_name": "All casted ballots"
              },
              {
                "value": "disabled",
                "display_name": "Disabled (no percents)"
              }
            ],
            "weight": 370,
            "group": "Motions",
            "subgroup": "Voting and ballot papers"
          }
        }`),
	"core/config:101": []byte(`{
          "id": 101,
          "key": "motion_poll_default_majority_method",
          "value": "simple",
          "data": null
        }`),
	"core/config:102": []byte(`{
          "id": 102,
          "key": "motion_poll_default_groups",
          "value": [],
          "data": {
            "defaultValue": [],
            "inputType": "groups",
            "label": "Default groups with voting rights",
            "helpText": "",
            "choices": null,
            "weight": 372,
            "group": "Motions",
            "subgroup": "Voting and ballot papers"
          }
        }`),
	"core/config:103": []byte(`{
          "id": 103,
          "key": "motions_pdf_ballot_papers_selection",
          "value": "CUSTOM_NUMBER",
          "data": {
            "defaultValue": "CUSTOM_NUMBER",
            "inputType": "choice",
            "label": "Number of ballot papers",
            "helpText": "",
            "choices": [
              {
                "value": "NUMBER_OF_DELEGATES",
                "display_name": "Number of all delegates"
              },
              {
                "value": "NUMBER_OF_ALL_PARTICIPANTS",
                "display_name": "Number of all participants"
              },
              {
                "value": "CUSTOM_NUMBER",
                "display_name": "Use the following custom number"
              }
            ],
            "weight": 373,
            "group": "Motions",
            "subgroup": "Voting and ballot papers"
          }
        }`),
	"core/config:104": []byte(`{
          "id": 104,
          "key": "motions_pdf_ballot_papers_number",
          "value": 8,
          "data": {
            "defaultValue": 8,
            "inputType": "integer",
            "label": "Custom number of ballot papers",
            "helpText": "",
            "choices": null,
            "weight": 374,
            "group": "Motions",
            "subgroup": "Voting and ballot papers"
          }
        }`),
	"core/config:105": []byte(`{
          "id": 105,
          "key": "motions_export_title",
          "value": "Motions",
          "data": {
            "defaultValue": "Motions",
            "inputType": "string",
            "label": "Title for PDF documents of motions",
            "helpText": "",
            "choices": null,
            "weight": 380,
            "group": "Motions",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:106": []byte(`{
          "id": 106,
          "key": "motions_export_preamble",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Preamble text for PDF documents of motions",
            "helpText": "",
            "choices": null,
            "weight": 382,
            "group": "Motions",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:107": []byte(`{
          "id": 107,
          "key": "motions_export_submitter_recommendation",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Show submitters and recommendation/state in table of contents",
            "helpText": "",
            "choices": null,
            "weight": 384,
            "group": "Motions",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:108": []byte(`{
          "id": 108,
          "key": "motions_export_follow_recommendation",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Show checkbox to record decision",
            "helpText": "",
            "choices": null,
            "weight": 386,
            "group": "Motions",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:109": []byte(`{
          "id": 109,
          "key": "assignment_poll_method",
          "value": "votes",
          "data": {
            "defaultValue": "votes",
            "inputType": "choice",
            "label": "Default election method",
            "helpText": "",
            "choices": [
              {
                "value": "votes",
                "display_name": "Yes per candidate"
              },
              {
                "value": "YN",
                "display_name": "Yes/No per candidate"
              },
              {
                "value": "YNA",
                "display_name": "Yes/No/Abstain per candidate"
              }
            ],
            "weight": 400,
            "group": "Elections",
            "subgroup": "Ballot"
          }
        }`),
	"core/config:11": []byte(`{
          "id": 11,
          "key": "general_system_conference_auto_connect",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Connect all users to live conference automatically",
            "helpText": "Server settings required to activate Jitsi Meet integration.",
            "choices": null,
            "weight": 141,
            "group": "General",
            "subgroup": "Live conference"
          }
        }`),
	"core/config:110": []byte(`{
          "id": 110,
          "key": "assignment_poll_default_type",
          "value": "analog",
          "data": {
            "defaultValue": "analog",
            "inputType": "choice",
            "label": "Default voting type",
            "helpText": "",
            "choices": [
              {
                "value": "analog",
                "display_name": "analog"
              },
              {
                "value": "named",
                "display_name": "nominal"
              },
              {
                "value": "pseudoanonymous",
                "display_name": "non-nominal"
              }
            ],
            "weight": 403,
            "group": "Elections",
            "subgroup": "Ballot"
          }
        }`),
	"core/config:111": []byte(`{
          "id": 111,
          "key": "assignment_poll_default_100_percent_base",
          "value": "valid",
          "data": {
            "defaultValue": "valid",
            "inputType": "choice",
            "label": "Default 100 % base of an election result",
            "helpText": "",
            "choices": [
              {
                "value": "YN",
                "display_name": "Yes/No per candidate"
              },
              {
                "value": "YNA",
                "display_name": "Yes/No/Abstain per candidate"
              },
              {
                "value": "votes",
                "display_name": "Sum of votes including general No/Abstain"
              },
              {
                "value": "valid",
                "display_name": "All valid ballots"
              },
              {
                "value": "cast",
                "display_name": "All casted ballots"
              },
              {
                "value": "disabled",
                "display_name": "Disabled (no percents)"
              }
            ],
            "weight": 405,
            "group": "Elections",
            "subgroup": "Ballot"
          }
        }`),
	"core/config:112": []byte(`{
          "id": 112,
          "key": "assignment_poll_default_groups",
          "value": [],
          "data": {
            "defaultValue": [],
            "inputType": "groups",
            "label": "Default groups with voting rights",
            "helpText": "",
            "choices": null,
            "weight": 410,
            "group": "Elections",
            "subgroup": "Ballot"
          }
        }`),
	"core/config:113": []byte(`{
          "id": 113,
          "key": "assignment_poll_default_majority_method",
          "value": "simple",
          "data": null
        }`),
	"core/config:114": []byte(`{
          "id": 114,
          "key": "assignment_poll_sort_poll_result_by_votes",
          "value": true,
          "data": {
            "defaultValue": true,
            "inputType": "boolean",
            "label": "Sort election results by amount of votes",
            "helpText": "",
            "choices": null,
            "weight": 420,
            "group": "Elections",
            "subgroup": "Ballot"
          }
        }`),
	"core/config:115": []byte(`{
          "id": 115,
          "key": "assignment_poll_add_candidates_to_list_of_speakers",
          "value": true,
          "data": {
            "defaultValue": true,
            "inputType": "boolean",
            "label": "Put all candidates on the list of speakers",
            "helpText": "",
            "choices": null,
            "weight": 425,
            "group": "Elections",
            "subgroup": "Ballot"
          }
        }`),
	"core/config:116": []byte(`{
          "id": 116,
          "key": "assignments_pdf_ballot_papers_selection",
          "value": "CUSTOM_NUMBER",
          "data": {
            "defaultValue": "CUSTOM_NUMBER",
            "inputType": "choice",
            "label": "Number of ballot papers",
            "helpText": "",
            "choices": [
              {
                "value": "NUMBER_OF_DELEGATES",
                "display_name": "Number of all delegates"
              },
              {
                "value": "NUMBER_OF_ALL_PARTICIPANTS",
                "display_name": "Number of all participants"
              },
              {
                "value": "CUSTOM_NUMBER",
                "display_name": "Use the following custom number"
              }
            ],
            "weight": 430,
            "group": "Elections",
            "subgroup": "Ballot papers"
          }
        }`),
	"core/config:117": []byte(`{
          "id": 117,
          "key": "assignments_pdf_ballot_papers_number",
          "value": 8,
          "data": {
            "defaultValue": 8,
            "inputType": "integer",
            "label": "Custom number of ballot papers",
            "helpText": "",
            "choices": null,
            "weight": 435,
            "group": "Elections",
            "subgroup": "Ballot papers"
          }
        }`),
	"core/config:118": []byte(`{
          "id": 118,
          "key": "assignments_pdf_title",
          "value": "Elections",
          "data": {
            "defaultValue": "Elections",
            "inputType": "string",
            "label": "Title for PDF document (all elections)",
            "helpText": "",
            "choices": null,
            "weight": 460,
            "group": "Elections",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:119": []byte(`{
          "id": 119,
          "key": "assignments_pdf_preamble",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Preamble text for PDF document (all elections)",
            "helpText": "",
            "choices": null,
            "weight": 470,
            "group": "Elections",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:12": []byte(`{
          "id": 12,
          "key": "general_system_conference_los_restriction",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Allow only current speakers and list of speakers managers to enter the live conference",
            "helpText": "Server settings required to activate Jitsi Meet integration.",
            "choices": null,
            "weight": 142,
            "group": "General",
            "subgroup": "Live conference"
          }
        }`),
	"core/config:13": []byte(`{
          "id": 13,
          "key": "general_system_stream_url",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Livestream url",
            "helpText": "Remove URL to deactivate livestream. Check extra group permission to see livestream.",
            "choices": null,
            "weight": 143,
            "group": "General",
            "subgroup": "Live conference"
          }
        }`),
	"core/config:14": []byte(`{
          "id": 14,
          "key": "general_system_enable_anonymous",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Allow access for anonymous guest users",
            "helpText": "",
            "choices": null,
            "weight": 150,
            "group": "General",
            "subgroup": "System"
          }
        }`),
	"core/config:15": []byte(`{
          "id": 15,
          "key": "general_login_info_text",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Show this text on the login page",
            "helpText": "",
            "choices": null,
            "weight": 152,
            "group": "General",
            "subgroup": "System"
          }
        }`),
	"core/config:16": []byte(`{
          "id": 16,
          "key": "openslides_theme",
          "value": "openslides-default-light-theme",
          "data": {
            "defaultValue": "openslides-default-light-theme",
            "inputType": "choice",
            "label": "OpenSlides Theme",
            "helpText": "",
            "choices": [
              {
                "value": "openslides-default-light-theme",
                "display_name": "OpenSlides Default"
              },
              {
                "value": "openslides-default-dark-theme",
                "display_name": "OpenSlides Dark"
              },
              {
                "value": "openslides-red-light-theme",
                "display_name": "OpenSlides Red"
              },
              {
                "value": "openslides-red-dark-theme",
                "display_name": "OpenSlides Red Dark"
              },
              {
                "value": "openslides-green-light-theme",
                "display_name": "OpenSlides Green"
              },
              {
                "value": "openslides-green-dark-theme",
                "display_name": "OpenSlides Green Dark"
              },
              {
                "value": "openslides-solarized-dark-theme",
                "display_name": "OpenSlides Solarized"
              }
            ],
            "weight": 154,
            "group": "General",
            "subgroup": "System"
          }
        }`),
	"core/config:17": []byte(`{
          "id": 17,
          "key": "general_csv_separator",
          "value": ",",
          "data": {
            "defaultValue": ",",
            "inputType": "string",
            "label": "Separator used for all csv exports and examples",
            "helpText": "",
            "choices": null,
            "weight": 160,
            "group": "General",
            "subgroup": "Export"
          }
        }`),
	"core/config:18": []byte(`{
          "id": 18,
          "key": "general_csv_encoding",
          "value": "utf-8",
          "data": {
            "defaultValue": "utf-8",
            "inputType": "choice",
            "label": "Default encoding for all csv exports",
            "helpText": "",
            "choices": [
              {
                "value": "utf-8",
                "display_name": "UTF-8"
              },
              {
                "value": "iso-8859-15",
                "display_name": "ISO-8859-15"
              }
            ],
            "weight": 162,
            "group": "General",
            "subgroup": "Export"
          }
        }`),
	"core/config:19": []byte(`{
          "id": 19,
          "key": "general_export_pdf_pagenumber_alignment",
          "value": "center",
          "data": {
            "defaultValue": "center",
            "inputType": "choice",
            "label": "Page number alignment in PDF",
            "helpText": "",
            "choices": [
              {
                "value": "left",
                "display_name": "Left"
              },
              {
                "value": "center",
                "display_name": "Center"
              },
              {
                "value": "right",
                "display_name": "Right"
              }
            ],
            "weight": 164,
            "group": "General",
            "subgroup": "Export"
          }
        }`),
	"core/config:2": []byte(`{
          "id": 2,
          "key": "general_event_name",
          "value": "OpenSlides",
          "data": {
            "defaultValue": "OpenSlides",
            "inputType": "string",
            "label": "Event name",
            "helpText": "",
            "choices": null,
            "weight": 110,
            "group": "General",
            "subgroup": "Event"
          }
        }`),
	"core/config:20": []byte(`{
          "id": 20,
          "key": "general_export_pdf_fontsize",
          "value": "10",
          "data": {
            "defaultValue": "10",
            "inputType": "choice",
            "label": "Standard font size in PDF",
            "helpText": "",
            "choices": [
              {
                "value": "10",
                "display_name": "10"
              },
              {
                "value": "11",
                "display_name": "11"
              },
              {
                "value": "12",
                "display_name": "12"
              }
            ],
            "weight": 166,
            "group": "General",
            "subgroup": "Export"
          }
        }`),
	"core/config:21": []byte(`{
          "id": 21,
          "key": "general_export_pdf_pagesize",
          "value": "A4",
          "data": {
            "defaultValue": "A4",
            "inputType": "choice",
            "label": "Standard page size in PDF",
            "helpText": "",
            "choices": [
              {
                "value": "A4",
                "display_name": "DIN A4"
              },
              {
                "value": "A5",
                "display_name": "DIN A5"
              }
            ],
            "weight": 168,
            "group": "General",
            "subgroup": "Export"
          }
        }`),
	"core/config:22": []byte(`{
          "id": 22,
          "key": "logos_available",
          "value": [
            "logo_projector_main",
            "logo_projector_header",
            "logo_web_header",
            "logo_pdf_header_L",
            "logo_pdf_header_R",
            "logo_pdf_footer_L",
            "logo_pdf_footer_R",
            "logo_pdf_ballot_paper"
          ],
          "data": null
        }`),
	"core/config:23": []byte(`{
          "id": 23,
          "key": "logo_projector_main",
          "value": {
            "display_name": "Projector logo",
            "path": ""
          },
          "data": null
        }`),
	"core/config:24": []byte(`{
          "id": 24,
          "key": "logo_projector_header",
          "value": {
            "display_name": "Projector header image",
            "path": ""
          },
          "data": null
        }`),
	"core/config:25": []byte(`{
          "id": 25,
          "key": "logo_web_header",
          "value": {
            "display_name": "Web interface header logo",
            "path": "/media/folder/in.jpg"
          },
          "data": null
        }`),
	"core/config:26": []byte(`{
          "id": 26,
          "key": "logo_pdf_header_L",
          "value": {
            "display_name": "PDF header logo (left)",
            "path": ""
          },
          "data": null
        }`),
	"core/config:27": []byte(`{
          "id": 27,
          "key": "logo_pdf_header_R",
          "value": {
            "display_name": "PDF header logo (right)",
            "path": ""
          },
          "data": null
        }`),
	"core/config:28": []byte(`{
          "id": 28,
          "key": "logo_pdf_footer_L",
          "value": {
            "display_name": "PDF footer logo (left)",
            "path": "/media/folder/in.jpg"
          },
          "data": null
        }`),
	"core/config:29": []byte(`{
          "id": 29,
          "key": "logo_pdf_footer_R",
          "value": {
            "display_name": "PDF footer logo (right)",
            "path": ""
          },
          "data": null
        }`),
	"core/config:3": []byte(`{
          "id": 3,
          "key": "general_event_description",
          "value": "Presentation and assembly system",
          "data": {
            "defaultValue": "Presentation and assembly system",
            "inputType": "string",
            "label": "Short description of event",
            "helpText": "",
            "choices": null,
            "weight": 115,
            "group": "General",
            "subgroup": "Event"
          }
        }`),
	"core/config:30": []byte(`{
          "id": 30,
          "key": "logo_pdf_ballot_paper",
          "value": {
            "display_name": "PDF ballot paper logo",
            "path": ""
          },
          "data": null
        }`),
	"core/config:31": []byte(`{
          "id": 31,
          "key": "fonts_available",
          "value": [
            "font_regular",
            "font_italic",
            "font_bold",
            "font_bold_italic",
            "font_monospace"
          ],
          "data": null
        }`),
	"core/config:32": []byte(`{
          "id": 32,
          "key": "font_regular",
          "value": {
            "display_name": "Font regular",
            "default": "assets/fonts/fira-sans-latin-400.woff",
            "path": ""
          },
          "data": null
        }`),
	"core/config:33": []byte(`{
          "id": 33,
          "key": "font_italic",
          "value": {
            "display_name": "Font italic",
            "default": "assets/fonts/fira-sans-latin-400italic.woff",
            "path": ""
          },
          "data": null
        }`),
	"core/config:34": []byte(`{
          "id": 34,
          "key": "font_bold",
          "value": {
            "display_name": "Font bold",
            "default": "assets/fonts/fira-sans-latin-500.woff",
            "path": ""
          },
          "data": null
        }`),
	"core/config:35": []byte(`{
          "id": 35,
          "key": "font_bold_italic",
          "value": {
            "display_name": "Font bold italic",
            "default": "assets/fonts/fira-sans-latin-500italic.woff",
            "path": ""
          },
          "data": null
        }`),
	"core/config:36": []byte(`{
          "id": 36,
          "key": "font_monospace",
          "value": {
            "display_name": "Font monospace",
            "default": "assets/fonts/roboto-condensed-bold.woff",
            "path": ""
          },
          "data": null
        }`),
	"core/config:37": []byte(`{
          "id": 37,
          "key": "translations",
          "value": [],
          "data": {
            "defaultValue": [],
            "inputType": "translations",
            "label": "Custom translations",
            "helpText": "",
            "choices": null,
            "weight": 1000,
            "group": "Custom translations",
            "subgroup": "General"
          }
        }`),
	"core/config:38": []byte(`{
          "id": 38,
          "key": "config_version",
          "value": 2,
          "data": null
        }`),
	"core/config:39": []byte(`{
          "id": 39,
          "key": "db_id",
          "value": "2c3bd736c87a48a4be1f0dc707702144",
          "data": null
        }`),
	"core/config:4": []byte(`{
          "id": 4,
          "key": "general_event_date",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Event date",
            "helpText": "",
            "choices": null,
            "weight": 120,
            "group": "General",
            "subgroup": "Event"
          }
        }`),
	"core/config:40": []byte(`{
          "id": 40,
          "key": "users_sort_by",
          "value": "first_name",
          "data": {
            "defaultValue": "first_name",
            "inputType": "choice",
            "label": "Sort name of participants by",
            "helpText": "",
            "choices": [
              {
                "value": "first_name",
                "display_name": "Given name"
              },
              {
                "value": "last_name",
                "display_name": "Surname"
              },
              {
                "value": "number",
                "display_name": "Participant number"
              }
            ],
            "weight": 510,
            "group": "Participants",
            "subgroup": "General"
          }
        }`),
	"core/config:41": []byte(`{
          "id": 41,
          "key": "users_enable_presence_view",
          "value": true,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Enable participant presence view",
            "helpText": "",
            "choices": null,
            "weight": 511,
            "group": "Participants",
            "subgroup": "General"
          }
        }`),
	"core/config:42": []byte(`{
          "id": 42,
          "key": "users_allow_self_set_present",
          "value": true,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Allow users to set themselves as present",
            "helpText": "e.g. for online meetings",
            "choices": null,
            "weight": 512,
            "group": "Participants",
            "subgroup": "General"
          }
        }`),
	"core/config:43": []byte(`{
          "id": 43,
          "key": "users_activate_vote_weight",
          "value": true,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Activate vote weight",
            "helpText": "",
            "choices": null,
            "weight": 513,
            "group": "Participants",
            "subgroup": "General"
          }
        }`),
	"core/config:44": []byte(`{
          "id": 44,
          "key": "users_pdf_welcometitle",
          "value": "Welcome to OpenSlides",
          "data": {
            "defaultValue": "Welcome to OpenSlides",
            "inputType": "string",
            "label": "Title for access data and welcome PDF",
            "helpText": "",
            "choices": null,
            "weight": 520,
            "group": "Participants",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:45": []byte(`{
          "id": 45,
          "key": "users_pdf_welcometext",
          "value": "[Place for your welcome and help text.]",
          "data": {
            "defaultValue": "[Place for your welcome and help text.]",
            "inputType": "string",
            "label": "Help text for access data and welcome PDF",
            "helpText": "",
            "choices": null,
            "weight": 530,
            "group": "Participants",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:46": []byte(`{
          "id": 46,
          "key": "users_pdf_url",
          "value": "http://example.com:8000",
          "data": {
            "defaultValue": "http://example.com:8000",
            "inputType": "string",
            "label": "System URL",
            "helpText": "Used for QRCode in PDF of access data.",
            "choices": null,
            "weight": 540,
            "group": "Participants",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:47": []byte(`{
          "id": 47,
          "key": "users_pdf_wlan_ssid",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "WLAN name (SSID)",
            "helpText": "Used for WLAN QRCode in PDF of access data.",
            "choices": null,
            "weight": 550,
            "group": "Participants",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:48": []byte(`{
          "id": 48,
          "key": "users_pdf_wlan_password",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "WLAN password",
            "helpText": "Used for WLAN QRCode in PDF of access data.",
            "choices": null,
            "weight": 560,
            "group": "Participants",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:49": []byte(`{
          "id": 49,
          "key": "users_pdf_wlan_encryption",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "choice",
            "label": "WLAN encryption",
            "helpText": "Used for WLAN QRCode in PDF of access data.",
            "choices": [
              {
                "value": "",
                "display_name": "---------"
              },
              {
                "value": "WEP",
                "display_name": "WEP"
              },
              {
                "value": "WPA",
                "display_name": "WPA/WPA2"
              },
              {
                "value": "nopass",
                "display_name": "No encryption"
              }
            ],
            "weight": 570,
            "group": "Participants",
            "subgroup": "PDF export"
          }
        }`),
	"core/config:5": []byte(`{
          "id": 5,
          "key": "general_event_location",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Event location",
            "helpText": "",
            "choices": null,
            "weight": 125,
            "group": "General",
            "subgroup": "Event"
          }
        }`),
	"core/config:50": []byte(`{
          "id": 50,
          "key": "users_email_sender",
          "value": "OpenSlides",
          "data": {
            "defaultValue": "OpenSlides",
            "inputType": "string",
            "label": "Sender name",
            "helpText": "The sender address is defined in the OpenSlides server settings and should modified by administrator only.",
            "choices": null,
            "weight": 600,
            "group": "Participants",
            "subgroup": "Email"
          }
        }`),
	"core/config:51": []byte(`{
          "id": 51,
          "key": "users_email_replyto",
          "value": "",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Reply address",
            "helpText": "",
            "choices": null,
            "weight": 601,
            "group": "Participants",
            "subgroup": "Email"
          }
        }`),
	"core/config:52": []byte(`{
          "id": 52,
          "key": "users_email_subject",
          "value": "OpenSlides access data",
          "data": {
            "defaultValue": "OpenSlides access data",
            "inputType": "string",
            "label": "Email subject",
            "helpText": "You can use {event_name} and {username} as placeholder.",
            "choices": null,
            "weight": 605,
            "group": "Participants",
            "subgroup": "Email"
          }
        }`),
	"core/config:53": []byte(`{
          "id": 53,
          "key": "users_email_body",
          "value": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
          "data": {
            "defaultValue": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
            "inputType": "text",
            "label": "Email body",
            "helpText": "Use these placeholders: {name}, {event_name}, {url}, {username}, {password}. The url referrs to the system url.",
            "choices": null,
            "weight": 610,
            "group": "Participants",
            "subgroup": "Email"
          }
        }`),
	"core/config:54": []byte(`{
          "id": 54,
          "key": "agenda_start_event_date_time",
          "value": null,
          "data": {
            "defaultValue": null,
            "inputType": "datetimepicker",
            "label": "Begin of event",
            "helpText": "Input format: DD.MM.YYYY HH:MM",
            "choices": null,
            "weight": 200,
            "group": "Agenda",
            "subgroup": "General"
          }
        }`),
	"core/config:55": []byte(`{
          "id": 55,
          "key": "agenda_show_subtitle",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Show subtitles in the agenda",
            "helpText": "",
            "choices": null,
            "weight": 201,
            "group": "Agenda",
            "subgroup": "General"
          }
        }`),
	"core/config:56": []byte(`{
          "id": 56,
          "key": "agenda_enable_numbering",
          "value": true,
          "data": {
            "defaultValue": true,
            "inputType": "boolean",
            "label": "Enable numbering for agenda items",
            "helpText": "",
            "choices": null,
            "weight": 205,
            "group": "Agenda",
            "subgroup": "Numbering"
          }
        }`),
	"core/config:57": []byte(`{
          "id": 57,
          "key": "agenda_number_prefix",
          "value": "TOP",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Numbering prefix for agenda items",
            "helpText": "This prefix will be set if you run the automatic agenda numbering.",
            "choices": null,
            "weight": 206,
            "group": "Agenda",
            "subgroup": "Numbering"
          }
        }`),
	"core/config:58": []byte(`{
          "id": 58,
          "key": "agenda_numeral_system",
          "value": "arabic",
          "data": {
            "defaultValue": "arabic",
            "inputType": "choice",
            "label": "Numeral system for agenda items",
            "helpText": "",
            "choices": [
              {
                "value": "arabic",
                "display_name": "Arabic"
              },
              {
                "value": "roman",
                "display_name": "Roman"
              }
            ],
            "weight": 207,
            "group": "Agenda",
            "subgroup": "Numbering"
          }
        }`),
	"core/config:59": []byte(`{
          "id": 59,
          "key": "agenda_item_creation",
          "value": "default_yes",
          "data": {
            "defaultValue": "default_yes",
            "inputType": "choice",
            "label": "Add to agenda",
            "helpText": "",
            "choices": [
              {
                "value": "always",
                "display_name": "Always"
              },
              {
                "value": "never",
                "display_name": "Never"
              },
              {
                "value": "default_yes",
                "display_name": "Ask, default yes"
              },
              {
                "value": "default_no",
                "display_name": "Ask, default no"
              }
            ],
            "weight": 210,
            "group": "Agenda",
            "subgroup": "Visibility"
          }
        }`),
	"core/config:6": []byte(`{
          "id": 6,
          "key": "general_event_legal_notice",
          "value": "<a href=\"http://www.openslides.org\">OpenSlides</a> is a free web based presentation and assembly system for visualizing and controlling agenda, motions and elections of an assembly.",
          "data": null
        }`),
	"core/config:60": []byte(`{
          "id": 60,
          "key": "agenda_new_items_default_visibility",
          "value": "2",
          "data": {
            "defaultValue": "2",
            "inputType": "choice",
            "label": "Default visibility for new agenda items (except topics)",
            "helpText": "",
            "choices": [
              {
                "value": "1",
                "display_name": "Public item"
              },
              {
                "value": "2",
                "display_name": "Internal item"
              },
              {
                "value": "3",
                "display_name": "Hidden item"
              }
            ],
            "weight": 211,
            "group": "Agenda",
            "subgroup": "Visibility"
          }
        }`),
	"core/config:61": []byte(`{
          "id": 61,
          "key": "agenda_hide_internal_items_on_projector",
          "value": true,
          "data": {
            "defaultValue": true,
            "inputType": "boolean",
            "label": "Hide internal items when projecting subitems",
            "helpText": "",
            "choices": null,
            "weight": 212,
            "group": "Agenda",
            "subgroup": "Visibility"
          }
        }`),
	"core/config:62": []byte(`{
          "id": 62,
          "key": "agenda_show_last_speakers",
          "value": 0,
          "data": {
            "defaultValue": 0,
            "inputType": "integer",
            "label": "Number of last speakers to be shown on the projector",
            "helpText": "",
            "choices": null,
            "weight": 220,
            "group": "Agenda",
            "subgroup": "List of speakers"
          }
        }`),
	"core/config:63": []byte(`{
          "id": 63,
          "key": "agenda_show_next_speakers",
          "value": -1,
          "data": {
            "defaultValue": -1,
            "inputType": "integer",
            "label": "Number of the next speakers to be shown on the projector",
            "helpText": "Enter number of the next shown speakers. Choose -1 to show all next speakers.",
            "choices": null,
            "weight": 222,
            "group": "Agenda",
            "subgroup": "List of speakers"
          }
        }`),
	"core/config:64": []byte(`{
          "id": 64,
          "key": "agenda_countdown_warning_time",
          "value": 0,
          "data": {
            "defaultValue": 0,
            "inputType": "integer",
            "label": "Show orange countdown in the last x seconds of speaking time",
            "helpText": "Enter duration in seconds. Choose 0 to disable warning color.",
            "choices": null,
            "weight": 224,
            "group": "Agenda",
            "subgroup": "List of speakers"
          }
        }`),
	"core/config:65": []byte(`{
          "id": 65,
          "key": "projector_default_countdown",
          "value": 60,
          "data": {
            "defaultValue": 60,
            "inputType": "integer",
            "label": "Predefined seconds of new countdowns",
            "helpText": "",
            "choices": null,
            "weight": 226,
            "group": "Agenda",
            "subgroup": "List of speakers"
          }
        }`),
	"core/config:66": []byte(`{
          "id": 66,
          "key": "agenda_couple_countdown_and_speakers",
          "value": true,
          "data": {
            "defaultValue": true,
            "inputType": "boolean",
            "label": "Couple countdown with the list of speakers",
            "helpText": "[Begin speech] starts the countdown, [End speech] stops the countdown.",
            "choices": null,
            "weight": 228,
            "group": "Agenda",
            "subgroup": "List of speakers"
          }
        }`),
	"core/config:67": []byte(`{
          "id": 67,
          "key": "agenda_hide_amount_of_speakers",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Hide the amount of speakers in subtitle of list of speakers slide",
            "helpText": "",
            "choices": null,
            "weight": 230,
            "group": "Agenda",
            "subgroup": "List of speakers"
          }
        }`),
	"core/config:68": []byte(`{
          "id": 68,
          "key": "agenda_present_speakers_only",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Only present participants can be added to the list of speakers",
            "helpText": "",
            "choices": null,
            "weight": 232,
            "group": "Agenda",
            "subgroup": "List of speakers"
          }
        }`),
	"core/config:69": []byte(`{
          "id": 69,
          "key": "agenda_show_first_contribution",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Show hint »first speech« in the list of speakers management view",
            "helpText": "",
            "choices": null,
            "weight": 234,
            "group": "Agenda",
            "subgroup": "List of speakers"
          }
        }`),
	"core/config:7": []byte(`{
          "id": 7,
          "key": "general_event_privacy_policy",
          "value": "",
          "data": null
        }`),
	"core/config:70": []byte(`{
          "id": 70,
          "key": "motions_workflow",
          "value": "1",
          "data": {
            "defaultValue": "1",
            "inputType": "choice",
            "label": "Workflow of new motions",
            "helpText": "",
            "choices": [
              {
                "value": "1",
                "display_name": "Simple Workflow"
              },
              {
                "value": "2",
                "display_name": "Complex Workflow"
              }
            ],
            "weight": 310,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:71": []byte(`{
          "id": 71,
          "key": "motions_statute_amendments_workflow",
          "value": "1",
          "data": {
            "defaultValue": "1",
            "inputType": "choice",
            "label": "Workflow of new statute amendments",
            "helpText": "",
            "choices": [
              {
                "value": "1",
                "display_name": "Simple Workflow"
              },
              {
                "value": "2",
                "display_name": "Complex Workflow"
              }
            ],
            "weight": 312,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:72": []byte(`{
          "id": 72,
          "key": "motions_amendments_workflow",
          "value": "1",
          "data": {
            "defaultValue": "1",
            "inputType": "choice",
            "label": "Workflow of new amendments",
            "helpText": "",
            "choices": [
              {
                "value": "1",
                "display_name": "Simple Workflow"
              },
              {
                "value": "2",
                "display_name": "Complex Workflow"
              }
            ],
            "weight": 314,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:73": []byte(`{
          "id": 73,
          "key": "motions_preamble",
          "value": "The assembly may decide:",
          "data": {
            "defaultValue": "The assembly may decide:",
            "inputType": "string",
            "label": "Motion preamble",
            "helpText": "",
            "choices": null,
            "weight": 320,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:74": []byte(`{
          "id": 74,
          "key": "motions_default_line_numbering",
          "value": "outside",
          "data": {
            "defaultValue": "outside",
            "inputType": "choice",
            "label": "Default line numbering",
            "helpText": "",
            "choices": [
              {
                "value": "outside",
                "display_name": "outside"
              },
              {
                "value": "inline",
                "display_name": "inline"
              },
              {
                "value": "none",
                "display_name": "Disabled"
              }
            ],
            "weight": 322,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:75": []byte(`{
          "id": 75,
          "key": "motions_line_length",
          "value": 85,
          "data": {
            "defaultValue": 85,
            "inputType": "integer",
            "label": "Line length",
            "helpText": "The maximum number of characters per line. Relevant when line numbering is enabled. Min: 40",
            "choices": null,
            "weight": 323,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:76": []byte(`{
          "id": 76,
          "key": "motions_reason_required",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Reason required for creating new motion",
            "helpText": "",
            "choices": null,
            "weight": 324,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:77": []byte(`{
          "id": 77,
          "key": "motions_disable_text_on_projector",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Hide motion text on projector",
            "helpText": "",
            "choices": null,
            "weight": 325,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:78": []byte(`{
          "id": 78,
          "key": "motions_disable_reason_on_projector",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Hide reason on projector",
            "helpText": "",
            "choices": null,
            "weight": 326,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:79": []byte(`{
          "id": 79,
          "key": "motions_disable_recommendation_on_projector",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Hide recommendation on projector",
            "helpText": "",
            "choices": null,
            "weight": 327,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:8": []byte(`{
          "id": 8,
          "key": "general_event_welcome_title",
          "value": "Welcome to OpenSlides",
          "data": null
        }`),
	"core/config:80": []byte(`{
          "id": 80,
          "key": "motions_hide_referring_motions",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Hide referring motions",
            "helpText": "",
            "choices": null,
            "weight": 328,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:81": []byte(`{
          "id": 81,
          "key": "motions_disable_sidebox_on_projector",
          "value": true,
          "data": {
            "defaultValue": true,
            "inputType": "boolean",
            "label": "Show meta information box below the title on projector",
            "helpText": "",
            "choices": null,
            "weight": 329,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:82": []byte(`{
          "id": 82,
          "key": "motions_show_sequential_numbers",
          "value": true,
          "data": {
            "defaultValue": true,
            "inputType": "boolean",
            "label": "Show the sequential number for a motion",
            "helpText": "In motion list, motion detail and PDF.",
            "choices": null,
            "weight": 330,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:83": []byte(`{
          "id": 83,
          "key": "motions_recommendations_by",
          "value": "ABK",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Name of recommender",
            "helpText": "Will be displayed as label before selected recommendation. Use an empty value to disable the recommendation system.",
            "choices": null,
            "weight": 332,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:84": []byte(`{
          "id": 84,
          "key": "motions_statute_recommendations_by",
          "value": "ABK2",
          "data": {
            "defaultValue": "",
            "inputType": "string",
            "label": "Name of recommender for statute amendments",
            "helpText": "Will be displayed as label before selected recommendation in statute amendments.",
            "choices": null,
            "weight": 333,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:85": []byte(`{
          "id": 85,
          "key": "motions_recommendation_text_mode",
          "value": "diff",
          "data": {
            "defaultValue": "diff",
            "inputType": "choice",
            "label": "Default text version for change recommendations",
            "helpText": "",
            "choices": [
              {
                "value": "original",
                "display_name": "Original version"
              },
              {
                "value": "changed",
                "display_name": "Changed version"
              },
              {
                "value": "diff",
                "display_name": "Diff version"
              },
              {
                "value": "agreed",
                "display_name": "Final version"
              }
            ],
            "weight": 334,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:86": []byte(`{
          "id": 86,
          "key": "motions_motions_sorting",
          "value": "identifier",
          "data": {
            "defaultValue": "identifier",
            "inputType": "choice",
            "label": "Sort motions by",
            "helpText": "",
            "choices": [
              {
                "value": "weight",
                "display_name": "Call list"
              },
              {
                "value": "identifier",
                "display_name": "Identifier"
              }
            ],
            "weight": 335,
            "group": "Motions",
            "subgroup": "General"
          }
        }`),
	"core/config:87": []byte(`{
          "id": 87,
          "key": "motions_identifier",
          "value": "per_category",
          "data": {
            "defaultValue": "per_category",
            "inputType": "choice",
            "label": "Identifier",
            "helpText": "",
            "choices": [
              {
                "value": "per_category",
                "display_name": "Numbered per category"
              },
              {
                "value": "serially_numbered",
                "display_name": "Serially numbered"
              },
              {
                "value": "manually",
                "display_name": "Set it manually"
              }
            ],
            "weight": 340,
            "group": "Motions",
            "subgroup": "Numbering"
          }
        }`),
	"core/config:88": []byte(`{
          "id": 88,
          "key": "motions_identifier_min_digits",
          "value": 1,
          "data": {
            "defaultValue": 1,
            "inputType": "integer",
            "label": "Number of minimal digits for identifier",
            "helpText": "Uses leading zeros to sort motions correctly by identifier.",
            "choices": null,
            "weight": 342,
            "group": "Motions",
            "subgroup": "Numbering"
          }
        }`),
	"core/config:89": []byte(`{
          "id": 89,
          "key": "motions_identifier_with_blank",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Allow blank in identifier",
            "helpText": "Blank between prefix and number, e.g. 'A 001'.",
            "choices": null,
            "weight": 344,
            "group": "Motions",
            "subgroup": "Numbering"
          }
        }`),
	"core/config:9": []byte(`{
          "id": 9,
          "key": "general_event_welcome_text",
          "value": "[Space for your welcome text.]",
          "data": null
        }`),
	"core/config:90": []byte(`{
          "id": 90,
          "key": "motions_statutes_enabled",
          "value": true,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Activate statute amendments",
            "helpText": "",
            "choices": null,
            "weight": 350,
            "group": "Motions",
            "subgroup": "Amendments"
          }
        }`),
	"core/config:91": []byte(`{
          "id": 91,
          "key": "motions_amendments_enabled",
          "value": true,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Activate amendments",
            "helpText": "",
            "choices": null,
            "weight": 351,
            "group": "Motions",
            "subgroup": "Amendments"
          }
        }`),
	"core/config:92": []byte(`{
          "id": 92,
          "key": "motions_amendments_main_table",
          "value": true,
          "data": {
            "defaultValue": true,
            "inputType": "boolean",
            "label": "Show amendments together with motions",
            "helpText": "",
            "choices": null,
            "weight": 352,
            "group": "Motions",
            "subgroup": "Amendments"
          }
        }`),
	"core/config:93": []byte(`{
          "id": 93,
          "key": "motions_amendments_prefix",
          "value": "Ä-",
          "data": {
            "defaultValue": "-",
            "inputType": "string",
            "label": "Prefix for the identifier for amendments",
            "helpText": "",
            "choices": null,
            "weight": 353,
            "group": "Motions",
            "subgroup": "Amendments"
          }
        }`),
	"core/config:94": []byte(`{
          "id": 94,
          "key": "motions_amendments_text_mode",
          "value": "paragraph",
          "data": {
            "defaultValue": "paragraph",
            "inputType": "choice",
            "label": "How to create new amendments",
            "helpText": "",
            "choices": [
              {
                "value": "freestyle",
                "display_name": "Empty text field"
              },
              {
                "value": "fulltext",
                "display_name": "Edit the whole motion text"
              },
              {
                "value": "paragraph",
                "display_name": "Paragraph-based, Diff-enabled"
              }
            ],
            "weight": 354,
            "group": "Motions",
            "subgroup": "Amendments"
          }
        }`),
	"core/config:95": []byte(`{
          "id": 95,
          "key": "motions_amendments_multiple_paragraphs",
          "value": true,
          "data": {
            "defaultValue": true,
            "inputType": "boolean",
            "label": "Amendments can change multiple paragraphs",
            "helpText": "",
            "choices": null,
            "weight": 355,
            "group": "Motions",
            "subgroup": "Amendments"
          }
        }`),
	"core/config:96": []byte(`{
          "id": 96,
          "key": "motions_amendments_of_amendments",
          "value": true,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Allow amendments of amendments",
            "helpText": "",
            "choices": null,
            "weight": 356,
            "group": "Motions",
            "subgroup": "Amendments"
          }
        }`),
	"core/config:97": []byte(`{
          "id": 97,
          "key": "motions_min_supporters",
          "value": 1,
          "data": {
            "defaultValue": 0,
            "inputType": "integer",
            "label": "Number of (minimum) required supporters for a motion",
            "helpText": "Choose 0 to disable the supporting system.",
            "choices": null,
            "weight": 360,
            "group": "Motions",
            "subgroup": "Supporters"
          }
        }`),
	"core/config:98": []byte(`{
          "id": 98,
          "key": "motions_remove_supporters",
          "value": false,
          "data": {
            "defaultValue": false,
            "inputType": "boolean",
            "label": "Remove all supporters of a motion if a submitter edits his motion in early state",
            "helpText": "",
            "choices": null,
            "weight": 361,
            "group": "Motions",
            "subgroup": "Supporters"
          }
        }`),
	"core/config:99": []byte(`{
          "id": 99,
          "key": "motion_poll_default_type",
          "value": "named",
          "data": {
            "defaultValue": "analog",
            "inputType": "choice",
            "label": "Default voting type",
            "helpText": "",
            "choices": [
              {
                "value": "analog",
                "display_name": "analog"
              },
              {
                "value": "named",
                "display_name": "nominal"
              },
              {
                "value": "pseudoanonymous",
                "display_name": "non-nominal"
              }
            ],
            "weight": 367,
            "group": "Motions",
            "subgroup": "Voting and ballot papers"
          }
        }`),
	"core/countdown:1": []byte(`{
          "id": 1,
          "title": "Default countdown",
          "description": "",
          "default_time": 60,
          "countdown_time": 1597141855.090748,
          "running": true
        }`),
	"core/projection-default:1": []byte(`{
          "id": 1,
          "name": "agenda_all_items",
          "display_name": "Agenda",
          "projector_id": 1
        }`),
	"core/projection-default:10": []byte(`{
          "id": 10,
          "name": "messages",
          "display_name": "Messages",
          "projector_id": 1
        }`),
	"core/projection-default:11": []byte(`{
          "id": 11,
          "name": "countdowns",
          "display_name": "Countdowns",
          "projector_id": 1
        }`),
	"core/projection-default:12": []byte(`{
          "id": 12,
          "name": "assignment_poll",
          "display_name": "Ballots",
          "projector_id": 1
        }`),
	"core/projection-default:13": []byte(`{
          "id": 13,
          "name": "motion_poll",
          "display_name": "Motion votes",
          "projector_id": 1
        }`),
	"core/projection-default:2": []byte(`{
          "id": 2,
          "name": "topics",
          "display_name": "Topics",
          "projector_id": 1
        }`),
	"core/projection-default:3": []byte(`{
          "id": 3,
          "name": "agenda_list_of_speakers",
          "display_name": "List of speakers",
          "projector_id": 1
        }`),
	"core/projection-default:4": []byte(`{
          "id": 4,
          "name": "agenda_current_list_of_speakers",
          "display_name": "Current list of speakers",
          "projector_id": 1
        }`),
	"core/projection-default:5": []byte(`{
          "id": 5,
          "name": "motions",
          "display_name": "Motions",
          "projector_id": 1
        }`),
	"core/projection-default:6": []byte(`{
          "id": 6,
          "name": "motionBlocks",
          "display_name": "Motion blocks",
          "projector_id": 1
        }`),
	"core/projection-default:7": []byte(`{
          "id": 7,
          "name": "assignments",
          "display_name": "Elections",
          "projector_id": 1
        }`),
	"core/projection-default:8": []byte(`{
          "id": 8,
          "name": "users",
          "display_name": "Participants",
          "projector_id": 1
        }`),
	"core/projection-default:9": []byte(`{
          "id": 9,
          "name": "mediafiles",
          "display_name": "Files",
          "projector_id": 1
        }`),
	"core/projector-message:1": []byte(`{
          "id": 1,
          "message": "<p>test</p>"
        }`),
	"core/projector:1": []byte(`{
          "id": 1,
          "elements": [
            {
              "name": "mediafiles/mediafile",
              "id": 3
            }
          ],
          "elements_preview": [],
          "elements_history": [
            [
              {
                "name": "assignments/assignment",
                "id": 1
              }
            ]
          ],
          "scale": 0,
          "scroll": 0,
          "name": "Default projector",
          "width": 1200,
          "aspect_ratio_numerator": 16,
          "aspect_ratio_denominator": 9,
          "reference_projector_id": 1,
          "projectiondefaults_id": [
            1,
            2,
            3,
            4,
            5,
            6,
            7,
            8,
            9,
            10,
            11,
            12,
            13
          ],
          "color": "#000000",
          "background_color": "#ffffff",
          "header_background_color": "#317796",
          "header_font_color": "#f5f5f5",
          "header_h1_color": "#317796",
          "chyron_background_color": "#317796",
          "chyron_font_color": "#ffffff",
          "show_header_footer": true,
          "show_title": true,
          "show_logo": true
        }`),
	"core/projector:2": []byte(`{
          "id": 2,
          "elements": [
            {
              "name": "core/clock",
              "stable": true
            },
            {
              "name": "agenda/current-list-of-speakers",
              "stable": false
            },
            {
              "name": "agenda/current-speaker-chyron",
              "stable": true
            },
            {
              "name": "agenda/current-list-of-speakers-overlay",
              "stable": true
            }
          ],
          "elements_preview": [],
          "elements_history": [],
          "scale": 0,
          "scroll": 0,
          "name": "sideprojector",
          "width": 1200,
          "aspect_ratio_numerator": 16,
          "aspect_ratio_denominator": 9,
          "reference_projector_id": 1,
          "projectiondefaults_id": [],
          "color": "#000000",
          "background_color": "#ffffff",
          "header_background_color": "#317796",
          "header_font_color": "#f5f5f5",
          "header_h1_color": "#317796",
          "chyron_background_color": "#317796",
          "chyron_font_color": "#ffffff",
          "show_header_footer": true,
          "show_title": true,
          "show_logo": true
        }`),
	"core/tag:1": []byte(`{
          "id": 1,
          "name": "T1"
        }`),
	"core/tag:2": []byte(`{
          "id": 2,
          "name": "T2"
        }`),
	"mediafiles/mediafile:1": []byte(`{
          "id": 1,
          "title": "folder",
          "media_url_prefix": "/media/",
          "filesize": "",
          "mimetype": "",
          "pdf_information": {},
          "access_groups_id": [
            3
          ],
          "create_timestamp": "2020-08-11T11:15:50.719490+02:00",
          "is_directory": true,
          "path": "folder/",
          "parent_id": null,
          "list_of_speakers_id": 4,
          "inherited_access_groups_id": [
            3
          ]
        }`),
	"mediafiles/mediafile:2": []byte(`{
          "id": 2,
          "title": "A.txt",
          "media_url_prefix": "/media/",
          "filesize": "< 1 kB",
          "mimetype": "text/plain",
          "pdf_information": {},
          "access_groups_id": [],
          "create_timestamp": "2020-08-11T11:16:07.013124+02:00",
          "is_directory": false,
          "path": "A.txt",
          "parent_id": null,
          "list_of_speakers_id": 5,
          "inherited_access_groups_id": true
        }`),
	"mediafiles/mediafile:3": []byte(`{
          "id": 3,
          "title": "in.jpg",
          "media_url_prefix": "/media/",
          "filesize": "122 kB",
          "mimetype": "image/jpeg",
          "pdf_information": {},
          "access_groups_id": [
            4
          ],
          "create_timestamp": "2020-08-11T11:16:29.302184+02:00",
          "is_directory": false,
          "path": "folder/in.jpg",
          "parent_id": 1,
          "list_of_speakers_id": 6,
          "inherited_access_groups_id": false
        }`),
	"mediafiles/mediafile:4": []byte(`{
          "id": 4,
          "title": "A.txt",
          "media_url_prefix": "/media/",
          "filesize": "< 1 kB",
          "mimetype": "text/plain",
          "pdf_information": {},
          "access_groups_id": [
            4
          ],
          "create_timestamp": "2020-08-11T12:27:14.168247+02:00",
          "is_directory": false,
          "path": "folder/A.txt",
          "parent_id": 1,
          "list_of_speakers_id": 11,
          "inherited_access_groups_id": false
        }`),
	"motions/category:1": []byte(`{
          "id": 1,
          "name": "first",
          "prefix": "A",
          "parent_id": null,
          "weight": 2,
          "level": 0
        }`),
	"motions/category:2": []byte(`{
          "id": 2,
          "name": "second",
          "prefix": "B",
          "parent_id": 1,
          "weight": 4,
          "level": 1
        }`),
	"motions/category:3": []byte(`{
          "id": 3,
          "name": "third",
          "prefix": "C",
          "parent_id": null,
          "weight": 6,
          "level": 0
        }`),
	"motions/motion-block:1": []byte(`{
          "id": 1,
          "title": "block",
          "agenda_item_id": 8,
          "list_of_speakers_id": 12,
          "internal": false,
          "motions_id": [
            1,
            2
          ]
        }`),
	"motions/motion-block:2": []byte(`{
          "id": 2,
          "title": "block internal",
          "agenda_item_id": null,
          "list_of_speakers_id": 13,
          "internal": true,
          "motions_id": [
            3
          ]
        }`),
	"motions/motion-change-recommendation:1": []byte(`{
          "id": 1,
          "motion_id": 1,
          "rejected": false,
          "internal": true,
          "type": 2,
          "other_description": "",
          "line_from": 9,
          "line_to": 11,
          "text": "<p class=\"os-split-before os-split-after\">eleifend ac, enim. Quisque rutrum. Aenean imperdiet.</p>",
          "creation_time": "2020-08-11T11:39:08.232146+02:00"
        }`),
	"motions/motion-comment-section:1": []byte(`{
          "id": 1,
          "name": "public",
          "read_groups_id": [
            2,
            3
          ],
          "write_groups_id": [
            2,
            3
          ],
          "weight": 10000
        }`),
	"motions/motion-comment-section:2": []byte(`{
          "id": 2,
          "name": "internal",
          "read_groups_id": [],
          "write_groups_id": [],
          "weight": 10000
        }`),
	"motions/motion-option:1": []byte(`{
          "id": 1,
          "yes": "0.000000",
          "no": "1.000000",
          "abstain": "0.000000",
          "poll_id": 1,
          "pollstate": 4
        }`),
	"motions/motion-option:2": []byte(`{
          "id": 2,
          "yes": "1.000000",
          "no": "0.000000",
          "abstain": "0.000000",
          "poll_id": 2,
          "pollstate": 4
        }`),
	"motions/motion-poll:1": []byte(`{
          "motion_id": 3,
          "pollmethod": "YNA",
          "state": 4,
          "type": "named",
          "title": "Abstimmung",
          "groups_id": [
            3
          ],
          "votesvalid": "1.000000",
          "votesinvalid": "0.000000",
          "votescast": "1.000000",
          "options_id": [
            1
          ],
          "id": 1,
          "onehundred_percent_base": "YNA",
          "majority_method": "simple",
          "voted_id": [
            5
          ]
        }`),
	"motions/motion-poll:2": []byte(`{
          "motion_id": 3,
          "pollmethod": "YNA",
          "state": 4,
          "type": "pseudoanonymous",
          "title": "Abstimmung (2)",
          "groups_id": [
            3
          ],
          "votesvalid": "1.000000",
          "votesinvalid": "0.000000",
          "votescast": "1.000000",
          "options_id": [
            2
          ],
          "id": 2,
          "onehundred_percent_base": "YNA",
          "majority_method": "simple",
          "voted_id": [
            5
          ]
        }`),
	"motions/motion-vote:1": []byte(`{
          "id": 1,
          "weight": "1.000000",
          "value": "N",
          "user_id": 5,
          "option_id": 1,
          "pollstate": 4
        }`),
	"motions/motion-vote:2": []byte(`{
          "id": 2,
          "weight": "1.000000",
          "value": "Y",
          "user_id": null,
          "option_id": 2,
          "pollstate": 4
        }`),
	"motions/motion:1": []byte(`{
          "id": 1,
          "identifier": null,
          "title": "Leadmotion1",
          "text": "<p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>",
          "amendment_paragraphs": null,
          "modified_final_version": "",
          "reason": "<p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>",
          "parent_id": null,
          "category_id": 3,
          "category_weight": 10000,
          "comments": [],
          "motion_block_id": 1,
          "origin": "",
          "submitters": [
            {
              "id": 2,
              "user_id": 4,
              "motion_id": 1,
              "weight": 1
            }
          ],
          "supporters_id": [],
          "state_id": 5,
          "state_extension": "Maybe",
          "state_restriction": [
            "is_submitter"
          ],
          "statute_paragraph_id": null,
          "workflow_id": 2,
          "recommendation_id": 7,
          "recommendation_extension": "if [motion:2] is acepted",
          "tags_id": [],
          "attachments_id": [
            2
          ],
          "agenda_item_id": 5,
          "list_of_speakers_id": 8,
          "sort_parent_id": null,
          "weight": 10000,
          "created": "2020-08-11T11:36:06.865840+02:00",
          "last_modified": "2020-08-11T12:42:00.402425+02:00",
          "change_recommendations_id": [
            1
          ],
          "amendments_id": [
            2
          ]
        }`),
	"motions/motion:2": []byte(`{
          "id": 2,
          "identifier": "Ä-1",
          "title": "Änderungsantrag zu Leadmotion1",
          "text": "",
          "amendment_paragraphs": [
            "<p>Lorem ipsum dolor sit amet, consectedfgddfgdf&nbsp; gdfgdfg dfg dfg ww ontes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>"
          ],
          "modified_final_version": "",
          "reason": "",
          "parent_id": 1,
          "category_id": 2,
          "category_weight": 10000,
          "comments": [],
          "motion_block_id": 1,
          "origin": "",
          "submitters": [
            {
              "id": 3,
              "user_id": 1,
              "motion_id": 2,
              "weight": 1
            }
          ],
          "supporters_id": [],
          "state_id": 1,
          "state_extension": null,
          "state_restriction": [],
          "statute_paragraph_id": null,
          "workflow_id": 1,
          "recommendation_id": null,
          "recommendation_extension": null,
          "tags_id": [],
          "attachments_id": [],
          "agenda_item_id": 6,
          "list_of_speakers_id": 9,
          "sort_parent_id": null,
          "weight": 10000,
          "created": "2020-08-11T11:39:35.025914+02:00",
          "last_modified": "2020-08-11T12:41:55.666495+02:00",
          "change_recommendations_id": [],
          "amendments_id": []
        }`),
	"motions/motion:3": []byte(`{
          "id": 3,
          "identifier": "2",
          "title": "Public",
          "text": "<p>a</p>",
          "amendment_paragraphs": null,
          "modified_final_version": "",
          "reason": "<p>a</p>",
          "parent_id": null,
          "category_id": 1,
          "category_weight": 10000,
          "comments": [
            {
              "id": 1,
              "comment": "<p>test</p>",
              "section_id": 1,
              "read_groups_id": [
                2,
                3
              ]
            },
            {
              "id": 2,
              "comment": "<p>test</p>",
              "section_id": 2,
              "read_groups_id": []
            }
          ],
          "motion_block_id": 2,
          "origin": "",
          "submitters": [
            {
              "id": 4,
              "user_id": 1,
              "motion_id": 3,
              "weight": 1
            }
          ],
          "supporters_id": [],
          "state_id": 1,
          "state_extension": null,
          "state_restriction": [],
          "statute_paragraph_id": null,
          "workflow_id": 1,
          "recommendation_id": null,
          "recommendation_extension": null,
          "tags_id": [
            1
          ],
          "attachments_id": [],
          "agenda_item_id": 7,
          "list_of_speakers_id": 10,
          "sort_parent_id": null,
          "weight": 10000,
          "created": "2020-08-11T12:24:45.289233+02:00",
          "last_modified": "2020-08-11T12:41:51.319986+02:00",
          "change_recommendations_id": [],
          "amendments_id": []
        }`),
	"motions/state:1": []byte(`{
          "id": 1,
          "name": "submitted",
          "recommendation_label": null,
          "css_class": "lightblue",
          "restriction": [],
          "allow_support": true,
          "allow_create_poll": true,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": 0,
          "show_recommendation_extension_field": false,
          "next_states_id": [
            2,
            3,
            4
          ],
          "workflow_id": 1
        }`),
	"motions/state:10": []byte(`{
          "id": 10,
          "name": "withdrawed",
          "recommendation_label": null,
          "css_class": "grey",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": -1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 2
        }`),
	"motions/state:11": []byte(`{
          "id": 11,
          "name": "adjourned",
          "recommendation_label": "Adjournment",
          "css_class": "grey",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": -1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 2
        }`),
	"motions/state:12": []byte(`{
          "id": 12,
          "name": "not concerned",
          "recommendation_label": "No concernment",
          "css_class": "grey",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": -1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 2
        }`),
	"motions/state:13": []byte(`{
          "id": 13,
          "name": "refered to committee",
          "recommendation_label": "Referral to committee",
          "css_class": "grey",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": -1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 2
        }`),
	"motions/state:14": []byte(`{
          "id": 14,
          "name": "needs review",
          "recommendation_label": null,
          "css_class": "grey",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": -1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 2
        }`),
	"motions/state:15": []byte(`{
          "id": 15,
          "name": "rejected (not authorized)",
          "recommendation_label": "Rejection (not authorized)",
          "css_class": "grey",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": -1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 2
        }`),
	"motions/state:2": []byte(`{
          "id": 2,
          "name": "accepted",
          "recommendation_label": "Acceptance",
          "css_class": "green",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": 1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 1
        }`),
	"motions/state:3": []byte(`{
          "id": 3,
          "name": "rejected",
          "recommendation_label": "Rejection",
          "css_class": "red",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": -1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 1
        }`),
	"motions/state:4": []byte(`{
          "id": 4,
          "name": "not decided",
          "recommendation_label": "No decision",
          "css_class": "grey",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": -1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 1
        }`),
	"motions/state:5": []byte(`{
          "id": 5,
          "name": "in progress",
          "recommendation_label": null,
          "css_class": "lightblue",
          "restriction": [
            "is_submitter"
          ],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": true,
          "dont_set_identifier": true,
          "show_state_extension_field": true,
          "merge_amendment_into_final": 1,
          "show_recommendation_extension_field": false,
          "next_states_id": [
            6,
            10
          ],
          "workflow_id": 2
        }`),
	"motions/state:6": []byte(`{
          "id": 6,
          "name": "submitted",
          "recommendation_label": null,
          "css_class": "lightblue",
          "restriction": [],
          "allow_support": true,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": true,
          "show_state_extension_field": false,
          "merge_amendment_into_final": 0,
          "show_recommendation_extension_field": false,
          "next_states_id": [
            7,
            10,
            15
          ],
          "workflow_id": 2
        }`),
	"motions/state:7": []byte(`{
          "id": 7,
          "name": "permitted",
          "recommendation_label": "Permission",
          "css_class": "lightblue",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": true,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": 0,
          "show_recommendation_extension_field": true,
          "next_states_id": [
            8,
            9,
            10,
            11,
            12,
            13,
            14
          ],
          "workflow_id": 2
        }`),
	"motions/state:8": []byte(`{
          "id": 8,
          "name": "accepted",
          "recommendation_label": "Acceptance",
          "css_class": "green",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": 1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 2
        }`),
	"motions/state:9": []byte(`{
          "id": 9,
          "name": "rejected",
          "recommendation_label": "Rejection",
          "css_class": "red",
          "restriction": [],
          "allow_support": false,
          "allow_create_poll": false,
          "allow_submitter_edit": false,
          "dont_set_identifier": false,
          "show_state_extension_field": false,
          "merge_amendment_into_final": -1,
          "show_recommendation_extension_field": false,
          "next_states_id": [],
          "workflow_id": 2
        }`),
	"motions/statute-paragraph:1": []byte(`{
          "id": 1,
          "title": "Statute",
          "text": "<p>test</p>",
          "weight": 10000
        }`),
	"motions/workflow:1": []byte(`{
          "id": 1,
          "name": "Simple Workflow",
          "states_id": [
            1,
            2,
            3,
            4
          ],
          "first_state_id": 1
        }`),
	"motions/workflow:2": []byte(`{
          "id": 2,
          "name": "Complex Workflow",
          "states_id": [
            5,
            6,
            7,
            8,
            9,
            10,
            11,
            12,
            13,
            14,
            15
          ],
          "first_state_id": 5
        }`),
	"topics/topic:1": []byte(`{
          "id": 1,
          "title": "Topic",
          "text": "",
          "attachments_id": [],
          "agenda_item_id": 1,
          "list_of_speakers_id": 1
        }`),
	"topics/topic:2": []byte(`{
          "id": 2,
          "title": "Hidden",
          "text": "",
          "attachments_id": [],
          "agenda_item_id": 2,
          "list_of_speakers_id": 2
        }`),
	"topics/topic:3": []byte(`{
          "id": 3,
          "title": "Internal",
          "text": "",
          "attachments_id": [],
          "agenda_item_id": 3,
          "list_of_speakers_id": 3
        }`),
	"topics/topic:4": []byte(`{
          "id": 4,
          "title": "Another public topic",
          "text": "",
          "attachments_id": [],
          "agenda_item_id": 9,
          "list_of_speakers_id": 14
        }`),
	"users/group:1": []byte(`{
          "id": 1,
          "name": "Default",
          "permissions": [
            "agenda.can_see",
            "agenda.can_see_internal_items",
            "agenda.can_see_list_of_speakers",
            "assignments.can_see",
            "core.can_see_frontpage",
            "core.can_see_projector",
            "mediafiles.can_see",
            "motions.can_see",
            "users.can_change_password"
          ]
        }`),
	"users/group:2": []byte(`{
          "id": 2,
          "name": "Admin",
          "permissions": []
        }`),
	"users/group:3": []byte(`{
          "id": 3,
          "name": "Delegates",
          "permissions": [
            "agenda.can_see",
            "agenda.can_see_internal_items",
            "agenda.can_see_list_of_speakers",
            "agenda.can_be_speaker",
            "assignments.can_nominate_other",
            "assignments.can_nominate_self",
            "assignments.can_see",
            "core.can_see_frontpage",
            "core.can_see_projector",
            "mediafiles.can_see",
            "motions.can_create",
            "motions.can_create_amendments",
            "motions.can_see",
            "motions.can_support",
            "users.can_change_password"
          ]
        }`),
	"users/group:4": []byte(`{
          "id": 4,
          "name": "Staff",
          "permissions": [
            "agenda.can_manage",
            "agenda.can_see",
            "agenda.can_see_internal_items",
            "agenda.can_manage_list_of_speakers",
            "agenda.can_see_list_of_speakers",
            "agenda.can_be_speaker",
            "assignments.can_manage",
            "assignments.can_nominate_other",
            "assignments.can_nominate_self",
            "assignments.can_see",
            "core.can_see_history",
            "core.can_manage_projector",
            "core.can_see_frontpage",
            "core.can_see_projector",
            "core.can_manage_tags",
            "mediafiles.can_manage",
            "mediafiles.can_see",
            "motions.can_create",
            "motions.can_create_amendments",
            "motions.can_manage",
            "motions.can_manage_metadata",
            "motions.can_see",
            "motions.can_see_internal",
            "users.can_change_password",
            "users.can_manage",
            "users.can_see_extra_data",
            "users.can_see_name"
          ]
        }`),
	"users/group:5": []byte(`{
          "id": 5,
          "name": "Committees",
          "permissions": [
            "agenda.can_see",
            "agenda.can_see_internal_items",
            "agenda.can_see_list_of_speakers",
            "assignments.can_see",
            "core.can_see_frontpage",
            "core.can_see_projector",
            "mediafiles.can_see",
            "motions.can_create",
            "motions.can_create_amendments",
            "motions.can_see",
            "motions.can_support",
            "users.can_change_password",
            "users.can_see_name"
          ]
        }`),
	"users/personal-note:1": []byte(`{
          "id": 1,
          "user_id": 1,
          "notes": {
            "motions/motion": {
              "3": {
                "note": "<p>test</p>",
                "star": false
              }
            }
          }
        }`),
	"users/personal-note:2": []byte(`{
          "id": 2,
          "user_id": 4,
          "notes": {
            "motions/motion": {
              "3": {
                "note": "<p>user a</p>",
                "star": false
              }
            }
          }
        }`),
	"users/user:1": []byte(`{
          "id": 1,
          "username": "admin",
          "title": "",
          "first_name": "",
          "last_name": "Administrator",
          "structure_level": "",
          "number": "",
          "about_me": "",
          "groups_id": [
            2
          ],
          "is_present": false,
          "is_committee": false,
          "vote_weight": "1.000000",
          "gender": "",
          "email": "",
          "last_email_send": null,
          "comment": "",
          "is_active": true,
          "auth_type": "default",
          "default_password": "admin",
          "session_auth_hash": "f5d40e336a5273852b538c2aa757fe0672c03846"
        }`),
	"users/user:2": []byte(`{
          "id": 2,
          "username": "candidate1",
          "title": "",
          "first_name": "candidate1",
          "last_name": "",
          "structure_level": "",
          "number": "",
          "about_me": "",
          "groups_id": [],
          "is_present": false,
          "is_committee": false,
          "vote_weight": "1.000000",
          "gender": "",
          "email": "",
          "last_email_send": null,
          "comment": "",
          "is_active": true,
          "auth_type": "default",
          "default_password": "8NpbvXCBDr",
          "session_auth_hash": "b1dd6c4ee19b31d45bf350a6ab77ce793280c4de"
        }`),
	"users/user:3": []byte(`{
          "id": 3,
          "username": "candidate2",
          "title": "",
          "first_name": "candidate2",
          "last_name": "",
          "structure_level": "",
          "number": "",
          "about_me": "",
          "groups_id": [],
          "is_present": false,
          "is_committee": false,
          "vote_weight": "1.000000",
          "gender": "",
          "email": "",
          "last_email_send": null,
          "comment": "",
          "is_active": true,
          "auth_type": "default",
          "default_password": "5YLEHrUUTG",
          "session_auth_hash": "741b6232747ed8949efcf627b6b5cabeca5a8341"
        }`),
	"users/user:4": []byte(`{
          "id": 4,
          "username": "a",
          "title": "",
          "first_name": "a",
          "last_name": "",
          "structure_level": "",
          "number": "",
          "about_me": "",
          "groups_id": [
            3
          ],
          "is_present": true,
          "is_committee": false,
          "vote_weight": "1.000000",
          "gender": "",
          "email": "",
          "last_email_send": null,
          "comment": "",
          "is_active": true,
          "auth_type": "default",
          "default_password": "a",
          "session_auth_hash": "ced21507b1f1414b8091fe705029f60ea6ae2ce2"
        }`),
	"users/user:5": []byte(`{
          "id": 5,
          "username": "b",
          "title": "",
          "first_name": "b",
          "last_name": "",
          "structure_level": "",
          "number": "",
          "about_me": "",
          "groups_id": [
            3
          ],
          "is_present": true,
          "is_committee": false,
          "vote_weight": "1.000000",
          "gender": "",
          "email": "",
          "last_email_send": null,
          "comment": "",
          "is_active": true,
          "auth_type": "default",
          "default_password": "b",
          "session_auth_hash": "776090ab96a2c4b20bca484dcd531f9c10925631"
        }`),
	"users/user:6": []byte(`{
          "id": 6,
          "username": "speaker1",
          "title": "title",
          "first_name": "speaker1",
          "last_name": "the last name",
          "structure_level": "layer X",
          "number": "3",
          "about_me": "",
          "groups_id": [],
          "is_present": true,
          "is_committee": false,
          "vote_weight": "1.000000",
          "gender": "",
          "email": "",
          "last_email_send": null,
          "comment": "",
          "is_active": true,
          "auth_type": "default",
          "default_password": "ZdbyxFDWpp",
          "session_auth_hash": "24cc3ec0d7b3d3028fb83dbc36bc7ded3b756e88"
        }`),
	"users/user:7": []byte(`{
          "id": 7,
          "username": "speaker2",
          "title": "",
          "first_name": "speaker2",
          "last_name": "",
          "structure_level": "",
          "number": "",
          "about_me": "",
          "groups_id": [],
          "is_present": true,
          "is_committee": false,
          "vote_weight": "1.000000",
          "gender": "",
          "email": "",
          "last_email_send": null,
          "comment": "",
          "is_active": true,
          "auth_type": "default",
          "default_password": "5HZr2zPM3x",
          "session_auth_hash": "54c0f8375d1af3b88125fddcf201817024ea0aa6"
        }`),
	"users/user:8": []byte(`{
          "id": 8,
          "username": "not required",
          "title": "",
          "first_name": "not required",
          "last_name": "",
          "structure_level": "",
          "number": "",
          "about_me": "",
          "groups_id": [],
          "is_present": true,
          "is_committee": false,
          "vote_weight": "1.000000",
          "gender": "",
          "email": "",
          "last_email_send": null,
          "comment": "",
          "is_active": true,
          "auth_type": "default",
          "default_password": "hq59FcgwLc",
          "session_auth_hash": "ad09d36968d35f31bf3397ee28433ae24494023c"
        }`),
}

var exampleRestrictedData = map[int]map[string]json.RawMessage{
	1: {
		"agenda/item:1": []byte(`{
            "id": 1,
            "item_number": "",
            "title_information": {
              "title": "Topic"
            },
            "comment": null,
            "closed": false,
            "type": 1,
            "is_internal": false,
            "is_hidden": false,
            "duration": null,
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            },
            "weight": 2,
            "parent_id": null,
            "level": 0,
            "tags_id": []
          }`),
		"agenda/item:2": []byte(`{
            "id": 2,
            "item_number": "",
            "title_information": {
              "title": "Hidden"
            },
            "comment": null,
            "closed": false,
            "type": 3,
            "is_internal": false,
            "is_hidden": true,
            "duration": null,
            "content_object": {
              "collection": "topics/topic",
              "id": 2
            },
            "weight": 4,
            "parent_id": null,
            "level": 0,
            "tags_id": []
          }`),
		"agenda/item:3": []byte(`{
            "id": 3,
            "item_number": "",
            "title_information": {
              "title": "Internal"
            },
            "comment": null,
            "closed": false,
            "type": 2,
            "is_internal": true,
            "is_hidden": false,
            "duration": null,
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            },
            "weight": 8,
            "parent_id": null,
            "level": 0,
            "tags_id": []
          }`),
		"agenda/item:4": []byte(`{
            "id": 4,
            "item_number": "",
            "title_information": {
              "title": "Assignment"
            },
            "comment": null,
            "closed": false,
            "type": 1,
            "is_internal": false,
            "is_hidden": true,
            "duration": null,
            "content_object": {
              "collection": "assignments/assignment",
              "id": 1
            },
            "weight": 6,
            "parent_id": 2,
            "level": 1,
            "tags_id": []
          }`),
		"agenda/item:5": []byte(`{
            "id": 5,
            "item_number": "",
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "comment": null,
            "closed": false,
            "type": 1,
            "is_internal": true,
            "is_hidden": false,
            "duration": null,
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            },
            "weight": 14,
            "parent_id": 3,
            "level": 1,
            "tags_id": []
          }`),
		"agenda/item:6": []byte(`{
            "id": 6,
            "item_number": "",
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "comment": null,
            "closed": false,
            "type": 1,
            "is_internal": true,
            "is_hidden": false,
            "duration": 0,
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            },
            "weight": 16,
            "parent_id": 3,
            "level": 1,
            "tags_id": []
          }`),
		"agenda/item:7": []byte(`{
            "id": 7,
            "item_number": "",
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "comment": null,
            "closed": false,
            "type": 2,
            "is_internal": true,
            "is_hidden": false,
            "duration": null,
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            },
            "weight": 18,
            "parent_id": 6,
            "level": 2,
            "tags_id": []
          }`),
		"agenda/item:8": []byte(`{
            "id": 8,
            "item_number": "",
            "title_information": {
              "title": "block"
            },
            "comment": null,
            "closed": false,
            "type": 3,
            "is_internal": true,
            "is_hidden": true,
            "duration": null,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 1
            },
            "weight": 10,
            "parent_id": 3,
            "level": 1,
            "tags_id": []
          }`),
		"agenda/item:9": []byte(`{
            "id": 9,
            "item_number": "",
            "title_information": {
              "title": "Another public topic"
            },
            "comment": null,
            "closed": false,
            "type": 1,
            "is_internal": true,
            "is_hidden": true,
            "duration": null,
            "content_object": {
              "collection": "topics/topic",
              "id": 4
            },
            "weight": 12,
            "parent_id": 8,
            "level": 2,
            "tags_id": []
          }`),
		"agenda/list-of-speakers:1": []byte(`{
            "id": 1,
            "title_information": {
              "title": "Topic"
            },
            "speakers": [
              {
                "id": 3,
                "user_id": 6,
                "begin_time": "2020-08-11T12:28:30.894574+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:10": []byte(`{
            "id": 10,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:11": []byte(`{
            "id": 11,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:12": []byte(`{
            "id": 12,
            "title_information": {
              "title": "block"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:13": []byte(`{
            "id": 13,
            "title_information": {
              "title": "block internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:14": []byte(`{
            "id": 14,
            "title_information": {
              "title": "Another public topic"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:2": []byte(`{
            "id": 2,
            "title_information": {
              "title": "Hidden"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:3": []byte(`{
            "id": 3,
            "title_information": {
              "title": "Internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:4": []byte(`{
            "id": 4,
            "title_information": {
              "title": "folder"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:5": []byte(`{
            "id": 5,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:6": []byte(`{
            "id": 6,
            "title_information": {
              "title": "in.jpg"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:7": []byte(`{
            "id": 7,
            "title_information": {
              "title": "Assignment"
            },
            "speakers": [
              {
                "id": 4,
                "user_id": 6,
                "begin_time": "2020-08-11T12:29:55.054553+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              },
              {
                "id": 6,
                "user_id": 7,
                "begin_time": null,
                "end_time": null,
                "weight": 2,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "assignments/assignment",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:8": []byte(`{
            "id": 8,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:9": []byte(`{
            "id": 9,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            }
          }`),
		"assignments/assignment-option:1": []byte(`{
            "user_id": 2,
            "weight": 1,
            "id": 1,
            "yes": "1.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-option:2": []byte(`{
            "user_id": 3,
            "weight": 3,
            "id": 2,
            "yes": "0.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-poll:1": []byte(`{
            "assignment_id": 1,
            "description": "",
            "pollmethod": "votes",
            "votes_amount": 1,
            "allow_multiple_votes_per_candidate": false,
            "global_no": true,
            "amount_global_no": "0.000000",
            "global_abstain": true,
            "amount_global_abstain": "0.000000",
            "state": 2,
            "type": "named",
            "title": "Wahlgang",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              1,
              2
            ],
            "id": 1,
            "onehundred_percent_base": "valid",
            "majority_method": "simple",
            "voted_id": [
              4
            ],
            "user_has_voted": false
          }`),
		"assignments/assignment-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "Y",
            "user_id": 4,
            "option_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment:1": []byte(`{
            "id": 1,
            "title": "Assignment",
            "description": "",
            "open_posts": 1,
            "phase": 0,
            "assignment_related_users": [
              {
                "id": 1,
                "user_id": 2,
                "weight": 1
              },
              {
                "id": 3,
                "user_id": 3,
                "weight": 3
              }
            ],
            "default_poll_description": "",
            "agenda_item_id": 4,
            "list_of_speakers_id": 7,
            "tags_id": [
              2
            ],
            "attachments_id": [],
            "number_poll_candidates": false,
            "polls_id": [
              1
            ]
          }`),
		"core/config:10": []byte(`{
            "id": 10,
            "key": "general_system_conference_show",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show live conference window",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 140,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:100": []byte(`{
            "id": 100,
            "key": "motion_poll_default_100_percent_base",
            "value": "YNA",
            "data": {
              "defaultValue": "YNA",
              "inputType": "choice",
              "label": "Default 100 % base of a voting result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 370,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:101": []byte(`{
            "id": 101,
            "key": "motion_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:102": []byte(`{
            "id": 102,
            "key": "motion_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 372,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:103": []byte(`{
            "id": 103,
            "key": "motions_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 373,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:104": []byte(`{
            "id": 104,
            "key": "motions_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 374,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:105": []byte(`{
            "id": 105,
            "key": "motions_export_title",
            "value": "Motions",
            "data": {
              "defaultValue": "Motions",
              "inputType": "string",
              "label": "Title for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 380,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:106": []byte(`{
            "id": 106,
            "key": "motions_export_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 382,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:107": []byte(`{
            "id": 107,
            "key": "motions_export_submitter_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show submitters and recommendation/state in table of contents",
              "helpText": "",
              "choices": null,
              "weight": 384,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:108": []byte(`{
            "id": 108,
            "key": "motions_export_follow_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show checkbox to record decision",
              "helpText": "",
              "choices": null,
              "weight": 386,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:109": []byte(`{
            "id": 109,
            "key": "assignment_poll_method",
            "value": "votes",
            "data": {
              "defaultValue": "votes",
              "inputType": "choice",
              "label": "Default election method",
              "helpText": "",
              "choices": [
                {
                  "value": "votes",
                  "display_name": "Yes per candidate"
                },
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                }
              ],
              "weight": 400,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:11": []byte(`{
            "id": 11,
            "key": "general_system_conference_auto_connect",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Connect all users to live conference automatically",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 141,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:110": []byte(`{
            "id": 110,
            "key": "assignment_poll_default_type",
            "value": "analog",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 403,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:111": []byte(`{
            "id": 111,
            "key": "assignment_poll_default_100_percent_base",
            "value": "valid",
            "data": {
              "defaultValue": "valid",
              "inputType": "choice",
              "label": "Default 100 % base of an election result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                },
                {
                  "value": "votes",
                  "display_name": "Sum of votes including general No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 405,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:112": []byte(`{
            "id": 112,
            "key": "assignment_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 410,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:113": []byte(`{
            "id": 113,
            "key": "assignment_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:114": []byte(`{
            "id": 114,
            "key": "assignment_poll_sort_poll_result_by_votes",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Sort election results by amount of votes",
              "helpText": "",
              "choices": null,
              "weight": 420,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:115": []byte(`{
            "id": 115,
            "key": "assignment_poll_add_candidates_to_list_of_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Put all candidates on the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 425,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:116": []byte(`{
            "id": 116,
            "key": "assignments_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 430,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:117": []byte(`{
            "id": 117,
            "key": "assignments_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 435,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:118": []byte(`{
            "id": 118,
            "key": "assignments_pdf_title",
            "value": "Elections",
            "data": {
              "defaultValue": "Elections",
              "inputType": "string",
              "label": "Title for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 460,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:119": []byte(`{
            "id": 119,
            "key": "assignments_pdf_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 470,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:12": []byte(`{
            "id": 12,
            "key": "general_system_conference_los_restriction",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow only current speakers and list of speakers managers to enter the live conference",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 142,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:13": []byte(`{
            "id": 13,
            "key": "general_system_stream_url",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Livestream url",
              "helpText": "Remove URL to deactivate livestream. Check extra group permission to see livestream.",
              "choices": null,
              "weight": 143,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:14": []byte(`{
            "id": 14,
            "key": "general_system_enable_anonymous",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow access for anonymous guest users",
              "helpText": "",
              "choices": null,
              "weight": 150,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:15": []byte(`{
            "id": 15,
            "key": "general_login_info_text",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Show this text on the login page",
              "helpText": "",
              "choices": null,
              "weight": 152,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:16": []byte(`{
            "id": 16,
            "key": "openslides_theme",
            "value": "openslides-default-light-theme",
            "data": {
              "defaultValue": "openslides-default-light-theme",
              "inputType": "choice",
              "label": "OpenSlides Theme",
              "helpText": "",
              "choices": [
                {
                  "value": "openslides-default-light-theme",
                  "display_name": "OpenSlides Default"
                },
                {
                  "value": "openslides-default-dark-theme",
                  "display_name": "OpenSlides Dark"
                },
                {
                  "value": "openslides-red-light-theme",
                  "display_name": "OpenSlides Red"
                },
                {
                  "value": "openslides-red-dark-theme",
                  "display_name": "OpenSlides Red Dark"
                },
                {
                  "value": "openslides-green-light-theme",
                  "display_name": "OpenSlides Green"
                },
                {
                  "value": "openslides-green-dark-theme",
                  "display_name": "OpenSlides Green Dark"
                },
                {
                  "value": "openslides-solarized-dark-theme",
                  "display_name": "OpenSlides Solarized"
                }
              ],
              "weight": 154,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:17": []byte(`{
            "id": 17,
            "key": "general_csv_separator",
            "value": ",",
            "data": {
              "defaultValue": ",",
              "inputType": "string",
              "label": "Separator used for all csv exports and examples",
              "helpText": "",
              "choices": null,
              "weight": 160,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:18": []byte(`{
            "id": 18,
            "key": "general_csv_encoding",
            "value": "utf-8",
            "data": {
              "defaultValue": "utf-8",
              "inputType": "choice",
              "label": "Default encoding for all csv exports",
              "helpText": "",
              "choices": [
                {
                  "value": "utf-8",
                  "display_name": "UTF-8"
                },
                {
                  "value": "iso-8859-15",
                  "display_name": "ISO-8859-15"
                }
              ],
              "weight": 162,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:19": []byte(`{
            "id": 19,
            "key": "general_export_pdf_pagenumber_alignment",
            "value": "center",
            "data": {
              "defaultValue": "center",
              "inputType": "choice",
              "label": "Page number alignment in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "left",
                  "display_name": "Left"
                },
                {
                  "value": "center",
                  "display_name": "Center"
                },
                {
                  "value": "right",
                  "display_name": "Right"
                }
              ],
              "weight": 164,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:2": []byte(`{
            "id": 2,
            "key": "general_event_name",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Event name",
              "helpText": "",
              "choices": null,
              "weight": 110,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:20": []byte(`{
            "id": 20,
            "key": "general_export_pdf_fontsize",
            "value": "10",
            "data": {
              "defaultValue": "10",
              "inputType": "choice",
              "label": "Standard font size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "10",
                  "display_name": "10"
                },
                {
                  "value": "11",
                  "display_name": "11"
                },
                {
                  "value": "12",
                  "display_name": "12"
                }
              ],
              "weight": 166,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:21": []byte(`{
            "id": 21,
            "key": "general_export_pdf_pagesize",
            "value": "A4",
            "data": {
              "defaultValue": "A4",
              "inputType": "choice",
              "label": "Standard page size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "A4",
                  "display_name": "DIN A4"
                },
                {
                  "value": "A5",
                  "display_name": "DIN A5"
                }
              ],
              "weight": 168,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:22": []byte(`{
            "id": 22,
            "key": "logos_available",
            "value": [
              "logo_projector_main",
              "logo_projector_header",
              "logo_web_header",
              "logo_pdf_header_L",
              "logo_pdf_header_R",
              "logo_pdf_footer_L",
              "logo_pdf_footer_R",
              "logo_pdf_ballot_paper"
            ],
            "data": null
          }`),
		"core/config:23": []byte(`{
            "id": 23,
            "key": "logo_projector_main",
            "value": {
              "display_name": "Projector logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:24": []byte(`{
            "id": 24,
            "key": "logo_projector_header",
            "value": {
              "display_name": "Projector header image",
              "path": ""
            },
            "data": null
          }`),
		"core/config:25": []byte(`{
            "id": 25,
            "key": "logo_web_header",
            "value": {
              "display_name": "Web interface header logo",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:26": []byte(`{
            "id": 26,
            "key": "logo_pdf_header_L",
            "value": {
              "display_name": "PDF header logo (left)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:27": []byte(`{
            "id": 27,
            "key": "logo_pdf_header_R",
            "value": {
              "display_name": "PDF header logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:28": []byte(`{
            "id": 28,
            "key": "logo_pdf_footer_L",
            "value": {
              "display_name": "PDF footer logo (left)",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:29": []byte(`{
            "id": 29,
            "key": "logo_pdf_footer_R",
            "value": {
              "display_name": "PDF footer logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:3": []byte(`{
            "id": 3,
            "key": "general_event_description",
            "value": "Presentation and assembly system",
            "data": {
              "defaultValue": "Presentation and assembly system",
              "inputType": "string",
              "label": "Short description of event",
              "helpText": "",
              "choices": null,
              "weight": 115,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:30": []byte(`{
            "id": 30,
            "key": "logo_pdf_ballot_paper",
            "value": {
              "display_name": "PDF ballot paper logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:31": []byte(`{
            "id": 31,
            "key": "fonts_available",
            "value": [
              "font_regular",
              "font_italic",
              "font_bold",
              "font_bold_italic",
              "font_monospace"
            ],
            "data": null
          }`),
		"core/config:32": []byte(`{
            "id": 32,
            "key": "font_regular",
            "value": {
              "display_name": "Font regular",
              "default": "assets/fonts/fira-sans-latin-400.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:33": []byte(`{
            "id": 33,
            "key": "font_italic",
            "value": {
              "display_name": "Font italic",
              "default": "assets/fonts/fira-sans-latin-400italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:34": []byte(`{
            "id": 34,
            "key": "font_bold",
            "value": {
              "display_name": "Font bold",
              "default": "assets/fonts/fira-sans-latin-500.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:35": []byte(`{
            "id": 35,
            "key": "font_bold_italic",
            "value": {
              "display_name": "Font bold italic",
              "default": "assets/fonts/fira-sans-latin-500italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:36": []byte(`{
            "id": 36,
            "key": "font_monospace",
            "value": {
              "display_name": "Font monospace",
              "default": "assets/fonts/roboto-condensed-bold.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:37": []byte(`{
            "id": 37,
            "key": "translations",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "translations",
              "label": "Custom translations",
              "helpText": "",
              "choices": null,
              "weight": 1000,
              "group": "Custom translations",
              "subgroup": "General"
            }
          }`),
		"core/config:38": []byte(`{
            "id": 38,
            "key": "config_version",
            "value": 2,
            "data": null
          }`),
		"core/config:39": []byte(`{
            "id": 39,
            "key": "db_id",
            "value": "2c3bd736c87a48a4be1f0dc707702144",
            "data": null
          }`),
		"core/config:4": []byte(`{
            "id": 4,
            "key": "general_event_date",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event date",
              "helpText": "",
              "choices": null,
              "weight": 120,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:40": []byte(`{
            "id": 40,
            "key": "users_sort_by",
            "value": "first_name",
            "data": {
              "defaultValue": "first_name",
              "inputType": "choice",
              "label": "Sort name of participants by",
              "helpText": "",
              "choices": [
                {
                  "value": "first_name",
                  "display_name": "Given name"
                },
                {
                  "value": "last_name",
                  "display_name": "Surname"
                },
                {
                  "value": "number",
                  "display_name": "Participant number"
                }
              ],
              "weight": 510,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:41": []byte(`{
            "id": 41,
            "key": "users_enable_presence_view",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Enable participant presence view",
              "helpText": "",
              "choices": null,
              "weight": 511,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:42": []byte(`{
            "id": 42,
            "key": "users_allow_self_set_present",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow users to set themselves as present",
              "helpText": "e.g. for online meetings",
              "choices": null,
              "weight": 512,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:43": []byte(`{
            "id": 43,
            "key": "users_activate_vote_weight",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate vote weight",
              "helpText": "",
              "choices": null,
              "weight": 513,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:44": []byte(`{
            "id": 44,
            "key": "users_pdf_welcometitle",
            "value": "Welcome to OpenSlides",
            "data": {
              "defaultValue": "Welcome to OpenSlides",
              "inputType": "string",
              "label": "Title for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 520,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:45": []byte(`{
            "id": 45,
            "key": "users_pdf_welcometext",
            "value": "[Place for your welcome and help text.]",
            "data": {
              "defaultValue": "[Place for your welcome and help text.]",
              "inputType": "string",
              "label": "Help text for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 530,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:46": []byte(`{
            "id": 46,
            "key": "users_pdf_url",
            "value": "http://example.com:8000",
            "data": {
              "defaultValue": "http://example.com:8000",
              "inputType": "string",
              "label": "System URL",
              "helpText": "Used for QRCode in PDF of access data.",
              "choices": null,
              "weight": 540,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:47": []byte(`{
            "id": 47,
            "key": "users_pdf_wlan_ssid",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN name (SSID)",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 550,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:48": []byte(`{
            "id": 48,
            "key": "users_pdf_wlan_password",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN password",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 560,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:49": []byte(`{
            "id": 49,
            "key": "users_pdf_wlan_encryption",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "choice",
              "label": "WLAN encryption",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": [
                {
                  "value": "",
                  "display_name": "---------"
                },
                {
                  "value": "WEP",
                  "display_name": "WEP"
                },
                {
                  "value": "WPA",
                  "display_name": "WPA/WPA2"
                },
                {
                  "value": "nopass",
                  "display_name": "No encryption"
                }
              ],
              "weight": 570,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:5": []byte(`{
            "id": 5,
            "key": "general_event_location",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event location",
              "helpText": "",
              "choices": null,
              "weight": 125,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:50": []byte(`{
            "id": 50,
            "key": "users_email_sender",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Sender name",
              "helpText": "The sender address is defined in the OpenSlides server settings and should modified by administrator only.",
              "choices": null,
              "weight": 600,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:51": []byte(`{
            "id": 51,
            "key": "users_email_replyto",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Reply address",
              "helpText": "",
              "choices": null,
              "weight": 601,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:52": []byte(`{
            "id": 52,
            "key": "users_email_subject",
            "value": "OpenSlides access data",
            "data": {
              "defaultValue": "OpenSlides access data",
              "inputType": "string",
              "label": "Email subject",
              "helpText": "You can use {event_name} and {username} as placeholder.",
              "choices": null,
              "weight": 605,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:53": []byte(`{
            "id": 53,
            "key": "users_email_body",
            "value": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
            "data": {
              "defaultValue": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
              "inputType": "text",
              "label": "Email body",
              "helpText": "Use these placeholders: {name}, {event_name}, {url}, {username}, {password}. The url referrs to the system url.",
              "choices": null,
              "weight": 610,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:54": []byte(`{
            "id": 54,
            "key": "agenda_start_event_date_time",
            "value": null,
            "data": {
              "defaultValue": null,
              "inputType": "datetimepicker",
              "label": "Begin of event",
              "helpText": "Input format: DD.MM.YYYY HH:MM",
              "choices": null,
              "weight": 200,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:55": []byte(`{
            "id": 55,
            "key": "agenda_show_subtitle",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show subtitles in the agenda",
              "helpText": "",
              "choices": null,
              "weight": 201,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:56": []byte(`{
            "id": 56,
            "key": "agenda_enable_numbering",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Enable numbering for agenda items",
              "helpText": "",
              "choices": null,
              "weight": 205,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:57": []byte(`{
            "id": 57,
            "key": "agenda_number_prefix",
            "value": "TOP",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Numbering prefix for agenda items",
              "helpText": "This prefix will be set if you run the automatic agenda numbering.",
              "choices": null,
              "weight": 206,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:58": []byte(`{
            "id": 58,
            "key": "agenda_numeral_system",
            "value": "arabic",
            "data": {
              "defaultValue": "arabic",
              "inputType": "choice",
              "label": "Numeral system for agenda items",
              "helpText": "",
              "choices": [
                {
                  "value": "arabic",
                  "display_name": "Arabic"
                },
                {
                  "value": "roman",
                  "display_name": "Roman"
                }
              ],
              "weight": 207,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:59": []byte(`{
            "id": 59,
            "key": "agenda_item_creation",
            "value": "default_yes",
            "data": {
              "defaultValue": "default_yes",
              "inputType": "choice",
              "label": "Add to agenda",
              "helpText": "",
              "choices": [
                {
                  "value": "always",
                  "display_name": "Always"
                },
                {
                  "value": "never",
                  "display_name": "Never"
                },
                {
                  "value": "default_yes",
                  "display_name": "Ask, default yes"
                },
                {
                  "value": "default_no",
                  "display_name": "Ask, default no"
                }
              ],
              "weight": 210,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:6": []byte(`{
            "id": 6,
            "key": "general_event_legal_notice",
            "value": "<a href=\"http://www.openslides.org\">OpenSlides</a> is a free web based presentation and assembly system for visualizing and controlling agenda, motions and elections of an assembly.",
            "data": null
          }`),
		"core/config:60": []byte(`{
            "id": 60,
            "key": "agenda_new_items_default_visibility",
            "value": "2",
            "data": {
              "defaultValue": "2",
              "inputType": "choice",
              "label": "Default visibility for new agenda items (except topics)",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Public item"
                },
                {
                  "value": "2",
                  "display_name": "Internal item"
                },
                {
                  "value": "3",
                  "display_name": "Hidden item"
                }
              ],
              "weight": 211,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:61": []byte(`{
            "id": 61,
            "key": "agenda_hide_internal_items_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Hide internal items when projecting subitems",
              "helpText": "",
              "choices": null,
              "weight": 212,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:62": []byte(`{
            "id": 62,
            "key": "agenda_show_last_speakers",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of last speakers to be shown on the projector",
              "helpText": "",
              "choices": null,
              "weight": 220,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:63": []byte(`{
            "id": 63,
            "key": "agenda_show_next_speakers",
            "value": -1,
            "data": {
              "defaultValue": -1,
              "inputType": "integer",
              "label": "Number of the next speakers to be shown on the projector",
              "helpText": "Enter number of the next shown speakers. Choose -1 to show all next speakers.",
              "choices": null,
              "weight": 222,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:64": []byte(`{
            "id": 64,
            "key": "agenda_countdown_warning_time",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Show orange countdown in the last x seconds of speaking time",
              "helpText": "Enter duration in seconds. Choose 0 to disable warning color.",
              "choices": null,
              "weight": 224,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:65": []byte(`{
            "id": 65,
            "key": "projector_default_countdown",
            "value": 60,
            "data": {
              "defaultValue": 60,
              "inputType": "integer",
              "label": "Predefined seconds of new countdowns",
              "helpText": "",
              "choices": null,
              "weight": 226,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:66": []byte(`{
            "id": 66,
            "key": "agenda_couple_countdown_and_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Couple countdown with the list of speakers",
              "helpText": "[Begin speech] starts the countdown, [End speech] stops the countdown.",
              "choices": null,
              "weight": 228,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:67": []byte(`{
            "id": 67,
            "key": "agenda_hide_amount_of_speakers",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide the amount of speakers in subtitle of list of speakers slide",
              "helpText": "",
              "choices": null,
              "weight": 230,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:68": []byte(`{
            "id": 68,
            "key": "agenda_present_speakers_only",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Only present participants can be added to the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 232,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:69": []byte(`{
            "id": 69,
            "key": "agenda_show_first_contribution",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show hint »first speech« in the list of speakers management view",
              "helpText": "",
              "choices": null,
              "weight": 234,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:7": []byte(`{
            "id": 7,
            "key": "general_event_privacy_policy",
            "value": "",
            "data": null
          }`),
		"core/config:70": []byte(`{
            "id": 70,
            "key": "motions_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new motions",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 310,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:71": []byte(`{
            "id": 71,
            "key": "motions_statute_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new statute amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 312,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:72": []byte(`{
            "id": 72,
            "key": "motions_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 314,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:73": []byte(`{
            "id": 73,
            "key": "motions_preamble",
            "value": "The assembly may decide:",
            "data": {
              "defaultValue": "The assembly may decide:",
              "inputType": "string",
              "label": "Motion preamble",
              "helpText": "",
              "choices": null,
              "weight": 320,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:74": []byte(`{
            "id": 74,
            "key": "motions_default_line_numbering",
            "value": "outside",
            "data": {
              "defaultValue": "outside",
              "inputType": "choice",
              "label": "Default line numbering",
              "helpText": "",
              "choices": [
                {
                  "value": "outside",
                  "display_name": "outside"
                },
                {
                  "value": "inline",
                  "display_name": "inline"
                },
                {
                  "value": "none",
                  "display_name": "Disabled"
                }
              ],
              "weight": 322,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:75": []byte(`{
            "id": 75,
            "key": "motions_line_length",
            "value": 85,
            "data": {
              "defaultValue": 85,
              "inputType": "integer",
              "label": "Line length",
              "helpText": "The maximum number of characters per line. Relevant when line numbering is enabled. Min: 40",
              "choices": null,
              "weight": 323,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:76": []byte(`{
            "id": 76,
            "key": "motions_reason_required",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Reason required for creating new motion",
              "helpText": "",
              "choices": null,
              "weight": 324,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:77": []byte(`{
            "id": 77,
            "key": "motions_disable_text_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide motion text on projector",
              "helpText": "",
              "choices": null,
              "weight": 325,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:78": []byte(`{
            "id": 78,
            "key": "motions_disable_reason_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide reason on projector",
              "helpText": "",
              "choices": null,
              "weight": 326,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:79": []byte(`{
            "id": 79,
            "key": "motions_disable_recommendation_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide recommendation on projector",
              "helpText": "",
              "choices": null,
              "weight": 327,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:8": []byte(`{
            "id": 8,
            "key": "general_event_welcome_title",
            "value": "Welcome to OpenSlides",
            "data": null
          }`),
		"core/config:80": []byte(`{
            "id": 80,
            "key": "motions_hide_referring_motions",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide referring motions",
              "helpText": "",
              "choices": null,
              "weight": 328,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:81": []byte(`{
            "id": 81,
            "key": "motions_disable_sidebox_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show meta information box below the title on projector",
              "helpText": "",
              "choices": null,
              "weight": 329,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:82": []byte(`{
            "id": 82,
            "key": "motions_show_sequential_numbers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show the sequential number for a motion",
              "helpText": "In motion list, motion detail and PDF.",
              "choices": null,
              "weight": 330,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:83": []byte(`{
            "id": 83,
            "key": "motions_recommendations_by",
            "value": "ABK",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender",
              "helpText": "Will be displayed as label before selected recommendation. Use an empty value to disable the recommendation system.",
              "choices": null,
              "weight": 332,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:84": []byte(`{
            "id": 84,
            "key": "motions_statute_recommendations_by",
            "value": "ABK2",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender for statute amendments",
              "helpText": "Will be displayed as label before selected recommendation in statute amendments.",
              "choices": null,
              "weight": 333,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:85": []byte(`{
            "id": 85,
            "key": "motions_recommendation_text_mode",
            "value": "diff",
            "data": {
              "defaultValue": "diff",
              "inputType": "choice",
              "label": "Default text version for change recommendations",
              "helpText": "",
              "choices": [
                {
                  "value": "original",
                  "display_name": "Original version"
                },
                {
                  "value": "changed",
                  "display_name": "Changed version"
                },
                {
                  "value": "diff",
                  "display_name": "Diff version"
                },
                {
                  "value": "agreed",
                  "display_name": "Final version"
                }
              ],
              "weight": 334,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:86": []byte(`{
            "id": 86,
            "key": "motions_motions_sorting",
            "value": "identifier",
            "data": {
              "defaultValue": "identifier",
              "inputType": "choice",
              "label": "Sort motions by",
              "helpText": "",
              "choices": [
                {
                  "value": "weight",
                  "display_name": "Call list"
                },
                {
                  "value": "identifier",
                  "display_name": "Identifier"
                }
              ],
              "weight": 335,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:87": []byte(`{
            "id": 87,
            "key": "motions_identifier",
            "value": "per_category",
            "data": {
              "defaultValue": "per_category",
              "inputType": "choice",
              "label": "Identifier",
              "helpText": "",
              "choices": [
                {
                  "value": "per_category",
                  "display_name": "Numbered per category"
                },
                {
                  "value": "serially_numbered",
                  "display_name": "Serially numbered"
                },
                {
                  "value": "manually",
                  "display_name": "Set it manually"
                }
              ],
              "weight": 340,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:88": []byte(`{
            "id": 88,
            "key": "motions_identifier_min_digits",
            "value": 1,
            "data": {
              "defaultValue": 1,
              "inputType": "integer",
              "label": "Number of minimal digits for identifier",
              "helpText": "Uses leading zeros to sort motions correctly by identifier.",
              "choices": null,
              "weight": 342,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:89": []byte(`{
            "id": 89,
            "key": "motions_identifier_with_blank",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow blank in identifier",
              "helpText": "Blank between prefix and number, e.g. 'A 001'.",
              "choices": null,
              "weight": 344,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:9": []byte(`{
            "id": 9,
            "key": "general_event_welcome_text",
            "value": "[Space for your welcome text.]",
            "data": null
          }`),
		"core/config:90": []byte(`{
            "id": 90,
            "key": "motions_statutes_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate statute amendments",
              "helpText": "",
              "choices": null,
              "weight": 350,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:91": []byte(`{
            "id": 91,
            "key": "motions_amendments_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate amendments",
              "helpText": "",
              "choices": null,
              "weight": 351,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:92": []byte(`{
            "id": 92,
            "key": "motions_amendments_main_table",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show amendments together with motions",
              "helpText": "",
              "choices": null,
              "weight": 352,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:93": []byte(`{
            "id": 93,
            "key": "motions_amendments_prefix",
            "value": "Ä-",
            "data": {
              "defaultValue": "-",
              "inputType": "string",
              "label": "Prefix for the identifier for amendments",
              "helpText": "",
              "choices": null,
              "weight": 353,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:94": []byte(`{
            "id": 94,
            "key": "motions_amendments_text_mode",
            "value": "paragraph",
            "data": {
              "defaultValue": "paragraph",
              "inputType": "choice",
              "label": "How to create new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "freestyle",
                  "display_name": "Empty text field"
                },
                {
                  "value": "fulltext",
                  "display_name": "Edit the whole motion text"
                },
                {
                  "value": "paragraph",
                  "display_name": "Paragraph-based, Diff-enabled"
                }
              ],
              "weight": 354,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:95": []byte(`{
            "id": 95,
            "key": "motions_amendments_multiple_paragraphs",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Amendments can change multiple paragraphs",
              "helpText": "",
              "choices": null,
              "weight": 355,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:96": []byte(`{
            "id": 96,
            "key": "motions_amendments_of_amendments",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow amendments of amendments",
              "helpText": "",
              "choices": null,
              "weight": 356,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:97": []byte(`{
            "id": 97,
            "key": "motions_min_supporters",
            "value": 1,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of (minimum) required supporters for a motion",
              "helpText": "Choose 0 to disable the supporting system.",
              "choices": null,
              "weight": 360,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:98": []byte(`{
            "id": 98,
            "key": "motions_remove_supporters",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Remove all supporters of a motion if a submitter edits his motion in early state",
              "helpText": "",
              "choices": null,
              "weight": 361,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:99": []byte(`{
            "id": 99,
            "key": "motion_poll_default_type",
            "value": "named",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 367,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/countdown:1": []byte(`{
            "id": 1,
            "title": "Default countdown",
            "description": "",
            "default_time": 60,
            "countdown_time": 1597141855.090748,
            "running": true
          }`),
		"core/projection-default:1": []byte(`{
            "id": 1,
            "name": "agenda_all_items",
            "display_name": "Agenda",
            "projector_id": 1
          }`),
		"core/projection-default:10": []byte(`{
            "id": 10,
            "name": "messages",
            "display_name": "Messages",
            "projector_id": 1
          }`),
		"core/projection-default:11": []byte(`{
            "id": 11,
            "name": "countdowns",
            "display_name": "Countdowns",
            "projector_id": 1
          }`),
		"core/projection-default:12": []byte(`{
            "id": 12,
            "name": "assignment_poll",
            "display_name": "Ballots",
            "projector_id": 1
          }`),
		"core/projection-default:13": []byte(`{
            "id": 13,
            "name": "motion_poll",
            "display_name": "Motion votes",
            "projector_id": 1
          }`),
		"core/projection-default:2": []byte(`{
            "id": 2,
            "name": "topics",
            "display_name": "Topics",
            "projector_id": 1
          }`),
		"core/projection-default:3": []byte(`{
            "id": 3,
            "name": "agenda_list_of_speakers",
            "display_name": "List of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:4": []byte(`{
            "id": 4,
            "name": "agenda_current_list_of_speakers",
            "display_name": "Current list of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:5": []byte(`{
            "id": 5,
            "name": "motions",
            "display_name": "Motions",
            "projector_id": 1
          }`),
		"core/projection-default:6": []byte(`{
            "id": 6,
            "name": "motionBlocks",
            "display_name": "Motion blocks",
            "projector_id": 1
          }`),
		"core/projection-default:7": []byte(`{
            "id": 7,
            "name": "assignments",
            "display_name": "Elections",
            "projector_id": 1
          }`),
		"core/projection-default:8": []byte(`{
            "id": 8,
            "name": "users",
            "display_name": "Participants",
            "projector_id": 1
          }`),
		"core/projection-default:9": []byte(`{
            "id": 9,
            "name": "mediafiles",
            "display_name": "Files",
            "projector_id": 1
          }`),
		"core/projector-message:1": []byte(`{
            "id": 1,
            "message": "<p>test</p>"
          }`),
		"core/projector:1": []byte(`{
            "id": 1,
            "elements": [
              {
                "name": "mediafiles/mediafile",
                "id": 3
              }
            ],
            "elements_preview": [],
            "elements_history": [
              [
                {
                  "name": "assignments/assignment",
                  "id": 1
                }
              ]
            ],
            "scale": 0,
            "scroll": 0,
            "name": "Default projector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [
              1,
              2,
              3,
              4,
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13
            ],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/projector:2": []byte(`{
            "id": 2,
            "elements": [
              {
                "name": "core/clock",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers",
                "stable": false
              },
              {
                "name": "agenda/current-speaker-chyron",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers-overlay",
                "stable": true
              }
            ],
            "elements_preview": [],
            "elements_history": [],
            "scale": 0,
            "scroll": 0,
            "name": "sideprojector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/tag:1": []byte(`{
            "id": 1,
            "name": "T1"
          }`),
		"core/tag:2": []byte(`{
            "id": 2,
            "name": "T2"
          }`),
		"mediafiles/mediafile:1": []byte(`{
            "id": 1,
            "title": "folder",
            "media_url_prefix": "/media/",
            "filesize": "",
            "mimetype": "",
            "pdf_information": {},
            "access_groups_id": [
              3
            ],
            "create_timestamp": "2020-08-11T11:15:50.719490+02:00",
            "is_directory": true,
            "path": "folder/",
            "parent_id": null,
            "list_of_speakers_id": 4,
            "inherited_access_groups_id": [
              3
            ]
          }`),
		"mediafiles/mediafile:2": []byte(`{
            "id": 2,
            "title": "A.txt",
            "media_url_prefix": "/media/",
            "filesize": "< 1 kB",
            "mimetype": "text/plain",
            "pdf_information": {},
            "access_groups_id": [],
            "create_timestamp": "2020-08-11T11:16:07.013124+02:00",
            "is_directory": false,
            "path": "A.txt",
            "parent_id": null,
            "list_of_speakers_id": 5,
            "inherited_access_groups_id": true
          }`),
		"mediafiles/mediafile:3": []byte(`{
            "id": 3,
            "title": "in.jpg",
            "media_url_prefix": "/media/",
            "filesize": "122 kB",
            "mimetype": "image/jpeg",
            "pdf_information": {},
            "access_groups_id": [
              4
            ],
            "create_timestamp": "2020-08-11T11:16:29.302184+02:00",
            "is_directory": false,
            "path": "folder/in.jpg",
            "parent_id": 1,
            "list_of_speakers_id": 6,
            "inherited_access_groups_id": false
          }`),
		"mediafiles/mediafile:4": []byte(`{
            "id": 4,
            "title": "A.txt",
            "media_url_prefix": "/media/",
            "filesize": "< 1 kB",
            "mimetype": "text/plain",
            "pdf_information": {},
            "access_groups_id": [
              4
            ],
            "create_timestamp": "2020-08-11T12:27:14.168247+02:00",
            "is_directory": false,
            "path": "folder/A.txt",
            "parent_id": 1,
            "list_of_speakers_id": 11,
            "inherited_access_groups_id": false
          }`),
		"motions/category:1": []byte(`{
            "id": 1,
            "name": "first",
            "prefix": "A",
            "parent_id": null,
            "weight": 2,
            "level": 0
          }`),
		"motions/category:2": []byte(`{
            "id": 2,
            "name": "second",
            "prefix": "B",
            "parent_id": 1,
            "weight": 4,
            "level": 1
          }`),
		"motions/category:3": []byte(`{
            "id": 3,
            "name": "third",
            "prefix": "C",
            "parent_id": null,
            "weight": 6,
            "level": 0
          }`),
		"motions/motion-block:1": []byte(`{
            "id": 1,
            "title": "block",
            "agenda_item_id": 8,
            "list_of_speakers_id": 12,
            "internal": false,
            "motions_id": [
              1,
              2
            ]
          }`),
		"motions/motion-block:2": []byte(`{
            "id": 2,
            "title": "block internal",
            "agenda_item_id": null,
            "list_of_speakers_id": 13,
            "internal": true,
            "motions_id": [
              3
            ]
          }`),
		"motions/motion-change-recommendation:1": []byte(`{
            "id": 1,
            "motion_id": 1,
            "rejected": false,
            "internal": true,
            "type": 2,
            "other_description": "",
            "line_from": 9,
            "line_to": 11,
            "text": "<p class=\"os-split-before os-split-after\">eleifend ac, enim. Quisque rutrum. Aenean imperdiet.</p>",
            "creation_time": "2020-08-11T11:39:08.232146+02:00"
          }`),
		"motions/motion-comment-section:1": []byte(`{
            "id": 1,
            "name": "public",
            "read_groups_id": [
              2,
              3
            ],
            "write_groups_id": [
              2,
              3
            ],
            "weight": 10000
          }`),
		"motions/motion-comment-section:2": []byte(`{
            "id": 2,
            "name": "internal",
            "read_groups_id": [],
            "write_groups_id": [],
            "weight": 10000
          }`),
		"motions/motion-option:1": []byte(`{
            "id": 1,
            "yes": "0.000000",
            "no": "1.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-option:2": []byte(`{
            "id": 2,
            "yes": "1.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 2,
            "pollstate": 4
          }`),
		"motions/motion-poll:1": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "named",
            "title": "Abstimmung",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              1
            ],
            "id": 1,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-poll:2": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "pseudoanonymous",
            "title": "Abstimmung (2)",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              2
            ],
            "id": 2,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "N",
            "user_id": 5,
            "option_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-vote:2": []byte(`{
            "id": 2,
            "weight": "1.000000",
            "value": "Y",
            "user_id": null,
            "option_id": 2,
            "pollstate": 4
          }`),
		"motions/motion:1": []byte(`{
            "id": 1,
            "identifier": null,
            "title": "Leadmotion1",
            "text": "<p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>",
            "parent_id": null,
            "category_id": 3,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 2,
                "user_id": 4,
                "motion_id": 1,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 5,
            "state_extension": "Maybe",
            "state_restriction": [
              "is_submitter"
            ],
            "statute_paragraph_id": null,
            "workflow_id": 2,
            "recommendation_id": 7,
            "recommendation_extension": "if [motion:2] is acepted",
            "tags_id": [],
            "attachments_id": [
              2
            ],
            "agenda_item_id": 5,
            "list_of_speakers_id": 8,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:36:06.865840+02:00",
            "last_modified": "2020-08-11T12:42:00.402425+02:00",
            "change_recommendations_id": [
              1
            ],
            "amendments_id": [
              2
            ]
          }`),
		"motions/motion:2": []byte(`{
            "id": 2,
            "identifier": "Ä-1",
            "title": "Änderungsantrag zu Leadmotion1",
            "text": "",
            "amendment_paragraphs": [
              "<p>Lorem ipsum dolor sit amet, consectedfgddfgdf&nbsp; gdfgdfg dfg dfg ww ontes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>"
            ],
            "modified_final_version": "",
            "reason": "",
            "parent_id": 1,
            "category_id": 2,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 3,
                "user_id": 1,
                "motion_id": 2,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [],
            "attachments_id": [],
            "agenda_item_id": 6,
            "list_of_speakers_id": 9,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:39:35.025914+02:00",
            "last_modified": "2020-08-11T12:41:55.666495+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/motion:3": []byte(`{
            "id": 3,
            "identifier": "2",
            "title": "Public",
            "text": "<p>a</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>a</p>",
            "parent_id": null,
            "category_id": 1,
            "category_weight": 10000,
            "comments": [
              {
                "id": 1,
                "comment": "<p>test</p>",
                "section_id": 1,
                "read_groups_id": [
                  2,
                  3
                ]
              },
              {
                "id": 2,
                "comment": "<p>test</p>",
                "section_id": 2,
                "read_groups_id": []
              }
            ],
            "motion_block_id": 2,
            "origin": "",
            "submitters": [
              {
                "id": 4,
                "user_id": 1,
                "motion_id": 3,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [
              1
            ],
            "attachments_id": [],
            "agenda_item_id": 7,
            "list_of_speakers_id": 10,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T12:24:45.289233+02:00",
            "last_modified": "2020-08-11T12:41:51.319986+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/state:1": []byte(`{
            "id": 1,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              2,
              3,
              4
            ],
            "workflow_id": 1
          }`),
		"motions/state:10": []byte(`{
            "id": 10,
            "name": "withdrawed",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:11": []byte(`{
            "id": 11,
            "name": "adjourned",
            "recommendation_label": "Adjournment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:12": []byte(`{
            "id": 12,
            "name": "not concerned",
            "recommendation_label": "No concernment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:13": []byte(`{
            "id": 13,
            "name": "refered to committee",
            "recommendation_label": "Referral to committee",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:14": []byte(`{
            "id": 14,
            "name": "needs review",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:15": []byte(`{
            "id": 15,
            "name": "rejected (not authorized)",
            "recommendation_label": "Rejection (not authorized)",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:2": []byte(`{
            "id": 2,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:3": []byte(`{
            "id": 3,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:4": []byte(`{
            "id": 4,
            "name": "not decided",
            "recommendation_label": "No decision",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:5": []byte(`{
            "id": 5,
            "name": "in progress",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [
              "is_submitter"
            ],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": true,
            "dont_set_identifier": true,
            "show_state_extension_field": true,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              6,
              10
            ],
            "workflow_id": 2
          }`),
		"motions/state:6": []byte(`{
            "id": 6,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": true,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              7,
              10,
              15
            ],
            "workflow_id": 2
          }`),
		"motions/state:7": []byte(`{
            "id": 7,
            "name": "permitted",
            "recommendation_label": "Permission",
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": true,
            "next_states_id": [
              8,
              9,
              10,
              11,
              12,
              13,
              14
            ],
            "workflow_id": 2
          }`),
		"motions/state:8": []byte(`{
            "id": 8,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:9": []byte(`{
            "id": 9,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/statute-paragraph:1": []byte(`{
            "id": 1,
            "title": "Statute",
            "text": "<p>test</p>",
            "weight": 10000
          }`),
		"motions/workflow:1": []byte(`{
            "id": 1,
            "name": "Simple Workflow",
            "states_id": [
              1,
              2,
              3,
              4
            ],
            "first_state_id": 1
          }`),
		"motions/workflow:2": []byte(`{
            "id": 2,
            "name": "Complex Workflow",
            "states_id": [
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13,
              14,
              15
            ],
            "first_state_id": 5
          }`),
		"topics/topic:1": []byte(`{
            "id": 1,
            "title": "Topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 1,
            "list_of_speakers_id": 1
          }`),
		"topics/topic:2": []byte(`{
            "id": 2,
            "title": "Hidden",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 2,
            "list_of_speakers_id": 2
          }`),
		"topics/topic:3": []byte(`{
            "id": 3,
            "title": "Internal",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 3,
            "list_of_speakers_id": 3
          }`),
		"topics/topic:4": []byte(`{
            "id": 4,
            "title": "Another public topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 9,
            "list_of_speakers_id": 14
          }`),
		"users/group:1": []byte(`{
            "id": 1,
            "name": "Default",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_see",
              "users.can_change_password"
            ]
          }`),
		"users/group:2": []byte(`{
            "id": 2,
            "name": "Admin",
            "permissions": []
          }`),
		"users/group:3": []byte(`{
            "id": 3,
            "name": "Delegates",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password"
            ]
          }`),
		"users/group:4": []byte(`{
            "id": 4,
            "name": "Staff",
            "permissions": [
              "agenda.can_manage",
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_manage_list_of_speakers",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_manage",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_history",
              "core.can_manage_projector",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "core.can_manage_tags",
              "mediafiles.can_manage",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_manage",
              "motions.can_manage_metadata",
              "motions.can_see",
              "motions.can_see_internal",
              "users.can_change_password",
              "users.can_manage",
              "users.can_see_extra_data",
              "users.can_see_name"
            ]
          }`),
		"users/group:5": []byte(`{
            "id": 5,
            "name": "Committees",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password",
              "users.can_see_name"
            ]
          }`),
		"users/personal-note:1": []byte(`{
            "id": 1,
            "user_id": 1,
            "notes": {
              "motions/motion": {
                "3": {
                  "note": "<p>test</p>",
                  "star": false
                }
              }
            }
          }`),
		"users/user:1": []byte(`{
            "id": 1,
            "vote_weight": "1.000000",
            "default_password": "admin",
            "about_me": "",
            "structure_level": "",
            "is_present": false,
            "email": "",
            "groups_id": [
              2
            ],
            "is_active": true,
            "number": "",
            "last_name": "Administrator",
            "is_committee": false,
            "title": "",
            "comment": "",
            "auth_type": "default",
            "last_email_send": null,
            "first_name": "",
            "username": "admin",
            "gender": ""
          }`),
		"users/user:2": []byte(`{
            "id": 2,
            "vote_weight": "1.000000",
            "default_password": "8NpbvXCBDr",
            "about_me": "",
            "structure_level": "",
            "is_present": false,
            "email": "",
            "groups_id": [],
            "is_active": true,
            "number": "",
            "last_name": "",
            "is_committee": false,
            "title": "",
            "comment": "",
            "auth_type": "default",
            "last_email_send": null,
            "first_name": "candidate1",
            "username": "candidate1",
            "gender": ""
          }`),
		"users/user:3": []byte(`{
            "id": 3,
            "vote_weight": "1.000000",
            "default_password": "5YLEHrUUTG",
            "about_me": "",
            "structure_level": "",
            "is_present": false,
            "email": "",
            "groups_id": [],
            "is_active": true,
            "number": "",
            "last_name": "",
            "is_committee": false,
            "title": "",
            "comment": "",
            "auth_type": "default",
            "last_email_send": null,
            "first_name": "candidate2",
            "username": "candidate2",
            "gender": ""
          }`),
		"users/user:4": []byte(`{
            "id": 4,
            "vote_weight": "1.000000",
            "default_password": "a",
            "about_me": "",
            "structure_level": "",
            "is_present": true,
            "email": "",
            "groups_id": [
              3
            ],
            "is_active": true,
            "number": "",
            "last_name": "",
            "is_committee": false,
            "title": "",
            "comment": "",
            "auth_type": "default",
            "last_email_send": null,
            "first_name": "a",
            "username": "a",
            "gender": ""
          }`),
		"users/user:5": []byte(`{
            "id": 5,
            "vote_weight": "1.000000",
            "default_password": "b",
            "about_me": "",
            "structure_level": "",
            "is_present": true,
            "email": "",
            "groups_id": [
              3
            ],
            "is_active": true,
            "number": "",
            "last_name": "",
            "is_committee": false,
            "title": "",
            "comment": "",
            "auth_type": "default",
            "last_email_send": null,
            "first_name": "b",
            "username": "b",
            "gender": ""
          }`),
		"users/user:6": []byte(`{
            "id": 6,
            "vote_weight": "1.000000",
            "default_password": "ZdbyxFDWpp",
            "about_me": "",
            "structure_level": "layer X",
            "is_present": true,
            "email": "",
            "groups_id": [],
            "is_active": true,
            "number": "3",
            "last_name": "the last name",
            "is_committee": false,
            "title": "title",
            "comment": "",
            "auth_type": "default",
            "last_email_send": null,
            "first_name": "speaker1",
            "username": "speaker1",
            "gender": ""
          }`),
		"users/user:7": []byte(`{
            "id": 7,
            "vote_weight": "1.000000",
            "default_password": "5HZr2zPM3x",
            "about_me": "",
            "structure_level": "",
            "is_present": true,
            "email": "",
            "groups_id": [],
            "is_active": true,
            "number": "",
            "last_name": "",
            "is_committee": false,
            "title": "",
            "comment": "",
            "auth_type": "default",
            "last_email_send": null,
            "first_name": "speaker2",
            "username": "speaker2",
            "gender": ""
          }`),
		"users/user:8": []byte(`{
            "id": 8,
            "vote_weight": "1.000000",
            "default_password": "hq59FcgwLc",
            "about_me": "",
            "structure_level": "",
            "is_present": true,
            "email": "",
            "groups_id": [],
            "is_active": true,
            "number": "",
            "last_name": "",
            "is_committee": false,
            "title": "",
            "comment": "",
            "auth_type": "default",
            "last_email_send": null,
            "first_name": "not required",
            "username": "not required",
            "gender": ""
          }`),
	},
	2: {
		"agenda/item:1": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            },
            "is_internal": false,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": null,
            "id": 1,
            "level": 0,
            "title_information": {
              "title": "Topic"
            },
            "closed": false,
            "weight": 2,
            "duration": null
          }`),
		"agenda/item:3": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": null,
            "id": 3,
            "level": 0,
            "title_information": {
              "title": "Internal"
            },
            "closed": false,
            "weight": 8,
            "duration": null
          }`),
		"agenda/item:5": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 5,
            "level": 1,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "closed": false,
            "weight": 14,
            "duration": null
          }`),
		"agenda/item:6": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 6,
            "level": 1,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "closed": false,
            "weight": 16,
            "duration": 0
          }`),
		"agenda/item:7": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": 6,
            "id": 7,
            "level": 2,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "closed": false,
            "weight": 18,
            "duration": null
          }`),
		"agenda/list-of-speakers:1": []byte(`{
            "id": 1,
            "title_information": {
              "title": "Topic"
            },
            "speakers": [
              {
                "id": 3,
                "user_id": 6,
                "begin_time": "2020-08-11T12:28:30.894574+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:10": []byte(`{
            "id": 10,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:11": []byte(`{
            "id": 11,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:12": []byte(`{
            "id": 12,
            "title_information": {
              "title": "block"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:13": []byte(`{
            "id": 13,
            "title_information": {
              "title": "block internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:14": []byte(`{
            "id": 14,
            "title_information": {
              "title": "Another public topic"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:2": []byte(`{
            "id": 2,
            "title_information": {
              "title": "Hidden"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:3": []byte(`{
            "id": 3,
            "title_information": {
              "title": "Internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:4": []byte(`{
            "id": 4,
            "title_information": {
              "title": "folder"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:5": []byte(`{
            "id": 5,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:6": []byte(`{
            "id": 6,
            "title_information": {
              "title": "in.jpg"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:7": []byte(`{
            "id": 7,
            "title_information": {
              "title": "Assignment"
            },
            "speakers": [
              {
                "id": 4,
                "user_id": 6,
                "begin_time": "2020-08-11T12:29:55.054553+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              },
              {
                "id": 6,
                "user_id": 7,
                "begin_time": null,
                "end_time": null,
                "weight": 2,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "assignments/assignment",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:8": []byte(`{
            "id": 8,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:9": []byte(`{
            "id": 9,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            }
          }`),
		"assignments/assignment-option:1": []byte(`{
            "user_id": 2,
            "weight": 1,
            "id": 1,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-option:2": []byte(`{
            "user_id": 3,
            "weight": 3,
            "id": 2,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-poll:1": []byte(`{
            "assignment_id": 1,
            "description": "",
            "pollmethod": "votes",
            "votes_amount": 1,
            "allow_multiple_votes_per_candidate": false,
            "global_no": true,
            "global_abstain": true,
            "state": 2,
            "type": "named",
            "title": "Wahlgang",
            "groups_id": [
              3
            ],
            "options_id": [
              1,
              2
            ],
            "id": 1,
            "onehundred_percent_base": "valid",
            "majority_method": "simple",
            "user_has_voted": false
          }`),
		"assignments/assignment:1": []byte(`{
            "id": 1,
            "title": "Assignment",
            "description": "",
            "open_posts": 1,
            "phase": 0,
            "assignment_related_users": [
              {
                "id": 1,
                "user_id": 2,
                "weight": 1
              },
              {
                "id": 3,
                "user_id": 3,
                "weight": 3
              }
            ],
            "default_poll_description": "",
            "agenda_item_id": 4,
            "list_of_speakers_id": 7,
            "tags_id": [
              2
            ],
            "attachments_id": [],
            "number_poll_candidates": false,
            "polls_id": [
              1
            ]
          }`),
		"core/config:10": []byte(`{
            "id": 10,
            "key": "general_system_conference_show",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show live conference window",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 140,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:100": []byte(`{
            "id": 100,
            "key": "motion_poll_default_100_percent_base",
            "value": "YNA",
            "data": {
              "defaultValue": "YNA",
              "inputType": "choice",
              "label": "Default 100 % base of a voting result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 370,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:101": []byte(`{
            "id": 101,
            "key": "motion_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:102": []byte(`{
            "id": 102,
            "key": "motion_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 372,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:103": []byte(`{
            "id": 103,
            "key": "motions_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 373,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:104": []byte(`{
            "id": 104,
            "key": "motions_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 374,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:105": []byte(`{
            "id": 105,
            "key": "motions_export_title",
            "value": "Motions",
            "data": {
              "defaultValue": "Motions",
              "inputType": "string",
              "label": "Title for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 380,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:106": []byte(`{
            "id": 106,
            "key": "motions_export_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 382,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:107": []byte(`{
            "id": 107,
            "key": "motions_export_submitter_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show submitters and recommendation/state in table of contents",
              "helpText": "",
              "choices": null,
              "weight": 384,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:108": []byte(`{
            "id": 108,
            "key": "motions_export_follow_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show checkbox to record decision",
              "helpText": "",
              "choices": null,
              "weight": 386,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:109": []byte(`{
            "id": 109,
            "key": "assignment_poll_method",
            "value": "votes",
            "data": {
              "defaultValue": "votes",
              "inputType": "choice",
              "label": "Default election method",
              "helpText": "",
              "choices": [
                {
                  "value": "votes",
                  "display_name": "Yes per candidate"
                },
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                }
              ],
              "weight": 400,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:11": []byte(`{
            "id": 11,
            "key": "general_system_conference_auto_connect",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Connect all users to live conference automatically",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 141,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:110": []byte(`{
            "id": 110,
            "key": "assignment_poll_default_type",
            "value": "analog",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 403,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:111": []byte(`{
            "id": 111,
            "key": "assignment_poll_default_100_percent_base",
            "value": "valid",
            "data": {
              "defaultValue": "valid",
              "inputType": "choice",
              "label": "Default 100 % base of an election result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                },
                {
                  "value": "votes",
                  "display_name": "Sum of votes including general No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 405,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:112": []byte(`{
            "id": 112,
            "key": "assignment_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 410,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:113": []byte(`{
            "id": 113,
            "key": "assignment_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:114": []byte(`{
            "id": 114,
            "key": "assignment_poll_sort_poll_result_by_votes",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Sort election results by amount of votes",
              "helpText": "",
              "choices": null,
              "weight": 420,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:115": []byte(`{
            "id": 115,
            "key": "assignment_poll_add_candidates_to_list_of_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Put all candidates on the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 425,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:116": []byte(`{
            "id": 116,
            "key": "assignments_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 430,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:117": []byte(`{
            "id": 117,
            "key": "assignments_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 435,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:118": []byte(`{
            "id": 118,
            "key": "assignments_pdf_title",
            "value": "Elections",
            "data": {
              "defaultValue": "Elections",
              "inputType": "string",
              "label": "Title for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 460,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:119": []byte(`{
            "id": 119,
            "key": "assignments_pdf_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 470,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:12": []byte(`{
            "id": 12,
            "key": "general_system_conference_los_restriction",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow only current speakers and list of speakers managers to enter the live conference",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 142,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:13": []byte(`{
            "id": 13,
            "key": "general_system_stream_url",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Livestream url",
              "helpText": "Remove URL to deactivate livestream. Check extra group permission to see livestream.",
              "choices": null,
              "weight": 143,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:14": []byte(`{
            "id": 14,
            "key": "general_system_enable_anonymous",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow access for anonymous guest users",
              "helpText": "",
              "choices": null,
              "weight": 150,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:15": []byte(`{
            "id": 15,
            "key": "general_login_info_text",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Show this text on the login page",
              "helpText": "",
              "choices": null,
              "weight": 152,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:16": []byte(`{
            "id": 16,
            "key": "openslides_theme",
            "value": "openslides-default-light-theme",
            "data": {
              "defaultValue": "openslides-default-light-theme",
              "inputType": "choice",
              "label": "OpenSlides Theme",
              "helpText": "",
              "choices": [
                {
                  "value": "openslides-default-light-theme",
                  "display_name": "OpenSlides Default"
                },
                {
                  "value": "openslides-default-dark-theme",
                  "display_name": "OpenSlides Dark"
                },
                {
                  "value": "openslides-red-light-theme",
                  "display_name": "OpenSlides Red"
                },
                {
                  "value": "openslides-red-dark-theme",
                  "display_name": "OpenSlides Red Dark"
                },
                {
                  "value": "openslides-green-light-theme",
                  "display_name": "OpenSlides Green"
                },
                {
                  "value": "openslides-green-dark-theme",
                  "display_name": "OpenSlides Green Dark"
                },
                {
                  "value": "openslides-solarized-dark-theme",
                  "display_name": "OpenSlides Solarized"
                }
              ],
              "weight": 154,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:17": []byte(`{
            "id": 17,
            "key": "general_csv_separator",
            "value": ",",
            "data": {
              "defaultValue": ",",
              "inputType": "string",
              "label": "Separator used for all csv exports and examples",
              "helpText": "",
              "choices": null,
              "weight": 160,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:18": []byte(`{
            "id": 18,
            "key": "general_csv_encoding",
            "value": "utf-8",
            "data": {
              "defaultValue": "utf-8",
              "inputType": "choice",
              "label": "Default encoding for all csv exports",
              "helpText": "",
              "choices": [
                {
                  "value": "utf-8",
                  "display_name": "UTF-8"
                },
                {
                  "value": "iso-8859-15",
                  "display_name": "ISO-8859-15"
                }
              ],
              "weight": 162,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:19": []byte(`{
            "id": 19,
            "key": "general_export_pdf_pagenumber_alignment",
            "value": "center",
            "data": {
              "defaultValue": "center",
              "inputType": "choice",
              "label": "Page number alignment in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "left",
                  "display_name": "Left"
                },
                {
                  "value": "center",
                  "display_name": "Center"
                },
                {
                  "value": "right",
                  "display_name": "Right"
                }
              ],
              "weight": 164,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:2": []byte(`{
            "id": 2,
            "key": "general_event_name",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Event name",
              "helpText": "",
              "choices": null,
              "weight": 110,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:20": []byte(`{
            "id": 20,
            "key": "general_export_pdf_fontsize",
            "value": "10",
            "data": {
              "defaultValue": "10",
              "inputType": "choice",
              "label": "Standard font size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "10",
                  "display_name": "10"
                },
                {
                  "value": "11",
                  "display_name": "11"
                },
                {
                  "value": "12",
                  "display_name": "12"
                }
              ],
              "weight": 166,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:21": []byte(`{
            "id": 21,
            "key": "general_export_pdf_pagesize",
            "value": "A4",
            "data": {
              "defaultValue": "A4",
              "inputType": "choice",
              "label": "Standard page size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "A4",
                  "display_name": "DIN A4"
                },
                {
                  "value": "A5",
                  "display_name": "DIN A5"
                }
              ],
              "weight": 168,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:22": []byte(`{
            "id": 22,
            "key": "logos_available",
            "value": [
              "logo_projector_main",
              "logo_projector_header",
              "logo_web_header",
              "logo_pdf_header_L",
              "logo_pdf_header_R",
              "logo_pdf_footer_L",
              "logo_pdf_footer_R",
              "logo_pdf_ballot_paper"
            ],
            "data": null
          }`),
		"core/config:23": []byte(`{
            "id": 23,
            "key": "logo_projector_main",
            "value": {
              "display_name": "Projector logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:24": []byte(`{
            "id": 24,
            "key": "logo_projector_header",
            "value": {
              "display_name": "Projector header image",
              "path": ""
            },
            "data": null
          }`),
		"core/config:25": []byte(`{
            "id": 25,
            "key": "logo_web_header",
            "value": {
              "display_name": "Web interface header logo",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:26": []byte(`{
            "id": 26,
            "key": "logo_pdf_header_L",
            "value": {
              "display_name": "PDF header logo (left)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:27": []byte(`{
            "id": 27,
            "key": "logo_pdf_header_R",
            "value": {
              "display_name": "PDF header logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:28": []byte(`{
            "id": 28,
            "key": "logo_pdf_footer_L",
            "value": {
              "display_name": "PDF footer logo (left)",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:29": []byte(`{
            "id": 29,
            "key": "logo_pdf_footer_R",
            "value": {
              "display_name": "PDF footer logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:3": []byte(`{
            "id": 3,
            "key": "general_event_description",
            "value": "Presentation and assembly system",
            "data": {
              "defaultValue": "Presentation and assembly system",
              "inputType": "string",
              "label": "Short description of event",
              "helpText": "",
              "choices": null,
              "weight": 115,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:30": []byte(`{
            "id": 30,
            "key": "logo_pdf_ballot_paper",
            "value": {
              "display_name": "PDF ballot paper logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:31": []byte(`{
            "id": 31,
            "key": "fonts_available",
            "value": [
              "font_regular",
              "font_italic",
              "font_bold",
              "font_bold_italic",
              "font_monospace"
            ],
            "data": null
          }`),
		"core/config:32": []byte(`{
            "id": 32,
            "key": "font_regular",
            "value": {
              "display_name": "Font regular",
              "default": "assets/fonts/fira-sans-latin-400.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:33": []byte(`{
            "id": 33,
            "key": "font_italic",
            "value": {
              "display_name": "Font italic",
              "default": "assets/fonts/fira-sans-latin-400italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:34": []byte(`{
            "id": 34,
            "key": "font_bold",
            "value": {
              "display_name": "Font bold",
              "default": "assets/fonts/fira-sans-latin-500.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:35": []byte(`{
            "id": 35,
            "key": "font_bold_italic",
            "value": {
              "display_name": "Font bold italic",
              "default": "assets/fonts/fira-sans-latin-500italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:36": []byte(`{
            "id": 36,
            "key": "font_monospace",
            "value": {
              "display_name": "Font monospace",
              "default": "assets/fonts/roboto-condensed-bold.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:37": []byte(`{
            "id": 37,
            "key": "translations",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "translations",
              "label": "Custom translations",
              "helpText": "",
              "choices": null,
              "weight": 1000,
              "group": "Custom translations",
              "subgroup": "General"
            }
          }`),
		"core/config:38": []byte(`{
            "id": 38,
            "key": "config_version",
            "value": 2,
            "data": null
          }`),
		"core/config:39": []byte(`{
            "id": 39,
            "key": "db_id",
            "value": "2c3bd736c87a48a4be1f0dc707702144",
            "data": null
          }`),
		"core/config:4": []byte(`{
            "id": 4,
            "key": "general_event_date",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event date",
              "helpText": "",
              "choices": null,
              "weight": 120,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:40": []byte(`{
            "id": 40,
            "key": "users_sort_by",
            "value": "first_name",
            "data": {
              "defaultValue": "first_name",
              "inputType": "choice",
              "label": "Sort name of participants by",
              "helpText": "",
              "choices": [
                {
                  "value": "first_name",
                  "display_name": "Given name"
                },
                {
                  "value": "last_name",
                  "display_name": "Surname"
                },
                {
                  "value": "number",
                  "display_name": "Participant number"
                }
              ],
              "weight": 510,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:41": []byte(`{
            "id": 41,
            "key": "users_enable_presence_view",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Enable participant presence view",
              "helpText": "",
              "choices": null,
              "weight": 511,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:42": []byte(`{
            "id": 42,
            "key": "users_allow_self_set_present",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow users to set themselves as present",
              "helpText": "e.g. for online meetings",
              "choices": null,
              "weight": 512,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:43": []byte(`{
            "id": 43,
            "key": "users_activate_vote_weight",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate vote weight",
              "helpText": "",
              "choices": null,
              "weight": 513,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:44": []byte(`{
            "id": 44,
            "key": "users_pdf_welcometitle",
            "value": "Welcome to OpenSlides",
            "data": {
              "defaultValue": "Welcome to OpenSlides",
              "inputType": "string",
              "label": "Title for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 520,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:45": []byte(`{
            "id": 45,
            "key": "users_pdf_welcometext",
            "value": "[Place for your welcome and help text.]",
            "data": {
              "defaultValue": "[Place for your welcome and help text.]",
              "inputType": "string",
              "label": "Help text for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 530,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:46": []byte(`{
            "id": 46,
            "key": "users_pdf_url",
            "value": "http://example.com:8000",
            "data": {
              "defaultValue": "http://example.com:8000",
              "inputType": "string",
              "label": "System URL",
              "helpText": "Used for QRCode in PDF of access data.",
              "choices": null,
              "weight": 540,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:47": []byte(`{
            "id": 47,
            "key": "users_pdf_wlan_ssid",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN name (SSID)",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 550,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:48": []byte(`{
            "id": 48,
            "key": "users_pdf_wlan_password",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN password",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 560,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:49": []byte(`{
            "id": 49,
            "key": "users_pdf_wlan_encryption",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "choice",
              "label": "WLAN encryption",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": [
                {
                  "value": "",
                  "display_name": "---------"
                },
                {
                  "value": "WEP",
                  "display_name": "WEP"
                },
                {
                  "value": "WPA",
                  "display_name": "WPA/WPA2"
                },
                {
                  "value": "nopass",
                  "display_name": "No encryption"
                }
              ],
              "weight": 570,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:5": []byte(`{
            "id": 5,
            "key": "general_event_location",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event location",
              "helpText": "",
              "choices": null,
              "weight": 125,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:50": []byte(`{
            "id": 50,
            "key": "users_email_sender",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Sender name",
              "helpText": "The sender address is defined in the OpenSlides server settings and should modified by administrator only.",
              "choices": null,
              "weight": 600,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:51": []byte(`{
            "id": 51,
            "key": "users_email_replyto",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Reply address",
              "helpText": "",
              "choices": null,
              "weight": 601,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:52": []byte(`{
            "id": 52,
            "key": "users_email_subject",
            "value": "OpenSlides access data",
            "data": {
              "defaultValue": "OpenSlides access data",
              "inputType": "string",
              "label": "Email subject",
              "helpText": "You can use {event_name} and {username} as placeholder.",
              "choices": null,
              "weight": 605,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:53": []byte(`{
            "id": 53,
            "key": "users_email_body",
            "value": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
            "data": {
              "defaultValue": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
              "inputType": "text",
              "label": "Email body",
              "helpText": "Use these placeholders: {name}, {event_name}, {url}, {username}, {password}. The url referrs to the system url.",
              "choices": null,
              "weight": 610,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:54": []byte(`{
            "id": 54,
            "key": "agenda_start_event_date_time",
            "value": null,
            "data": {
              "defaultValue": null,
              "inputType": "datetimepicker",
              "label": "Begin of event",
              "helpText": "Input format: DD.MM.YYYY HH:MM",
              "choices": null,
              "weight": 200,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:55": []byte(`{
            "id": 55,
            "key": "agenda_show_subtitle",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show subtitles in the agenda",
              "helpText": "",
              "choices": null,
              "weight": 201,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:56": []byte(`{
            "id": 56,
            "key": "agenda_enable_numbering",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Enable numbering for agenda items",
              "helpText": "",
              "choices": null,
              "weight": 205,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:57": []byte(`{
            "id": 57,
            "key": "agenda_number_prefix",
            "value": "TOP",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Numbering prefix for agenda items",
              "helpText": "This prefix will be set if you run the automatic agenda numbering.",
              "choices": null,
              "weight": 206,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:58": []byte(`{
            "id": 58,
            "key": "agenda_numeral_system",
            "value": "arabic",
            "data": {
              "defaultValue": "arabic",
              "inputType": "choice",
              "label": "Numeral system for agenda items",
              "helpText": "",
              "choices": [
                {
                  "value": "arabic",
                  "display_name": "Arabic"
                },
                {
                  "value": "roman",
                  "display_name": "Roman"
                }
              ],
              "weight": 207,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:59": []byte(`{
            "id": 59,
            "key": "agenda_item_creation",
            "value": "default_yes",
            "data": {
              "defaultValue": "default_yes",
              "inputType": "choice",
              "label": "Add to agenda",
              "helpText": "",
              "choices": [
                {
                  "value": "always",
                  "display_name": "Always"
                },
                {
                  "value": "never",
                  "display_name": "Never"
                },
                {
                  "value": "default_yes",
                  "display_name": "Ask, default yes"
                },
                {
                  "value": "default_no",
                  "display_name": "Ask, default no"
                }
              ],
              "weight": 210,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:6": []byte(`{
            "id": 6,
            "key": "general_event_legal_notice",
            "value": "<a href=\"http://www.openslides.org\">OpenSlides</a> is a free web based presentation and assembly system for visualizing and controlling agenda, motions and elections of an assembly.",
            "data": null
          }`),
		"core/config:60": []byte(`{
            "id": 60,
            "key": "agenda_new_items_default_visibility",
            "value": "2",
            "data": {
              "defaultValue": "2",
              "inputType": "choice",
              "label": "Default visibility for new agenda items (except topics)",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Public item"
                },
                {
                  "value": "2",
                  "display_name": "Internal item"
                },
                {
                  "value": "3",
                  "display_name": "Hidden item"
                }
              ],
              "weight": 211,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:61": []byte(`{
            "id": 61,
            "key": "agenda_hide_internal_items_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Hide internal items when projecting subitems",
              "helpText": "",
              "choices": null,
              "weight": 212,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:62": []byte(`{
            "id": 62,
            "key": "agenda_show_last_speakers",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of last speakers to be shown on the projector",
              "helpText": "",
              "choices": null,
              "weight": 220,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:63": []byte(`{
            "id": 63,
            "key": "agenda_show_next_speakers",
            "value": -1,
            "data": {
              "defaultValue": -1,
              "inputType": "integer",
              "label": "Number of the next speakers to be shown on the projector",
              "helpText": "Enter number of the next shown speakers. Choose -1 to show all next speakers.",
              "choices": null,
              "weight": 222,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:64": []byte(`{
            "id": 64,
            "key": "agenda_countdown_warning_time",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Show orange countdown in the last x seconds of speaking time",
              "helpText": "Enter duration in seconds. Choose 0 to disable warning color.",
              "choices": null,
              "weight": 224,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:65": []byte(`{
            "id": 65,
            "key": "projector_default_countdown",
            "value": 60,
            "data": {
              "defaultValue": 60,
              "inputType": "integer",
              "label": "Predefined seconds of new countdowns",
              "helpText": "",
              "choices": null,
              "weight": 226,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:66": []byte(`{
            "id": 66,
            "key": "agenda_couple_countdown_and_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Couple countdown with the list of speakers",
              "helpText": "[Begin speech] starts the countdown, [End speech] stops the countdown.",
              "choices": null,
              "weight": 228,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:67": []byte(`{
            "id": 67,
            "key": "agenda_hide_amount_of_speakers",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide the amount of speakers in subtitle of list of speakers slide",
              "helpText": "",
              "choices": null,
              "weight": 230,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:68": []byte(`{
            "id": 68,
            "key": "agenda_present_speakers_only",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Only present participants can be added to the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 232,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:69": []byte(`{
            "id": 69,
            "key": "agenda_show_first_contribution",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show hint »first speech« in the list of speakers management view",
              "helpText": "",
              "choices": null,
              "weight": 234,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:7": []byte(`{
            "id": 7,
            "key": "general_event_privacy_policy",
            "value": "",
            "data": null
          }`),
		"core/config:70": []byte(`{
            "id": 70,
            "key": "motions_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new motions",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 310,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:71": []byte(`{
            "id": 71,
            "key": "motions_statute_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new statute amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 312,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:72": []byte(`{
            "id": 72,
            "key": "motions_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 314,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:73": []byte(`{
            "id": 73,
            "key": "motions_preamble",
            "value": "The assembly may decide:",
            "data": {
              "defaultValue": "The assembly may decide:",
              "inputType": "string",
              "label": "Motion preamble",
              "helpText": "",
              "choices": null,
              "weight": 320,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:74": []byte(`{
            "id": 74,
            "key": "motions_default_line_numbering",
            "value": "outside",
            "data": {
              "defaultValue": "outside",
              "inputType": "choice",
              "label": "Default line numbering",
              "helpText": "",
              "choices": [
                {
                  "value": "outside",
                  "display_name": "outside"
                },
                {
                  "value": "inline",
                  "display_name": "inline"
                },
                {
                  "value": "none",
                  "display_name": "Disabled"
                }
              ],
              "weight": 322,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:75": []byte(`{
            "id": 75,
            "key": "motions_line_length",
            "value": 85,
            "data": {
              "defaultValue": 85,
              "inputType": "integer",
              "label": "Line length",
              "helpText": "The maximum number of characters per line. Relevant when line numbering is enabled. Min: 40",
              "choices": null,
              "weight": 323,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:76": []byte(`{
            "id": 76,
            "key": "motions_reason_required",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Reason required for creating new motion",
              "helpText": "",
              "choices": null,
              "weight": 324,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:77": []byte(`{
            "id": 77,
            "key": "motions_disable_text_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide motion text on projector",
              "helpText": "",
              "choices": null,
              "weight": 325,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:78": []byte(`{
            "id": 78,
            "key": "motions_disable_reason_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide reason on projector",
              "helpText": "",
              "choices": null,
              "weight": 326,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:79": []byte(`{
            "id": 79,
            "key": "motions_disable_recommendation_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide recommendation on projector",
              "helpText": "",
              "choices": null,
              "weight": 327,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:8": []byte(`{
            "id": 8,
            "key": "general_event_welcome_title",
            "value": "Welcome to OpenSlides",
            "data": null
          }`),
		"core/config:80": []byte(`{
            "id": 80,
            "key": "motions_hide_referring_motions",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide referring motions",
              "helpText": "",
              "choices": null,
              "weight": 328,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:81": []byte(`{
            "id": 81,
            "key": "motions_disable_sidebox_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show meta information box below the title on projector",
              "helpText": "",
              "choices": null,
              "weight": 329,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:82": []byte(`{
            "id": 82,
            "key": "motions_show_sequential_numbers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show the sequential number for a motion",
              "helpText": "In motion list, motion detail and PDF.",
              "choices": null,
              "weight": 330,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:83": []byte(`{
            "id": 83,
            "key": "motions_recommendations_by",
            "value": "ABK",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender",
              "helpText": "Will be displayed as label before selected recommendation. Use an empty value to disable the recommendation system.",
              "choices": null,
              "weight": 332,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:84": []byte(`{
            "id": 84,
            "key": "motions_statute_recommendations_by",
            "value": "ABK2",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender for statute amendments",
              "helpText": "Will be displayed as label before selected recommendation in statute amendments.",
              "choices": null,
              "weight": 333,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:85": []byte(`{
            "id": 85,
            "key": "motions_recommendation_text_mode",
            "value": "diff",
            "data": {
              "defaultValue": "diff",
              "inputType": "choice",
              "label": "Default text version for change recommendations",
              "helpText": "",
              "choices": [
                {
                  "value": "original",
                  "display_name": "Original version"
                },
                {
                  "value": "changed",
                  "display_name": "Changed version"
                },
                {
                  "value": "diff",
                  "display_name": "Diff version"
                },
                {
                  "value": "agreed",
                  "display_name": "Final version"
                }
              ],
              "weight": 334,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:86": []byte(`{
            "id": 86,
            "key": "motions_motions_sorting",
            "value": "identifier",
            "data": {
              "defaultValue": "identifier",
              "inputType": "choice",
              "label": "Sort motions by",
              "helpText": "",
              "choices": [
                {
                  "value": "weight",
                  "display_name": "Call list"
                },
                {
                  "value": "identifier",
                  "display_name": "Identifier"
                }
              ],
              "weight": 335,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:87": []byte(`{
            "id": 87,
            "key": "motions_identifier",
            "value": "per_category",
            "data": {
              "defaultValue": "per_category",
              "inputType": "choice",
              "label": "Identifier",
              "helpText": "",
              "choices": [
                {
                  "value": "per_category",
                  "display_name": "Numbered per category"
                },
                {
                  "value": "serially_numbered",
                  "display_name": "Serially numbered"
                },
                {
                  "value": "manually",
                  "display_name": "Set it manually"
                }
              ],
              "weight": 340,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:88": []byte(`{
            "id": 88,
            "key": "motions_identifier_min_digits",
            "value": 1,
            "data": {
              "defaultValue": 1,
              "inputType": "integer",
              "label": "Number of minimal digits for identifier",
              "helpText": "Uses leading zeros to sort motions correctly by identifier.",
              "choices": null,
              "weight": 342,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:89": []byte(`{
            "id": 89,
            "key": "motions_identifier_with_blank",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow blank in identifier",
              "helpText": "Blank between prefix and number, e.g. 'A 001'.",
              "choices": null,
              "weight": 344,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:9": []byte(`{
            "id": 9,
            "key": "general_event_welcome_text",
            "value": "[Space for your welcome text.]",
            "data": null
          }`),
		"core/config:90": []byte(`{
            "id": 90,
            "key": "motions_statutes_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate statute amendments",
              "helpText": "",
              "choices": null,
              "weight": 350,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:91": []byte(`{
            "id": 91,
            "key": "motions_amendments_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate amendments",
              "helpText": "",
              "choices": null,
              "weight": 351,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:92": []byte(`{
            "id": 92,
            "key": "motions_amendments_main_table",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show amendments together with motions",
              "helpText": "",
              "choices": null,
              "weight": 352,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:93": []byte(`{
            "id": 93,
            "key": "motions_amendments_prefix",
            "value": "Ä-",
            "data": {
              "defaultValue": "-",
              "inputType": "string",
              "label": "Prefix for the identifier for amendments",
              "helpText": "",
              "choices": null,
              "weight": 353,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:94": []byte(`{
            "id": 94,
            "key": "motions_amendments_text_mode",
            "value": "paragraph",
            "data": {
              "defaultValue": "paragraph",
              "inputType": "choice",
              "label": "How to create new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "freestyle",
                  "display_name": "Empty text field"
                },
                {
                  "value": "fulltext",
                  "display_name": "Edit the whole motion text"
                },
                {
                  "value": "paragraph",
                  "display_name": "Paragraph-based, Diff-enabled"
                }
              ],
              "weight": 354,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:95": []byte(`{
            "id": 95,
            "key": "motions_amendments_multiple_paragraphs",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Amendments can change multiple paragraphs",
              "helpText": "",
              "choices": null,
              "weight": 355,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:96": []byte(`{
            "id": 96,
            "key": "motions_amendments_of_amendments",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow amendments of amendments",
              "helpText": "",
              "choices": null,
              "weight": 356,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:97": []byte(`{
            "id": 97,
            "key": "motions_min_supporters",
            "value": 1,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of (minimum) required supporters for a motion",
              "helpText": "Choose 0 to disable the supporting system.",
              "choices": null,
              "weight": 360,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:98": []byte(`{
            "id": 98,
            "key": "motions_remove_supporters",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Remove all supporters of a motion if a submitter edits his motion in early state",
              "helpText": "",
              "choices": null,
              "weight": 361,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:99": []byte(`{
            "id": 99,
            "key": "motion_poll_default_type",
            "value": "named",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 367,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/countdown:1": []byte(`{
            "id": 1,
            "title": "Default countdown",
            "description": "",
            "default_time": 60,
            "countdown_time": 1597141855.090748,
            "running": true
          }`),
		"core/projection-default:1": []byte(`{
            "id": 1,
            "name": "agenda_all_items",
            "display_name": "Agenda",
            "projector_id": 1
          }`),
		"core/projection-default:10": []byte(`{
            "id": 10,
            "name": "messages",
            "display_name": "Messages",
            "projector_id": 1
          }`),
		"core/projection-default:11": []byte(`{
            "id": 11,
            "name": "countdowns",
            "display_name": "Countdowns",
            "projector_id": 1
          }`),
		"core/projection-default:12": []byte(`{
            "id": 12,
            "name": "assignment_poll",
            "display_name": "Ballots",
            "projector_id": 1
          }`),
		"core/projection-default:13": []byte(`{
            "id": 13,
            "name": "motion_poll",
            "display_name": "Motion votes",
            "projector_id": 1
          }`),
		"core/projection-default:2": []byte(`{
            "id": 2,
            "name": "topics",
            "display_name": "Topics",
            "projector_id": 1
          }`),
		"core/projection-default:3": []byte(`{
            "id": 3,
            "name": "agenda_list_of_speakers",
            "display_name": "List of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:4": []byte(`{
            "id": 4,
            "name": "agenda_current_list_of_speakers",
            "display_name": "Current list of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:5": []byte(`{
            "id": 5,
            "name": "motions",
            "display_name": "Motions",
            "projector_id": 1
          }`),
		"core/projection-default:6": []byte(`{
            "id": 6,
            "name": "motionBlocks",
            "display_name": "Motion blocks",
            "projector_id": 1
          }`),
		"core/projection-default:7": []byte(`{
            "id": 7,
            "name": "assignments",
            "display_name": "Elections",
            "projector_id": 1
          }`),
		"core/projection-default:8": []byte(`{
            "id": 8,
            "name": "users",
            "display_name": "Participants",
            "projector_id": 1
          }`),
		"core/projection-default:9": []byte(`{
            "id": 9,
            "name": "mediafiles",
            "display_name": "Files",
            "projector_id": 1
          }`),
		"core/projector-message:1": []byte(`{
            "id": 1,
            "message": "<p>test</p>"
          }`),
		"core/projector:1": []byte(`{
            "id": 1,
            "elements": [
              {
                "name": "mediafiles/mediafile",
                "id": 3
              }
            ],
            "elements_preview": [],
            "elements_history": [
              [
                {
                  "name": "assignments/assignment",
                  "id": 1
                }
              ]
            ],
            "scale": 0,
            "scroll": 0,
            "name": "Default projector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [
              1,
              2,
              3,
              4,
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13
            ],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/projector:2": []byte(`{
            "id": 2,
            "elements": [
              {
                "name": "core/clock",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers",
                "stable": false
              },
              {
                "name": "agenda/current-speaker-chyron",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers-overlay",
                "stable": true
              }
            ],
            "elements_preview": [],
            "elements_history": [],
            "scale": 0,
            "scroll": 0,
            "name": "sideprojector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/tag:1": []byte(`{
            "id": 1,
            "name": "T1"
          }`),
		"core/tag:2": []byte(`{
            "id": 2,
            "name": "T2"
          }`),
		"mediafiles/mediafile:2": []byte(`{
            "id": 2,
            "title": "A.txt",
            "media_url_prefix": "/media/",
            "filesize": "< 1 kB",
            "mimetype": "text/plain",
            "pdf_information": {},
            "access_groups_id": [],
            "create_timestamp": "2020-08-11T11:16:07.013124+02:00",
            "is_directory": false,
            "path": "A.txt",
            "parent_id": null,
            "list_of_speakers_id": 5,
            "inherited_access_groups_id": true
          }`),
		"motions/category:1": []byte(`{
            "id": 1,
            "name": "first",
            "prefix": "A",
            "parent_id": null,
            "weight": 2,
            "level": 0
          }`),
		"motions/category:2": []byte(`{
            "id": 2,
            "name": "second",
            "prefix": "B",
            "parent_id": 1,
            "weight": 4,
            "level": 1
          }`),
		"motions/category:3": []byte(`{
            "id": 3,
            "name": "third",
            "prefix": "C",
            "parent_id": null,
            "weight": 6,
            "level": 0
          }`),
		"motions/motion-block:1": []byte(`{
            "id": 1,
            "title": "block",
            "agenda_item_id": 8,
            "list_of_speakers_id": 12,
            "internal": false,
            "motions_id": [
              1,
              2
            ]
          }`),
		"motions/motion-option:1": []byte(`{
            "id": 1,
            "yes": "0.000000",
            "no": "1.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-option:2": []byte(`{
            "id": 2,
            "yes": "1.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 2,
            "pollstate": 4
          }`),
		"motions/motion-poll:1": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "named",
            "title": "Abstimmung",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              1
            ],
            "id": 1,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-poll:2": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "pseudoanonymous",
            "title": "Abstimmung (2)",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              2
            ],
            "id": 2,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "N",
            "user_id": 5,
            "option_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-vote:2": []byte(`{
            "id": 2,
            "weight": "1.000000",
            "value": "Y",
            "user_id": null,
            "option_id": 2,
            "pollstate": 4
          }`),
		"motions/motion:2": []byte(`{
            "id": 2,
            "identifier": "Ä-1",
            "title": "Änderungsantrag zu Leadmotion1",
            "text": "",
            "amendment_paragraphs": [
              "<p>Lorem ipsum dolor sit amet, consectedfgddfgdf&nbsp; gdfgdfg dfg dfg ww ontes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>"
            ],
            "modified_final_version": "",
            "reason": "",
            "parent_id": 1,
            "category_id": 2,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 3,
                "user_id": 1,
                "motion_id": 2,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [],
            "attachments_id": [],
            "agenda_item_id": 6,
            "list_of_speakers_id": 9,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:39:35.025914+02:00",
            "last_modified": "2020-08-11T12:41:55.666495+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/motion:3": []byte(`{
            "id": 3,
            "identifier": "2",
            "title": "Public",
            "text": "<p>a</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>a</p>",
            "parent_id": null,
            "category_id": 1,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 2,
            "origin": "",
            "submitters": [
              {
                "id": 4,
                "user_id": 1,
                "motion_id": 3,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [
              1
            ],
            "attachments_id": [],
            "agenda_item_id": 7,
            "list_of_speakers_id": 10,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T12:24:45.289233+02:00",
            "last_modified": "2020-08-11T12:41:51.319986+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/state:1": []byte(`{
            "id": 1,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              2,
              3,
              4
            ],
            "workflow_id": 1
          }`),
		"motions/state:10": []byte(`{
            "id": 10,
            "name": "withdrawed",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:11": []byte(`{
            "id": 11,
            "name": "adjourned",
            "recommendation_label": "Adjournment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:12": []byte(`{
            "id": 12,
            "name": "not concerned",
            "recommendation_label": "No concernment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:13": []byte(`{
            "id": 13,
            "name": "refered to committee",
            "recommendation_label": "Referral to committee",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:14": []byte(`{
            "id": 14,
            "name": "needs review",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:15": []byte(`{
            "id": 15,
            "name": "rejected (not authorized)",
            "recommendation_label": "Rejection (not authorized)",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:2": []byte(`{
            "id": 2,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:3": []byte(`{
            "id": 3,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:4": []byte(`{
            "id": 4,
            "name": "not decided",
            "recommendation_label": "No decision",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:5": []byte(`{
            "id": 5,
            "name": "in progress",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [
              "is_submitter"
            ],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": true,
            "dont_set_identifier": true,
            "show_state_extension_field": true,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              6,
              10
            ],
            "workflow_id": 2
          }`),
		"motions/state:6": []byte(`{
            "id": 6,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": true,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              7,
              10,
              15
            ],
            "workflow_id": 2
          }`),
		"motions/state:7": []byte(`{
            "id": 7,
            "name": "permitted",
            "recommendation_label": "Permission",
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": true,
            "next_states_id": [
              8,
              9,
              10,
              11,
              12,
              13,
              14
            ],
            "workflow_id": 2
          }`),
		"motions/state:8": []byte(`{
            "id": 8,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:9": []byte(`{
            "id": 9,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/statute-paragraph:1": []byte(`{
            "id": 1,
            "title": "Statute",
            "text": "<p>test</p>",
            "weight": 10000
          }`),
		"motions/workflow:1": []byte(`{
            "id": 1,
            "name": "Simple Workflow",
            "states_id": [
              1,
              2,
              3,
              4
            ],
            "first_state_id": 1
          }`),
		"motions/workflow:2": []byte(`{
            "id": 2,
            "name": "Complex Workflow",
            "states_id": [
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13,
              14,
              15
            ],
            "first_state_id": 5
          }`),
		"topics/topic:1": []byte(`{
            "id": 1,
            "title": "Topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 1,
            "list_of_speakers_id": 1
          }`),
		"topics/topic:2": []byte(`{
            "id": 2,
            "title": "Hidden",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 2,
            "list_of_speakers_id": 2
          }`),
		"topics/topic:3": []byte(`{
            "id": 3,
            "title": "Internal",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 3,
            "list_of_speakers_id": 3
          }`),
		"topics/topic:4": []byte(`{
            "id": 4,
            "title": "Another public topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 9,
            "list_of_speakers_id": 14
          }`),
		"users/group:1": []byte(`{
            "id": 1,
            "name": "Default",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_see",
              "users.can_change_password"
            ]
          }`),
		"users/group:2": []byte(`{
            "id": 2,
            "name": "Admin",
            "permissions": []
          }`),
		"users/group:3": []byte(`{
            "id": 3,
            "name": "Delegates",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password"
            ]
          }`),
		"users/group:4": []byte(`{
            "id": 4,
            "name": "Staff",
            "permissions": [
              "agenda.can_manage",
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_manage_list_of_speakers",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_manage",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_history",
              "core.can_manage_projector",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "core.can_manage_tags",
              "mediafiles.can_manage",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_manage",
              "motions.can_manage_metadata",
              "motions.can_see",
              "motions.can_see_internal",
              "users.can_change_password",
              "users.can_manage",
              "users.can_see_extra_data",
              "users.can_see_name"
            ]
          }`),
		"users/group:5": []byte(`{
            "id": 5,
            "name": "Committees",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password",
              "users.can_see_name"
            ]
          }`),
		"users/user:1": []byte(`{
            "first_name": "",
            "username": "admin",
            "about_me": "",
            "id": 1,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              2
            ],
            "number": "",
            "last_name": "Administrator",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:2": []byte(`{
            "first_name": "candidate1",
            "username": "candidate1",
            "about_me": "",
            "id": 2,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "groups_id": [],
            "title": "",
            "email": "",
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:3": []byte(`{
            "first_name": "candidate2",
            "username": "candidate2",
            "about_me": "",
            "id": 3,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:4": []byte(`{
            "first_name": "a",
            "username": "a",
            "about_me": "",
            "id": 4,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:5": []byte(`{
            "first_name": "b",
            "username": "b",
            "about_me": "",
            "id": 5,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:6": []byte(`{
            "first_name": "speaker1",
            "username": "speaker1",
            "about_me": "",
            "id": 6,
            "structure_level": "layer X",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "title",
            "groups_id": [],
            "number": "3",
            "last_name": "the last name",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:7": []byte(`{
            "first_name": "speaker2",
            "username": "speaker2",
            "about_me": "",
            "id": 7,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
	},
	3: {
		"agenda/item:1": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            },
            "is_internal": false,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": null,
            "id": 1,
            "level": 0,
            "title_information": {
              "title": "Topic"
            },
            "closed": false,
            "weight": 2,
            "duration": null
          }`),
		"agenda/item:3": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": null,
            "id": 3,
            "level": 0,
            "title_information": {
              "title": "Internal"
            },
            "closed": false,
            "weight": 8,
            "duration": null
          }`),
		"agenda/item:5": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 5,
            "level": 1,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "closed": false,
            "weight": 14,
            "duration": null
          }`),
		"agenda/item:6": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 6,
            "level": 1,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "closed": false,
            "weight": 16,
            "duration": 0
          }`),
		"agenda/item:7": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": 6,
            "id": 7,
            "level": 2,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "closed": false,
            "weight": 18,
            "duration": null
          }`),
		"agenda/list-of-speakers:1": []byte(`{
            "id": 1,
            "title_information": {
              "title": "Topic"
            },
            "speakers": [
              {
                "id": 3,
                "user_id": 6,
                "begin_time": "2020-08-11T12:28:30.894574+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:10": []byte(`{
            "id": 10,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:11": []byte(`{
            "id": 11,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:12": []byte(`{
            "id": 12,
            "title_information": {
              "title": "block"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:13": []byte(`{
            "id": 13,
            "title_information": {
              "title": "block internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:14": []byte(`{
            "id": 14,
            "title_information": {
              "title": "Another public topic"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:2": []byte(`{
            "id": 2,
            "title_information": {
              "title": "Hidden"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:3": []byte(`{
            "id": 3,
            "title_information": {
              "title": "Internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:4": []byte(`{
            "id": 4,
            "title_information": {
              "title": "folder"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:5": []byte(`{
            "id": 5,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:6": []byte(`{
            "id": 6,
            "title_information": {
              "title": "in.jpg"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:7": []byte(`{
            "id": 7,
            "title_information": {
              "title": "Assignment"
            },
            "speakers": [
              {
                "id": 4,
                "user_id": 6,
                "begin_time": "2020-08-11T12:29:55.054553+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              },
              {
                "id": 6,
                "user_id": 7,
                "begin_time": null,
                "end_time": null,
                "weight": 2,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "assignments/assignment",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:8": []byte(`{
            "id": 8,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:9": []byte(`{
            "id": 9,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            }
          }`),
		"assignments/assignment-option:1": []byte(`{
            "user_id": 2,
            "weight": 1,
            "id": 1,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-option:2": []byte(`{
            "user_id": 3,
            "weight": 3,
            "id": 2,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-poll:1": []byte(`{
            "assignment_id": 1,
            "description": "",
            "pollmethod": "votes",
            "votes_amount": 1,
            "allow_multiple_votes_per_candidate": false,
            "global_no": true,
            "global_abstain": true,
            "state": 2,
            "type": "named",
            "title": "Wahlgang",
            "groups_id": [
              3
            ],
            "options_id": [
              1,
              2
            ],
            "id": 1,
            "onehundred_percent_base": "valid",
            "majority_method": "simple",
            "user_has_voted": false
          }`),
		"assignments/assignment:1": []byte(`{
            "id": 1,
            "title": "Assignment",
            "description": "",
            "open_posts": 1,
            "phase": 0,
            "assignment_related_users": [
              {
                "id": 1,
                "user_id": 2,
                "weight": 1
              },
              {
                "id": 3,
                "user_id": 3,
                "weight": 3
              }
            ],
            "default_poll_description": "",
            "agenda_item_id": 4,
            "list_of_speakers_id": 7,
            "tags_id": [
              2
            ],
            "attachments_id": [],
            "number_poll_candidates": false,
            "polls_id": [
              1
            ]
          }`),
		"core/config:10": []byte(`{
            "id": 10,
            "key": "general_system_conference_show",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show live conference window",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 140,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:100": []byte(`{
            "id": 100,
            "key": "motion_poll_default_100_percent_base",
            "value": "YNA",
            "data": {
              "defaultValue": "YNA",
              "inputType": "choice",
              "label": "Default 100 % base of a voting result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 370,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:101": []byte(`{
            "id": 101,
            "key": "motion_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:102": []byte(`{
            "id": 102,
            "key": "motion_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 372,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:103": []byte(`{
            "id": 103,
            "key": "motions_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 373,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:104": []byte(`{
            "id": 104,
            "key": "motions_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 374,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:105": []byte(`{
            "id": 105,
            "key": "motions_export_title",
            "value": "Motions",
            "data": {
              "defaultValue": "Motions",
              "inputType": "string",
              "label": "Title for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 380,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:106": []byte(`{
            "id": 106,
            "key": "motions_export_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 382,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:107": []byte(`{
            "id": 107,
            "key": "motions_export_submitter_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show submitters and recommendation/state in table of contents",
              "helpText": "",
              "choices": null,
              "weight": 384,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:108": []byte(`{
            "id": 108,
            "key": "motions_export_follow_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show checkbox to record decision",
              "helpText": "",
              "choices": null,
              "weight": 386,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:109": []byte(`{
            "id": 109,
            "key": "assignment_poll_method",
            "value": "votes",
            "data": {
              "defaultValue": "votes",
              "inputType": "choice",
              "label": "Default election method",
              "helpText": "",
              "choices": [
                {
                  "value": "votes",
                  "display_name": "Yes per candidate"
                },
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                }
              ],
              "weight": 400,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:11": []byte(`{
            "id": 11,
            "key": "general_system_conference_auto_connect",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Connect all users to live conference automatically",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 141,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:110": []byte(`{
            "id": 110,
            "key": "assignment_poll_default_type",
            "value": "analog",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 403,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:111": []byte(`{
            "id": 111,
            "key": "assignment_poll_default_100_percent_base",
            "value": "valid",
            "data": {
              "defaultValue": "valid",
              "inputType": "choice",
              "label": "Default 100 % base of an election result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                },
                {
                  "value": "votes",
                  "display_name": "Sum of votes including general No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 405,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:112": []byte(`{
            "id": 112,
            "key": "assignment_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 410,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:113": []byte(`{
            "id": 113,
            "key": "assignment_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:114": []byte(`{
            "id": 114,
            "key": "assignment_poll_sort_poll_result_by_votes",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Sort election results by amount of votes",
              "helpText": "",
              "choices": null,
              "weight": 420,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:115": []byte(`{
            "id": 115,
            "key": "assignment_poll_add_candidates_to_list_of_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Put all candidates on the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 425,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:116": []byte(`{
            "id": 116,
            "key": "assignments_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 430,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:117": []byte(`{
            "id": 117,
            "key": "assignments_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 435,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:118": []byte(`{
            "id": 118,
            "key": "assignments_pdf_title",
            "value": "Elections",
            "data": {
              "defaultValue": "Elections",
              "inputType": "string",
              "label": "Title for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 460,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:119": []byte(`{
            "id": 119,
            "key": "assignments_pdf_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 470,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:12": []byte(`{
            "id": 12,
            "key": "general_system_conference_los_restriction",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow only current speakers and list of speakers managers to enter the live conference",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 142,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:13": []byte(`{
            "id": 13,
            "key": "general_system_stream_url",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Livestream url",
              "helpText": "Remove URL to deactivate livestream. Check extra group permission to see livestream.",
              "choices": null,
              "weight": 143,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:14": []byte(`{
            "id": 14,
            "key": "general_system_enable_anonymous",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow access for anonymous guest users",
              "helpText": "",
              "choices": null,
              "weight": 150,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:15": []byte(`{
            "id": 15,
            "key": "general_login_info_text",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Show this text on the login page",
              "helpText": "",
              "choices": null,
              "weight": 152,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:16": []byte(`{
            "id": 16,
            "key": "openslides_theme",
            "value": "openslides-default-light-theme",
            "data": {
              "defaultValue": "openslides-default-light-theme",
              "inputType": "choice",
              "label": "OpenSlides Theme",
              "helpText": "",
              "choices": [
                {
                  "value": "openslides-default-light-theme",
                  "display_name": "OpenSlides Default"
                },
                {
                  "value": "openslides-default-dark-theme",
                  "display_name": "OpenSlides Dark"
                },
                {
                  "value": "openslides-red-light-theme",
                  "display_name": "OpenSlides Red"
                },
                {
                  "value": "openslides-red-dark-theme",
                  "display_name": "OpenSlides Red Dark"
                },
                {
                  "value": "openslides-green-light-theme",
                  "display_name": "OpenSlides Green"
                },
                {
                  "value": "openslides-green-dark-theme",
                  "display_name": "OpenSlides Green Dark"
                },
                {
                  "value": "openslides-solarized-dark-theme",
                  "display_name": "OpenSlides Solarized"
                }
              ],
              "weight": 154,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:17": []byte(`{
            "id": 17,
            "key": "general_csv_separator",
            "value": ",",
            "data": {
              "defaultValue": ",",
              "inputType": "string",
              "label": "Separator used for all csv exports and examples",
              "helpText": "",
              "choices": null,
              "weight": 160,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:18": []byte(`{
            "id": 18,
            "key": "general_csv_encoding",
            "value": "utf-8",
            "data": {
              "defaultValue": "utf-8",
              "inputType": "choice",
              "label": "Default encoding for all csv exports",
              "helpText": "",
              "choices": [
                {
                  "value": "utf-8",
                  "display_name": "UTF-8"
                },
                {
                  "value": "iso-8859-15",
                  "display_name": "ISO-8859-15"
                }
              ],
              "weight": 162,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:19": []byte(`{
            "id": 19,
            "key": "general_export_pdf_pagenumber_alignment",
            "value": "center",
            "data": {
              "defaultValue": "center",
              "inputType": "choice",
              "label": "Page number alignment in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "left",
                  "display_name": "Left"
                },
                {
                  "value": "center",
                  "display_name": "Center"
                },
                {
                  "value": "right",
                  "display_name": "Right"
                }
              ],
              "weight": 164,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:2": []byte(`{
            "id": 2,
            "key": "general_event_name",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Event name",
              "helpText": "",
              "choices": null,
              "weight": 110,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:20": []byte(`{
            "id": 20,
            "key": "general_export_pdf_fontsize",
            "value": "10",
            "data": {
              "defaultValue": "10",
              "inputType": "choice",
              "label": "Standard font size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "10",
                  "display_name": "10"
                },
                {
                  "value": "11",
                  "display_name": "11"
                },
                {
                  "value": "12",
                  "display_name": "12"
                }
              ],
              "weight": 166,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:21": []byte(`{
            "id": 21,
            "key": "general_export_pdf_pagesize",
            "value": "A4",
            "data": {
              "defaultValue": "A4",
              "inputType": "choice",
              "label": "Standard page size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "A4",
                  "display_name": "DIN A4"
                },
                {
                  "value": "A5",
                  "display_name": "DIN A5"
                }
              ],
              "weight": 168,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:22": []byte(`{
            "id": 22,
            "key": "logos_available",
            "value": [
              "logo_projector_main",
              "logo_projector_header",
              "logo_web_header",
              "logo_pdf_header_L",
              "logo_pdf_header_R",
              "logo_pdf_footer_L",
              "logo_pdf_footer_R",
              "logo_pdf_ballot_paper"
            ],
            "data": null
          }`),
		"core/config:23": []byte(`{
            "id": 23,
            "key": "logo_projector_main",
            "value": {
              "display_name": "Projector logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:24": []byte(`{
            "id": 24,
            "key": "logo_projector_header",
            "value": {
              "display_name": "Projector header image",
              "path": ""
            },
            "data": null
          }`),
		"core/config:25": []byte(`{
            "id": 25,
            "key": "logo_web_header",
            "value": {
              "display_name": "Web interface header logo",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:26": []byte(`{
            "id": 26,
            "key": "logo_pdf_header_L",
            "value": {
              "display_name": "PDF header logo (left)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:27": []byte(`{
            "id": 27,
            "key": "logo_pdf_header_R",
            "value": {
              "display_name": "PDF header logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:28": []byte(`{
            "id": 28,
            "key": "logo_pdf_footer_L",
            "value": {
              "display_name": "PDF footer logo (left)",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:29": []byte(`{
            "id": 29,
            "key": "logo_pdf_footer_R",
            "value": {
              "display_name": "PDF footer logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:3": []byte(`{
            "id": 3,
            "key": "general_event_description",
            "value": "Presentation and assembly system",
            "data": {
              "defaultValue": "Presentation and assembly system",
              "inputType": "string",
              "label": "Short description of event",
              "helpText": "",
              "choices": null,
              "weight": 115,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:30": []byte(`{
            "id": 30,
            "key": "logo_pdf_ballot_paper",
            "value": {
              "display_name": "PDF ballot paper logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:31": []byte(`{
            "id": 31,
            "key": "fonts_available",
            "value": [
              "font_regular",
              "font_italic",
              "font_bold",
              "font_bold_italic",
              "font_monospace"
            ],
            "data": null
          }`),
		"core/config:32": []byte(`{
            "id": 32,
            "key": "font_regular",
            "value": {
              "display_name": "Font regular",
              "default": "assets/fonts/fira-sans-latin-400.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:33": []byte(`{
            "id": 33,
            "key": "font_italic",
            "value": {
              "display_name": "Font italic",
              "default": "assets/fonts/fira-sans-latin-400italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:34": []byte(`{
            "id": 34,
            "key": "font_bold",
            "value": {
              "display_name": "Font bold",
              "default": "assets/fonts/fira-sans-latin-500.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:35": []byte(`{
            "id": 35,
            "key": "font_bold_italic",
            "value": {
              "display_name": "Font bold italic",
              "default": "assets/fonts/fira-sans-latin-500italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:36": []byte(`{
            "id": 36,
            "key": "font_monospace",
            "value": {
              "display_name": "Font monospace",
              "default": "assets/fonts/roboto-condensed-bold.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:37": []byte(`{
            "id": 37,
            "key": "translations",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "translations",
              "label": "Custom translations",
              "helpText": "",
              "choices": null,
              "weight": 1000,
              "group": "Custom translations",
              "subgroup": "General"
            }
          }`),
		"core/config:38": []byte(`{
            "id": 38,
            "key": "config_version",
            "value": 2,
            "data": null
          }`),
		"core/config:39": []byte(`{
            "id": 39,
            "key": "db_id",
            "value": "2c3bd736c87a48a4be1f0dc707702144",
            "data": null
          }`),
		"core/config:4": []byte(`{
            "id": 4,
            "key": "general_event_date",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event date",
              "helpText": "",
              "choices": null,
              "weight": 120,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:40": []byte(`{
            "id": 40,
            "key": "users_sort_by",
            "value": "first_name",
            "data": {
              "defaultValue": "first_name",
              "inputType": "choice",
              "label": "Sort name of participants by",
              "helpText": "",
              "choices": [
                {
                  "value": "first_name",
                  "display_name": "Given name"
                },
                {
                  "value": "last_name",
                  "display_name": "Surname"
                },
                {
                  "value": "number",
                  "display_name": "Participant number"
                }
              ],
              "weight": 510,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:41": []byte(`{
            "id": 41,
            "key": "users_enable_presence_view",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Enable participant presence view",
              "helpText": "",
              "choices": null,
              "weight": 511,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:42": []byte(`{
            "id": 42,
            "key": "users_allow_self_set_present",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow users to set themselves as present",
              "helpText": "e.g. for online meetings",
              "choices": null,
              "weight": 512,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:43": []byte(`{
            "id": 43,
            "key": "users_activate_vote_weight",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate vote weight",
              "helpText": "",
              "choices": null,
              "weight": 513,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:44": []byte(`{
            "id": 44,
            "key": "users_pdf_welcometitle",
            "value": "Welcome to OpenSlides",
            "data": {
              "defaultValue": "Welcome to OpenSlides",
              "inputType": "string",
              "label": "Title for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 520,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:45": []byte(`{
            "id": 45,
            "key": "users_pdf_welcometext",
            "value": "[Place for your welcome and help text.]",
            "data": {
              "defaultValue": "[Place for your welcome and help text.]",
              "inputType": "string",
              "label": "Help text for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 530,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:46": []byte(`{
            "id": 46,
            "key": "users_pdf_url",
            "value": "http://example.com:8000",
            "data": {
              "defaultValue": "http://example.com:8000",
              "inputType": "string",
              "label": "System URL",
              "helpText": "Used for QRCode in PDF of access data.",
              "choices": null,
              "weight": 540,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:47": []byte(`{
            "id": 47,
            "key": "users_pdf_wlan_ssid",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN name (SSID)",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 550,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:48": []byte(`{
            "id": 48,
            "key": "users_pdf_wlan_password",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN password",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 560,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:49": []byte(`{
            "id": 49,
            "key": "users_pdf_wlan_encryption",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "choice",
              "label": "WLAN encryption",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": [
                {
                  "value": "",
                  "display_name": "---------"
                },
                {
                  "value": "WEP",
                  "display_name": "WEP"
                },
                {
                  "value": "WPA",
                  "display_name": "WPA/WPA2"
                },
                {
                  "value": "nopass",
                  "display_name": "No encryption"
                }
              ],
              "weight": 570,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:5": []byte(`{
            "id": 5,
            "key": "general_event_location",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event location",
              "helpText": "",
              "choices": null,
              "weight": 125,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:50": []byte(`{
            "id": 50,
            "key": "users_email_sender",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Sender name",
              "helpText": "The sender address is defined in the OpenSlides server settings and should modified by administrator only.",
              "choices": null,
              "weight": 600,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:51": []byte(`{
            "id": 51,
            "key": "users_email_replyto",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Reply address",
              "helpText": "",
              "choices": null,
              "weight": 601,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:52": []byte(`{
            "id": 52,
            "key": "users_email_subject",
            "value": "OpenSlides access data",
            "data": {
              "defaultValue": "OpenSlides access data",
              "inputType": "string",
              "label": "Email subject",
              "helpText": "You can use {event_name} and {username} as placeholder.",
              "choices": null,
              "weight": 605,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:53": []byte(`{
            "id": 53,
            "key": "users_email_body",
            "value": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
            "data": {
              "defaultValue": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
              "inputType": "text",
              "label": "Email body",
              "helpText": "Use these placeholders: {name}, {event_name}, {url}, {username}, {password}. The url referrs to the system url.",
              "choices": null,
              "weight": 610,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:54": []byte(`{
            "id": 54,
            "key": "agenda_start_event_date_time",
            "value": null,
            "data": {
              "defaultValue": null,
              "inputType": "datetimepicker",
              "label": "Begin of event",
              "helpText": "Input format: DD.MM.YYYY HH:MM",
              "choices": null,
              "weight": 200,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:55": []byte(`{
            "id": 55,
            "key": "agenda_show_subtitle",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show subtitles in the agenda",
              "helpText": "",
              "choices": null,
              "weight": 201,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:56": []byte(`{
            "id": 56,
            "key": "agenda_enable_numbering",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Enable numbering for agenda items",
              "helpText": "",
              "choices": null,
              "weight": 205,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:57": []byte(`{
            "id": 57,
            "key": "agenda_number_prefix",
            "value": "TOP",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Numbering prefix for agenda items",
              "helpText": "This prefix will be set if you run the automatic agenda numbering.",
              "choices": null,
              "weight": 206,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:58": []byte(`{
            "id": 58,
            "key": "agenda_numeral_system",
            "value": "arabic",
            "data": {
              "defaultValue": "arabic",
              "inputType": "choice",
              "label": "Numeral system for agenda items",
              "helpText": "",
              "choices": [
                {
                  "value": "arabic",
                  "display_name": "Arabic"
                },
                {
                  "value": "roman",
                  "display_name": "Roman"
                }
              ],
              "weight": 207,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:59": []byte(`{
            "id": 59,
            "key": "agenda_item_creation",
            "value": "default_yes",
            "data": {
              "defaultValue": "default_yes",
              "inputType": "choice",
              "label": "Add to agenda",
              "helpText": "",
              "choices": [
                {
                  "value": "always",
                  "display_name": "Always"
                },
                {
                  "value": "never",
                  "display_name": "Never"
                },
                {
                  "value": "default_yes",
                  "display_name": "Ask, default yes"
                },
                {
                  "value": "default_no",
                  "display_name": "Ask, default no"
                }
              ],
              "weight": 210,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:6": []byte(`{
            "id": 6,
            "key": "general_event_legal_notice",
            "value": "<a href=\"http://www.openslides.org\">OpenSlides</a> is a free web based presentation and assembly system for visualizing and controlling agenda, motions and elections of an assembly.",
            "data": null
          }`),
		"core/config:60": []byte(`{
            "id": 60,
            "key": "agenda_new_items_default_visibility",
            "value": "2",
            "data": {
              "defaultValue": "2",
              "inputType": "choice",
              "label": "Default visibility for new agenda items (except topics)",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Public item"
                },
                {
                  "value": "2",
                  "display_name": "Internal item"
                },
                {
                  "value": "3",
                  "display_name": "Hidden item"
                }
              ],
              "weight": 211,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:61": []byte(`{
            "id": 61,
            "key": "agenda_hide_internal_items_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Hide internal items when projecting subitems",
              "helpText": "",
              "choices": null,
              "weight": 212,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:62": []byte(`{
            "id": 62,
            "key": "agenda_show_last_speakers",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of last speakers to be shown on the projector",
              "helpText": "",
              "choices": null,
              "weight": 220,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:63": []byte(`{
            "id": 63,
            "key": "agenda_show_next_speakers",
            "value": -1,
            "data": {
              "defaultValue": -1,
              "inputType": "integer",
              "label": "Number of the next speakers to be shown on the projector",
              "helpText": "Enter number of the next shown speakers. Choose -1 to show all next speakers.",
              "choices": null,
              "weight": 222,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:64": []byte(`{
            "id": 64,
            "key": "agenda_countdown_warning_time",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Show orange countdown in the last x seconds of speaking time",
              "helpText": "Enter duration in seconds. Choose 0 to disable warning color.",
              "choices": null,
              "weight": 224,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:65": []byte(`{
            "id": 65,
            "key": "projector_default_countdown",
            "value": 60,
            "data": {
              "defaultValue": 60,
              "inputType": "integer",
              "label": "Predefined seconds of new countdowns",
              "helpText": "",
              "choices": null,
              "weight": 226,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:66": []byte(`{
            "id": 66,
            "key": "agenda_couple_countdown_and_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Couple countdown with the list of speakers",
              "helpText": "[Begin speech] starts the countdown, [End speech] stops the countdown.",
              "choices": null,
              "weight": 228,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:67": []byte(`{
            "id": 67,
            "key": "agenda_hide_amount_of_speakers",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide the amount of speakers in subtitle of list of speakers slide",
              "helpText": "",
              "choices": null,
              "weight": 230,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:68": []byte(`{
            "id": 68,
            "key": "agenda_present_speakers_only",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Only present participants can be added to the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 232,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:69": []byte(`{
            "id": 69,
            "key": "agenda_show_first_contribution",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show hint »first speech« in the list of speakers management view",
              "helpText": "",
              "choices": null,
              "weight": 234,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:7": []byte(`{
            "id": 7,
            "key": "general_event_privacy_policy",
            "value": "",
            "data": null
          }`),
		"core/config:70": []byte(`{
            "id": 70,
            "key": "motions_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new motions",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 310,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:71": []byte(`{
            "id": 71,
            "key": "motions_statute_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new statute amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 312,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:72": []byte(`{
            "id": 72,
            "key": "motions_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 314,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:73": []byte(`{
            "id": 73,
            "key": "motions_preamble",
            "value": "The assembly may decide:",
            "data": {
              "defaultValue": "The assembly may decide:",
              "inputType": "string",
              "label": "Motion preamble",
              "helpText": "",
              "choices": null,
              "weight": 320,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:74": []byte(`{
            "id": 74,
            "key": "motions_default_line_numbering",
            "value": "outside",
            "data": {
              "defaultValue": "outside",
              "inputType": "choice",
              "label": "Default line numbering",
              "helpText": "",
              "choices": [
                {
                  "value": "outside",
                  "display_name": "outside"
                },
                {
                  "value": "inline",
                  "display_name": "inline"
                },
                {
                  "value": "none",
                  "display_name": "Disabled"
                }
              ],
              "weight": 322,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:75": []byte(`{
            "id": 75,
            "key": "motions_line_length",
            "value": 85,
            "data": {
              "defaultValue": 85,
              "inputType": "integer",
              "label": "Line length",
              "helpText": "The maximum number of characters per line. Relevant when line numbering is enabled. Min: 40",
              "choices": null,
              "weight": 323,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:76": []byte(`{
            "id": 76,
            "key": "motions_reason_required",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Reason required for creating new motion",
              "helpText": "",
              "choices": null,
              "weight": 324,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:77": []byte(`{
            "id": 77,
            "key": "motions_disable_text_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide motion text on projector",
              "helpText": "",
              "choices": null,
              "weight": 325,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:78": []byte(`{
            "id": 78,
            "key": "motions_disable_reason_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide reason on projector",
              "helpText": "",
              "choices": null,
              "weight": 326,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:79": []byte(`{
            "id": 79,
            "key": "motions_disable_recommendation_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide recommendation on projector",
              "helpText": "",
              "choices": null,
              "weight": 327,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:8": []byte(`{
            "id": 8,
            "key": "general_event_welcome_title",
            "value": "Welcome to OpenSlides",
            "data": null
          }`),
		"core/config:80": []byte(`{
            "id": 80,
            "key": "motions_hide_referring_motions",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide referring motions",
              "helpText": "",
              "choices": null,
              "weight": 328,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:81": []byte(`{
            "id": 81,
            "key": "motions_disable_sidebox_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show meta information box below the title on projector",
              "helpText": "",
              "choices": null,
              "weight": 329,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:82": []byte(`{
            "id": 82,
            "key": "motions_show_sequential_numbers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show the sequential number for a motion",
              "helpText": "In motion list, motion detail and PDF.",
              "choices": null,
              "weight": 330,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:83": []byte(`{
            "id": 83,
            "key": "motions_recommendations_by",
            "value": "ABK",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender",
              "helpText": "Will be displayed as label before selected recommendation. Use an empty value to disable the recommendation system.",
              "choices": null,
              "weight": 332,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:84": []byte(`{
            "id": 84,
            "key": "motions_statute_recommendations_by",
            "value": "ABK2",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender for statute amendments",
              "helpText": "Will be displayed as label before selected recommendation in statute amendments.",
              "choices": null,
              "weight": 333,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:85": []byte(`{
            "id": 85,
            "key": "motions_recommendation_text_mode",
            "value": "diff",
            "data": {
              "defaultValue": "diff",
              "inputType": "choice",
              "label": "Default text version for change recommendations",
              "helpText": "",
              "choices": [
                {
                  "value": "original",
                  "display_name": "Original version"
                },
                {
                  "value": "changed",
                  "display_name": "Changed version"
                },
                {
                  "value": "diff",
                  "display_name": "Diff version"
                },
                {
                  "value": "agreed",
                  "display_name": "Final version"
                }
              ],
              "weight": 334,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:86": []byte(`{
            "id": 86,
            "key": "motions_motions_sorting",
            "value": "identifier",
            "data": {
              "defaultValue": "identifier",
              "inputType": "choice",
              "label": "Sort motions by",
              "helpText": "",
              "choices": [
                {
                  "value": "weight",
                  "display_name": "Call list"
                },
                {
                  "value": "identifier",
                  "display_name": "Identifier"
                }
              ],
              "weight": 335,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:87": []byte(`{
            "id": 87,
            "key": "motions_identifier",
            "value": "per_category",
            "data": {
              "defaultValue": "per_category",
              "inputType": "choice",
              "label": "Identifier",
              "helpText": "",
              "choices": [
                {
                  "value": "per_category",
                  "display_name": "Numbered per category"
                },
                {
                  "value": "serially_numbered",
                  "display_name": "Serially numbered"
                },
                {
                  "value": "manually",
                  "display_name": "Set it manually"
                }
              ],
              "weight": 340,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:88": []byte(`{
            "id": 88,
            "key": "motions_identifier_min_digits",
            "value": 1,
            "data": {
              "defaultValue": 1,
              "inputType": "integer",
              "label": "Number of minimal digits for identifier",
              "helpText": "Uses leading zeros to sort motions correctly by identifier.",
              "choices": null,
              "weight": 342,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:89": []byte(`{
            "id": 89,
            "key": "motions_identifier_with_blank",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow blank in identifier",
              "helpText": "Blank between prefix and number, e.g. 'A 001'.",
              "choices": null,
              "weight": 344,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:9": []byte(`{
            "id": 9,
            "key": "general_event_welcome_text",
            "value": "[Space for your welcome text.]",
            "data": null
          }`),
		"core/config:90": []byte(`{
            "id": 90,
            "key": "motions_statutes_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate statute amendments",
              "helpText": "",
              "choices": null,
              "weight": 350,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:91": []byte(`{
            "id": 91,
            "key": "motions_amendments_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate amendments",
              "helpText": "",
              "choices": null,
              "weight": 351,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:92": []byte(`{
            "id": 92,
            "key": "motions_amendments_main_table",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show amendments together with motions",
              "helpText": "",
              "choices": null,
              "weight": 352,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:93": []byte(`{
            "id": 93,
            "key": "motions_amendments_prefix",
            "value": "Ä-",
            "data": {
              "defaultValue": "-",
              "inputType": "string",
              "label": "Prefix for the identifier for amendments",
              "helpText": "",
              "choices": null,
              "weight": 353,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:94": []byte(`{
            "id": 94,
            "key": "motions_amendments_text_mode",
            "value": "paragraph",
            "data": {
              "defaultValue": "paragraph",
              "inputType": "choice",
              "label": "How to create new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "freestyle",
                  "display_name": "Empty text field"
                },
                {
                  "value": "fulltext",
                  "display_name": "Edit the whole motion text"
                },
                {
                  "value": "paragraph",
                  "display_name": "Paragraph-based, Diff-enabled"
                }
              ],
              "weight": 354,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:95": []byte(`{
            "id": 95,
            "key": "motions_amendments_multiple_paragraphs",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Amendments can change multiple paragraphs",
              "helpText": "",
              "choices": null,
              "weight": 355,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:96": []byte(`{
            "id": 96,
            "key": "motions_amendments_of_amendments",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow amendments of amendments",
              "helpText": "",
              "choices": null,
              "weight": 356,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:97": []byte(`{
            "id": 97,
            "key": "motions_min_supporters",
            "value": 1,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of (minimum) required supporters for a motion",
              "helpText": "Choose 0 to disable the supporting system.",
              "choices": null,
              "weight": 360,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:98": []byte(`{
            "id": 98,
            "key": "motions_remove_supporters",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Remove all supporters of a motion if a submitter edits his motion in early state",
              "helpText": "",
              "choices": null,
              "weight": 361,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:99": []byte(`{
            "id": 99,
            "key": "motion_poll_default_type",
            "value": "named",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 367,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/countdown:1": []byte(`{
            "id": 1,
            "title": "Default countdown",
            "description": "",
            "default_time": 60,
            "countdown_time": 1597141855.090748,
            "running": true
          }`),
		"core/projection-default:1": []byte(`{
            "id": 1,
            "name": "agenda_all_items",
            "display_name": "Agenda",
            "projector_id": 1
          }`),
		"core/projection-default:10": []byte(`{
            "id": 10,
            "name": "messages",
            "display_name": "Messages",
            "projector_id": 1
          }`),
		"core/projection-default:11": []byte(`{
            "id": 11,
            "name": "countdowns",
            "display_name": "Countdowns",
            "projector_id": 1
          }`),
		"core/projection-default:12": []byte(`{
            "id": 12,
            "name": "assignment_poll",
            "display_name": "Ballots",
            "projector_id": 1
          }`),
		"core/projection-default:13": []byte(`{
            "id": 13,
            "name": "motion_poll",
            "display_name": "Motion votes",
            "projector_id": 1
          }`),
		"core/projection-default:2": []byte(`{
            "id": 2,
            "name": "topics",
            "display_name": "Topics",
            "projector_id": 1
          }`),
		"core/projection-default:3": []byte(`{
            "id": 3,
            "name": "agenda_list_of_speakers",
            "display_name": "List of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:4": []byte(`{
            "id": 4,
            "name": "agenda_current_list_of_speakers",
            "display_name": "Current list of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:5": []byte(`{
            "id": 5,
            "name": "motions",
            "display_name": "Motions",
            "projector_id": 1
          }`),
		"core/projection-default:6": []byte(`{
            "id": 6,
            "name": "motionBlocks",
            "display_name": "Motion blocks",
            "projector_id": 1
          }`),
		"core/projection-default:7": []byte(`{
            "id": 7,
            "name": "assignments",
            "display_name": "Elections",
            "projector_id": 1
          }`),
		"core/projection-default:8": []byte(`{
            "id": 8,
            "name": "users",
            "display_name": "Participants",
            "projector_id": 1
          }`),
		"core/projection-default:9": []byte(`{
            "id": 9,
            "name": "mediafiles",
            "display_name": "Files",
            "projector_id": 1
          }`),
		"core/projector-message:1": []byte(`{
            "id": 1,
            "message": "<p>test</p>"
          }`),
		"core/projector:1": []byte(`{
            "id": 1,
            "elements": [
              {
                "name": "mediafiles/mediafile",
                "id": 3
              }
            ],
            "elements_preview": [],
            "elements_history": [
              [
                {
                  "name": "assignments/assignment",
                  "id": 1
                }
              ]
            ],
            "scale": 0,
            "scroll": 0,
            "name": "Default projector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [
              1,
              2,
              3,
              4,
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13
            ],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/projector:2": []byte(`{
            "id": 2,
            "elements": [
              {
                "name": "core/clock",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers",
                "stable": false
              },
              {
                "name": "agenda/current-speaker-chyron",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers-overlay",
                "stable": true
              }
            ],
            "elements_preview": [],
            "elements_history": [],
            "scale": 0,
            "scroll": 0,
            "name": "sideprojector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/tag:1": []byte(`{
            "id": 1,
            "name": "T1"
          }`),
		"core/tag:2": []byte(`{
            "id": 2,
            "name": "T2"
          }`),
		"mediafiles/mediafile:2": []byte(`{
            "id": 2,
            "title": "A.txt",
            "media_url_prefix": "/media/",
            "filesize": "< 1 kB",
            "mimetype": "text/plain",
            "pdf_information": {},
            "access_groups_id": [],
            "create_timestamp": "2020-08-11T11:16:07.013124+02:00",
            "is_directory": false,
            "path": "A.txt",
            "parent_id": null,
            "list_of_speakers_id": 5,
            "inherited_access_groups_id": true
          }`),
		"motions/category:1": []byte(`{
            "id": 1,
            "name": "first",
            "prefix": "A",
            "parent_id": null,
            "weight": 2,
            "level": 0
          }`),
		"motions/category:2": []byte(`{
            "id": 2,
            "name": "second",
            "prefix": "B",
            "parent_id": 1,
            "weight": 4,
            "level": 1
          }`),
		"motions/category:3": []byte(`{
            "id": 3,
            "name": "third",
            "prefix": "C",
            "parent_id": null,
            "weight": 6,
            "level": 0
          }`),
		"motions/motion-block:1": []byte(`{
            "id": 1,
            "title": "block",
            "agenda_item_id": 8,
            "list_of_speakers_id": 12,
            "internal": false,
            "motions_id": [
              1,
              2
            ]
          }`),
		"motions/motion-option:1": []byte(`{
            "id": 1,
            "yes": "0.000000",
            "no": "1.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-option:2": []byte(`{
            "id": 2,
            "yes": "1.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 2,
            "pollstate": 4
          }`),
		"motions/motion-poll:1": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "named",
            "title": "Abstimmung",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              1
            ],
            "id": 1,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-poll:2": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "pseudoanonymous",
            "title": "Abstimmung (2)",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              2
            ],
            "id": 2,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "N",
            "user_id": 5,
            "option_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-vote:2": []byte(`{
            "id": 2,
            "weight": "1.000000",
            "value": "Y",
            "user_id": null,
            "option_id": 2,
            "pollstate": 4
          }`),
		"motions/motion:2": []byte(`{
            "id": 2,
            "identifier": "Ä-1",
            "title": "Änderungsantrag zu Leadmotion1",
            "text": "",
            "amendment_paragraphs": [
              "<p>Lorem ipsum dolor sit amet, consectedfgddfgdf&nbsp; gdfgdfg dfg dfg ww ontes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>"
            ],
            "modified_final_version": "",
            "reason": "",
            "parent_id": 1,
            "category_id": 2,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 3,
                "user_id": 1,
                "motion_id": 2,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [],
            "attachments_id": [],
            "agenda_item_id": 6,
            "list_of_speakers_id": 9,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:39:35.025914+02:00",
            "last_modified": "2020-08-11T12:41:55.666495+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/motion:3": []byte(`{
            "id": 3,
            "identifier": "2",
            "title": "Public",
            "text": "<p>a</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>a</p>",
            "parent_id": null,
            "category_id": 1,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 2,
            "origin": "",
            "submitters": [
              {
                "id": 4,
                "user_id": 1,
                "motion_id": 3,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [
              1
            ],
            "attachments_id": [],
            "agenda_item_id": 7,
            "list_of_speakers_id": 10,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T12:24:45.289233+02:00",
            "last_modified": "2020-08-11T12:41:51.319986+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/state:1": []byte(`{
            "id": 1,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              2,
              3,
              4
            ],
            "workflow_id": 1
          }`),
		"motions/state:10": []byte(`{
            "id": 10,
            "name": "withdrawed",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:11": []byte(`{
            "id": 11,
            "name": "adjourned",
            "recommendation_label": "Adjournment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:12": []byte(`{
            "id": 12,
            "name": "not concerned",
            "recommendation_label": "No concernment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:13": []byte(`{
            "id": 13,
            "name": "refered to committee",
            "recommendation_label": "Referral to committee",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:14": []byte(`{
            "id": 14,
            "name": "needs review",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:15": []byte(`{
            "id": 15,
            "name": "rejected (not authorized)",
            "recommendation_label": "Rejection (not authorized)",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:2": []byte(`{
            "id": 2,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:3": []byte(`{
            "id": 3,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:4": []byte(`{
            "id": 4,
            "name": "not decided",
            "recommendation_label": "No decision",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:5": []byte(`{
            "id": 5,
            "name": "in progress",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [
              "is_submitter"
            ],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": true,
            "dont_set_identifier": true,
            "show_state_extension_field": true,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              6,
              10
            ],
            "workflow_id": 2
          }`),
		"motions/state:6": []byte(`{
            "id": 6,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": true,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              7,
              10,
              15
            ],
            "workflow_id": 2
          }`),
		"motions/state:7": []byte(`{
            "id": 7,
            "name": "permitted",
            "recommendation_label": "Permission",
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": true,
            "next_states_id": [
              8,
              9,
              10,
              11,
              12,
              13,
              14
            ],
            "workflow_id": 2
          }`),
		"motions/state:8": []byte(`{
            "id": 8,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:9": []byte(`{
            "id": 9,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/statute-paragraph:1": []byte(`{
            "id": 1,
            "title": "Statute",
            "text": "<p>test</p>",
            "weight": 10000
          }`),
		"motions/workflow:1": []byte(`{
            "id": 1,
            "name": "Simple Workflow",
            "states_id": [
              1,
              2,
              3,
              4
            ],
            "first_state_id": 1
          }`),
		"motions/workflow:2": []byte(`{
            "id": 2,
            "name": "Complex Workflow",
            "states_id": [
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13,
              14,
              15
            ],
            "first_state_id": 5
          }`),
		"topics/topic:1": []byte(`{
            "id": 1,
            "title": "Topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 1,
            "list_of_speakers_id": 1
          }`),
		"topics/topic:2": []byte(`{
            "id": 2,
            "title": "Hidden",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 2,
            "list_of_speakers_id": 2
          }`),
		"topics/topic:3": []byte(`{
            "id": 3,
            "title": "Internal",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 3,
            "list_of_speakers_id": 3
          }`),
		"topics/topic:4": []byte(`{
            "id": 4,
            "title": "Another public topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 9,
            "list_of_speakers_id": 14
          }`),
		"users/group:1": []byte(`{
            "id": 1,
            "name": "Default",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_see",
              "users.can_change_password"
            ]
          }`),
		"users/group:2": []byte(`{
            "id": 2,
            "name": "Admin",
            "permissions": []
          }`),
		"users/group:3": []byte(`{
            "id": 3,
            "name": "Delegates",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password"
            ]
          }`),
		"users/group:4": []byte(`{
            "id": 4,
            "name": "Staff",
            "permissions": [
              "agenda.can_manage",
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_manage_list_of_speakers",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_manage",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_history",
              "core.can_manage_projector",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "core.can_manage_tags",
              "mediafiles.can_manage",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_manage",
              "motions.can_manage_metadata",
              "motions.can_see",
              "motions.can_see_internal",
              "users.can_change_password",
              "users.can_manage",
              "users.can_see_extra_data",
              "users.can_see_name"
            ]
          }`),
		"users/group:5": []byte(`{
            "id": 5,
            "name": "Committees",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password",
              "users.can_see_name"
            ]
          }`),
		"users/user:1": []byte(`{
            "first_name": "",
            "username": "admin",
            "about_me": "",
            "id": 1,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              2
            ],
            "number": "",
            "last_name": "Administrator",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:2": []byte(`{
            "first_name": "candidate1",
            "username": "candidate1",
            "about_me": "",
            "id": 2,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:3": []byte(`{
            "first_name": "candidate2",
            "username": "candidate2",
            "about_me": "",
            "id": 3,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "groups_id": [],
            "title": "",
            "email": "",
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:4": []byte(`{
            "first_name": "a",
            "username": "a",
            "about_me": "",
            "id": 4,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:5": []byte(`{
            "first_name": "b",
            "username": "b",
            "about_me": "",
            "id": 5,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:6": []byte(`{
            "first_name": "speaker1",
            "username": "speaker1",
            "about_me": "",
            "id": 6,
            "structure_level": "layer X",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "title",
            "groups_id": [],
            "number": "3",
            "last_name": "the last name",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:7": []byte(`{
            "first_name": "speaker2",
            "username": "speaker2",
            "about_me": "",
            "id": 7,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
	},
	4: {
		"agenda/item:1": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            },
            "is_internal": false,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": null,
            "id": 1,
            "level": 0,
            "title_information": {
              "title": "Topic"
            },
            "closed": false,
            "weight": 2,
            "duration": null
          }`),
		"agenda/item:3": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": null,
            "id": 3,
            "level": 0,
            "title_information": {
              "title": "Internal"
            },
            "closed": false,
            "weight": 8,
            "duration": null
          }`),
		"agenda/item:5": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 5,
            "level": 1,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "closed": false,
            "weight": 14,
            "duration": null
          }`),
		"agenda/item:6": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 6,
            "level": 1,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "closed": false,
            "weight": 16,
            "duration": 0
          }`),
		"agenda/item:7": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": 6,
            "id": 7,
            "level": 2,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "closed": false,
            "weight": 18,
            "duration": null
          }`),
		"agenda/list-of-speakers:1": []byte(`{
            "id": 1,
            "title_information": {
              "title": "Topic"
            },
            "speakers": [
              {
                "id": 3,
                "user_id": 6,
                "begin_time": "2020-08-11T12:28:30.894574+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:10": []byte(`{
            "id": 10,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:11": []byte(`{
            "id": 11,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:12": []byte(`{
            "id": 12,
            "title_information": {
              "title": "block"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:13": []byte(`{
            "id": 13,
            "title_information": {
              "title": "block internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:14": []byte(`{
            "id": 14,
            "title_information": {
              "title": "Another public topic"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:2": []byte(`{
            "id": 2,
            "title_information": {
              "title": "Hidden"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:3": []byte(`{
            "id": 3,
            "title_information": {
              "title": "Internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:4": []byte(`{
            "id": 4,
            "title_information": {
              "title": "folder"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:5": []byte(`{
            "id": 5,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:6": []byte(`{
            "id": 6,
            "title_information": {
              "title": "in.jpg"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:7": []byte(`{
            "id": 7,
            "title_information": {
              "title": "Assignment"
            },
            "speakers": [
              {
                "id": 4,
                "user_id": 6,
                "begin_time": "2020-08-11T12:29:55.054553+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              },
              {
                "id": 6,
                "user_id": 7,
                "begin_time": null,
                "end_time": null,
                "weight": 2,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "assignments/assignment",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:8": []byte(`{
            "id": 8,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:9": []byte(`{
            "id": 9,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            }
          }`),
		"assignments/assignment-option:1": []byte(`{
            "user_id": 2,
            "weight": 1,
            "id": 1,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-option:2": []byte(`{
            "user_id": 3,
            "weight": 3,
            "id": 2,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-poll:1": []byte(`{
            "assignment_id": 1,
            "description": "",
            "pollmethod": "votes",
            "votes_amount": 1,
            "allow_multiple_votes_per_candidate": false,
            "global_no": true,
            "global_abstain": true,
            "state": 2,
            "type": "named",
            "title": "Wahlgang",
            "groups_id": [
              3
            ],
            "options_id": [
              1,
              2
            ],
            "id": 1,
            "onehundred_percent_base": "valid",
            "majority_method": "simple",
            "user_has_voted": true
          }`),
		"assignments/assignment-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "Y",
            "user_id": 4,
            "option_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment:1": []byte(`{
            "id": 1,
            "title": "Assignment",
            "description": "",
            "open_posts": 1,
            "phase": 0,
            "assignment_related_users": [
              {
                "id": 1,
                "user_id": 2,
                "weight": 1
              },
              {
                "id": 3,
                "user_id": 3,
                "weight": 3
              }
            ],
            "default_poll_description": "",
            "agenda_item_id": 4,
            "list_of_speakers_id": 7,
            "tags_id": [
              2
            ],
            "attachments_id": [],
            "number_poll_candidates": false,
            "polls_id": [
              1
            ]
          }`),
		"core/config:10": []byte(`{
            "id": 10,
            "key": "general_system_conference_show",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show live conference window",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 140,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:100": []byte(`{
            "id": 100,
            "key": "motion_poll_default_100_percent_base",
            "value": "YNA",
            "data": {
              "defaultValue": "YNA",
              "inputType": "choice",
              "label": "Default 100 % base of a voting result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 370,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:101": []byte(`{
            "id": 101,
            "key": "motion_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:102": []byte(`{
            "id": 102,
            "key": "motion_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 372,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:103": []byte(`{
            "id": 103,
            "key": "motions_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 373,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:104": []byte(`{
            "id": 104,
            "key": "motions_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 374,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:105": []byte(`{
            "id": 105,
            "key": "motions_export_title",
            "value": "Motions",
            "data": {
              "defaultValue": "Motions",
              "inputType": "string",
              "label": "Title for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 380,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:106": []byte(`{
            "id": 106,
            "key": "motions_export_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 382,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:107": []byte(`{
            "id": 107,
            "key": "motions_export_submitter_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show submitters and recommendation/state in table of contents",
              "helpText": "",
              "choices": null,
              "weight": 384,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:108": []byte(`{
            "id": 108,
            "key": "motions_export_follow_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show checkbox to record decision",
              "helpText": "",
              "choices": null,
              "weight": 386,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:109": []byte(`{
            "id": 109,
            "key": "assignment_poll_method",
            "value": "votes",
            "data": {
              "defaultValue": "votes",
              "inputType": "choice",
              "label": "Default election method",
              "helpText": "",
              "choices": [
                {
                  "value": "votes",
                  "display_name": "Yes per candidate"
                },
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                }
              ],
              "weight": 400,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:11": []byte(`{
            "id": 11,
            "key": "general_system_conference_auto_connect",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Connect all users to live conference automatically",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 141,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:110": []byte(`{
            "id": 110,
            "key": "assignment_poll_default_type",
            "value": "analog",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 403,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:111": []byte(`{
            "id": 111,
            "key": "assignment_poll_default_100_percent_base",
            "value": "valid",
            "data": {
              "defaultValue": "valid",
              "inputType": "choice",
              "label": "Default 100 % base of an election result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                },
                {
                  "value": "votes",
                  "display_name": "Sum of votes including general No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 405,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:112": []byte(`{
            "id": 112,
            "key": "assignment_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 410,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:113": []byte(`{
            "id": 113,
            "key": "assignment_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:114": []byte(`{
            "id": 114,
            "key": "assignment_poll_sort_poll_result_by_votes",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Sort election results by amount of votes",
              "helpText": "",
              "choices": null,
              "weight": 420,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:115": []byte(`{
            "id": 115,
            "key": "assignment_poll_add_candidates_to_list_of_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Put all candidates on the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 425,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:116": []byte(`{
            "id": 116,
            "key": "assignments_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 430,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:117": []byte(`{
            "id": 117,
            "key": "assignments_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 435,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:118": []byte(`{
            "id": 118,
            "key": "assignments_pdf_title",
            "value": "Elections",
            "data": {
              "defaultValue": "Elections",
              "inputType": "string",
              "label": "Title for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 460,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:119": []byte(`{
            "id": 119,
            "key": "assignments_pdf_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 470,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:12": []byte(`{
            "id": 12,
            "key": "general_system_conference_los_restriction",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow only current speakers and list of speakers managers to enter the live conference",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 142,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:13": []byte(`{
            "id": 13,
            "key": "general_system_stream_url",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Livestream url",
              "helpText": "Remove URL to deactivate livestream. Check extra group permission to see livestream.",
              "choices": null,
              "weight": 143,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:14": []byte(`{
            "id": 14,
            "key": "general_system_enable_anonymous",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow access for anonymous guest users",
              "helpText": "",
              "choices": null,
              "weight": 150,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:15": []byte(`{
            "id": 15,
            "key": "general_login_info_text",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Show this text on the login page",
              "helpText": "",
              "choices": null,
              "weight": 152,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:16": []byte(`{
            "id": 16,
            "key": "openslides_theme",
            "value": "openslides-default-light-theme",
            "data": {
              "defaultValue": "openslides-default-light-theme",
              "inputType": "choice",
              "label": "OpenSlides Theme",
              "helpText": "",
              "choices": [
                {
                  "value": "openslides-default-light-theme",
                  "display_name": "OpenSlides Default"
                },
                {
                  "value": "openslides-default-dark-theme",
                  "display_name": "OpenSlides Dark"
                },
                {
                  "value": "openslides-red-light-theme",
                  "display_name": "OpenSlides Red"
                },
                {
                  "value": "openslides-red-dark-theme",
                  "display_name": "OpenSlides Red Dark"
                },
                {
                  "value": "openslides-green-light-theme",
                  "display_name": "OpenSlides Green"
                },
                {
                  "value": "openslides-green-dark-theme",
                  "display_name": "OpenSlides Green Dark"
                },
                {
                  "value": "openslides-solarized-dark-theme",
                  "display_name": "OpenSlides Solarized"
                }
              ],
              "weight": 154,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:17": []byte(`{
            "id": 17,
            "key": "general_csv_separator",
            "value": ",",
            "data": {
              "defaultValue": ",",
              "inputType": "string",
              "label": "Separator used for all csv exports and examples",
              "helpText": "",
              "choices": null,
              "weight": 160,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:18": []byte(`{
            "id": 18,
            "key": "general_csv_encoding",
            "value": "utf-8",
            "data": {
              "defaultValue": "utf-8",
              "inputType": "choice",
              "label": "Default encoding for all csv exports",
              "helpText": "",
              "choices": [
                {
                  "value": "utf-8",
                  "display_name": "UTF-8"
                },
                {
                  "value": "iso-8859-15",
                  "display_name": "ISO-8859-15"
                }
              ],
              "weight": 162,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:19": []byte(`{
            "id": 19,
            "key": "general_export_pdf_pagenumber_alignment",
            "value": "center",
            "data": {
              "defaultValue": "center",
              "inputType": "choice",
              "label": "Page number alignment in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "left",
                  "display_name": "Left"
                },
                {
                  "value": "center",
                  "display_name": "Center"
                },
                {
                  "value": "right",
                  "display_name": "Right"
                }
              ],
              "weight": 164,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:2": []byte(`{
            "id": 2,
            "key": "general_event_name",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Event name",
              "helpText": "",
              "choices": null,
              "weight": 110,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:20": []byte(`{
            "id": 20,
            "key": "general_export_pdf_fontsize",
            "value": "10",
            "data": {
              "defaultValue": "10",
              "inputType": "choice",
              "label": "Standard font size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "10",
                  "display_name": "10"
                },
                {
                  "value": "11",
                  "display_name": "11"
                },
                {
                  "value": "12",
                  "display_name": "12"
                }
              ],
              "weight": 166,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:21": []byte(`{
            "id": 21,
            "key": "general_export_pdf_pagesize",
            "value": "A4",
            "data": {
              "defaultValue": "A4",
              "inputType": "choice",
              "label": "Standard page size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "A4",
                  "display_name": "DIN A4"
                },
                {
                  "value": "A5",
                  "display_name": "DIN A5"
                }
              ],
              "weight": 168,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:22": []byte(`{
            "id": 22,
            "key": "logos_available",
            "value": [
              "logo_projector_main",
              "logo_projector_header",
              "logo_web_header",
              "logo_pdf_header_L",
              "logo_pdf_header_R",
              "logo_pdf_footer_L",
              "logo_pdf_footer_R",
              "logo_pdf_ballot_paper"
            ],
            "data": null
          }`),
		"core/config:23": []byte(`{
            "id": 23,
            "key": "logo_projector_main",
            "value": {
              "display_name": "Projector logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:24": []byte(`{
            "id": 24,
            "key": "logo_projector_header",
            "value": {
              "display_name": "Projector header image",
              "path": ""
            },
            "data": null
          }`),
		"core/config:25": []byte(`{
            "id": 25,
            "key": "logo_web_header",
            "value": {
              "display_name": "Web interface header logo",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:26": []byte(`{
            "id": 26,
            "key": "logo_pdf_header_L",
            "value": {
              "display_name": "PDF header logo (left)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:27": []byte(`{
            "id": 27,
            "key": "logo_pdf_header_R",
            "value": {
              "display_name": "PDF header logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:28": []byte(`{
            "id": 28,
            "key": "logo_pdf_footer_L",
            "value": {
              "display_name": "PDF footer logo (left)",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:29": []byte(`{
            "id": 29,
            "key": "logo_pdf_footer_R",
            "value": {
              "display_name": "PDF footer logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:3": []byte(`{
            "id": 3,
            "key": "general_event_description",
            "value": "Presentation and assembly system",
            "data": {
              "defaultValue": "Presentation and assembly system",
              "inputType": "string",
              "label": "Short description of event",
              "helpText": "",
              "choices": null,
              "weight": 115,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:30": []byte(`{
            "id": 30,
            "key": "logo_pdf_ballot_paper",
            "value": {
              "display_name": "PDF ballot paper logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:31": []byte(`{
            "id": 31,
            "key": "fonts_available",
            "value": [
              "font_regular",
              "font_italic",
              "font_bold",
              "font_bold_italic",
              "font_monospace"
            ],
            "data": null
          }`),
		"core/config:32": []byte(`{
            "id": 32,
            "key": "font_regular",
            "value": {
              "display_name": "Font regular",
              "default": "assets/fonts/fira-sans-latin-400.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:33": []byte(`{
            "id": 33,
            "key": "font_italic",
            "value": {
              "display_name": "Font italic",
              "default": "assets/fonts/fira-sans-latin-400italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:34": []byte(`{
            "id": 34,
            "key": "font_bold",
            "value": {
              "display_name": "Font bold",
              "default": "assets/fonts/fira-sans-latin-500.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:35": []byte(`{
            "id": 35,
            "key": "font_bold_italic",
            "value": {
              "display_name": "Font bold italic",
              "default": "assets/fonts/fira-sans-latin-500italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:36": []byte(`{
            "id": 36,
            "key": "font_monospace",
            "value": {
              "display_name": "Font monospace",
              "default": "assets/fonts/roboto-condensed-bold.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:37": []byte(`{
            "id": 37,
            "key": "translations",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "translations",
              "label": "Custom translations",
              "helpText": "",
              "choices": null,
              "weight": 1000,
              "group": "Custom translations",
              "subgroup": "General"
            }
          }`),
		"core/config:38": []byte(`{
            "id": 38,
            "key": "config_version",
            "value": 2,
            "data": null
          }`),
		"core/config:39": []byte(`{
            "id": 39,
            "key": "db_id",
            "value": "2c3bd736c87a48a4be1f0dc707702144",
            "data": null
          }`),
		"core/config:4": []byte(`{
            "id": 4,
            "key": "general_event_date",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event date",
              "helpText": "",
              "choices": null,
              "weight": 120,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:40": []byte(`{
            "id": 40,
            "key": "users_sort_by",
            "value": "first_name",
            "data": {
              "defaultValue": "first_name",
              "inputType": "choice",
              "label": "Sort name of participants by",
              "helpText": "",
              "choices": [
                {
                  "value": "first_name",
                  "display_name": "Given name"
                },
                {
                  "value": "last_name",
                  "display_name": "Surname"
                },
                {
                  "value": "number",
                  "display_name": "Participant number"
                }
              ],
              "weight": 510,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:41": []byte(`{
            "id": 41,
            "key": "users_enable_presence_view",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Enable participant presence view",
              "helpText": "",
              "choices": null,
              "weight": 511,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:42": []byte(`{
            "id": 42,
            "key": "users_allow_self_set_present",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow users to set themselves as present",
              "helpText": "e.g. for online meetings",
              "choices": null,
              "weight": 512,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:43": []byte(`{
            "id": 43,
            "key": "users_activate_vote_weight",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate vote weight",
              "helpText": "",
              "choices": null,
              "weight": 513,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:44": []byte(`{
            "id": 44,
            "key": "users_pdf_welcometitle",
            "value": "Welcome to OpenSlides",
            "data": {
              "defaultValue": "Welcome to OpenSlides",
              "inputType": "string",
              "label": "Title for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 520,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:45": []byte(`{
            "id": 45,
            "key": "users_pdf_welcometext",
            "value": "[Place for your welcome and help text.]",
            "data": {
              "defaultValue": "[Place for your welcome and help text.]",
              "inputType": "string",
              "label": "Help text for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 530,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:46": []byte(`{
            "id": 46,
            "key": "users_pdf_url",
            "value": "http://example.com:8000",
            "data": {
              "defaultValue": "http://example.com:8000",
              "inputType": "string",
              "label": "System URL",
              "helpText": "Used for QRCode in PDF of access data.",
              "choices": null,
              "weight": 540,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:47": []byte(`{
            "id": 47,
            "key": "users_pdf_wlan_ssid",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN name (SSID)",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 550,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:48": []byte(`{
            "id": 48,
            "key": "users_pdf_wlan_password",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN password",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 560,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:49": []byte(`{
            "id": 49,
            "key": "users_pdf_wlan_encryption",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "choice",
              "label": "WLAN encryption",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": [
                {
                  "value": "",
                  "display_name": "---------"
                },
                {
                  "value": "WEP",
                  "display_name": "WEP"
                },
                {
                  "value": "WPA",
                  "display_name": "WPA/WPA2"
                },
                {
                  "value": "nopass",
                  "display_name": "No encryption"
                }
              ],
              "weight": 570,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:5": []byte(`{
            "id": 5,
            "key": "general_event_location",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event location",
              "helpText": "",
              "choices": null,
              "weight": 125,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:50": []byte(`{
            "id": 50,
            "key": "users_email_sender",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Sender name",
              "helpText": "The sender address is defined in the OpenSlides server settings and should modified by administrator only.",
              "choices": null,
              "weight": 600,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:51": []byte(`{
            "id": 51,
            "key": "users_email_replyto",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Reply address",
              "helpText": "",
              "choices": null,
              "weight": 601,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:52": []byte(`{
            "id": 52,
            "key": "users_email_subject",
            "value": "OpenSlides access data",
            "data": {
              "defaultValue": "OpenSlides access data",
              "inputType": "string",
              "label": "Email subject",
              "helpText": "You can use {event_name} and {username} as placeholder.",
              "choices": null,
              "weight": 605,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:53": []byte(`{
            "id": 53,
            "key": "users_email_body",
            "value": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
            "data": {
              "defaultValue": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
              "inputType": "text",
              "label": "Email body",
              "helpText": "Use these placeholders: {name}, {event_name}, {url}, {username}, {password}. The url referrs to the system url.",
              "choices": null,
              "weight": 610,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:54": []byte(`{
            "id": 54,
            "key": "agenda_start_event_date_time",
            "value": null,
            "data": {
              "defaultValue": null,
              "inputType": "datetimepicker",
              "label": "Begin of event",
              "helpText": "Input format: DD.MM.YYYY HH:MM",
              "choices": null,
              "weight": 200,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:55": []byte(`{
            "id": 55,
            "key": "agenda_show_subtitle",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show subtitles in the agenda",
              "helpText": "",
              "choices": null,
              "weight": 201,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:56": []byte(`{
            "id": 56,
            "key": "agenda_enable_numbering",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Enable numbering for agenda items",
              "helpText": "",
              "choices": null,
              "weight": 205,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:57": []byte(`{
            "id": 57,
            "key": "agenda_number_prefix",
            "value": "TOP",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Numbering prefix for agenda items",
              "helpText": "This prefix will be set if you run the automatic agenda numbering.",
              "choices": null,
              "weight": 206,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:58": []byte(`{
            "id": 58,
            "key": "agenda_numeral_system",
            "value": "arabic",
            "data": {
              "defaultValue": "arabic",
              "inputType": "choice",
              "label": "Numeral system for agenda items",
              "helpText": "",
              "choices": [
                {
                  "value": "arabic",
                  "display_name": "Arabic"
                },
                {
                  "value": "roman",
                  "display_name": "Roman"
                }
              ],
              "weight": 207,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:59": []byte(`{
            "id": 59,
            "key": "agenda_item_creation",
            "value": "default_yes",
            "data": {
              "defaultValue": "default_yes",
              "inputType": "choice",
              "label": "Add to agenda",
              "helpText": "",
              "choices": [
                {
                  "value": "always",
                  "display_name": "Always"
                },
                {
                  "value": "never",
                  "display_name": "Never"
                },
                {
                  "value": "default_yes",
                  "display_name": "Ask, default yes"
                },
                {
                  "value": "default_no",
                  "display_name": "Ask, default no"
                }
              ],
              "weight": 210,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:6": []byte(`{
            "id": 6,
            "key": "general_event_legal_notice",
            "value": "<a href=\"http://www.openslides.org\">OpenSlides</a> is a free web based presentation and assembly system for visualizing and controlling agenda, motions and elections of an assembly.",
            "data": null
          }`),
		"core/config:60": []byte(`{
            "id": 60,
            "key": "agenda_new_items_default_visibility",
            "value": "2",
            "data": {
              "defaultValue": "2",
              "inputType": "choice",
              "label": "Default visibility for new agenda items (except topics)",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Public item"
                },
                {
                  "value": "2",
                  "display_name": "Internal item"
                },
                {
                  "value": "3",
                  "display_name": "Hidden item"
                }
              ],
              "weight": 211,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:61": []byte(`{
            "id": 61,
            "key": "agenda_hide_internal_items_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Hide internal items when projecting subitems",
              "helpText": "",
              "choices": null,
              "weight": 212,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:62": []byte(`{
            "id": 62,
            "key": "agenda_show_last_speakers",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of last speakers to be shown on the projector",
              "helpText": "",
              "choices": null,
              "weight": 220,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:63": []byte(`{
            "id": 63,
            "key": "agenda_show_next_speakers",
            "value": -1,
            "data": {
              "defaultValue": -1,
              "inputType": "integer",
              "label": "Number of the next speakers to be shown on the projector",
              "helpText": "Enter number of the next shown speakers. Choose -1 to show all next speakers.",
              "choices": null,
              "weight": 222,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:64": []byte(`{
            "id": 64,
            "key": "agenda_countdown_warning_time",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Show orange countdown in the last x seconds of speaking time",
              "helpText": "Enter duration in seconds. Choose 0 to disable warning color.",
              "choices": null,
              "weight": 224,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:65": []byte(`{
            "id": 65,
            "key": "projector_default_countdown",
            "value": 60,
            "data": {
              "defaultValue": 60,
              "inputType": "integer",
              "label": "Predefined seconds of new countdowns",
              "helpText": "",
              "choices": null,
              "weight": 226,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:66": []byte(`{
            "id": 66,
            "key": "agenda_couple_countdown_and_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Couple countdown with the list of speakers",
              "helpText": "[Begin speech] starts the countdown, [End speech] stops the countdown.",
              "choices": null,
              "weight": 228,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:67": []byte(`{
            "id": 67,
            "key": "agenda_hide_amount_of_speakers",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide the amount of speakers in subtitle of list of speakers slide",
              "helpText": "",
              "choices": null,
              "weight": 230,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:68": []byte(`{
            "id": 68,
            "key": "agenda_present_speakers_only",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Only present participants can be added to the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 232,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:69": []byte(`{
            "id": 69,
            "key": "agenda_show_first_contribution",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show hint »first speech« in the list of speakers management view",
              "helpText": "",
              "choices": null,
              "weight": 234,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:7": []byte(`{
            "id": 7,
            "key": "general_event_privacy_policy",
            "value": "",
            "data": null
          }`),
		"core/config:70": []byte(`{
            "id": 70,
            "key": "motions_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new motions",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 310,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:71": []byte(`{
            "id": 71,
            "key": "motions_statute_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new statute amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 312,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:72": []byte(`{
            "id": 72,
            "key": "motions_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 314,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:73": []byte(`{
            "id": 73,
            "key": "motions_preamble",
            "value": "The assembly may decide:",
            "data": {
              "defaultValue": "The assembly may decide:",
              "inputType": "string",
              "label": "Motion preamble",
              "helpText": "",
              "choices": null,
              "weight": 320,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:74": []byte(`{
            "id": 74,
            "key": "motions_default_line_numbering",
            "value": "outside",
            "data": {
              "defaultValue": "outside",
              "inputType": "choice",
              "label": "Default line numbering",
              "helpText": "",
              "choices": [
                {
                  "value": "outside",
                  "display_name": "outside"
                },
                {
                  "value": "inline",
                  "display_name": "inline"
                },
                {
                  "value": "none",
                  "display_name": "Disabled"
                }
              ],
              "weight": 322,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:75": []byte(`{
            "id": 75,
            "key": "motions_line_length",
            "value": 85,
            "data": {
              "defaultValue": 85,
              "inputType": "integer",
              "label": "Line length",
              "helpText": "The maximum number of characters per line. Relevant when line numbering is enabled. Min: 40",
              "choices": null,
              "weight": 323,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:76": []byte(`{
            "id": 76,
            "key": "motions_reason_required",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Reason required for creating new motion",
              "helpText": "",
              "choices": null,
              "weight": 324,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:77": []byte(`{
            "id": 77,
            "key": "motions_disable_text_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide motion text on projector",
              "helpText": "",
              "choices": null,
              "weight": 325,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:78": []byte(`{
            "id": 78,
            "key": "motions_disable_reason_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide reason on projector",
              "helpText": "",
              "choices": null,
              "weight": 326,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:79": []byte(`{
            "id": 79,
            "key": "motions_disable_recommendation_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide recommendation on projector",
              "helpText": "",
              "choices": null,
              "weight": 327,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:8": []byte(`{
            "id": 8,
            "key": "general_event_welcome_title",
            "value": "Welcome to OpenSlides",
            "data": null
          }`),
		"core/config:80": []byte(`{
            "id": 80,
            "key": "motions_hide_referring_motions",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide referring motions",
              "helpText": "",
              "choices": null,
              "weight": 328,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:81": []byte(`{
            "id": 81,
            "key": "motions_disable_sidebox_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show meta information box below the title on projector",
              "helpText": "",
              "choices": null,
              "weight": 329,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:82": []byte(`{
            "id": 82,
            "key": "motions_show_sequential_numbers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show the sequential number for a motion",
              "helpText": "In motion list, motion detail and PDF.",
              "choices": null,
              "weight": 330,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:83": []byte(`{
            "id": 83,
            "key": "motions_recommendations_by",
            "value": "ABK",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender",
              "helpText": "Will be displayed as label before selected recommendation. Use an empty value to disable the recommendation system.",
              "choices": null,
              "weight": 332,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:84": []byte(`{
            "id": 84,
            "key": "motions_statute_recommendations_by",
            "value": "ABK2",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender for statute amendments",
              "helpText": "Will be displayed as label before selected recommendation in statute amendments.",
              "choices": null,
              "weight": 333,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:85": []byte(`{
            "id": 85,
            "key": "motions_recommendation_text_mode",
            "value": "diff",
            "data": {
              "defaultValue": "diff",
              "inputType": "choice",
              "label": "Default text version for change recommendations",
              "helpText": "",
              "choices": [
                {
                  "value": "original",
                  "display_name": "Original version"
                },
                {
                  "value": "changed",
                  "display_name": "Changed version"
                },
                {
                  "value": "diff",
                  "display_name": "Diff version"
                },
                {
                  "value": "agreed",
                  "display_name": "Final version"
                }
              ],
              "weight": 334,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:86": []byte(`{
            "id": 86,
            "key": "motions_motions_sorting",
            "value": "identifier",
            "data": {
              "defaultValue": "identifier",
              "inputType": "choice",
              "label": "Sort motions by",
              "helpText": "",
              "choices": [
                {
                  "value": "weight",
                  "display_name": "Call list"
                },
                {
                  "value": "identifier",
                  "display_name": "Identifier"
                }
              ],
              "weight": 335,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:87": []byte(`{
            "id": 87,
            "key": "motions_identifier",
            "value": "per_category",
            "data": {
              "defaultValue": "per_category",
              "inputType": "choice",
              "label": "Identifier",
              "helpText": "",
              "choices": [
                {
                  "value": "per_category",
                  "display_name": "Numbered per category"
                },
                {
                  "value": "serially_numbered",
                  "display_name": "Serially numbered"
                },
                {
                  "value": "manually",
                  "display_name": "Set it manually"
                }
              ],
              "weight": 340,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:88": []byte(`{
            "id": 88,
            "key": "motions_identifier_min_digits",
            "value": 1,
            "data": {
              "defaultValue": 1,
              "inputType": "integer",
              "label": "Number of minimal digits for identifier",
              "helpText": "Uses leading zeros to sort motions correctly by identifier.",
              "choices": null,
              "weight": 342,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:89": []byte(`{
            "id": 89,
            "key": "motions_identifier_with_blank",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow blank in identifier",
              "helpText": "Blank between prefix and number, e.g. 'A 001'.",
              "choices": null,
              "weight": 344,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:9": []byte(`{
            "id": 9,
            "key": "general_event_welcome_text",
            "value": "[Space for your welcome text.]",
            "data": null
          }`),
		"core/config:90": []byte(`{
            "id": 90,
            "key": "motions_statutes_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate statute amendments",
              "helpText": "",
              "choices": null,
              "weight": 350,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:91": []byte(`{
            "id": 91,
            "key": "motions_amendments_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate amendments",
              "helpText": "",
              "choices": null,
              "weight": 351,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:92": []byte(`{
            "id": 92,
            "key": "motions_amendments_main_table",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show amendments together with motions",
              "helpText": "",
              "choices": null,
              "weight": 352,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:93": []byte(`{
            "id": 93,
            "key": "motions_amendments_prefix",
            "value": "Ä-",
            "data": {
              "defaultValue": "-",
              "inputType": "string",
              "label": "Prefix for the identifier for amendments",
              "helpText": "",
              "choices": null,
              "weight": 353,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:94": []byte(`{
            "id": 94,
            "key": "motions_amendments_text_mode",
            "value": "paragraph",
            "data": {
              "defaultValue": "paragraph",
              "inputType": "choice",
              "label": "How to create new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "freestyle",
                  "display_name": "Empty text field"
                },
                {
                  "value": "fulltext",
                  "display_name": "Edit the whole motion text"
                },
                {
                  "value": "paragraph",
                  "display_name": "Paragraph-based, Diff-enabled"
                }
              ],
              "weight": 354,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:95": []byte(`{
            "id": 95,
            "key": "motions_amendments_multiple_paragraphs",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Amendments can change multiple paragraphs",
              "helpText": "",
              "choices": null,
              "weight": 355,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:96": []byte(`{
            "id": 96,
            "key": "motions_amendments_of_amendments",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow amendments of amendments",
              "helpText": "",
              "choices": null,
              "weight": 356,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:97": []byte(`{
            "id": 97,
            "key": "motions_min_supporters",
            "value": 1,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of (minimum) required supporters for a motion",
              "helpText": "Choose 0 to disable the supporting system.",
              "choices": null,
              "weight": 360,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:98": []byte(`{
            "id": 98,
            "key": "motions_remove_supporters",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Remove all supporters of a motion if a submitter edits his motion in early state",
              "helpText": "",
              "choices": null,
              "weight": 361,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:99": []byte(`{
            "id": 99,
            "key": "motion_poll_default_type",
            "value": "named",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 367,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/countdown:1": []byte(`{
            "id": 1,
            "title": "Default countdown",
            "description": "",
            "default_time": 60,
            "countdown_time": 1597141855.090748,
            "running": true
          }`),
		"core/projection-default:1": []byte(`{
            "id": 1,
            "name": "agenda_all_items",
            "display_name": "Agenda",
            "projector_id": 1
          }`),
		"core/projection-default:10": []byte(`{
            "id": 10,
            "name": "messages",
            "display_name": "Messages",
            "projector_id": 1
          }`),
		"core/projection-default:11": []byte(`{
            "id": 11,
            "name": "countdowns",
            "display_name": "Countdowns",
            "projector_id": 1
          }`),
		"core/projection-default:12": []byte(`{
            "id": 12,
            "name": "assignment_poll",
            "display_name": "Ballots",
            "projector_id": 1
          }`),
		"core/projection-default:13": []byte(`{
            "id": 13,
            "name": "motion_poll",
            "display_name": "Motion votes",
            "projector_id": 1
          }`),
		"core/projection-default:2": []byte(`{
            "id": 2,
            "name": "topics",
            "display_name": "Topics",
            "projector_id": 1
          }`),
		"core/projection-default:3": []byte(`{
            "id": 3,
            "name": "agenda_list_of_speakers",
            "display_name": "List of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:4": []byte(`{
            "id": 4,
            "name": "agenda_current_list_of_speakers",
            "display_name": "Current list of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:5": []byte(`{
            "id": 5,
            "name": "motions",
            "display_name": "Motions",
            "projector_id": 1
          }`),
		"core/projection-default:6": []byte(`{
            "id": 6,
            "name": "motionBlocks",
            "display_name": "Motion blocks",
            "projector_id": 1
          }`),
		"core/projection-default:7": []byte(`{
            "id": 7,
            "name": "assignments",
            "display_name": "Elections",
            "projector_id": 1
          }`),
		"core/projection-default:8": []byte(`{
            "id": 8,
            "name": "users",
            "display_name": "Participants",
            "projector_id": 1
          }`),
		"core/projection-default:9": []byte(`{
            "id": 9,
            "name": "mediafiles",
            "display_name": "Files",
            "projector_id": 1
          }`),
		"core/projector-message:1": []byte(`{
            "id": 1,
            "message": "<p>test</p>"
          }`),
		"core/projector:1": []byte(`{
            "id": 1,
            "elements": [
              {
                "name": "mediafiles/mediafile",
                "id": 3
              }
            ],
            "elements_preview": [],
            "elements_history": [
              [
                {
                  "name": "assignments/assignment",
                  "id": 1
                }
              ]
            ],
            "scale": 0,
            "scroll": 0,
            "name": "Default projector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [
              1,
              2,
              3,
              4,
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13
            ],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/projector:2": []byte(`{
            "id": 2,
            "elements": [
              {
                "name": "core/clock",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers",
                "stable": false
              },
              {
                "name": "agenda/current-speaker-chyron",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers-overlay",
                "stable": true
              }
            ],
            "elements_preview": [],
            "elements_history": [],
            "scale": 0,
            "scroll": 0,
            "name": "sideprojector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/tag:1": []byte(`{
            "id": 1,
            "name": "T1"
          }`),
		"core/tag:2": []byte(`{
            "id": 2,
            "name": "T2"
          }`),
		"mediafiles/mediafile:1": []byte(`{
            "id": 1,
            "title": "folder",
            "media_url_prefix": "/media/",
            "filesize": "",
            "mimetype": "",
            "pdf_information": {},
            "access_groups_id": [
              3
            ],
            "create_timestamp": "2020-08-11T11:15:50.719490+02:00",
            "is_directory": true,
            "path": "folder/",
            "parent_id": null,
            "list_of_speakers_id": 4,
            "inherited_access_groups_id": [
              3
            ]
          }`),
		"mediafiles/mediafile:2": []byte(`{
            "id": 2,
            "title": "A.txt",
            "media_url_prefix": "/media/",
            "filesize": "< 1 kB",
            "mimetype": "text/plain",
            "pdf_information": {},
            "access_groups_id": [],
            "create_timestamp": "2020-08-11T11:16:07.013124+02:00",
            "is_directory": false,
            "path": "A.txt",
            "parent_id": null,
            "list_of_speakers_id": 5,
            "inherited_access_groups_id": true
          }`),
		"motions/category:1": []byte(`{
            "id": 1,
            "name": "first",
            "prefix": "A",
            "parent_id": null,
            "weight": 2,
            "level": 0
          }`),
		"motions/category:2": []byte(`{
            "id": 2,
            "name": "second",
            "prefix": "B",
            "parent_id": 1,
            "weight": 4,
            "level": 1
          }`),
		"motions/category:3": []byte(`{
            "id": 3,
            "name": "third",
            "prefix": "C",
            "parent_id": null,
            "weight": 6,
            "level": 0
          }`),
		"motions/motion-block:1": []byte(`{
            "id": 1,
            "title": "block",
            "agenda_item_id": 8,
            "list_of_speakers_id": 12,
            "internal": false,
            "motions_id": [
              1,
              2
            ]
          }`),
		"motions/motion-comment-section:1": []byte(`{
            "id": 1,
            "name": "public",
            "read_groups_id": [
              2,
              3
            ],
            "write_groups_id": [
              2,
              3
            ],
            "weight": 10000
          }`),
		"motions/motion-option:1": []byte(`{
            "id": 1,
            "yes": "0.000000",
            "no": "1.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-option:2": []byte(`{
            "id": 2,
            "yes": "1.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 2,
            "pollstate": 4
          }`),
		"motions/motion-poll:1": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "named",
            "title": "Abstimmung",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              1
            ],
            "id": 1,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-poll:2": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "pseudoanonymous",
            "title": "Abstimmung (2)",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              2
            ],
            "id": 2,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "N",
            "user_id": 5,
            "option_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-vote:2": []byte(`{
            "id": 2,
            "weight": "1.000000",
            "value": "Y",
            "user_id": null,
            "option_id": 2,
            "pollstate": 4
          }`),
		"motions/motion:1": []byte(`{
            "id": 1,
            "identifier": null,
            "title": "Leadmotion1",
            "text": "<p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>",
            "parent_id": null,
            "category_id": 3,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 2,
                "user_id": 4,
                "motion_id": 1,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 5,
            "state_extension": "Maybe",
            "state_restriction": [
              "is_submitter"
            ],
            "statute_paragraph_id": null,
            "workflow_id": 2,
            "recommendation_id": 7,
            "recommendation_extension": "if [motion:2] is acepted",
            "tags_id": [],
            "attachments_id": [
              2
            ],
            "agenda_item_id": 5,
            "list_of_speakers_id": 8,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:36:06.865840+02:00",
            "last_modified": "2020-08-11T12:42:00.402425+02:00",
            "change_recommendations_id": [
              1
            ],
            "amendments_id": [
              2
            ]
          }`),
		"motions/motion:2": []byte(`{
            "id": 2,
            "identifier": "Ä-1",
            "title": "Änderungsantrag zu Leadmotion1",
            "text": "",
            "amendment_paragraphs": [
              "<p>Lorem ipsum dolor sit amet, consectedfgddfgdf&nbsp; gdfgdfg dfg dfg ww ontes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>"
            ],
            "modified_final_version": "",
            "reason": "",
            "parent_id": 1,
            "category_id": 2,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 3,
                "user_id": 1,
                "motion_id": 2,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [],
            "attachments_id": [],
            "agenda_item_id": 6,
            "list_of_speakers_id": 9,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:39:35.025914+02:00",
            "last_modified": "2020-08-11T12:41:55.666495+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/motion:3": []byte(`{
            "id": 3,
            "identifier": "2",
            "title": "Public",
            "text": "<p>a</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>a</p>",
            "parent_id": null,
            "category_id": 1,
            "category_weight": 10000,
            "comments": [
              {
                "id": 1,
                "comment": "<p>test</p>",
                "section_id": 1,
                "read_groups_id": [
                  2,
                  3
                ]
              }
            ],
            "motion_block_id": 2,
            "origin": "",
            "submitters": [
              {
                "id": 4,
                "user_id": 1,
                "motion_id": 3,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [
              1
            ],
            "attachments_id": [],
            "agenda_item_id": 7,
            "list_of_speakers_id": 10,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T12:24:45.289233+02:00",
            "last_modified": "2020-08-11T12:41:51.319986+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/state:1": []byte(`{
            "id": 1,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              2,
              3,
              4
            ],
            "workflow_id": 1
          }`),
		"motions/state:10": []byte(`{
            "id": 10,
            "name": "withdrawed",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:11": []byte(`{
            "id": 11,
            "name": "adjourned",
            "recommendation_label": "Adjournment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:12": []byte(`{
            "id": 12,
            "name": "not concerned",
            "recommendation_label": "No concernment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:13": []byte(`{
            "id": 13,
            "name": "refered to committee",
            "recommendation_label": "Referral to committee",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:14": []byte(`{
            "id": 14,
            "name": "needs review",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:15": []byte(`{
            "id": 15,
            "name": "rejected (not authorized)",
            "recommendation_label": "Rejection (not authorized)",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:2": []byte(`{
            "id": 2,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:3": []byte(`{
            "id": 3,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:4": []byte(`{
            "id": 4,
            "name": "not decided",
            "recommendation_label": "No decision",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:5": []byte(`{
            "id": 5,
            "name": "in progress",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [
              "is_submitter"
            ],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": true,
            "dont_set_identifier": true,
            "show_state_extension_field": true,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              6,
              10
            ],
            "workflow_id": 2
          }`),
		"motions/state:6": []byte(`{
            "id": 6,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": true,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              7,
              10,
              15
            ],
            "workflow_id": 2
          }`),
		"motions/state:7": []byte(`{
            "id": 7,
            "name": "permitted",
            "recommendation_label": "Permission",
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": true,
            "next_states_id": [
              8,
              9,
              10,
              11,
              12,
              13,
              14
            ],
            "workflow_id": 2
          }`),
		"motions/state:8": []byte(`{
            "id": 8,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:9": []byte(`{
            "id": 9,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/statute-paragraph:1": []byte(`{
            "id": 1,
            "title": "Statute",
            "text": "<p>test</p>",
            "weight": 10000
          }`),
		"motions/workflow:1": []byte(`{
            "id": 1,
            "name": "Simple Workflow",
            "states_id": [
              1,
              2,
              3,
              4
            ],
            "first_state_id": 1
          }`),
		"motions/workflow:2": []byte(`{
            "id": 2,
            "name": "Complex Workflow",
            "states_id": [
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13,
              14,
              15
            ],
            "first_state_id": 5
          }`),
		"topics/topic:1": []byte(`{
            "id": 1,
            "title": "Topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 1,
            "list_of_speakers_id": 1
          }`),
		"topics/topic:2": []byte(`{
            "id": 2,
            "title": "Hidden",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 2,
            "list_of_speakers_id": 2
          }`),
		"topics/topic:3": []byte(`{
            "id": 3,
            "title": "Internal",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 3,
            "list_of_speakers_id": 3
          }`),
		"topics/topic:4": []byte(`{
            "id": 4,
            "title": "Another public topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 9,
            "list_of_speakers_id": 14
          }`),
		"users/group:1": []byte(`{
            "id": 1,
            "name": "Default",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_see",
              "users.can_change_password"
            ]
          }`),
		"users/group:2": []byte(`{
            "id": 2,
            "name": "Admin",
            "permissions": []
          }`),
		"users/group:3": []byte(`{
            "id": 3,
            "name": "Delegates",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password"
            ]
          }`),
		"users/group:4": []byte(`{
            "id": 4,
            "name": "Staff",
            "permissions": [
              "agenda.can_manage",
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_manage_list_of_speakers",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_manage",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_history",
              "core.can_manage_projector",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "core.can_manage_tags",
              "mediafiles.can_manage",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_manage",
              "motions.can_manage_metadata",
              "motions.can_see",
              "motions.can_see_internal",
              "users.can_change_password",
              "users.can_manage",
              "users.can_see_extra_data",
              "users.can_see_name"
            ]
          }`),
		"users/group:5": []byte(`{
            "id": 5,
            "name": "Committees",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password",
              "users.can_see_name"
            ]
          }`),
		"users/personal-note:2": []byte(`{
            "id": 2,
            "user_id": 4,
            "notes": {
              "motions/motion": {
                "3": {
                  "note": "<p>user a</p>",
                  "star": false
                }
              }
            }
          }`),
		"users/user:1": []byte(`{
            "first_name": "",
            "username": "admin",
            "about_me": "",
            "id": 1,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              2
            ],
            "number": "",
            "last_name": "Administrator",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:2": []byte(`{
            "first_name": "candidate1",
            "username": "candidate1",
            "about_me": "",
            "id": 2,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:3": []byte(`{
            "first_name": "candidate2",
            "username": "candidate2",
            "about_me": "",
            "id": 3,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:4": []byte(`{
            "first_name": "a",
            "username": "a",
            "about_me": "",
            "id": 4,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "groups_id": [
              3
            ],
            "title": "",
            "email": "",
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:5": []byte(`{
            "first_name": "b",
            "username": "b",
            "about_me": "",
            "id": 5,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:6": []byte(`{
            "first_name": "speaker1",
            "username": "speaker1",
            "about_me": "",
            "id": 6,
            "structure_level": "layer X",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "title",
            "groups_id": [],
            "number": "3",
            "last_name": "the last name",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:7": []byte(`{
            "first_name": "speaker2",
            "username": "speaker2",
            "about_me": "",
            "id": 7,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
	},
	5: {
		"agenda/item:1": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            },
            "is_internal": false,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": null,
            "id": 1,
            "level": 0,
            "title_information": {
              "title": "Topic"
            },
            "closed": false,
            "weight": 2,
            "duration": null
          }`),
		"agenda/item:3": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": null,
            "id": 3,
            "level": 0,
            "title_information": {
              "title": "Internal"
            },
            "closed": false,
            "weight": 8,
            "duration": null
          }`),
		"agenda/item:5": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 5,
            "level": 1,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "closed": false,
            "weight": 14,
            "duration": null
          }`),
		"agenda/item:6": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 6,
            "level": 1,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "closed": false,
            "weight": 16,
            "duration": 0
          }`),
		"agenda/item:7": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": 6,
            "id": 7,
            "level": 2,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "closed": false,
            "weight": 18,
            "duration": null
          }`),
		"agenda/list-of-speakers:1": []byte(`{
            "id": 1,
            "title_information": {
              "title": "Topic"
            },
            "speakers": [
              {
                "id": 3,
                "user_id": 6,
                "begin_time": "2020-08-11T12:28:30.894574+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:10": []byte(`{
            "id": 10,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:11": []byte(`{
            "id": 11,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:12": []byte(`{
            "id": 12,
            "title_information": {
              "title": "block"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:13": []byte(`{
            "id": 13,
            "title_information": {
              "title": "block internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:14": []byte(`{
            "id": 14,
            "title_information": {
              "title": "Another public topic"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:2": []byte(`{
            "id": 2,
            "title_information": {
              "title": "Hidden"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:3": []byte(`{
            "id": 3,
            "title_information": {
              "title": "Internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:4": []byte(`{
            "id": 4,
            "title_information": {
              "title": "folder"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:5": []byte(`{
            "id": 5,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:6": []byte(`{
            "id": 6,
            "title_information": {
              "title": "in.jpg"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:7": []byte(`{
            "id": 7,
            "title_information": {
              "title": "Assignment"
            },
            "speakers": [
              {
                "id": 4,
                "user_id": 6,
                "begin_time": "2020-08-11T12:29:55.054553+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              },
              {
                "id": 6,
                "user_id": 7,
                "begin_time": null,
                "end_time": null,
                "weight": 2,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "assignments/assignment",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:8": []byte(`{
            "id": 8,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:9": []byte(`{
            "id": 9,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            }
          }`),
		"assignments/assignment-option:1": []byte(`{
            "user_id": 2,
            "weight": 1,
            "id": 1,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-option:2": []byte(`{
            "user_id": 3,
            "weight": 3,
            "id": 2,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-poll:1": []byte(`{
            "assignment_id": 1,
            "description": "",
            "pollmethod": "votes",
            "votes_amount": 1,
            "allow_multiple_votes_per_candidate": false,
            "global_no": true,
            "global_abstain": true,
            "state": 2,
            "type": "named",
            "title": "Wahlgang",
            "groups_id": [
              3
            ],
            "options_id": [
              1,
              2
            ],
            "id": 1,
            "onehundred_percent_base": "valid",
            "majority_method": "simple",
            "user_has_voted": false
          }`),
		"assignments/assignment:1": []byte(`{
            "id": 1,
            "title": "Assignment",
            "description": "",
            "open_posts": 1,
            "phase": 0,
            "assignment_related_users": [
              {
                "id": 1,
                "user_id": 2,
                "weight": 1
              },
              {
                "id": 3,
                "user_id": 3,
                "weight": 3
              }
            ],
            "default_poll_description": "",
            "agenda_item_id": 4,
            "list_of_speakers_id": 7,
            "tags_id": [
              2
            ],
            "attachments_id": [],
            "number_poll_candidates": false,
            "polls_id": [
              1
            ]
          }`),
		"core/config:10": []byte(`{
            "id": 10,
            "key": "general_system_conference_show",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show live conference window",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 140,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:100": []byte(`{
            "id": 100,
            "key": "motion_poll_default_100_percent_base",
            "value": "YNA",
            "data": {
              "defaultValue": "YNA",
              "inputType": "choice",
              "label": "Default 100 % base of a voting result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 370,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:101": []byte(`{
            "id": 101,
            "key": "motion_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:102": []byte(`{
            "id": 102,
            "key": "motion_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 372,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:103": []byte(`{
            "id": 103,
            "key": "motions_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 373,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:104": []byte(`{
            "id": 104,
            "key": "motions_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 374,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:105": []byte(`{
            "id": 105,
            "key": "motions_export_title",
            "value": "Motions",
            "data": {
              "defaultValue": "Motions",
              "inputType": "string",
              "label": "Title for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 380,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:106": []byte(`{
            "id": 106,
            "key": "motions_export_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 382,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:107": []byte(`{
            "id": 107,
            "key": "motions_export_submitter_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show submitters and recommendation/state in table of contents",
              "helpText": "",
              "choices": null,
              "weight": 384,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:108": []byte(`{
            "id": 108,
            "key": "motions_export_follow_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show checkbox to record decision",
              "helpText": "",
              "choices": null,
              "weight": 386,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:109": []byte(`{
            "id": 109,
            "key": "assignment_poll_method",
            "value": "votes",
            "data": {
              "defaultValue": "votes",
              "inputType": "choice",
              "label": "Default election method",
              "helpText": "",
              "choices": [
                {
                  "value": "votes",
                  "display_name": "Yes per candidate"
                },
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                }
              ],
              "weight": 400,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:11": []byte(`{
            "id": 11,
            "key": "general_system_conference_auto_connect",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Connect all users to live conference automatically",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 141,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:110": []byte(`{
            "id": 110,
            "key": "assignment_poll_default_type",
            "value": "analog",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 403,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:111": []byte(`{
            "id": 111,
            "key": "assignment_poll_default_100_percent_base",
            "value": "valid",
            "data": {
              "defaultValue": "valid",
              "inputType": "choice",
              "label": "Default 100 % base of an election result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                },
                {
                  "value": "votes",
                  "display_name": "Sum of votes including general No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 405,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:112": []byte(`{
            "id": 112,
            "key": "assignment_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 410,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:113": []byte(`{
            "id": 113,
            "key": "assignment_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:114": []byte(`{
            "id": 114,
            "key": "assignment_poll_sort_poll_result_by_votes",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Sort election results by amount of votes",
              "helpText": "",
              "choices": null,
              "weight": 420,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:115": []byte(`{
            "id": 115,
            "key": "assignment_poll_add_candidates_to_list_of_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Put all candidates on the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 425,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:116": []byte(`{
            "id": 116,
            "key": "assignments_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 430,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:117": []byte(`{
            "id": 117,
            "key": "assignments_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 435,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:118": []byte(`{
            "id": 118,
            "key": "assignments_pdf_title",
            "value": "Elections",
            "data": {
              "defaultValue": "Elections",
              "inputType": "string",
              "label": "Title for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 460,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:119": []byte(`{
            "id": 119,
            "key": "assignments_pdf_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 470,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:12": []byte(`{
            "id": 12,
            "key": "general_system_conference_los_restriction",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow only current speakers and list of speakers managers to enter the live conference",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 142,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:13": []byte(`{
            "id": 13,
            "key": "general_system_stream_url",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Livestream url",
              "helpText": "Remove URL to deactivate livestream. Check extra group permission to see livestream.",
              "choices": null,
              "weight": 143,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:14": []byte(`{
            "id": 14,
            "key": "general_system_enable_anonymous",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow access for anonymous guest users",
              "helpText": "",
              "choices": null,
              "weight": 150,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:15": []byte(`{
            "id": 15,
            "key": "general_login_info_text",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Show this text on the login page",
              "helpText": "",
              "choices": null,
              "weight": 152,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:16": []byte(`{
            "id": 16,
            "key": "openslides_theme",
            "value": "openslides-default-light-theme",
            "data": {
              "defaultValue": "openslides-default-light-theme",
              "inputType": "choice",
              "label": "OpenSlides Theme",
              "helpText": "",
              "choices": [
                {
                  "value": "openslides-default-light-theme",
                  "display_name": "OpenSlides Default"
                },
                {
                  "value": "openslides-default-dark-theme",
                  "display_name": "OpenSlides Dark"
                },
                {
                  "value": "openslides-red-light-theme",
                  "display_name": "OpenSlides Red"
                },
                {
                  "value": "openslides-red-dark-theme",
                  "display_name": "OpenSlides Red Dark"
                },
                {
                  "value": "openslides-green-light-theme",
                  "display_name": "OpenSlides Green"
                },
                {
                  "value": "openslides-green-dark-theme",
                  "display_name": "OpenSlides Green Dark"
                },
                {
                  "value": "openslides-solarized-dark-theme",
                  "display_name": "OpenSlides Solarized"
                }
              ],
              "weight": 154,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:17": []byte(`{
            "id": 17,
            "key": "general_csv_separator",
            "value": ",",
            "data": {
              "defaultValue": ",",
              "inputType": "string",
              "label": "Separator used for all csv exports and examples",
              "helpText": "",
              "choices": null,
              "weight": 160,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:18": []byte(`{
            "id": 18,
            "key": "general_csv_encoding",
            "value": "utf-8",
            "data": {
              "defaultValue": "utf-8",
              "inputType": "choice",
              "label": "Default encoding for all csv exports",
              "helpText": "",
              "choices": [
                {
                  "value": "utf-8",
                  "display_name": "UTF-8"
                },
                {
                  "value": "iso-8859-15",
                  "display_name": "ISO-8859-15"
                }
              ],
              "weight": 162,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:19": []byte(`{
            "id": 19,
            "key": "general_export_pdf_pagenumber_alignment",
            "value": "center",
            "data": {
              "defaultValue": "center",
              "inputType": "choice",
              "label": "Page number alignment in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "left",
                  "display_name": "Left"
                },
                {
                  "value": "center",
                  "display_name": "Center"
                },
                {
                  "value": "right",
                  "display_name": "Right"
                }
              ],
              "weight": 164,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:2": []byte(`{
            "id": 2,
            "key": "general_event_name",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Event name",
              "helpText": "",
              "choices": null,
              "weight": 110,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:20": []byte(`{
            "id": 20,
            "key": "general_export_pdf_fontsize",
            "value": "10",
            "data": {
              "defaultValue": "10",
              "inputType": "choice",
              "label": "Standard font size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "10",
                  "display_name": "10"
                },
                {
                  "value": "11",
                  "display_name": "11"
                },
                {
                  "value": "12",
                  "display_name": "12"
                }
              ],
              "weight": 166,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:21": []byte(`{
            "id": 21,
            "key": "general_export_pdf_pagesize",
            "value": "A4",
            "data": {
              "defaultValue": "A4",
              "inputType": "choice",
              "label": "Standard page size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "A4",
                  "display_name": "DIN A4"
                },
                {
                  "value": "A5",
                  "display_name": "DIN A5"
                }
              ],
              "weight": 168,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:22": []byte(`{
            "id": 22,
            "key": "logos_available",
            "value": [
              "logo_projector_main",
              "logo_projector_header",
              "logo_web_header",
              "logo_pdf_header_L",
              "logo_pdf_header_R",
              "logo_pdf_footer_L",
              "logo_pdf_footer_R",
              "logo_pdf_ballot_paper"
            ],
            "data": null
          }`),
		"core/config:23": []byte(`{
            "id": 23,
            "key": "logo_projector_main",
            "value": {
              "display_name": "Projector logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:24": []byte(`{
            "id": 24,
            "key": "logo_projector_header",
            "value": {
              "display_name": "Projector header image",
              "path": ""
            },
            "data": null
          }`),
		"core/config:25": []byte(`{
            "id": 25,
            "key": "logo_web_header",
            "value": {
              "display_name": "Web interface header logo",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:26": []byte(`{
            "id": 26,
            "key": "logo_pdf_header_L",
            "value": {
              "display_name": "PDF header logo (left)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:27": []byte(`{
            "id": 27,
            "key": "logo_pdf_header_R",
            "value": {
              "display_name": "PDF header logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:28": []byte(`{
            "id": 28,
            "key": "logo_pdf_footer_L",
            "value": {
              "display_name": "PDF footer logo (left)",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:29": []byte(`{
            "id": 29,
            "key": "logo_pdf_footer_R",
            "value": {
              "display_name": "PDF footer logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:3": []byte(`{
            "id": 3,
            "key": "general_event_description",
            "value": "Presentation and assembly system",
            "data": {
              "defaultValue": "Presentation and assembly system",
              "inputType": "string",
              "label": "Short description of event",
              "helpText": "",
              "choices": null,
              "weight": 115,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:30": []byte(`{
            "id": 30,
            "key": "logo_pdf_ballot_paper",
            "value": {
              "display_name": "PDF ballot paper logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:31": []byte(`{
            "id": 31,
            "key": "fonts_available",
            "value": [
              "font_regular",
              "font_italic",
              "font_bold",
              "font_bold_italic",
              "font_monospace"
            ],
            "data": null
          }`),
		"core/config:32": []byte(`{
            "id": 32,
            "key": "font_regular",
            "value": {
              "display_name": "Font regular",
              "default": "assets/fonts/fira-sans-latin-400.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:33": []byte(`{
            "id": 33,
            "key": "font_italic",
            "value": {
              "display_name": "Font italic",
              "default": "assets/fonts/fira-sans-latin-400italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:34": []byte(`{
            "id": 34,
            "key": "font_bold",
            "value": {
              "display_name": "Font bold",
              "default": "assets/fonts/fira-sans-latin-500.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:35": []byte(`{
            "id": 35,
            "key": "font_bold_italic",
            "value": {
              "display_name": "Font bold italic",
              "default": "assets/fonts/fira-sans-latin-500italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:36": []byte(`{
            "id": 36,
            "key": "font_monospace",
            "value": {
              "display_name": "Font monospace",
              "default": "assets/fonts/roboto-condensed-bold.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:37": []byte(`{
            "id": 37,
            "key": "translations",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "translations",
              "label": "Custom translations",
              "helpText": "",
              "choices": null,
              "weight": 1000,
              "group": "Custom translations",
              "subgroup": "General"
            }
          }`),
		"core/config:38": []byte(`{
            "id": 38,
            "key": "config_version",
            "value": 2,
            "data": null
          }`),
		"core/config:39": []byte(`{
            "id": 39,
            "key": "db_id",
            "value": "2c3bd736c87a48a4be1f0dc707702144",
            "data": null
          }`),
		"core/config:4": []byte(`{
            "id": 4,
            "key": "general_event_date",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event date",
              "helpText": "",
              "choices": null,
              "weight": 120,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:40": []byte(`{
            "id": 40,
            "key": "users_sort_by",
            "value": "first_name",
            "data": {
              "defaultValue": "first_name",
              "inputType": "choice",
              "label": "Sort name of participants by",
              "helpText": "",
              "choices": [
                {
                  "value": "first_name",
                  "display_name": "Given name"
                },
                {
                  "value": "last_name",
                  "display_name": "Surname"
                },
                {
                  "value": "number",
                  "display_name": "Participant number"
                }
              ],
              "weight": 510,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:41": []byte(`{
            "id": 41,
            "key": "users_enable_presence_view",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Enable participant presence view",
              "helpText": "",
              "choices": null,
              "weight": 511,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:42": []byte(`{
            "id": 42,
            "key": "users_allow_self_set_present",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow users to set themselves as present",
              "helpText": "e.g. for online meetings",
              "choices": null,
              "weight": 512,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:43": []byte(`{
            "id": 43,
            "key": "users_activate_vote_weight",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate vote weight",
              "helpText": "",
              "choices": null,
              "weight": 513,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:44": []byte(`{
            "id": 44,
            "key": "users_pdf_welcometitle",
            "value": "Welcome to OpenSlides",
            "data": {
              "defaultValue": "Welcome to OpenSlides",
              "inputType": "string",
              "label": "Title for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 520,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:45": []byte(`{
            "id": 45,
            "key": "users_pdf_welcometext",
            "value": "[Place for your welcome and help text.]",
            "data": {
              "defaultValue": "[Place for your welcome and help text.]",
              "inputType": "string",
              "label": "Help text for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 530,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:46": []byte(`{
            "id": 46,
            "key": "users_pdf_url",
            "value": "http://example.com:8000",
            "data": {
              "defaultValue": "http://example.com:8000",
              "inputType": "string",
              "label": "System URL",
              "helpText": "Used for QRCode in PDF of access data.",
              "choices": null,
              "weight": 540,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:47": []byte(`{
            "id": 47,
            "key": "users_pdf_wlan_ssid",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN name (SSID)",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 550,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:48": []byte(`{
            "id": 48,
            "key": "users_pdf_wlan_password",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN password",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 560,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:49": []byte(`{
            "id": 49,
            "key": "users_pdf_wlan_encryption",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "choice",
              "label": "WLAN encryption",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": [
                {
                  "value": "",
                  "display_name": "---------"
                },
                {
                  "value": "WEP",
                  "display_name": "WEP"
                },
                {
                  "value": "WPA",
                  "display_name": "WPA/WPA2"
                },
                {
                  "value": "nopass",
                  "display_name": "No encryption"
                }
              ],
              "weight": 570,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:5": []byte(`{
            "id": 5,
            "key": "general_event_location",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event location",
              "helpText": "",
              "choices": null,
              "weight": 125,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:50": []byte(`{
            "id": 50,
            "key": "users_email_sender",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Sender name",
              "helpText": "The sender address is defined in the OpenSlides server settings and should modified by administrator only.",
              "choices": null,
              "weight": 600,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:51": []byte(`{
            "id": 51,
            "key": "users_email_replyto",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Reply address",
              "helpText": "",
              "choices": null,
              "weight": 601,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:52": []byte(`{
            "id": 52,
            "key": "users_email_subject",
            "value": "OpenSlides access data",
            "data": {
              "defaultValue": "OpenSlides access data",
              "inputType": "string",
              "label": "Email subject",
              "helpText": "You can use {event_name} and {username} as placeholder.",
              "choices": null,
              "weight": 605,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:53": []byte(`{
            "id": 53,
            "key": "users_email_body",
            "value": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
            "data": {
              "defaultValue": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
              "inputType": "text",
              "label": "Email body",
              "helpText": "Use these placeholders: {name}, {event_name}, {url}, {username}, {password}. The url referrs to the system url.",
              "choices": null,
              "weight": 610,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:54": []byte(`{
            "id": 54,
            "key": "agenda_start_event_date_time",
            "value": null,
            "data": {
              "defaultValue": null,
              "inputType": "datetimepicker",
              "label": "Begin of event",
              "helpText": "Input format: DD.MM.YYYY HH:MM",
              "choices": null,
              "weight": 200,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:55": []byte(`{
            "id": 55,
            "key": "agenda_show_subtitle",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show subtitles in the agenda",
              "helpText": "",
              "choices": null,
              "weight": 201,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:56": []byte(`{
            "id": 56,
            "key": "agenda_enable_numbering",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Enable numbering for agenda items",
              "helpText": "",
              "choices": null,
              "weight": 205,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:57": []byte(`{
            "id": 57,
            "key": "agenda_number_prefix",
            "value": "TOP",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Numbering prefix for agenda items",
              "helpText": "This prefix will be set if you run the automatic agenda numbering.",
              "choices": null,
              "weight": 206,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:58": []byte(`{
            "id": 58,
            "key": "agenda_numeral_system",
            "value": "arabic",
            "data": {
              "defaultValue": "arabic",
              "inputType": "choice",
              "label": "Numeral system for agenda items",
              "helpText": "",
              "choices": [
                {
                  "value": "arabic",
                  "display_name": "Arabic"
                },
                {
                  "value": "roman",
                  "display_name": "Roman"
                }
              ],
              "weight": 207,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:59": []byte(`{
            "id": 59,
            "key": "agenda_item_creation",
            "value": "default_yes",
            "data": {
              "defaultValue": "default_yes",
              "inputType": "choice",
              "label": "Add to agenda",
              "helpText": "",
              "choices": [
                {
                  "value": "always",
                  "display_name": "Always"
                },
                {
                  "value": "never",
                  "display_name": "Never"
                },
                {
                  "value": "default_yes",
                  "display_name": "Ask, default yes"
                },
                {
                  "value": "default_no",
                  "display_name": "Ask, default no"
                }
              ],
              "weight": 210,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:6": []byte(`{
            "id": 6,
            "key": "general_event_legal_notice",
            "value": "<a href=\"http://www.openslides.org\">OpenSlides</a> is a free web based presentation and assembly system for visualizing and controlling agenda, motions and elections of an assembly.",
            "data": null
          }`),
		"core/config:60": []byte(`{
            "id": 60,
            "key": "agenda_new_items_default_visibility",
            "value": "2",
            "data": {
              "defaultValue": "2",
              "inputType": "choice",
              "label": "Default visibility for new agenda items (except topics)",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Public item"
                },
                {
                  "value": "2",
                  "display_name": "Internal item"
                },
                {
                  "value": "3",
                  "display_name": "Hidden item"
                }
              ],
              "weight": 211,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:61": []byte(`{
            "id": 61,
            "key": "agenda_hide_internal_items_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Hide internal items when projecting subitems",
              "helpText": "",
              "choices": null,
              "weight": 212,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:62": []byte(`{
            "id": 62,
            "key": "agenda_show_last_speakers",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of last speakers to be shown on the projector",
              "helpText": "",
              "choices": null,
              "weight": 220,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:63": []byte(`{
            "id": 63,
            "key": "agenda_show_next_speakers",
            "value": -1,
            "data": {
              "defaultValue": -1,
              "inputType": "integer",
              "label": "Number of the next speakers to be shown on the projector",
              "helpText": "Enter number of the next shown speakers. Choose -1 to show all next speakers.",
              "choices": null,
              "weight": 222,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:64": []byte(`{
            "id": 64,
            "key": "agenda_countdown_warning_time",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Show orange countdown in the last x seconds of speaking time",
              "helpText": "Enter duration in seconds. Choose 0 to disable warning color.",
              "choices": null,
              "weight": 224,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:65": []byte(`{
            "id": 65,
            "key": "projector_default_countdown",
            "value": 60,
            "data": {
              "defaultValue": 60,
              "inputType": "integer",
              "label": "Predefined seconds of new countdowns",
              "helpText": "",
              "choices": null,
              "weight": 226,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:66": []byte(`{
            "id": 66,
            "key": "agenda_couple_countdown_and_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Couple countdown with the list of speakers",
              "helpText": "[Begin speech] starts the countdown, [End speech] stops the countdown.",
              "choices": null,
              "weight": 228,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:67": []byte(`{
            "id": 67,
            "key": "agenda_hide_amount_of_speakers",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide the amount of speakers in subtitle of list of speakers slide",
              "helpText": "",
              "choices": null,
              "weight": 230,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:68": []byte(`{
            "id": 68,
            "key": "agenda_present_speakers_only",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Only present participants can be added to the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 232,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:69": []byte(`{
            "id": 69,
            "key": "agenda_show_first_contribution",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show hint »first speech« in the list of speakers management view",
              "helpText": "",
              "choices": null,
              "weight": 234,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:7": []byte(`{
            "id": 7,
            "key": "general_event_privacy_policy",
            "value": "",
            "data": null
          }`),
		"core/config:70": []byte(`{
            "id": 70,
            "key": "motions_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new motions",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 310,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:71": []byte(`{
            "id": 71,
            "key": "motions_statute_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new statute amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 312,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:72": []byte(`{
            "id": 72,
            "key": "motions_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 314,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:73": []byte(`{
            "id": 73,
            "key": "motions_preamble",
            "value": "The assembly may decide:",
            "data": {
              "defaultValue": "The assembly may decide:",
              "inputType": "string",
              "label": "Motion preamble",
              "helpText": "",
              "choices": null,
              "weight": 320,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:74": []byte(`{
            "id": 74,
            "key": "motions_default_line_numbering",
            "value": "outside",
            "data": {
              "defaultValue": "outside",
              "inputType": "choice",
              "label": "Default line numbering",
              "helpText": "",
              "choices": [
                {
                  "value": "outside",
                  "display_name": "outside"
                },
                {
                  "value": "inline",
                  "display_name": "inline"
                },
                {
                  "value": "none",
                  "display_name": "Disabled"
                }
              ],
              "weight": 322,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:75": []byte(`{
            "id": 75,
            "key": "motions_line_length",
            "value": 85,
            "data": {
              "defaultValue": 85,
              "inputType": "integer",
              "label": "Line length",
              "helpText": "The maximum number of characters per line. Relevant when line numbering is enabled. Min: 40",
              "choices": null,
              "weight": 323,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:76": []byte(`{
            "id": 76,
            "key": "motions_reason_required",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Reason required for creating new motion",
              "helpText": "",
              "choices": null,
              "weight": 324,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:77": []byte(`{
            "id": 77,
            "key": "motions_disable_text_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide motion text on projector",
              "helpText": "",
              "choices": null,
              "weight": 325,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:78": []byte(`{
            "id": 78,
            "key": "motions_disable_reason_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide reason on projector",
              "helpText": "",
              "choices": null,
              "weight": 326,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:79": []byte(`{
            "id": 79,
            "key": "motions_disable_recommendation_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide recommendation on projector",
              "helpText": "",
              "choices": null,
              "weight": 327,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:8": []byte(`{
            "id": 8,
            "key": "general_event_welcome_title",
            "value": "Welcome to OpenSlides",
            "data": null
          }`),
		"core/config:80": []byte(`{
            "id": 80,
            "key": "motions_hide_referring_motions",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide referring motions",
              "helpText": "",
              "choices": null,
              "weight": 328,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:81": []byte(`{
            "id": 81,
            "key": "motions_disable_sidebox_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show meta information box below the title on projector",
              "helpText": "",
              "choices": null,
              "weight": 329,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:82": []byte(`{
            "id": 82,
            "key": "motions_show_sequential_numbers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show the sequential number for a motion",
              "helpText": "In motion list, motion detail and PDF.",
              "choices": null,
              "weight": 330,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:83": []byte(`{
            "id": 83,
            "key": "motions_recommendations_by",
            "value": "ABK",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender",
              "helpText": "Will be displayed as label before selected recommendation. Use an empty value to disable the recommendation system.",
              "choices": null,
              "weight": 332,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:84": []byte(`{
            "id": 84,
            "key": "motions_statute_recommendations_by",
            "value": "ABK2",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender for statute amendments",
              "helpText": "Will be displayed as label before selected recommendation in statute amendments.",
              "choices": null,
              "weight": 333,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:85": []byte(`{
            "id": 85,
            "key": "motions_recommendation_text_mode",
            "value": "diff",
            "data": {
              "defaultValue": "diff",
              "inputType": "choice",
              "label": "Default text version for change recommendations",
              "helpText": "",
              "choices": [
                {
                  "value": "original",
                  "display_name": "Original version"
                },
                {
                  "value": "changed",
                  "display_name": "Changed version"
                },
                {
                  "value": "diff",
                  "display_name": "Diff version"
                },
                {
                  "value": "agreed",
                  "display_name": "Final version"
                }
              ],
              "weight": 334,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:86": []byte(`{
            "id": 86,
            "key": "motions_motions_sorting",
            "value": "identifier",
            "data": {
              "defaultValue": "identifier",
              "inputType": "choice",
              "label": "Sort motions by",
              "helpText": "",
              "choices": [
                {
                  "value": "weight",
                  "display_name": "Call list"
                },
                {
                  "value": "identifier",
                  "display_name": "Identifier"
                }
              ],
              "weight": 335,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:87": []byte(`{
            "id": 87,
            "key": "motions_identifier",
            "value": "per_category",
            "data": {
              "defaultValue": "per_category",
              "inputType": "choice",
              "label": "Identifier",
              "helpText": "",
              "choices": [
                {
                  "value": "per_category",
                  "display_name": "Numbered per category"
                },
                {
                  "value": "serially_numbered",
                  "display_name": "Serially numbered"
                },
                {
                  "value": "manually",
                  "display_name": "Set it manually"
                }
              ],
              "weight": 340,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:88": []byte(`{
            "id": 88,
            "key": "motions_identifier_min_digits",
            "value": 1,
            "data": {
              "defaultValue": 1,
              "inputType": "integer",
              "label": "Number of minimal digits for identifier",
              "helpText": "Uses leading zeros to sort motions correctly by identifier.",
              "choices": null,
              "weight": 342,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:89": []byte(`{
            "id": 89,
            "key": "motions_identifier_with_blank",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow blank in identifier",
              "helpText": "Blank between prefix and number, e.g. 'A 001'.",
              "choices": null,
              "weight": 344,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:9": []byte(`{
            "id": 9,
            "key": "general_event_welcome_text",
            "value": "[Space for your welcome text.]",
            "data": null
          }`),
		"core/config:90": []byte(`{
            "id": 90,
            "key": "motions_statutes_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate statute amendments",
              "helpText": "",
              "choices": null,
              "weight": 350,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:91": []byte(`{
            "id": 91,
            "key": "motions_amendments_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate amendments",
              "helpText": "",
              "choices": null,
              "weight": 351,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:92": []byte(`{
            "id": 92,
            "key": "motions_amendments_main_table",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show amendments together with motions",
              "helpText": "",
              "choices": null,
              "weight": 352,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:93": []byte(`{
            "id": 93,
            "key": "motions_amendments_prefix",
            "value": "Ä-",
            "data": {
              "defaultValue": "-",
              "inputType": "string",
              "label": "Prefix for the identifier for amendments",
              "helpText": "",
              "choices": null,
              "weight": 353,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:94": []byte(`{
            "id": 94,
            "key": "motions_amendments_text_mode",
            "value": "paragraph",
            "data": {
              "defaultValue": "paragraph",
              "inputType": "choice",
              "label": "How to create new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "freestyle",
                  "display_name": "Empty text field"
                },
                {
                  "value": "fulltext",
                  "display_name": "Edit the whole motion text"
                },
                {
                  "value": "paragraph",
                  "display_name": "Paragraph-based, Diff-enabled"
                }
              ],
              "weight": 354,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:95": []byte(`{
            "id": 95,
            "key": "motions_amendments_multiple_paragraphs",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Amendments can change multiple paragraphs",
              "helpText": "",
              "choices": null,
              "weight": 355,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:96": []byte(`{
            "id": 96,
            "key": "motions_amendments_of_amendments",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow amendments of amendments",
              "helpText": "",
              "choices": null,
              "weight": 356,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:97": []byte(`{
            "id": 97,
            "key": "motions_min_supporters",
            "value": 1,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of (minimum) required supporters for a motion",
              "helpText": "Choose 0 to disable the supporting system.",
              "choices": null,
              "weight": 360,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:98": []byte(`{
            "id": 98,
            "key": "motions_remove_supporters",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Remove all supporters of a motion if a submitter edits his motion in early state",
              "helpText": "",
              "choices": null,
              "weight": 361,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:99": []byte(`{
            "id": 99,
            "key": "motion_poll_default_type",
            "value": "named",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 367,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/countdown:1": []byte(`{
            "id": 1,
            "title": "Default countdown",
            "description": "",
            "default_time": 60,
            "countdown_time": 1597141855.090748,
            "running": true
          }`),
		"core/projection-default:1": []byte(`{
            "id": 1,
            "name": "agenda_all_items",
            "display_name": "Agenda",
            "projector_id": 1
          }`),
		"core/projection-default:10": []byte(`{
            "id": 10,
            "name": "messages",
            "display_name": "Messages",
            "projector_id": 1
          }`),
		"core/projection-default:11": []byte(`{
            "id": 11,
            "name": "countdowns",
            "display_name": "Countdowns",
            "projector_id": 1
          }`),
		"core/projection-default:12": []byte(`{
            "id": 12,
            "name": "assignment_poll",
            "display_name": "Ballots",
            "projector_id": 1
          }`),
		"core/projection-default:13": []byte(`{
            "id": 13,
            "name": "motion_poll",
            "display_name": "Motion votes",
            "projector_id": 1
          }`),
		"core/projection-default:2": []byte(`{
            "id": 2,
            "name": "topics",
            "display_name": "Topics",
            "projector_id": 1
          }`),
		"core/projection-default:3": []byte(`{
            "id": 3,
            "name": "agenda_list_of_speakers",
            "display_name": "List of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:4": []byte(`{
            "id": 4,
            "name": "agenda_current_list_of_speakers",
            "display_name": "Current list of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:5": []byte(`{
            "id": 5,
            "name": "motions",
            "display_name": "Motions",
            "projector_id": 1
          }`),
		"core/projection-default:6": []byte(`{
            "id": 6,
            "name": "motionBlocks",
            "display_name": "Motion blocks",
            "projector_id": 1
          }`),
		"core/projection-default:7": []byte(`{
            "id": 7,
            "name": "assignments",
            "display_name": "Elections",
            "projector_id": 1
          }`),
		"core/projection-default:8": []byte(`{
            "id": 8,
            "name": "users",
            "display_name": "Participants",
            "projector_id": 1
          }`),
		"core/projection-default:9": []byte(`{
            "id": 9,
            "name": "mediafiles",
            "display_name": "Files",
            "projector_id": 1
          }`),
		"core/projector-message:1": []byte(`{
            "id": 1,
            "message": "<p>test</p>"
          }`),
		"core/projector:1": []byte(`{
            "id": 1,
            "elements": [
              {
                "name": "mediafiles/mediafile",
                "id": 3
              }
            ],
            "elements_preview": [],
            "elements_history": [
              [
                {
                  "name": "assignments/assignment",
                  "id": 1
                }
              ]
            ],
            "scale": 0,
            "scroll": 0,
            "name": "Default projector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [
              1,
              2,
              3,
              4,
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13
            ],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/projector:2": []byte(`{
            "id": 2,
            "elements": [
              {
                "name": "core/clock",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers",
                "stable": false
              },
              {
                "name": "agenda/current-speaker-chyron",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers-overlay",
                "stable": true
              }
            ],
            "elements_preview": [],
            "elements_history": [],
            "scale": 0,
            "scroll": 0,
            "name": "sideprojector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/tag:1": []byte(`{
            "id": 1,
            "name": "T1"
          }`),
		"core/tag:2": []byte(`{
            "id": 2,
            "name": "T2"
          }`),
		"mediafiles/mediafile:1": []byte(`{
            "id": 1,
            "title": "folder",
            "media_url_prefix": "/media/",
            "filesize": "",
            "mimetype": "",
            "pdf_information": {},
            "access_groups_id": [
              3
            ],
            "create_timestamp": "2020-08-11T11:15:50.719490+02:00",
            "is_directory": true,
            "path": "folder/",
            "parent_id": null,
            "list_of_speakers_id": 4,
            "inherited_access_groups_id": [
              3
            ]
          }`),
		"mediafiles/mediafile:2": []byte(`{
            "id": 2,
            "title": "A.txt",
            "media_url_prefix": "/media/",
            "filesize": "< 1 kB",
            "mimetype": "text/plain",
            "pdf_information": {},
            "access_groups_id": [],
            "create_timestamp": "2020-08-11T11:16:07.013124+02:00",
            "is_directory": false,
            "path": "A.txt",
            "parent_id": null,
            "list_of_speakers_id": 5,
            "inherited_access_groups_id": true
          }`),
		"motions/category:1": []byte(`{
            "id": 1,
            "name": "first",
            "prefix": "A",
            "parent_id": null,
            "weight": 2,
            "level": 0
          }`),
		"motions/category:2": []byte(`{
            "id": 2,
            "name": "second",
            "prefix": "B",
            "parent_id": 1,
            "weight": 4,
            "level": 1
          }`),
		"motions/category:3": []byte(`{
            "id": 3,
            "name": "third",
            "prefix": "C",
            "parent_id": null,
            "weight": 6,
            "level": 0
          }`),
		"motions/motion-block:1": []byte(`{
            "id": 1,
            "title": "block",
            "agenda_item_id": 8,
            "list_of_speakers_id": 12,
            "internal": false,
            "motions_id": [
              1,
              2
            ]
          }`),
		"motions/motion-comment-section:1": []byte(`{
            "id": 1,
            "name": "public",
            "read_groups_id": [
              2,
              3
            ],
            "write_groups_id": [
              2,
              3
            ],
            "weight": 10000
          }`),
		"motions/motion-option:1": []byte(`{
            "id": 1,
            "yes": "0.000000",
            "no": "1.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-option:2": []byte(`{
            "id": 2,
            "yes": "1.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 2,
            "pollstate": 4
          }`),
		"motions/motion-poll:1": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "named",
            "title": "Abstimmung",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              1
            ],
            "id": 1,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": true
          }`),
		"motions/motion-poll:2": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "pseudoanonymous",
            "title": "Abstimmung (2)",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              2
            ],
            "id": 2,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": true
          }`),
		"motions/motion-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "N",
            "user_id": 5,
            "option_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-vote:2": []byte(`{
            "id": 2,
            "weight": "1.000000",
            "value": "Y",
            "user_id": null,
            "option_id": 2,
            "pollstate": 4
          }`),
		"motions/motion:2": []byte(`{
            "id": 2,
            "identifier": "Ä-1",
            "title": "Änderungsantrag zu Leadmotion1",
            "text": "",
            "amendment_paragraphs": [
              "<p>Lorem ipsum dolor sit amet, consectedfgddfgdf&nbsp; gdfgdfg dfg dfg ww ontes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>"
            ],
            "modified_final_version": "",
            "reason": "",
            "parent_id": 1,
            "category_id": 2,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 3,
                "user_id": 1,
                "motion_id": 2,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [],
            "attachments_id": [],
            "agenda_item_id": 6,
            "list_of_speakers_id": 9,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:39:35.025914+02:00",
            "last_modified": "2020-08-11T12:41:55.666495+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/motion:3": []byte(`{
            "id": 3,
            "identifier": "2",
            "title": "Public",
            "text": "<p>a</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>a</p>",
            "parent_id": null,
            "category_id": 1,
            "category_weight": 10000,
            "comments": [
              {
                "id": 1,
                "comment": "<p>test</p>",
                "section_id": 1,
                "read_groups_id": [
                  2,
                  3
                ]
              }
            ],
            "motion_block_id": 2,
            "origin": "",
            "submitters": [
              {
                "id": 4,
                "user_id": 1,
                "motion_id": 3,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [
              1
            ],
            "attachments_id": [],
            "agenda_item_id": 7,
            "list_of_speakers_id": 10,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T12:24:45.289233+02:00",
            "last_modified": "2020-08-11T12:41:51.319986+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/state:1": []byte(`{
            "id": 1,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              2,
              3,
              4
            ],
            "workflow_id": 1
          }`),
		"motions/state:10": []byte(`{
            "id": 10,
            "name": "withdrawed",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:11": []byte(`{
            "id": 11,
            "name": "adjourned",
            "recommendation_label": "Adjournment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:12": []byte(`{
            "id": 12,
            "name": "not concerned",
            "recommendation_label": "No concernment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:13": []byte(`{
            "id": 13,
            "name": "refered to committee",
            "recommendation_label": "Referral to committee",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:14": []byte(`{
            "id": 14,
            "name": "needs review",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:15": []byte(`{
            "id": 15,
            "name": "rejected (not authorized)",
            "recommendation_label": "Rejection (not authorized)",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:2": []byte(`{
            "id": 2,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:3": []byte(`{
            "id": 3,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:4": []byte(`{
            "id": 4,
            "name": "not decided",
            "recommendation_label": "No decision",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:5": []byte(`{
            "id": 5,
            "name": "in progress",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [
              "is_submitter"
            ],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": true,
            "dont_set_identifier": true,
            "show_state_extension_field": true,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              6,
              10
            ],
            "workflow_id": 2
          }`),
		"motions/state:6": []byte(`{
            "id": 6,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": true,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              7,
              10,
              15
            ],
            "workflow_id": 2
          }`),
		"motions/state:7": []byte(`{
            "id": 7,
            "name": "permitted",
            "recommendation_label": "Permission",
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": true,
            "next_states_id": [
              8,
              9,
              10,
              11,
              12,
              13,
              14
            ],
            "workflow_id": 2
          }`),
		"motions/state:8": []byte(`{
            "id": 8,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:9": []byte(`{
            "id": 9,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/statute-paragraph:1": []byte(`{
            "id": 1,
            "title": "Statute",
            "text": "<p>test</p>",
            "weight": 10000
          }`),
		"motions/workflow:1": []byte(`{
            "id": 1,
            "name": "Simple Workflow",
            "states_id": [
              1,
              2,
              3,
              4
            ],
            "first_state_id": 1
          }`),
		"motions/workflow:2": []byte(`{
            "id": 2,
            "name": "Complex Workflow",
            "states_id": [
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13,
              14,
              15
            ],
            "first_state_id": 5
          }`),
		"topics/topic:1": []byte(`{
            "id": 1,
            "title": "Topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 1,
            "list_of_speakers_id": 1
          }`),
		"topics/topic:2": []byte(`{
            "id": 2,
            "title": "Hidden",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 2,
            "list_of_speakers_id": 2
          }`),
		"topics/topic:3": []byte(`{
            "id": 3,
            "title": "Internal",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 3,
            "list_of_speakers_id": 3
          }`),
		"topics/topic:4": []byte(`{
            "id": 4,
            "title": "Another public topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 9,
            "list_of_speakers_id": 14
          }`),
		"users/group:1": []byte(`{
            "id": 1,
            "name": "Default",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_see",
              "users.can_change_password"
            ]
          }`),
		"users/group:2": []byte(`{
            "id": 2,
            "name": "Admin",
            "permissions": []
          }`),
		"users/group:3": []byte(`{
            "id": 3,
            "name": "Delegates",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password"
            ]
          }`),
		"users/group:4": []byte(`{
            "id": 4,
            "name": "Staff",
            "permissions": [
              "agenda.can_manage",
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_manage_list_of_speakers",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_manage",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_history",
              "core.can_manage_projector",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "core.can_manage_tags",
              "mediafiles.can_manage",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_manage",
              "motions.can_manage_metadata",
              "motions.can_see",
              "motions.can_see_internal",
              "users.can_change_password",
              "users.can_manage",
              "users.can_see_extra_data",
              "users.can_see_name"
            ]
          }`),
		"users/group:5": []byte(`{
            "id": 5,
            "name": "Committees",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password",
              "users.can_see_name"
            ]
          }`),
		"users/user:1": []byte(`{
            "first_name": "",
            "username": "admin",
            "about_me": "",
            "id": 1,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              2
            ],
            "number": "",
            "last_name": "Administrator",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:2": []byte(`{
            "first_name": "candidate1",
            "username": "candidate1",
            "about_me": "",
            "id": 2,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:3": []byte(`{
            "first_name": "candidate2",
            "username": "candidate2",
            "about_me": "",
            "id": 3,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:4": []byte(`{
            "first_name": "a",
            "username": "a",
            "about_me": "",
            "id": 4,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:5": []byte(`{
            "first_name": "b",
            "username": "b",
            "about_me": "",
            "id": 5,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "groups_id": [
              3
            ],
            "title": "",
            "email": "",
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:6": []byte(`{
            "first_name": "speaker1",
            "username": "speaker1",
            "about_me": "",
            "id": 6,
            "structure_level": "layer X",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "title",
            "groups_id": [],
            "number": "3",
            "last_name": "the last name",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:7": []byte(`{
            "first_name": "speaker2",
            "username": "speaker2",
            "about_me": "",
            "id": 7,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
	},
	6: {
		"agenda/item:1": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            },
            "is_internal": false,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": null,
            "id": 1,
            "level": 0,
            "title_information": {
              "title": "Topic"
            },
            "closed": false,
            "weight": 2,
            "duration": null
          }`),
		"agenda/item:3": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": null,
            "id": 3,
            "level": 0,
            "title_information": {
              "title": "Internal"
            },
            "closed": false,
            "weight": 8,
            "duration": null
          }`),
		"agenda/item:5": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 5,
            "level": 1,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "closed": false,
            "weight": 14,
            "duration": null
          }`),
		"agenda/item:6": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 6,
            "level": 1,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "closed": false,
            "weight": 16,
            "duration": 0
          }`),
		"agenda/item:7": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": 6,
            "id": 7,
            "level": 2,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "closed": false,
            "weight": 18,
            "duration": null
          }`),
		"agenda/list-of-speakers:1": []byte(`{
            "id": 1,
            "title_information": {
              "title": "Topic"
            },
            "speakers": [
              {
                "id": 3,
                "user_id": 6,
                "begin_time": "2020-08-11T12:28:30.894574+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:10": []byte(`{
            "id": 10,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:11": []byte(`{
            "id": 11,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:12": []byte(`{
            "id": 12,
            "title_information": {
              "title": "block"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:13": []byte(`{
            "id": 13,
            "title_information": {
              "title": "block internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:14": []byte(`{
            "id": 14,
            "title_information": {
              "title": "Another public topic"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:2": []byte(`{
            "id": 2,
            "title_information": {
              "title": "Hidden"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:3": []byte(`{
            "id": 3,
            "title_information": {
              "title": "Internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:4": []byte(`{
            "id": 4,
            "title_information": {
              "title": "folder"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:5": []byte(`{
            "id": 5,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:6": []byte(`{
            "id": 6,
            "title_information": {
              "title": "in.jpg"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:7": []byte(`{
            "id": 7,
            "title_information": {
              "title": "Assignment"
            },
            "speakers": [
              {
                "id": 4,
                "user_id": 6,
                "begin_time": "2020-08-11T12:29:55.054553+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              },
              {
                "id": 6,
                "user_id": 7,
                "begin_time": null,
                "end_time": null,
                "weight": 2,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "assignments/assignment",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:8": []byte(`{
            "id": 8,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:9": []byte(`{
            "id": 9,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            }
          }`),
		"assignments/assignment-option:1": []byte(`{
            "user_id": 2,
            "weight": 1,
            "id": 1,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-option:2": []byte(`{
            "user_id": 3,
            "weight": 3,
            "id": 2,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-poll:1": []byte(`{
            "assignment_id": 1,
            "description": "",
            "pollmethod": "votes",
            "votes_amount": 1,
            "allow_multiple_votes_per_candidate": false,
            "global_no": true,
            "global_abstain": true,
            "state": 2,
            "type": "named",
            "title": "Wahlgang",
            "groups_id": [
              3
            ],
            "options_id": [
              1,
              2
            ],
            "id": 1,
            "onehundred_percent_base": "valid",
            "majority_method": "simple",
            "user_has_voted": false
          }`),
		"assignments/assignment:1": []byte(`{
            "id": 1,
            "title": "Assignment",
            "description": "",
            "open_posts": 1,
            "phase": 0,
            "assignment_related_users": [
              {
                "id": 1,
                "user_id": 2,
                "weight": 1
              },
              {
                "id": 3,
                "user_id": 3,
                "weight": 3
              }
            ],
            "default_poll_description": "",
            "agenda_item_id": 4,
            "list_of_speakers_id": 7,
            "tags_id": [
              2
            ],
            "attachments_id": [],
            "number_poll_candidates": false,
            "polls_id": [
              1
            ]
          }`),
		"core/config:10": []byte(`{
            "id": 10,
            "key": "general_system_conference_show",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show live conference window",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 140,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:100": []byte(`{
            "id": 100,
            "key": "motion_poll_default_100_percent_base",
            "value": "YNA",
            "data": {
              "defaultValue": "YNA",
              "inputType": "choice",
              "label": "Default 100 % base of a voting result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 370,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:101": []byte(`{
            "id": 101,
            "key": "motion_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:102": []byte(`{
            "id": 102,
            "key": "motion_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 372,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:103": []byte(`{
            "id": 103,
            "key": "motions_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 373,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:104": []byte(`{
            "id": 104,
            "key": "motions_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 374,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:105": []byte(`{
            "id": 105,
            "key": "motions_export_title",
            "value": "Motions",
            "data": {
              "defaultValue": "Motions",
              "inputType": "string",
              "label": "Title for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 380,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:106": []byte(`{
            "id": 106,
            "key": "motions_export_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 382,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:107": []byte(`{
            "id": 107,
            "key": "motions_export_submitter_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show submitters and recommendation/state in table of contents",
              "helpText": "",
              "choices": null,
              "weight": 384,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:108": []byte(`{
            "id": 108,
            "key": "motions_export_follow_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show checkbox to record decision",
              "helpText": "",
              "choices": null,
              "weight": 386,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:109": []byte(`{
            "id": 109,
            "key": "assignment_poll_method",
            "value": "votes",
            "data": {
              "defaultValue": "votes",
              "inputType": "choice",
              "label": "Default election method",
              "helpText": "",
              "choices": [
                {
                  "value": "votes",
                  "display_name": "Yes per candidate"
                },
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                }
              ],
              "weight": 400,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:11": []byte(`{
            "id": 11,
            "key": "general_system_conference_auto_connect",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Connect all users to live conference automatically",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 141,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:110": []byte(`{
            "id": 110,
            "key": "assignment_poll_default_type",
            "value": "analog",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 403,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:111": []byte(`{
            "id": 111,
            "key": "assignment_poll_default_100_percent_base",
            "value": "valid",
            "data": {
              "defaultValue": "valid",
              "inputType": "choice",
              "label": "Default 100 % base of an election result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                },
                {
                  "value": "votes",
                  "display_name": "Sum of votes including general No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 405,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:112": []byte(`{
            "id": 112,
            "key": "assignment_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 410,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:113": []byte(`{
            "id": 113,
            "key": "assignment_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:114": []byte(`{
            "id": 114,
            "key": "assignment_poll_sort_poll_result_by_votes",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Sort election results by amount of votes",
              "helpText": "",
              "choices": null,
              "weight": 420,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:115": []byte(`{
            "id": 115,
            "key": "assignment_poll_add_candidates_to_list_of_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Put all candidates on the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 425,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:116": []byte(`{
            "id": 116,
            "key": "assignments_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 430,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:117": []byte(`{
            "id": 117,
            "key": "assignments_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 435,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:118": []byte(`{
            "id": 118,
            "key": "assignments_pdf_title",
            "value": "Elections",
            "data": {
              "defaultValue": "Elections",
              "inputType": "string",
              "label": "Title for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 460,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:119": []byte(`{
            "id": 119,
            "key": "assignments_pdf_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 470,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:12": []byte(`{
            "id": 12,
            "key": "general_system_conference_los_restriction",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow only current speakers and list of speakers managers to enter the live conference",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 142,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:13": []byte(`{
            "id": 13,
            "key": "general_system_stream_url",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Livestream url",
              "helpText": "Remove URL to deactivate livestream. Check extra group permission to see livestream.",
              "choices": null,
              "weight": 143,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:14": []byte(`{
            "id": 14,
            "key": "general_system_enable_anonymous",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow access for anonymous guest users",
              "helpText": "",
              "choices": null,
              "weight": 150,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:15": []byte(`{
            "id": 15,
            "key": "general_login_info_text",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Show this text on the login page",
              "helpText": "",
              "choices": null,
              "weight": 152,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:16": []byte(`{
            "id": 16,
            "key": "openslides_theme",
            "value": "openslides-default-light-theme",
            "data": {
              "defaultValue": "openslides-default-light-theme",
              "inputType": "choice",
              "label": "OpenSlides Theme",
              "helpText": "",
              "choices": [
                {
                  "value": "openslides-default-light-theme",
                  "display_name": "OpenSlides Default"
                },
                {
                  "value": "openslides-default-dark-theme",
                  "display_name": "OpenSlides Dark"
                },
                {
                  "value": "openslides-red-light-theme",
                  "display_name": "OpenSlides Red"
                },
                {
                  "value": "openslides-red-dark-theme",
                  "display_name": "OpenSlides Red Dark"
                },
                {
                  "value": "openslides-green-light-theme",
                  "display_name": "OpenSlides Green"
                },
                {
                  "value": "openslides-green-dark-theme",
                  "display_name": "OpenSlides Green Dark"
                },
                {
                  "value": "openslides-solarized-dark-theme",
                  "display_name": "OpenSlides Solarized"
                }
              ],
              "weight": 154,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:17": []byte(`{
            "id": 17,
            "key": "general_csv_separator",
            "value": ",",
            "data": {
              "defaultValue": ",",
              "inputType": "string",
              "label": "Separator used for all csv exports and examples",
              "helpText": "",
              "choices": null,
              "weight": 160,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:18": []byte(`{
            "id": 18,
            "key": "general_csv_encoding",
            "value": "utf-8",
            "data": {
              "defaultValue": "utf-8",
              "inputType": "choice",
              "label": "Default encoding for all csv exports",
              "helpText": "",
              "choices": [
                {
                  "value": "utf-8",
                  "display_name": "UTF-8"
                },
                {
                  "value": "iso-8859-15",
                  "display_name": "ISO-8859-15"
                }
              ],
              "weight": 162,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:19": []byte(`{
            "id": 19,
            "key": "general_export_pdf_pagenumber_alignment",
            "value": "center",
            "data": {
              "defaultValue": "center",
              "inputType": "choice",
              "label": "Page number alignment in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "left",
                  "display_name": "Left"
                },
                {
                  "value": "center",
                  "display_name": "Center"
                },
                {
                  "value": "right",
                  "display_name": "Right"
                }
              ],
              "weight": 164,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:2": []byte(`{
            "id": 2,
            "key": "general_event_name",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Event name",
              "helpText": "",
              "choices": null,
              "weight": 110,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:20": []byte(`{
            "id": 20,
            "key": "general_export_pdf_fontsize",
            "value": "10",
            "data": {
              "defaultValue": "10",
              "inputType": "choice",
              "label": "Standard font size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "10",
                  "display_name": "10"
                },
                {
                  "value": "11",
                  "display_name": "11"
                },
                {
                  "value": "12",
                  "display_name": "12"
                }
              ],
              "weight": 166,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:21": []byte(`{
            "id": 21,
            "key": "general_export_pdf_pagesize",
            "value": "A4",
            "data": {
              "defaultValue": "A4",
              "inputType": "choice",
              "label": "Standard page size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "A4",
                  "display_name": "DIN A4"
                },
                {
                  "value": "A5",
                  "display_name": "DIN A5"
                }
              ],
              "weight": 168,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:22": []byte(`{
            "id": 22,
            "key": "logos_available",
            "value": [
              "logo_projector_main",
              "logo_projector_header",
              "logo_web_header",
              "logo_pdf_header_L",
              "logo_pdf_header_R",
              "logo_pdf_footer_L",
              "logo_pdf_footer_R",
              "logo_pdf_ballot_paper"
            ],
            "data": null
          }`),
		"core/config:23": []byte(`{
            "id": 23,
            "key": "logo_projector_main",
            "value": {
              "display_name": "Projector logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:24": []byte(`{
            "id": 24,
            "key": "logo_projector_header",
            "value": {
              "display_name": "Projector header image",
              "path": ""
            },
            "data": null
          }`),
		"core/config:25": []byte(`{
            "id": 25,
            "key": "logo_web_header",
            "value": {
              "display_name": "Web interface header logo",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:26": []byte(`{
            "id": 26,
            "key": "logo_pdf_header_L",
            "value": {
              "display_name": "PDF header logo (left)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:27": []byte(`{
            "id": 27,
            "key": "logo_pdf_header_R",
            "value": {
              "display_name": "PDF header logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:28": []byte(`{
            "id": 28,
            "key": "logo_pdf_footer_L",
            "value": {
              "display_name": "PDF footer logo (left)",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:29": []byte(`{
            "id": 29,
            "key": "logo_pdf_footer_R",
            "value": {
              "display_name": "PDF footer logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:3": []byte(`{
            "id": 3,
            "key": "general_event_description",
            "value": "Presentation and assembly system",
            "data": {
              "defaultValue": "Presentation and assembly system",
              "inputType": "string",
              "label": "Short description of event",
              "helpText": "",
              "choices": null,
              "weight": 115,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:30": []byte(`{
            "id": 30,
            "key": "logo_pdf_ballot_paper",
            "value": {
              "display_name": "PDF ballot paper logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:31": []byte(`{
            "id": 31,
            "key": "fonts_available",
            "value": [
              "font_regular",
              "font_italic",
              "font_bold",
              "font_bold_italic",
              "font_monospace"
            ],
            "data": null
          }`),
		"core/config:32": []byte(`{
            "id": 32,
            "key": "font_regular",
            "value": {
              "display_name": "Font regular",
              "default": "assets/fonts/fira-sans-latin-400.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:33": []byte(`{
            "id": 33,
            "key": "font_italic",
            "value": {
              "display_name": "Font italic",
              "default": "assets/fonts/fira-sans-latin-400italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:34": []byte(`{
            "id": 34,
            "key": "font_bold",
            "value": {
              "display_name": "Font bold",
              "default": "assets/fonts/fira-sans-latin-500.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:35": []byte(`{
            "id": 35,
            "key": "font_bold_italic",
            "value": {
              "display_name": "Font bold italic",
              "default": "assets/fonts/fira-sans-latin-500italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:36": []byte(`{
            "id": 36,
            "key": "font_monospace",
            "value": {
              "display_name": "Font monospace",
              "default": "assets/fonts/roboto-condensed-bold.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:37": []byte(`{
            "id": 37,
            "key": "translations",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "translations",
              "label": "Custom translations",
              "helpText": "",
              "choices": null,
              "weight": 1000,
              "group": "Custom translations",
              "subgroup": "General"
            }
          }`),
		"core/config:38": []byte(`{
            "id": 38,
            "key": "config_version",
            "value": 2,
            "data": null
          }`),
		"core/config:39": []byte(`{
            "id": 39,
            "key": "db_id",
            "value": "2c3bd736c87a48a4be1f0dc707702144",
            "data": null
          }`),
		"core/config:4": []byte(`{
            "id": 4,
            "key": "general_event_date",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event date",
              "helpText": "",
              "choices": null,
              "weight": 120,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:40": []byte(`{
            "id": 40,
            "key": "users_sort_by",
            "value": "first_name",
            "data": {
              "defaultValue": "first_name",
              "inputType": "choice",
              "label": "Sort name of participants by",
              "helpText": "",
              "choices": [
                {
                  "value": "first_name",
                  "display_name": "Given name"
                },
                {
                  "value": "last_name",
                  "display_name": "Surname"
                },
                {
                  "value": "number",
                  "display_name": "Participant number"
                }
              ],
              "weight": 510,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:41": []byte(`{
            "id": 41,
            "key": "users_enable_presence_view",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Enable participant presence view",
              "helpText": "",
              "choices": null,
              "weight": 511,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:42": []byte(`{
            "id": 42,
            "key": "users_allow_self_set_present",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow users to set themselves as present",
              "helpText": "e.g. for online meetings",
              "choices": null,
              "weight": 512,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:43": []byte(`{
            "id": 43,
            "key": "users_activate_vote_weight",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate vote weight",
              "helpText": "",
              "choices": null,
              "weight": 513,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:44": []byte(`{
            "id": 44,
            "key": "users_pdf_welcometitle",
            "value": "Welcome to OpenSlides",
            "data": {
              "defaultValue": "Welcome to OpenSlides",
              "inputType": "string",
              "label": "Title for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 520,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:45": []byte(`{
            "id": 45,
            "key": "users_pdf_welcometext",
            "value": "[Place for your welcome and help text.]",
            "data": {
              "defaultValue": "[Place for your welcome and help text.]",
              "inputType": "string",
              "label": "Help text for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 530,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:46": []byte(`{
            "id": 46,
            "key": "users_pdf_url",
            "value": "http://example.com:8000",
            "data": {
              "defaultValue": "http://example.com:8000",
              "inputType": "string",
              "label": "System URL",
              "helpText": "Used for QRCode in PDF of access data.",
              "choices": null,
              "weight": 540,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:47": []byte(`{
            "id": 47,
            "key": "users_pdf_wlan_ssid",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN name (SSID)",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 550,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:48": []byte(`{
            "id": 48,
            "key": "users_pdf_wlan_password",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN password",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 560,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:49": []byte(`{
            "id": 49,
            "key": "users_pdf_wlan_encryption",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "choice",
              "label": "WLAN encryption",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": [
                {
                  "value": "",
                  "display_name": "---------"
                },
                {
                  "value": "WEP",
                  "display_name": "WEP"
                },
                {
                  "value": "WPA",
                  "display_name": "WPA/WPA2"
                },
                {
                  "value": "nopass",
                  "display_name": "No encryption"
                }
              ],
              "weight": 570,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:5": []byte(`{
            "id": 5,
            "key": "general_event_location",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event location",
              "helpText": "",
              "choices": null,
              "weight": 125,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:50": []byte(`{
            "id": 50,
            "key": "users_email_sender",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Sender name",
              "helpText": "The sender address is defined in the OpenSlides server settings and should modified by administrator only.",
              "choices": null,
              "weight": 600,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:51": []byte(`{
            "id": 51,
            "key": "users_email_replyto",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Reply address",
              "helpText": "",
              "choices": null,
              "weight": 601,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:52": []byte(`{
            "id": 52,
            "key": "users_email_subject",
            "value": "OpenSlides access data",
            "data": {
              "defaultValue": "OpenSlides access data",
              "inputType": "string",
              "label": "Email subject",
              "helpText": "You can use {event_name} and {username} as placeholder.",
              "choices": null,
              "weight": 605,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:53": []byte(`{
            "id": 53,
            "key": "users_email_body",
            "value": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
            "data": {
              "defaultValue": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
              "inputType": "text",
              "label": "Email body",
              "helpText": "Use these placeholders: {name}, {event_name}, {url}, {username}, {password}. The url referrs to the system url.",
              "choices": null,
              "weight": 610,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:54": []byte(`{
            "id": 54,
            "key": "agenda_start_event_date_time",
            "value": null,
            "data": {
              "defaultValue": null,
              "inputType": "datetimepicker",
              "label": "Begin of event",
              "helpText": "Input format: DD.MM.YYYY HH:MM",
              "choices": null,
              "weight": 200,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:55": []byte(`{
            "id": 55,
            "key": "agenda_show_subtitle",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show subtitles in the agenda",
              "helpText": "",
              "choices": null,
              "weight": 201,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:56": []byte(`{
            "id": 56,
            "key": "agenda_enable_numbering",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Enable numbering for agenda items",
              "helpText": "",
              "choices": null,
              "weight": 205,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:57": []byte(`{
            "id": 57,
            "key": "agenda_number_prefix",
            "value": "TOP",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Numbering prefix for agenda items",
              "helpText": "This prefix will be set if you run the automatic agenda numbering.",
              "choices": null,
              "weight": 206,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:58": []byte(`{
            "id": 58,
            "key": "agenda_numeral_system",
            "value": "arabic",
            "data": {
              "defaultValue": "arabic",
              "inputType": "choice",
              "label": "Numeral system for agenda items",
              "helpText": "",
              "choices": [
                {
                  "value": "arabic",
                  "display_name": "Arabic"
                },
                {
                  "value": "roman",
                  "display_name": "Roman"
                }
              ],
              "weight": 207,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:59": []byte(`{
            "id": 59,
            "key": "agenda_item_creation",
            "value": "default_yes",
            "data": {
              "defaultValue": "default_yes",
              "inputType": "choice",
              "label": "Add to agenda",
              "helpText": "",
              "choices": [
                {
                  "value": "always",
                  "display_name": "Always"
                },
                {
                  "value": "never",
                  "display_name": "Never"
                },
                {
                  "value": "default_yes",
                  "display_name": "Ask, default yes"
                },
                {
                  "value": "default_no",
                  "display_name": "Ask, default no"
                }
              ],
              "weight": 210,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:6": []byte(`{
            "id": 6,
            "key": "general_event_legal_notice",
            "value": "<a href=\"http://www.openslides.org\">OpenSlides</a> is a free web based presentation and assembly system for visualizing and controlling agenda, motions and elections of an assembly.",
            "data": null
          }`),
		"core/config:60": []byte(`{
            "id": 60,
            "key": "agenda_new_items_default_visibility",
            "value": "2",
            "data": {
              "defaultValue": "2",
              "inputType": "choice",
              "label": "Default visibility for new agenda items (except topics)",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Public item"
                },
                {
                  "value": "2",
                  "display_name": "Internal item"
                },
                {
                  "value": "3",
                  "display_name": "Hidden item"
                }
              ],
              "weight": 211,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:61": []byte(`{
            "id": 61,
            "key": "agenda_hide_internal_items_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Hide internal items when projecting subitems",
              "helpText": "",
              "choices": null,
              "weight": 212,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:62": []byte(`{
            "id": 62,
            "key": "agenda_show_last_speakers",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of last speakers to be shown on the projector",
              "helpText": "",
              "choices": null,
              "weight": 220,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:63": []byte(`{
            "id": 63,
            "key": "agenda_show_next_speakers",
            "value": -1,
            "data": {
              "defaultValue": -1,
              "inputType": "integer",
              "label": "Number of the next speakers to be shown on the projector",
              "helpText": "Enter number of the next shown speakers. Choose -1 to show all next speakers.",
              "choices": null,
              "weight": 222,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:64": []byte(`{
            "id": 64,
            "key": "agenda_countdown_warning_time",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Show orange countdown in the last x seconds of speaking time",
              "helpText": "Enter duration in seconds. Choose 0 to disable warning color.",
              "choices": null,
              "weight": 224,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:65": []byte(`{
            "id": 65,
            "key": "projector_default_countdown",
            "value": 60,
            "data": {
              "defaultValue": 60,
              "inputType": "integer",
              "label": "Predefined seconds of new countdowns",
              "helpText": "",
              "choices": null,
              "weight": 226,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:66": []byte(`{
            "id": 66,
            "key": "agenda_couple_countdown_and_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Couple countdown with the list of speakers",
              "helpText": "[Begin speech] starts the countdown, [End speech] stops the countdown.",
              "choices": null,
              "weight": 228,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:67": []byte(`{
            "id": 67,
            "key": "agenda_hide_amount_of_speakers",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide the amount of speakers in subtitle of list of speakers slide",
              "helpText": "",
              "choices": null,
              "weight": 230,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:68": []byte(`{
            "id": 68,
            "key": "agenda_present_speakers_only",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Only present participants can be added to the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 232,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:69": []byte(`{
            "id": 69,
            "key": "agenda_show_first_contribution",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show hint »first speech« in the list of speakers management view",
              "helpText": "",
              "choices": null,
              "weight": 234,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:7": []byte(`{
            "id": 7,
            "key": "general_event_privacy_policy",
            "value": "",
            "data": null
          }`),
		"core/config:70": []byte(`{
            "id": 70,
            "key": "motions_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new motions",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 310,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:71": []byte(`{
            "id": 71,
            "key": "motions_statute_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new statute amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 312,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:72": []byte(`{
            "id": 72,
            "key": "motions_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 314,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:73": []byte(`{
            "id": 73,
            "key": "motions_preamble",
            "value": "The assembly may decide:",
            "data": {
              "defaultValue": "The assembly may decide:",
              "inputType": "string",
              "label": "Motion preamble",
              "helpText": "",
              "choices": null,
              "weight": 320,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:74": []byte(`{
            "id": 74,
            "key": "motions_default_line_numbering",
            "value": "outside",
            "data": {
              "defaultValue": "outside",
              "inputType": "choice",
              "label": "Default line numbering",
              "helpText": "",
              "choices": [
                {
                  "value": "outside",
                  "display_name": "outside"
                },
                {
                  "value": "inline",
                  "display_name": "inline"
                },
                {
                  "value": "none",
                  "display_name": "Disabled"
                }
              ],
              "weight": 322,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:75": []byte(`{
            "id": 75,
            "key": "motions_line_length",
            "value": 85,
            "data": {
              "defaultValue": 85,
              "inputType": "integer",
              "label": "Line length",
              "helpText": "The maximum number of characters per line. Relevant when line numbering is enabled. Min: 40",
              "choices": null,
              "weight": 323,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:76": []byte(`{
            "id": 76,
            "key": "motions_reason_required",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Reason required for creating new motion",
              "helpText": "",
              "choices": null,
              "weight": 324,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:77": []byte(`{
            "id": 77,
            "key": "motions_disable_text_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide motion text on projector",
              "helpText": "",
              "choices": null,
              "weight": 325,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:78": []byte(`{
            "id": 78,
            "key": "motions_disable_reason_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide reason on projector",
              "helpText": "",
              "choices": null,
              "weight": 326,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:79": []byte(`{
            "id": 79,
            "key": "motions_disable_recommendation_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide recommendation on projector",
              "helpText": "",
              "choices": null,
              "weight": 327,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:8": []byte(`{
            "id": 8,
            "key": "general_event_welcome_title",
            "value": "Welcome to OpenSlides",
            "data": null
          }`),
		"core/config:80": []byte(`{
            "id": 80,
            "key": "motions_hide_referring_motions",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide referring motions",
              "helpText": "",
              "choices": null,
              "weight": 328,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:81": []byte(`{
            "id": 81,
            "key": "motions_disable_sidebox_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show meta information box below the title on projector",
              "helpText": "",
              "choices": null,
              "weight": 329,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:82": []byte(`{
            "id": 82,
            "key": "motions_show_sequential_numbers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show the sequential number for a motion",
              "helpText": "In motion list, motion detail and PDF.",
              "choices": null,
              "weight": 330,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:83": []byte(`{
            "id": 83,
            "key": "motions_recommendations_by",
            "value": "ABK",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender",
              "helpText": "Will be displayed as label before selected recommendation. Use an empty value to disable the recommendation system.",
              "choices": null,
              "weight": 332,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:84": []byte(`{
            "id": 84,
            "key": "motions_statute_recommendations_by",
            "value": "ABK2",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender for statute amendments",
              "helpText": "Will be displayed as label before selected recommendation in statute amendments.",
              "choices": null,
              "weight": 333,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:85": []byte(`{
            "id": 85,
            "key": "motions_recommendation_text_mode",
            "value": "diff",
            "data": {
              "defaultValue": "diff",
              "inputType": "choice",
              "label": "Default text version for change recommendations",
              "helpText": "",
              "choices": [
                {
                  "value": "original",
                  "display_name": "Original version"
                },
                {
                  "value": "changed",
                  "display_name": "Changed version"
                },
                {
                  "value": "diff",
                  "display_name": "Diff version"
                },
                {
                  "value": "agreed",
                  "display_name": "Final version"
                }
              ],
              "weight": 334,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:86": []byte(`{
            "id": 86,
            "key": "motions_motions_sorting",
            "value": "identifier",
            "data": {
              "defaultValue": "identifier",
              "inputType": "choice",
              "label": "Sort motions by",
              "helpText": "",
              "choices": [
                {
                  "value": "weight",
                  "display_name": "Call list"
                },
                {
                  "value": "identifier",
                  "display_name": "Identifier"
                }
              ],
              "weight": 335,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:87": []byte(`{
            "id": 87,
            "key": "motions_identifier",
            "value": "per_category",
            "data": {
              "defaultValue": "per_category",
              "inputType": "choice",
              "label": "Identifier",
              "helpText": "",
              "choices": [
                {
                  "value": "per_category",
                  "display_name": "Numbered per category"
                },
                {
                  "value": "serially_numbered",
                  "display_name": "Serially numbered"
                },
                {
                  "value": "manually",
                  "display_name": "Set it manually"
                }
              ],
              "weight": 340,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:88": []byte(`{
            "id": 88,
            "key": "motions_identifier_min_digits",
            "value": 1,
            "data": {
              "defaultValue": 1,
              "inputType": "integer",
              "label": "Number of minimal digits for identifier",
              "helpText": "Uses leading zeros to sort motions correctly by identifier.",
              "choices": null,
              "weight": 342,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:89": []byte(`{
            "id": 89,
            "key": "motions_identifier_with_blank",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow blank in identifier",
              "helpText": "Blank between prefix and number, e.g. 'A 001'.",
              "choices": null,
              "weight": 344,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:9": []byte(`{
            "id": 9,
            "key": "general_event_welcome_text",
            "value": "[Space for your welcome text.]",
            "data": null
          }`),
		"core/config:90": []byte(`{
            "id": 90,
            "key": "motions_statutes_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate statute amendments",
              "helpText": "",
              "choices": null,
              "weight": 350,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:91": []byte(`{
            "id": 91,
            "key": "motions_amendments_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate amendments",
              "helpText": "",
              "choices": null,
              "weight": 351,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:92": []byte(`{
            "id": 92,
            "key": "motions_amendments_main_table",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show amendments together with motions",
              "helpText": "",
              "choices": null,
              "weight": 352,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:93": []byte(`{
            "id": 93,
            "key": "motions_amendments_prefix",
            "value": "Ä-",
            "data": {
              "defaultValue": "-",
              "inputType": "string",
              "label": "Prefix for the identifier for amendments",
              "helpText": "",
              "choices": null,
              "weight": 353,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:94": []byte(`{
            "id": 94,
            "key": "motions_amendments_text_mode",
            "value": "paragraph",
            "data": {
              "defaultValue": "paragraph",
              "inputType": "choice",
              "label": "How to create new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "freestyle",
                  "display_name": "Empty text field"
                },
                {
                  "value": "fulltext",
                  "display_name": "Edit the whole motion text"
                },
                {
                  "value": "paragraph",
                  "display_name": "Paragraph-based, Diff-enabled"
                }
              ],
              "weight": 354,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:95": []byte(`{
            "id": 95,
            "key": "motions_amendments_multiple_paragraphs",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Amendments can change multiple paragraphs",
              "helpText": "",
              "choices": null,
              "weight": 355,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:96": []byte(`{
            "id": 96,
            "key": "motions_amendments_of_amendments",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow amendments of amendments",
              "helpText": "",
              "choices": null,
              "weight": 356,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:97": []byte(`{
            "id": 97,
            "key": "motions_min_supporters",
            "value": 1,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of (minimum) required supporters for a motion",
              "helpText": "Choose 0 to disable the supporting system.",
              "choices": null,
              "weight": 360,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:98": []byte(`{
            "id": 98,
            "key": "motions_remove_supporters",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Remove all supporters of a motion if a submitter edits his motion in early state",
              "helpText": "",
              "choices": null,
              "weight": 361,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:99": []byte(`{
            "id": 99,
            "key": "motion_poll_default_type",
            "value": "named",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 367,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/countdown:1": []byte(`{
            "id": 1,
            "title": "Default countdown",
            "description": "",
            "default_time": 60,
            "countdown_time": 1597141855.090748,
            "running": true
          }`),
		"core/projection-default:1": []byte(`{
            "id": 1,
            "name": "agenda_all_items",
            "display_name": "Agenda",
            "projector_id": 1
          }`),
		"core/projection-default:10": []byte(`{
            "id": 10,
            "name": "messages",
            "display_name": "Messages",
            "projector_id": 1
          }`),
		"core/projection-default:11": []byte(`{
            "id": 11,
            "name": "countdowns",
            "display_name": "Countdowns",
            "projector_id": 1
          }`),
		"core/projection-default:12": []byte(`{
            "id": 12,
            "name": "assignment_poll",
            "display_name": "Ballots",
            "projector_id": 1
          }`),
		"core/projection-default:13": []byte(`{
            "id": 13,
            "name": "motion_poll",
            "display_name": "Motion votes",
            "projector_id": 1
          }`),
		"core/projection-default:2": []byte(`{
            "id": 2,
            "name": "topics",
            "display_name": "Topics",
            "projector_id": 1
          }`),
		"core/projection-default:3": []byte(`{
            "id": 3,
            "name": "agenda_list_of_speakers",
            "display_name": "List of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:4": []byte(`{
            "id": 4,
            "name": "agenda_current_list_of_speakers",
            "display_name": "Current list of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:5": []byte(`{
            "id": 5,
            "name": "motions",
            "display_name": "Motions",
            "projector_id": 1
          }`),
		"core/projection-default:6": []byte(`{
            "id": 6,
            "name": "motionBlocks",
            "display_name": "Motion blocks",
            "projector_id": 1
          }`),
		"core/projection-default:7": []byte(`{
            "id": 7,
            "name": "assignments",
            "display_name": "Elections",
            "projector_id": 1
          }`),
		"core/projection-default:8": []byte(`{
            "id": 8,
            "name": "users",
            "display_name": "Participants",
            "projector_id": 1
          }`),
		"core/projection-default:9": []byte(`{
            "id": 9,
            "name": "mediafiles",
            "display_name": "Files",
            "projector_id": 1
          }`),
		"core/projector-message:1": []byte(`{
            "id": 1,
            "message": "<p>test</p>"
          }`),
		"core/projector:1": []byte(`{
            "id": 1,
            "elements": [
              {
                "name": "mediafiles/mediafile",
                "id": 3
              }
            ],
            "elements_preview": [],
            "elements_history": [
              [
                {
                  "name": "assignments/assignment",
                  "id": 1
                }
              ]
            ],
            "scale": 0,
            "scroll": 0,
            "name": "Default projector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [
              1,
              2,
              3,
              4,
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13
            ],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/projector:2": []byte(`{
            "id": 2,
            "elements": [
              {
                "name": "core/clock",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers",
                "stable": false
              },
              {
                "name": "agenda/current-speaker-chyron",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers-overlay",
                "stable": true
              }
            ],
            "elements_preview": [],
            "elements_history": [],
            "scale": 0,
            "scroll": 0,
            "name": "sideprojector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/tag:1": []byte(`{
            "id": 1,
            "name": "T1"
          }`),
		"core/tag:2": []byte(`{
            "id": 2,
            "name": "T2"
          }`),
		"mediafiles/mediafile:2": []byte(`{
            "id": 2,
            "title": "A.txt",
            "media_url_prefix": "/media/",
            "filesize": "< 1 kB",
            "mimetype": "text/plain",
            "pdf_information": {},
            "access_groups_id": [],
            "create_timestamp": "2020-08-11T11:16:07.013124+02:00",
            "is_directory": false,
            "path": "A.txt",
            "parent_id": null,
            "list_of_speakers_id": 5,
            "inherited_access_groups_id": true
          }`),
		"motions/category:1": []byte(`{
            "id": 1,
            "name": "first",
            "prefix": "A",
            "parent_id": null,
            "weight": 2,
            "level": 0
          }`),
		"motions/category:2": []byte(`{
            "id": 2,
            "name": "second",
            "prefix": "B",
            "parent_id": 1,
            "weight": 4,
            "level": 1
          }`),
		"motions/category:3": []byte(`{
            "id": 3,
            "name": "third",
            "prefix": "C",
            "parent_id": null,
            "weight": 6,
            "level": 0
          }`),
		"motions/motion-block:1": []byte(`{
            "id": 1,
            "title": "block",
            "agenda_item_id": 8,
            "list_of_speakers_id": 12,
            "internal": false,
            "motions_id": [
              1,
              2
            ]
          }`),
		"motions/motion-option:1": []byte(`{
            "id": 1,
            "yes": "0.000000",
            "no": "1.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-option:2": []byte(`{
            "id": 2,
            "yes": "1.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 2,
            "pollstate": 4
          }`),
		"motions/motion-poll:1": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "named",
            "title": "Abstimmung",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              1
            ],
            "id": 1,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-poll:2": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "pseudoanonymous",
            "title": "Abstimmung (2)",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              2
            ],
            "id": 2,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "N",
            "user_id": 5,
            "option_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-vote:2": []byte(`{
            "id": 2,
            "weight": "1.000000",
            "value": "Y",
            "user_id": null,
            "option_id": 2,
            "pollstate": 4
          }`),
		"motions/motion:2": []byte(`{
            "id": 2,
            "identifier": "Ä-1",
            "title": "Änderungsantrag zu Leadmotion1",
            "text": "",
            "amendment_paragraphs": [
              "<p>Lorem ipsum dolor sit amet, consectedfgddfgdf&nbsp; gdfgdfg dfg dfg ww ontes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>"
            ],
            "modified_final_version": "",
            "reason": "",
            "parent_id": 1,
            "category_id": 2,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 3,
                "user_id": 1,
                "motion_id": 2,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [],
            "attachments_id": [],
            "agenda_item_id": 6,
            "list_of_speakers_id": 9,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:39:35.025914+02:00",
            "last_modified": "2020-08-11T12:41:55.666495+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/motion:3": []byte(`{
            "id": 3,
            "identifier": "2",
            "title": "Public",
            "text": "<p>a</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>a</p>",
            "parent_id": null,
            "category_id": 1,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 2,
            "origin": "",
            "submitters": [
              {
                "id": 4,
                "user_id": 1,
                "motion_id": 3,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [
              1
            ],
            "attachments_id": [],
            "agenda_item_id": 7,
            "list_of_speakers_id": 10,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T12:24:45.289233+02:00",
            "last_modified": "2020-08-11T12:41:51.319986+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/state:1": []byte(`{
            "id": 1,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              2,
              3,
              4
            ],
            "workflow_id": 1
          }`),
		"motions/state:10": []byte(`{
            "id": 10,
            "name": "withdrawed",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:11": []byte(`{
            "id": 11,
            "name": "adjourned",
            "recommendation_label": "Adjournment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:12": []byte(`{
            "id": 12,
            "name": "not concerned",
            "recommendation_label": "No concernment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:13": []byte(`{
            "id": 13,
            "name": "refered to committee",
            "recommendation_label": "Referral to committee",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:14": []byte(`{
            "id": 14,
            "name": "needs review",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:15": []byte(`{
            "id": 15,
            "name": "rejected (not authorized)",
            "recommendation_label": "Rejection (not authorized)",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:2": []byte(`{
            "id": 2,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:3": []byte(`{
            "id": 3,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:4": []byte(`{
            "id": 4,
            "name": "not decided",
            "recommendation_label": "No decision",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:5": []byte(`{
            "id": 5,
            "name": "in progress",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [
              "is_submitter"
            ],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": true,
            "dont_set_identifier": true,
            "show_state_extension_field": true,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              6,
              10
            ],
            "workflow_id": 2
          }`),
		"motions/state:6": []byte(`{
            "id": 6,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": true,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              7,
              10,
              15
            ],
            "workflow_id": 2
          }`),
		"motions/state:7": []byte(`{
            "id": 7,
            "name": "permitted",
            "recommendation_label": "Permission",
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": true,
            "next_states_id": [
              8,
              9,
              10,
              11,
              12,
              13,
              14
            ],
            "workflow_id": 2
          }`),
		"motions/state:8": []byte(`{
            "id": 8,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:9": []byte(`{
            "id": 9,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/statute-paragraph:1": []byte(`{
            "id": 1,
            "title": "Statute",
            "text": "<p>test</p>",
            "weight": 10000
          }`),
		"motions/workflow:1": []byte(`{
            "id": 1,
            "name": "Simple Workflow",
            "states_id": [
              1,
              2,
              3,
              4
            ],
            "first_state_id": 1
          }`),
		"motions/workflow:2": []byte(`{
            "id": 2,
            "name": "Complex Workflow",
            "states_id": [
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13,
              14,
              15
            ],
            "first_state_id": 5
          }`),
		"topics/topic:1": []byte(`{
            "id": 1,
            "title": "Topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 1,
            "list_of_speakers_id": 1
          }`),
		"topics/topic:2": []byte(`{
            "id": 2,
            "title": "Hidden",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 2,
            "list_of_speakers_id": 2
          }`),
		"topics/topic:3": []byte(`{
            "id": 3,
            "title": "Internal",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 3,
            "list_of_speakers_id": 3
          }`),
		"topics/topic:4": []byte(`{
            "id": 4,
            "title": "Another public topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 9,
            "list_of_speakers_id": 14
          }`),
		"users/group:1": []byte(`{
            "id": 1,
            "name": "Default",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_see",
              "users.can_change_password"
            ]
          }`),
		"users/group:2": []byte(`{
            "id": 2,
            "name": "Admin",
            "permissions": []
          }`),
		"users/group:3": []byte(`{
            "id": 3,
            "name": "Delegates",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password"
            ]
          }`),
		"users/group:4": []byte(`{
            "id": 4,
            "name": "Staff",
            "permissions": [
              "agenda.can_manage",
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_manage_list_of_speakers",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_manage",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_history",
              "core.can_manage_projector",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "core.can_manage_tags",
              "mediafiles.can_manage",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_manage",
              "motions.can_manage_metadata",
              "motions.can_see",
              "motions.can_see_internal",
              "users.can_change_password",
              "users.can_manage",
              "users.can_see_extra_data",
              "users.can_see_name"
            ]
          }`),
		"users/group:5": []byte(`{
            "id": 5,
            "name": "Committees",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password",
              "users.can_see_name"
            ]
          }`),
		"users/user:1": []byte(`{
            "first_name": "",
            "username": "admin",
            "about_me": "",
            "id": 1,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              2
            ],
            "number": "",
            "last_name": "Administrator",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:2": []byte(`{
            "first_name": "candidate1",
            "username": "candidate1",
            "about_me": "",
            "id": 2,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:3": []byte(`{
            "first_name": "candidate2",
            "username": "candidate2",
            "about_me": "",
            "id": 3,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:4": []byte(`{
            "first_name": "a",
            "username": "a",
            "about_me": "",
            "id": 4,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:5": []byte(`{
            "first_name": "b",
            "username": "b",
            "about_me": "",
            "id": 5,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:6": []byte(`{
            "first_name": "speaker1",
            "username": "speaker1",
            "about_me": "",
            "id": 6,
            "structure_level": "layer X",
            "is_present": true,
            "vote_weight": "1.000000",
            "groups_id": [],
            "title": "title",
            "email": "",
            "number": "3",
            "last_name": "the last name",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:7": []byte(`{
            "first_name": "speaker2",
            "username": "speaker2",
            "about_me": "",
            "id": 7,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
	},
	7: {
		"agenda/item:1": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            },
            "is_internal": false,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": null,
            "id": 1,
            "level": 0,
            "title_information": {
              "title": "Topic"
            },
            "closed": false,
            "weight": 2,
            "duration": null
          }`),
		"agenda/item:3": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": null,
            "id": 3,
            "level": 0,
            "title_information": {
              "title": "Internal"
            },
            "closed": false,
            "weight": 8,
            "duration": null
          }`),
		"agenda/item:5": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 5,
            "level": 1,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "closed": false,
            "weight": 14,
            "duration": null
          }`),
		"agenda/item:6": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 6,
            "level": 1,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "closed": false,
            "weight": 16,
            "duration": 0
          }`),
		"agenda/item:7": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": 6,
            "id": 7,
            "level": 2,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "closed": false,
            "weight": 18,
            "duration": null
          }`),
		"agenda/list-of-speakers:1": []byte(`{
            "id": 1,
            "title_information": {
              "title": "Topic"
            },
            "speakers": [
              {
                "id": 3,
                "user_id": 6,
                "begin_time": "2020-08-11T12:28:30.894574+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:10": []byte(`{
            "id": 10,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:11": []byte(`{
            "id": 11,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:12": []byte(`{
            "id": 12,
            "title_information": {
              "title": "block"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:13": []byte(`{
            "id": 13,
            "title_information": {
              "title": "block internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:14": []byte(`{
            "id": 14,
            "title_information": {
              "title": "Another public topic"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:2": []byte(`{
            "id": 2,
            "title_information": {
              "title": "Hidden"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:3": []byte(`{
            "id": 3,
            "title_information": {
              "title": "Internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:4": []byte(`{
            "id": 4,
            "title_information": {
              "title": "folder"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:5": []byte(`{
            "id": 5,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:6": []byte(`{
            "id": 6,
            "title_information": {
              "title": "in.jpg"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:7": []byte(`{
            "id": 7,
            "title_information": {
              "title": "Assignment"
            },
            "speakers": [
              {
                "id": 4,
                "user_id": 6,
                "begin_time": "2020-08-11T12:29:55.054553+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              },
              {
                "id": 6,
                "user_id": 7,
                "begin_time": null,
                "end_time": null,
                "weight": 2,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "assignments/assignment",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:8": []byte(`{
            "id": 8,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:9": []byte(`{
            "id": 9,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            }
          }`),
		"assignments/assignment-option:1": []byte(`{
            "user_id": 2,
            "weight": 1,
            "id": 1,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-option:2": []byte(`{
            "user_id": 3,
            "weight": 3,
            "id": 2,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-poll:1": []byte(`{
            "assignment_id": 1,
            "description": "",
            "pollmethod": "votes",
            "votes_amount": 1,
            "allow_multiple_votes_per_candidate": false,
            "global_no": true,
            "global_abstain": true,
            "state": 2,
            "type": "named",
            "title": "Wahlgang",
            "groups_id": [
              3
            ],
            "options_id": [
              1,
              2
            ],
            "id": 1,
            "onehundred_percent_base": "valid",
            "majority_method": "simple",
            "user_has_voted": false
          }`),
		"assignments/assignment:1": []byte(`{
            "id": 1,
            "title": "Assignment",
            "description": "",
            "open_posts": 1,
            "phase": 0,
            "assignment_related_users": [
              {
                "id": 1,
                "user_id": 2,
                "weight": 1
              },
              {
                "id": 3,
                "user_id": 3,
                "weight": 3
              }
            ],
            "default_poll_description": "",
            "agenda_item_id": 4,
            "list_of_speakers_id": 7,
            "tags_id": [
              2
            ],
            "attachments_id": [],
            "number_poll_candidates": false,
            "polls_id": [
              1
            ]
          }`),
		"core/config:10": []byte(`{
            "id": 10,
            "key": "general_system_conference_show",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show live conference window",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 140,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:100": []byte(`{
            "id": 100,
            "key": "motion_poll_default_100_percent_base",
            "value": "YNA",
            "data": {
              "defaultValue": "YNA",
              "inputType": "choice",
              "label": "Default 100 % base of a voting result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 370,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:101": []byte(`{
            "id": 101,
            "key": "motion_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:102": []byte(`{
            "id": 102,
            "key": "motion_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 372,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:103": []byte(`{
            "id": 103,
            "key": "motions_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 373,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:104": []byte(`{
            "id": 104,
            "key": "motions_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 374,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:105": []byte(`{
            "id": 105,
            "key": "motions_export_title",
            "value": "Motions",
            "data": {
              "defaultValue": "Motions",
              "inputType": "string",
              "label": "Title for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 380,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:106": []byte(`{
            "id": 106,
            "key": "motions_export_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 382,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:107": []byte(`{
            "id": 107,
            "key": "motions_export_submitter_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show submitters and recommendation/state in table of contents",
              "helpText": "",
              "choices": null,
              "weight": 384,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:108": []byte(`{
            "id": 108,
            "key": "motions_export_follow_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show checkbox to record decision",
              "helpText": "",
              "choices": null,
              "weight": 386,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:109": []byte(`{
            "id": 109,
            "key": "assignment_poll_method",
            "value": "votes",
            "data": {
              "defaultValue": "votes",
              "inputType": "choice",
              "label": "Default election method",
              "helpText": "",
              "choices": [
                {
                  "value": "votes",
                  "display_name": "Yes per candidate"
                },
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                }
              ],
              "weight": 400,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:11": []byte(`{
            "id": 11,
            "key": "general_system_conference_auto_connect",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Connect all users to live conference automatically",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 141,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:110": []byte(`{
            "id": 110,
            "key": "assignment_poll_default_type",
            "value": "analog",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 403,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:111": []byte(`{
            "id": 111,
            "key": "assignment_poll_default_100_percent_base",
            "value": "valid",
            "data": {
              "defaultValue": "valid",
              "inputType": "choice",
              "label": "Default 100 % base of an election result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                },
                {
                  "value": "votes",
                  "display_name": "Sum of votes including general No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 405,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:112": []byte(`{
            "id": 112,
            "key": "assignment_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 410,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:113": []byte(`{
            "id": 113,
            "key": "assignment_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:114": []byte(`{
            "id": 114,
            "key": "assignment_poll_sort_poll_result_by_votes",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Sort election results by amount of votes",
              "helpText": "",
              "choices": null,
              "weight": 420,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:115": []byte(`{
            "id": 115,
            "key": "assignment_poll_add_candidates_to_list_of_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Put all candidates on the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 425,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:116": []byte(`{
            "id": 116,
            "key": "assignments_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 430,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:117": []byte(`{
            "id": 117,
            "key": "assignments_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 435,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:118": []byte(`{
            "id": 118,
            "key": "assignments_pdf_title",
            "value": "Elections",
            "data": {
              "defaultValue": "Elections",
              "inputType": "string",
              "label": "Title for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 460,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:119": []byte(`{
            "id": 119,
            "key": "assignments_pdf_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 470,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:12": []byte(`{
            "id": 12,
            "key": "general_system_conference_los_restriction",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow only current speakers and list of speakers managers to enter the live conference",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 142,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:13": []byte(`{
            "id": 13,
            "key": "general_system_stream_url",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Livestream url",
              "helpText": "Remove URL to deactivate livestream. Check extra group permission to see livestream.",
              "choices": null,
              "weight": 143,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:14": []byte(`{
            "id": 14,
            "key": "general_system_enable_anonymous",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow access for anonymous guest users",
              "helpText": "",
              "choices": null,
              "weight": 150,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:15": []byte(`{
            "id": 15,
            "key": "general_login_info_text",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Show this text on the login page",
              "helpText": "",
              "choices": null,
              "weight": 152,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:16": []byte(`{
            "id": 16,
            "key": "openslides_theme",
            "value": "openslides-default-light-theme",
            "data": {
              "defaultValue": "openslides-default-light-theme",
              "inputType": "choice",
              "label": "OpenSlides Theme",
              "helpText": "",
              "choices": [
                {
                  "value": "openslides-default-light-theme",
                  "display_name": "OpenSlides Default"
                },
                {
                  "value": "openslides-default-dark-theme",
                  "display_name": "OpenSlides Dark"
                },
                {
                  "value": "openslides-red-light-theme",
                  "display_name": "OpenSlides Red"
                },
                {
                  "value": "openslides-red-dark-theme",
                  "display_name": "OpenSlides Red Dark"
                },
                {
                  "value": "openslides-green-light-theme",
                  "display_name": "OpenSlides Green"
                },
                {
                  "value": "openslides-green-dark-theme",
                  "display_name": "OpenSlides Green Dark"
                },
                {
                  "value": "openslides-solarized-dark-theme",
                  "display_name": "OpenSlides Solarized"
                }
              ],
              "weight": 154,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:17": []byte(`{
            "id": 17,
            "key": "general_csv_separator",
            "value": ",",
            "data": {
              "defaultValue": ",",
              "inputType": "string",
              "label": "Separator used for all csv exports and examples",
              "helpText": "",
              "choices": null,
              "weight": 160,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:18": []byte(`{
            "id": 18,
            "key": "general_csv_encoding",
            "value": "utf-8",
            "data": {
              "defaultValue": "utf-8",
              "inputType": "choice",
              "label": "Default encoding for all csv exports",
              "helpText": "",
              "choices": [
                {
                  "value": "utf-8",
                  "display_name": "UTF-8"
                },
                {
                  "value": "iso-8859-15",
                  "display_name": "ISO-8859-15"
                }
              ],
              "weight": 162,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:19": []byte(`{
            "id": 19,
            "key": "general_export_pdf_pagenumber_alignment",
            "value": "center",
            "data": {
              "defaultValue": "center",
              "inputType": "choice",
              "label": "Page number alignment in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "left",
                  "display_name": "Left"
                },
                {
                  "value": "center",
                  "display_name": "Center"
                },
                {
                  "value": "right",
                  "display_name": "Right"
                }
              ],
              "weight": 164,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:2": []byte(`{
            "id": 2,
            "key": "general_event_name",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Event name",
              "helpText": "",
              "choices": null,
              "weight": 110,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:20": []byte(`{
            "id": 20,
            "key": "general_export_pdf_fontsize",
            "value": "10",
            "data": {
              "defaultValue": "10",
              "inputType": "choice",
              "label": "Standard font size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "10",
                  "display_name": "10"
                },
                {
                  "value": "11",
                  "display_name": "11"
                },
                {
                  "value": "12",
                  "display_name": "12"
                }
              ],
              "weight": 166,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:21": []byte(`{
            "id": 21,
            "key": "general_export_pdf_pagesize",
            "value": "A4",
            "data": {
              "defaultValue": "A4",
              "inputType": "choice",
              "label": "Standard page size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "A4",
                  "display_name": "DIN A4"
                },
                {
                  "value": "A5",
                  "display_name": "DIN A5"
                }
              ],
              "weight": 168,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:22": []byte(`{
            "id": 22,
            "key": "logos_available",
            "value": [
              "logo_projector_main",
              "logo_projector_header",
              "logo_web_header",
              "logo_pdf_header_L",
              "logo_pdf_header_R",
              "logo_pdf_footer_L",
              "logo_pdf_footer_R",
              "logo_pdf_ballot_paper"
            ],
            "data": null
          }`),
		"core/config:23": []byte(`{
            "id": 23,
            "key": "logo_projector_main",
            "value": {
              "display_name": "Projector logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:24": []byte(`{
            "id": 24,
            "key": "logo_projector_header",
            "value": {
              "display_name": "Projector header image",
              "path": ""
            },
            "data": null
          }`),
		"core/config:25": []byte(`{
            "id": 25,
            "key": "logo_web_header",
            "value": {
              "display_name": "Web interface header logo",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:26": []byte(`{
            "id": 26,
            "key": "logo_pdf_header_L",
            "value": {
              "display_name": "PDF header logo (left)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:27": []byte(`{
            "id": 27,
            "key": "logo_pdf_header_R",
            "value": {
              "display_name": "PDF header logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:28": []byte(`{
            "id": 28,
            "key": "logo_pdf_footer_L",
            "value": {
              "display_name": "PDF footer logo (left)",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:29": []byte(`{
            "id": 29,
            "key": "logo_pdf_footer_R",
            "value": {
              "display_name": "PDF footer logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:3": []byte(`{
            "id": 3,
            "key": "general_event_description",
            "value": "Presentation and assembly system",
            "data": {
              "defaultValue": "Presentation and assembly system",
              "inputType": "string",
              "label": "Short description of event",
              "helpText": "",
              "choices": null,
              "weight": 115,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:30": []byte(`{
            "id": 30,
            "key": "logo_pdf_ballot_paper",
            "value": {
              "display_name": "PDF ballot paper logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:31": []byte(`{
            "id": 31,
            "key": "fonts_available",
            "value": [
              "font_regular",
              "font_italic",
              "font_bold",
              "font_bold_italic",
              "font_monospace"
            ],
            "data": null
          }`),
		"core/config:32": []byte(`{
            "id": 32,
            "key": "font_regular",
            "value": {
              "display_name": "Font regular",
              "default": "assets/fonts/fira-sans-latin-400.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:33": []byte(`{
            "id": 33,
            "key": "font_italic",
            "value": {
              "display_name": "Font italic",
              "default": "assets/fonts/fira-sans-latin-400italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:34": []byte(`{
            "id": 34,
            "key": "font_bold",
            "value": {
              "display_name": "Font bold",
              "default": "assets/fonts/fira-sans-latin-500.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:35": []byte(`{
            "id": 35,
            "key": "font_bold_italic",
            "value": {
              "display_name": "Font bold italic",
              "default": "assets/fonts/fira-sans-latin-500italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:36": []byte(`{
            "id": 36,
            "key": "font_monospace",
            "value": {
              "display_name": "Font monospace",
              "default": "assets/fonts/roboto-condensed-bold.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:37": []byte(`{
            "id": 37,
            "key": "translations",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "translations",
              "label": "Custom translations",
              "helpText": "",
              "choices": null,
              "weight": 1000,
              "group": "Custom translations",
              "subgroup": "General"
            }
          }`),
		"core/config:38": []byte(`{
            "id": 38,
            "key": "config_version",
            "value": 2,
            "data": null
          }`),
		"core/config:39": []byte(`{
            "id": 39,
            "key": "db_id",
            "value": "2c3bd736c87a48a4be1f0dc707702144",
            "data": null
          }`),
		"core/config:4": []byte(`{
            "id": 4,
            "key": "general_event_date",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event date",
              "helpText": "",
              "choices": null,
              "weight": 120,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:40": []byte(`{
            "id": 40,
            "key": "users_sort_by",
            "value": "first_name",
            "data": {
              "defaultValue": "first_name",
              "inputType": "choice",
              "label": "Sort name of participants by",
              "helpText": "",
              "choices": [
                {
                  "value": "first_name",
                  "display_name": "Given name"
                },
                {
                  "value": "last_name",
                  "display_name": "Surname"
                },
                {
                  "value": "number",
                  "display_name": "Participant number"
                }
              ],
              "weight": 510,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:41": []byte(`{
            "id": 41,
            "key": "users_enable_presence_view",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Enable participant presence view",
              "helpText": "",
              "choices": null,
              "weight": 511,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:42": []byte(`{
            "id": 42,
            "key": "users_allow_self_set_present",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow users to set themselves as present",
              "helpText": "e.g. for online meetings",
              "choices": null,
              "weight": 512,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:43": []byte(`{
            "id": 43,
            "key": "users_activate_vote_weight",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate vote weight",
              "helpText": "",
              "choices": null,
              "weight": 513,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:44": []byte(`{
            "id": 44,
            "key": "users_pdf_welcometitle",
            "value": "Welcome to OpenSlides",
            "data": {
              "defaultValue": "Welcome to OpenSlides",
              "inputType": "string",
              "label": "Title for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 520,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:45": []byte(`{
            "id": 45,
            "key": "users_pdf_welcometext",
            "value": "[Place for your welcome and help text.]",
            "data": {
              "defaultValue": "[Place for your welcome and help text.]",
              "inputType": "string",
              "label": "Help text for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 530,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:46": []byte(`{
            "id": 46,
            "key": "users_pdf_url",
            "value": "http://example.com:8000",
            "data": {
              "defaultValue": "http://example.com:8000",
              "inputType": "string",
              "label": "System URL",
              "helpText": "Used for QRCode in PDF of access data.",
              "choices": null,
              "weight": 540,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:47": []byte(`{
            "id": 47,
            "key": "users_pdf_wlan_ssid",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN name (SSID)",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 550,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:48": []byte(`{
            "id": 48,
            "key": "users_pdf_wlan_password",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN password",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 560,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:49": []byte(`{
            "id": 49,
            "key": "users_pdf_wlan_encryption",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "choice",
              "label": "WLAN encryption",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": [
                {
                  "value": "",
                  "display_name": "---------"
                },
                {
                  "value": "WEP",
                  "display_name": "WEP"
                },
                {
                  "value": "WPA",
                  "display_name": "WPA/WPA2"
                },
                {
                  "value": "nopass",
                  "display_name": "No encryption"
                }
              ],
              "weight": 570,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:5": []byte(`{
            "id": 5,
            "key": "general_event_location",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event location",
              "helpText": "",
              "choices": null,
              "weight": 125,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:50": []byte(`{
            "id": 50,
            "key": "users_email_sender",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Sender name",
              "helpText": "The sender address is defined in the OpenSlides server settings and should modified by administrator only.",
              "choices": null,
              "weight": 600,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:51": []byte(`{
            "id": 51,
            "key": "users_email_replyto",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Reply address",
              "helpText": "",
              "choices": null,
              "weight": 601,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:52": []byte(`{
            "id": 52,
            "key": "users_email_subject",
            "value": "OpenSlides access data",
            "data": {
              "defaultValue": "OpenSlides access data",
              "inputType": "string",
              "label": "Email subject",
              "helpText": "You can use {event_name} and {username} as placeholder.",
              "choices": null,
              "weight": 605,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:53": []byte(`{
            "id": 53,
            "key": "users_email_body",
            "value": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
            "data": {
              "defaultValue": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
              "inputType": "text",
              "label": "Email body",
              "helpText": "Use these placeholders: {name}, {event_name}, {url}, {username}, {password}. The url referrs to the system url.",
              "choices": null,
              "weight": 610,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:54": []byte(`{
            "id": 54,
            "key": "agenda_start_event_date_time",
            "value": null,
            "data": {
              "defaultValue": null,
              "inputType": "datetimepicker",
              "label": "Begin of event",
              "helpText": "Input format: DD.MM.YYYY HH:MM",
              "choices": null,
              "weight": 200,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:55": []byte(`{
            "id": 55,
            "key": "agenda_show_subtitle",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show subtitles in the agenda",
              "helpText": "",
              "choices": null,
              "weight": 201,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:56": []byte(`{
            "id": 56,
            "key": "agenda_enable_numbering",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Enable numbering for agenda items",
              "helpText": "",
              "choices": null,
              "weight": 205,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:57": []byte(`{
            "id": 57,
            "key": "agenda_number_prefix",
            "value": "TOP",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Numbering prefix for agenda items",
              "helpText": "This prefix will be set if you run the automatic agenda numbering.",
              "choices": null,
              "weight": 206,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:58": []byte(`{
            "id": 58,
            "key": "agenda_numeral_system",
            "value": "arabic",
            "data": {
              "defaultValue": "arabic",
              "inputType": "choice",
              "label": "Numeral system for agenda items",
              "helpText": "",
              "choices": [
                {
                  "value": "arabic",
                  "display_name": "Arabic"
                },
                {
                  "value": "roman",
                  "display_name": "Roman"
                }
              ],
              "weight": 207,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:59": []byte(`{
            "id": 59,
            "key": "agenda_item_creation",
            "value": "default_yes",
            "data": {
              "defaultValue": "default_yes",
              "inputType": "choice",
              "label": "Add to agenda",
              "helpText": "",
              "choices": [
                {
                  "value": "always",
                  "display_name": "Always"
                },
                {
                  "value": "never",
                  "display_name": "Never"
                },
                {
                  "value": "default_yes",
                  "display_name": "Ask, default yes"
                },
                {
                  "value": "default_no",
                  "display_name": "Ask, default no"
                }
              ],
              "weight": 210,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:6": []byte(`{
            "id": 6,
            "key": "general_event_legal_notice",
            "value": "<a href=\"http://www.openslides.org\">OpenSlides</a> is a free web based presentation and assembly system for visualizing and controlling agenda, motions and elections of an assembly.",
            "data": null
          }`),
		"core/config:60": []byte(`{
            "id": 60,
            "key": "agenda_new_items_default_visibility",
            "value": "2",
            "data": {
              "defaultValue": "2",
              "inputType": "choice",
              "label": "Default visibility for new agenda items (except topics)",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Public item"
                },
                {
                  "value": "2",
                  "display_name": "Internal item"
                },
                {
                  "value": "3",
                  "display_name": "Hidden item"
                }
              ],
              "weight": 211,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:61": []byte(`{
            "id": 61,
            "key": "agenda_hide_internal_items_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Hide internal items when projecting subitems",
              "helpText": "",
              "choices": null,
              "weight": 212,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:62": []byte(`{
            "id": 62,
            "key": "agenda_show_last_speakers",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of last speakers to be shown on the projector",
              "helpText": "",
              "choices": null,
              "weight": 220,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:63": []byte(`{
            "id": 63,
            "key": "agenda_show_next_speakers",
            "value": -1,
            "data": {
              "defaultValue": -1,
              "inputType": "integer",
              "label": "Number of the next speakers to be shown on the projector",
              "helpText": "Enter number of the next shown speakers. Choose -1 to show all next speakers.",
              "choices": null,
              "weight": 222,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:64": []byte(`{
            "id": 64,
            "key": "agenda_countdown_warning_time",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Show orange countdown in the last x seconds of speaking time",
              "helpText": "Enter duration in seconds. Choose 0 to disable warning color.",
              "choices": null,
              "weight": 224,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:65": []byte(`{
            "id": 65,
            "key": "projector_default_countdown",
            "value": 60,
            "data": {
              "defaultValue": 60,
              "inputType": "integer",
              "label": "Predefined seconds of new countdowns",
              "helpText": "",
              "choices": null,
              "weight": 226,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:66": []byte(`{
            "id": 66,
            "key": "agenda_couple_countdown_and_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Couple countdown with the list of speakers",
              "helpText": "[Begin speech] starts the countdown, [End speech] stops the countdown.",
              "choices": null,
              "weight": 228,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:67": []byte(`{
            "id": 67,
            "key": "agenda_hide_amount_of_speakers",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide the amount of speakers in subtitle of list of speakers slide",
              "helpText": "",
              "choices": null,
              "weight": 230,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:68": []byte(`{
            "id": 68,
            "key": "agenda_present_speakers_only",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Only present participants can be added to the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 232,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:69": []byte(`{
            "id": 69,
            "key": "agenda_show_first_contribution",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show hint »first speech« in the list of speakers management view",
              "helpText": "",
              "choices": null,
              "weight": 234,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:7": []byte(`{
            "id": 7,
            "key": "general_event_privacy_policy",
            "value": "",
            "data": null
          }`),
		"core/config:70": []byte(`{
            "id": 70,
            "key": "motions_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new motions",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 310,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:71": []byte(`{
            "id": 71,
            "key": "motions_statute_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new statute amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 312,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:72": []byte(`{
            "id": 72,
            "key": "motions_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 314,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:73": []byte(`{
            "id": 73,
            "key": "motions_preamble",
            "value": "The assembly may decide:",
            "data": {
              "defaultValue": "The assembly may decide:",
              "inputType": "string",
              "label": "Motion preamble",
              "helpText": "",
              "choices": null,
              "weight": 320,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:74": []byte(`{
            "id": 74,
            "key": "motions_default_line_numbering",
            "value": "outside",
            "data": {
              "defaultValue": "outside",
              "inputType": "choice",
              "label": "Default line numbering",
              "helpText": "",
              "choices": [
                {
                  "value": "outside",
                  "display_name": "outside"
                },
                {
                  "value": "inline",
                  "display_name": "inline"
                },
                {
                  "value": "none",
                  "display_name": "Disabled"
                }
              ],
              "weight": 322,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:75": []byte(`{
            "id": 75,
            "key": "motions_line_length",
            "value": 85,
            "data": {
              "defaultValue": 85,
              "inputType": "integer",
              "label": "Line length",
              "helpText": "The maximum number of characters per line. Relevant when line numbering is enabled. Min: 40",
              "choices": null,
              "weight": 323,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:76": []byte(`{
            "id": 76,
            "key": "motions_reason_required",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Reason required for creating new motion",
              "helpText": "",
              "choices": null,
              "weight": 324,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:77": []byte(`{
            "id": 77,
            "key": "motions_disable_text_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide motion text on projector",
              "helpText": "",
              "choices": null,
              "weight": 325,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:78": []byte(`{
            "id": 78,
            "key": "motions_disable_reason_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide reason on projector",
              "helpText": "",
              "choices": null,
              "weight": 326,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:79": []byte(`{
            "id": 79,
            "key": "motions_disable_recommendation_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide recommendation on projector",
              "helpText": "",
              "choices": null,
              "weight": 327,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:8": []byte(`{
            "id": 8,
            "key": "general_event_welcome_title",
            "value": "Welcome to OpenSlides",
            "data": null
          }`),
		"core/config:80": []byte(`{
            "id": 80,
            "key": "motions_hide_referring_motions",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide referring motions",
              "helpText": "",
              "choices": null,
              "weight": 328,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:81": []byte(`{
            "id": 81,
            "key": "motions_disable_sidebox_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show meta information box below the title on projector",
              "helpText": "",
              "choices": null,
              "weight": 329,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:82": []byte(`{
            "id": 82,
            "key": "motions_show_sequential_numbers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show the sequential number for a motion",
              "helpText": "In motion list, motion detail and PDF.",
              "choices": null,
              "weight": 330,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:83": []byte(`{
            "id": 83,
            "key": "motions_recommendations_by",
            "value": "ABK",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender",
              "helpText": "Will be displayed as label before selected recommendation. Use an empty value to disable the recommendation system.",
              "choices": null,
              "weight": 332,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:84": []byte(`{
            "id": 84,
            "key": "motions_statute_recommendations_by",
            "value": "ABK2",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender for statute amendments",
              "helpText": "Will be displayed as label before selected recommendation in statute amendments.",
              "choices": null,
              "weight": 333,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:85": []byte(`{
            "id": 85,
            "key": "motions_recommendation_text_mode",
            "value": "diff",
            "data": {
              "defaultValue": "diff",
              "inputType": "choice",
              "label": "Default text version for change recommendations",
              "helpText": "",
              "choices": [
                {
                  "value": "original",
                  "display_name": "Original version"
                },
                {
                  "value": "changed",
                  "display_name": "Changed version"
                },
                {
                  "value": "diff",
                  "display_name": "Diff version"
                },
                {
                  "value": "agreed",
                  "display_name": "Final version"
                }
              ],
              "weight": 334,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:86": []byte(`{
            "id": 86,
            "key": "motions_motions_sorting",
            "value": "identifier",
            "data": {
              "defaultValue": "identifier",
              "inputType": "choice",
              "label": "Sort motions by",
              "helpText": "",
              "choices": [
                {
                  "value": "weight",
                  "display_name": "Call list"
                },
                {
                  "value": "identifier",
                  "display_name": "Identifier"
                }
              ],
              "weight": 335,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:87": []byte(`{
            "id": 87,
            "key": "motions_identifier",
            "value": "per_category",
            "data": {
              "defaultValue": "per_category",
              "inputType": "choice",
              "label": "Identifier",
              "helpText": "",
              "choices": [
                {
                  "value": "per_category",
                  "display_name": "Numbered per category"
                },
                {
                  "value": "serially_numbered",
                  "display_name": "Serially numbered"
                },
                {
                  "value": "manually",
                  "display_name": "Set it manually"
                }
              ],
              "weight": 340,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:88": []byte(`{
            "id": 88,
            "key": "motions_identifier_min_digits",
            "value": 1,
            "data": {
              "defaultValue": 1,
              "inputType": "integer",
              "label": "Number of minimal digits for identifier",
              "helpText": "Uses leading zeros to sort motions correctly by identifier.",
              "choices": null,
              "weight": 342,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:89": []byte(`{
            "id": 89,
            "key": "motions_identifier_with_blank",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow blank in identifier",
              "helpText": "Blank between prefix and number, e.g. 'A 001'.",
              "choices": null,
              "weight": 344,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:9": []byte(`{
            "id": 9,
            "key": "general_event_welcome_text",
            "value": "[Space for your welcome text.]",
            "data": null
          }`),
		"core/config:90": []byte(`{
            "id": 90,
            "key": "motions_statutes_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate statute amendments",
              "helpText": "",
              "choices": null,
              "weight": 350,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:91": []byte(`{
            "id": 91,
            "key": "motions_amendments_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate amendments",
              "helpText": "",
              "choices": null,
              "weight": 351,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:92": []byte(`{
            "id": 92,
            "key": "motions_amendments_main_table",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show amendments together with motions",
              "helpText": "",
              "choices": null,
              "weight": 352,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:93": []byte(`{
            "id": 93,
            "key": "motions_amendments_prefix",
            "value": "Ä-",
            "data": {
              "defaultValue": "-",
              "inputType": "string",
              "label": "Prefix for the identifier for amendments",
              "helpText": "",
              "choices": null,
              "weight": 353,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:94": []byte(`{
            "id": 94,
            "key": "motions_amendments_text_mode",
            "value": "paragraph",
            "data": {
              "defaultValue": "paragraph",
              "inputType": "choice",
              "label": "How to create new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "freestyle",
                  "display_name": "Empty text field"
                },
                {
                  "value": "fulltext",
                  "display_name": "Edit the whole motion text"
                },
                {
                  "value": "paragraph",
                  "display_name": "Paragraph-based, Diff-enabled"
                }
              ],
              "weight": 354,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:95": []byte(`{
            "id": 95,
            "key": "motions_amendments_multiple_paragraphs",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Amendments can change multiple paragraphs",
              "helpText": "",
              "choices": null,
              "weight": 355,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:96": []byte(`{
            "id": 96,
            "key": "motions_amendments_of_amendments",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow amendments of amendments",
              "helpText": "",
              "choices": null,
              "weight": 356,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:97": []byte(`{
            "id": 97,
            "key": "motions_min_supporters",
            "value": 1,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of (minimum) required supporters for a motion",
              "helpText": "Choose 0 to disable the supporting system.",
              "choices": null,
              "weight": 360,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:98": []byte(`{
            "id": 98,
            "key": "motions_remove_supporters",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Remove all supporters of a motion if a submitter edits his motion in early state",
              "helpText": "",
              "choices": null,
              "weight": 361,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:99": []byte(`{
            "id": 99,
            "key": "motion_poll_default_type",
            "value": "named",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 367,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/countdown:1": []byte(`{
            "id": 1,
            "title": "Default countdown",
            "description": "",
            "default_time": 60,
            "countdown_time": 1597141855.090748,
            "running": true
          }`),
		"core/projection-default:1": []byte(`{
            "id": 1,
            "name": "agenda_all_items",
            "display_name": "Agenda",
            "projector_id": 1
          }`),
		"core/projection-default:10": []byte(`{
            "id": 10,
            "name": "messages",
            "display_name": "Messages",
            "projector_id": 1
          }`),
		"core/projection-default:11": []byte(`{
            "id": 11,
            "name": "countdowns",
            "display_name": "Countdowns",
            "projector_id": 1
          }`),
		"core/projection-default:12": []byte(`{
            "id": 12,
            "name": "assignment_poll",
            "display_name": "Ballots",
            "projector_id": 1
          }`),
		"core/projection-default:13": []byte(`{
            "id": 13,
            "name": "motion_poll",
            "display_name": "Motion votes",
            "projector_id": 1
          }`),
		"core/projection-default:2": []byte(`{
            "id": 2,
            "name": "topics",
            "display_name": "Topics",
            "projector_id": 1
          }`),
		"core/projection-default:3": []byte(`{
            "id": 3,
            "name": "agenda_list_of_speakers",
            "display_name": "List of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:4": []byte(`{
            "id": 4,
            "name": "agenda_current_list_of_speakers",
            "display_name": "Current list of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:5": []byte(`{
            "id": 5,
            "name": "motions",
            "display_name": "Motions",
            "projector_id": 1
          }`),
		"core/projection-default:6": []byte(`{
            "id": 6,
            "name": "motionBlocks",
            "display_name": "Motion blocks",
            "projector_id": 1
          }`),
		"core/projection-default:7": []byte(`{
            "id": 7,
            "name": "assignments",
            "display_name": "Elections",
            "projector_id": 1
          }`),
		"core/projection-default:8": []byte(`{
            "id": 8,
            "name": "users",
            "display_name": "Participants",
            "projector_id": 1
          }`),
		"core/projection-default:9": []byte(`{
            "id": 9,
            "name": "mediafiles",
            "display_name": "Files",
            "projector_id": 1
          }`),
		"core/projector-message:1": []byte(`{
            "id": 1,
            "message": "<p>test</p>"
          }`),
		"core/projector:1": []byte(`{
            "id": 1,
            "elements": [
              {
                "name": "mediafiles/mediafile",
                "id": 3
              }
            ],
            "elements_preview": [],
            "elements_history": [
              [
                {
                  "name": "assignments/assignment",
                  "id": 1
                }
              ]
            ],
            "scale": 0,
            "scroll": 0,
            "name": "Default projector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [
              1,
              2,
              3,
              4,
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13
            ],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/projector:2": []byte(`{
            "id": 2,
            "elements": [
              {
                "name": "core/clock",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers",
                "stable": false
              },
              {
                "name": "agenda/current-speaker-chyron",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers-overlay",
                "stable": true
              }
            ],
            "elements_preview": [],
            "elements_history": [],
            "scale": 0,
            "scroll": 0,
            "name": "sideprojector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/tag:1": []byte(`{
            "id": 1,
            "name": "T1"
          }`),
		"core/tag:2": []byte(`{
            "id": 2,
            "name": "T2"
          }`),
		"mediafiles/mediafile:2": []byte(`{
            "id": 2,
            "title": "A.txt",
            "media_url_prefix": "/media/",
            "filesize": "< 1 kB",
            "mimetype": "text/plain",
            "pdf_information": {},
            "access_groups_id": [],
            "create_timestamp": "2020-08-11T11:16:07.013124+02:00",
            "is_directory": false,
            "path": "A.txt",
            "parent_id": null,
            "list_of_speakers_id": 5,
            "inherited_access_groups_id": true
          }`),
		"motions/category:1": []byte(`{
            "id": 1,
            "name": "first",
            "prefix": "A",
            "parent_id": null,
            "weight": 2,
            "level": 0
          }`),
		"motions/category:2": []byte(`{
            "id": 2,
            "name": "second",
            "prefix": "B",
            "parent_id": 1,
            "weight": 4,
            "level": 1
          }`),
		"motions/category:3": []byte(`{
            "id": 3,
            "name": "third",
            "prefix": "C",
            "parent_id": null,
            "weight": 6,
            "level": 0
          }`),
		"motions/motion-block:1": []byte(`{
            "id": 1,
            "title": "block",
            "agenda_item_id": 8,
            "list_of_speakers_id": 12,
            "internal": false,
            "motions_id": [
              1,
              2
            ]
          }`),
		"motions/motion-option:1": []byte(`{
            "id": 1,
            "yes": "0.000000",
            "no": "1.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-option:2": []byte(`{
            "id": 2,
            "yes": "1.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 2,
            "pollstate": 4
          }`),
		"motions/motion-poll:1": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "named",
            "title": "Abstimmung",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              1
            ],
            "id": 1,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-poll:2": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "pseudoanonymous",
            "title": "Abstimmung (2)",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              2
            ],
            "id": 2,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "N",
            "user_id": 5,
            "option_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-vote:2": []byte(`{
            "id": 2,
            "weight": "1.000000",
            "value": "Y",
            "user_id": null,
            "option_id": 2,
            "pollstate": 4
          }`),
		"motions/motion:2": []byte(`{
            "id": 2,
            "identifier": "Ä-1",
            "title": "Änderungsantrag zu Leadmotion1",
            "text": "",
            "amendment_paragraphs": [
              "<p>Lorem ipsum dolor sit amet, consectedfgddfgdf&nbsp; gdfgdfg dfg dfg ww ontes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>"
            ],
            "modified_final_version": "",
            "reason": "",
            "parent_id": 1,
            "category_id": 2,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 3,
                "user_id": 1,
                "motion_id": 2,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [],
            "attachments_id": [],
            "agenda_item_id": 6,
            "list_of_speakers_id": 9,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:39:35.025914+02:00",
            "last_modified": "2020-08-11T12:41:55.666495+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/motion:3": []byte(`{
            "id": 3,
            "identifier": "2",
            "title": "Public",
            "text": "<p>a</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>a</p>",
            "parent_id": null,
            "category_id": 1,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 2,
            "origin": "",
            "submitters": [
              {
                "id": 4,
                "user_id": 1,
                "motion_id": 3,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [
              1
            ],
            "attachments_id": [],
            "agenda_item_id": 7,
            "list_of_speakers_id": 10,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T12:24:45.289233+02:00",
            "last_modified": "2020-08-11T12:41:51.319986+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/state:1": []byte(`{
            "id": 1,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              2,
              3,
              4
            ],
            "workflow_id": 1
          }`),
		"motions/state:10": []byte(`{
            "id": 10,
            "name": "withdrawed",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:11": []byte(`{
            "id": 11,
            "name": "adjourned",
            "recommendation_label": "Adjournment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:12": []byte(`{
            "id": 12,
            "name": "not concerned",
            "recommendation_label": "No concernment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:13": []byte(`{
            "id": 13,
            "name": "refered to committee",
            "recommendation_label": "Referral to committee",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:14": []byte(`{
            "id": 14,
            "name": "needs review",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:15": []byte(`{
            "id": 15,
            "name": "rejected (not authorized)",
            "recommendation_label": "Rejection (not authorized)",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:2": []byte(`{
            "id": 2,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:3": []byte(`{
            "id": 3,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:4": []byte(`{
            "id": 4,
            "name": "not decided",
            "recommendation_label": "No decision",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:5": []byte(`{
            "id": 5,
            "name": "in progress",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [
              "is_submitter"
            ],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": true,
            "dont_set_identifier": true,
            "show_state_extension_field": true,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              6,
              10
            ],
            "workflow_id": 2
          }`),
		"motions/state:6": []byte(`{
            "id": 6,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": true,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              7,
              10,
              15
            ],
            "workflow_id": 2
          }`),
		"motions/state:7": []byte(`{
            "id": 7,
            "name": "permitted",
            "recommendation_label": "Permission",
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": true,
            "next_states_id": [
              8,
              9,
              10,
              11,
              12,
              13,
              14
            ],
            "workflow_id": 2
          }`),
		"motions/state:8": []byte(`{
            "id": 8,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:9": []byte(`{
            "id": 9,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/statute-paragraph:1": []byte(`{
            "id": 1,
            "title": "Statute",
            "text": "<p>test</p>",
            "weight": 10000
          }`),
		"motions/workflow:1": []byte(`{
            "id": 1,
            "name": "Simple Workflow",
            "states_id": [
              1,
              2,
              3,
              4
            ],
            "first_state_id": 1
          }`),
		"motions/workflow:2": []byte(`{
            "id": 2,
            "name": "Complex Workflow",
            "states_id": [
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13,
              14,
              15
            ],
            "first_state_id": 5
          }`),
		"topics/topic:1": []byte(`{
            "id": 1,
            "title": "Topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 1,
            "list_of_speakers_id": 1
          }`),
		"topics/topic:2": []byte(`{
            "id": 2,
            "title": "Hidden",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 2,
            "list_of_speakers_id": 2
          }`),
		"topics/topic:3": []byte(`{
            "id": 3,
            "title": "Internal",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 3,
            "list_of_speakers_id": 3
          }`),
		"topics/topic:4": []byte(`{
            "id": 4,
            "title": "Another public topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 9,
            "list_of_speakers_id": 14
          }`),
		"users/group:1": []byte(`{
            "id": 1,
            "name": "Default",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_see",
              "users.can_change_password"
            ]
          }`),
		"users/group:2": []byte(`{
            "id": 2,
            "name": "Admin",
            "permissions": []
          }`),
		"users/group:3": []byte(`{
            "id": 3,
            "name": "Delegates",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password"
            ]
          }`),
		"users/group:4": []byte(`{
            "id": 4,
            "name": "Staff",
            "permissions": [
              "agenda.can_manage",
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_manage_list_of_speakers",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_manage",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_history",
              "core.can_manage_projector",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "core.can_manage_tags",
              "mediafiles.can_manage",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_manage",
              "motions.can_manage_metadata",
              "motions.can_see",
              "motions.can_see_internal",
              "users.can_change_password",
              "users.can_manage",
              "users.can_see_extra_data",
              "users.can_see_name"
            ]
          }`),
		"users/group:5": []byte(`{
            "id": 5,
            "name": "Committees",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password",
              "users.can_see_name"
            ]
          }`),
		"users/user:1": []byte(`{
            "first_name": "",
            "username": "admin",
            "about_me": "",
            "id": 1,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              2
            ],
            "number": "",
            "last_name": "Administrator",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:2": []byte(`{
            "first_name": "candidate1",
            "username": "candidate1",
            "about_me": "",
            "id": 2,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:3": []byte(`{
            "first_name": "candidate2",
            "username": "candidate2",
            "about_me": "",
            "id": 3,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:4": []byte(`{
            "first_name": "a",
            "username": "a",
            "about_me": "",
            "id": 4,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:5": []byte(`{
            "first_name": "b",
            "username": "b",
            "about_me": "",
            "id": 5,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:6": []byte(`{
            "first_name": "speaker1",
            "username": "speaker1",
            "about_me": "",
            "id": 6,
            "structure_level": "layer X",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "title",
            "groups_id": [],
            "number": "3",
            "last_name": "the last name",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:7": []byte(`{
            "first_name": "speaker2",
            "username": "speaker2",
            "about_me": "",
            "id": 7,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "groups_id": [],
            "title": "",
            "email": "",
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
	},
	8: {
		"agenda/item:1": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            },
            "is_internal": false,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": null,
            "id": 1,
            "level": 0,
            "title_information": {
              "title": "Topic"
            },
            "closed": false,
            "weight": 2,
            "duration": null
          }`),
		"agenda/item:3": []byte(`{
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": null,
            "id": 3,
            "level": 0,
            "title_information": {
              "title": "Internal"
            },
            "closed": false,
            "weight": 8,
            "duration": null
          }`),
		"agenda/item:5": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 5,
            "level": 1,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "closed": false,
            "weight": 14,
            "duration": null
          }`),
		"agenda/item:6": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 1,
            "is_hidden": false,
            "parent_id": 3,
            "id": 6,
            "level": 1,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "closed": false,
            "weight": 16,
            "duration": 0
          }`),
		"agenda/item:7": []byte(`{
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            },
            "is_internal": true,
            "tags_id": [],
            "item_number": "",
            "type": 2,
            "is_hidden": false,
            "parent_id": 6,
            "id": 7,
            "level": 2,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "closed": false,
            "weight": 18,
            "duration": null
          }`),
		"agenda/list-of-speakers:1": []byte(`{
            "id": 1,
            "title_information": {
              "title": "Topic"
            },
            "speakers": [
              {
                "id": 3,
                "user_id": 6,
                "begin_time": "2020-08-11T12:28:30.894574+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:10": []byte(`{
            "id": 10,
            "title_information": {
              "title": "Public",
              "identifier": "2"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:11": []byte(`{
            "id": 11,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:12": []byte(`{
            "id": 12,
            "title_information": {
              "title": "block"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:13": []byte(`{
            "id": 13,
            "title_information": {
              "title": "block internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion-block",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:14": []byte(`{
            "id": 14,
            "title_information": {
              "title": "Another public topic"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 4
            }
          }`),
		"agenda/list-of-speakers:2": []byte(`{
            "id": 2,
            "title_information": {
              "title": "Hidden"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:3": []byte(`{
            "id": 3,
            "title_information": {
              "title": "Internal"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "topics/topic",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:4": []byte(`{
            "id": 4,
            "title_information": {
              "title": "folder"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:5": []byte(`{
            "id": 5,
            "title_information": {
              "title": "A.txt"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 2
            }
          }`),
		"agenda/list-of-speakers:6": []byte(`{
            "id": 6,
            "title_information": {
              "title": "in.jpg"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "mediafiles/mediafile",
              "id": 3
            }
          }`),
		"agenda/list-of-speakers:7": []byte(`{
            "id": 7,
            "title_information": {
              "title": "Assignment"
            },
            "speakers": [
              {
                "id": 4,
                "user_id": 6,
                "begin_time": "2020-08-11T12:29:55.054553+02:00",
                "end_time": null,
                "weight": null,
                "marked": false
              },
              {
                "id": 6,
                "user_id": 7,
                "begin_time": null,
                "end_time": null,
                "weight": 2,
                "marked": false
              }
            ],
            "closed": false,
            "content_object": {
              "collection": "assignments/assignment",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:8": []byte(`{
            "id": 8,
            "title_information": {
              "title": "Leadmotion1",
              "identifier": null
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 1
            }
          }`),
		"agenda/list-of-speakers:9": []byte(`{
            "id": 9,
            "title_information": {
              "title": "Änderungsantrag zu Leadmotion1",
              "identifier": "Ä-1"
            },
            "speakers": [],
            "closed": false,
            "content_object": {
              "collection": "motions/motion",
              "id": 2
            }
          }`),
		"assignments/assignment-option:1": []byte(`{
            "user_id": 2,
            "weight": 1,
            "id": 1,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-option:2": []byte(`{
            "user_id": 3,
            "weight": 3,
            "id": 2,
            "poll_id": 1,
            "pollstate": 2
          }`),
		"assignments/assignment-poll:1": []byte(`{
            "assignment_id": 1,
            "description": "",
            "pollmethod": "votes",
            "votes_amount": 1,
            "allow_multiple_votes_per_candidate": false,
            "global_no": true,
            "global_abstain": true,
            "state": 2,
            "type": "named",
            "title": "Wahlgang",
            "groups_id": [
              3
            ],
            "options_id": [
              1,
              2
            ],
            "id": 1,
            "onehundred_percent_base": "valid",
            "majority_method": "simple",
            "user_has_voted": false
          }`),
		"assignments/assignment:1": []byte(`{
            "id": 1,
            "title": "Assignment",
            "description": "",
            "open_posts": 1,
            "phase": 0,
            "assignment_related_users": [
              {
                "id": 1,
                "user_id": 2,
                "weight": 1
              },
              {
                "id": 3,
                "user_id": 3,
                "weight": 3
              }
            ],
            "default_poll_description": "",
            "agenda_item_id": 4,
            "list_of_speakers_id": 7,
            "tags_id": [
              2
            ],
            "attachments_id": [],
            "number_poll_candidates": false,
            "polls_id": [
              1
            ]
          }`),
		"core/config:10": []byte(`{
            "id": 10,
            "key": "general_system_conference_show",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show live conference window",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 140,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:100": []byte(`{
            "id": 100,
            "key": "motion_poll_default_100_percent_base",
            "value": "YNA",
            "data": {
              "defaultValue": "YNA",
              "inputType": "choice",
              "label": "Default 100 % base of a voting result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 370,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:101": []byte(`{
            "id": 101,
            "key": "motion_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:102": []byte(`{
            "id": 102,
            "key": "motion_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 372,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:103": []byte(`{
            "id": 103,
            "key": "motions_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 373,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:104": []byte(`{
            "id": 104,
            "key": "motions_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 374,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/config:105": []byte(`{
            "id": 105,
            "key": "motions_export_title",
            "value": "Motions",
            "data": {
              "defaultValue": "Motions",
              "inputType": "string",
              "label": "Title for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 380,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:106": []byte(`{
            "id": 106,
            "key": "motions_export_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF documents of motions",
              "helpText": "",
              "choices": null,
              "weight": 382,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:107": []byte(`{
            "id": 107,
            "key": "motions_export_submitter_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show submitters and recommendation/state in table of contents",
              "helpText": "",
              "choices": null,
              "weight": 384,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:108": []byte(`{
            "id": 108,
            "key": "motions_export_follow_recommendation",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show checkbox to record decision",
              "helpText": "",
              "choices": null,
              "weight": 386,
              "group": "Motions",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:109": []byte(`{
            "id": 109,
            "key": "assignment_poll_method",
            "value": "votes",
            "data": {
              "defaultValue": "votes",
              "inputType": "choice",
              "label": "Default election method",
              "helpText": "",
              "choices": [
                {
                  "value": "votes",
                  "display_name": "Yes per candidate"
                },
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                }
              ],
              "weight": 400,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:11": []byte(`{
            "id": 11,
            "key": "general_system_conference_auto_connect",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Connect all users to live conference automatically",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 141,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:110": []byte(`{
            "id": 110,
            "key": "assignment_poll_default_type",
            "value": "analog",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 403,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:111": []byte(`{
            "id": 111,
            "key": "assignment_poll_default_100_percent_base",
            "value": "valid",
            "data": {
              "defaultValue": "valid",
              "inputType": "choice",
              "label": "Default 100 % base of an election result",
              "helpText": "",
              "choices": [
                {
                  "value": "YN",
                  "display_name": "Yes/No per candidate"
                },
                {
                  "value": "YNA",
                  "display_name": "Yes/No/Abstain per candidate"
                },
                {
                  "value": "votes",
                  "display_name": "Sum of votes including general No/Abstain"
                },
                {
                  "value": "valid",
                  "display_name": "All valid ballots"
                },
                {
                  "value": "cast",
                  "display_name": "All casted ballots"
                },
                {
                  "value": "disabled",
                  "display_name": "Disabled (no percents)"
                }
              ],
              "weight": 405,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:112": []byte(`{
            "id": 112,
            "key": "assignment_poll_default_groups",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "groups",
              "label": "Default groups with voting rights",
              "helpText": "",
              "choices": null,
              "weight": 410,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:113": []byte(`{
            "id": 113,
            "key": "assignment_poll_default_majority_method",
            "value": "simple",
            "data": null
          }`),
		"core/config:114": []byte(`{
            "id": 114,
            "key": "assignment_poll_sort_poll_result_by_votes",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Sort election results by amount of votes",
              "helpText": "",
              "choices": null,
              "weight": 420,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:115": []byte(`{
            "id": 115,
            "key": "assignment_poll_add_candidates_to_list_of_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Put all candidates on the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 425,
              "group": "Elections",
              "subgroup": "Ballot"
            }
          }`),
		"core/config:116": []byte(`{
            "id": 116,
            "key": "assignments_pdf_ballot_papers_selection",
            "value": "CUSTOM_NUMBER",
            "data": {
              "defaultValue": "CUSTOM_NUMBER",
              "inputType": "choice",
              "label": "Number of ballot papers",
              "helpText": "",
              "choices": [
                {
                  "value": "NUMBER_OF_DELEGATES",
                  "display_name": "Number of all delegates"
                },
                {
                  "value": "NUMBER_OF_ALL_PARTICIPANTS",
                  "display_name": "Number of all participants"
                },
                {
                  "value": "CUSTOM_NUMBER",
                  "display_name": "Use the following custom number"
                }
              ],
              "weight": 430,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:117": []byte(`{
            "id": 117,
            "key": "assignments_pdf_ballot_papers_number",
            "value": 8,
            "data": {
              "defaultValue": 8,
              "inputType": "integer",
              "label": "Custom number of ballot papers",
              "helpText": "",
              "choices": null,
              "weight": 435,
              "group": "Elections",
              "subgroup": "Ballot papers"
            }
          }`),
		"core/config:118": []byte(`{
            "id": 118,
            "key": "assignments_pdf_title",
            "value": "Elections",
            "data": {
              "defaultValue": "Elections",
              "inputType": "string",
              "label": "Title for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 460,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:119": []byte(`{
            "id": 119,
            "key": "assignments_pdf_preamble",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Preamble text for PDF document (all elections)",
              "helpText": "",
              "choices": null,
              "weight": 470,
              "group": "Elections",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:12": []byte(`{
            "id": 12,
            "key": "general_system_conference_los_restriction",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow only current speakers and list of speakers managers to enter the live conference",
              "helpText": "Server settings required to activate Jitsi Meet integration.",
              "choices": null,
              "weight": 142,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:13": []byte(`{
            "id": 13,
            "key": "general_system_stream_url",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Livestream url",
              "helpText": "Remove URL to deactivate livestream. Check extra group permission to see livestream.",
              "choices": null,
              "weight": 143,
              "group": "General",
              "subgroup": "Live conference"
            }
          }`),
		"core/config:14": []byte(`{
            "id": 14,
            "key": "general_system_enable_anonymous",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow access for anonymous guest users",
              "helpText": "",
              "choices": null,
              "weight": 150,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:15": []byte(`{
            "id": 15,
            "key": "general_login_info_text",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Show this text on the login page",
              "helpText": "",
              "choices": null,
              "weight": 152,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:16": []byte(`{
            "id": 16,
            "key": "openslides_theme",
            "value": "openslides-default-light-theme",
            "data": {
              "defaultValue": "openslides-default-light-theme",
              "inputType": "choice",
              "label": "OpenSlides Theme",
              "helpText": "",
              "choices": [
                {
                  "value": "openslides-default-light-theme",
                  "display_name": "OpenSlides Default"
                },
                {
                  "value": "openslides-default-dark-theme",
                  "display_name": "OpenSlides Dark"
                },
                {
                  "value": "openslides-red-light-theme",
                  "display_name": "OpenSlides Red"
                },
                {
                  "value": "openslides-red-dark-theme",
                  "display_name": "OpenSlides Red Dark"
                },
                {
                  "value": "openslides-green-light-theme",
                  "display_name": "OpenSlides Green"
                },
                {
                  "value": "openslides-green-dark-theme",
                  "display_name": "OpenSlides Green Dark"
                },
                {
                  "value": "openslides-solarized-dark-theme",
                  "display_name": "OpenSlides Solarized"
                }
              ],
              "weight": 154,
              "group": "General",
              "subgroup": "System"
            }
          }`),
		"core/config:17": []byte(`{
            "id": 17,
            "key": "general_csv_separator",
            "value": ",",
            "data": {
              "defaultValue": ",",
              "inputType": "string",
              "label": "Separator used for all csv exports and examples",
              "helpText": "",
              "choices": null,
              "weight": 160,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:18": []byte(`{
            "id": 18,
            "key": "general_csv_encoding",
            "value": "utf-8",
            "data": {
              "defaultValue": "utf-8",
              "inputType": "choice",
              "label": "Default encoding for all csv exports",
              "helpText": "",
              "choices": [
                {
                  "value": "utf-8",
                  "display_name": "UTF-8"
                },
                {
                  "value": "iso-8859-15",
                  "display_name": "ISO-8859-15"
                }
              ],
              "weight": 162,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:19": []byte(`{
            "id": 19,
            "key": "general_export_pdf_pagenumber_alignment",
            "value": "center",
            "data": {
              "defaultValue": "center",
              "inputType": "choice",
              "label": "Page number alignment in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "left",
                  "display_name": "Left"
                },
                {
                  "value": "center",
                  "display_name": "Center"
                },
                {
                  "value": "right",
                  "display_name": "Right"
                }
              ],
              "weight": 164,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:2": []byte(`{
            "id": 2,
            "key": "general_event_name",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Event name",
              "helpText": "",
              "choices": null,
              "weight": 110,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:20": []byte(`{
            "id": 20,
            "key": "general_export_pdf_fontsize",
            "value": "10",
            "data": {
              "defaultValue": "10",
              "inputType": "choice",
              "label": "Standard font size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "10",
                  "display_name": "10"
                },
                {
                  "value": "11",
                  "display_name": "11"
                },
                {
                  "value": "12",
                  "display_name": "12"
                }
              ],
              "weight": 166,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:21": []byte(`{
            "id": 21,
            "key": "general_export_pdf_pagesize",
            "value": "A4",
            "data": {
              "defaultValue": "A4",
              "inputType": "choice",
              "label": "Standard page size in PDF",
              "helpText": "",
              "choices": [
                {
                  "value": "A4",
                  "display_name": "DIN A4"
                },
                {
                  "value": "A5",
                  "display_name": "DIN A5"
                }
              ],
              "weight": 168,
              "group": "General",
              "subgroup": "Export"
            }
          }`),
		"core/config:22": []byte(`{
            "id": 22,
            "key": "logos_available",
            "value": [
              "logo_projector_main",
              "logo_projector_header",
              "logo_web_header",
              "logo_pdf_header_L",
              "logo_pdf_header_R",
              "logo_pdf_footer_L",
              "logo_pdf_footer_R",
              "logo_pdf_ballot_paper"
            ],
            "data": null
          }`),
		"core/config:23": []byte(`{
            "id": 23,
            "key": "logo_projector_main",
            "value": {
              "display_name": "Projector logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:24": []byte(`{
            "id": 24,
            "key": "logo_projector_header",
            "value": {
              "display_name": "Projector header image",
              "path": ""
            },
            "data": null
          }`),
		"core/config:25": []byte(`{
            "id": 25,
            "key": "logo_web_header",
            "value": {
              "display_name": "Web interface header logo",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:26": []byte(`{
            "id": 26,
            "key": "logo_pdf_header_L",
            "value": {
              "display_name": "PDF header logo (left)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:27": []byte(`{
            "id": 27,
            "key": "logo_pdf_header_R",
            "value": {
              "display_name": "PDF header logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:28": []byte(`{
            "id": 28,
            "key": "logo_pdf_footer_L",
            "value": {
              "display_name": "PDF footer logo (left)",
              "path": "/media/folder/in.jpg"
            },
            "data": null
          }`),
		"core/config:29": []byte(`{
            "id": 29,
            "key": "logo_pdf_footer_R",
            "value": {
              "display_name": "PDF footer logo (right)",
              "path": ""
            },
            "data": null
          }`),
		"core/config:3": []byte(`{
            "id": 3,
            "key": "general_event_description",
            "value": "Presentation and assembly system",
            "data": {
              "defaultValue": "Presentation and assembly system",
              "inputType": "string",
              "label": "Short description of event",
              "helpText": "",
              "choices": null,
              "weight": 115,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:30": []byte(`{
            "id": 30,
            "key": "logo_pdf_ballot_paper",
            "value": {
              "display_name": "PDF ballot paper logo",
              "path": ""
            },
            "data": null
          }`),
		"core/config:31": []byte(`{
            "id": 31,
            "key": "fonts_available",
            "value": [
              "font_regular",
              "font_italic",
              "font_bold",
              "font_bold_italic",
              "font_monospace"
            ],
            "data": null
          }`),
		"core/config:32": []byte(`{
            "id": 32,
            "key": "font_regular",
            "value": {
              "display_name": "Font regular",
              "default": "assets/fonts/fira-sans-latin-400.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:33": []byte(`{
            "id": 33,
            "key": "font_italic",
            "value": {
              "display_name": "Font italic",
              "default": "assets/fonts/fira-sans-latin-400italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:34": []byte(`{
            "id": 34,
            "key": "font_bold",
            "value": {
              "display_name": "Font bold",
              "default": "assets/fonts/fira-sans-latin-500.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:35": []byte(`{
            "id": 35,
            "key": "font_bold_italic",
            "value": {
              "display_name": "Font bold italic",
              "default": "assets/fonts/fira-sans-latin-500italic.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:36": []byte(`{
            "id": 36,
            "key": "font_monospace",
            "value": {
              "display_name": "Font monospace",
              "default": "assets/fonts/roboto-condensed-bold.woff",
              "path": ""
            },
            "data": null
          }`),
		"core/config:37": []byte(`{
            "id": 37,
            "key": "translations",
            "value": [],
            "data": {
              "defaultValue": [],
              "inputType": "translations",
              "label": "Custom translations",
              "helpText": "",
              "choices": null,
              "weight": 1000,
              "group": "Custom translations",
              "subgroup": "General"
            }
          }`),
		"core/config:38": []byte(`{
            "id": 38,
            "key": "config_version",
            "value": 2,
            "data": null
          }`),
		"core/config:39": []byte(`{
            "id": 39,
            "key": "db_id",
            "value": "2c3bd736c87a48a4be1f0dc707702144",
            "data": null
          }`),
		"core/config:4": []byte(`{
            "id": 4,
            "key": "general_event_date",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event date",
              "helpText": "",
              "choices": null,
              "weight": 120,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:40": []byte(`{
            "id": 40,
            "key": "users_sort_by",
            "value": "first_name",
            "data": {
              "defaultValue": "first_name",
              "inputType": "choice",
              "label": "Sort name of participants by",
              "helpText": "",
              "choices": [
                {
                  "value": "first_name",
                  "display_name": "Given name"
                },
                {
                  "value": "last_name",
                  "display_name": "Surname"
                },
                {
                  "value": "number",
                  "display_name": "Participant number"
                }
              ],
              "weight": 510,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:41": []byte(`{
            "id": 41,
            "key": "users_enable_presence_view",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Enable participant presence view",
              "helpText": "",
              "choices": null,
              "weight": 511,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:42": []byte(`{
            "id": 42,
            "key": "users_allow_self_set_present",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow users to set themselves as present",
              "helpText": "e.g. for online meetings",
              "choices": null,
              "weight": 512,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:43": []byte(`{
            "id": 43,
            "key": "users_activate_vote_weight",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate vote weight",
              "helpText": "",
              "choices": null,
              "weight": 513,
              "group": "Participants",
              "subgroup": "General"
            }
          }`),
		"core/config:44": []byte(`{
            "id": 44,
            "key": "users_pdf_welcometitle",
            "value": "Welcome to OpenSlides",
            "data": {
              "defaultValue": "Welcome to OpenSlides",
              "inputType": "string",
              "label": "Title for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 520,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:45": []byte(`{
            "id": 45,
            "key": "users_pdf_welcometext",
            "value": "[Place for your welcome and help text.]",
            "data": {
              "defaultValue": "[Place for your welcome and help text.]",
              "inputType": "string",
              "label": "Help text for access data and welcome PDF",
              "helpText": "",
              "choices": null,
              "weight": 530,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:46": []byte(`{
            "id": 46,
            "key": "users_pdf_url",
            "value": "http://example.com:8000",
            "data": {
              "defaultValue": "http://example.com:8000",
              "inputType": "string",
              "label": "System URL",
              "helpText": "Used for QRCode in PDF of access data.",
              "choices": null,
              "weight": 540,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:47": []byte(`{
            "id": 47,
            "key": "users_pdf_wlan_ssid",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN name (SSID)",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 550,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:48": []byte(`{
            "id": 48,
            "key": "users_pdf_wlan_password",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "WLAN password",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": null,
              "weight": 560,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:49": []byte(`{
            "id": 49,
            "key": "users_pdf_wlan_encryption",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "choice",
              "label": "WLAN encryption",
              "helpText": "Used for WLAN QRCode in PDF of access data.",
              "choices": [
                {
                  "value": "",
                  "display_name": "---------"
                },
                {
                  "value": "WEP",
                  "display_name": "WEP"
                },
                {
                  "value": "WPA",
                  "display_name": "WPA/WPA2"
                },
                {
                  "value": "nopass",
                  "display_name": "No encryption"
                }
              ],
              "weight": 570,
              "group": "Participants",
              "subgroup": "PDF export"
            }
          }`),
		"core/config:5": []byte(`{
            "id": 5,
            "key": "general_event_location",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Event location",
              "helpText": "",
              "choices": null,
              "weight": 125,
              "group": "General",
              "subgroup": "Event"
            }
          }`),
		"core/config:50": []byte(`{
            "id": 50,
            "key": "users_email_sender",
            "value": "OpenSlides",
            "data": {
              "defaultValue": "OpenSlides",
              "inputType": "string",
              "label": "Sender name",
              "helpText": "The sender address is defined in the OpenSlides server settings and should modified by administrator only.",
              "choices": null,
              "weight": 600,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:51": []byte(`{
            "id": 51,
            "key": "users_email_replyto",
            "value": "",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Reply address",
              "helpText": "",
              "choices": null,
              "weight": 601,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:52": []byte(`{
            "id": 52,
            "key": "users_email_subject",
            "value": "OpenSlides access data",
            "data": {
              "defaultValue": "OpenSlides access data",
              "inputType": "string",
              "label": "Email subject",
              "helpText": "You can use {event_name} and {username} as placeholder.",
              "choices": null,
              "weight": 605,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:53": []byte(`{
            "id": 53,
            "key": "users_email_body",
            "value": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
            "data": {
              "defaultValue": "Dear {name},\n\nthis is your personal OpenSlides login:\n\n    {url}\n    username: {username}\n    password: {password}\n\nThis email was generated automatically.",
              "inputType": "text",
              "label": "Email body",
              "helpText": "Use these placeholders: {name}, {event_name}, {url}, {username}, {password}. The url referrs to the system url.",
              "choices": null,
              "weight": 610,
              "group": "Participants",
              "subgroup": "Email"
            }
          }`),
		"core/config:54": []byte(`{
            "id": 54,
            "key": "agenda_start_event_date_time",
            "value": null,
            "data": {
              "defaultValue": null,
              "inputType": "datetimepicker",
              "label": "Begin of event",
              "helpText": "Input format: DD.MM.YYYY HH:MM",
              "choices": null,
              "weight": 200,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:55": []byte(`{
            "id": 55,
            "key": "agenda_show_subtitle",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show subtitles in the agenda",
              "helpText": "",
              "choices": null,
              "weight": 201,
              "group": "Agenda",
              "subgroup": "General"
            }
          }`),
		"core/config:56": []byte(`{
            "id": 56,
            "key": "agenda_enable_numbering",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Enable numbering for agenda items",
              "helpText": "",
              "choices": null,
              "weight": 205,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:57": []byte(`{
            "id": 57,
            "key": "agenda_number_prefix",
            "value": "TOP",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Numbering prefix for agenda items",
              "helpText": "This prefix will be set if you run the automatic agenda numbering.",
              "choices": null,
              "weight": 206,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:58": []byte(`{
            "id": 58,
            "key": "agenda_numeral_system",
            "value": "arabic",
            "data": {
              "defaultValue": "arabic",
              "inputType": "choice",
              "label": "Numeral system for agenda items",
              "helpText": "",
              "choices": [
                {
                  "value": "arabic",
                  "display_name": "Arabic"
                },
                {
                  "value": "roman",
                  "display_name": "Roman"
                }
              ],
              "weight": 207,
              "group": "Agenda",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:59": []byte(`{
            "id": 59,
            "key": "agenda_item_creation",
            "value": "default_yes",
            "data": {
              "defaultValue": "default_yes",
              "inputType": "choice",
              "label": "Add to agenda",
              "helpText": "",
              "choices": [
                {
                  "value": "always",
                  "display_name": "Always"
                },
                {
                  "value": "never",
                  "display_name": "Never"
                },
                {
                  "value": "default_yes",
                  "display_name": "Ask, default yes"
                },
                {
                  "value": "default_no",
                  "display_name": "Ask, default no"
                }
              ],
              "weight": 210,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:6": []byte(`{
            "id": 6,
            "key": "general_event_legal_notice",
            "value": "<a href=\"http://www.openslides.org\">OpenSlides</a> is a free web based presentation and assembly system for visualizing and controlling agenda, motions and elections of an assembly.",
            "data": null
          }`),
		"core/config:60": []byte(`{
            "id": 60,
            "key": "agenda_new_items_default_visibility",
            "value": "2",
            "data": {
              "defaultValue": "2",
              "inputType": "choice",
              "label": "Default visibility for new agenda items (except topics)",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Public item"
                },
                {
                  "value": "2",
                  "display_name": "Internal item"
                },
                {
                  "value": "3",
                  "display_name": "Hidden item"
                }
              ],
              "weight": 211,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:61": []byte(`{
            "id": 61,
            "key": "agenda_hide_internal_items_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Hide internal items when projecting subitems",
              "helpText": "",
              "choices": null,
              "weight": 212,
              "group": "Agenda",
              "subgroup": "Visibility"
            }
          }`),
		"core/config:62": []byte(`{
            "id": 62,
            "key": "agenda_show_last_speakers",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of last speakers to be shown on the projector",
              "helpText": "",
              "choices": null,
              "weight": 220,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:63": []byte(`{
            "id": 63,
            "key": "agenda_show_next_speakers",
            "value": -1,
            "data": {
              "defaultValue": -1,
              "inputType": "integer",
              "label": "Number of the next speakers to be shown on the projector",
              "helpText": "Enter number of the next shown speakers. Choose -1 to show all next speakers.",
              "choices": null,
              "weight": 222,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:64": []byte(`{
            "id": 64,
            "key": "agenda_countdown_warning_time",
            "value": 0,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Show orange countdown in the last x seconds of speaking time",
              "helpText": "Enter duration in seconds. Choose 0 to disable warning color.",
              "choices": null,
              "weight": 224,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:65": []byte(`{
            "id": 65,
            "key": "projector_default_countdown",
            "value": 60,
            "data": {
              "defaultValue": 60,
              "inputType": "integer",
              "label": "Predefined seconds of new countdowns",
              "helpText": "",
              "choices": null,
              "weight": 226,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:66": []byte(`{
            "id": 66,
            "key": "agenda_couple_countdown_and_speakers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Couple countdown with the list of speakers",
              "helpText": "[Begin speech] starts the countdown, [End speech] stops the countdown.",
              "choices": null,
              "weight": 228,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:67": []byte(`{
            "id": 67,
            "key": "agenda_hide_amount_of_speakers",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide the amount of speakers in subtitle of list of speakers slide",
              "helpText": "",
              "choices": null,
              "weight": 230,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:68": []byte(`{
            "id": 68,
            "key": "agenda_present_speakers_only",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Only present participants can be added to the list of speakers",
              "helpText": "",
              "choices": null,
              "weight": 232,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:69": []byte(`{
            "id": 69,
            "key": "agenda_show_first_contribution",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Show hint »first speech« in the list of speakers management view",
              "helpText": "",
              "choices": null,
              "weight": 234,
              "group": "Agenda",
              "subgroup": "List of speakers"
            }
          }`),
		"core/config:7": []byte(`{
            "id": 7,
            "key": "general_event_privacy_policy",
            "value": "",
            "data": null
          }`),
		"core/config:70": []byte(`{
            "id": 70,
            "key": "motions_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new motions",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 310,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:71": []byte(`{
            "id": 71,
            "key": "motions_statute_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new statute amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 312,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:72": []byte(`{
            "id": 72,
            "key": "motions_amendments_workflow",
            "value": "1",
            "data": {
              "defaultValue": "1",
              "inputType": "choice",
              "label": "Workflow of new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "1",
                  "display_name": "Simple Workflow"
                },
                {
                  "value": "2",
                  "display_name": "Complex Workflow"
                }
              ],
              "weight": 314,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:73": []byte(`{
            "id": 73,
            "key": "motions_preamble",
            "value": "The assembly may decide:",
            "data": {
              "defaultValue": "The assembly may decide:",
              "inputType": "string",
              "label": "Motion preamble",
              "helpText": "",
              "choices": null,
              "weight": 320,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:74": []byte(`{
            "id": 74,
            "key": "motions_default_line_numbering",
            "value": "outside",
            "data": {
              "defaultValue": "outside",
              "inputType": "choice",
              "label": "Default line numbering",
              "helpText": "",
              "choices": [
                {
                  "value": "outside",
                  "display_name": "outside"
                },
                {
                  "value": "inline",
                  "display_name": "inline"
                },
                {
                  "value": "none",
                  "display_name": "Disabled"
                }
              ],
              "weight": 322,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:75": []byte(`{
            "id": 75,
            "key": "motions_line_length",
            "value": 85,
            "data": {
              "defaultValue": 85,
              "inputType": "integer",
              "label": "Line length",
              "helpText": "The maximum number of characters per line. Relevant when line numbering is enabled. Min: 40",
              "choices": null,
              "weight": 323,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:76": []byte(`{
            "id": 76,
            "key": "motions_reason_required",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Reason required for creating new motion",
              "helpText": "",
              "choices": null,
              "weight": 324,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:77": []byte(`{
            "id": 77,
            "key": "motions_disable_text_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide motion text on projector",
              "helpText": "",
              "choices": null,
              "weight": 325,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:78": []byte(`{
            "id": 78,
            "key": "motions_disable_reason_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide reason on projector",
              "helpText": "",
              "choices": null,
              "weight": 326,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:79": []byte(`{
            "id": 79,
            "key": "motions_disable_recommendation_on_projector",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide recommendation on projector",
              "helpText": "",
              "choices": null,
              "weight": 327,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:8": []byte(`{
            "id": 8,
            "key": "general_event_welcome_title",
            "value": "Welcome to OpenSlides",
            "data": null
          }`),
		"core/config:80": []byte(`{
            "id": 80,
            "key": "motions_hide_referring_motions",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Hide referring motions",
              "helpText": "",
              "choices": null,
              "weight": 328,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:81": []byte(`{
            "id": 81,
            "key": "motions_disable_sidebox_on_projector",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show meta information box below the title on projector",
              "helpText": "",
              "choices": null,
              "weight": 329,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:82": []byte(`{
            "id": 82,
            "key": "motions_show_sequential_numbers",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show the sequential number for a motion",
              "helpText": "In motion list, motion detail and PDF.",
              "choices": null,
              "weight": 330,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:83": []byte(`{
            "id": 83,
            "key": "motions_recommendations_by",
            "value": "ABK",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender",
              "helpText": "Will be displayed as label before selected recommendation. Use an empty value to disable the recommendation system.",
              "choices": null,
              "weight": 332,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:84": []byte(`{
            "id": 84,
            "key": "motions_statute_recommendations_by",
            "value": "ABK2",
            "data": {
              "defaultValue": "",
              "inputType": "string",
              "label": "Name of recommender for statute amendments",
              "helpText": "Will be displayed as label before selected recommendation in statute amendments.",
              "choices": null,
              "weight": 333,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:85": []byte(`{
            "id": 85,
            "key": "motions_recommendation_text_mode",
            "value": "diff",
            "data": {
              "defaultValue": "diff",
              "inputType": "choice",
              "label": "Default text version for change recommendations",
              "helpText": "",
              "choices": [
                {
                  "value": "original",
                  "display_name": "Original version"
                },
                {
                  "value": "changed",
                  "display_name": "Changed version"
                },
                {
                  "value": "diff",
                  "display_name": "Diff version"
                },
                {
                  "value": "agreed",
                  "display_name": "Final version"
                }
              ],
              "weight": 334,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:86": []byte(`{
            "id": 86,
            "key": "motions_motions_sorting",
            "value": "identifier",
            "data": {
              "defaultValue": "identifier",
              "inputType": "choice",
              "label": "Sort motions by",
              "helpText": "",
              "choices": [
                {
                  "value": "weight",
                  "display_name": "Call list"
                },
                {
                  "value": "identifier",
                  "display_name": "Identifier"
                }
              ],
              "weight": 335,
              "group": "Motions",
              "subgroup": "General"
            }
          }`),
		"core/config:87": []byte(`{
            "id": 87,
            "key": "motions_identifier",
            "value": "per_category",
            "data": {
              "defaultValue": "per_category",
              "inputType": "choice",
              "label": "Identifier",
              "helpText": "",
              "choices": [
                {
                  "value": "per_category",
                  "display_name": "Numbered per category"
                },
                {
                  "value": "serially_numbered",
                  "display_name": "Serially numbered"
                },
                {
                  "value": "manually",
                  "display_name": "Set it manually"
                }
              ],
              "weight": 340,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:88": []byte(`{
            "id": 88,
            "key": "motions_identifier_min_digits",
            "value": 1,
            "data": {
              "defaultValue": 1,
              "inputType": "integer",
              "label": "Number of minimal digits for identifier",
              "helpText": "Uses leading zeros to sort motions correctly by identifier.",
              "choices": null,
              "weight": 342,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:89": []byte(`{
            "id": 89,
            "key": "motions_identifier_with_blank",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow blank in identifier",
              "helpText": "Blank between prefix and number, e.g. 'A 001'.",
              "choices": null,
              "weight": 344,
              "group": "Motions",
              "subgroup": "Numbering"
            }
          }`),
		"core/config:9": []byte(`{
            "id": 9,
            "key": "general_event_welcome_text",
            "value": "[Space for your welcome text.]",
            "data": null
          }`),
		"core/config:90": []byte(`{
            "id": 90,
            "key": "motions_statutes_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate statute amendments",
              "helpText": "",
              "choices": null,
              "weight": 350,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:91": []byte(`{
            "id": 91,
            "key": "motions_amendments_enabled",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Activate amendments",
              "helpText": "",
              "choices": null,
              "weight": 351,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:92": []byte(`{
            "id": 92,
            "key": "motions_amendments_main_table",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Show amendments together with motions",
              "helpText": "",
              "choices": null,
              "weight": 352,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:93": []byte(`{
            "id": 93,
            "key": "motions_amendments_prefix",
            "value": "Ä-",
            "data": {
              "defaultValue": "-",
              "inputType": "string",
              "label": "Prefix for the identifier for amendments",
              "helpText": "",
              "choices": null,
              "weight": 353,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:94": []byte(`{
            "id": 94,
            "key": "motions_amendments_text_mode",
            "value": "paragraph",
            "data": {
              "defaultValue": "paragraph",
              "inputType": "choice",
              "label": "How to create new amendments",
              "helpText": "",
              "choices": [
                {
                  "value": "freestyle",
                  "display_name": "Empty text field"
                },
                {
                  "value": "fulltext",
                  "display_name": "Edit the whole motion text"
                },
                {
                  "value": "paragraph",
                  "display_name": "Paragraph-based, Diff-enabled"
                }
              ],
              "weight": 354,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:95": []byte(`{
            "id": 95,
            "key": "motions_amendments_multiple_paragraphs",
            "value": true,
            "data": {
              "defaultValue": true,
              "inputType": "boolean",
              "label": "Amendments can change multiple paragraphs",
              "helpText": "",
              "choices": null,
              "weight": 355,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:96": []byte(`{
            "id": 96,
            "key": "motions_amendments_of_amendments",
            "value": true,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Allow amendments of amendments",
              "helpText": "",
              "choices": null,
              "weight": 356,
              "group": "Motions",
              "subgroup": "Amendments"
            }
          }`),
		"core/config:97": []byte(`{
            "id": 97,
            "key": "motions_min_supporters",
            "value": 1,
            "data": {
              "defaultValue": 0,
              "inputType": "integer",
              "label": "Number of (minimum) required supporters for a motion",
              "helpText": "Choose 0 to disable the supporting system.",
              "choices": null,
              "weight": 360,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:98": []byte(`{
            "id": 98,
            "key": "motions_remove_supporters",
            "value": false,
            "data": {
              "defaultValue": false,
              "inputType": "boolean",
              "label": "Remove all supporters of a motion if a submitter edits his motion in early state",
              "helpText": "",
              "choices": null,
              "weight": 361,
              "group": "Motions",
              "subgroup": "Supporters"
            }
          }`),
		"core/config:99": []byte(`{
            "id": 99,
            "key": "motion_poll_default_type",
            "value": "named",
            "data": {
              "defaultValue": "analog",
              "inputType": "choice",
              "label": "Default voting type",
              "helpText": "",
              "choices": [
                {
                  "value": "analog",
                  "display_name": "analog"
                },
                {
                  "value": "named",
                  "display_name": "nominal"
                },
                {
                  "value": "pseudoanonymous",
                  "display_name": "non-nominal"
                }
              ],
              "weight": 367,
              "group": "Motions",
              "subgroup": "Voting and ballot papers"
            }
          }`),
		"core/countdown:1": []byte(`{
            "id": 1,
            "title": "Default countdown",
            "description": "",
            "default_time": 60,
            "countdown_time": 1597141855.090748,
            "running": true
          }`),
		"core/projection-default:1": []byte(`{
            "id": 1,
            "name": "agenda_all_items",
            "display_name": "Agenda",
            "projector_id": 1
          }`),
		"core/projection-default:10": []byte(`{
            "id": 10,
            "name": "messages",
            "display_name": "Messages",
            "projector_id": 1
          }`),
		"core/projection-default:11": []byte(`{
            "id": 11,
            "name": "countdowns",
            "display_name": "Countdowns",
            "projector_id": 1
          }`),
		"core/projection-default:12": []byte(`{
            "id": 12,
            "name": "assignment_poll",
            "display_name": "Ballots",
            "projector_id": 1
          }`),
		"core/projection-default:13": []byte(`{
            "id": 13,
            "name": "motion_poll",
            "display_name": "Motion votes",
            "projector_id": 1
          }`),
		"core/projection-default:2": []byte(`{
            "id": 2,
            "name": "topics",
            "display_name": "Topics",
            "projector_id": 1
          }`),
		"core/projection-default:3": []byte(`{
            "id": 3,
            "name": "agenda_list_of_speakers",
            "display_name": "List of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:4": []byte(`{
            "id": 4,
            "name": "agenda_current_list_of_speakers",
            "display_name": "Current list of speakers",
            "projector_id": 1
          }`),
		"core/projection-default:5": []byte(`{
            "id": 5,
            "name": "motions",
            "display_name": "Motions",
            "projector_id": 1
          }`),
		"core/projection-default:6": []byte(`{
            "id": 6,
            "name": "motionBlocks",
            "display_name": "Motion blocks",
            "projector_id": 1
          }`),
		"core/projection-default:7": []byte(`{
            "id": 7,
            "name": "assignments",
            "display_name": "Elections",
            "projector_id": 1
          }`),
		"core/projection-default:8": []byte(`{
            "id": 8,
            "name": "users",
            "display_name": "Participants",
            "projector_id": 1
          }`),
		"core/projection-default:9": []byte(`{
            "id": 9,
            "name": "mediafiles",
            "display_name": "Files",
            "projector_id": 1
          }`),
		"core/projector-message:1": []byte(`{
            "id": 1,
            "message": "<p>test</p>"
          }`),
		"core/projector:1": []byte(`{
            "id": 1,
            "elements": [
              {
                "name": "mediafiles/mediafile",
                "id": 3
              }
            ],
            "elements_preview": [],
            "elements_history": [
              [
                {
                  "name": "assignments/assignment",
                  "id": 1
                }
              ]
            ],
            "scale": 0,
            "scroll": 0,
            "name": "Default projector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [
              1,
              2,
              3,
              4,
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13
            ],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/projector:2": []byte(`{
            "id": 2,
            "elements": [
              {
                "name": "core/clock",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers",
                "stable": false
              },
              {
                "name": "agenda/current-speaker-chyron",
                "stable": true
              },
              {
                "name": "agenda/current-list-of-speakers-overlay",
                "stable": true
              }
            ],
            "elements_preview": [],
            "elements_history": [],
            "scale": 0,
            "scroll": 0,
            "name": "sideprojector",
            "width": 1200,
            "aspect_ratio_numerator": 16,
            "aspect_ratio_denominator": 9,
            "reference_projector_id": 1,
            "projectiondefaults_id": [],
            "color": "#000000",
            "background_color": "#ffffff",
            "header_background_color": "#317796",
            "header_font_color": "#f5f5f5",
            "header_h1_color": "#317796",
            "chyron_background_color": "#317796",
            "chyron_font_color": "#ffffff",
            "show_header_footer": true,
            "show_title": true,
            "show_logo": true
          }`),
		"core/tag:1": []byte(`{
            "id": 1,
            "name": "T1"
          }`),
		"core/tag:2": []byte(`{
            "id": 2,
            "name": "T2"
          }`),
		"mediafiles/mediafile:2": []byte(`{
            "id": 2,
            "title": "A.txt",
            "media_url_prefix": "/media/",
            "filesize": "< 1 kB",
            "mimetype": "text/plain",
            "pdf_information": {},
            "access_groups_id": [],
            "create_timestamp": "2020-08-11T11:16:07.013124+02:00",
            "is_directory": false,
            "path": "A.txt",
            "parent_id": null,
            "list_of_speakers_id": 5,
            "inherited_access_groups_id": true
          }`),
		"motions/category:1": []byte(`{
            "id": 1,
            "name": "first",
            "prefix": "A",
            "parent_id": null,
            "weight": 2,
            "level": 0
          }`),
		"motions/category:2": []byte(`{
            "id": 2,
            "name": "second",
            "prefix": "B",
            "parent_id": 1,
            "weight": 4,
            "level": 1
          }`),
		"motions/category:3": []byte(`{
            "id": 3,
            "name": "third",
            "prefix": "C",
            "parent_id": null,
            "weight": 6,
            "level": 0
          }`),
		"motions/motion-block:1": []byte(`{
            "id": 1,
            "title": "block",
            "agenda_item_id": 8,
            "list_of_speakers_id": 12,
            "internal": false,
            "motions_id": [
              1,
              2
            ]
          }`),
		"motions/motion-option:1": []byte(`{
            "id": 1,
            "yes": "0.000000",
            "no": "1.000000",
            "abstain": "0.000000",
            "poll_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-option:2": []byte(`{
            "id": 2,
            "yes": "1.000000",
            "no": "0.000000",
            "abstain": "0.000000",
            "poll_id": 2,
            "pollstate": 4
          }`),
		"motions/motion-poll:1": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "named",
            "title": "Abstimmung",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              1
            ],
            "id": 1,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-poll:2": []byte(`{
            "motion_id": 3,
            "pollmethod": "YNA",
            "state": 4,
            "type": "pseudoanonymous",
            "title": "Abstimmung (2)",
            "groups_id": [
              3
            ],
            "votesvalid": "1.000000",
            "votesinvalid": "0.000000",
            "votescast": "1.000000",
            "options_id": [
              2
            ],
            "id": 2,
            "onehundred_percent_base": "YNA",
            "majority_method": "simple",
            "voted_id": [
              5
            ],
            "user_has_voted": false
          }`),
		"motions/motion-vote:1": []byte(`{
            "id": 1,
            "weight": "1.000000",
            "value": "N",
            "user_id": 5,
            "option_id": 1,
            "pollstate": 4
          }`),
		"motions/motion-vote:2": []byte(`{
            "id": 2,
            "weight": "1.000000",
            "value": "Y",
            "user_id": null,
            "option_id": 2,
            "pollstate": 4
          }`),
		"motions/motion:2": []byte(`{
            "id": 2,
            "identifier": "Ä-1",
            "title": "Änderungsantrag zu Leadmotion1",
            "text": "",
            "amendment_paragraphs": [
              "<p>Lorem ipsum dolor sit amet, consectedfgddfgdf&nbsp; gdfgdfg dfg dfg ww ontes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi.<br>Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.<br>Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Fusce vulputate eleifend sapien. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nullam accumsan lorem in dui. Cras ultricies mi eu turpis hendrerit fringilla. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In ac dui quis mi consectetuer lacinia. Nam pretium turpis et arcu. Duis arcu tortor, suscipit eget, imperdiet nec, imperdiet iaculis, ipsum. Sed aliquam ultrices mauris. Integer ante arcu, accumsan a, consectetuer eget, posuere ut, mauris. Praesent adipiscing. Phasellus ullamcorper ipsum rutrum nunc. Nunc nonummy metus. Vestibulum volutpat pretium libero. Cras id dui. Aenean ut</p>"
            ],
            "modified_final_version": "",
            "reason": "",
            "parent_id": 1,
            "category_id": 2,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 1,
            "origin": "",
            "submitters": [
              {
                "id": 3,
                "user_id": 1,
                "motion_id": 2,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [],
            "attachments_id": [],
            "agenda_item_id": 6,
            "list_of_speakers_id": 9,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T11:39:35.025914+02:00",
            "last_modified": "2020-08-11T12:41:55.666495+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/motion:3": []byte(`{
            "id": 3,
            "identifier": "2",
            "title": "Public",
            "text": "<p>a</p>",
            "amendment_paragraphs": null,
            "modified_final_version": "",
            "reason": "<p>a</p>",
            "parent_id": null,
            "category_id": 1,
            "category_weight": 10000,
            "comments": [],
            "motion_block_id": 2,
            "origin": "",
            "submitters": [
              {
                "id": 4,
                "user_id": 1,
                "motion_id": 3,
                "weight": 1
              }
            ],
            "supporters_id": [],
            "state_id": 1,
            "state_extension": null,
            "state_restriction": [],
            "statute_paragraph_id": null,
            "workflow_id": 1,
            "recommendation_id": null,
            "recommendation_extension": null,
            "tags_id": [
              1
            ],
            "attachments_id": [],
            "agenda_item_id": 7,
            "list_of_speakers_id": 10,
            "sort_parent_id": null,
            "weight": 10000,
            "created": "2020-08-11T12:24:45.289233+02:00",
            "last_modified": "2020-08-11T12:41:51.319986+02:00",
            "change_recommendations_id": [],
            "amendments_id": []
          }`),
		"motions/state:1": []byte(`{
            "id": 1,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              2,
              3,
              4
            ],
            "workflow_id": 1
          }`),
		"motions/state:10": []byte(`{
            "id": 10,
            "name": "withdrawed",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:11": []byte(`{
            "id": 11,
            "name": "adjourned",
            "recommendation_label": "Adjournment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:12": []byte(`{
            "id": 12,
            "name": "not concerned",
            "recommendation_label": "No concernment",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:13": []byte(`{
            "id": 13,
            "name": "refered to committee",
            "recommendation_label": "Referral to committee",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:14": []byte(`{
            "id": 14,
            "name": "needs review",
            "recommendation_label": null,
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:15": []byte(`{
            "id": 15,
            "name": "rejected (not authorized)",
            "recommendation_label": "Rejection (not authorized)",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:2": []byte(`{
            "id": 2,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:3": []byte(`{
            "id": 3,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:4": []byte(`{
            "id": 4,
            "name": "not decided",
            "recommendation_label": "No decision",
            "css_class": "grey",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 1
          }`),
		"motions/state:5": []byte(`{
            "id": 5,
            "name": "in progress",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [
              "is_submitter"
            ],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": true,
            "dont_set_identifier": true,
            "show_state_extension_field": true,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              6,
              10
            ],
            "workflow_id": 2
          }`),
		"motions/state:6": []byte(`{
            "id": 6,
            "name": "submitted",
            "recommendation_label": null,
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": true,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": true,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": false,
            "next_states_id": [
              7,
              10,
              15
            ],
            "workflow_id": 2
          }`),
		"motions/state:7": []byte(`{
            "id": 7,
            "name": "permitted",
            "recommendation_label": "Permission",
            "css_class": "lightblue",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": true,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 0,
            "show_recommendation_extension_field": true,
            "next_states_id": [
              8,
              9,
              10,
              11,
              12,
              13,
              14
            ],
            "workflow_id": 2
          }`),
		"motions/state:8": []byte(`{
            "id": 8,
            "name": "accepted",
            "recommendation_label": "Acceptance",
            "css_class": "green",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": 1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/state:9": []byte(`{
            "id": 9,
            "name": "rejected",
            "recommendation_label": "Rejection",
            "css_class": "red",
            "restriction": [],
            "allow_support": false,
            "allow_create_poll": false,
            "allow_submitter_edit": false,
            "dont_set_identifier": false,
            "show_state_extension_field": false,
            "merge_amendment_into_final": -1,
            "show_recommendation_extension_field": false,
            "next_states_id": [],
            "workflow_id": 2
          }`),
		"motions/statute-paragraph:1": []byte(`{
            "id": 1,
            "title": "Statute",
            "text": "<p>test</p>",
            "weight": 10000
          }`),
		"motions/workflow:1": []byte(`{
            "id": 1,
            "name": "Simple Workflow",
            "states_id": [
              1,
              2,
              3,
              4
            ],
            "first_state_id": 1
          }`),
		"motions/workflow:2": []byte(`{
            "id": 2,
            "name": "Complex Workflow",
            "states_id": [
              5,
              6,
              7,
              8,
              9,
              10,
              11,
              12,
              13,
              14,
              15
            ],
            "first_state_id": 5
          }`),
		"topics/topic:1": []byte(`{
            "id": 1,
            "title": "Topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 1,
            "list_of_speakers_id": 1
          }`),
		"topics/topic:2": []byte(`{
            "id": 2,
            "title": "Hidden",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 2,
            "list_of_speakers_id": 2
          }`),
		"topics/topic:3": []byte(`{
            "id": 3,
            "title": "Internal",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 3,
            "list_of_speakers_id": 3
          }`),
		"topics/topic:4": []byte(`{
            "id": 4,
            "title": "Another public topic",
            "text": "",
            "attachments_id": [],
            "agenda_item_id": 9,
            "list_of_speakers_id": 14
          }`),
		"users/group:1": []byte(`{
            "id": 1,
            "name": "Default",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_see",
              "users.can_change_password"
            ]
          }`),
		"users/group:2": []byte(`{
            "id": 2,
            "name": "Admin",
            "permissions": []
          }`),
		"users/group:3": []byte(`{
            "id": 3,
            "name": "Delegates",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password"
            ]
          }`),
		"users/group:4": []byte(`{
            "id": 4,
            "name": "Staff",
            "permissions": [
              "agenda.can_manage",
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_manage_list_of_speakers",
              "agenda.can_see_list_of_speakers",
              "agenda.can_be_speaker",
              "assignments.can_manage",
              "assignments.can_nominate_other",
              "assignments.can_nominate_self",
              "assignments.can_see",
              "core.can_see_history",
              "core.can_manage_projector",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "core.can_manage_tags",
              "mediafiles.can_manage",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_manage",
              "motions.can_manage_metadata",
              "motions.can_see",
              "motions.can_see_internal",
              "users.can_change_password",
              "users.can_manage",
              "users.can_see_extra_data",
              "users.can_see_name"
            ]
          }`),
		"users/group:5": []byte(`{
            "id": 5,
            "name": "Committees",
            "permissions": [
              "agenda.can_see",
              "agenda.can_see_internal_items",
              "agenda.can_see_list_of_speakers",
              "assignments.can_see",
              "core.can_see_frontpage",
              "core.can_see_projector",
              "mediafiles.can_see",
              "motions.can_create",
              "motions.can_create_amendments",
              "motions.can_see",
              "motions.can_support",
              "users.can_change_password",
              "users.can_see_name"
            ]
          }`),
		"users/user:1": []byte(`{
            "first_name": "",
            "username": "admin",
            "about_me": "",
            "id": 1,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              2
            ],
            "number": "",
            "last_name": "Administrator",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:2": []byte(`{
            "first_name": "candidate1",
            "username": "candidate1",
            "about_me": "",
            "id": 2,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:3": []byte(`{
            "first_name": "candidate2",
            "username": "candidate2",
            "about_me": "",
            "id": 3,
            "structure_level": "",
            "is_present": false,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:4": []byte(`{
            "first_name": "a",
            "username": "a",
            "about_me": "",
            "id": 4,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:5": []byte(`{
            "first_name": "b",
            "username": "b",
            "about_me": "",
            "id": 5,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [
              3
            ],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:6": []byte(`{
            "first_name": "speaker1",
            "username": "speaker1",
            "about_me": "",
            "id": 6,
            "structure_level": "layer X",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "title",
            "groups_id": [],
            "number": "3",
            "last_name": "the last name",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:7": []byte(`{
            "first_name": "speaker2",
            "username": "speaker2",
            "about_me": "",
            "id": 7,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "title": "",
            "groups_id": [],
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
		"users/user:8": []byte(`{
            "first_name": "not required",
            "username": "not required",
            "about_me": "",
            "id": 8,
            "structure_level": "",
            "is_present": true,
            "vote_weight": "1.000000",
            "groups_id": [],
            "title": "",
            "email": "",
            "number": "",
            "last_name": "",
            "is_committee": false,
            "gender": ""
          }`),
	},
}
