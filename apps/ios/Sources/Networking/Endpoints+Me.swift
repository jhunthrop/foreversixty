import Foundation
import Models

public extension Endpoint where Response == Me {
    /// `GET /v1/me`: the account, its characters, guilds and entitlements.
    static var me: Endpoint { Endpoint(method: .get, path: "/v1/me") }
}
