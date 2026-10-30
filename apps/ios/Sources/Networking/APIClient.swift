import Foundation

/// The typed API client. Builds the request, attaches the device bearer token,
/// unwraps the `{ok,data,error,request_id}` envelope and maps failures to `APIError`.
public struct APIClient: Sendable {
    public static let requestIDHeader = "X-Request-ID"

    public let baseURL: URL
    private let transport: any HTTPTransport
    private let tokens: any BearerTokenSource

    public init(baseURL: URL, transport: any HTTPTransport, tokens: any BearerTokenSource = NoBearerToken()) {
        self.baseURL = baseURL
        self.transport = transport
        self.tokens = tokens
    }

    /// The client for a configuration: live over URLSession, or bundled fixtures.
    public static func make(
        _ configuration: APIConfiguration,
        tokens: any BearerTokenSource,
        bundle: Bundle = .main
    ) -> APIClient {
        let transport: any HTTPTransport = switch configuration.mode {
        case .live: URLSessionTransport()
        case .fixture: FixtureTransport.bundled(in: bundle) ?? FixtureTransport(directory: bundle.bundleURL)
        }
        return APIClient(baseURL: configuration.baseURL, transport: transport, tokens: tokens)
    }

    public func send<Response>(_ endpoint: Endpoint<Response>) async throws(APIError) -> Response {
        let request = try makeRequest(endpoint)
        let (data, response) = try await transport.send(request)
        return try decode(Response.self, data: data, response: response)
    }

    func makeRequest<Response>(_ endpoint: Endpoint<Response>) throws(APIError) -> URLRequest {
        var components = URLComponents(url: baseURL, resolvingAgainstBaseURL: false)
        components?.path = baseURL.path.trimmingSuffix("/") + endpoint.path
        components?.queryItems = endpoint.query.isEmpty ? nil : endpoint.query
        guard let url = components?.url else {
            throw .transport("bad URL for \(endpoint.path)")
        }
        var request = URLRequest(url: url)
        request.httpMethod = endpoint.method.rawValue
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        if let body = endpoint.body {
            request.httpBody = body
            request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        }
        if endpoint.authenticated, let token = tokens.bearerToken() {
            request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        }
        return request
    }

    func decode<Response: Decodable & Sendable>(
        _: Response.Type,
        data: Data,
        response: HTTPURLResponse
    ) throws(APIError) -> Response {
        let status = response.statusCode
        let success = (200..<300).contains(status)
        if data.isEmpty, success, let empty = Empty() as? Response {
            return empty
        }
        let envelope: Envelope<Response>
        do {
            envelope = try WireFormat.decoder().decode(Envelope<Response>.self, from: data)
        } catch {
            if status == 401 { throw .unauthorized(requestID: nil) }
            if !success { throw .unexpectedStatus(status) }
            throw .decoding(String(describing: error))
        }
        if status == 401 {
            throw .unauthorized(requestID: envelope.requestId)
        }
        guard envelope.ok, success else {
            let body = envelope.error ?? ErrorBody(code: "unknown", message: "Something went wrong.")
            throw .server(status: status, body: body, requestID: envelope.requestId)
        }
        if let payload = envelope.data { return payload }
        if let empty = Empty() as? Response { return empty }
        throw .decoding("envelope ok with no data")
    }
}

private extension String {
    func trimmingSuffix(_ suffix: String) -> String {
        hasSuffix(suffix) ? String(dropLast(suffix.count)) : self
    }
}
