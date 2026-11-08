import Auth
import DesignSystem
import Models
import SwiftUI
import UIKit

/// Home's account states around the signed-in home (spec F1.1 to F1.4 and F1.3's no-characters
/// state). Signed in with characters, it draws `signedIn`, the home lane's screen.
struct HomeAccountGate<SignedIn: View>: View {
    let session: SessionStore
    let onPickClass: () -> Void
    @ViewBuilder let signedIn: () -> SignedIn
    @Environment(\.openURL) private var openURL

    var body: some View {
        switch session.phase {
        case .signedOut, .signingIn:
            SignedOutPanel(session: session, variant: .home, onPickClass: onPickClass)
        case .importing:
            ImportView()
        case let .signedIn(me):
            if let me, me.characters.isEmpty {
                NoCharactersView { openURL(AccountLinks.setup) }
                    .refreshable { await session.refresh() }
            } else {
                signedIn()
            }
        }
    }
}

/// F1.1 bound to the session: the sign-in tap, the spinner and the notice.
struct SignedOutPanel: View {
    let session: SessionStore
    var variant: SignedOutView.Variant = .home
    let onPickClass: () -> Void

    var body: some View {
        SignedOutView(
            variant: variant,
            notice: notice,
            noticeIsFailure: noticeIsFailure,
            isSigningIn: session.phase == .signingIn,
            onSignIn: {
                session.dismissNotice()
                Task { await session.signIn(deviceName: UIDevice.current.name) }
            },
            onPickClass: {
                session.dismissNotice()
                onPickClass()
            }
        )
    }

    private var signedOutNotice: SignedOutNotice? {
        if case let .signedOut(notice?) = session.phase { return notice }
        return nil
    }

    private var notice: String? { signedOutNotice.map(AccountCopy.notice) }

    private var noticeIsFailure: Bool { signedOutNotice.map(AccountCopy.isFailure) ?? true }
}

/// The menu's Settings (flows spec 3, F10): the account rows the You tab had (characters, Discord,
/// plan, devices, live surfaces, calendar, offline data, help, about, sign out, delete account),
/// pushed under the tab the menu was opened on. Signed out the menu has no Settings row.
struct SettingsScreen: View {
    let session: SessionStore
    /// You's version line and About's version (`AccountCopy.version`).
    let version: String
    /// Lane F8: See plans and Manage billing, where the store allows them.
    let plansLink: PlansLink
    /// Lane F14: whether the widgets switch is on, so Settings lists Live surfaces.
    var liveSurfaces = false
    /// Lane F14: the Live Activities redraw when a switch changes.
    var onLiveSettingsChange: () -> Void = {}
    /// Lane F16: Calendar, nil while `guild_extras` is off.
    var calendar: (extras: GuildExtrasDependencies, sync: CalendarSyncController)?
    /// Lane F15: Offline data's model, nil while `system` is off.
    var offlineData: (@MainActor () -> OfflineDataModel)?
    @State private var deviceCount: Int?

    var body: some View {
        if let me = session.phase.me {
            let summary = YouSummary(me: me, current: session.currentCharacter, deviceCount: deviceCount, version: version)
            YouView(
                summary: summary, mayLink: plansLink.mayLink, showsLiveSurfaces: liveSurfaces, showsCalendar: calendar != nil,
                showsOfflineData: offlineData != nil
            ) {
                Task { await session.signOut() }
            }
            .navigationDestination(for: YouRoute.self) { route in
                AccountRouteScreen(
                    route: route, me: me, summary: summary, session: session, plansLink: plansLink,
                    onLiveSettingsChange: onLiveSettingsChange, calendar: calendar, offlineData: offlineData
                ) { deviceCount = $0 }
            }
            .task { deviceCount = (try? await session.devices())?.count }
            .task { await plansLink.resolve() }
        } else {
            ProgressView().frame(maxWidth: .infinity, maxHeight: .infinity).background(FSColor.bg)
        }
    }
}

/// One account screen, by route: Settings' rows and the menu's Plan, Devices and Offline data.
struct AccountRouteScreen: View {
    let route: YouRoute
    let me: Me
    let summary: YouSummary
    let session: SessionStore
    let plansLink: PlansLink
    var onLiveSettingsChange: () -> Void = {}
    var calendar: (extras: GuildExtrasDependencies, sync: CalendarSyncController)?
    var offlineData: (@MainActor () -> OfflineDataModel)?
    var onDevices: (Int) -> Void = { _ in }

    var body: some View {
        switch route {
        case .plan: PlanView(plan: summary.plan, mayLink: plansLink.mayLink)
        case .devices: DevicesScreen(session: session, onCountChange: onDevices)
        case .deleteAccount: DeleteAccountScreen(session: session)
        case .about: AboutView(version: summary.version)
        case .liveSurfaces:
            LiveSurfacesScreen(showsSim: me.entitlements.serverSims, onChange: onLiveSettingsChange)
        case .calendar:
            if let calendar {
                CalendarRoute(extras: calendar.extras, sync: calendar.sync, account: String(me.user.id))
            }
        case .offlineData:
            if let offlineData {
                OfflineDataScreen(model: offlineData(), now: { Date() })
            }
        }
    }
}
