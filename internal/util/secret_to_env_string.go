package util

func BuildEnvStringFromSecret(secretMap map[string]interface{}) string {
	result := ""
	for k := range secretMap {
		secretValue := k + "=" + secretMap[k].(string) + "\n"
		result = result + secretValue
	}

	return result
}
