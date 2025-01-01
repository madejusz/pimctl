package models
type X509CertificateCRLValidationConfigurationState int

const (
    DISABLED_X509CERTIFICATECRLVALIDATIONCONFIGURATIONSTATE X509CertificateCRLValidationConfigurationState = iota
    ENABLED_X509CERTIFICATECRLVALIDATIONCONFIGURATIONSTATE
    UNKNOWNFUTUREVALUE_X509CERTIFICATECRLVALIDATIONCONFIGURATIONSTATE
)

func (i X509CertificateCRLValidationConfigurationState) String() string {
    return []string{"disabled", "enabled", "unknownFutureValue"}[i]
}
func ParseX509CertificateCRLValidationConfigurationState(v string) (any, error) {
    result := DISABLED_X509CERTIFICATECRLVALIDATIONCONFIGURATIONSTATE
    switch v {
        case "disabled":
            result = DISABLED_X509CERTIFICATECRLVALIDATIONCONFIGURATIONSTATE
        case "enabled":
            result = ENABLED_X509CERTIFICATECRLVALIDATIONCONFIGURATIONSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_X509CERTIFICATECRLVALIDATIONCONFIGURATIONSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeX509CertificateCRLValidationConfigurationState(values []X509CertificateCRLValidationConfigurationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i X509CertificateCRLValidationConfigurationState) isMultiValue() bool {
    return false
}
