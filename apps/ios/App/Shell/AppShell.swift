import DesignSystem
import Models
import Observation
import SwiftUI

/// The four tabs, the app's jobs (flows spec 3, owner 2026-11-08).
enum AppTab: String, CaseIterable, Identifiable {
    case home, guild, gear, activity

    var id: String { rawValue }

    var title: String {
        switch self {
        case .home: Copy.tabHome
        case .guild: Copy.tabGuild
        case .gear: Copy.tabGear
        case .activity: Copy.tabActivity
        }
    }

    /// The tab icon (Brand.xcassets `tab-<name>`): the owner's rune stone, helm and war horn, and
    /// Guild's the site's own Guild door, the banner (spec 2026-11-10-guild-tab.md section 1).
    var icon: String { "tab-\(rawValue)" }
}

/// Where a menu row goes (flows spec 3): the tools and the account chores, pushed under the tab
/// the menu was opened on.
enum MenuDestination: String, Hashable, Identifiable, CaseIterable {
    case guides, tiers, patchNotes, sims, reports, planner, rankings
    case notifications, plan, devices, offlineData, settings

    var id: String { rawValue }

    /// The tools a signed-out visitor can use.
    static let signedOutTools: [MenuDestination] = [.guides, .tiers, .patchNotes, .planner, .rankings]
    static let tools: [MenuDestination] = [.guides, .tiers, .patchNotes, .sims, .reports, .planner, .rankings]
    static let account: [MenuDestination] = [.notifications, .plan, .devices, .offlineData, .settings]
}

/// The shell's state: the selected tab, the profile menu, and what the menu pushed on each tab.
@MainActor
@Observable
final class AppShell {
    var selection: AppTab
    var menuOpen = false
    private(set) var pushed: [AppTab: MenuDestination] = [:]
    /// The tab badges (flows spec 3): Raid, the unanswered next raid or an officer's open items;
    /// Activity, the unread count. Zero shows none.
    var badges: [AppTab: Int] = [:]

    init(selection: AppTab = .home) {
        self.selection = selection
    }

    func openMenu() { menuOpen = true }

    func closeMenu() { menuOpen = false }

    /// A menu row: closes the menu and pushes the row's screen on the tab it was opened over.
    func open(_ destination: MenuDestination) {
        push(destination, on: selection)
        menuOpen = false
    }

    /// A link to a menu screen (`/app/guides`, `/app/you/notifications`): selects the tab and
    /// pushes the screen on it.
    func push(_ destination: MenuDestination, on tab: AppTab) {
        selection = tab
        pushed[tab] = destination
    }

    func binding(for tab: AppTab) -> Binding<MenuDestination?> {
        Binding(get: { self.pushed[tab] }, set: { self.pushed[tab] = $0 })
    }
}

/// What a tab root needs from the shell: which tab it is, the selected character for the profile
/// button, and the screens the menu pushes. RootView puts one in each tab's environment.
struct TabChrome {
    let tab: AppTab
    let shell: AppShell
    let character: AccountCharacter?
    let destination: (MenuDestination) -> AnyView
}

extension EnvironmentValues {
    @Entry var tabChrome: TabChrome?
}

/// A tab root's chrome (flows spec 3): the profile button top left, the left-edge swipe that opens
/// the menu (roots only: on a pushed screen the edge is iOS Back), and the menu's pushes. Applied
/// inside each tab's NavigationStack, on its root; a no-op outside the shell (tests, previews).
struct TabRootChrome: ViewModifier {
    @Environment(\.tabChrome) private var chrome

    /// The strip along the leading edge that takes the swipe.
    static let edgeWidth: CGFloat = 20
    /// How far the swipe travels before the menu opens.
    static let openDistance: CGFloat = 40

    func body(content: Content) -> some View {
        if let chrome {
            content
                .toolbar {
                    ToolbarItem(placement: .topBarLeading) {
                        ProfileButton(character: chrome.character) { chrome.shell.openMenu() }
                    }
                }
                .overlay(alignment: .leading) {
                    Color.clear
                        .frame(width: Self.edgeWidth)
                        .contentShape(.rect)
                        .gesture(DragGesture(minimumDistance: 10).onEnded { value in
                            if value.translation.width > Self.openDistance { chrome.shell.openMenu() }
                        })
                        .accessibilityHidden(true)
                }
                .navigationDestination(item: chrome.shell.binding(for: chrome.tab)) { chrome.destination($0) }
        } else {
            content
        }
    }
}

extension View {
    /// Marks a tab's root screen: the profile button, the edge swipe and the menu's pushes.
    func tabRootChrome() -> some View { modifier(TabRootChrome()) }
}
