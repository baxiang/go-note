package service

import "errors"

// Service Define a service interface
type CalService interface {

	// Add calculate a+b
	Add(a, b int) int

	// Subtract calculate a-b
	Subtract(a, b int) int

	// Multiply calculate a*b
	Multiply(a, b int) int

	// Divide calculate a/b
	Divide(a, b int) (int, error)

	// HealthCheck check service health status
	HealthCheck() bool
}

//ArithmeticService implement Service interface
type CalServiceImpl struct {
}

// Add implement Add method
func (c CalServiceImpl) Add(a, b int) int {
	return a + b
}

// Subtract implement Subtract method
func (c CalServiceImpl) Subtract(a, b int) int {
	return a - b
}

// Multiply implement Multiply method
func (c CalServiceImpl) Multiply(a, b int) int {
	return a * b
}

// Divide implement Divide method
func (c CalServiceImpl) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("the dividend can not be zero!")
	}

	return a / b, nil
}

// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true。
func (c CalServiceImpl) HealthCheck() bool {
	return true
}

// ServiceMiddleware define service middleware
type ServiceMiddleware func(CalService) CalService
