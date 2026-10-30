import Foundation

/// The API's minimum-version gate: `426` with code `update_required` and the store page in
/// `error.fields.store_url` (docs/contracts/app-v1.md "Forced update").
public enum UpdateRequired {
    public static let status = 426
    public static let code = "update_required"
    public static let storeURLField = "store_url"

    /// One gate answer. `storeURL` is nil when the gate sent no https link.
    public struct Answer: Sendable, Equatable {
        public let storeURL: URL?
    }

    /// The gate's answer when `data` is one, else nil. Read with a plain decoder:
    /// WireFormat's snake-case strategy would rename the `store_url` key inside `fields`.
    public static func answer(status: Int, data: Data) -> Answer? {
        guard status == Self.status,
              let envelope = try? JSONDecoder().decode(GateEnvelope.self, from: data),
              envelope.error?.code == Self.code
        else { return nil }
        let raw = envelope.error?.fields?[storeURLField]
        let url = raw.flatMap(URL.init(string:)).flatMap { $0.scheme == "https" ? $0 : nil }
        return Answer(storeURL: url)
    }

    private struct GateEnvelope: Decodable {
        struct Body: Decodable {
            let code: String
            let fields: [String: String]?
        }

        let error: Body?
    }
}

/// Told when any request answers `update_required`; the app blocks on it.
public protocol UpdateRequiredSink: Sendable {
    func updateRequired(storeURL: URL?)
}

/// Adds the client identity headers to every request and reports the update gate.
/// Wraps the live or fixture transport, so APIClient itself stays header-agnostic.
public struct IdentifiedTransport: HTTPTransport {
    private let base: any HTTPTransport
    private let identity: ClientIdentity
    private let sink: (any UpdateRequiredSink)?

    public init(_ base: any HTTPTransport, identity: ClientIdentity, sink: (any UpdateRequiredSink)? = nil) {
        self.base = base
        self.identity = identity
        self.sink = sink
    }

    public func send(_ request: URLRequest) async throws(APIError) -> (Data, HTTPURLResponse) {
        var identified = request
        identified.setValue(identity.appHeaderValue, forHTTPHeaderField: ClientIdentity.appHeader)
        identified.setValue(identity.installID, forHTTPHeaderField: ClientIdentity.installHeader)
        let (data, response) = try await base.send(identified)
        if let sink, let gate = UpdateRequired.answer(status: response.statusCode, data: data) {
            sink.updateRequired(storeURL: gate.storeURL)
        }
        return (data, response)
    }
}
