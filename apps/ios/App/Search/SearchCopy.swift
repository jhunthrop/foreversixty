import Foundation
import Models

/// Every F12 sentence built from data (spec §5.4); the words live in the copy catalog.
enum SearchWords {
    static func className(_ slug: String) -> String {
        GameClass(rawValue: slug.lowercased())?.name ?? slug.capitalized
    }

    /// `Elemental Shaman`.
    static func specClass(spec: String, classSlug: String) -> String {
        Copy.bisSpecChip(spec: spec, className: className(classSlug))
    }

    static func scope(_ scope: SearchScope) -> String {
        switch scope {
        case .all: Copy.searchScopeAll
        case .items: Copy.searchScopeItems
        case .characters: Copy.searchScopeCharacters
        case .guides: Copy.searchScopeGuides
        case .talents: Copy.searchScopeTalents
        }
    }

    static func group(_ group: SearchGroupKind) -> String {
        switch group {
        case .guides: Copy.searchGroupGuides
        case .talents: Copy.searchGroupTalents
        case .items: Copy.searchGroupItems
        case .characters: Copy.searchGroupCharacters
        }
    }

    /// `Wrist · item level 57 · requires level 52`; no level part for an item anyone can wear.
    static func itemMeta(_ item: SearchItem) -> String {
        let slot = ItemVocabulary.slotLabel(item.slot)
        return item.requiredLevel > 0
            ? Copy.searchItemMeta(slot: slot, ilvl: item.itemLevel, level: item.requiredLevel)
            : Copy.searchItemMetaNoLevel(slot: slot, ilvl: item.itemLevel)
    }

    /// `Warrior · Level 58 · US PvP`; the class or level the API does not know is left out.
    static func characterMeta(_ row: CharacterSearchRow) -> String {
        let className = row.className.map(className) ?? ""
        let region = row.region.uppercased()
        let ruleset = AccountCopy.ruleset(row.ruleset)
        let meta = row.level.map { Copy.searchCharacterMeta(className: className, level: $0, region: region, ruleset: ruleset) }
            ?? Copy.searchCharacterMetaNoLevel(className: className, region: region, ruleset: ruleset)
        return row.className == nil ? meta.trimmingCharacters(in: CharacterSet(charactersIn: " ·")) : meta
    }

    /// `Guide · DPS`.
    static func guideMeta(_ guide: SearchGuide) -> String {
        Copy.searchGuideMeta(role: BisLabels.role(guide.role))
    }

    /// `Elemental Shaman · Elemental · 3 ranks`.
    static func talentMeta(_ talent: SearchTalent) -> String {
        Copy.searchTalentMeta(
            spec: talent.specName, className: className(talent.className), tree: talent.tree,
            ranks: Copy.searchTalentRanks(n: talent.maxRank)
        )
    }

    /// Recent's second line: what kind of thing it is.
    static func kind(_ target: SearchTarget) -> String {
        switch target {
        case .item: Copy.searchRecentItem
        case .character: Copy.searchRecentCharacter
        case .guide: Copy.searchRecentGuide
        case .talent: Copy.searchRecentTalent
        }
    }
}

/// The site pages a search result opens in the browser until its app screen exists:
/// guides (F9), character pages and rankings (F11), the planner (F13).
enum SearchLinks {
    static let site = URL(string: "https://foreversixty.gg")!
    static let rankings = site.appending(path: "rankings")
    static let planner = HomeLinks.planner

    static func guide(className: String, spec: String) -> URL {
        site.appending(path: "guides/\(className)/\(spec)")
    }

    /// `key` is `{region}/{ruleset}/{slug}`, the page's own path.
    static func character(key: String) -> URL {
        site.appending(path: "character/\(key)")
    }

    /// A spec key's guide path segment: `hunter-beast-mastery` → `beast-mastery`.
    static func guideSpec(specKey: String, className: String) -> String {
        specKey.hasPrefix(className + "-") ? String(specKey.dropFirst(className.count + 1)) : specKey
    }
}
