package converter

import (
	"testing"

	pb "github.com/s1ntezc0der/event-horizon-backend/services/game/proto"
)

func TestSubmitScoreFromProto(t *testing.T) {
	m := SubmitScoreFromProto(&pb.SubmitScoreRequest{UserId: "u1", GameId: "hexagon", Level: 2, Score: 100})
	if m.UserID != "u1" || m.GameID != "hexagon" || m.Score != 100 {
		t.Fatalf("%+v", m)
	}
}
