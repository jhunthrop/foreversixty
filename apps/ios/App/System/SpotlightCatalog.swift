import Foundation
import Models

/// One Spotlight row's words and where its tap goes (spec §5.7, board F15.3).
struct SpotlightRow: Codable, Equatable, Sendable {
    let id: String
    let title: String
    let subtitle: String
}

/// What the app puts in Spotlight: your characters (up to 9) and the first 50 watched, the next
/// 5 planned raids and the guides saved on the phone. Items are never indexed. Pure.
struct SpotlightCatalog: Codable, Equatable, Sendable {
    static let ownLimit = 9
    static let watchedLimit = 50
    static let raidLimit = 5

    let characters: [SpotlightRow]
    let raids: [SpotlightRow]
    let guides: [SpotlightRow]

    static let empty = SpotlightCatalog(characters: [], raids: [], guides: [])

    /// A saved guide: its spec and when its file was saved.
    struct SavedGuide: Equatable, Sendable {
        let spec: SpecInfo
        let savedAt: Date
    }

    static func build(
        own: [AccountCharacter], watched: [WatchedCharacter], planned: [RaidView], guildID: Int64?,
        guides: [SavedGuide], answers: IntentAnswers, now: Date
    ) -> SpotlightCatalog {
        let mine = own.prefix(ownLimit).map(row)
        let ownKeys = Set(mine.map(\.id))
        let others = watched.filter { !ownKeys.contains($0.key) }.prefix(watchedLimit).map(row)
        let raids = guildID.map { guild in
            planned.filter { $0.status != "cancelled" && $0.startsAt > now }
                .sorted { $0.startsAt < $1.startsAt }
                .prefix(raidLimit)
                .map { raidRow($0, guild: guild, answers: answers) }
        } ?? []
        return SpotlightCatalog(characters: mine + others, raids: raids, guides: guides.map { guideRow($0, clock: answers.clock) })
    }

    /// `Elemental Shaman · 60 · Olympus XXVII`.
    static func row(_ character: AccountCharacter) -> SpotlightRow {
        let className = SearchWords.className(character.className ?? "")
        let level = character.level.map(String.init) ?? ""
        let spec = character.spec ?? ""
        let sub = character.guild.map {
            Copy.spotlightCharacterSub(spec: spec, className: className, level: level, guild: $0.name)
        } ?? Copy.spotlightCharacterSubNoGuild(spec: spec, className: className, level: level)
        return SpotlightRow(id: character.key, title: character.name, subtitle: tidy(sub))
    }

    /// A watched character has no spec on the list: `Shaman · 60`.
    static func row(_ character: WatchedCharacter) -> SpotlightRow {
        let className = SearchWords.className(character.className ?? "")
        let sub = Copy.spotlightCharacterSubNoGuild(spec: "", className: className, level: character.level.map(String.init) ?? "")
        return SpotlightRow(id: character.key, title: character.name, subtitle: tidy(sub))
    }

    /// `Onyxia's Lair · Thu 8 pm`, `Raid · Thu 10 Dec`. The id is `{guild}/{raid}`.
    static func raidRow(_ raid: RaidView, guild: Int64, answers: IntentAnswers) -> SpotlightRow {
        let date = "\(answers.clock.shortWeekday(raid.startsAt)) \(answers.clock.dayMonth(raid.startsAt))"
        return SpotlightRow(
            id: "\(guild)/\(raid.id)", title: answers.raidTitle(zone: raid.zoneName, startsAt: raid.startsAt),
            subtitle: Copy.spotlightRaidSub(date: date)
        )
    }

    /// `Elemental Shaman guide`, `Guide · Updated 24 Sep 2026`. The id is the guide file's key.
    static func guideRow(_ guide: SavedGuide, clock: WidgetClock) -> SpotlightRow {
        let className = SearchWords.className(guide.spec.classSlug)
        return SpotlightRow(
            id: guide.spec.spec, title: Copy.shareGuide(spec: guide.spec.name, className: className),
            subtitle: Copy.spotlightGuideSub(date: clock.fullDate(guide.savedAt))
        )
    }

    /// An empty part leaves no stray space or separator: `Shaman · 60`, never ` Shaman · 60 · `.
    static func tidy(_ text: String) -> String {
        text.components(separatedBy: " · ")
            .map { $0.trimmingCharacters(in: .whitespaces) }
            .filter { !$0.isEmpty }
            .joined(separator: " · ")
    }
}
