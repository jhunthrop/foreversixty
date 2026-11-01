import Auth
import Foundation
import Routing
import Models

/// One Forever Score component bar (spec F11.4): its length is the value over the API's cap.
struct ScoreBar: Equatable, Sendable, Identifiable {
    let label: String
    let value: Int
    let cap: Int

    var id: String { label }

    var fraction: Double {
        guard cap > 0 else { return 0 }
        return min(1, max(0, Double(value) / Double(cap)))
    }

    var a11y: String { Copy.charA11yComponent(component: label, value: value, cap: cap) }
}

/// A Bests or Recent fights row.
struct FightLine: Equatable, Sendable, Identifiable {
    let id: String
    let encounter: String
    let meta: String
    let value: String
    let unit: String
    let percentile: Int?
    let kill: Bool
    let reportID: String
}

/// A Builds seen row.
struct BuildLine: Equatable, Sendable {
    let trees: String
    let meta: String
}

/// The character page's words and figures (spec F11.4 to F11.7), from `CharacterProfile`.
enum CharacterPage {
    /// The planner on the character's class (`/app/planner/{class}`): the class only, since
    /// rankings carry no build code or race. Nil when the class is not known.
    static func plannerLink(_ profile: CharacterProfile) -> DeepLink? {
        guard let slug = profile.character.className?.lowercased(), !slug.isEmpty else { return nil }
        return .planner(className: slug)
    }

    private static let separator = " · "
    private static let treeSeparator = " / "

    /// The page's spec: the gear spec (`shaman-elemental`) by its catalog name, else the spec of
    /// the newest best or fight.
    static func spec(_ profile: CharacterProfile, specs: SpecCatalog) -> String? {
        if let key = profile.gearSpec, let known = specs.spec(key) { return known.name }
        return (profile.best + profile.history).compactMap(\.spec).first
    }

    /// `Orc Elemental Shaman · Level 60`; the parts the API has, in that order.
    static func descriptor(_ profile: CharacterProfile, specs: SpecCatalog) -> String {
        let header = profile.character
        let race = header.race.map(HomePresenter.words) ?? ""
        let spec = spec(profile, specs: specs) ?? ""
        let className = header.className.map(RankingsCopy.className) ?? ""
        if let level = header.level {
            return Copy.charDescriptor(race: race, spec: spec, className: className, level: level).collapsingSpaces
        }
        return Copy.charDescriptorNoLevel(race: race, spec: spec, className: className).collapsingSpaces
    }

    /// `Hordebreakers · US Normal`.
    static func whereLine(_ header: CharacterProfileHeader) -> String {
        let region = header.region.uppercased()
        let ruleset = RankingsCopy.ruleset(header.ruleset)
        guard let guild = header.guild?.name else { return Copy.charWhereNoGuild(region: region, ruleset: ruleset) }
        return Copy.charWhere(guild: guild, region: region, ruleset: ruleset)
    }

    static func isYours(_ header: CharacterProfileHeader, keys: Set<String>) -> Bool {
        keys.contains(header.key)
    }

    /// A rated fight exists: never a zero score (F11.6).
    static func hasScore(_ profile: CharacterProfile) -> Bool { profile.score != nil }

    static func score(_ profile: CharacterProfile) -> String {
        (profile.score ?? 0).formatted(.number.grouping(.automatic).locale(Locale(identifier: "en_US")))
    }

    static func scoreA11y(_ profile: CharacterProfile) -> String {
        Copy.charA11yScore(score: score(profile), tier: tierName(profile.tier))
    }

    static func tierName(_ tier: String?) -> String {
        switch tier {
        case "common": Copy.charTierCommon
        case "uncommon": Copy.charTierUncommon
        case "rare": Copy.charTierRare
        case "epic": Copy.charTierEpic
        case "legendary": Copy.charTierLegendary
        default: ""
        }
    }

