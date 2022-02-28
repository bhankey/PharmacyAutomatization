package tokenrepo

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (r *TokenRepo) InsertRefreshToken(ctx context.Context, token entities.RefreshToken) error {
	const query = `
		INSERT INTO refresh_tokens(user_id, refresh_token, user_agent, ip, finger_print)
							VALUE ($1, $2, $3, $4, $5)
`

	if _, err := r.master.ExecContext(ctx, query, token.UserID, token.Token, token.UserAgent, token.IP, token.FingerPrint); err != nil {
		return fmt.Errorf("failed to insert refresh token: %w", err)
	}

	return nil
}
