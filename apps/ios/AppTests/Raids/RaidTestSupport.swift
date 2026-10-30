@testable import ForeverSixty
import Foundation
import Models
import Networking

/// The contract fixtures, read in place from docs/contracts/fixtures/app (the simulator reads the
/// host's files). Every helper throws, so a missing or misnamed fixture fails only its own test.
enum RaidFixture {
    struct Missing: Error, CustomStringConvertible {
        let name: String
        let reason: String
        var description: String { "fixture \(name): \(reason)" }
    }

    static let directory: URL = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // Raids
        .deletingLastPathComponent() // AppTests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "docs/contracts/fixtures/app")

    static func data(_ name: String) throws -> Data {
        do {
            return try Data(contentsOf: directory.appending(path: "\(name).json"))
        } catch {
            throw Missing(name: name, reason: "\(error)")
        }
    }

    static func load<Value: Decodable>(_ type: Value.Type, _ name: String) throws -> Value {
        try WireFormat.decoder().decode(Envelope<Value>.self, from: data(name)).data
    }

    static func error(_ name: String, status: Int) throws -> APIError {
        let body = try WireFormat.decoder().decode(ErrorEnvelope.self, from: data(name)).error
        return .server(status: status, body: body, requestID: nil)
    }

    static func raid(_ name: String = "guild.raid") throws -> RaidView { try load(RaidView.self, name) }

    /// The same load as the API would fail it, for the fake service's typed-throws closures.
    static func api<Value: Decodable>(_ type: Value.Type, _ name: String) throws(APIError) -> Value {
        do { return try load(type, name) } catch { throw .missingFixture("\(error)") }
    }

    private struct Envelope<Value: Decodable>: Decodable { let data: Value }
    private struct ErrorEnvelope: Decodable { let error: ErrorBody }
}

/// Fixed instants: two days before the fixture raid (invites 2026-12-20 00:45Z, pull 01:00Z).
enum RaidTime {
    static let chicago = TimeZone(identifier: "America/Chicago") ?? .gmt
    static let clock = RaidClock(zone: chicago)
    static let now = date("2026-12-17T22:30:00Z")

    /// A fixed instant; a typo reads as the reference date rather than crashing the host.
    static func date(_ text: String) -> Date { (try? Date.ISO8601FormatStyle().parse(text)) ?? .distantPast }
}

/// A scripted raid service: each call answers from its closure; every sign-up is recorded.
final class FakeRaidService: RaidService, @unchecked Sendable {
    var meAnswer: () throws(APIError) -> Me = { () throws(APIError) in try RaidFixture.api(Me.self, "me") }
    var raidsAnswer: () throws(APIError) -> GuildRaids = { () throws(APIError) in try RaidFixture.api(GuildRaids.self, "guild.raids") }
    var raidAnswer: () throws(APIError) -> RaidView = { try RaidFixture.api(RaidView.self, "guild.raid") }
    var signupAnswer: (SignupRequest) throws(APIError) -> SignupAnswer = { _ throws(APIError) in try RaidFixture.api(SignupAnswer.self, "guild.signup") }
    var lockAnswer: () throws(APIError) -> RaidView = { () throws(APIError) in try RaidFixture.api(RaidView.self, "guild.raid.locked") }
    var nightAnswer: () throws(APIError) -> RaidNightSummary = { () throws(APIError) in try RaidFixture.api(RaidNightSummary.self, "me.raid-night") }
    private(set) var signups: [(key: String, request: SignupRequest)] = []
    private(set) var locks: [String] = []

    func me() async throws(APIError) -> Me { try meAnswer() }
    func raids(guild: Int64) async throws(APIError) -> GuildRaids { try raidsAnswer() }
    func raid(guild: Int64, raid: Int64) async throws(APIError) -> RaidView { try raidAnswer() }

    func signup(guild: Int64, raid: Int64, characterKey: String, _ request: SignupRequest) async throws(APIError) -> SignupAnswer {
        signups.append((characterKey, request))
        return try signupAnswer(request)
    }

    func lock(guild: Int64, raid: Int64) async throws(APIError) -> RaidView { locks.append("lock"); return try lockAnswer() }
    func unlock(guild: Int64, raid: Int64) async throws(APIError) -> RaidView {
        locks.append("unlock")
        return try RaidFixture.api(RaidView.self, "guild.raid.officer")
    }
    func night(raid: Int64) async throws(APIError) -> RaidNightSummary { try nightAnswer() }
}

@MainActor
final class RecordingHaptics: RaidHaptics {
    private(set) var calls: [String] = []
    func selection() { calls.append("selection") }
    func success() { calls.append("success") }
    func error() { calls.append("error") }
}

struct FakePermission: NotificationPermission {
    let undetermined: Bool
    func isUndetermined() async -> Bool { undetermined }
    func request() async {}
}

final class MemoryAskStore: ReminderAskStore, @unchecked Sendable {
    var shown = 0
    var last: Date?
    func policy() -> ReminderAskPolicy { ReminderAskPolicy(shown: shown, lastShown: last) }
    func recordShown(at date: Date) { shown += 1; last = date }
}

/// An in-memory cache, for the offline states.
final class MemoryRaidCache: RaidCache, @unchecked Sendable {
    private var store: [String: (Data, Date)] = [:]

    func read<Value: Decodable>(_ type: Value.Type, key: String) -> Cached<Value>? {
        guard let (data, date) = store[key], let value = try? WireFormat.decoder().decode(Value.self, from: data) else { return nil }
        return Cached(value: value, savedAt: date)
    }

    func write(_ value: some Encodable, key: String, at date: Date) {
        guard let data = try? WireFormat.encoder().encode(value) else { return }
        store[key] = (data, date)
    }
}

@MainActor
func makeController(
    _ raid: RaidView? = nil,
    service: FakeRaidService = FakeRaidService(),
    haptics: RecordingHaptics = RecordingHaptics(),
    undetermined: Bool = false,
    askStore: MemoryAskStore = MemoryAskStore(),
    now: Date = RaidTime.now
) throws -> SignupController {
    try SignupController(guild: 2, raid: raid ?? RaidFixture.raid(), dependencies: .init(
        service: service, haptics: haptics, permission: FakePermission(undetermined: undetermined),
        askStore: askStore, now: { now }, hold: nil
    ))
}
