module doorayctl

go 1.22.2

replace github.com/dooray-go/dooray v0.2.0 => ../dooray-go

//replace github.com/thediveo/klo v1.1.0 => ../klo

require (
	github.com/dooray-go/dooray v0.2.0
	github.com/spf13/cobra v1.10.1
	github.com/zbum/klo v1.2.0
)

require (
	github.com/clipperhouse/uax29/v2 v2.2.0 // indirect
	github.com/fvbommel/sortorder v1.1.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.19 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	k8s.io/client-go v0.30.5 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
