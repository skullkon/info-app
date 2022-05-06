package information

import "context"

type Repository interface {
	Insert(ctx context.Context, info []Info) error
	GetAll(ctx context.Context) ([]Info, error)
	GetRating(ctx context.Context) ([]string, error)
	GetRatingWithParam(ctx context.Context, param string) ([]string, error)
}
