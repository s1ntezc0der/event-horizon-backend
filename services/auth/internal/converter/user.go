package converter

import (
	pb "github.com/s1ntezc0der/event-horizon-backend/services/auth/proto"
	"github.com/s1ntezc0der/event-horizon-backend/services/auth/internal/model"
)

// UserToProto maps a domain user to GetUserResponse fields (without scores).
func UserToProto(u *model.User) *pb.GetUserResponse {
	if u == nil {
		return nil
	}
	return &pb.GetUserResponse{
		UserId:   u.ID,
		Email:    u.Email,
		Nickname: u.Nickname,
		Role:     u.Role,
	}
}

// UserToProtoWithScores attaches best scores to the response.
func UserToProtoWithScores(u *model.User, bestScores map[string]int32, totalScore int32) *pb.GetUserResponse {
	resp := UserToProto(u)
	if resp == nil {
		return nil
	}
	resp.BestScores = bestScores
	resp.TotalScore = totalScore
	return resp
}
