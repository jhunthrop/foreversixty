import Foundation
import Models

/// One static file of the site and the type it decodes to. The members are generated from
/// api/openapi.yaml's `app-static` operations (Generated/APIEndpoints.swift).
public struct StaticResource<Response: Decodable & Sendable>: Sendable {
    /// Absolute path under `/data/`, e.g. "/data/active-build.json".
    public let path: String

    public init(path: String) {
        precondition(path.hasPrefix("/data/"), "Static paths live under /data/: \(path)")
        self.path = path
    }
}

/// Reads the site's static data (`https://foreversixty.gg/data/...`, docs/contracts/app-v1.md
/// "Static files"). These files are bare JSON with no envelope, so they never go through
/// `APIClient`; fixture mode answers them from the `cdn.*` fixtures.
public struct StaticDataClient: Sendable {
    public static let siteURL = URL(string: "https://foreversixty.gg")!

    public let baseURL: URL
    private let transport: any HTTPTransport

    public init(baseURL: URL = Self.siteURL, transport: any HTTPTransport) {
        self.baseURL = baseURL
        self.transport = transport
    }

    public func get<Response>(_ resource: StaticResource<Response>) async throws(APIError) -> Response {
        try await fetch(resource).value
    }

    /// The decoded file and the body it came from, for a caller that keeps the body
    /// (`CachedStaticDataClient`).
    public func fetch<Response>(
        _ resource: StaticResource<Response>
    ) async throws(APIError) -> (value: Response, body: Data) {
        var request = URLRequest(url: baseURL.appending(path: resource.path))
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        let (data, response) = try await transport.send(request)
        guard (200..<300).contains(response.statusCode) else {
            throw .unexpectedStatus(response.statusCode)
        }
        return try (Self.decode(resource, from: data), data)
    }

    /// Decodes a body of `resource`, fetched now or saved earlier.
    public static func decode<Response>(
        _ resource: StaticResource<Response>,
        from data: Data
    ) throws(APIError) -> Response {
        do {
            return try WireFormat.decoder().decode(Response.self, from: data)
        } catch {
            throw .decoding("\(resource.path): \(error)")
        }
    }

    /// The build to read the lists and BiS from.
    public func activeBuild() async throws(APIError) -> ActiveBuild {
        try await get(.getActiveBuild)
    }
}
