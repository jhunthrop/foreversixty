import Foundation

/// The signed-in account, as `GET /v1/me` serves `user` (api/internal/auth/store.go `User`).
public struct User: Codable, Sendable, Equatable {
    public let id: Int64
    /// Nil for an account signed in by email.
    public let battletag: String?
    public let email: String?
    public let role: String
    public let anonymize: Bool

    public init(id: Int64, battletag: String?, email: String?, role: String, anonymize: Bool) {
        self.id = id
        self.battletag = battletag
        self.email = email
        self.role = role
        self.anonymize = anonymize
    }
}
