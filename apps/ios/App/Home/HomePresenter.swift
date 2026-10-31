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
    let faction: String?
    /// Worn item ids by slot from the last export, for the gear rows the answer does not name.
    let gear: [String: Int]

    init(_ character: AccountCharacter) {
        self.init(
            key: character.key, region: character.region, ruleset: character.ruleset, name: character.name,
            race: character.race, className: character.className, specName: character.spec, level: character.level,
            guild: character.guild?.name, faction: character.faction, gear: character.build?.gear ?? [:]
        )
    }

    init(
        key: String, region: String, ruleset: String, name: String, race: String?, className: String?,
        specName: String?, level: Int?, guild: String?, faction: String?, gear: [String: Int] = [:]
    ) {
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
    }

    /// The path segment of the character routes (`velaren` in `us/normal/velaren`).
    var slug: String { key.split(separator: "/").last.map(String.init) ?? name.lowercased() }
}

/// One upgrade row (spec F2.1 item 2).
struct HomeUpgradeRow: Equatable, Identifiable {
    enum Gain: Equatable {
        /// `+5.3` over `DPS`.
        case measured(figure: String, unit: String)
        /// The server's words for a gain it did not model (`Gain not modelled yet`).
        case words(String)
    }

    let slot: String
    let pick: HomeItemLook
    let wornID: Int?
    let meta: String
    let replaces: String
    let gain: Gain
    let sourceLine: String

    var id: String { slot }
}

/// What the answer region says (spec F2.1 to F2.4 and the empty states).
enum HomeAnswerBody: Equatable {
    case upgrades(rows: [HomeUpgradeRow], footer: String)
    case allMatch(String)
    case message(String, buttons: [HomeAction], note: String?)
}

/// A button on an answer sentence.
struct HomeAction: Equatable {
    enum Kind: Equatable { case getAddon, openList, openPlanner }
    let kind: Kind
    let title: String
    let primary: Bool
}

/// One "Next for {name}" row.
struct HomeNextRow: Equatable, Identifiable {
    enum Kind: Equatable { case talents, sims }
    enum Tag: Equatable { case matches, differs }
    let kind: Kind
    let title: String
    let line: String
    let tag: Tag?
    var id: Kind { kind }
}

/// Everything Home draws for one answer. Built by `HomePresenter`, rendered by `HomeAnswerView`.
struct HomeAnswerModel: Equatable {
    let heroMeta: String
    let count: String?
    let gain: (figure: String, unit: String)?
    let note: String?
    let synced: String?
    let bandHead: String?
    let body: HomeAnswerBody
    let already: String?
    let nextTitle: String
    let next: [HomeNextRow]

    static func == (lhs: Self, rhs: Self) -> Bool {
        lhs.heroMeta == rhs.heroMeta && lhs.count == rhs.count && lhs.gain?.figure == rhs.gain?.figure
            && lhs.gain?.unit == rhs.gain?.unit && lhs.note == rhs.note && lhs.synced == rhs.synced
            && lhs.bandHead == rhs.bandHead && lhs.body == rhs.body && lhs.already == rhs.already
            && lhs.nextTitle == rhs.nextTitle && lhs.next == rhs.next
    }
}

/// Turns the API's home answer into words. It formats and never ranks: the rows keep the
/// server's order and every gain is the server's (plan guard rail "one behaviour, one place").
struct HomePresenter {
    static let visibleRows = 5

    let character: HomeCharacter
    let finishedSims: [FinishedSim]
    let time: HomeTime

    func model(_ home: CharacterHome) -> HomeAnswerModel {
        let isAnswer = home.state == "ok" && home.upgradeCount > 0
        return HomeAnswerModel(
            heroMeta: heroMeta,
            count: isAnswer ? Copy.homeAnswerUpgrades(n: home.upgradeCount) : nil,
            gain: isAnswer ? home.totalGain.map { (ItemNumbers.signedTenths($0), home.unitWord ?? "") } : nil,
            note: isAnswer ? note(home) : nil,
            synced: synced(home),
            bandHead: bandHead(home),
            body: body(home),
            already: already(home),
            nextTitle: Copy.homeNextTitle(name: character.name),
            next: next(home)
        )
    }

    /// The Next rows alone, for the error state with nothing cached.
    func nextRowsWithoutAnswer() -> [HomeNextRow] { [simsRow(spec: nil)] }

    /// `Orc Beast Mastery Hunter · Level 60 · Olympus XXVII`.
    var heroMeta: String {
        let who = [character.race, character.specName, character.className]
            .compactMap { $0.map(Self.words) }
            .joined(separator: " ")
        return [who.isEmpty ? nil : who, character.level.map { Copy.homeHeroLevel(n: $0) }, character.guild]
            .compactMap(\.self)
            .joined(separator: " · ")
    }

    /// `60 · raid` for the raid list, the band's words (`50 to 59`) otherwise.
    func bandHead(_ home: CharacterHome) -> String? { Self.bandHead(home, character: character) }

    static func bandHead(_ home: CharacterHome, character: HomeCharacter) -> String? {
        range(home, character: character).map(BisLabels.head)
    }

    /// The list Home answers from: the character's level through the site's band rule
    /// (`BisRange.forLevel`), the rule BiS opens the character's list with. The server's band
    /// stands in only when the level is unknown, and goes through the same rule.
    static func range(_ home: CharacterHome, character: HomeCharacter) -> BisRange? {
        (character.level ?? home.band).map(BisRange.forLevel)
    }

