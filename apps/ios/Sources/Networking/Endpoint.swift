import Foundation

public enum HTTPMethod: String, Sendable {
    case get = "GET"
    case post = "POST"
    case put = "PUT"
    case patch = "PATCH"
    case delete = "DELETE"
}

/// One API route and the payload type its envelope carries. Feature lanes add
/// their routes as static members in an `extension Endpoint`.
public struct Endpoint<Response: Decodable & Sendable>: Sendable {
    public let method: HTTPMethod
    /// Absolute API path, e.g. "/v1/me".
    public let path: String
    public let query: [URLQueryItem]
    public let body: Data?
    /// False for the sign-in exchange, which runs before a token exists.
    public let authenticated: Bool

    public init(
        method: HTTPMethod,
        path: String,
        query: [URLQueryItem] = [],
        body: Data? = nil,
        authenticated: Bool = true
    ) {
        precondition(path.hasPrefix("/"), "Endpoint path must be absolute: \(path)")
        self.method = method
        self.path = path
        self.query = query
        self.body = body
        self.authenticated = authenticated
    }

    /// A JSON request body in the API's wire format.
    public static func json(
        _ method: HTTPMethod,
        _ path: String,
        body: some Encodable,
        authenticated: Bool = true
    ) throws(APIError) -> Endpoint {
        do {
            let data = try WireFormat.encoder().encode(body)
            return Endpoint(method: method, path: path, body: data, authenticated: authenticated)
        } catch {
            throw .decoding("could not encode \(path) body: \(error)")
        }
    }
}
