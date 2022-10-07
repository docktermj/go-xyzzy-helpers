/*
Package helper ...
*/
package g2configuration

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/docktermj/go-xyzzy-helpers/logger"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getOsEnv(variableName string) (string, error) {
	var err error = nil
	result, isSet := os.LookupEnv(variableName)
	if !isSet {
		err = logger.BuildError(MessageIdFormat, 2990, "Environment variable not set.", variableName)
	}
	return result, err
}

func buildSpecificDatabaseUrl(databaseUrl string) (string, error) {
	result := ""
	parsedUrl, err := url.Parse(databaseUrl)
	if err != nil {
		return "", err
	}

	switch parsedUrl.Scheme {
	case "db2":
		result = "FIXME: Not implemented"
	case "mssql":
		result = "FIXME: Not implemented"
	case "mysql":
		result = "FIXME: Not implemented"
	case "postgresql":
		result = fmt.Sprintf(
			"%s://%s@%s:%s/",
			parsedUrl.Scheme,
			parsedUrl.User,
			parsedUrl.Host,
			string(parsedUrl.Path[1:]),
		)
	case "sqlite3":
		result = "FIXME: Not implemented"
	default:
		result = ""
	}

	return result, err
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Get a Senzing configuration for a "system install" with single database.
func BuildSimpleSystemConfigurationJson(senzingDatabaseUrl string) (string, error) {
	var err error = nil

	if len(senzingDatabaseUrl) == 0 {

		// If SENZING_ENGINE_CONFIGURATION_JSON is set, use it.

		senzingEngineConfigurationJson, err := getOsEnv("SENZING_ENGINE_CONFIGURATION_JSON")
		if err == nil {
			return senzingEngineConfigurationJson, err
		}

		senzingDatabaseUrl, err = getOsEnv("SENZING_TOOLS_DATABASE_URL")
		if err != nil {
			return "", err
		}
	}

	specificDatabaseUrl, specificDatabaseUrlErr := buildSpecificDatabaseUrl(senzingDatabaseUrl)
	if specificDatabaseUrlErr != nil {
		return "", specificDatabaseUrlErr
	}

	resultStruct := G2Configuration{
		Pipeline: G2ConfigurationPipeline{
			ConfigPath:   "/etc/opt/senzing",
			ResourcePath: "/opt/senzing/g2/resources",
			SupportPath:  "/opt/senzing/data",
		},
		Sql: G2ConfigurationSql{
			Connection: specificDatabaseUrl,
		},
	}

	resultBytes, _ := json.Marshal(resultStruct)
	return string(resultBytes), err
}
