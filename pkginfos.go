//
// Get information from packages installed on the system
//

package pkginfos

import (

	b "bytes"
	s "strings"
	"os/exec"
)

type DebPackage struct {

	Name				string `json:"name"`
	Status			string `json:"status"`
	Version			string `json:"version"`
	Size				string `json:"size"`
	Arch				string `json:"arch"`
	Maintainer	string `json:"maintainer"`
}

// use dpkg-query to list packages
func GetPackages() []DebPackage {

	var (
		cmdOut []byte
		err error
		pkgs []DebPackage
		p DebPackage
	)

	cmd := "dpkg-query"
	args := []string{"-W","-f", "${Package};${Status};${Version};${Installed-Size};${Architecture};${Maintainer}|"}

	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		return pkgs
	}

	// parse return datas from dpkg-query
	for _, dpkg := range b.Split(cmdOut, []byte{'|'}) {

		if(len(dpkg)<1) { continue } // Do I really need this ?!

		p = DebPackage{}
		pkg := s.Split(string(dpkg),";")

		p.Name = pkg[0]
		p.Status = pkg[1]
		p.Version = pkg[2]
		p.Size = pkg[3]
		p.Arch = pkg[4]
		p.Maintainer = pkg[5]

		pkgs = append(pkgs,p)
	}

	return pkgs
}
