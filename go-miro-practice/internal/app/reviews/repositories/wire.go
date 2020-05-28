// +build wireinject

package repositories

import (
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/database"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var resProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateReviewRepository(f string) (ReviewsRepository, error) {
	panic(wire.Build(resProviderSet))
}