import Foundation
import Testing

/// Labels and headings are sentence case on native (spec 2026-11-07 ruling N1; owner 2026-11-10: the apps
/// set caps where the site does not). Outside Sources/DesignSystem no source may force capitals: a code
/// written in capitals (a region, a metric unit) goes through `Codes`, so the one place that uppercases
/// is reviewed once. Android: `CaseGuardTest.kt`.
@Suite struct CaseGuardTests {
    /// apps/ios, from this file's place in AppTests/System.
    private static let ios = URL(fileURLWithPath: #filePath).deletingLastPathComponent().deletingLastPathComponent().deletingLastPathComponent()

    /// The app's and the extensions' sources; Sources/DesignSystem owns `Codes`.
    private static let scanned = ["App", "Sources", "Widgets", "Shared", "NotificationContent"]
    private static let designSystem = "/Sources/DesignSystem/"

    /// Swift's and Foundation's uppercasing and SwiftUI's caps transform.
    private static var forcedCaps: Regex<Substring> { #/\.uppercased\(|localizedUppercase|textCase\(\.uppercase|\.uppercaseSmallCaps/# }

    /// Empty on purpose: a new entry needs the owner's reason.
    private static let allowed: [String] = []

    private static func sources() -> [URL] {
        scanned.flatMap { folder -> [URL] in
            let root = ios.appendingPathComponent(folder)
            guard let walker = FileManager.default.enumerator(at: root, includingPropertiesForKeys: nil) else { return [] }
            return walker.compactMap { $0 as? URL }
                .filter { $0.pathExtension == "swift" && !$0.path.contains(designSystem) }
        }
    }

    @Test func noSourceOutsideTheDesignSystemForcesCapitals() throws {
        let offenders = try Self.sources().flatMap { file -> [String] in
            let lines = try String(contentsOf: file, encoding: .utf8).components(separatedBy: "\n")
            return lines.enumerated().compactMap { index, line in
                line.contains(Self.forcedCaps) ? "\(file.path.replacingOccurrences(of: Self.ios.path + "/", with: "")):\(index + 1)" : nil
            }
        }
        #expect(offenders == Self.allowed, "Forced capitals (use sentence case, or Codes for a code):\n\(offenders.joined(separator: "\n"))")
    }

    @Test func theGuardSeesTheAppsSources() {
        #expect(Self.sources().count > 200)
    }
}
