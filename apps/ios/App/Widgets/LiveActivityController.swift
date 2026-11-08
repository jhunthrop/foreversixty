import ActivityKit
import Foundation
import Models
import Networking

/// The server side of a Live Activity: its push token kept by `POST /v1/me/live-activities` and
/// dropped by `DELETE` when the activity ends, and the install's push-to-start token kept by
/// `POST /v1/me/live-activities/start-token`, so `raid.starting` can start the raid's activity.
protocol LiveRegistrations: Sendable {
    func register(kind: String, objectID: String, token: Data) async -> Int64?
    func unregister(_ id: Int64) async
    @discardableResult
    func storeStartToken(kind: String, token: Data) async -> Bool
}

struct LiveRegistrationCalls: LiveRegistrations {
    let client: APIClient

    func register(kind: String, objectID: String, token: Data) async -> Int64? {
        let request = LiveActivityRequest(kind: kind, platform: "ios", pushToken: token.hexString, objectId: objectID)
        guard let endpoint = try? Endpoint<LiveActivityRegistered>.registerLiveActivity(request) else { return nil }
        return try? await client.send(endpoint).id
    }

    func unregister(_ id: Int64) async {
        _ = try? await client.send(.deleteLiveActivity(id: id))
    }

    func storeStartToken(kind: String, token: Data) async -> Bool {
        let request = LiveStartTokenRequest(platform: "ios", kind: kind, token: token.hexString)
        guard let endpoint = try? Endpoint<LiveStartTokenStored>.storeLiveStartToken(request) else { return false }
        return (try? await client.send(endpoint)) != nil
    }
}

extension Data {
    /// An APNs token as the API takes it: lowercase hex.
    var hexString: String { map { String(format: "%02x", $0) }.joined() }
}

/// Starts, updates and ends the raid countdown and the server-sim Live Activities (spec §5.6).
/// The server starts the raid's activity by push at `raid.starting` and ends it on the night's
/// summary; the app starts it too while it runs inside the window, and every update token it sees
/// goes to the server.
@MainActor
final class LiveActivityController {
    private let registrations: any LiveRegistrations
    private let settings: LiveSurfaceSettings
    private let now: @Sendable () -> Date
    private let notificationsAllowed: @Sendable () async -> Bool
    private var registered: [String: Int64] = [:]
    /// The activities whose update tokens are being followed (`raid-7`, `sim-…`).
    private var following: Set<String> = []
    private var lastSimUpdate: [String: Date] = [:]
    /// The push-to-start and remote-activity follows, while signed in with notifications allowed.
    private var remoteFollow: Task<Void, Never>?

    /// The sim activity updates at most this often (spec §5.6).
    static let simUpdateEvery: TimeInterval = 15
    static let cancelledFor: TimeInterval = 5 * 60
    static let finishedFor: TimeInterval = 60 * 60

    init(
        registrations: any LiveRegistrations, settings: LiveSurfaceSettings = .shared,
        now: @escaping @Sendable () -> Date = { Date() },
        notificationsAllowed: @escaping @Sendable () async -> Bool = PushRegistrar.systemAuthorized
    ) {
        self.registrations = registrations
        self.settings = settings
        self.now = now
        self.notificationsAllowed = notificationsAllowed
    }

    // MARK: Remote start (flows F2)

    /// While signed in with notifications allowed: the install's push-to-start tokens go to the API,
    /// and each activity the server starts has its update tokens registered like one the app started.
    /// Signed out (or not allowed) stops following.
    func followRemoteStarts(signedIn: Bool) async {
        remoteFollow?.cancel()
        remoteFollow = nil
        guard signedIn, settings.raidCountdown, await notificationsAllowed() else { return }
        let registrations = registrations
        let startTokens = RaidActivities.startTokens()
        let started = RaidActivities.started()
        remoteFollow = Task { [weak self] in
            async let stored: Void = {
                for await token in startTokens { await registrations.storeStartToken(kind: "raid", token: token) }
            }()
            for await (raidID, tokens) in started {
                self?.follow(tokens, kind: "raid", objectID: raidID)
            }
            await stored
        }
    }

    // MARK: Raid

