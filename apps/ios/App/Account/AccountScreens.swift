import Auth
import DesignSystem
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

/// The You tab: the account section when signed in, the sign-in panel when not.
struct YouScreen: View {
    let session: SessionStore
    /// You's version line and About's version (`AccountCopy.version`).
    let version: String
    let onPickClass: () -> Void
    /// Lane F8: See plans and Manage billing, where the store allows them.
    let plansLink: PlansLink
    /// Lane F14: whether the widgets switch is on, so You lists Live surfaces.
    var liveSurfaces = false
    /// Lane F14: the Live Activities redraw when a switch changes.
    var onLiveSettingsChange: () -> Void = {}
    /// Lane F16: You › Calendar, nil while `guild_extras` is off.
    var calendar: (extras: GuildExtrasDependencies, sync: CalendarSyncController)?
    @State private var deviceCount: Int?

    var body: some View {
        NavigationStack {
            if let me = session.phase.me {
                let summary = YouSummary(me: me, current: session.currentCharacter, deviceCount: deviceCount, version: version)
                YouView(summary: summary, mayLink: plansLink.mayLink, showsLiveSurfaces: liveSurfaces, showsCalendar: calendar != nil) {
                    Task { await session.signOut() }
                }
                .navigationDestination(for: YouRoute.self) { route in
                    switch route {
                    case .plan: PlanView(plan: summary.plan, mayLink: plansLink.mayLink)
                    case .devices: DevicesScreen(session: session) { deviceCount = $0 }
                    case .deleteAccount: DeleteAccountScreen(session: session)
                    case .about: AboutView(version: version)
                    case .liveSurfaces:
                        LiveSurfacesScreen(showsSim: me.entitlements.serverSims, onChange: onLiveSettingsChange)
                    case .calendar:
                        if let calendar {
                            CalendarRoute(extras: calendar.extras, sync: calendar.sync, account: String(me.user.id))
                        }
                    }
                }
                .task { deviceCount = (try? await session.devices())?.count }
                .task { await plansLink.resolve() }
            } else if session.phase.isSignedIn {
                ProgressView().frame(maxWidth: .infinity, maxHeight: .infinity).background(FSColor.bg)
            } else {
                SignedOutPanel(session: session, onPickClass: onPickClass)
            }
        }
    }
}
