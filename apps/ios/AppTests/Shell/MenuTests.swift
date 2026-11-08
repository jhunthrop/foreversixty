@testable import ForeverSixty
import Foundation
import Models
import SwiftUI
import Testing
import UIKit

/// The mocks' account for the menu (design/mocks/ws_app_flows.py CHARACTERS and LOW): Thandrel the
/// main, five more at 10 and up, two below; Battle.net portraits from design/mocks/data/avatars.
enum MenuTestData {
    static let now = Date(timeIntervalSince1970: 1_798_740_000)
    static let guild = CharacterGuild(id: 2, name: "Olympus XXVII", rank: "Raider", verified: true)

    static func character(
        _ name: String, _ className: String, _ spec: String?, _ level: Int, faction: String,
        avatar: Bool = true, guild: CharacterGuild? = nil
    ) -> AccountCharacter {
        let url = avatar ? HomeFixtures.repoRoot.appending(path: "design/mocks/data/avatars/\(name.lowercased()).jpg").absoluteString : nil
        return AccountCharacter(
            key: "us/normal/\(name.lowercased())", region: "us", ruleset: "normal", name: name, source: "bnet",
            className: className, level: level, faction: faction, avatarUrl: url, spec: spec, guild: guild
        )
    }

    static let thandrel = character("Thandrel", "warrior", "Fury", 60, faction: "alliance", guild: guild)
    static let others = [
        character("Lightvow", "paladin", "Holy", 60, faction: "alliance", guild: guild),
        character("Kaelor", "mage", "Frost", 47, faction: "horde"),
        character("Felsnap", "hunter", "Beast Mastery", 34, faction: "alliance"),
        character("Brewmist", "shaman", "Enhancement", 22, faction: "horde"),
        character("Ashmantle", "warlock", "Affliction", 18, faction: "alliance", avatar: false),
    ]
    static let low = [
        character("Pockets", "rogue", nil, 6, faction: "alliance", avatar: false),
        character("Bankalt", "mage", nil, 1, faction: "alliance", avatar: false),
    ]

    static func me(_ characters: [AccountCharacter]) -> Me {
        Me(
            user: AppUser(id: 1, battletag: "Thandrel#1234", role: "user", anonymize: false),
            characters: characters,
            entitlements: Entitlements(
                serverSims: true, retention: false, multiCompare: false, history: false, notifications: true,
                officerViews: false, rosterCheck: false, logCoach: false, supporterMark: false, personalPortal: false
            ),
            discordLinkAvailable: false,
            mainCharacterKey: thandrel.key,
            bnetImportedAt: now.addingTimeInterval(-2 * 3600)
        )
    }

    static let facts = ProfileMenuSummary.AccountFacts(
        notificationGroups: ["Raids", "Gear and sims", "Logs"], plan: "Premium · Founder price", devices: 3, offlineBytes: 26_600_000
    )
    static let version = "Version 0.1.0 (7) · data 1.60.1.70790"

    static func summary(_ characters: [AccountCharacter]?) -> ProfileMenuSummary {
        ProfileMenuSummary(
            me: characters.map(me), current: characters == nil ? nil : thandrel, facts: facts, version: version, now: now
        )
    }
}

/// Reads the mocks' portrait stand-ins from disk, at once.
struct FilePortraitSource: PortraitSource {
    func cached(_ url: URL) -> UIImage? { UIImage(contentsOfFile: url.path) }
    func load(_ url: URL) async -> UIImage? { cached(url) }
}

@Suite("Profile menu")
struct ProfileMenuSummaryTests {
    @Test("You is the selected character; Switch lists the rest at 10 and up, the lower ones fold")
    func characters() {
        let summary = MenuTestData.summary([MenuTestData.thandrel] + MenuTestData.others + MenuTestData.low)
        #expect(summary.you?.name == "Thandrel")
        #expect(summary.isMain)
        #expect(summary.others.map(\.name) == ["Lightvow", "Kaelor", "Felsnap", "Brewmist", "Ashmantle"])
        #expect(summary.lowLevel.map(\.name) == ["Pockets", "Bankalt"])
        #expect(summary.stamp == "Updated 2 h ago from Battle.net")
    }

    @Test("the tools and the account rows, in the owner's order, with their lines")
    func rows() {
        let summary = MenuTestData.summary([MenuTestData.thandrel])
        #expect(summary.tools.map(\.title) == ["Guides", "Tiers", "Patch notes", "Sims", "Logs and reports", "Planner", "Rankings"])
        #expect(summary.account.map(\.line) == [
            "Raids, gear and sims, logs on", "Premium · Founder price", "3 signed in", "26.6 MB on this phone", "Sign out, delete account",
        ])
    }

    @Test("signed out: no You, no account rows, only the tools that work without an account")
    func signedOut() {
        let summary = MenuTestData.summary(nil)
        #expect(summary.you == nil)
        #expect(summary.account.isEmpty)
        #expect(summary.tools.map(\.destination) == [.guides, .tiers, .patchNotes, .planner, .rankings])
    }

    @Test("every notification group off reads All off")
    func allOff() {
        #expect(ProfileMenuSummary.notificationsLine([]) == "All off")
        #expect(ProfileMenuSummary.notificationsLine(["Raids"]) == "Raids on")
    }
}

@MainActor
@Suite("Shell")
struct AppShellTests {
    @Test("a menu row closes the menu and pushes on the tab it was opened over")
    func open() {
        let shell = AppShell(selection: .gear)
        shell.openMenu()
        shell.open(.sims)
        #expect(!shell.menuOpen)
        #expect(shell.pushed[.gear] == .sims)
        #expect(shell.pushed[.home] == nil)
    }

    @Test("a link to a menu screen selects its tab and pushes there")
    func push() {
        let shell = AppShell(selection: .activity)
        shell.push(.guides, on: .home)
        #expect(shell.selection == .home)
        #expect(shell.pushed[.home] == .guides)
    }

    @Test("account links land under Home")
    func accountLinks() {
        #expect(RootView.accountDestination(.notifications) == .notifications)
        #expect(RootView.accountDestination(.plan) == .plan)
        #expect(RootView.accountDestination(.guides) == nil)
    }
}

@MainActor
@Suite("Menu parity")
struct MenuParityTests {
    /// The drawer: 85% of the 393 pt screen.
    static let size = CGSize(width: 393 * ProfileDrawer<EmptyView>.widthShare, height: 852)

    private func menu(_ characters: [AccountCharacter]?, low: Bool = false) -> some View {
        ProfileMenuView(summary: MenuTestData.summary(characters), showsLowLevel: low)
            .environment(\.portraitSource, FilePortraitSource())
    }

    private var everyone: [AccountCharacter] { [MenuTestData.thandrel] + MenuTestData.others + MenuTestData.low }

    @Test("ios-menu")
    func full() async throws {
        try await ParityCheck.assert(menu(everyone), screen: "menu", size: Self.size)
    }

    @Test("ios-menu-one")
    func one() async throws {
        try await ParityCheck.assert(menu([MenuTestData.thandrel]), screen: "menu-one", size: Self.size)
    }

    @Test("ios-menu-low")
    func low() async throws {
        try await ParityCheck.assert(menu(everyone, low: true), screen: "menu-low", size: Self.size)
    }

    @Test("ios-menu-out")
    func out() async throws {
        try await ParityCheck.assert(menu(nil), screen: "menu-out", size: Self.size)
    }
}
