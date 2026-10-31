import Foundation

/// One route answered from a contract fixture. `pattern` is the path with `{name}` segments;
/// `fixture` is lane C's file stem (`<endpoint>[.<variant>]`, docs/contracts/app-v1.md), and
/// may name a path segment (`cdn.bis.{spec}`). A nil fixture answers an empty body (a 204).
public struct FixtureRoute: Sendable, Hashable {
    public let method: HTTPMethod
    public let pattern: String
    public let fixture: String?
    public let status: Int

    public init(_ method: HTTPMethod, _ pattern: String, fixture: String?, status: Int = 200) {
        self.method = method
        self.pattern = pattern
        self.fixture = fixture
        self.status = status
    }

    /// The path's `{name}` values when this route serves the request, else nil.
    func match(method: String, path: String) -> [String: String]? {
        guard method == self.method.rawValue else { return nil }
        let want = Self.segments(pattern)
        let got = Self.segments(path)
        guard want.count == got.count else { return nil }
        var params: [String: String] = [:]
        for (expected, actual) in zip(want, got) {
            if let (name, suffix) = Self.placeholder(expected) {
                guard actual.hasSuffix(suffix), actual.count > suffix.count else { return nil }
                params[name] = String(actual.dropLast(suffix.count))
            } else if expected != actual {
                return nil
            }
        }
        return params
    }

    /// `{spec}.json` is the placeholder `spec` followed by the literal `.json`.
    private static func placeholder(_ segment: Substring) -> (String, Substring)? {
        guard segment.hasPrefix("{"), let close = segment.firstIndex(of: "}") else { return nil }
        let name = segment[segment.index(after: segment.startIndex)..<close]
        return (String(name), segment[segment.index(after: close)...])
    }

    private static func segments(_ path: String) -> [Substring] {
        path.split(separator: "/", omittingEmptySubsequences: true)
    }
}

/// Which state of an endpoint fixture mode answers: lane C's variants and error variants.
public enum FixtureVariant: Sendable, Hashable {
    /// A success variant: `me` with `.named("empty")` is `me.empty.json`.
    case named(String)
    /// An error variant: `me.patch` with `.error("invalid", status: 400)` is
    /// `me.patch.error-invalid.json`.
    case error(String, status: Int)
    /// The shared `error-unauthorized.json`: a missing, revoked or expired token on any route.
    case unauthorized

    /// The file stem for a route's own stem.
    func file(_ stem: String) -> String {
        switch self {
        case let .named(variant): "\(stem).\(variant)"
        case let .error(code, _): "\(stem).error-\(code)"
        case .unauthorized: "error-unauthorized"
        }
    }

    /// The status to answer with; nil keeps the route's success status.
    var status: Int? {
        switch self {
        case .named: nil
        case let .error(_, status): status
        case .unauthorized: 401
        }
    }
}

/// The route-to-fixture table for docs/contracts/fixtures/app, generated from api/openapi.yaml
/// (`ContractRoutes`, web/scripts/gen-mobile-api.mjs): a route's default fixture is the first of
/// its `x-fixtures`, or its `x-fixture-route`. API fixtures are whole envelopes; `cdn.*` fixtures
/// are the static files as served. A new route joins by being added to the spec.
public enum FixtureCatalog {
    public static let api: [FixtureRoute] = ContractRoutes.api

    /// The site's static data (`https://foreversixty.gg/data/...`); bare JSON, no envelope.
    public static let cdn: [FixtureRoute] = ContractRoutes.staticData

    public static let app: [FixtureRoute] = api + cdn
}

/// Answers requests from fixture files in a directory: the app bundle's copy in fixture mode,
/// the repo's docs/contracts/fixtures/app in tests. No network. `variants` picks a state per
/// route stem (`["me": .named("empty")]`); every other route answers its success fixture.
public struct FixtureTransport: HTTPTransport {
    private let directory: URL
    private let routes: [FixtureRoute]
    private let variants: [String: FixtureVariant]

    public init(
        directory: URL,
        routes: [FixtureRoute] = FixtureCatalog.app,
        variants: [String: FixtureVariant] = [:]
    ) {
        self.directory = directory
        self.routes = routes
        self.variants = variants
    }

    /// The folder project.yml copies docs/contracts/fixtures/app into. A folder reference keeps
    /// its name on disk in the bundle (XcodeGen's `name` only labels the group), so it is `app`.
    public static let bundledFolder = "app"

    /// The bundled copy of the contract fixtures.
    public static func bundled(in bundle: Bundle = .main) -> FixtureTransport? {
        bundle.url(forResource: bundledFolder, withExtension: nil).map { FixtureTransport(directory: $0) }
    }

    public func send(_ request: URLRequest) async throws(APIError) -> (Data, HTTPURLResponse) {
        let method = request.httpMethod ?? HTTPMethod.get.rawValue
        let path = request.url?.path ?? ""
        let key = "\(method) \(path)"
        guard let (route, params) = resolve(method: method, path: path) else {
            throw .missingFixture(key)
        }
        let (data, status) = try answer(route, params: params, key: key)
        guard let url = request.url,
              let response = HTTPURLResponse(url: url, statusCode: status, httpVersion: "HTTP/1.1", headerFields: nil)
        else {
            throw .transport("could not build a fixture response for \(key)")
        }
        return (data, response)
    }

    /// The fixture file a request reads, or nil for a 204 or an unknown route (for tests).
    public func fileName(method: HTTPMethod, path: String) -> String? {
        guard let (route, params) = resolve(method: method.rawValue, path: path) else { return nil }
        return file(for: route, params: params).map { "\($0).json" }
    }

    private func resolve(method: String, path: String) -> (FixtureRoute, [String: String])? {
        for route in routes {
            if let params = route.match(method: method, path: path) { return (route, params) }
        }
        return nil
    }

    private func file(for route: FixtureRoute, params: [String: String]) -> String? {
        guard let stem = route.fixture else { return nil }
        let named = params.reduce(stem) { $0.replacingOccurrences(of: "{\($1.key)}", with: $1.value) }
        return variants[stem].map { $0.file(named) } ?? named
    }

    private func answer(_ route: FixtureRoute, params: [String: String], key: String) throws(APIError) -> (Data, Int) {
        let variant = route.fixture.flatMap { variants[$0] }
        let status = variant?.status ?? route.status
        guard let stem = file(for: route, params: params) else { return (Data(), status) }
        guard let data = try? Data(contentsOf: directory.appending(path: "\(stem).json")) else {
            throw .missingFixture("\(key) -> \(stem).json")
        }
        return (data, status)
    }
}
