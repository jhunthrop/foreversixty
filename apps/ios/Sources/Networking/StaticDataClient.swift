import Foundation
import Models

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

    public func get<Response: Decodable & Sendable>(_: Response.Type, path: String) async throws(APIError) -> Response {
        precondition(path.hasPrefix("/data/"), "Static paths live under /data/: \(path)")
        var request = URLRequest(url: baseURL.appending(path: path))
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        let (data, response) = try await transport.send(request)
        guard (200..<300).contains(response.statusCode) else {
            throw .unexpectedStatus(response.statusCode)
        }
        do {
            return try WireFormat.decoder().decode(Response.self, from: data)
        } catch {
            throw .decoding("\(path): \(error)")
        }
    }

    /// The build to read the lists and BiS from.
    public func activeBuild() async throws(APIError) -> ActiveBuild {
        try await get(ActiveBuild.self, path: "/data/active-build.json")
    }
}
