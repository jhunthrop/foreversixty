import Foundation

/// What leaves the device in a crash or hang report, made safe first. The same rules run on
/// Android (core/telemetry Scrubber.kt); apps/shared/privacy.md is what they promise.
///
/// - A URL keeps scheme, host and path and loses its query and fragment, so the sign-in
///   callback's `code` and `state` never travel.
/// - Free text loses device tokens (`fsd_…`), bearer credentials and any `code`, `state`,
///   `token`, `code_verifier` or `key` parameter value.
public enum Scrubber {
    public static let redacted = "[redacted]"

    private static let secretPatterns: [(NSRegularExpression, String)] = [
        (#"fsd_[A-Za-z0-9_\-]+"#, "fsd_\(redacted)"),
        (#"(?i)\bbearer\s+[A-Za-z0-9._~+/=\-]+"#, "Bearer \(redacted)"),
        (#"(?i)\b(code|code_verifier|state|token|access_token|refresh_token|key)=[^&\s#"']+"#, "$1=\(redacted)"),
    ].map { pattern, template in
        // The patterns are constants; a bad one is a programming error caught by the tests.
        (try! NSRegularExpression(pattern: pattern), template) // swiftlint:disable:this force_try
    }

    /// The URL without its query and fragment. Text that is not an absolute URL is scrubbed as text.
    public static func url(_ raw: String) -> String {
        guard var components = URLComponents(string: raw), components.scheme != nil, components.host != nil else {
            return text(raw)
        }
        components.query = nil
        components.fragment = nil
        components.user = nil
        components.password = nil
        return components.string ?? text(raw)
    }

    /// Free text with every credential-shaped run replaced.
    public static func text(_ raw: String) -> String {
        secretPatterns.reduce(raw) { result, rule in
            let range = NSRange(result.startIndex..., in: result)
            return rule.0.stringByReplacingMatches(in: result, range: range, withTemplate: rule.1)
        }
    }
}
