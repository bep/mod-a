package p

var (
	name    = "a"
	version = "v1.1.0"
)

func Version() string {
	return name + ": " + version
}
