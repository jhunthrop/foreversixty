import Auth
import Foundation
import Models

/// The answer card (spec F11.1): your rank with its colour and words, or not ranked yet.
enum RankAnswer: Equatable, Sendable {
    /// `name` is the ranked character's name when it is not the current character (F11.3).
    case ranked(ordinal: String, tier: ParseTier, unit: String, meta: String, name: String?)
    case unranked(title: String, meta: String)
}

/// One phase chip: the API's id and its words.
struct PhaseChoice: Equatable, Sendable, Identifiable {
    let id: String
    let label: String
}

/// Every F11 board sentence built from data (spec §5.3); the words live in the copy catalog.
enum RankingsCopy {
    private static let separator = " · "
    private static let percent = 100.0
    private static let nonBreaking = "\u{00A0}"

    static let regions = ["us", "eu", "kr", "tw", "cn"]
    static let rulesets = ["normal", "pvp", "rp", "hardcore"]
    static let factions = ["alliance", "horde"]

    static func metric(_ metric: RankMetric) -> String {
        switch metric {
        case .dps: Copy.rankMeasureDamage
        case .hps: Copy.rankMeasureHealing
        case .damageTaken: Copy.rankMeasureTaken
        case .execution: Copy.rankMeasureExecution
        }
    }

    /// The word under a row's value.
    static func unit(_ metric: RankMetric) -> String {
        switch metric {
        case .dps: Copy.rankUnitDps
        case .hps: Copy.rankUnitHps
        case .damageTaken: Copy.rankUnitDtps
        case .execution: Copy.rankUnitExecution
        }
    }

    /// The API's per-second unit word (`dps`, `hps`, `dtps`, `damage_taken`) on a character's fights.
    static func fightUnit(_ metric: String) -> String {
        switch metric {
        case "hps": Copy.rankUnitHps
        case "dtps", RankMetric.damageTaken.rawValue: Copy.rankUnitDtps
        case RankMetric.execution.rawValue: Copy.rankUnitExecution
        default: Copy.rankUnitDps
        }
    }

    static func spokenUnit(_ metric: RankMetric) -> String {
        switch metric {
        case .dps: Copy.rankA11yPerSecondDps
        case .hps: Copy.rankA11yPerSecondHps
        case .damageTaken: Copy.rankA11yPerSecondDtps
        case .execution: Copy.rankA11yExecuted
        }
    }

    /// A ruleset's words; empty is every ruleset.
    static func ruleset(_ id: String) -> String {
        switch id {
        case "": Copy.rankRulesetAll
        case "normal": Copy.rankRulesetNormal
        case "pvp": Copy.rankRulesetPvp
        case "rp": Copy.rankRulesetRp
        case "hardcore": Copy.rankRulesetHardcore
        default: id
        }
    }

    static func faction(_ id: String) -> String {
        switch id {
        case "alliance": Copy.rankFactionAlliance
        case "horde": Copy.rankFactionHorde
        default: Copy.rankFactionBoth
        }
    }

    /// A phase's words (phases.ts `phaseLabel`); an id the app does not know prints as sent.
    static func phase(_ id: String) -> String {
        switch id {
        case "pre-beta": Copy.rankPhasePreBeta
        case "beta": Copy.rankPhaseBeta
        case "launch": Copy.rankPhaseLaunch
        case "raids-1": Copy.rankPhaseRaids1
        default: id
        }
    }

    /// The phase chips, newest first.
    static func phaseChoices(_ phases: [PhaseBoundary]) -> [PhaseChoice] {
        phases.sorted { $0.start > $1.start }.map { PhaseChoice(id: $0.name, label: phase($0.name)) }
    }

    /// `Shaman` from `shaman`.
    static func className(_ slug: String) -> String {
        slug.prefix(1).uppercased() + slug.dropFirst().lowercased()
    }

