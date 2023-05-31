package main

import (
	"fmt"
	xds "github.com/cncf/xds/go/xds/type/v3"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/api"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/http"
	"google.golang.org/protobuf/types/known/anypb"
)

const Name = "simple"

func init() {
		http.RegisterHttpFilterConfigParser(&parser{})
	http.RegisterHttpFilterConfigFactory(Name, ConfigFactory)
}

type config struct {
	echoBody string
}

type parser struct {
}

// func (p *parser) Parse(any *anypb.Any) (interface{}, error) {
func (p *parser) Parse(any *anypb.Any) interface{} {
	configStruct := &xds.TypedStruct{}
	if err := any.UnmarshalTo(configStruct); err != nil {
		return nil
		// return nil, err
	}

	v := configStruct.Value
	conf := &config{}
	prefix, ok := v.AsMap()["prefix_localreply_body"]
	if !ok {
		fmt.Println("Missing prefix_localreply_body")
		// return nil, errors.New("missing prefix_localreply_body")
		return nil
	}
	if str, ok := prefix.(string); ok {
		conf.echoBody = str
	} else {
		fmt.Println("prefix_localreply_body: expect string while got %T", prefix)
		// return nil, fmt.Errorf("prefix_localreply_body: expect string while got %T", prefix)
		return nil
	}
	fmt.Println("------------ Returning the config")
	fmt.Println(conf)
	return conf
}

func (p *parser) Merge(parent interface{}, child interface{}) interface{} {
	parentConfig := parent.(*config)
	childConfig := child.(*config)

	// copy one, do not update parentConfig directly.
	newConfig := *parentConfig
	if childConfig.echoBody != "" {
		newConfig.echoBody = childConfig.echoBody
	}
	return &newConfig
}

func ConfigFactory(c interface{}) api.StreamFilterFactory {
	conf, ok := c.(*config)
	if !ok {
		fmt.Println(conf)
		panic("unexpected config type")
	}

	return func(callbacks api.FilterCallbackHandler) api.StreamFilter {
		return &filter{
			callbacks: callbacks,
			config:    conf,
		}
	}
}

func main() {}
