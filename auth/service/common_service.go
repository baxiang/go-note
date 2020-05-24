package service


type Service interface {
	SimpleData(username string) string
	AdminData(username string)string
	HealthCheck()bool
}

type commonService struct {

}

func NewCommonService()*commonService{
	return &commonService{}
}

func (s *commonService) SimpleData(username string) string {
	return "hello " + username + " ,simple data, with simple authority"
}

func (s *commonService) AdminData(username string) string {
	return "hello " + username + " ,admin data, with admin authority"

}

// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true
func (s *commonService) HealthCheck() bool {
	return true
}