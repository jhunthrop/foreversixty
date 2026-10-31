@testable import ForeverSixty
import Foundation
import Models
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// F8's states at iPhone 17 size (spec 5.8.1 acceptance list): You free in a link region and
/// elsewhere, Premium, the Plan screen in both regions, a seat, and You free at AX5. The same
/// images are the review captures in design/reviews/2026-10-31-mobile-premium.
@MainActor
@Suite("Premium snapshots")
struct PremiumSnapshotTests {
    @Test func youFreeLinkRegion() throws {
        let summary = try summary("me.json")
        snap(NavigationStack { YouView(summary: summary, mayLink: true) {} }, "you-free-us")
    }

    @Test func youFreeElsewhere() throws {
        let summary = try summary("me.json")
        snap(NavigationStack { YouView(summary: summary, mayLink: false) {} }, "you-free-gb")
    }

    @Test func youFreeLinkRegionLargest() throws {
        let summary = try summary("me.json")
        snap(NavigationStack { YouView(summary: summary, mayLink: true) {} }, "you-free-us-ax5", size: .accessibility5)
    }

    @Test func youPremiumLargest() throws {
        let summary = try summary("me.premium.json")
        snap(NavigationStack { YouView(summary: summary, mayLink: true) {} }, "you-premium-ax5", size: .accessibility5)
    }

    @Test func youSeat() throws {
        let summary = try summary("me.seat.json")
        snap(NavigationStack { YouView(summary: summary, mayLink: false) {} }, "you-seat")
    }

    @Test func planLinkRegion() throws {
        let plan = try plan("me.premium.json")
        snap(NavigationStack { PlanView(plan: plan, mayLink: true) }, "plan-us")
    }

    @Test func planElsewhere() throws {
        let plan = try plan("me.premium.json")
        snap(NavigationStack { PlanView(plan: plan, mayLink: false) }, "plan-gb")
    }

    @Test func planSeat() throws {
        let plan = try plan("me.seat.json")
        snap(NavigationStack { PlanView(plan: plan, mayLink: true) }, "plan-seat")
    }

    @Test func planLargest() throws {
        let plan = try plan("me.premium.json")
        snap(NavigationStack { PlanView(plan: plan, mayLink: true) }, "plan-us-ax5", size: .accessibility5)
    }

    private func summary(_ file: String) throws -> YouSummary {
        let me = try AccountSnapshotTests.me(file)
        return YouSummary(me: me, current: me.mainCharacter, deviceCount: 3, version: AccountSnapshotTests.version, now: PremiumTests.now)
    }

    private func plan(_ file: String) throws -> PlanSummary {
        PlanSummary(me: try AccountSnapshotTests.me(file), now: PremiumTests.now)
    }

    private func snap(_ view: some View, _ name: String, size: DynamicTypeSize = .large) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = UIContentSizeCategory(size)
        }
        assertSnapshot(
            of: view.environment(\.dynamicTypeSize, size).preferredColorScheme(.dark),
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: AccountSnapshotTests.iPhone17), traits: traits),
            named: name,
            testName: "premium"
        )
    }
}
