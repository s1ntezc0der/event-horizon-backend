package hexagons

import (
	"fmt"
)

type Move struct {
	FromX     int
	FromY     int
	ToX       int
	ToY       int
	Timestamp int32
}

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateMoves(seed string, moves []Move, finalScore int) (bool, int, error) {
	logger.Info("Validation started", "seed", seed, "moves_count", len(moves))

	// Pancaker FE (tray → empty hex) does not replay the old MoveTile board.
	// Same MVP policy as memory: empty moves → trust client score (capped).
	if len(moves) == 0 {
		if finalScore < 0 {
			finalScore = 0
		}
		if finalScore > 100_000 {
			finalScore = 100_000
		}
		logger.Info("Hexagon empty moves — trusting client score", "final_score", finalScore)
		return true, finalScore, nil
	}

	board, err := GenerateInitialBoard(seed, 3)
	if err != nil {
		logger.Error("Failed to generate board", "error", err)
		return false, 0, err
	}

	for i, move := range moves {
		from := Coord{Q: move.FromX, R: move.FromY}
		to := Coord{Q: move.ToX, R: move.ToY}
		logger.Debug("Applying move", "index", i, "from", from, "to", to)

		if err := board.MoveTile(from, to); err != nil {
			logger.Error("Move failed", "index", i, "error", err)
			return false, 0, fmt.Errorf("move %d invalid: %w", i, err)
		}
	}

	calculatedScore := board.CalculateScore()
	logger.Info("Validation completed", "calculated_score", calculatedScore, "final_score", finalScore)
	logger.Info("Final board state",
		"total_score", board.TotalScore,
		"tiles_count", len(board.Tiles),
	)

	return true, calculatedScore, nil
}
