import Auth
import Foundation
import Networking

/// The app's dependencies, built once from Info.plist. Feature lanes read the
/// client and token store from here; tests and previews build their own.
struct AppEnvironment {
    let configuration: APIConfiguration
    let tokens: any TokenStore
    let client: APIClient
    /// Flips when the API says this build is below its minimum; RootView then blocks.
    let updateGate: UpdateGate
    /// Read at launch (RootView) and cached; feature lanes ask it `isEnabled(_:)`.
    let appConfig: AppConfigStore
    /// The account (lane F1): signed in or out, the current character, sign-in and sign-out.
    /// Feature lanes read it from the SwiftUI environment (`@Environment(SessionStore.self)`).
    let session: SessionStore
    /// This build's version and build number (apps/VERSION, lane P2); You's version line reads it.
    let identity: ClientIdentity
    /// Lane F8: whether See plans may link to web checkout here (decision P, fail closed).
    let plansLink: PlansLink
    /// Lane F14: the App Group snapshot every widget and Live Activity draws from.
    let widgets: WidgetRefresher
    /// Lane F14: the raid countdown and server-sim Live Activities.
    let liveActivities: LiveActivityController

    @MainActor
    init(
        configuration: APIConfiguration,
        tokens: any TokenStore,
        identity: ClientIdentity,
        updateGate: UpdateGate = UpdateGate(),
        bundle: Bundle = .main,
        configCache: any AppConfigCache = UserDefaultsAppConfigCache(),
        account: AccountStorage = .device,
        storefront: any StorefrontReading = StorefrontLaunch.reader
    ) {
        let sessionTokens = SessionTokenSource(tokens)
        self.identity = identity
        self.configuration = configuration
        self.tokens = sessionTokens
        self.updateGate = updateGate
        self.client = APIClient.make(
            configuration, tokens: sessionTokens, identity: identity, updateGate: updateGate, bundle: bundle
        )
        let appConfig = AppConfigStore(client: client, cache: configCache)
        self.appConfig = appConfig
        let mode = configuration.mode
        self.plansLink = PlansLink(appConfig: appConfig, storefront: storefront) {
            appConfig.isAvailable(.premiumLink, mode: mode, isDebugBuild: Self.isDebugBuild)
        }
        let widgets = WidgetRefresher(
            calls: LiveWidgetCalls(client: client),
            iconBuild: { appConfig.current?.activeBuild ?? BundledBuild.read() },
            enabled: { appConfig.isAvailable(.widgets, mode: mode, isDebugBuild: Self.isDebugBuild) }
        )
        let liveActivities = LiveActivityController(registrations: LiveRegistrationCalls(client: client))
        widgets.afterRefresh = { snapshot in await liveActivities.syncRaid(snapshot) }
        self.widgets = widgets
        self.liveActivities = liveActivities
        self.session = SessionStore(
            client: client,
            tokens: sessionTokens,
            web: account.web(),
            picked: account.picked,
            deviceID: account.deviceID,
            clearCaches: {
                URLCache.shared.removeAllCachedResponses()
                FileHomeCache.standard()?.removeAll()
                DiskRaidCache.standard().removeAll()
                DiskRaidCache.sims().removeAll()
                DiskRaidCache.reports().removeAll()
                widgets.clear()
                Task { await liveActivities.endAll() }
            }
        )
    }

    /// Whether a feature's screens show: on in debug and fixture builds, else as the config says.
    func isAvailable(_ feature: Feature) -> Bool {
        appConfig.isAvailable(feature, mode: configuration.mode, isDebugBuild: Self.isDebugBuild)
    }

    #if DEBUG
    private static let isDebugBuild = true
    #else
    private static let isDebugBuild = false
    #endif

    /// Fixture mode keeps its token in memory so a fixture session never reaches the Keychain.
    @MainActor
    static func fromBundle(_ bundle: Bundle = .main) -> AppEnvironment {
        let configuration: APIConfiguration
        do {
            configuration = try APIConfiguration(infoDictionary: bundle.infoDictionary ?? [:])
        } catch {
            preconditionFailure("Build is misconfigured: \(error)")
        }
        let tokens: any TokenStore = switch configuration.mode {
        // Lane F14: written to the App Group's keychain group so the widget's sign-up intent reads it.
        case .live: KeychainTokenStore(accessGroup: AppGroup.containerURL == nil ? nil : AppGroup.id)
        case .fixture: InMemoryTokenStore(token: "fsd_fixture")
        }
        let account: AccountStorage = switch configuration.mode {
        case .live: .device
        case .fixture: .fixture(bundle: bundle)
        }
        let identity = ClientIdentity.bundled(bundle, installID: InstallID.current())
        return AppEnvironment(
            configuration: configuration, tokens: tokens, identity: identity, bundle: bundle, account: account
        )
    }

    // TODO(lane A2, F7): register the APNs device token with the push_tokens
    // endpoint after sign-in and when the system rotates it; unregister on sign-out.
}
