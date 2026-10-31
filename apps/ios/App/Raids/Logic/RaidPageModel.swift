import Foundation
import Models
import Networking
import Observation

/// What a deep link asks the raid page to do on open (links.json: `spec=ask`, `state=in`).
struct RaidOpenAction: Hashable, Sendable {
    let state: SignupState
}

/// Loads one raid for the raid page (cache first), and keeps it live in its last 24 h.
@MainActor
@Observable
final class RaidPageModel {
    enum State {
        case loading
        case loaded(SignupController)
        case failed
        case notAvailable
    }

    let guild: Int64
    let raidID: Int64
    private(set) var state: State = .loading
    private(set) var offlineSince: Date?

    private let deps: SignupController.Dependencies
    private let cache: any RaidCache
    private let onChange: (RaidView) -> Void
    private var openAction: RaidOpenAction?

    /// While the raid starts within this window the page refetches every 30 s (spec F4.2 Live).
    static let liveWindow: TimeInterval = 24 * 3600
    static let livePeriod: Duration = .seconds(30)

    init(
        guild: Int64, raid: Int64, deps: SignupController.Dependencies, cache: any RaidCache,
        openAction: RaidOpenAction? = nil, onChange: @escaping (RaidView) -> Void = { _ in }
    ) {
        self.guild = guild
        self.raidID = raid
        self.deps = deps
        self.cache = cache
        self.openAction = openAction
        self.onChange = onChange
    }

    /// A page built straight from a view (previews and snapshot tests).
    init(controller: SignupController, deps: SignupController.Dependencies, offlineSince: Date? = nil) {
        guild = controller.guild
        raidID = controller.raid.id
        self.deps = deps
        cache = NoRaidCache()
        onChange = { _ in }
        state = .loaded(controller)
        self.offlineSince = offlineSince
        controller.offline = offlineSince != nil
    }

    private var cacheKey: String { "guild-\(guild)-raid-\(raidID)" }

    func load() async {
        if case .loaded = state { await refresh(); return }
        if let cached = cache.read(RaidView.self, key: cacheKey) { show(cached.value) }
        do {
            let view = try await deps.service.raid(guild: guild, raid: raidID)
            cache.write(view, key: cacheKey, at: deps.now())
            offlineSince = nil
            show(view)
            await runOpenAction()
        } catch {
            fail(error)
        }
    }

    private func refresh() async {
        guard case let .loaded(controller) = state else { return }
        await controller.refresh()
        if controller.offline {
            offlineSince = offlineSince ?? cache.read(RaidView.self, key: cacheKey)?.savedAt ?? deps.now()
        } else {
            offlineSince = nil
            cache.write(controller.raid, key: cacheKey, at: deps.now())
        }
    }

    private func show(_ view: RaidView) {
        if case let .loaded(controller) = state {
            controller.replace(with: view)
            return
        }
        let controller = SignupController(guild: guild, raid: view, dependencies: deps)
        controller.onRaidChanged = { [weak self] view in
            guard let self else { return }
            self.cache.write(view, key: self.cacheKey, at: self.deps.now())
            self.onChange(view)
        }
        state = .loaded(controller)
    }

    private func fail(_ error: APIError) {
        if case let .loaded(controller) = state {
            if SignupRefusal.isOffline(error) {
                controller.offline = true
                offlineSince = cache.read(RaidView.self, key: cacheKey)?.savedAt ?? deps.now()
            }
            return
        }
        if case let .server(status, _, _) = error, status == 404 {
            state = .notAvailable
        } else {
            state = .failed
        }
    }

    /// The link's press: In for the current character, which asks the spec when it has none.
    private func runOpenAction() async {
        guard let action = openAction, case let .loaded(controller) = state,
              let character = controller.raid.viewer.characters.first(where: \.verified)
        else { return }
        openAction = nil
        await controller.press(character.characterKey, action.state)
    }

    /// Refetches every 30 s while the raid starts within 24 h; the caller's task cancels it.
    func stayLive() async {
        while !Task.isCancelled {
            try? await Task.sleep(for: Self.livePeriod)
            guard case let .loaded(controller) = state else { continue }
            let untilStart = controller.raid.startsAt.timeIntervalSince(deps.now())
            if untilStart < Self.liveWindow, controller.facts.isSignable { await refresh() }
        }
    }
}
