import Foundation

/// docs/contracts/fixtures/app in the repo, read in place so tests always see
/// lane C's current fixtures (the simulator can read the host filesystem).
public enum FixtureLocation {
    public static let directory: URL = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // TestSupport
        .deletingLastPathComponent() // Tests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "docs/contracts/fixtures/app")

    public static func data(_ file: String) throws -> Data {
        try Data(contentsOf: directory.appending(path: file))
    }
}
