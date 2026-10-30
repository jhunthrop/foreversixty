@testable import ForeverSixty
import SnapshotTesting
import SwiftUI
import Testing

/// Reference images are recorded on the iPhone 17 simulator (iOS 27.1, Xcode 27.1)
/// in __Snapshots__. Re-record after an intended change with
/// `SNAPSHOT_TESTING_RECORD=all` in the test environment, and review the PNG diff.
@MainActor
@Suite("Home snapshot")
struct HomeSnapshotTests {
    @Test("placeholder home with a loaded character")
    func loadedHome() {
        let view = HomeView(state: .loaded(HomeCard(
            name: "Obnoxious",
            detail: "Level 42 · Arms Warrior",
            guild: "Fixture Guild"
        )))
        .preferredColorScheme(.dark)

        assertSnapshot(
            of: view,
            as: .image(
                precision: 0.99,
                perceptualPrecision: 0.98,
                layout: .device(config: .iPhone13),
                traits: UITraitCollection(userInterfaceStyle: .dark)
            )
        )
    }
}
