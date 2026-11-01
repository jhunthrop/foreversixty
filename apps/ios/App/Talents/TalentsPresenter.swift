import Foundation
import Models

/// Which build the tree draws (spec F6 §5.1 segments).
enum TalentBuildKind: Equatable, Sendable, CaseIterable {
    case guide, yours
}

/// Your export, as the screen can use it (spec F6 §4.3, §6).
enum TalentExport: Equatable, Sendable {
    /// Ranks per tree, read on the active build's trees.
    case ranks([[Int]])
    case none
    /// Made on a build whose trees differ from the active build's.
    case stale
}

/// The guide build, as the screen can use it (spec F6 §4.2, §6).
enum TalentGuideBuild: Equatable, Sendable {
    case build(talents: String, band: Int)
    /// No band, no file or no entry: `talents.noList`.
    case none
    /// The character has no spec, so there is no BiS file to read.
    case noSpec
    /// The BiS file failed to load.
    case failed
    /// Lane F15 (spec F13.5 Compare with mine): a shared build's ranks per tree in the guide's
    /// place, read on the active trees.
    case shared(treeRanks: [[Int]])
}

/// The tag kinds the status line uses (spec F6 §5.1 item 3).
enum TalentTagKind: Equatable, Sendable {
    case matches, differs, muted
}

/// The status block: a tag (or none) and the sentence; `isError` draws it as the failed line.
struct TalentStatusLine: Equatable, Sendable {
    let tag: String?
    let tagKind: TalentTagKind
    let sentence: String
    var isError = false

    var accessibilityLabel: String { [tag, sentence].compactMap(\.self).joined(separator: ". ") }
}

/// One tree tab: `{Tree} {points}`, plus the dot for a tree holding a dashed cell.
struct TalentTabModel: Equatable, Sendable, Identifiable {
    let index: Int
    let name: String
    let points: Int
    let differing: Int
    /// What the compare is against; nil reads as the guide build.
    var words: TalentCompareWords?

    var id: Int { index }
    var accessibilityValue: String {
        guard differing > 0 else { return Copy.talentsPoints(n: points) }
        return words?.tabValue(points: points, differing: differing) ?? Copy.talentsA11yTabDiffers(n: points, k: differing)
    }
}

/// One drawn cell: the talent, its rank and state, and on Your build its signed difference.
struct TalentCellModel: Equatable, Sendable, Identifiable {
    let talent: Talent
    let rank: Int
    let state: TalentCellState
    /// Your rank minus the guide's, on Your build when they differ; nil otherwise.
    let delta: Int?
    let guideRank: Int?
    /// What the compare is against; nil reads as the guide build.
    var words: TalentCompareWords?

    var id: Int { talent.id }
    var badge: String { "\(rank)/\(talent.maxRank)" }
    var chip: String? { delta.map { $0 > 0 ? "+\($0)" : "\u{2212}\(-$0)" } }
    var accessibilityLabel: String {
        guard let guideRank, delta != nil else { return Copy.talentsA11yCell(name: talent.name, n: rank, max: talent.maxRank) }
        return (words ?? .guide(spec: "", band: "")).cellLabel(name: talent.name, rank: rank, max: talent.maxRank, other: guideRank)
    }
}

/// The rank sheet (spec F6 §7).
struct TalentSheetModel: Equatable, Sendable, Identifiable {
    let talent: Talent
    let rank: Int
    let state: TalentCellState
    let current: String?
    let next: String?
    let requirements: [String]
    let guideLine: String?
    /// True when your rank is over the guide's: the guide line's ring is wipe, else gold.
    let isOver: Bool

    var id: Int { talent.id }
    var rankLine: String { Copy.talentsRank(n: rank, max: talent.maxRank) }
}

/// Everything the Talents screen shows, pure over the files and the character (spec F6). The
/// counts, dashes and totals all come from one pass over the drawn ranks.
struct TalentsPresenter: Equatable, Sendable {
    let file: TalentsFile
    let trees: [TalentTree]
    let guideRanks: [Int: Int]?
    let yourRanks: [Int: Int]?
    let guide: TalentGuideBuild
    let export: TalentExport
    let level: Int?
    let specName: String

