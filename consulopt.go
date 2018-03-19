package consulib

import (
	consulapi "github.com/hashicorp/consul/api"
	"fmt"
)


func ServiceRegster(serviceHttp,sId,sName,host,router string,port int ,tags []string) error{
	config := consulapi.DefaultConfig()
	if serviceHttp != "" {
		config.Address = serviceHttp
	}
	client,err := consulapi.NewClient(config)
	if err != nil {
		return  err
	}
	registration := new(consulapi.AgentServiceRegistration)

	registration.ID =  sId
	registration.Name = sName
	registration.Port = port
	registration.Tags = tags
	registration.Address = host

	check := &consulapi.AgentServiceCheck{
		HTTP:  fmt.Sprintf("http://%s:%d%s",registration.Address,port,router),
		Timeout:                        "20s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "60s", //check失败后30秒删除本服务
	}


	registration.Check = check


	err = client.Agent().ServiceRegister(registration)

	if err !=nil {
		return  err
	}

	return  nil


}


func FindServiceByServiceName(serviceName,serviceHttp string)string{
	if serviceName == "" || len(serviceName) == 0 {
		return ""
	}
	config := consulapi.DefaultConfig()
	if serviceHttp != "" {
		config.Address = serviceHttp
	}
	client,err := consulapi.NewClient(config)
	if err != nil {
		return ""
	}
	services,_ := client.Agent().Services()

	if _,found := services[serviceName]; found {
	   return  serviceHttp
	}
	return  ""
}

