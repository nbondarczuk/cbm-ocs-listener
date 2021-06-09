package common

// start up of the environment
func StartUp() {
	// Start a SQL DB session to e used by repositories
	createOracleDbSession()
}
