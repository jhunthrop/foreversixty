@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// A fake API for the widget calls: answers from fixtures, or fails.
final class FakeWidgetCalls: WidgetCalls, @unchecked Sendable {
    private let lock = NSLock()
    var answerError: APIError?
    var snapshotError: APIError?
    private(set) var answered: [String] = []
    var fresh: Widget?

    func snapshot(character: String?) async throws(APIError) -> Widget {
        if let snapshotError { throw snapshotError }
        if let fresh { return fresh }
        do { return try WidgetFixtures.widget("widget.get") } catch { throw .missingFixture("\(error)") }
    }

    func answer(guild: Int64, raid: Int64, characterKey: String, state: String) async throws(APIError) {
        lock.withLock { answered.append(state) }
        if let answerError { throw answerError }
    }
}

@Suite("Widget logic")
struct WidgetLogicTests {
    private func tempStore() -> WidgetSnapshotStore {
        WidgetSnapshotStore(directory: FileManager.default.temporaryDirectory.appending(path: "widgets-\(UUID().uuidString)"))
    }

    @Test("the snapshot round-trips through the shared file, and sign-out empties it")
    func storeRoundTrip() throws {
        let store = tempStore()
        let snapshot = try WidgetFixtures.snapshot()
        try store.save(snapshot)
        #expect(store.load()?.widget == snapshot.widget)
        #expect(store.snapshot(for: "us/pvp/lightningbolltz")?.widget == snapshot.widget)
        #expect(store.snapshot(for: "us/pvp/someone") == nil)
        store.saveCharacters([WidgetCharacterChoice(key: "us/pvp/a", name: "A")])
        #expect(store.characters().map(\.name) == ["A"])
        store.clear()
        #expect(store.load() == nil && store.characters().isEmpty)
    }

    @Test("a press saves: drawn at once, then the server's snapshot")
    func pressSaves() async throws {
        let store = tempStore()
        try store.save(try WidgetFixtures.snapshot(patch: ["my": .object(["state": .null, "spec": .string("Elemental"), "role": .string("dps")])]))
        let calls = FakeWidgetCalls()
        calls.fresh = try WidgetFixtures.widget("widget.get", patch: ["signed_up": .number(25)])
        let signup = WidgetSignup(store: store, calls: calls, now: { WidgetTime.now }, reload: {})
        #expect(await signup.answer("in"))
        #expect(calls.answered == ["in"])
        let saved = try #require(store.load())
        #expect(saved.pending == nil && saved.widget.nextRaid?.signedUp == 25 && saved.widget.nextRaid?.my.state == "in")
    }

    @Test("a failed press redraws the previous answer with Not saved for ten seconds")
    func pressFails() async throws {
        let store = tempStore()
        try store.save(try WidgetFixtures.snapshot())
        let calls = FakeWidgetCalls()
        calls.answerError = .transport("offline")
        let signup = WidgetSignup(store: store, calls: calls, now: { WidgetTime.now }, reload: {})
        #expect(await signup.answer("out") == false)
        let saved = try #require(store.load())
        #expect(saved.pending == nil && saved.notSavedState == "out")
        #expect(saved.notSavedUntil == WidgetTime.now.addingTimeInterval(WidgetSnapshot.notSavedFor))
        let figures = WidgetFigures(saved, now: WidgetTime.now, zone: WidgetTime.chicago)
        #expect(figures.checked == "in" && figures.notSaved == "out")
    }

    @Test("a saved press stands when the fresh snapshot cannot be read")
    func pressSavedOffline() async throws {
        let store = tempStore()
        try store.save(try WidgetFixtures.snapshot())
        let calls = FakeWidgetCalls()
        calls.snapshotError = .transport("offline")
        #expect(await WidgetSignup(store: store, calls: calls, now: { WidgetTime.now }, reload: {}).answer("late"))
        #expect(store.load()?.widget.nextRaid?.my.state == "late")
    }

    @Test("timeline entries step 30 minutes, hourly past two days, and stop at the horizon")
    func timeline() throws {
        let snapshot = try WidgetFixtures.snapshot()
        let raid = try #require(snapshot.widget.nextRaid)
        let dates = WidgetTimeline.dates(now: WidgetTime.now, snapshot: snapshot)
        #expect(dates.first == WidgetTime.now)
        #expect(dates.contains(raid.startsAt) && dates.contains(raid.endsAt))
        #expect(dates.contains(raid.startsAt.addingTimeInterval(-30 * 60)))
        #expect(dates.contains(raid.startsAt.addingTimeInterval(-49 * 3600)))
        #expect(!dates.contains(raid.startsAt.addingTimeInterval(-(48 * 3600 + 30 * 60))))
        #expect(dates == dates.sorted())
        #expect(dates.allSatisfy { $0 <= WidgetTime.now.addingTimeInterval(WidgetTimeline.horizon) })
    }

