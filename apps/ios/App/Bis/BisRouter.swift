import Foundation
import Observation
import Routing

/// A request for the BiS tab to show one list (a guide's Gear row, a `/app/bis` link).
struct BisRequest: Equatable, Sendable {
    let spec: String
    let range: BisRange
    /// Each request is new, so the same row tapped twice applies twice.
    let id = UUID()
}

/// The BiS tab's outside requests, applied by `BisScreen` over its own model.
@MainActor
@Observable
final class BisRouter {
    private(set) var request: BisRequest?
    static let raidLinkBand = "p1"

    /// A guide's Gear row: that spec at the band with the raid preset.
    func openRaid(spec: String, band: Int) {
        request = BisRequest(spec: spec, range: BisRange(band: band, preset: BisRange.raidPreset))
    }

    static func handles(_ link: DeepLink) -> Bool {
        if case .bis = link { return true }
        return false
    }

    /// `/app/bis/{spec}/{band}`: `p1` is the Phase 1 raid list, a number its bare band.
    func open(_ link: DeepLink) {
        guard case let .bis(spec, band) = link else { return }
        if band == Self.raidLinkBand {
            openRaid(spec: spec, band: BisRange.lastBand)
        } else if let level = Int(band) {
            request = BisRequest(spec: spec, range: BisRange(band: BisRange.forLevel(level).band, preset: BisRange.barePreset))
        }
    }
}
