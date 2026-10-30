import Foundation
import Networking

/// Answers every request with one canned response and remembers the requests.
public final class RecordingTransport: HTTPTransport, @unchecked Sendable {
    private let lock = NSLock()
    private var recorded: [URLRequest] = []
    private let status: Int
    private let body: Data

    public init(status: Int = 200, body: String) {
        self.status = status
        self.body = Data(body.utf8)
    }

    public var requests: [URLRequest] { lock.withLock { recorded } }

    public func send(_ request: URLRequest) async throws(APIError) -> (Data, HTTPURLResponse) {
        lock.withLock { recorded.append(request) }
        let response = HTTPURLResponse(url: request.url!, statusCode: status, httpVersion: nil, headerFields: nil)!
        return (body, response)
    }
}

public struct FixedToken: BearerTokenSource {
    public let token: String?
    public init(token: String?) { self.token = token }
    public func bearerToken() -> String? { token }
    public func discardBearerToken() {}
}
