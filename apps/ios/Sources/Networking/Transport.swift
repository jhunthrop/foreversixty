import Foundation

/// Sends one request and returns the raw body and status. The client's only
/// seam to the outside: live mode uses URLSession, fixture mode reads files,
/// tests inject their own.
public protocol HTTPTransport: Sendable {
    func send(_ request: URLRequest) async throws(APIError) -> (Data, HTTPURLResponse)
}

/// The live transport.
public struct URLSessionTransport: HTTPTransport {
    private let session: URLSession

    public init(session: URLSession = .shared) {
        self.session = session
    }

    public func send(_ request: URLRequest) async throws(APIError) -> (Data, HTTPURLResponse) {
        let data: Data
        let response: URLResponse
        do {
            (data, response) = try await session.data(for: request)
        } catch {
            throw .transport(error.localizedDescription)
        }
        guard let http = response as? HTTPURLResponse else {
            throw .transport("not an HTTP response")
        }
        return (data, http)
    }
}

/// Supplies the device bearer token (`fsd_…`) for authenticated requests.
/// Auth's token store conforms; Networking never reads the Keychain itself.
public protocol BearerTokenSource: Sendable {
    func bearerToken() -> String?
}

/// No token: every authenticated request goes out without one and the API answers 401.
public struct NoBearerToken: BearerTokenSource {
    public init() {}
    public func bearerToken() -> String? { nil }
}