    /// Brings the raid countdown in line with the snapshot: started inside the window for an In
    /// or Late player, updated while it runs, `Cancelled` for 5 minutes, ended otherwise.
    func syncRaid(_ snapshot: WidgetSnapshot?) async {
        let running = RaidActivities.runningIDs()
        guard let snapshot, let raid = snapshot.widget.nextRaid, let character = snapshot.widget.character else {
            for id in running { await endRaid(id, state: nil, after: nil) }
            return
        }
        let id = String(raid.raidId)
        let figures = WidgetFigures(snapshot, now: now(), zone: .current)
        let state = RaidActivityAttributes.ContentState(raid, myState: figures.checked)
        let stale = raid.startsAt.addingTimeInterval(WidgetFigures.liveTail)
        for other in running where other != id { await endRaid(other, state: nil, after: nil) }
        let isRunning = running.contains(id)
        if raid.status == "cancelled", isRunning {
            await endRaid(id, state: state, after: now().addingTimeInterval(Self.cancelledFor))
        } else if figures.liveActive, settings.raidCountdown {
            if isRunning {
                await RaidActivities.update(id, state: state, stale: stale)
            } else if let tokens = RaidActivities.start(RaidActivityAttributes(raid, characterName: character.name), state: state, stale: stale) {
                follow(tokens, kind: "raid", objectID: id)
            }
        } else if isRunning {
            await endRaid(id, state: nil, after: nil)
        }
    }

    // MARK: Sim

    /// Follows the in-flight server runs the sims poll knows: one activity each, updated at most
    /// every 15 seconds, `Finished` (an hour) or `Failed` at the end. Only for `server_sims`.
    func syncSims(_ runs: [SimLiveRun], entitled: Bool) async {
        guard entitled, settings.simProgress else { return }
        let running = SimActivities.runningIDs()
        for run in runs {
            let stale = now().addingTimeInterval(Self.simUpdateEvery * 4)
            let isRunning = run.state.state == SimActivityAttributes.ContentState.running
            if running.contains(run.id) {
                if isRunning {
                    guard now().timeIntervalSince(lastSimUpdate[run.id] ?? .distantPast) >= Self.simUpdateEvery else { continue }
                    lastSimUpdate[run.id] = now()
                    await SimActivities.update(run.id, state: run.state, stale: stale)
                } else {
                    let after = run.state.state == SimActivityAttributes.ContentState.finished ? Self.finishedFor : Self.cancelledFor
                    await SimActivities.end(run.id, state: run.state, after: now().addingTimeInterval(after))
                    await forget("sim-\(run.id)")
                }
            } else if isRunning, let tokens = SimActivities.start(SimActivityAttributes(simID: run.id, title: run.title), state: run.state, stale: stale) {
                lastSimUpdate[run.id] = now()
                follow(tokens, kind: "sim", objectID: run.id)
            }
        }
    }

    /// Sign-out: every activity ends at once.
    func endAll() async {
        for id in RaidActivities.runningIDs() { await endRaid(id, state: nil, after: nil) }
        for id in SimActivities.runningIDs() {
            await SimActivities.end(id, state: nil, after: nil)
            await forget("sim-\(id)")
        }
    }

    // MARK: Shared

    private func endRaid(_ id: String, state: RaidActivityAttributes.ContentState?, after: Date?) async {
        await RaidActivities.end(id, state: state, after: after)
        await forget("raid-\(id)")
    }

    /// Keeps each push token the system hands the activity on the server.
    /// Keeps each push token the system hands the activity on the server, once per activity: one
    /// the app started also appears among the remote starts.
    private func follow(_ tokens: AsyncStream<Data>, kind: String, objectID: String) {
        let key = "\(kind)-\(objectID)"
        guard following.insert(key).inserted else { return }
        let registrations = registrations
        Task { [weak self] in
            for await token in tokens {
                let id = await registrations.register(kind: kind, objectID: objectID, token: token)
                self?.keep(id, for: key)
            }
            self?.following.remove(key)
        }
    }

    private func keep(_ id: Int64?, for key: String) {
        if let id { registered[key] = id }
    }

    private func forget(_ key: String) async {
        if let id = registered.removeValue(forKey: key) { await registrations.unregister(id) }
    }
}

/// ActivityKit calls for the raid countdown, by raid id. `Activity` is not Sendable, so it never
/// leaves these nonisolated functions; the controller passes ids and states only.
enum RaidActivities {
    nonisolated static func runningIDs() -> [String] {
        Activity<RaidActivityAttributes>.activities.filter { $0.activityState == .active }.map(\.attributes.raidID)
    }

    nonisolated static func start(_ attributes: RaidActivityAttributes, state: RaidActivityAttributes.ContentState, stale: Date?) -> AsyncStream<Data>? {
        guard ActivityAuthorizationInfo().areActivitiesEnabled,
              let activity = try? Activity.request(attributes: attributes, content: ActivityContent(state: state, staleDate: stale), pushType: .token)
        else { return nil }
        return tokens(activity.id)
    }

    /// The install's push-to-start tokens, as the system hands them out (they rotate).
    nonisolated static func startTokens() -> AsyncStream<Data> {
        AsyncStream { continuation in
            let task = Task {
                for await token in Activity<RaidActivityAttributes>.pushToStartTokenUpdates { continuation.yield(token) }
                continuation.finish()
            }
            continuation.onTermination = { _ in task.cancel() }
        }
    }

