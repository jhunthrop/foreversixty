@testable import ForeverSixty
import Models
import Networking
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// Every F4 screen and state on iPhone 17 (402 x 874), at the default size and at AX5 (spec 15).
/// Recorded like HomeSnapshotTests (`SNAPSHOT_TESTING_RECORD=all`); the PNGs double as the review
/// captures in design/reviews/2026-10-31-mobile-raids.
@MainActor
@Suite("Raids snapshots")
struct RaidSnapshotTests {
    static let phone = CGSize(width: 402, height: 874)
    static let safeArea = UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0)

    private func snap(_ view: some View, _ name: String, height: CGFloat = 874, large: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let config = ViewImageConfig(safeArea: Self.safeArea, size: CGSize(width: Self.phone.width, height: height), traits: traits)
        let sized = view
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    // MARK: Raids tab (F4.1, F4.13 to F4.17)

    private func tab(_ state: RaidsTabState, controller: SignupController? = nil) -> some View {
        NavigationStack {
            RaidsTabContentView(
                state: state, clock: RaidTime.clock, now: { RaidTime.now }, nextController: controller,
                openRaid: { _, _ in }, openNight: { _ in }, openWeb: { _ in }, retry: {}
            )
            .navigationTitle(Copy.tabRaids)
            .toolbarTitleDisplayMode(.inlineLarge)
        }
    }

    private func loadedContent(offline: Date? = nil, planned: [RaidView]? = nil) throws -> RaidsTabContent {
        let me = try RaidFixture.load(Me.self, "me")
        let raids = try RaidFixture.load(GuildRaids.self, "guild.raids")
        let character = try #require(me.mainCharacter)
        return RaidsTabContent(
            head: GuildHead(character: try #require(character.guild), guild: me.guilds.first),
            character: character, planned: planned ?? raids.planned ?? [], nights: raids.rows,
            waiting: raids.waitingForApproval, offlineSince: offline
        )
    }

    @Test func raidsTab_loaded() throws {
        let content = try loadedContent()
        snap(tab(.loaded(content), controller: try makeController(content.planned.first)), "raids-tab-loaded-ios", height: 1300)
    }

    @Test func raidsTab_loaded_xl() throws {
        let content = try loadedContent()
        snap(tab(.loaded(content), controller: try makeController(content.planned.first)), "raids-tab-loaded-xl-ios", height: 2600, large: true)
    }

    @Test func raidsTab_noGuild() throws { snap(tab(.noGuild(name: "Yellfury")), "raids-tab-no-guild-ios") }

    @Test func raidsTab_waiting() throws {
        let me = try RaidFixture.load(Me.self, "me.waiting")
        let character = try #require(me.mainCharacter)
        let head = GuildHead(character: try #require(character.guild), guild: me.guilds.first)
        snap(tab(.waiting(head: head, character: character)), "raids-tab-waiting-ios")
    }

    @Test func raidsTab_empty() throws { snap(tab(.loaded(try loadedContent(planned: []))), "raids-tab-no-raid-ios") }

    @Test func raidsTab_error() throws { snap(tab(.failed), "raids-tab-error-ios") }

    @Test func raidsTab_offline() throws {
        let content = try loadedContent(offline: RaidTime.now.addingTimeInterval(-3600))
        let controller = try makeController(content.planned.first)
        controller.offline = true
        snap(tab(.loaded(content), controller: controller), "raids-tab-offline-ios")
    }

    @Test func raidsTab_signedOut() throws { snap(tab(.signedOut), "raids-tab-signed-out-ios") }

    @Test func raidsTab_loading() throws { snap(tab(.loading), "raids-tab-loading-ios") }

    // MARK: Raid page (F4.2, F4.6 to F4.12, F4.18 to F4.21)

    private func page(_ controller: SignupController, offline: Date? = nil) -> some View {
        NavigationStack {
            RaidPageContent(
                controller: controller, clock: RaidTime.clock, now: { RaidTime.now }, currentKey: nil,
                offlineSince: offline, guildName: "OLYMPUS XXVII", openNight: { _ in }
            )
        }
    }

    @Test func raidPage_member() throws { snap(page(try makeController()), "raid-page-member-ios") }

    @Test func raidPage_full() throws { snap(page(try makeController()), "raid-page-full-ios", height: 3000) }

    @Test func raidPage_xl() throws { snap(page(try makeController()), "raid-page-xl-ios", height: 2400, large: true) }

    @Test func raidPage_unsigned() async throws {
        snap(page(try makeController(try RaidFixture.raid("guild.raid.ask-spec"))), "raid-page-unsigned-ios")
    }

    @Test func raidPage_noInvite() throws { snap(page(try makeController(try RaidFixture.raid("guild.raid.no-invite"))), "raid-page-no-invite-ios") }

    @Test func raidPage_locked() throws { snap(page(try makeController(try RaidFixture.raid("guild.raid.locked"))), "raid-page-locked-ios") }

    @Test func raidPage_cancelled() throws { snap(page(try makeController(try RaidFixture.raid("guild.raid.cancelled"))), "raid-page-cancelled-ios") }

    @Test func raidPage_offline() throws {
        let controller = try makeController()
        controller.offline = true
        snap(page(controller, offline: RaidTime.now.addingTimeInterval(-600)), "raid-page-offline-ios")
    }

    @Test func raidPage_saveFailed() async throws {
        let service = FakeRaidService()
        service.signupAnswer = { _ throws(APIError) in throw .unexpectedStatus(500) }
        let controller = try makeController(service: service)
        await controller.press("us/pvp/obnoxious-yell", .late)
        snap(page(controller), "raid-page-save-failed-ios")
    }

    @Test func raidPage_rateLimited() async throws {
        let service = FakeRaidService()
        service.signupAnswer = { _ throws(APIError) in throw (try? RaidFixture.error("guild.signup.error-rate-limited", status: 429)) ?? .unexpectedStatus(0) }
        let controller = try makeController(service: service)
        await controller.press("us/pvp/obnoxious-yell", .bench)
        snap(page(controller), "raid-page-rate-limited-ios")
    }

    @Test func raidPage_raceLocked() async throws {
        let service = FakeRaidService()
        service.signupAnswer = { _ throws(APIError) in throw (try? RaidFixture.error("guild.signup.error-raid-locked", status: 409)) ?? .unexpectedStatus(0) }
        let controller = try makeController(service: service)
        await controller.press("us/pvp/obnoxious-yell", .late)
        snap(page(controller), "raid-page-race-locked-ios")
    }

    @Test func raidPage_moved() async throws {
        let service = FakeRaidService()
        let answer = SignupAnswer(raid: try RaidFixture.raid("guild.raid.two-characters"), moved: "Yellfury")
        service.signupAnswer = { _ in answer }
        let controller = try makeController(try RaidFixture.raid("guild.raid.two-characters"), service: service)
        await controller.press("us/pvp/obnoxious-yell", .signedIn)
        snap(page(controller), "raid-page-moved-ios", height: 1100)
    }

    @Test func raidPage_reminderAsk() async throws {
        let controller = try makeController(undetermined: true)
        await controller.press("us/pvp/obnoxious-yell", .signedIn)
        snap(page(controller), "raid-page-reminder-ask-ios", height: 1100)
    }

    @Test func raidPage_specPending() async throws {
        let controller = try makeController(try RaidFixture.raid("guild.raid.ask-spec"))
        await controller.press("us/pvp/lightningbolltz", .signedIn)
        snap(page(controller), "raid-page-spec-pending-ios")
    }

    @Test func raidPage_done() throws { snap(page(try makeController(try RaidFixture.raid("guild.raid.done"))), "raid-page-done-ios") }

    @Test func raidPage_officer() throws { snap(page(try makeController(try RaidFixture.raid("guild.raid.officer"))), "raid-page-officer-ios", height: 3200) }

    @Test func raidPage_officerLocked() async throws {
        let service = FakeRaidService()
        service.lockAnswer = { () throws(APIError) in try RaidFixture.api(RaidView.self, "guild.raid.officer-locked") }
        let controller = try makeController(try RaidFixture.raid("guild.raid.officer"), service: service)
        await controller.toggleLock()
        snap(page(controller), "raid-page-officer-locked-ios")
    }

    // MARK: Sheets and Your raid night (F4.3, F4.22, F4.23)

    private func sheet(_ raid: String, ask: SpecAsk, key: String) throws -> some View {
        let controller = try makeController(try RaidFixture.raid(raid))
        return SpecSheet(ask: ask, character: try #require(controller.character(key)), raidTitle: "Onyxia's Lair", date: "Sat 19 Dec") { _ in }
    }

    @Test func specSheet_ask() throws {
        snap(try sheet("guild.raid.ask-spec", ask: .ask(characterKey: "us/pvp/lightningbolltz", state: .signedIn), key: "us/pvp/lightningbolltz"), "spec-sheet-ask-ios")
    }

    @Test func specSheet_ask_xl() throws {
        snap(try sheet("guild.raid.ask-spec", ask: .ask(characterKey: "us/pvp/lightningbolltz", state: .signedIn), key: "us/pvp/lightningbolltz"), "spec-sheet-ask-xl-ios", height: 1600, large: true)
    }

    @Test func specSheet_change() throws {
        snap(try sheet("guild.raid", ask: .change(characterKey: "us/pvp/obnoxious-yell"), key: "us/pvp/obnoxious-yell"), "spec-sheet-change-ios")
    }

    private func night(_ fixture: String) throws -> some View {
        let night = try RaidFixture.load(RaidNightSummary.self, fixture)
        let model = RaidNightModel(raid: 29, service: FakeRaidService(), cache: NoRaidCache(), now: { RaidTime.now }, initial: .loaded(night))
        return NavigationStack { RaidNightView(model: model, clock: RaidTime.clock, now: { RaidTime.now }) }
    }

    @Test func raidNight_premium() throws { snap(try night("me.raid-night"), "raid-night-premium-ios") }

    @Test func raidNight_free() throws { snap(try night("me.raid-night.free"), "raid-night-free-ios") }
}
