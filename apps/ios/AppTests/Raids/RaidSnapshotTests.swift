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
        // The repo's game icons, as the raid flows' tests paint them.
        let sized = view
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.raidIconBuild, HomeFlowTestData.build)
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(precision: SnapshotTolerance.precision, perceptualPrecision: SnapshotTolerance.perceptualPrecision, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    // MARK: Raid page (F4.2, F4.6 to F4.12, F4.18 to F4.21)

    private func page(_ controller: SignupController, offline: Date? = nil) -> some View {
        PushedScreen(back: AppTab.guild.title) {
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
        await controller.press("us/pvp/gorvash", .late)
        snap(page(controller), "raid-page-save-failed-ios")
    }

    @Test func raidPage_rateLimited() async throws {
        let service = FakeRaidService()
        service.signupAnswer = { _ throws(APIError) in throw (try? RaidFixture.error("guild.signup.error-rate-limited", status: 429)) ?? .unexpectedStatus(0) }
        let controller = try makeController(service: service)
        await controller.press("us/pvp/gorvash", .bench)
        snap(page(controller), "raid-page-rate-limited-ios")
    }

    @Test func raidPage_raceLocked() async throws {
        let service = FakeRaidService()
        service.signupAnswer = { _ throws(APIError) in throw (try? RaidFixture.error("guild.signup.error-raid-locked", status: 409)) ?? .unexpectedStatus(0) }
        let controller = try makeController(service: service)
        await controller.press("us/pvp/gorvash", .late)
        snap(page(controller), "raid-page-race-locked-ios")
    }

    @Test func raidPage_moved() async throws {
        let service = FakeRaidService()
        let answer = try RaidFixture.load(SignupAnswer.self, "guild.signup.moved")
        service.signupAnswer = { _ in answer }
        let controller = try makeController(try RaidFixture.raid("guild.raid.two-characters"), service: service)
        await controller.press("us/pvp/yellfury", .signedIn)
        snap(page(controller), "raid-page-moved-ios", height: 1100)
    }

    @Test func raidPage_reminderAsk() async throws {
        let controller = try makeController(undetermined: true)
        await controller.press("us/pvp/gorvash", .signedIn)
        snap(page(controller), "raid-page-reminder-ask-ios", height: 1100)
    }

    @Test func raidPage_specPending() async throws {
        let controller = try makeController(try RaidFixture.raid("guild.raid.ask-spec"))
        await controller.press("us/pvp/lightningbolltz", .signedIn)
        snap(page(controller), "raid-page-spec-pending-ios")
    }

    @Test func raidPage_done() throws { snap(page(try makeController(try RaidFixture.raid("guild.raid.done"))), "raid-page-done-ios") }

    @Test func raidPage_officer() throws { snap(page(try makeController(try RaidFixture.raid("guild.raid.officer"))), "raid-page-officer-ios", height: 3200) }

    @Test func raidPage_officer_xl() throws {
        snap(page(try makeController(try RaidFixture.raid("guild.raid.officer"))), "raid-page-officer-xl-ios", height: 3600, large: true)
    }

    @Test func raidPage_officerLocked() async throws {
        let service = FakeRaidService()
        service.lockAnswer = { () throws(APIError) in try RaidFixture.api(RaidView.self, "guild.raid.officer-locked") }
        let controller = try makeController(try RaidFixture.raid("guild.raid.officer"), service: service)
        await controller.toggleLock()
        snap(page(controller), "raid-page-officer-locked-ios")
    }

    // MARK: Sheets and Your raid night (F4.3, F4.22, F4.23)

    /// The spec sheet at its medium detent over the raid page, with the grabber, as it presents.
    private func sheet(_ raid: String, ask: SpecAsk, key: String, height: CGFloat = Self.sheetHeight) throws -> some View {
        let controller = try makeController(try RaidFixture.raid(raid))
        let character = try #require(controller.character(key))
        return SheetFrame(height: height, behind: page(controller)) {
            SpecSheet(ask: ask, character: character, raidTitle: "Onyxia's Lair", date: "Sat 19 Dec") { _ in }
        }
    }

    /// The medium detent on the 874 pt phone, and the sheet's height on the taller AX5 canvas.
    static let sheetHeight: CGFloat = 520
    static let sheetHeightLarge: CGFloat = 1300

    @Test func specSheet_ask() throws {
        snap(try sheet("guild.raid.ask-spec", ask: .ask(characterKey: "us/pvp/lightningbolltz", state: .signedIn), key: "us/pvp/lightningbolltz"), "spec-sheet-ask-ios")
    }

    @Test func specSheet_ask_xl() throws {
        snap(try sheet("guild.raid.ask-spec", ask: .ask(characterKey: "us/pvp/lightningbolltz", state: .signedIn), key: "us/pvp/lightningbolltz", height: Self.sheetHeightLarge), "spec-sheet-ask-xl-ios", height: 1600, large: true)
    }

    @Test func specSheet_change() throws {
        snap(try sheet("guild.raid", ask: .change(characterKey: "us/pvp/gorvash"), key: "us/pvp/gorvash"), "spec-sheet-change-ios")
    }

    private func night(_ fixture: String) throws -> some View {
        let night = try RaidFixture.load(RaidNightSummary.self, fixture)
        let model = RaidNightModel(raid: 29, service: FakeRaidService(), cache: NoRaidCache(), now: { RaidTime.now }, initial: .loaded(night))
        return PushedScreen(back: AppTab.guild.title) { RaidNightView(model: model, clock: RaidTime.clock, now: { RaidTime.now }) }
    }

    @Test func raidNight_premium() throws { snap(try night("me.raid-night"), "raid-night-premium-ios") }

    @Test func raidNight_free() throws { snap(try night("me.raid-night.free"), "raid-night-free-ios") }

    @Test func raidNight_xl() throws { snap(try night("me.raid-night"), "raid-night-xl-ios", height: 2400, large: true) }
}
