package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/skullkon/info-app/internal/domain"
	"github.com/skullkon/info-app/pkg/client"
	"github.com/skullkon/info-app/pkg/logging"
	"strings"
)

type Repository struct {
	client client.Client
	logger *logging.Logger
}

type Result struct {
	Os string `ch:"os"`
}

func NewRepository(client client.Client, logger *logging.Logger) *Repository {
	return &Repository{
		client: client,
		logger: logger,
	}
}

func (r *Repository) Insert(ctx context.Context, info []domain.Info) error {
	batch, err := r.client.PrepareBatch(ctx, "INSERT INTO info")
	if err != nil {
		return err
	}

	for i := 0; i < len(info); i++ {
		err := batch.Append(
			info[i].Id,
			info[i].Ip,
			info[i].TypeOfDevice,
			info[i].Os,
			info[i].OsVersion,
			info[i].Browser,
			info[i].BrowserVersion,
			info[i].Brand,
			info[i].Model,
			info[i].Resolution,
			info[i].Time,
		)
		if err != nil {
			fmt.Print(err)
			return err
		}
	}

	err = batch.Send()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]domain.Info, error) {
	var test []domain.Info
	err := r.client.Select(ctx, &test, "select * from info")
	if err != nil {
		return nil, err
	}
	return test, nil
}

func (r *Repository) GetRating(ctx context.Context, attr string) ([]string, error) {
	var res Result
	var answer []string
	//trying to protect against sql injection
	attribute := strings.Split(attr, " ")
	query := fmt.Sprintf("SELECT %s FROM info Group by %s ORDER BY count(id) DESC LIMIT 100", attribute[0], attribute[0])

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&res.Os)
		if err != nil {
			return nil, err
		}
		answer = append(answer, res.Os)
	}

	return answer, nil
}

func (r *Repository) GetRatingWithParam(ctx context.Context, column string, value string, attr string) ([]string, error) {
	var res Result
	var answer []string
	//trying to protect against sql injection except for binding params
	attribute := strings.Split(attr, " ")
	clm := strings.Split(column, " ")
	val := strings.Split(value, " ")

	if len(clm) != 1 && len(val) != 1 {
		return nil, errors.New("something goes wrong")
	}

	query := fmt.Sprintf("SELECT %s FROM info WHERE %s = '%s' Group by %s ORDER BY count(id) DESC LIMIT 100", attribute[0], clm[0], val[0], attribute[0])

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&res.Os)
		if err != nil {
			return nil, err
		}
		answer = append(answer, res.Os)
	}

	return answer, nil
}
