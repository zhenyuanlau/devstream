package reposcaffolding

import (
	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
	"github.com/devstream-io/devstream/pkg/util/log"
	"github.com/devstream-io/devstream/pkg/util/scm"
)

func GetDynamicStatus(options configmanager.RawOptions) (statemanager.ResourceStatus, error) {
	opts, err := NewOptions(options)
	if err != nil {
		return nil, err
	}

	scmClient, err := scm.NewClientWithAuth(opts.DestinationRepo)
	if err != nil {
		log.Debugf("reposcaffolding status init repo failed: %+v", err)
		return nil, err
	}

	repoInfo, err := scmClient.DescribeRepo()
	if err != nil {
		log.Debugf("reposcaffolding status describe repo failed: %+v", err)
		return nil, err
	}

	resStatus := statemanager.ResourceStatus{
		"repo":     repoInfo.Repo,
		"owner":    repoInfo.Owner,
		"org":      repoInfo.Org,
		"repoURL":  repoInfo.CloneURL,
		"repoType": repoInfo.RepoType,
		"source":   opts.SourceRepo.BuildScmURL(),
	}

	resStatus.SetOutputs(statemanager.ResourceOutputs{
		"repo":    repoInfo.Repo,
		"org":     repoInfo.Org,
		"owner":   repoInfo.Owner,
		"repoURL": repoInfo.CloneURL,
	})
	return resStatus, nil
}
