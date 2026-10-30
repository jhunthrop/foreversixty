import Foundation

/// One route answered from a contract fixture file.
public struct FixtureRoute: Sendable, Hashable {
    public let method: HTTPMethod
    public let path: String
    public let file: String
    public let status: Int

    public init(_ method: HTTPMethod, _ path: String, file: String, status: Int = 200) {
        self.method = method
        self.path = path
        self.file = file
        self.status = status
    }
}

/// The route-to-fixture table for docs/contracts/fixtures/app (lane C owns the
/// files; docs/contracts/app-v1.md names them). Each fixture is a full response
/// body, envelope included. Feature lanes add a row per route they call.
public enum FixtureCatalog {
    public static let app: [FixtureRoute] = [
        FixtureRoute(.get, "/v1/me", file: "me.json"),
        FixtureRoute(.post, "/v1/auth/app/exchange", file: "auth-app-exchange.json"),
    ]
}

/// Answers requests from fixture files in a directory: the app bundle's copy
/// in fixture mode, the repo's docs/contracts/fixtures/app in tests. No network.
public struct FixtureTransport: HTTPTransport {
    private let directory: URL
    private let routes: [String: FixtureRoute]

    public init(directory: URL, routes: [FixtureRoute] = FixtureCatalog.app) {
        self.directory = directory
        self.routes = Dictionary(routes.map { (Self.key($0.method.rawValue, $0.path), $0) }) { _, last in last }
    }

    /// The bundled copy (project.yml copies docs/contracts/fixtures/app in as "Fixtures").
    public static func bundled(in bundle: Bundle = .main) -> FixtureTransport? {
        bundle.url(forResource: "Fixtures", withExtension: nil).map { FixtureTransport(directory: $0) }
    }

    public func send(_ request: URLRequest) async throws(APIError) -> (Data, HTTPURLResponse) {
        let method = request.httpMethod ?? HTTPMethod.get.rawValue
        let path = request.url?.path ?? ""
        let key = Self.key(method, path)
        guard let route = routes[key] else {
            throw .missingFixture(key)
        }
        let file = directory.appending(path: route.file)
        guard let data = try? Data(contentsOf: file) else {
            throw .missingFixture("\(key) -> \(route.file)")
        }
        guard let url = request.url,
              let response = HTTPURLResponse(url: url, statusCode: route.status, httpVersion: "HTTP/1.1", headerFields: nil)
        else {
            throw .transport("could not build a fixture response for \(key)")
        }
        return (data, response)
    }

    private static func key(_ method: String, _ path: String) -> String {
        "\(method) \(path)"
    }
}