    /// The tier's item quality, for its colour.
    static func tierQuality(_ tier: String?) -> Int? {
        switch tier {
        case "common": 1
        case "uncommon": 2
        case "rare": 3
        case "epic": 4
        case "legendary": 5
        default: nil
        }
    }

    /// `As of 17 Dec` from the API's `2026-12-17`.
    static func asOf(_ day: String?) -> String? {
        guard let day, let date = dayFormatter.date(from: day) else { return nil }
        return Copy.charAsOf(date: utcClock.shortDate(date))
    }

    /// The five bars, in the spec's order.
    static func bars(_ profile: CharacterProfile) -> [ScoreBar] {
        guard let values = profile.components, let caps = profile.componentCaps else { return [] }
        return [
            ScoreBar(label: Copy.charComponentGear, value: values.gear, cap: caps.gear),
            ScoreBar(label: Copy.charComponentExperience, value: values.experience, cap: caps.experience),
            ScoreBar(label: Copy.charComponentParses, value: values.parses, cap: caps.parses),
            ScoreBar(label: Copy.charComponentAttendance, value: values.attendance, cap: caps.attendance),
            ScoreBar(label: Copy.charComponentDiscipline, value: values.discipline, cap: caps.discipline),
        ]
    }

    /// The score history oldest first, each point 0 (lowest) to 1 (highest); flat is all 0.5.
    static func sparkline(_ history: [ScorePoint]) -> [Double] {
        let scores = history.sorted { $0.date < $1.date }.map { Double($0.score) }
        guard let low = scores.min(), let high = scores.max() else { return [] }
        guard high > low else { return scores.map { _ in 0.5 } }
        return scores.map { ($0 - low) / (high - low) }
    }

    static func bests(_ profile: CharacterProfile, now: Date) -> [FightLine] {
        let ruleset = RankingsCopy.ruleset(profile.character.ruleset)
        return profile.best.map { fight in
            line(fight, meta: Copy.charBestMeta(ruleset: ruleset, age: RelativeAge(from: fight.foughtAt, to: now).phrase))
        }
    }

    static func recent(_ profile: CharacterProfile, now: Date) -> [FightLine] {
        profile.history.map { fight in
            let result = fight.kill == false ? Copy.charWipe : Copy.charKill
            return line(fight, meta: Copy.charFightMeta(result: result, age: RelativeAge(from: fight.foughtAt, to: now).phrase))
        }
    }

    private static func line(_ fight: CharacterFight, meta: String) -> FightLine {
        FightLine(
            id: "\(fight.reportId)-\(fight.fightIndex)", encounter: fight.encounter, meta: meta,
            value: ReportsCopy.figure(fight.value), unit: RankingsCopy.fightUnit(fight.metric),
            percentile: fight.percentile.map { Int($0.rounded()) }, kill: fight.kill != false, reportID: fight.reportId
        )
    }

    static func builds(_ profile: CharacterProfile, clock: ReportsClock) -> [BuildLine] {
        profile.buildsSeen.map { build in
            BuildLine(
                trees: build.talentSplit.split(separator: "/").joined(separator: treeSeparator),
                meta: Copy.charBuildMeta(spec: build.spec, date: clock.shortDate(build.firstSeen))
            )
        }
    }

    /// `22.7 points`; nil without a gear export.
    static func gap(_ total: Double?) -> String? {
        total.map { Copy.charGapPoints(points: String(format: "%.1f", $0)) }
    }

    private static let utcClock = ReportsClock(zone: TimeZone(identifier: "UTC") ?? .gmt)

    private static var dayFormatter: DateFormatter {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_US_POSIX")
        formatter.timeZone = TimeZone(identifier: "UTC")
        formatter.dateFormat = "yyyy-MM-dd"
        return formatter
    }
}

private extension String {
    /// A descriptor with a part missing keeps single spaces and no stray edge.
    var collapsingSpaces: String {
        split(separator: " ", omittingEmptySubsequences: true).joined(separator: " ")
    }
}
