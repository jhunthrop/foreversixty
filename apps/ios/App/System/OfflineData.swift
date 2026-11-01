import Foundation

/// The five kinds of saved data You › Offline data lists (spec §5.7 F15.4), in its order.
enum OfflineCategory: String, CaseIterable, Sendable {
    case guides, bis, talents, search, icons

    var title: String {
        switch self {
        case .guides: Copy.offlineDataGuides
        case .bis: Copy.offlineDataBis
        case .talents: Copy.offlineDataTalents
        case .search: Copy.offlineDataSearch
        case .icons: Copy.offlineDataIcons
        }
    }

    var systemImage: String {
        switch self {
        case .guides: "book"
        case .bis: "shield"
        case .talents: "point.3.connected.trianglepath.dotted"
        case .search: "magnifyingglass"
        case .icons: "square.grid.2x2"
        }
    }

    /// Which category a saved file under a build belongs to; nil for files the screen does not
    /// list (the active build and dates at the root).
    static func of(relativePath: String) -> OfflineCategory? {
        let parts = relativePath.split(separator: "/").map(String.init)
        guard parts.count >= 2 else { return nil }
        switch parts[1] {
        // The guide files alone, as the Guides root's `Saved on this phone` counts them; the index,
        // tiers and changelog are shared files and join the search index's row.
        case "guides" where parts.last != "index.json": return .guides
        case "guides", "tiers.json", "changelog.json": return .search
        case "bis": return .bis
        case "talents": return .talents
        case "search.json", "items", "item-uses", "sets.json", "loot.json", "classes.json": return .search
        default: return nil
        }
    }

    /// The class a saved file is for: `shaman` from `guides/shaman-elemental.json` and
    /// `talents/shaman.json`; nil for shared files (`guides/index.json`, `tiers.json`).
    static func classSlug(relativePath: String) -> String? {
        let parts = relativePath.split(separator: "/").map(String.init)
        guard parts.count == 3, ["guides", "bis", "talents"].contains(parts[1]) else { return nil }
        let stem = parts[2].replacingOccurrences(of: ".json", with: "")
        guard stem != "index" else { return nil }
        return String(stem.prefix { $0 != "-" })
    }
}

/// One row of the screen: `Guides`, `Shaman, Warrior · 2.4 MB · Updated today`.
struct OfflineRow: Equatable, Sendable, Identifiable {
    let category: OfflineCategory
    let bytes: Int
    let updated: Date?
    /// The saved classes' names, for guides, lists and talents.
    let classes: [String]

    var id: String { category.rawValue }
}

/// The screen's content, pure over the saved files (spec §5.7 F15.4). No build ids, hashes or
/// file counts anywhere (tenet 7).
struct OfflineReport: Equatable, Sendable {
    let rows: [OfflineRow]

    var totalBytes: Int { rows.reduce(0) { $0 + $1.bytes } }
    var isEmpty: Bool { rows.isEmpty }

    static func build(files: [SavedStaticFile], iconBytes: Int, iconsUpdated: Date?) -> OfflineReport {
        var rows = OfflineCategory.allCases.filter { $0 != .icons }.compactMap { category -> OfflineRow? in
            let mine = files.filter { OfflineCategory.of(relativePath: $0.relativePath) == category }
            guard !mine.isEmpty else { return nil }
            let classes = Set(mine.compactMap { OfflineCategory.classSlug(relativePath: $0.relativePath) })
            return OfflineRow(
                category: category, bytes: mine.reduce(0) { $0 + $1.bytes }, updated: mine.map(\.savedAt).max(),
                classes: category == .search ? [] : classes.map(SearchWords.className).sorted()
            )
        }
        if iconBytes > 0 { rows.append(OfflineRow(category: .icons, bytes: iconBytes, updated: iconsUpdated, classes: [])) }
        return OfflineReport(rows: rows)
    }
}

/// The screen's words: sizes, days and the row lines. Pure; the zone comes in.
struct OfflineWords: Sendable {
    let calendar: Calendar

