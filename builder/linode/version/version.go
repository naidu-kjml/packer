package version

import (
	"github.com/hashicorp/packer/helper/version"
	packerVersion "github.com/hashicorp/packer/version"
)

var LinodePluginVersion = version.PluginVersion{
	Version:           packerVersion.Version,
	VersionPrerelease: packerVersion.VersionPrerelease,
}