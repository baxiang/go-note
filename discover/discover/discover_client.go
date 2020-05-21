package discover

import "log"

type DiscoveryClient interface {

	/*
	 * @Date 2020/5/18 9:45 下午
	 * @Description method
	 * @Param
	 * @Return
	 */
	Register(serviceName string,
			instanceId string,
			instanceHost string,
			instancePort string,
			healthCheckUrl string,
		    meta map[string]string,
		    logger *log.Logger,
		)bool
	/*
	 * @Date 2020/5/18 9:45 下午
	 * @Description 服务注销接口
	 * @Param
	 * @Return
	 */
	DeRegister(instanceId string,
		logger *log.Logger,
		)bool

	DiscoverServices(serverName string,
		logger *log.Logger,
		)[]interface{}

}