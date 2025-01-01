package security
type SensorType int

const (
    ADCONNECTINTEGRATED_SENSORTYPE SensorType = iota
    ADCSINTEGRATED_SENSORTYPE
    ADFSINTEGRATED_SENSORTYPE
    DOMAINCONTROLLERINTEGRATED_SENSORTYPE
    DOMAINCONTROLLERSTANDALONE_SENSORTYPE
    UNKNOWNFUTUREVALUE_SENSORTYPE
)

func (i SensorType) String() string {
    return []string{"adConnectIntegrated", "adcsIntegrated", "adfsIntegrated", "domainControllerIntegrated", "domainControllerStandalone", "unknownFutureValue"}[i]
}
func ParseSensorType(v string) (any, error) {
    result := ADCONNECTINTEGRATED_SENSORTYPE
    switch v {
        case "adConnectIntegrated":
            result = ADCONNECTINTEGRATED_SENSORTYPE
        case "adcsIntegrated":
            result = ADCSINTEGRATED_SENSORTYPE
        case "adfsIntegrated":
            result = ADFSINTEGRATED_SENSORTYPE
        case "domainControllerIntegrated":
            result = DOMAINCONTROLLERINTEGRATED_SENSORTYPE
        case "domainControllerStandalone":
            result = DOMAINCONTROLLERSTANDALONE_SENSORTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SENSORTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSensorType(values []SensorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SensorType) isMultiValue() bool {
    return false
}
