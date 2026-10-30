import Foundation

/// Where requests go: the live API, or bundled contract fixtures.
public enum APIMode: String, Sendable {
    case live
    case fixture
}

/// The API's base URL and mode, read from the build settings
/// `API_BASE_URL` and `API_MODE` through Info.plist (project.yml).
public struct APIConfiguration: Sendable, Equatable {
    public static let baseURLKey = "FSAPIBaseURL"
    public static let modeKey = "FSAPIMode"

    public let baseURL: URL
    public let mode: APIMode

    public init(baseURL: URL, mode: APIMode) {
        self.baseURL = baseURL
        self.mode = mode
    }

    /// Reads the Info.plist keys; fails fast on a missing or malformed value,
    /// since a build without them is misconfigured.
    public init(infoDictionary: [String: Any]) throws(ConfigurationError) {
        guard let raw = infoDictionary[Self.baseURLKey] as? String, !raw.isEmpty else {
            throw .missing(Self.baseURLKey)
        }
        guard let url = URL(string: raw), url.scheme == "https" || url.host == "localhost" else {
            throw .invalid(Self.baseURLKey, raw)
        }
        let rawMode = infoDictionary[Self.modeKey] as? String ?? APIMode.live.rawValue
        guard let mode = APIMode(rawValue: rawMode) else {
            throw .invalid(Self.modeKey, rawMode)
        }
        self.init(baseURL: url, mode: mode)
    }
}

public enum ConfigurationError: Error, Equatable, CustomStringConvertible {
    case missing(String)
    case invalid(String, String)

    public var description: String {
        switch self {
        case let .missing(key): "Info.plist is missing \(key)"
        case let .invalid(key, value): "Info.plist \(key) is not valid: \(value)"
        }
    }
}