    static let tenMegabytes = 10_000_000
    private static let kilo = 1_000.0
    private static let mega = 1_000_000.0

    /// `2.4 MB`, `41 MB`, `410 KB`: decimal units, one decimal under 10 MB, whole numbers above.
    static func size(_ bytes: Int) -> String {
        let posix = Locale(identifier: "en_US_POSIX")
        if bytes >= tenMegabytes {
            return "\((Double(bytes) / mega).formatted(.number.precision(.fractionLength(0)).locale(posix))) MB"
        }
        if Double(bytes) >= mega {
            return "\((Double(bytes) / mega).formatted(.number.precision(.fractionLength(1)).locale(posix))) MB"
        }
        let kilobytes = bytes == 0 ? 0 : max(1, Int((Double(bytes) / kilo).rounded()))
        return "\(kilobytes) KB"
    }

    /// `today`, `yesterday`, else `24 Sep`.
    func day(_ date: Date, now: Date) -> String {
        if calendar.isDate(date, inSameDayAs: now) { return Copy.offlineDataToday }
        if let yesterday = calendar.date(byAdding: .day, value: -1, to: now), calendar.isDate(date, inSameDayAs: yesterday) {
            return Copy.offlineDataYesterday
        }
        return format("d MMM", date)
    }

    /// `9:41 pm`.
    func time(_ date: Date) -> String { format("h:mm a", date) }

    /// `Shaman, Warrior · 2.4 MB · Updated today`; the search index and icons have no what.
    func meta(_ row: OfflineRow, now: Date) -> String {
        let size = Self.size(row.bytes)
        // The board shows Item icons with their size only.
        guard row.category != .icons, let updated = row.updated else { return size }
        let when = day(updated, now: now)
        if row.classes.isEmpty { return Copy.offlineDataRowMetaNoWhat(size: size, when: when) }
        return Copy.offlineDataRowMeta(what: row.classes.joined(separator: ", "), size: size, when: when)
    }

    /// The footer: `Checked today at 9:41 pm. ...`, or `Not checked yet. ...`.
    func footer(checkedAt: Date?, now: Date) -> String {
        guard let checkedAt else { return Copy.offlineDataFooterNever }
        return Copy.offlineDataFooter(when: Copy.offlineDataCheckedAt(day: day(checkedAt, now: now), time: time(checkedAt)))
    }

    private func format(_ pattern: String, _ date: Date) -> String {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_US_POSIX")
        formatter.calendar = calendar
        formatter.timeZone = calendar.timeZone
        formatter.amSymbol = "am"
        formatter.pmSymbol = "pm"
        formatter.dateFormat = pattern
        return formatter.string(from: date)
    }
}

/// The two Updates switches (spec F15.4), both on by default, kept on this phone.
struct OfflineSettings: Sendable {
    static let keepKey = "offlineData.keepClassesReady"
    static let wifiKey = "offlineData.wifiOnly"
    static let checkedKey = "offlineData.checkedAt"

    nonisolated(unsafe) let defaults: UserDefaults

    init(defaults: UserDefaults = .standard) { self.defaults = defaults }

    var keepClassesReady: Bool {
        get { defaults.object(forKey: Self.keepKey) as? Bool ?? true }
        nonmutating set { defaults.set(newValue, forKey: Self.keepKey) }
    }

    var wifiOnly: Bool {
        get { defaults.object(forKey: Self.wifiKey) as? Bool ?? true }
        nonmutating set { defaults.set(newValue, forKey: Self.wifiKey) }
    }

    var checkedAt: Date? {
        get { defaults.object(forKey: Self.checkedKey) as? Date }
        nonmutating set { defaults.set(newValue, forKey: Self.checkedKey) }
    }

    /// A download that waits for Wi-Fi (icons, new guides) may run on this connection now.
    func mayDownload(on connectivity: any Connectivity) -> Bool {
        !wifiOnly || !connectivity.isExpensive
    }
}
