package main

import (
	"log"
	"samples/eureka"
	"samples/feign"
)

func main() {

	// 1.1 configure maps, app => app's urls
	//feign.DefaultFeign.UseUrls(map[string][]string{
	//	"MS-REGISTRY-URLS": {"http://192.168.20.236:9001", "http://192.168.20.237:9001"},
	//})

	// 1.2 use discovery client
	config := eureka.GetDefaultEurekaClientConfig()
	config.Region = "region-cn-hd-1"
	config.AvailabilityZones = map[string]string{
		"region-cn-hd-1": "zone-cn-hz-1",
	}

	config.ServiceUrl = map[string]string{
		"zone-cn-hz-1": "http://127.0.0.1:8751/instance",
	}

	eureka.DefaultClient.Config(config).
		Register("APP_ID_CLIENT_FROM_DNS", 9000).
		Run()

	// 2.1 request by registry apps
	log.Println("----> request by registy apps")
	// assign a eureka client explicitly, or use the eureka.DefaultClient by default.
	feign.DefaultFeign.UseDiscoveryClient(eureka.DefaultClient)
	res, err := feign.DefaultFeign.App("puzzle").R().SetHeaders(map[string]string{
		"Content-Type": "application/json",
	}).Get("/articles/page")
	if err != nil {
		log.Println("err=", err.Error())
		return
	}

	log.Println("res=", string(res.Body()))
	//
	//// 2.2 request by configuration apps
	//log.Println("----> request by configuration apps")
	//res, err = feign.DefaultFeign.App("MS-REGISTRY-URLS").R().SetHeaders(map[string]string{
	//	"Content-Type": "application/json",
	//}).Get("/eureka/apps/WM")
	//if err != nil {
	//	log.Println("err=", err.Error())
	//	return
	//}
	//
	//log.Println("res=", string(res.Body()))

	select {}
}
