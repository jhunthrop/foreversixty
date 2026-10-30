import Foundation

/// One signed-in install: a companion or an app (api/internal/auth/store.go `Device`).
public struct Device: Codable, Sendable, Equatable, Identifiable {
    public let id: String
    public let name: String
    /// "ios", "android", or a companion's OS.
    public let platform: String
    /// "companion" or "app"; `GET /v1/devices` sends it, the sign-in exchange does not.
    public let kind: String?
    /// Sent by `GET /v1/devices`; the sign-in exchange carries only id, name and platform.
    public let createdAt: Date?
    public let lastSeenAt: Date?
}
