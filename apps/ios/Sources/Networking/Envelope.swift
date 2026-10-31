import Foundation

/// Every API response body: `{ok, data, error, request_id}`
/// (api/internal/httpx/envelope.go).
public struct Envelope<Payload: Decodable & Sendable>: Decodable, Sendable {
    public let ok: Bool
    public let data: Payload?
    public let error: ErrorBody?
    public let requestId: String?
}

/// The envelope's error half (Go `httpx.ErrorBody`).
public struct ErrorBody: Codable, Sendable, Equatable {
    public let code: String
    public let message: String
    public let fields: [String: String]?

    public init(code: String, message: String, fields: [String: String]? = nil) {
        self.code = code
        self.message = message
        self.fields = fields
    }
}

/// A response with no payload worth reading (a 204, or `data: null`).
public struct Empty: Decodable, Sendable, Equatable {
    public init() {}
}