    /// Each raid activity as it appears (one a push started included): its raid id and update tokens.
    nonisolated static func started() -> AsyncStream<(String, AsyncStream<Data>)> {
        AsyncStream { continuation in
            let task = Task {
                for await activity in Activity<RaidActivityAttributes>.activityUpdates {
                    continuation.yield((activity.attributes.raidID, tokens(activity.id)))
                }
                continuation.finish()
            }
            continuation.onTermination = { _ in task.cancel() }
        }
    }

    nonisolated static func tokens(_ activityID: String) -> AsyncStream<Data> {
        AsyncStream { continuation in
            let task = Task {
                guard let activity = Activity<RaidActivityAttributes>.activities.first(where: { $0.id == activityID }) else {
                    continuation.finish()
                    return
                }
                for await token in activity.pushTokenUpdates { continuation.yield(token) }
                continuation.finish()
            }
            continuation.onTermination = { _ in task.cancel() }
        }
    }

    nonisolated static func update(_ id: String, state: RaidActivityAttributes.ContentState, stale: Date?) async {
        for activity in Activity<RaidActivityAttributes>.activities where activity.attributes.raidID == id {
            await activity.update(ActivityContent(state: state, staleDate: stale))
        }
    }

    nonisolated static func end(_ id: String, state: RaidActivityAttributes.ContentState?, after: Date?) async {
        for activity in Activity<RaidActivityAttributes>.activities where activity.attributes.raidID == id {
            await activity.end(state.map { ActivityContent(state: $0, staleDate: nil) }, dismissalPolicy: after.map { .after($0) } ?? .immediate)
        }
    }
}

/// ActivityKit calls for a server run's activity, by sim id.
enum SimActivities {
    nonisolated static func runningIDs() -> [String] {
        Activity<SimActivityAttributes>.activities.filter { $0.activityState == .active }.map(\.attributes.simID)
    }

    nonisolated static func start(_ attributes: SimActivityAttributes, state: SimActivityAttributes.ContentState, stale: Date?) -> AsyncStream<Data>? {
        guard ActivityAuthorizationInfo().areActivitiesEnabled,
              let activity = try? Activity.request(attributes: attributes, content: ActivityContent(state: state, staleDate: stale), pushType: .token)
        else { return nil }
        let activityID = activity.id
        return AsyncStream { continuation in
            let task = Task {
                guard let activity = Activity<SimActivityAttributes>.activities.first(where: { $0.id == activityID }) else {
                    continuation.finish()
                    return
                }
                for await token in activity.pushTokenUpdates { continuation.yield(token) }
                continuation.finish()
            }
            continuation.onTermination = { _ in task.cancel() }
        }
    }

    nonisolated static func update(_ id: String, state: SimActivityAttributes.ContentState, stale: Date?) async {
        for activity in Activity<SimActivityAttributes>.activities where activity.attributes.simID == id {
            await activity.update(ActivityContent(state: state, staleDate: stale))
        }
    }

    nonisolated static func end(_ id: String, state: SimActivityAttributes.ContentState?, after: Date?) async {
        for activity in Activity<SimActivityAttributes>.activities where activity.attributes.simID == id {
            await activity.end(state.map { ActivityContent(state: $0, staleDate: nil) }, dismissalPolicy: after.map { .after($0) } ?? .immediate)
        }
    }
}

/// One server run as the sims poll sees it, in the Live Activity's terms.
struct SimLiveRun: Equatable, Sendable {
    let id: String
    let title: String
    let state: SimActivityAttributes.ContentState

    /// From a list row and its newest poll answer (SimProgress), with the row's estimate.
    init(id: String, title: String, progress: SimProgress?, estimateSeconds: Int?, elapsedSeconds: Int, dps: Double?) {
        self.id = id
        self.title = title
        let remaining = estimateSeconds.map { max(0, $0 - elapsedSeconds) }
        switch progress?.state {
        case "done":
            state = .init(state: SimActivityAttributes.ContentState.finished, done: progress?.combosDone ?? 0,
                          total: progress?.combosTotal, etaSec: nil, dps: progress?.dps ?? dps)
        case "error":
            state = .init(state: SimActivityAttributes.ContentState.failed, done: progress?.combosDone ?? 0,
                          total: nil, etaSec: nil, dps: nil)
        default:
            let total = progress.map(\.combosTotal).flatMap { $0 > 1 ? $0 : nil }
            state = .init(state: SimActivityAttributes.ContentState.running, done: progress?.combosDone ?? 0,
                          total: total, etaSec: remaining, dps: nil)
        }
    }

    init(id: String, title: String, state: SimActivityAttributes.ContentState) {
        self.id = id
        self.title = title
        self.state = state
    }
}
