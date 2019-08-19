package ignition

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/coreos/ignition/v2/config/v3_0/types"
)

func TestIngnitionFilesystem(t *testing.T) {
	testIgnition(t, `
		data "ignition_filesystem" "qux" {
			device = "/qux"
			format = "ext4"
		}

		data "ignition_filesystem" "baz" {
			device = "/baz"
			format = "ext4"
			wipe_filesystem = true
			label = "root"
			uuid = "qux"
			options = ["rw"]
		}

		data "ignition_config" "test" {
			filesystems = [
				"${data.ignition_filesystem.qux.id}",
				"${data.ignition_filesystem.baz.id}",
			]
		}
	`, func(c *types.Config) error {
		if len(c.Storage.Filesystems) != 2 {
			return fmt.Errorf("disks, found %d", len(c.Storage.Filesystems))
		}

		f := c.Storage.Filesystems[0]
		if f.Device != "/qux" {
			return fmt.Errorf("device, found %q", f.Device)
		}

		if string(*f.Format) != "ext4" {
			return fmt.Errorf("format, found %q", *f.Format)
		}

		f = c.Storage.Filesystems[1]

		if f.Device != "/baz" {
			return fmt.Errorf("device, found %q", f.Device)
		}

		if *f.Format != "ext4" {
			return fmt.Errorf("format, found %q", *f.Format)
		}

		if *f.Label != "root" {
			return fmt.Errorf("label, found %q", *f.Label)
		}

		if *f.UUID != "qux" {
			return fmt.Errorf("uuid, found %q", *f.UUID)
		}

		if *f.WipeFilesystem != true {
			return fmt.Errorf("wipe_filesystem, found %t", *f.WipeFilesystem)
		}

		if len(f.Options) != 1 || f.Options[0] != "rw" {
			return fmt.Errorf("options, found %q", f.Options)
		}

		return nil
	})
}

func TestIngnitionFilesystemInvalidPath(t *testing.T) {
	testIgnitionError(t, `
		data "ignition_filesystem" "foo" {
			device = "/foo"
			format = "ext4"
			path = "foo"
		}

		data "ignition_config" "test" {
			filesystems = [
				"${data.ignition_filesystem.foo.id}",
			]
		}
	`, regexp.MustCompile("absolute"))
}
