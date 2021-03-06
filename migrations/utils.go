package migrations

import (
	"strings"
	"testing"
	"time"

	"code.gitea.io/sdk/gitea"
	"github.com/stretchr/testify/assert"
)

// DemoMigratory is been used for testing
var DemoMigratory = &Migratory{
	Options: Options{
		AuthUsername: "demo",
		AuthPassword: "demo",
		NewOwnerID:   1,
	},
	Client: gitea.NewClient("http://gitea:3000", "8bffa364d5a4b2f18421426da0baf6ccddd16d6b"),
	repository: &gitea.Repository{
		Name: "demo",
		Owner: &gitea.User{
			UserName: "demo",
		},
	},
	migratedMilestones: make(map[int64]int64),
	migratedLabels:     make(map[int64]int64),
}

var demoTime = time.Date(2018, 01, 01, 01, 01, 01, 01, time.UTC)

func assertNoError(t *testing.T, err error) {
	if err != nil && strings.Contains(err.Error(), "lookup gitea") {
		t.Skip("gitea instance is not running")
	} else {
		assert.NoError(t, err)
	}
}
