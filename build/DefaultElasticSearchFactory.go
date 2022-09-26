package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	elog "github.com/pip-services3-gox/pip-services3-elasticsearch-gox/log"
)

/*
DefaultElasticSearchFactory are creates ElasticSearch components by their descriptors.
See ElasticSearchLogger
*/
type DefaultElasticSearchFactory struct {
	*cbuild.Factory
}

// NewDefaultElasticSearchFactory create a new instance of the factory.
// Retruns *DefaultElasticSearchFactory
// pointer on new factory
func NewDefaultElasticSearchFactory() *DefaultElasticSearchFactory {
	c := DefaultElasticSearchFactory{}
	c.Factory = cbuild.NewFactory()

	elasticSearchLoggerDescriptor := cref.NewDescriptor("pip-services", "logger", "elasticsearch", "*", "1.0")

	c.RegisterType(elasticSearchLoggerDescriptor, elog.NewElasticSearchLogger)

	return &c
}
