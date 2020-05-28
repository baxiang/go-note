// +build wireinject

package controllers


var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)


func CreateReviewsController(cf string, sto repositories.ReviewsRepository) (*ReviewsController, error) {
	panic(wire.Build(testProviderSet))
}