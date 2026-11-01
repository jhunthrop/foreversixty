import Foundation
import Models

/// The guide files' `YYYY-MM-DD` stamps as the spec prints them: `1 Nov 2026` and `1 Nov`.
enum GuideDates {
    private static let isoFormat = "yyyy-MM-dd"
    private static let longFormat = "d MMM yyyy"
    private static let shortFormat = "d MMM"
    private static let monthDayFormat = "MM/dd"

    /// `1 Nov 2026`; the stamp itself when it is not a date.
    static func long(_ iso: String) -> String { reformat(iso, longFormat) }
    /// `1 Nov`.
    static func short(_ iso: String) -> String { reformat(iso, shortFormat) }
    /// `10/28`, a changelog row's date.
    static func monthDay(_ iso: String) -> String { reformat(iso, monthDayFormat) }

    /// Today's `YYYY-MM-DD` in `zone`, to compare with a stamp.
    static func iso(_ date: Date, zone: TimeZone = .current) -> String {
        formatter(isoFormat, zone: zone).string(from: date)
    }

    private static func reformat(_ iso: String, _ format: String) -> String {
        guard let date = formatter(isoFormat, zone: .gmt).date(from: iso) else { return iso }
        return formatter(format, zone: .gmt).string(from: date)
    }

    private static func formatter(_ format: String, zone: TimeZone) -> DateFormatter {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_GB")
        formatter.timeZone = zone
        formatter.dateFormat = format
        return formatter
    }
}

/// A role word as the guides print it: `DPS`, `Healer`, `Tank`.
enum GuideRole {
    static let dps = "dps"
    static let healer = "healer"
    static let tank = "tank"

    static func label(_ role: String) -> String {
        switch role {
        case healer: Copy.guidesRoleHealer
        case tank: Copy.guidesRoleTank
        default: Copy.guidesRoleDps
        }
    }
}

/// One cell of the `All classes` grid.
struct GuideClassCell: Identifiable, Equatable, Sendable {
    let classSlug: String
    let name: String

    var id: String { classSlug }
}

/// One spec's guide as the index names it.
struct GuideSpecRef: Identifiable, Hashable, Sendable {
    let classSlug: String
    let className: String
    /// The guide key, `shaman-elemental`.
    let spec: String
    let specSlug: String
    let specName: String
    let role: String
    let icon: String?
    let updated: String

    var id: String { spec }
    /// `Elemental Shaman`.
    var title: String { Copy.bisSpecChip(spec: specName, className: className) }
}

/// The index read as the Guides tab draws it (spec F9.1, F9.6).
enum GuidesIndexLogic {
    static func cells(_ index: GuideIndex) -> [GuideClassCell] {
        index.classes.map { GuideClassCell(classSlug: $0.classSlug, name: $0.className) }
    }

    static func specs(of classSlug: String, in index: GuideIndex) -> [GuideSpecRef] {
        index.classes.filter { $0.classSlug == classSlug }.flatMap(refs)
    }

    static func spec(_ key: String, in index: GuideIndex) -> GuideSpecRef? {
        index.classes.flatMap(refs).first { $0.spec == key }
    }

    /// The current character's guide: its class and its spec's tree name; nil with no character.
    static func yourGuide(_ index: GuideIndex, character: AccountCharacter?) -> GuideSpecRef? {
        guard let classSlug = character?.className, let specName = character?.spec else { return nil }
        return specs(of: classSlug, in: index).first { $0.specName.caseInsensitiveCompare(specName) == .orderedSame }
    }

    private static func refs(_ entry: GuideIndexClass) -> [GuideSpecRef] {
        entry.specs.map {
            GuideSpecRef(
                classSlug: entry.classSlug, className: entry.className, spec: $0.spec, specSlug: $0.specSlug,
                specName: $0.specName, role: $0.role, icon: $0.icon, updated: $0.updated
            )
        }
    }
}

/// The guide files saved on this phone, by guide key, with their size on disk.
struct SavedGuides: Equatable, Sendable {
    let bytes: [String: Int]

    static let none = SavedGuides(bytes: [:])

    func isSaved(_ spec: String) -> Bool { bytes[spec] != nil }

    /// `Saved on this phone: Warrior, Shaman · 32 KB. …`, classes in index order.
    func footer(_ index: GuideIndex) -> String {
        guard !bytes.isEmpty else { return Copy.guidesFooterNone }
        let classes = index.classes
            .filter { entry in entry.specs.contains { isSaved($0.spec) } }
            .map(\.className)
        let size = ByteCountFormatter.string(fromByteCount: Int64(bytes.values.reduce(0, +)), countStyle: .file)
        return Copy.guidesFooter(classes: classes.joined(separator: GuideCopy.listSeparator), size: size)
    }
}

/// The guide sections the strip names, in reading order (spec F9.2).
enum GuideSection: String, CaseIterable, Identifiable, Sendable {
    case overview, talents, rotation, stats, gear, races

    var id: String { rawValue }

    var title: String {
        switch self {
        case .overview: Copy.guidesSectionOverview
        case .talents: Copy.guidesSectionTalents
        case .rotation: Copy.guidesSectionRotation
        case .stats: Copy.guidesSectionStats
        case .gear: Copy.guidesSectionGear
        case .races: Copy.guidesSectionRaces
        }
    }