    /// The filter line (spec F11.1): `Onyxia · Normal · All regions · Phase 1 · Damage`, then a
    /// faction and `Today only` when chosen.
    static func filterLine(_ filters: RankingsFilters, encounterName: String) -> String {
        [
            encounterName,
            ruleset(filters.ruleset),
            filters.region.isEmpty ? Copy.rankLineAllRegions : filters.region.uppercased(),
            filters.phase.isEmpty ? "" : phase(filters.phase),
            metric(filters.metric),
            filters.faction.isEmpty ? "" : faction(filters.faction),
            filters.todayOnly ? Copy.rankFilterToday : "",
        ]
        .filter { !$0.isEmpty }
        .joined(separator: separator)
    }

    /// A row's value: per-second figures whole and grouped; execution as a percent.
    static func value(_ row: RankingRow, metric: RankMetric) -> String {
        guard metric == .execution else { return ReportsCopy.figure(row.value) }
        let score = row.executionScore ?? row.value
        return "\(Int((score * percent).rounded()))%"
    }

    /// `Hordebreakers · 40 players · 3:13 · 2 d ago`.
    static func rowMeta(_ row: RankingRow, now: Date) -> String {
        let time = ReportsCopy.duration(Int64(row.durationMs))
        // The age never breaks inside itself (`4 d` / `ago`) when the meta wraps.
        let age = RelativeAge(from: row.foughtAt, to: now).phrase.replacingOccurrences(of: " ", with: nonBreaking)
        guard let guild = row.guild?.name else { return Copy.rankRowMetaNoGuild(size: row.size, time: time, age: age) }
        return Copy.rankRowMeta(guild: guild, size: row.size, time: time, age: age)
    }

    /// `Rank 14 of 212, Lightningbolltz, 398 damage per second, Hordebreakers, your character`.
    static func a11y(_ row: RankingRow, total: Int, metric: RankMetric, yours: Bool) -> String {
        let base = Copy.rankA11yRow(
            rank: row.rank, total: total, name: row.player.name, value: value(row, metric: metric),
            unit: spokenUnit(metric), guild: row.guild?.name ?? ""
        )
        return yours ? base + ", " + Copy.rankA11yYours : base
    }

    /// The answer card; nil signed out (no current character).
    static func answer(page: RankingsPage, filters: RankingsFilters, encounterName: String, current: RankingsCharacter?) -> RankAnswer? {
        guard let current else { return nil }
        guard let placed = page.characterRank else {
            return .unranked(title: Copy.rankNone(encounter: encounterName), meta: Copy.rankNoneMeta)
        }
        let named = placed.row.player.name == current.name ? nil : placed.row.player.name
        let rulesetWords = ruleset(filters.ruleset)
        let phaseWords = phase(filters.phase)
        let meta = named.map {
            Copy.rankMetaNamed(name: $0, encounter: encounterName, ruleset: rulesetWords, phase: phaseWords)
        } ?? Copy.rankMeta(encounter: encounterName, ruleset: rulesetWords, phase: phaseWords)
        return .ranked(
            ordinal: ReportsCopy.ordinal(placed.rank), tier: RankColor.tier(rank: placed.rank, total: page.total),
            unit: unit(page: page, filters: filters), meta: meta, name: named
        )
    }

    /// `of 212 Elemental Shaman characters`, or `kills` when the API did not dedupe.
    static func unit(page: RankingsPage, filters: RankingsFilters) -> String {
        let className = className(filters.classSlug)
        return page.deduped
            ? Copy.rankUnit(total: page.total, spec: filters.spec, className: className)
            : Copy.rankUnitKills(total: page.total, spec: filters.spec, className: className)
    }

    /// Home's Rankings row sub (spec F11.1).
    static func homeLine(page: RankingsPage, filters: RankingsFilters, encounterName: String) -> String {
        guard let placed = page.characterRank else { return Copy.homeRankingsNone }
        return Copy.homeRankingsLine(
            ordinal: ReportsCopy.ordinal(placed.rank), total: page.total, spec: filters.spec,
            className: className(filters.classSlug), encounter: encounterName
        )
    }

    /// `Onyxia's Lair · 40 players · 38 ranked kills`.
    static func encounterMeta(_ encounter: RankedEncounter) -> String {
        let kills = Copy.rankKills(n: encounter.kills)
        guard let zone = encounter.zone, let size = encounter.size else { return kills }
        return Copy.rankEncounterMeta(zone: zone, size: size, kills: kills)
    }
}
