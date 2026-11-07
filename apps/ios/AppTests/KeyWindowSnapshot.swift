import SnapshotTesting
import SwiftUI

extension Snapshotting where Value: SwiftUI.View, Format == UIImage {
    /// A snapshot drawn through the key window (UIKit controls and iOS 26 glass bar items render),
    /// compared on the CPU only.
    ///
    /// Drawing through the window goes through the host Mac's GPU, so a lane's Mac and the CI runner
    /// can differ by a level on a pixel. Any byte difference sends swift-snapshot-testing to its
    /// perceptual compare, a CoreImage/Metal pass that runs on that same GPU and, on the runner's
    /// virtual GPU, reports nonsense (5% of pixels and a 0.45 precision for a one-pixel, one-level
    /// difference; 2026-11-02, Xcode 27.1 and iOS 26.5 on both sides). With `perceptualPrecision`
    /// at 1 the compare counts differing bytes instead, at the suites' precision (SnapshotTolerance).
    static func keyWindowImage(layout: SwiftUISnapshotLayout, traits: UITraitCollection) -> Snapshotting {
        .image(drawHierarchyInKeyWindow: true, precision: SnapshotTolerance.precision, perceptualPrecision: 1, layout: layout, traits: traits)
    }
}
