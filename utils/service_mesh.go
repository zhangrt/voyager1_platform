package utils

import "github.com/zhangrt/voyager1_platform/global"

//  get request path prefix by service name
func GetPrefixByServiceName(namespaceName string, serviceName string) string {
	url := ""
	namespaces := global.GS_CONFIG.ServiceMesh.Namespace
	for index := range namespaces {
		if namespaces[index].Name == namespaceName {

			services := namespaces[index].Services
			for index := range services {

				if services[index].Name == serviceName {
					url = "http://" + services[index].Host + ":" + services[index].Port + services[index].Prefix + "/"
					break
				}

			}

			break
		}
	}

	return url
}
