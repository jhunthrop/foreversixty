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

    @Test func accountDeleted() throws {
        let date = try #require(ISO8601DateFormatter().date(from: "2026-11-29T21:41:00Z"))
        snap(signedOutView(notice: AccountCopy.notice(.accountDeleted(completeBy: date)), failure: false), "signedout-deleted")
    }

    @Test func noCharacters() { snap(NoCharactersView {}, "home-no-characters") }

    @Test func signedOutLargest() { snap(signedOutView(notice: nil), "signedout-ax5", size: .accessibility5) }

    // MARK: F1.5

    @Test func switcher() throws {
        let me = try Self.threeCharacters()
        snap(switcherHost(me), "switcher-default")
    }

    @Test func switcherLargest() throws {
        let me = try Self.threeCharacters()
        snap(switcherHost(me), "switcher-ax5", size: .accessibility5)
    }

    // MARK: You, Devices, Delete account

    @Test func youFree() throws {
        let summary = try Self.summary("me.json")
        snap(NavigationStack { YouView(summary: summary) {} }, "you-free")
    }

    @Test func youPremium() throws {
        let summary = try Self.summary("me.premium.json")
        snap(NavigationStack { YouView(summary: summary) {} }, "you-premium")
    }

    @Test func about() {
        snap(NavigationStack { AboutView(version: Self.version) }, "about-default")
    }

    @Test func devices() throws {
        let rows = try Self.deviceRows()
        snap(NavigationStack { DevicesView(state: .loaded(rows), failedRevoke: nil, onRevoke: { _ in }, onRetry: {}) }, "devices-default")
    }

    @Test func devicesRevokeFailed() throws {
        let rows = try Self.deviceRows()
        let view = DevicesView(state: .loaded(rows), failedRevoke: rows[1].id, onRevoke: { _ in }, onRetry: {})
        snap(NavigationStack { view }, "devices-revoke-failed")
    }

    @Test func deleteFree() throws { snap(try deleteView("me.json"), "delete-free") }

    @Test func deletePremium() throws { snap(try deleteView("me.premium.json"), "delete-premium") }

    @Test func deleteGuildPayer() throws { snap(try deleteView("me.guild-payer.json"), "delete-guild-payer") }

    @Test func deleteSeat() throws { snap(try deleteView("me.seat.json"), "delete-seat") }

    @Test func deleteFailed() throws { snap(try deleteView("me.json", failed: true), "delete-failed") }

    private func deleteView(_ file: String, failed: Bool = false) throws -> some View {
        let consequences = DeletionConsequence.of(try Self.me(file))
        return NavigationStack {
            DeleteAccountView(consequences: consequences, isDeleting: false, failed: failed, onDelete: {})
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

    /// The selector open over a dimmed screen, and the sheet drawn as it presents.
    private func switcherHost(_ me: Me) -> some View {
        let current = me.characters.first
        return VStack(spacing: 0) {
            HStack {
                if let current { CharacterSelector(character: current, isOpen: true) {} }
                Spacer()
            }
            .padding(16)
            SwitcherSheet(me: me, currentKey: current?.key) { _ in }
                .clipShape(.rect(topLeadingRadius: 28, topTrailingRadius: 28))
        }
        .frame(maxHeight: .infinity, alignment: .top)
        .background(Color.black)
    }

    private func snap(_ view: some View, _ name: String, size: DynamicTypeSize = .large) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = UIContentSizeCategory(size)
        }
        assertSnapshot(
            of: view.environment(\.dynamicTypeSize, size).preferredColorScheme(.dark),
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: Self.iPhone17), traits: traits),
            named: name,
            testName: "account"
        )
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

    /// me.json plus a third character with no guild, as the spec's F1.5 capture asks.
    static func threeCharacters() throws -> Me {
        let data = try Data(contentsOf: fixtures.appending(path: "me.json"))
        var root = try #require(try JSONSerialization.jsonObject(with: data) as? [String: Any])
        var me = try #require(root["data"] as? [String: Any])
        var characters = try #require(me["characters"] as? [[String: Any]])
        characters.append([
            "key": "us/pvp/sablewhisper", "region": "us", "ruleset": "pvp", "name": "Sablewhisper",
            "class": "priest", "spec": "Shadow", "race": "undead", "level": 23, "source": "bnet",
        ])
        me["characters"] = characters
        root["data"] = me
        let patched = try JSONSerialization.data(withJSONObject: root)
        return try #require(try WireFormat.decoder().decode(Envelope<Me>.self, from: patched).data)
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
