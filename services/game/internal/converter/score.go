package converter

import (
	"github.com/s1ntezc0der/event-horizon-backend/services/game/internal/model"
	pb "github.com/s1ntezc0der/event-horizon-backend/services/game/proto"
)

func SubmitScoreFromProto(req *pb.SubmitScoreRequest) model.ScoreSubmission {
	if req == nil {
		return model.ScoreSubmission{}
	}
	return model.ScoreSubmission{
		UserID:   req.UserId,
		GameID:   req.GameId,
		Level:    req.Level,
		Score:    req.Score,
		Nickname: req.Nickname,
	}
}
