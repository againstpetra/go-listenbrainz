// Copyright (C) 2019 Luiz de Milon (kori)

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package listenbrainz

import (
	"reflect"
	"testing"
	"time"
)

func TestGetSubmissionTime(t *testing.T) {
	var tests = []struct {
		Length, Result time.Duration
	}{
		{time.Duration(-1), time.Duration(0)},
		{4 * time.Minute, 2 * time.Minute},
		{5 * time.Minute, 2*time.Minute + 30*time.Second},
	}

	for _, test := range tests {
		st, err := GetSubmissionTime(test.Length)
		if err != nil {
			t.Log("Test failed successfully at:", test.Length, ":", err)
		}
		if st != test.Result {
			t.Error("Expected", test.Result, "got", st)
		}
	}
}

func TestFormatPlayingNow(t *testing.T) {
	additional_info := AdditionalInfo{
		MediaPlayer:             "Rythembox",
		SubmissionClient:        "Rhythmbox ListenBrainz Plugin",
		SubmissionClientVersion: "1.0",
		ReleaseMbid:             "bf9e91ea-8029-4a04-a26a-224e00a83266",
		ArtistMbid:              []string{"db92a151-1ac2-438b-bc43-b82e149ddd50"},
		RecordingMbid:           "98255a8c-017a-4bc7-8dd6-1fa36124572b",
		Tags:                    []string{"you", "just", "got", "rick", "rolled!"},
		DurationMs:              222000,
	}
	track := Track{
		Title:          "b",
		Artist:         "a",
		Album:          "c",
		AdditionalInfo: additional_info,
	}

	ts := Submission{
		ListenType: "playing_now",
		Payloads: Payloads{
			Payload{
				Track: track,
			},
		},
	}

	s := FormatPlayingNow(track)

	if !reflect.DeepEqual(ts, s) {
		t.Error("Expected", ts, "got", s)
	}
}

func TestFormatSingle(t *testing.T) {
	track := Track{
		Title:  "b",
		Artist: "a",
		Album:  "c",
	}

	time := time.Now().Unix()

	ts := Submission{
		ListenType: "single",
		Payloads: Payloads{
			Payload{
				ListenedAt: time,
				Track:      track,
			},
		},
	}

	s := FormatSingle(track, time)

	if !reflect.DeepEqual(ts, s) {
		t.Error("Expected", ts, "got", s)
	}
}