    /// Exactly the sections the guide has; a name the app does not know is left out.
    static func of(_ guide: GuideFile) -> [GuideSection] {
        guide.sections.compactMap(GuideSection.init(rawValue:))
    }

    /// `Talents, section 2 of 6`.
    static func a11y(_ section: GuideSection, in sections: [GuideSection]) -> String {
        let position = (sections.firstIndex(of: section) ?? 0) + 1
        return Copy.guidesA11ySection(name: section.title, n: position, total: sections.count)
    }
}

/// One numbered stat chip.
struct StatChip: Identifiable, Equatable, Sendable {
    let number: Int
    let stat: String

    var id: Int { number }
    var a11y: String { Copy.guidesA11yStat(n: number, stat: stat) }
}

/// Every F9.2 sentence built from a guide file; the words live in the copy catalog.
enum GuideCopy {
    static let siteURL = URL(string: "https://foreversixty.gg")!
    private static let guidesPath = "guides"
    private static let treeSeparator = " / "
    /// Names in a sentence: `Dwarf, Orc`.
    static let listSeparator = ", "

    static func title(_ guide: GuideFile) -> String {
        Copy.bisSpecChip(spec: guide.specName, className: guide.className)
    }

    /// `Updated 1 Nov 2026 · Recommended races: Dwarf, Orc`.
    static func headMeta(_ guide: GuideFile) -> String {
        let date = GuideDates.long(guide.updated)
        let races = guide.recommendedRaces.map(\.name)
        guard !races.isEmpty else { return Copy.guidesHeadMetaNoRaces(date: date) }
        return Copy.guidesHeadMeta(date: date, races: races.joined(separator: listSeparator))
    }

    /// `DPS · Updated 1 Nov 2026`, the Your guide card.
    static func cardMeta(_ spec: GuideSpecRef) -> String {
        Copy.guidesCardMeta(role: GuideRole.label(spec.role), date: GuideDates.long(spec.updated))
    }

    /// `Healer · Updated 1 Nov`, a class page row.
    static func rowMeta(_ spec: GuideSpecRef) -> String {
        Copy.guidesCardMeta(role: GuideRole.label(spec.role), date: GuideDates.short(spec.updated))
    }

    /// The web page the share sheet sends: `https://foreversixty.gg/guides/{class}/{spec}`.
    static func shareURL(_ guide: GuideFile) -> URL {
        siteURL.appending(path: guidesPath).appending(path: guide.classSlug).appending(path: guide.specSlug)
    }

    static func statChips(_ guide: GuideFile) -> [StatChip] {
        guide.statPriority.enumerated().map { StatChip(number: $0.offset + 1, stat: $0.element) }
    }

    /// `5 ranks`, `1 rank`; nil when the file has no rank count.
    static func ranks(_ talent: GuideKeyTalent) -> String? {
        talent.maxRank.map { Copy.guidesRanks(n: $0) }
    }

    /// `31 / 10 / 10`.
    static func totals(_ build: GuideBuild) -> String {
        build.totals.map(String.init).joined(separator: treeSeparator)
    }

    /// `Best in slot for Elemental · 60 · Raid-ready, Phase 1`.
    static func gearRow(_ guide: GuideFile) -> String? {
        guide.gear.map { Copy.guidesGearRow(spec: guide.specName, preset: $0.presetLabel) }
    }
}

/// A build row of the Talents section.
struct GuideBuildRow: Identifiable, Equatable, Sendable {
    let title: String
    let totals: String

    var id: String { title }
}

enum GuideBuildRows {
    /// The raid build, then the leveling build. A null raid build is the leveling build itself,
    /// so its one row is the raid build.
    static func rows(_ talents: GuideTalents) -> [GuideBuildRow] {
        guard let raid = talents.raidBuild else {
            return talents.levelingBuild.map { [GuideBuildRow(title: Copy.guidesRaidBuild, totals: GuideCopy.totals($0))] } ?? []
        }
        let leveling = talents.levelingBuild.map { GuideBuildRow(title: Copy.guidesLevelingBuild, totals: GuideCopy.totals($0)) }
        return [GuideBuildRow(title: Copy.guidesRaidBuild, totals: GuideCopy.totals(raid))] + (leveling.map { [$0] } ?? [])
    }
}

/// A numbered rotation row.
struct RotationRow: Identifiable, Equatable, Sendable {
    let number: Int
    let spell: GuideSpell

    var id: Int { number }
    /// `Rank 4`, or nil when the file names no rank.
    var rankText: String? { spell.rank.map { Copy.guidesRank(n: $0) } }
}

/// One list of the rotation: Opener, Priority or Cooldowns.
struct RotationGroup: Identifiable, Equatable, Sendable {
    let title: String
    let rows: [RotationRow]

    var id: String { title }
}

enum RotationRows {
    /// The guide's own lists in its order, each numbered from 1; an empty list is left out.
    static func groups(_ rotation: GuideRotation?) -> [RotationGroup] {
        guard let rotation else { return [] }
        return [
            (Copy.guidesOpener, rotation.opener),
            (Copy.guidesPriority, rotation.priority),
            (Copy.guidesCooldowns, rotation.cooldowns),
        ]
        .filter { !$0.1.isEmpty }
        .map { title, spells in
            RotationGroup(title: title, rows: spells.enumerated().map { RotationRow(number: $0.offset + 1, spell: $0.element) })
        }
    }
}
