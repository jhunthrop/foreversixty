import Auth

extension RelativeAge {
    /// The one relative-age phrase (`2 h ago`): Home's sync age and You's device rows.
    var phrase: String {
        switch self {
        case .now: Copy.commonAgeNow
        case let .minutes(n): Copy.commonAgeMinutes(n: n)
        case let .hours(n): Copy.commonAgeHours(n: n)
        case let .days(n): Copy.commonAgeDays(n: n)
        }
    }
}
