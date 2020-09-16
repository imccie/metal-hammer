module github.com/metal-stack/metal-hammer

go 1.15

require (
	github.com/beevik/ntp v0.3.0
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/frankban/quicktest v1.10.2 // indirect
	github.com/go-openapi/errors v0.19.7
	github.com/go-openapi/runtime v0.19.21
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.11
	github.com/google/gopacket v1.1.18
	github.com/google/uuid v1.1.2
	github.com/inconshreveable/log15 v0.0.0-20200109203555-b30bc20e4fd1
	github.com/jaypipes/ghw v0.6.1
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mdlayher/ethernet v0.0.0-20190606142754-0394541c37b7
	github.com/mdlayher/lldp v0.0.0-20150915211757-afd9f83164c5
	github.com/mdlayher/raw v0.0.0-20191009151244-50f2db8cc065
	github.com/metal-stack/go-hal v0.1.10
	github.com/metal-stack/metal-api v0.9.0
	github.com/metal-stack/v v1.0.2
	// archiver must stay in version v2.1.0, see replace below
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/nwaples/rardecode v1.1.0 // indirect
	github.com/pierrec/lz4 v2.5.2+incompatible
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	github.com/u-root/u-root v7.0.0+incompatible
	github.com/vishvananda/netlink v1.1.0
	golang.org/x/sys v0.0.0-20200916084744-dbad9cb7cb7a
	google.golang.org/grpc v1.32.0
	gopkg.in/cheggaaa/pb.v1 v1.0.28
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/mholt/archiver => github.com/mholt/archiver v2.1.0+incompatible