    init(file: TalentsFile, guide: TalentGuideBuild, export: TalentExport, level: Int?, specName: String) {
        self.file = file
        trees = TalentRules.trees(file)
        self.guide = guide
        self.export = export
        self.level = level
        self.specName = specName
        switch guide {
        case let .build(talents, _): guideRanks = TalentRules.rankByID(file, treeRanks: TalentRules.ranks(ofBand: talents))
        case let .shared(treeRanks): guideRanks = TalentRules.rankByID(file, treeRanks: treeRanks)
        default: guideRanks = nil
        }
        if case let .ranks(ranks) = export {
            yourRanks = TalentRules.rankByID(file, treeRanks: ranks)
        } else {
            yourRanks = nil
        }
    }

    // MARK: The compare

    var diff: TalentDiff? {
        guard let guideRanks, let yourRanks else { return nil }
        return TalentRules.compare(file, yours: yourRanks, guide: guideRanks)
    }

    var bandLabel: String? {
        guard case let .build(_, band) = guide else { return nil }
        return TalentGuide.label(band: band)
    }

    /// The compare's words: against the band's guide build, or a shared build.
    var words: TalentCompareWords? {
        switch guide {
        case let .build(_, band): .guide(spec: specName, band: TalentGuide.label(band: band))
        case .shared: .sharedBuild
        default: nil
        }
    }

    var yourTotal: Int? { yourRanks.map { $0.values.reduce(0, +) } }

    /// The builds that can be drawn; the segments show only when both can.
    var builds: [TalentBuildKind] {
        [guideRanks == nil ? nil : TalentBuildKind.guide, yourRanks == nil ? nil : .yours].compactMap(\.self)
    }

    var showsSegments: Bool { builds.count == 2 }

    /// Your build when it can be drawn, else the guide (spec F6 §5.4).
    var initialBuild: TalentBuildKind { yourRanks == nil ? .guide : .yours }

    /// Nothing to draw: no export and no guide.
    var hasTree: Bool { !builds.isEmpty }

    func ranks(_ build: TalentBuildKind) -> [Int: Int] {
        (build == .yours ? yourRanks : guideRanks) ?? yourRanks ?? guideRanks ?? [:]
    }

    // MARK: Status and notes

    var status: TalentStatusLine {
        switch (export, guide) {
        case (.ranks, .failed):
            return TalentStatusLine(tag: nil, tagKind: .muted, sentence: Copy.talentsGuideFailed, isError: true)
        case (.ranks, .noSpec):
            return TalentStatusLine(tag: Copy.talentsYours, tagKind: .muted, sentence: Copy.talentsPickSpec)
        case (.ranks, .none):
            return TalentStatusLine(tag: Copy.talentsYours, tagKind: .muted, sentence: Copy.talentsNoList)
        case (.ranks, .build), (.ranks, .shared):
            return compareStatus
        case (.stale, _):
            return TalentStatusLine(tag: guideRanks == nil ? nil : words?.side, tagKind: .muted, sentence: Copy.talentsStaleExport)
        case (.none, _):
            return TalentStatusLine(tag: guideRanks == nil ? nil : words?.side, tagKind: .muted, sentence: Copy.homeNextTalentsNone)
        }
    }

    private var compareStatus: TalentStatusLine {
        guard let diff, let words else {
            return TalentStatusLine(tag: Copy.talentsYours, tagKind: .muted, sentence: Copy.talentsNoList)
        }
        switch diff.status {
        case .matches:
            return TalentStatusLine(tag: Copy.talentsMatches, tagKind: .matches, sentence: words.match(beyond: diff.extraCount > 0))
        case let .differs(count):
            return TalentStatusLine(tag: Copy.talentsDiffers, tagKind: .differs, sentence: words.differ(n: count, of: yourTotal ?? 0))
        case let .toTake(count):
            return TalentStatusLine(tag: Copy.talentsDiffers, tagKind: .differs, sentence: words.toTake(count))
        }
    }

    /// The secondary lines under the status, in order (spec F6 §5.1 item 4).
    var notes: [String] {
        var lines: [String] = []
        if let yourTotal, let level {
            let unspent = max(TalentRules.pointsAvailable(atLevel: level) - yourTotal, 0)
            if unspent > 0 { lines.append(Copy.talentsUnspent(n: unspent)) }
        }
        if let diff, let words, diff.diffCount > 0, diff.missingCount > 0 {
            lines.append(words.missing(diff.missingCount))
        }
        if let diff, let words, diff.extraCount > 0 {
            lines.append(words.beyond(diff.extraCount))
        }
        // No export and no guide either: say both, and draw no tree (spec F6 §6, No export).
        if !hasTree { lines.append(Copy.talentsNoList) }
        return lines
    }

