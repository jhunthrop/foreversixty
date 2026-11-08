@testable import ForeverSixty
import DesignSystem
import Foundation
import Models
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// The mocks' account (design/mocks/app/ios-first-*.html): six characters, two in Olympus XXVII.
enum FirstRunTestData {
    static let guild = CharacterGuild(id: 2, name: "Olympus XXVII", rank: "Raider", verified: true)

    static func character(_ name: String, _ className: String, _ spec: String, _ level: Int, guild: CharacterGuild? = nil) -> AccountCharacter {
        AccountCharacter(
            key: "us/normal/\(name.lowercased())", region: "us", ruleset: "normal", name: name, source: "bnet",
            className: className, level: level, faction: "alliance", spec: spec, guild: guild
        )
    }

    static let characters = [
        character("Brewmist", "shaman", "Enhancement", 22),
        character("Thandrel", "warrior", "Fury", 60, guild: guild),
        character("Lightvow", "paladin", "Holy", 60, guild: guild),
        character("Kaelor", "mage", "Frost", 47),
        character("Felsnap", "hunter", "Beast Mastery", 34),
        character("Lowbie", "rogue", "Combat", 9),
        character("Ashmantle", "warlock", "Affliction", 18),
    ]

    static func me(main: String? = nil, characters: [AccountCharacter] = characters) -> Me {
        Me(
            user: AppUser(id: 1, battletag: "Thandrel#1234", role: "user", anonymize: false),
            characters: characters,
            entitlements: Entitlements(
                serverSims: false, retention: false, multiCompare: false, history: false, notifications: false,
                officerViews: false, rosterCheck: false, logCoach: false, supporterMark: false, personalPortal: false
            ),
            discordLinkAvailable: false,
            mainCharacterKey: main
        )
    }
}

@MainActor
@Suite("First run (flows spec F8)")
struct FirstRunModelTests {
    final class Permission: NotificationPermission, @unchecked Sendable {
        var undetermined = true
        var requested = 0
        func isUndetermined() async -> Bool { undetermined }
        func request() async { requested += 1 }
    }

    final class Asked: FirstRunAskStore, @unchecked Sendable {
        var asked = false
        func wasAsked() -> Bool { asked }
        func markAsked() { asked = true }
    }

    struct Refused: Error {}

    private func model(fails: Bool = false, permission: Permission = Permission(), asked: Asked = Asked()) -> (FirstRunModel, [String]) {
        var saved: [String] = []
        let model = FirstRunModel(dependencies: .init(
            setMain: { key in
                if fails { throw Refused() }
                saved.append(key)
            },
            permission: permission,
            asked: asked
        ))
        return (model, saved)
    }

    @Test("Pick your main offers level 10 and up, the highest first, the first of them preselected")
    func candidates() {
        let names = FirstRunModel.candidates(FirstRunTestData.me()).map(\.name)
        #expect(names == ["Thandrel", "Lightvow", "Kaelor", "Felsnap", "Brewmist", "Ashmantle"])
        let (model, _) = model()
        model.update(me: FirstRunTestData.me())
        #expect(model.step == .pickMain)
        #expect(model.selection == "us/normal/thandrel")
    }

    @Test("an account with a main, or with nobody at 10, is never asked")
    func noAsk() {
        let (model, _) = model()
        model.update(me: FirstRunTestData.me(main: "us/normal/kaelor"))
        #expect(model.step == nil)
        model.update(me: FirstRunTestData.me(characters: [FirstRunTestData.character("Lowbie", "rogue", "Combat", 9)]))
        #expect(model.step == nil)
    }

    @Test("a guild main leads to the notification ask; Turn on asks the system once")
    func guildAsk() async {
        let permission = Permission()
        let asked = Asked()
        let (model, _) = model(permission: permission, asked: asked)
        let me = FirstRunTestData.me()
        model.update(me: me)
        await model.continueWithMain(me)
        #expect(model.step == .permission(guild: "Olympus XXVII"))
        await model.turnOnNotifications()
        #expect(model.step == nil)
        #expect(permission.requested == 1)
        #expect(asked.asked)
    }

    @Test("no guild, an earlier ask or a decided system: Home follows the pick")
    func noGuildAsk() async {
        let (noGuild, _) = model()
        let me = FirstRunTestData.me()
        noGuild.update(me: me)
        noGuild.selection = "us/normal/kaelor"
        await noGuild.continueWithMain(me)
        #expect(noGuild.step == nil)

        let asked = Asked()
        asked.asked = true
        let (again, _) = model(asked: asked)
        again.update(me: me)
        await again.continueWithMain(me)
        #expect(again.step == nil)

        let decided = Permission()
        decided.undetermined = false
        let (system, _) = model(permission: decided)
        system.update(me: me)
        await system.continueWithMain(me)
        #expect(system.step == nil)
    }

    @Test("a refused save keeps the picker with its line; Not now never asks the system")
    func failure() async {
        let (model, _) = model(fails: true)
        let me = FirstRunTestData.me()
        model.update(me: me)
        await model.continueWithMain(me)
        #expect(model.step == .pickMain)
        #expect(model.failed)

        let permission = Permission()
        let (ask, _) = self.model(permission: permission)
        ask.update(me: me)
        await ask.continueWithMain(me)
        ask.notNow()
        #expect(ask.step == nil)
        #expect(permission.requested == 0)
    }
}

/// The parity states (design/parity.md "Native"): regions against the mock's region file, the
/// text golden and the default and AX5 image goldens.
@MainActor
@Suite("First run parity")
struct FirstRunParityTests {
    static let device = CGSize(width: 393, height: 852)

    private var pickMain: some View {
        PickMainView(
            characters: FirstRunModel.candidates(FirstRunTestData.me()), selection: "us/normal/thandrel",
            saving: false, failed: false, onSelect: { _ in }, onContinue: {}
        )
    }

    private var permission: some View {
        FirstPermissionView(guild: "Olympus XXVII", onTurnOn: {}, onNotNow: {})
    }

    @Test("ios-first-pick-main")
    func pickMainParity() async throws {
        try await ParityCheck.assert(pickMain, screen: "first-pick-main", size: Self.device)
    }

    @Test("ios-first-permission")
    func permissionParity() async throws {
        try await ParityCheck.assert(permission, screen: "first-permission", size: Self.device)
    }

    @Test("ios-first-permission-long: a long guild name wraps over two balanced lines")
    func permissionLongParity() async throws {
        let view = FirstPermissionView(guild: "Ranchers of Thunderbluff", onTurnOn: {}, onNotNow: {})
        try await ParityCheck.assert(view, screen: "first-permission-long", size: Self.device)
    }

    @Test("the balanced break makes the longer line shortest")
    func balanced() {
        #expect(BalancedLines.two("Ranchers of Thunderbluff", role: FSType.title) == "Ranchers of\nThunderbluff")
        #expect(BalancedLines.two("Olympus", role: FSType.title) == "Olympus")
    }
}
