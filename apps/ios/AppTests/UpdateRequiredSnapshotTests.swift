@testable import ForeverSixty
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// The blocking update screen; recorded and re-recorded like HomeSnapshotTests.
@MainActor
@Suite("Update required")
struct UpdateRequiredSnapshotTests {
    @Test("the blocking update screen")
    func updateScreen() {
        let view = UpdateRequiredView(storeURL: UpdateGate.fallbackStoreURL).preferredColorScheme(.dark)
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

    @Test("the gate keeps the store link it was sent, and the fallback when it was sent none")
    func gateStoreURL() {
        let gate = UpdateGate()
        #expect(!gate.isRequired)
        gate.require(storeURL: nil)
        #expect(gate.isRequired)
        #expect(gate.storeURL == UpdateGate.fallbackStoreURL)
        let store = URL(string: "https://apps.apple.com/app/id1")!
        gate.require(storeURL: store)
        #expect(gate.storeURL == store)
    }
}
