package migrations

import (
	"code.gitea.io/sdk/gitea"
	"github.com/google/go-github/github"
)

// Repository migrates a GitHub Repository to a Gitea Repository without migrating issues etc.
func (m *Migratory) Repository(gr *github.Repository) (*gitea.Repository, error) {
	var err error
	m.repository, err = m.Client.MigrateRepo(gitea.MigrateRepoOption{
		Description:  gr.GetDescription(),
		AuthPassword: m.AuthPassword,
		AuthUsername: m.AuthUsername,
		CloneAddr:    gr.GetCloneURL(),
		RepoName:     gr.GetName(),
		UID:          int(m.NewOwnerID),
		Private:      m.Private,
	})
	return m.repository, err
}
