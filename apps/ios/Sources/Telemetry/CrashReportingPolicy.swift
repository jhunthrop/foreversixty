import Foundation

/// Whether this build reports crashes and hangs, and with what. Reporting runs only in a
/// release build talking to the live API with a DSN configured: never in debug, never in
/// fixture mode, never with the placeholder empty DSN.
public struct CrashReportingPolicy: Sendable, Equatable {
    public static let dsnKey = "FSSentryDSN"
    /// A main-thread stall this long is reported as a hang.
    public static let hangThreshold: TimeInterval = 2

    public let dsn: String
    public let release: String
    public let environment: String

    /// Nil when reporting must stay off for this build.
    public init?(dsn: String?, isDebugBuild: Bool, isLiveAPI: Bool, version: String, build: String) {
        guard !isDebugBuild, isLiveAPI, let dsn, !dsn.isEmpty, dsn.hasPrefix("https://") else { return nil }
        self.dsn = dsn
        self.release = "gg.foreversixty.app@\(version)+\(build)"
        self.environment = "production"
    }
}
