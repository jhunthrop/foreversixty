import DesignSystem
import Foundation
import SwiftUI
import UIKit

/// The native parity gate's region side (design/parity.md "Native"): render a screen at the
/// device size, collect the regions and inventory its primitives report (`parityRegion`,
/// `parityItem`), and compare them with the mock's region file
/// (design/mocks/app/regions/ios-<screen>.json) on the web gate's rules.
enum ParityRegions {
    /// x and y within 8 pt; width and height within 8 pt or 5%, whichever is larger.
    static let positionTolerance: CGFloat = 8
    static let sizeTolerance: CGFloat = 8
    static let sizeRatioTolerance: CGFloat = 0.05

    static let regionsDirectory = HomeFixtures.repoRoot.appending(path: "design/mocks/app/regions")

    /// One region file: the device and, per region, its box `[x, y, w, h]` and inventory.
    struct Reference: Decodable, Equatable {
        struct Device: Decodable, Equatable {
            let width: CGFloat
            let height: CGFloat
        }

        struct Region: Decodable, Equatable {
            let box: [CGFloat]
            let inventory: [String: Int]
        }

        let device: Device
        let regions: [String: Region]
    }

    /// What a render reported.
    struct Capture: Equatable {
        var regions: [String: CGRect] = [:]
        var items: [ParityItem] = []
    }

    static func reference(platform: String = "ios", screen: String) throws -> Reference {
        let url = regionsDirectory.appending(path: "\(platform)-\(screen).json")
        return try JSONDecoder().decode(Reference.self, from: Data(contentsOf: url))
    }

    /// Renders `view` at `size` in a window and reads what it reports. The view must call
    /// `.parityContent()` on its content.
    @MainActor
    static func capture(of view: some View, size: CGSize) -> Capture {
        let box = CaptureBox()
        let reporting = view.onParityCapture(regions: { box.set(regions: $0) }, items: { box.set(items: $0) })
        return ParityHost.render(reporting, size: size) { _ in box.capture }
    }

    /// Every way the capture breaks the reference; empty means it holds.
    static func compare(_ capture: Capture, with reference: Reference) -> [String] {
        var diffs: [String] = []
        for name in Set(reference.regions.keys).subtracting(capture.regions.keys).sorted() {
            diffs.append("\(name): in the mock, missing from the build")
        }
        for name in Set(capture.regions.keys).subtracting(reference.regions.keys).sorted() {
            diffs.append("\(name): in the build, not in the mock")
        }
        for (name, want) in reference.regions.sorted(by: { $0.key < $1.key }) {
            guard let got = capture.regions[name], want.box.count == 4 else { continue }
            diffs += boxDiffs(name, got: got, want: want.box)
            let counted = inventory(of: got, items: capture.items)
            if counted != want.inventory {
                diffs.append("\(name): inventory \(describe(counted)), the mock has \(describe(want.inventory))")
            }
        }
        return diffs
    }

    /// The items whose centre lies inside `region`, counted by kind (kinds with none left out).
    static func inventory(of region: CGRect, items: [ParityItem]) -> [String: Int] {
        var counts: [String: Int] = [:]
        for item in items where region.contains(CGPoint(x: item.frame.midX, y: item.frame.midY)) {
            counts[item.kind.rawValue, default: 0] += 1
        }
        return counts
    }

    private static func boxDiffs(_ name: String, got: CGRect, want: [CGFloat]) -> [String] {
        let checks: [(String, CGFloat, CGFloat, CGFloat)] = [
            ("x", got.minX, want[0], positionTolerance),
            ("y", got.minY, want[1], positionTolerance),
            ("width", got.width, want[2], max(sizeTolerance, want[2] * sizeRatioTolerance)),
            ("height", got.height, want[3], max(sizeTolerance, want[3] * sizeRatioTolerance)),
        ]
        return checks.compactMap { field, value, target, tolerance in
            abs(value - target) <= tolerance
                ? nil
                : "\(name): \(field) \(Int(value.rounded())), the mock has \(Int(target.rounded())) (within \(Int(tolerance.rounded())))"
        }
    }

    private static func describe(_ counts: [String: Int]) -> String {
        counts.isEmpty ? "{}" : counts.sorted { $0.key < $1.key }.map { "\($0.key) \($0.value)" }.joined(separator: ", ")
    }
}

/// Collects the preference callbacks, which may arrive off the main actor.
private final class CaptureBox: @unchecked Sendable {
    private let lock = NSLock()
    private var value = ParityRegions.Capture()

    var capture: ParityRegions.Capture {
        lock.withLock { value }
    }

    func set(regions: [String: CGRect]) {
        lock.withLock { value.regions = regions }
    }

    func set(items: [ParityItem]) {
        lock.withLock { value.items = items }
    }
}
