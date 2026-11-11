import Auth
import Models
import SwiftUI

/// The profile menu and the screens its rows push (flows spec 2 and 3).
extension RootView {
    /// The drawer's content for the account as it stands.
    var menu: some View {
        let session = environment.session
        let me = signInAvailable ? session.phase.me : nil
        let summary = ProfileMenuSummary(
            me: me, current: me == nil ? nil : session.currentCharacter, facts: menuFacts(me),
            version: AccountCopy.version(environment.identity, data: environment.appConfig.current?.activeBuild), now: Date()
        )
        return ProfileMenuView(
            summary: summary, signingIn: session.phase == .signingIn,
            actions: ProfileMenuActions(
                pick: { character in
                    session.pick(character.key)
                    shell.closeMenu()
                },
                refresh: {
                    #if DEBUG
                    await MenuDebugLaunch.holdRefresh()
                    #endif
                    return await session.refresh()
                },
                open: { shell.open($0) },
                signIn: { Task { await session.signIn(deviceName: UIDevice.current.name) } },
                help: { openURL(AccountLinks.help) },
                privacy: { openURL(AccountLinks.privacy) },
                loadDevices: { [client = environment.client] in try? await client.send(.listDevices) },
                loadNotifications: { [client = environment.client] in
                    let settings = try? await LiveNotificationSettingsService(client: client).load()
                    #if DEBUG
                        return MenuDebugLaunch.notifications(settings)
                    #else
                        return settings
                    #endif
                }
            )
        )
        .environment(\.portraitSource, RemotePortraitSource.shared)
    }

    /// The Account rows' second lines from what the app already knows.
    private func menuFacts(_ me: Me?) -> ProfileMenuSummary.AccountFacts {
        guard let me else { return .init() }
        return .init(
            notifications: nil,
            plan: PlanSummary(me: me, now: Date()).row,
            devices: nil,
            offlineBytes: nil
        )
    }

    /// A menu row's screen, pushed under the tab the menu was opened on.
    @ViewBuilder
    func menuScreen(_ destination: MenuDestination) -> some View {
        let session = signInAvailable ? environment.session : nil
        switch destination {
        case .guides, .tiers, .patchNotes:
            GuidesScreen(
                dependencies: .live(environment, session: session), session: session,
                tiersOn: environment.isAvailable(.guideTiers), initialSegment: Self.segment(destination),
                router: guidesRouter,
                openBis: { gear in
                    shell.selection = .gear
                    bisRouter.openRaid(spec: gear.spec, band: gear.band)
                }
            )
        case .sims:
            if let deps = HomeDependencies.live(environment).sims {
                SimsScreen(dependencies: deps, entitled: session?.phase.me?.entitlements.serverSims ?? false, router: simsRouter)
            }
        case .reports:
            if let deps = HomeDependencies.live(environment).reports {
                ReportsScreen(
                    dependencies: deps, account: ReportsAccount.of(session?.phase.me, current: session?.currentCharacter),
                    router: reportsRouter
                )
            }
        case .rankings:
            if let deps = HomeDependencies.live(environment).rankings {
                let me = session?.phase.me
                RankingsScreen(
                    dependencies: deps, account: RankingsAccount.of(me, current: session?.currentCharacter),
                    reportsAccount: ReportsAccount.of(me, current: session?.currentCharacter), encounter: nil
                )
            }
        case .planner:
            YourBuildsScreen(dependencies: plannerRouter.dependencies, router: plannerRouter, newBuild: plannerNewBuild())
        case .notifications:
            NotificationSettingsScreen(service: LiveNotificationSettingsService(client: environment.client))
        case .plan, .devices, .offlineData:
            if let me = session?.phase.me {
                AccountRouteScreen(
                    route: Self.accountRoute(destination), me: me,
                    summary: YouSummary(me: me, current: session?.currentCharacter, deviceCount: nil, version: menuVersion),
                    session: environment.session, plansLink: environment.plansLink, offlineData: offlineData
                )
            }
        case .settings:
            SettingsScreen(
                session: environment.session, version: menuVersion, plansLink: environment.plansLink,
                liveSurfaces: environment.isAvailable(.widgets),
                onLiveSettingsChange: { Task { await environment.liveActivities.syncRaid(WidgetSnapshotStore.shared().load()) } },
                calendar: guildExtras.map { ($0, calendarSync) },
                offlineData: offlineData
            )
        }
    }

    private var menuVersion: String {
        AccountCopy.version(environment.identity, data: environment.appConfig.current?.activeBuild)
    }

    static func segment(_ destination: MenuDestination) -> GuidesSegment {
        switch destination {
        case .tiers: .tiers
        case .patchNotes: .changes
        default: .guides
        }
    }

    static func accountRoute(_ destination: MenuDestination) -> YouRoute {
        switch destination {
        case .devices: .devices
        case .offlineData: .offlineData
        default: .plan
        }
    }
}
