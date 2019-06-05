package cloudflare

import (
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// sharedClient returns a common Cloudflare client setup needed for the
// sweeper functions.
func sharedClient() (*cloudflare.API, error) {
	client, err := cloudflare.New(os.Getenv("CLOUDFLARE_TOKEN"), os.Getenv("CLOUDFLARE_EMAIL"))

	if err != nil {
		return client, err
	}

	return client, nil
}
