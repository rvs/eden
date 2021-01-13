package defaults

//ImageOptions indicates parameters of predefined VMs
type ImageOptions struct {
	Size   int64
	Sha256 string
}

//ImageStore contains image options for images used inside tests
var ImageStore = map[string]*ImageOptions {
	"https://cloud-images.ubuntu.com/releases/groovy/release-20201022.1/ubuntu-20.10-server-cloudimg-amd64.img": {
		Size: 558760448,
		Sha256: "ef3ed6aaf9c8fe1d063d556ace6c4dfbb51920d12ba8312e09a1baf3b3eedf3d",
	},
	"https://cloud-images.ubuntu.com/releases/groovy/release-20201022.1/ubuntu-20.10-server-cloudimg-arm64.img": {
		Size: 525336576,
		Sha256: "c64a5e20dd61cc112de2a47d8b0a3ec30a553fe5fe54ca0a5f83c840778aa300",
	},
}