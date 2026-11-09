import SnapshotTesting
import SwiftUI
import UIKit

extension Snapshotting where Value: SwiftUI.View, Format == UIImage {
    /// A snapshot drawn through the key window (UIKit controls and iOS 26 glass bar items render),
    /// compared on the CPU only, at the suites' tolerance (SnapshotTolerance.maxChannelDelta).
    ///
    /// Drawing through the window goes through the host Mac's GPU, so a lane's Mac and the CI runner
    /// can differ by a level or three on an anti-aliased pixel (a crest's edge, 2026-11-08: 7,054
    /// pixels, at most 3 of 255). swift-snapshot-testing's perceptual compare would absorb that, but
    /// it is a CoreImage/Metal pass that, on the runner's virtual GPU, reports nonsense (5% of pixels
    /// and a 0.45 precision for a one-pixel, one-level difference; 2026-11-02, Xcode 27.1 and iOS
    /// 26.5 on both sides). So the bytes are compared here instead: every channel of every pixel
    /// within the delta the suites' perceptual precision allows, and anything beyond it falls through
    /// to the library's own report (its reference, failure and difference attachments).
    static func keyWindowImage(layout: SwiftUISnapshotLayout, traits: UITraitCollection) -> Snapshotting {
        var strategy = Snapshotting.image(
            drawHierarchyInKeyWindow: true, precision: SnapshotTolerance.precision, perceptualPrecision: 1,
            layout: layout, traits: traits
        )
        let exact = strategy.diffing.diffV2
        // The exact compare first: it is the library's own fast path, and the per-channel read runs
        // only on a miss.
        strategy.diffing.diffV2 = { reference, taken in
            guard let failure = exact(reference, taken) else { return nil }
            return SnapshotTolerance.withinChannelDelta(reference, taken) ? nil : failure
        }
        return strategy
    }
}
