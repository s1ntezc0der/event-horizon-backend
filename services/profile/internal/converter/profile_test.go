package converter

import (
	"testing"

	"github.com/s1ntezc0der/event-horizon-backend/services/profile/internal/model"
)

func TestProfileToProto(t *testing.T) {
	p := ProfileToProto(&model.Profile{UserID: "u1", Email: "e", Nickname: "n", Lamps: 2})
	if p.UserId != "u1" || p.Lamps != 2 {
		t.Fatalf("%+v", p)
	}
}
