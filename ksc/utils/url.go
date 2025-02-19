package utils

import (
	"fmt"
	"os"
)

type UrlInfo struct {
	UseSSL      bool
	Locate      bool
	UseInternal bool
}

type ServiceInfo struct {
	Service string
	Region  string
}

const (
	// EnvApiDomain is the environment variable name for the API domain. If set, it will override the default domain.
	EnvApiDomain   = "KSC_GALAXY_API_DOMAIN"
	defaultDomain  = "example.com"
	urlTpl         = "%s://%s.api.%s"
	internalPrefix = "internal"
	http           = "http"
	https          = "https"
)

func Url(urlInfo *UrlInfo, info ServiceInfo) string {
	protocol := Protocol(urlInfo.UseSSL)
	var reqDomain string
	if val, ok := os.LookupEnv(EnvApiDomain); ok && val != "" {
		reqDomain = val
	} else {
		reqDomain = defaultDomain
	}
	var reqUrlPrefix string
	if urlInfo.Locate && &info.Region != nil {
		reqUrlPrefix = info.Service + "." + info.Region
	} else if info.Service != "" {
		reqUrlPrefix = info.Service
	}
	if reqUrlPrefix == "" || urlInfo.UseInternal {
		reqUrlPrefix = internalPrefix
	}
	return fmt.Sprintf(urlTpl, protocol, reqUrlPrefix, reqDomain)
}

func Protocol(useSSL bool) string {
	if useSSL {
		return https
	}
	return http
}
