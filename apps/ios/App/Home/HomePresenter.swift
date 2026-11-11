import Foundation
import Models

/// The character Home is about, from `/v1/me` (lane F1's selector picks it).
struct HomeCharacter: Equatable, Sendable {
    let key: String
    let region: String
    let ruleset: String
    let name: String
    let race: String?
    let className: String?
    /// The primary tree's name as the API gives it (`Beast Mastery`).
    let specName: String?
    let level: Int?
    let guild: String?
    /// The guild's id, for its raids; nil outside a guild.
    let guildID: Int64?
    let faction: String?
    /// Worn item ids by slot from the last export, for the gear rows the answer does not name.
    let gear: [String: Int]
    /// The last export's talent ranks and the build it was made on (lane F6 draws them).
    let talents: BuildTalents?
    let dataBuild: String?

    init(_ character: AccountCharacter) {
        self.init(
            key: character.key, region: character.region, ruleset: character.ruleset, name: character.name,
            race: character.race, className: character.className, specName: character.spec, level: character.level,
            guild: character.guild?.name, guildID: character.guild?.id, faction: character.faction,
            gear: character.build?.gear ?? [:],
            talents: character.build?.talents, dataBuild: character.build?.dataBuild
        )
    }

    init(
        key: String, region: String, ruleset: String, name: String, race: String?, className: String?,
        specName: String?, level: Int?, guild: String?, guildID: Int64? = nil, faction: String?, gear: [String: Int] = [:],
        talents: BuildTalents? = nil, dataBuild: String? = nil
    ) {
        self.guildID = guildID
        self.key = key
        self.region = region
        self.ruleset = ruleset
        self.name = name
        self.race = race
        self.className = className
        self.specName = specName
        self.level = level
        self.guild = guild
        self.faction = faction
        self.gear = gear
        self.talents = talents
        self.dataBuild = dataBuild
    }

    /// The path segment of the character routes (`velaren` in `us/normal/velaren`).
    var slug: String { key.split(separator: "/").last.map(String.init) ?? name.lowercased() }
}

/// The home answer's words that outlive the old answer card: the band, the rows' order, what a row
/// replaces. Formats only; every gain and pick is the server's.
enum HomePresenter {
    /// The rows the loading skeleton draws.
    static let visibleRows = 5


    static func bandHead(_ home: CharacterHome, character: HomeCharacter) -> String? {
        range(home, character: character).map(BisLabels.head)
    }


    /// The list Home answers from: the character's level through the site's band rule
    /// (`BisRange.forLevel`), the rule BiS opens the character's list with. The server's band
    /// stands in only when the level is unknown, and goes through the same rule.
    static func range(_ home: CharacterHome, character: HomeCharacter) -> BisRange? {
        (character.level ?? home.band).map(BisRange.forLevel)
    }


    /// The rows in the coordinator's order (F2 fix round, 2026-10-31), for presentation only:
    /// measured gains largest first, then gains not modelled yet, then empty slots. The gains
    /// are the server's; this only arranges them. Equal rows keep the server's order.
    static func presentationOrder(_ upgrades: [HomeUpgrade]) -> [HomeUpgrade] {
        func group(_ upgrade: HomeUpgrade) -> Int {
            if upgrade.modelled, upgrade.gain != nil { return 0 }
            return upgrade.worn == nil ? 2 : 1
        }
        return upgrades.enumerated().sorted { lhs, rhs in
            let (gl, gr) = (group(lhs.element), group(rhs.element))
            if gl != gr { return gl < gr }
            if gl == 0, let a = lhs.element.gain, let b = rhs.element.gain, a != b { return a > b }
            return lhs.offset < rhs.offset
        }.map(\.element)
    }


    /// `Replaces Belt of the Archmage, ilvl 62`; the server's own label for an item our data
    /// does not know (`Not in our item data yet`); `Empty slot` when nothing is worn.
    static func replaces(_ worn: HomeWorn?) -> String {
        guard let worn else { return Copy.homeUpgradesEmpty }
        guard worn.known, let name = worn.name, let level = worn.itemLevel else { return worn.label }
        return Copy.homeUpgradesReplaces(item: name, n: level)
    }


    /// `night_elf` → `Night Elf`; the API's slugs are lowercase.
    static func words(_ slug: String) -> String {
        slug.split(whereSeparator: { $0 == "_" || $0 == "-" || $0 == " " })
            .map { $0.prefix(1).capitalized + $0.dropFirst() }
            .joined(separator: " ")
    }
}

extension ItemTooltipRequest {
    /// A tooltip for an item on Home or Your gear, in the answer's build and spec and Home's band.
    /// The band label is the band head (`60 · raid`), so the pick footer names the raid list without a preset id.
    static func forHome(
        _ look: HomeItemLook, home: CharacterHome, character: HomeCharacter, sourceLine: String? = nil, wornID: Int? = nil
    ) -> ItemTooltipRequest {
        let range = HomePresenter.range(home, character: character)
        return ItemTooltipRequest(
            itemID: look.id, fallback: look, build: home.dataBuild, classSlug: classSlug(home: home, character: character),
            spec: home.spec, specName: character.specName, band: range?.band, bandLabel: HomePresenter.bandHead(home, character: character),
            preset: range?.preset, faction: character.faction, sourceLine: sourceLine, wornID: wornID
        )
    }

    /// The items file to read: the class part of the spec slug (`hunter-beast-mastery`), else the character's class.
    static func classSlug(home: CharacterHome, character: HomeCharacter) -> String {
        home.spec?.split(separator: "-").first.map(String.init) ?? character.className ?? ""
    }
}
