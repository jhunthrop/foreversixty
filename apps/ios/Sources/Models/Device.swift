import Foundation

/// One signed-in install: a companion or an app (api/internal/auth/store.go `Device`).
public struct Device: Codable, Sendable, Equatable, Identifiable {
    public let id: String
    public let name: String
    /// "ios", "android", or a companion's OS.
    public let platform: String
    public let createdAt: Date
    public let lastSeenAt: Date?
}
