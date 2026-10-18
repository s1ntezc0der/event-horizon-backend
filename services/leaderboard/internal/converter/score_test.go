package converter

import (
	"testing"

	"github.com/s1ntezc0der/event-horizon-backend/services/leaderboard/internal/model"
)

func TestScoreEntryToProto(t *testing.T) {
	p := ScoreEntryToProto(model.ScoreEntry{Rank: 1, UserID: "u", Nickname: "n", Score: 9})
	if p.Rank != 1 || p.UserId != "u" {
		t.Fatalf("%+v", p)
	}
}
