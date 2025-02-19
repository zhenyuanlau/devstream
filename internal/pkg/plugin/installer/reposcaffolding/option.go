package reposcaffolding

import (
	"github.com/mitchellh/mapstructure"

	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/pkg/util/scm/git"
)

type Options struct {
	SourceRepo      *git.RepoInfo `validate:"required" mapstructure:"sourceRepo"`
	DestinationRepo *git.RepoInfo `validate:"required" mapstructure:"destinationRepo"`
	Vars            map[string]interface{}
}

func NewOptions(options configmanager.RawOptions) (*Options, error) {
	var opts Options
	if err := mapstructure.Decode(options, &opts); err != nil {
		return nil, err
	}
	return &opts, nil
}

func (opts *Options) renderTplConfig() map[string]interface{} {
	renderConfig := opts.DestinationRepo.BuildRepoRenderConfig()
	for k, v := range opts.Vars {
		renderConfig[k] = v
	}
	return renderConfig
}
