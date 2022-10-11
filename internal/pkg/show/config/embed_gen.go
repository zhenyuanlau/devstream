// Code generated by gen_embed_var.go; DO NOT EDIT.
package config

import _ "embed"

//go:embed default.yaml
var DefaultConfig string

// plugin default config
var (

	//go:embed plugins/argocd.yaml
	ArgocdDefaultConfig string

	//go:embed plugins/argocdapp.yaml
	ArgocdappDefaultConfig string

	//go:embed plugins/artifactory.yaml
	ArtifactoryDefaultConfig string

	//go:embed plugins/ci-generic.yaml
	CiGenericDefaultConfig string

	//go:embed plugins/devlake-config.yaml
	DevlakeConfigDefaultConfig string

	//go:embed plugins/devlake.yaml
	DevlakeDefaultConfig string

	//go:embed plugins/githubactions-golang.yaml
	GithubactionsGolangDefaultConfig string

	//go:embed plugins/githubactions-nodejs.yaml
	GithubactionsNodejsDefaultConfig string

	//go:embed plugins/githubactions-python.yaml
	GithubactionsPythonDefaultConfig string

	//go:embed plugins/gitlab-ce-docker.yaml
	GitlabCeDockerDefaultConfig string

	//go:embed plugins/gitlabci-generic.yaml
	GitlabciGenericDefaultConfig string

	//go:embed plugins/gitlabci-golang.yaml
	GitlabciGolangDefaultConfig string

	//go:embed plugins/gitlabci-java.yaml
	GitlabciJavaDefaultConfig string

	//go:embed plugins/harbor-docker.yaml
	HarborDockerDefaultConfig string

	//go:embed plugins/harbor.yaml
	HarborDefaultConfig string

	//go:embed plugins/hashicorp-vault.yaml
	HashicorpVaultDefaultConfig string

	//go:embed plugins/helm-generic.yaml
	HelmGenericDefaultConfig string

	//go:embed plugins/jenkins-pipeline.yaml
	JenkinsPipelineDefaultConfig string

	//go:embed plugins/jenkins.yaml
	JenkinsDefaultConfig string

	//go:embed plugins/jira-github-integ.yaml
	JiraGithubIntegDefaultConfig string

	//go:embed plugins/kube-prometheus.yaml
	KubePrometheusDefaultConfig string

	//go:embed plugins/openldap.yaml
	OpenldapDefaultConfig string

	//go:embed plugins/repo-scaffolding.yaml
	RepoScaffoldingDefaultConfig string

	//go:embed plugins/sonarqube.yaml
	SonarqubeDefaultConfig string

	//go:embed plugins/tekton.yaml
	TektonDefaultConfig string

	//go:embed plugins/trello-github-integ.yaml
	TrelloGithubIntegDefaultConfig string

	//go:embed plugins/trello.yaml
	TrelloDefaultConfig string

	//go:embed plugins/zentao.yaml
	ZentaoDefaultConfig string
)

var pluginDefaultConfigs = map[string]string{
	"argocd":               ArgocdDefaultConfig,
	"argocdapp":            ArgocdappDefaultConfig,
	"artifactory":          ArtifactoryDefaultConfig,
	"ci-generic":           CiGenericDefaultConfig,
	"devlake-config":       DevlakeConfigDefaultConfig,
	"devlake":              DevlakeDefaultConfig,
	"githubactions-golang": GithubactionsGolangDefaultConfig,
	"githubactions-nodejs": GithubactionsNodejsDefaultConfig,
	"githubactions-python": GithubactionsPythonDefaultConfig,
	"gitlab-ce-docker":     GitlabCeDockerDefaultConfig,
	"gitlabci-generic":     GitlabciGenericDefaultConfig,
	"gitlabci-golang":      GitlabciGolangDefaultConfig,
	"gitlabci-java":        GitlabciJavaDefaultConfig,
	"harbor-docker":        HarborDockerDefaultConfig,
	"harbor":               HarborDefaultConfig,
	"hashicorp-vault":      HashicorpVaultDefaultConfig,
	"helm-generic":         HelmGenericDefaultConfig,
	"jenkins-pipeline":     JenkinsPipelineDefaultConfig,
	"jenkins":              JenkinsDefaultConfig,
	"jira-github-integ":    JiraGithubIntegDefaultConfig,
	"kube-prometheus":      KubePrometheusDefaultConfig,
	"openldap":             OpenldapDefaultConfig,
	"repo-scaffolding":     RepoScaffoldingDefaultConfig,
	"sonarqube":            SonarqubeDefaultConfig,
	"tekton":               TektonDefaultConfig,
	"trello-github-integ":  TrelloGithubIntegDefaultConfig,
	"trello":               TrelloDefaultConfig,
	"zentao":               ZentaoDefaultConfig,
}

//go:embed templates/quickstart.yaml
var QuickStart string

//go:embed templates/gitops.yaml
var GitOps string
