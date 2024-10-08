package imagesource

import (
	"errors"
	"net/http"

	"github.com/distribution/distribution/v3"
	registryclient "github.com/distribution/distribution/v3/registry/client"
	"github.com/distribution/reference"
)

func NewDryRun(ref TypedImageReference) (distribution.Repository, error) {
	named, err := reference.WithName(ref.Ref.RepositoryName())
	if err != nil {
		return nil, err
	}
	return registryclient.NewRepository(named, ref.Ref.RegistryURL().String(), dryRunRoundTripper)
}

var dryRunRoundTripper = errorRoundTripper{errors.New("dry-run repository is not available")}

type errorRoundTripper struct {
	err error
}

func (rt errorRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, rt.err
}