    @Test("the silent widget.refresh push is recognised from its fixture")
    func refreshPush() throws {
        let data = try RaidFixture.data("surface.widget-refresh.apns")
        let userInfo = try #require(JSONSerialization.jsonObject(with: data) as? [AnyHashable: Any])
        #expect(WidgetRefresher.isRefreshPush(userInfo))
        #expect(!WidgetRefresher.isRefreshPush(["kind": "raid.posted"]))
    }

    @Test("the Live Activity shapes decode the API's push fixtures")
    func liveShapes() throws {
        let raid = try #require(JSONSerialization.jsonObject(with: RaidFixture.data("surface.live-raid-start.apns")) as? [String: Any])
        let aps = try #require(raid["aps"] as? [String: Any])
        let attributes = try JSONDecoder().decode(RaidActivityAttributes.self, from: JSONSerialization.data(withJSONObject: aps["attributes"] ?? [:]))
        let state = try JSONDecoder().decode(RaidActivityAttributes.ContentState.self, from: JSONSerialization.data(withJSONObject: aps["content-state"] ?? [:]))
        #expect(aps["attributes-type"] as? String == "RaidActivityAttributes")
        #expect(attributes.zoneName == "Onyxia's Lair" && attributes.characterName == "Lightningbolltz")
        #expect(state.signedUp == 38 && state.size == 40 && state.myState == "in" && state.inviteAt == nil)
        let rawStart = (aps["content-state"] as? [String: Any])?["starts_at"] as? Int ?? 0
        #expect(state.startDate == Date(timeIntervalSince1970: TimeInterval(rawStart)))

        let sim = try #require(JSONSerialization.jsonObject(with: RaidFixture.data("surface.live-sim-update.apns")) as? [String: Any])
        let simAps = try #require(sim["aps"] as? [String: Any])
        let simState = try JSONDecoder().decode(SimActivityAttributes.ContentState.self, from: JSONSerialization.data(withJSONObject: simAps["content-state"] ?? [:]))
        #expect(simState.state == "running" && simState.done == 40 && simState.total == 64 && simState.etaSec == 120)
        let encoded = try JSONSerialization.jsonObject(with: JSONEncoder().encode(simState)) as? [String: Any]
        #expect(Set(encoded?.keys.map { $0 } ?? []) == ["state", "done", "total", "eta_sec"])
    }

    @Test("the sim activity's words: running, finished, failed")
    func simWords() {
        let running = SimLiveFigures(title: "Elemental · Top Gear", state: .init(state: "running", done: 40, total: 64, etaSec: 120, dps: nil))
        #expect(running.detail == "40 of 64 combinations" && running.left == "about 2 min left" && running.fraction == 40.0 / 64)
        let finished = SimLiveFigures(title: "t", state: .init(state: "finished", done: 64, total: 64, etaSec: nil, dps: 412))
        #expect(finished.status == Copy.liveFinished && finished.value == "412" && finished.valueRest == "DPS · tap for the result")
        let failed = SimLiveFigures(title: "t", state: .init(state: "failed", done: 0, total: nil, etaSec: nil, dps: nil))
        #expect(failed.status == Copy.liveFailed && failed.detail == Copy.liveOpenRun)
        let plain = SimLiveRun(id: "x", title: "t", progress: SimProgress(state: "running", iterationsDone: 10, stage: 1, combosDone: 0, combosTotal: 1),
                               estimateSeconds: 300, elapsedSeconds: 60, dps: nil)
        #expect(plain.state.total == nil && plain.state.etaSec == 240)
    }

    @Test("the Live surfaces switches are on by default")
    func settings() {
        let suite = "widgets-\(UUID().uuidString)"
        let defaults = UserDefaults(suiteName: suite) ?? .standard
        let settings = LiveSurfaceSettings(suiteName: suite)
        #expect(settings.raidCountdown && settings.simProgress)
        settings.set(raidCountdown: false)
        #expect(!settings.raidCountdown && settings.simProgress)
        defaults.removePersistentDomain(forName: suite)
    }
}
