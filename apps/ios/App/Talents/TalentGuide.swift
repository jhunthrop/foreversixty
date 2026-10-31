import Foundation
import Models

/// Which guide build the screen compares with, and whether an export can be drawn on the active
/// build's trees (spec F6 §4.2, §4.3).
enum TalentGuide {
    static let firstBand = 20
    static let lastBand = 60
    private static let bandStep = 10
    private static let raidPreset = "raid"
    private static let barePreset = "bare"

    /// The band of a level: 60 from 60, `floor(level / 10) * 10` from 20 to 59, none under 20.
    static func band(forLevel level: Int?) -> Int? {
        guard let level, level >= firstBand else { return nil }
        return min(lastBand, level / bandStep * bandStep)
    }

    /// The planner's `selectBand`: the band's raid entry for the faction, else bare, else the first.
    static func entry(in file: BisFile, band: Int, faction: String?) -> BisBand? {
        let entries = file.bands.filter { $0.band == band && $0.faction == faction }
        return entries.first { $0.preset == raidPreset } ?? entries.first { $0.preset == barePreset } ?? entries.first
    }

    /// The band's copy label: `60`, or `50 to 59` (the planner's `bandLabel`).
    static func label(band: Int) -> String {
        band >= lastBand ? Copy.bisRangeTop(n: band) : Copy.bisRangeSpan(from: band, to: band + bandStep - 1)
    }

    /// True when an export made on `older` can be read on `active`'s trees: every tree lists the
    /// same talent ids in the same order, so each digit lands on the same talent.
    static func sameLayout(_ older: TalentsFile, _ active: TalentsFile) -> Bool {
        layout(older) == layout(active)
    }

    private static func layout(_ file: TalentsFile) -> [[Int]] {
        TalentRules.trees(file).map { $0.talents.map(\.id) }
    }
}