    // MARK: Tabs and the tree

    func tabs(_ build: TalentBuildKind) -> [TalentTabModel] {
        let ranks = ranks(build)
        let points = TalentRules.pointsPerTree(file, ranks: ranks)
        let differing = differingIDs(build)
        return trees.enumerated().map { index, tree in
            TalentTabModel(
                index: index, name: tree.name, points: points[index],
                differing: tree.talents.filter { differing.contains($0.id) }.count, words: words
            )
        }
    }

    /// The tree that opens: the most points in the build first shown (spec F6 §5.3).
    var openingTree: Int {
        TalentRules.openingTree(TalentRules.pointsPerTree(file, ranks: ranks(initialBuild)))
    }

    func cells(_ build: TalentBuildKind, tree index: Int) -> [TalentCellModel] {
        guard trees.indices.contains(index) else { return [] }
        let tree = trees[index]
        let ranks = ranks(build)
        let marks = build == .yours ? guideRanks : nil
        return tree.talents.map { talent in
            let rank = ranks[talent.id] ?? 0
            let guideRank = marks?[talent.id]
            let delta = guideRank.flatMap { rank == $0 ? nil : rank - $0 }
            return TalentCellModel(
                talent: talent, rank: rank, state: TalentRules.cellState(talent, in: tree, ranks: ranks),
                delta: delta, guideRank: guideRank, words: words
            )
        }
    }

    func connectors(_ build: TalentBuildKind, tree index: Int) -> [TalentConnector] {
        guard trees.indices.contains(index) else { return [] }
        return TalentRules.connectors(trees[index], ranks: ranks(build))
    }

    func legend(_ build: TalentBuildKind) -> String {
        differingIDs(build).isEmpty ? Copy.talentsTapHint : words?.legend ?? Copy.talentsLegend
    }

    private func differingIDs(_ build: TalentBuildKind) -> Set<Int> {
        build == .yours ? diff?.differing ?? [] : []
    }

    // MARK: The sheet

    func sheet(_ build: TalentBuildKind, talentID: Int) -> TalentSheetModel? {
        guard let tree = trees.first(where: { $0.talents.contains { $0.id == talentID } }),
              let talent = tree.talents.first(where: { $0.id == talentID })
        else { return nil }
        let ranks = ranks(build)
        let rank = ranks[talent.id] ?? 0
        let description = { (index: Int) in talent.ranks.indices.contains(index) ? talent.ranks[index].description : nil }
        let guideRank = build == .yours ? guideRanks?[talent.id] : nil
        let differs = guideRank.map { $0 != rank } ?? false
        return TalentSheetModel(
            talent: talent,
            rank: rank,
            state: TalentRules.cellState(talent, in: tree, ranks: ranks),
            current: rank > 0 ? description(rank - 1) : nil,
            next: rank < talent.maxRank ? description(rank) : nil,
            requirements: Self.requirements(talent, in: tree, ranks: ranks),
            guideLine: differs ? guideRank.map { (words ?? .guide(spec: "", band: "")).otherRank($0, max: talent.maxRank) } : nil,
            isOver: differs && rank > (guideRank ?? 0)
        )
    }

    /// The game's requirement lines, only when unmet in the drawn build (spec F6 §7).
    static func requirements(_ talent: Talent, in tree: TalentTree, ranks: [Int: Int]) -> [String] {
        var lines: [String] = []
        let needed = TalentRules.pointsPerTier * talent.tier
        let inTree = tree.talents.reduce(0) { $0 + (ranks[$1.id] ?? 0) }
        if inTree < needed { lines.append(Copy.talentsRequiresTree(n: needed, tree: tree.name)) }
        if !TalentRules.prerequisiteMet(talent, ranks: ranks), let prereqID = talent.prereqTalentId,
           let prereq = tree.talents.first(where: { $0.id == prereqID }) {
            let points = talent.prereqRank ?? 1
            lines.append(points == 1 ? Copy.talentsRequiresTalentOne(talent: prereq.name)
                : Copy.talentsRequiresTalent(n: points, talent: prereq.name))
        }
        return lines
    }
}
