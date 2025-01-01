package models
type AppManagementRestrictionState int

const (
    ENABLED_APPMANAGEMENTRESTRICTIONSTATE AppManagementRestrictionState = iota
    DISABLED_APPMANAGEMENTRESTRICTIONSTATE
    UNKNOWNFUTUREVALUE_APPMANAGEMENTRESTRICTIONSTATE
)

func (i AppManagementRestrictionState) String() string {
    return []string{"enabled", "disabled", "unknownFutureValue"}[i]
}
func ParseAppManagementRestrictionState(v string) (any, error) {
    result := ENABLED_APPMANAGEMENTRESTRICTIONSTATE
    switch v {
        case "enabled":
            result = ENABLED_APPMANAGEMENTRESTRICTIONSTATE
        case "disabled":
            result = DISABLED_APPMANAGEMENTRESTRICTIONSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPMANAGEMENTRESTRICTIONSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppManagementRestrictionState(values []AppManagementRestrictionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppManagementRestrictionState) isMultiValue() bool {
    return false
}
