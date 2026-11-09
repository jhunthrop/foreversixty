import Testing
import UIKit

/// The key-window snapshots' CPU compare reads past GPU anti-aliasing noise and nothing more.
@Suite("Snapshot tolerance")
struct SnapshotToleranceTests {
    private func image(_ level: CGFloat, size: CGSize = CGSize(width: 8, height: 8)) -> UIImage {
        let format = UIGraphicsImageRendererFormat()
        format.scale = 1
        return UIGraphicsImageRenderer(size: size, format: format).image { context in
            UIColor(white: level / 255, alpha: 1).setFill()
            context.fill(CGRect(origin: .zero, size: size))
        }
    }

    @Test func aFewLevelsAreNoise() {
        #expect(SnapshotTolerance.maxChannelDelta == 5)
        #expect(SnapshotTolerance.withinChannelDelta(image(100), image(103)))
    }

    @Test func aRealChangeIsNot() {
        #expect(!SnapshotTolerance.withinChannelDelta(image(100), image(140)))
        #expect(!SnapshotTolerance.withinChannelDelta(image(100), image(100, size: CGSize(width: 9, height: 8))))
    }
}
