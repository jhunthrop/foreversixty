import Auth
@testable import ForeverSixty
import Foundation
import Models
import Networking
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// F1's states at iPhone 17 size (spec §15), at the default text size and AX5 for the switcher
/// and the signed-out screen. The same images are the review captures in
/// design/reviews/2026-10-31-mobile-signin. Re-record with `SNAPSHOT_TESTING_RECORD=all`.
@MainActor
@Suite("Account snapshots")
struct AccountSnapshotTests {
    // MARK: F1.1 to F1.4

    @Test func signedOut() { snap(signedOutView(notice: nil), "signedout-default") }

    @Test func signingIn() { snap(signedOutView(notice: nil, signingIn: true), "signin-browser-open") }

    @Test func importing() { snap(ImportView(), "import-default") }

    @Test func signInFailed() {
        snap(signedOutView(notice: AccountCopy.notice(.signInFailed(.battleNet))), "signin-failed")
    }

    @Test func sessionEnded() { snap(signedOutView(notice: AccountCopy.notice(.sessionEnded)), "signedout-401") }

    @Test func accountDeleted() {
        snap(signedOutView(notice: AccountCopy.notice(.accountDeleted), failure: false), "signedout-deleted")
    }

    @Test func noCharacters() { snap(NoCharactersView {}, "home-no-characters") }

    @Test func signedOutLargest() { snap(signedOutView(notice: nil), "signedout-ax5", size: .accessibility5) }

    // MARK: You, Devices, Delete account

    @Test func youFree() throws {
        let summary = try Self.summary("me.json")
        snap(PushedScreen(back: AppTab.home.title) { YouView(summary: summary) {} }, "you-free")
    }

    @Test func youPremium() throws {
        let summary = try Self.summary("me.premium.json")
        snap(PushedScreen(back: AppTab.home.title) { YouView(summary: summary) {} }, "you-premium")
    }

    @Test func about() {
        snap(PushedScreen(back: Copy.youTitle) { AboutView(version: Self.version) }, "about-default")
    }

    @Test func devices() throws {
        let rows = try Self.deviceRows()
        snap(PushedScreen(back: Copy.youTitle) { DevicesView(state: .loaded(rows), failedRevoke: nil, onRevoke: { _ in }, onRetry: {}) }, "devices-default")
    }

    @Test func devicesRevokeFailed() throws {
        let rows = try Self.deviceRows()
        let view = DevicesView(state: .loaded(rows), failedRevoke: rows[1].id, onRevoke: { _ in }, onRetry: {})
        snap(PushedScreen(back: Copy.youTitle) { view }, "devices-revoke-failed")
    }

    @Test func deleteFree() throws { snap(try deleteView("me.json"), "delete-free") }

    @Test func deletePremium() throws { snap(try deleteView("me.premium.json"), "delete-premium") }

    @Test func deleteSeat() throws { snap(try deleteView("me.seat.json"), "delete-seat") }

    @Test func deleteFailed() throws { snap(try deleteView("me.json", outcome: .failed), "delete-failed") }

    @Test func deleteGuildObligation() throws {
        let message = "You hold the claim on Olympus XXVII, so deleting your account would cancel its Guild plan "
            + "for the whole guild right away, with no refund. Release the claim on your account page first."
        snap(try deleteView("me.guild-payer.json", outcome: .refused(message: message)), "delete-guild-obligation")
    }

    @Test func deleteMismatch() throws {
        snap(try deleteView("me.json", typed: "Yell#16", outcome: .confirmMismatch), "delete-mismatch")
    }

    @Test func deleteReady() throws { snap(try deleteView("me.json", typed: "yell#1604"), "delete-ready") }

    // To the end, as Android's scrolled captures: the whole page on a tall canvas, so the field, its
    // failure and the button show (an off-screen scroll to the bottom renders blank).

    @Test func deleteMismatchEnd() throws {
        snap(try deleteView("me.json", typed: "Yell#16", outcome: .confirmMismatch), "delete-mismatch-end", height: Self.wholePage)
    }

    @Test func deleteReadyEnd() throws { snap(try deleteView("me.json", typed: "yell#1604"), "delete-ready-end", height: Self.wholePage) }

    @Test func deleteFailedEnd() throws { snap(try deleteView("me.json", outcome: .failed), "delete-failed-end", height: Self.wholePage) }

    /// A canvas tall enough for the whole delete page.
    static let wholePage: CGFloat = 1600

    private func deleteView(_ file: String, typed: String = "", outcome: AccountActionOutcome? = nil) throws -> some View {
        let me = try Self.me(file)
        return PushedScreen(back: Copy.youTitle) {
            DeleteAccountView(
                consequences: DeletionConsequence.of(me),
                battleTag: me.user.battletag,
                typed: .constant(typed),
                isDeleting: false,
                outcome: outcome,
                onDelete: {}
            )
        }
    }

    static let version = "Version 0.1.0 (1) · data 1.60.1.70658"

    static func summary(_ file: String) throws -> YouSummary {
        let me = try me(file)
        return YouSummary(me: me, current: me.mainCharacter, deviceCount: 3, version: version)
    }

    // MARK: Helpers

    private func signedOutView(notice: String?, signingIn: Bool = false, failure: Bool = true) -> some View {
        SignedOutView(notice: notice, noticeIsFailure: failure, isSigningIn: signingIn, onSignIn: {}, onPickClass: {})
    }

    private func snap(_ view: some View, _ name: String, size: DynamicTypeSize = .large, height: CGFloat? = nil) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = UIContentSizeCategory(size)
        }
        assertSnapshot(
            of: view.environment(\.dynamicTypeSize, size).preferredColorScheme(.dark),
            as: .image(precision: SnapshotTolerance.precision, perceptualPrecision: SnapshotTolerance.perceptualPrecision, layout: .device(config: Self.device(height: height)), traits: traits),
            named: name,
            testName: "account"
        )
    }

    /// iPhone 17, or the same width at `height` for a whole-page capture.
    static func device(height: CGFloat?) -> ViewImageConfig {
        guard let height else { return iPhone17 }
        return ViewImageConfig(safeArea: iPhone17.safeArea, size: CGSize(width: iPhone17.size?.width ?? 0, height: height), traits: iPhone17.traits)
    }

    /// iPhone 17: 402 x 874 pt with its safe area.
    static let iPhone17 = ViewImageConfig(
        safeArea: UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0),
        size: CGSize(width: 402, height: 874),
        traits: UITraitCollection(userInterfaceStyle: .dark)
    )

    static func me(_ file: String) throws -> Me {
        let data = try Data(contentsOf: fixtures.appending(path: file))
        return try #require(try WireFormat.decoder().decode(Envelope<Me>.self, from: data).data)
    }

    static func deviceRows() throws -> [DeviceRow] {
        let data = try Data(contentsOf: fixtures.appending(path: "devices.list.json"))
        let devices = try #require(try WireFormat.decoder().decode(Envelope<[Device]>.self, from: data).data)
        let now = try #require(ISO8601DateFormatter().date(from: "2026-10-30T21:41:00Z"))
        return DeviceList.rows(devices, thisDevice: "dvc_8Qk2nXw4", now: now)
    }

    /// docs/contracts/fixtures/app, read in place.
    static let fixtures = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent().deletingLastPathComponent().deletingLastPathComponent().deletingLastPathComponent()
        .appending(path: "docs/contracts/fixtures/app")
}
