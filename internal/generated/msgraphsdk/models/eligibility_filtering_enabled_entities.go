package models
import (
    "math"
    "strings"
)
type EligibilityFilteringEnabledEntities int

const (
    NONE_ELIGIBILITYFILTERINGENABLEDENTITIES = 1
    SWAPREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES = 2
    OFFERSHIFTREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES = 4
    UNKNOWNFUTUREVALUE_ELIGIBILITYFILTERINGENABLEDENTITIES = 8
    TIMEOFFREASON_ELIGIBILITYFILTERINGENABLEDENTITIES = 16
)

func (i EligibilityFilteringEnabledEntities) String() string {
    var values []string
    options := []string{"none", "swapRequest", "offerShiftRequest", "unknownFutureValue", "timeOffReason"}
    for p := 0; p < 5; p++ {
        mantis := EligibilityFilteringEnabledEntities(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseEligibilityFilteringEnabledEntities(v string) (any, error) {
    var result EligibilityFilteringEnabledEntities
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_ELIGIBILITYFILTERINGENABLEDENTITIES
            case "swapRequest":
                result |= SWAPREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
            case "offerShiftRequest":
                result |= OFFERSHIFTREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ELIGIBILITYFILTERINGENABLEDENTITIES
            case "timeOffReason":
                result |= TIMEOFFREASON_ELIGIBILITYFILTERINGENABLEDENTITIES
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeEligibilityFilteringEnabledEntities(values []EligibilityFilteringEnabledEntities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EligibilityFilteringEnabledEntities) isMultiValue() bool {
    return true
}
