// The g2configuration package is a set of method to help with common tasks.
//
// The purpose of a g2configuration object is:
//   • ...
//   • ...
//   • ...
package g2configuration

import ()

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6510%04d"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2ConfigurationPipeline struct {
	ConfigPath   string `json:"CONFIGPATH"`
	ResourcePath string `json:"RESOURCEPATH"`
	SupportPath  string `json:"SUPPORTPATH"`
}

type G2ConfigurationSql struct {
	Connection string `json:"CONNECTION"`
}

type G2Configuration struct {
	Pipeline G2ConfigurationPipeline `json:"PIPELINE"`
	Sql      G2ConfigurationSql      `json:"SQL"`
}
