@testable import ForeverSixty
import DesignSystem
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// Lane M0's design-system sheet (App/System/DesignSystemSheet.swift) on iPhone 17's width
/// (402 pt), the whole sheet in one image, drawn through the key window so the tab bar and its
/// icons render. The PNG is the lane's review capture (design/reviews/2026-11-10-m0).
@MainActor
@Suite("Design system sheet")
struct DesignSystemSheetSnapshotTests {
    private static let width: CGFloat = 402
    private static let height: CGFloat = 2700

    @Test("every round-5 surface, button, primitive, the paperdoll header and the tab bar")
    func sheet() {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = .large
        }
        let portrait = UIImage(contentsOfFile: HomeFixtures.repoRoot.appending(path: "design/mocks/data/avatars/thandrel.jpg").path)
        let view = DesignSystemSheet(build: BundledBuild.read() ?? "", portrait: portrait)
            .frame(width: Self.width, height: Self.height, alignment: .top)
            .environment(\.itemIconSource, RepoItemIconSource())
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: view,
            as: .keyWindowImage(layout: .fixed(width: Self.width, height: Self.height), traits: traits),
            testName: "design-system-sheet-ios"
        )
    }
}
