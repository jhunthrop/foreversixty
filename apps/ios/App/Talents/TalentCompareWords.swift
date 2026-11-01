import Foundation

/// What F6's compare measures your build against, and so the words it uses (lane F15, spec
/// F13.5): the guide build of a band (`talents.*`) or a shared build (`plan.cmp.*`). One compare
/// draws both; only these words change.
enum TalentCompareWords: Equatable, Sendable {
    case guide(spec: String, band: String)
    case sharedBuild

    /// The segment over the other side: `Guide build` or `Shared build`.
    var side: String {
        switch self {
        case .guide: Copy.talentsGuide
        case .sharedBuild: Copy.planCmpSide
        }
    }

    var yours: String {
        switch self {
        case .guide: Copy.talentsYours
        case .sharedBuild: Copy.planCmpYours
        }
    }

    func match(beyond: Bool) -> String {
        switch (self, beyond) {
        case let (.guide(spec, band), true): Copy.talentsMatchBeyond(spec: spec, band: band)
        case let (.guide(spec, band), false): Copy.talentsMatch(spec: spec, band: band)
        case (.sharedBuild, true): Copy.planCmpMatchBeyond
        case (.sharedBuild, false): Copy.planCmpMatch
        }
    }

    func differ(n: Int, of total: Int) -> String {
        switch self {
        case let .guide(_, band): Copy.talentsDiffer(n: n, m: total, band: band)
        case .sharedBuild: Copy.planCmpDiffer(n: n, m: total)
        }
    }

    func toTake(_ n: Int) -> String {
        switch self {
        case let .guide(_, band): Copy.talentsToTake(n: n, band: band)
        case .sharedBuild: Copy.planCmpToTake(n: n)
        }
    }

    func missing(_ n: Int) -> String {
        switch self {
        case .guide: Copy.talentsMissing(n: n)
        case .sharedBuild: Copy.planCmpMissing(n: n)
        }
    }

    func beyond(_ n: Int) -> String {
        switch self {
        case .guide: Copy.talentsBeyond(n: n)
        case .sharedBuild: Copy.planCmpBeyond(n: n)
        }
    }

    var legend: String {
        switch self {
        case .guide: Copy.talentsLegend
        case .sharedBuild: Copy.planCmpLegend
        }
    }

    func otherRank(_ g: Int, max: Int) -> String {
        switch self {
        case .guide: Copy.talentsGuideRank(g: g, max: max)
        case .sharedBuild: Copy.planCmpRank(g: g, max: max)
        }
    }

    func tabValue(points: Int, differing: Int) -> String {
        switch self {
        case .guide: Copy.talentsA11yTabDiffers(n: points, k: differing)
        case .sharedBuild: Copy.planCmpA11yTab(n: points, k: differing)
        }
    }

    func cellLabel(name: String, rank: Int, max: Int, other: Int) -> String {
        switch self {
        case .guide: Copy.talentsA11yCellGuide(name: name, n: rank, max: max, g: other)
        case .sharedBuild: Copy.planCmpA11yCell(name: name, n: rank, max: max, g: other)
        }
    }
}
