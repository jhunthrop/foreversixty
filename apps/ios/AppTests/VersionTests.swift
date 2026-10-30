import Foundation
import Testing

/// The built app carries exactly the version in apps/VERSION, the file both apps read.
@Suite("App version")
struct VersionTests {
    private static let versionFile = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // AppTests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .appending(path: "VERSION")

    /// The `KEY = value` lines of apps/VERSION, skipping `//` comments.
    private static func settings() throws -> [String: String] {
        let text = try String(contentsOf: versionFile, encoding: .utf8)
        var values: [String: String] = [:]
        for line in text.split(separator: "\n") where !line.hasPrefix("//") {
            let parts = line.split(separator: "=", maxSplits: 1).map { $0.trimmingCharacters(in: .whitespaces) }
            if parts.count == 2 { values[parts[0]] = parts[1] }
        }
        return values
    }

    @Test("CFBundleShortVersionString and CFBundleVersion come from apps/VERSION")
    func bundleMatchesVersionFile() throws {
        let file = try Self.settings()
        let info = Bundle.main.infoDictionary ?? [:]
        #expect(info["CFBundleShortVersionString"] as? String == file["APP_VERSION"])
        #expect(info["CFBundleVersion"] as? String == file["APP_BUILD"])
    }
}