    /// The band's words for the sentences (`50 to 59`, `60`), empty when no band is known.
    private func bandWords(_ home: CharacterHome) -> String? {
        Self.range(home, character: character).map(BisLabels.band)
    }

    private func note(_ home: CharacterHome) -> String? {
        guard let label = bandWords(home) else { return nil }
        let base = Copy.homeAnswerNote(band: label)
        return home.notModelledCount > 0 ? base + Copy.homeAnswerNotModelled(n: home.notModelledCount) : base
    }

    private func synced(_ home: CharacterHome) -> String? {
        guard let at = home.syncedAt else { return nil }
        let age = time.age(since: at)
        return home.syncedFrom == "blizzard" ? Copy.homeAnswerSyncedBnet(age: age) : Copy.homeAnswerSyncedAddon(age: age)
    }

    private func body(_ home: CharacterHome) -> HomeAnswerBody {
        let spec = character.specName ?? ""
        let band = bandWords(home) ?? ""
        switch home.state {
        case "no_gear":
            return .message(Copy.homeAnswerNoGear, buttons: [
                HomeAction(kind: .getAddon, title: Copy.signinGetAddon, primary: true),
                HomeAction(kind: .openList, title: Copy.homeAnswerListButton(spec: spec, band: band), primary: false),
            ], note: Copy.homeAnswerNoGearNote)
        case "under_20":
            return .message(Copy.homeAnswerUnder20, buttons: [
                HomeAction(kind: .openList, title: Copy.homeAnswerUnder20Button(spec: spec), primary: false),
            ], note: nil)
        case "no_spec":
            return .message(Copy.homeAnswerPickSpec, buttons: [
                HomeAction(kind: .openPlanner, title: Copy.homeAnswerOpenPlanner, primary: false),
            ], note: nil)
        case "no_list":
            return .message(Copy.homeAnswerNoList, buttons: [], note: nil)
        default:
            guard home.upgradeCount > 0 else { return .allMatch(Copy.homeAnswerAllMatch(band: band)) }
            let rows = Self.presentationOrder(home.upgrades).prefix(Self.visibleRows).map { row($0, unit: home.unitWord ?? "") }
            let more = home.upgradeCount - rows.count
            return .upgrades(rows: rows, footer: more > 0 ? Copy.homeUpgradesMore(n: more) : Copy.homeUpgradesAllSlots)
        }
    }

    private func row(_ upgrade: HomeUpgrade, unit: String) -> HomeUpgradeRow {
        let pick = HomeItemLook(
            id: upgrade.pick.id, name: upgrade.pick.name, quality: upgrade.pick.quality,
            icon: upgrade.pick.icon, itemLevel: upgrade.pick.itemLevel
        )
        return HomeUpgradeRow(
            slot: upgrade.slot,
            pick: pick,
            wornID: upgrade.worn?.id,
            meta: "\(upgrade.slotLabel) · \(upgrade.pick.sourceLine)",
            replaces: Self.replaces(upgrade.worn),
            gain: upgrade.modelled && upgrade.gain != nil
                ? .measured(figure: ItemNumbers.signedTenths(upgrade.gain ?? 0), unit: unit)
                : .words(upgrade.gainWords),
            sourceLine: upgrade.pick.sourceLine
        )
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

    private func already(_ home: CharacterHome) -> String? {
        guard !home.alreadyBest.isEmpty else { return nil }
        return Copy.homeUpgradesAlready(slots: home.alreadyBest.map(\.slotLabel).joined(separator: ", "))
    }

    private func next(_ home: CharacterHome) -> [HomeNextRow] {
        [talentsRow(home), simsRow(spec: home.spec)]
    }

    private func talentsRow(_ home: CharacterHome) -> HomeNextRow {
        let band = bandWords(home) ?? ""
        let talents = home.talents
        switch talents.state {
        case "optimized":
            return HomeNextRow(kind: .talents, title: Copy.homeNextTalents,
                               line: Copy.homeNextTalentsMatch(spec: character.specName ?? "", band: band), tag: .matches)
        case "differs":
            return HomeNextRow(kind: .talents, title: Copy.homeNextTalents,
                               line: Copy.homeNextTalentsDiffer(n: talents.pointsDiffer ?? 0, m: talents.bandPoints ?? 0, band: band),
                               tag: .differs)
        case "no_build":
            return HomeNextRow(kind: .talents, title: Copy.homeNextTalents, line: Copy.homeNextTalentsNoBuild, tag: nil)
        default:
            return HomeNextRow(kind: .talents, title: Copy.homeNextTalents, line: Copy.homeNextTalentsNone, tag: nil)
        }
    }

    /// The newest unseen finished server run for this spec (`/v1/me` `finished_sims`, newest first).
    private func simsRow(spec: String?) -> HomeNextRow {
        let sim = finishedSims.first { $0.spec == spec }
        let line = sim.map { Copy.homeNextSimFinished(spec: character.specName ?? "", dps: String(Int($0.dps.rounded()))) }
        return HomeNextRow(kind: .sims, title: Copy.homeNextSims, line: line ?? Copy.homeNextNoSim, tag: nil)
    }

    /// `night_elf` → `Night Elf`; the API's slugs are lowercase.
    static func words(_ slug: String) -> String {
        slug.split(whereSeparator: { $0 == "_" || $0 == "-" || $0 == " " })
            .map { $0.prefix(1).uppercased() + $0.dropFirst() }
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
