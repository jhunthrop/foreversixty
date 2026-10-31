import Auth
import DesignSystem
import Models
import Networking
import Observation
import Routing
import SwiftUI

/// The raids feature's dependencies, built once from the app's environment.
@MainActor
struct RaidsDependencies {
    let service: any RaidService
    let cache: any RaidCache
    let haptics: any RaidHaptics
    let permission: any NotificationPermission
    let askStore: any ReminderAskStore
    let now: @Sendable () -> Date
    let clock: RaidClock
    /// The active data build, for the game icons.
    let iconBuild: @Sendable () async -> String?

    static func live(_ environment: AppEnvironment) -> RaidsDependencies {
        // The active build names every icon URL. In fixture mode it comes from the bundled
        // cdn.active-build fixture, as Android's fixture engine answers it; live, from the site's
        // /data/active-build.json (published by web/scripts/sync-data.mjs).
        let staticTransport: any HTTPTransport = switch environment.configuration.mode {
        case .fixture: FixtureTransport.bundled() ?? URLSessionTransport()
        case .live: URLSessionTransport()
        }
        let statics = StaticDataClient(transport: staticTransport)
        var service: any RaidService = LiveRaidService(client: environment.client)
        #if DEBUG
        if let fixture = RaidDebugLaunch.service(baseURL: environment.configuration.baseURL) { service = fixture }
        #endif
        return RaidsDependencies(
            service: service,
            cache: DiskRaidCache.standard(),
            haptics: SystemRaidHaptics(),
            permission: SystemNotificationPermission(),
            askStore: UserDefaultsReminderAskStore(),
            now: { Date() },
            clock: RaidClock(),
            iconBuild: { try? await statics.activeBuild().build }
        )
    }

    var signup: SignupController.Dependencies {
        SignupController.Dependencies(
            service: service, haptics: haptics, permission: permission, askStore: askStore, now: now, hold: .seconds(4)
        )
    }
}

/// A screen pushed on the Raids stack.
enum RaidsRoute: Hashable {
    case raid(guild: Int64, raid: Int64, roster: Bool, open: RaidOpenAction?)
    case night(raid: Int64)
}

/// The Raids tab's navigation: deep links and notifications push here (spec 5.0, 6.4).
@MainActor
@Observable
final class RaidsRouter {
    var path: [RaidsRoute] = []

    /// Opens a raid link on a fresh stack, so back goes to the tab's root.
    func open(_ link: DeepLink) {
        switch link {
        case let .raid(guild, raid, spec, state):
            let press = spec == "ask" ? SignupState.signedIn : state.flatMap(SignupState.init)
            path = [.raid(guild: guild, raid: raid, roster: false, open: press.map(RaidOpenAction.init))]
        case let .raidNight(guild, raid):
            path = [.raid(guild: guild, raid: raid, roster: false, open: nil), .night(raid: raid)]
        case .guild:
            path = []
        default:
            break
        }
    }

    /// True for the links the Raids tab opens.
    static func handles(_ link: DeepLink) -> Bool {
        switch link {
        case .raid, .raidNight, .guild: true
        default: false
        }
    }
}

/// F4: the Raids tab root. Signed in (lane F1), it follows the session's current character's
/// guild under the same selector as Home and BiS; otherwise `/v1/me`'s main character.
struct RaidsTab: View {
    let dependencies: RaidsDependencies
    @Bindable var router: RaidsRouter
    /// Lane F1's account, nil while its feature switch is off.
    let session: SessionStore?
    @State private var model: RaidsTabModel
    @State private var nextController: SignupController?
    @State private var iconBuild: String?
    @Environment(\.openURL) private var openURL

    init(dependencies: RaidsDependencies, router: RaidsRouter, session: SessionStore? = nil) {
        self.dependencies = dependencies
        self.router = router
        self.session = session
        _model = State(initialValue: RaidsTabModel(service: dependencies.service, cache: dependencies.cache, now: dependencies.now))
    }

    var body: some View {
        NavigationStack(path: $router.path) {
            RaidsTabContentView(
                state: model.state, clock: dependencies.clock, now: dependencies.now, nextController: nextController,
                openRaid: { raid, roster in router.path.append(.raid(guild: guildID ?? 0, raid: raid.id, roster: roster, open: nil)) },
                openNight: { router.path.append(.night(raid: $0)) },
                openWeb: { openURL($0) },
                retry: { Task { await reload() } }
            )
            .refreshable { await reload() }
            .navigationTitle(Copy.tabRaids)
            .toolbarTitleDisplayMode(.inline)
            .characterSelector(isOn: session != nil)
            .searchButton()
            // Under the selector the bar matches Home and BiS: no title beside it (the title still
            // names the back button on the pushed raid page).
            .toolbar(removing: session != nil ? .title : nil)
            .navigationDestination(for: RaidsRoute.self) { destination($0) }
        }
        .environment(\.raidIconBuild, iconBuild)
        .task { iconBuild = await dependencies.iconBuild() }
        .task(id: session?.currentCharacter?.key) {
            model.currentKey = session?.currentCharacter?.key
            await reload()
        }
    }

    private var guildID: Int64? {
        if case let .loaded(content) = model.state { return content.head.id }
        return nil
    }

    private var content: RaidsTabContent? {
        if case let .loaded(content) = model.state { return content }
        return nil
    }

    private func reload() async {
        await model.load()
        syncNextController()
    }

    /// The card's controller follows the next raid; a sign-up on it updates the tab's list.
    private func syncNextController() {
        guard let content, let next = PlannedRaids(content.planned, now: dependencies.now()).next else {
            nextController = nil
            return
        }
        if let controller = nextController, controller.raid.id == next.id {
            controller.replace(with: next)
            controller.offline = content.offlineSince != nil
            return
        }
        let controller = SignupController(guild: content.head.id, raid: next, dependencies: dependencies.signup)
        controller.offline = content.offlineSince != nil
        controller.onRaidChanged = { [model] view in model.replace(view) }
        nextController = controller
    }

    @ViewBuilder
    private func destination(_ route: RaidsRoute) -> some View {
        switch route {
        case let .raid(guild, raid, roster, open):
            RaidPageScreen(
                model: RaidPageModel(
                    guild: guild, raid: raid, deps: dependencies.signup, cache: dependencies.cache, openAction: open,
                    onChange: { [model] view in model.replace(view) }
                ),
                clock: dependencies.clock, now: dependencies.now, currentKey: content?.character.key,
                guildName: content?.head.name ?? "", guildURL: content?.head.webURL,
                openNight: { router.path.append(.night(raid: $0)) },
                back: { router.path.removeAll() },
                scrollToRoster: roster
            )
        case let .night(raid):
            RaidNightView(
                model: RaidNightModel(raid: raid, service: dependencies.service, cache: dependencies.cache, now: dependencies.now),
                clock: dependencies.clock, now: dependencies.now
            )
        }
    }
}
