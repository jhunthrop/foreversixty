import Foundation
import Models

public extension Endpoint where Response == CharacterProfile {
    /// `GET /v1/characters/{region}/{ruleset}/{name}`; `key` is `region/ruleset/slug`.
    static func character(key: String) -> Endpoint {
        Endpoint(method: .get, path: "/v1/characters/\(key)", authenticated: false)
    }
}
