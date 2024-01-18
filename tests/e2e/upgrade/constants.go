package upgrade

// The constants used in the upgrade tests are defined here
const (
	// the defaultChainID used for testing
	defaultChainID = "gurux_9000-1"

	// LocalVersionTag defines the docker image ImageTag when building locally
	LocalVersionTag = "latest"

	// tharsisRepo is the docker hub repository that contains the Gurux images pulled during tests
	tharsisRepo = "tharsishq/gurux"

	// upgradesPath is the relative path from this folder to the app/upgrades folder
	upgradesPath = "../../../app/upgrades"

	// versionSeparator is used to separate versions in the INITIAL_VERSION and TARGET_VERSION
	// environment vars
	versionSeparator = "/"
)
