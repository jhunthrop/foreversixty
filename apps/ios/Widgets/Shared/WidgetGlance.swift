import Foundation
import Models

/// The words the restyled widgets draw (flows spec F1, F5, E7; mocks ios-widgets-*): a label, the
/// big figure, a name and the lines under it, derived once from the snapshot and the clock. The
/// raid figures underneath stay WidgetFigures' (the golden vectors both apps share); this adds the
/// glance's own words on top.
struct RaidGlance: Equatable, Sendable {
    enum Kind: Equatable, Sendable {
        /// A raid the character can see: its countdown, zone, time and answer.
        case raid
        /// Before the first raids open: the countdown to them.
        case firstRaids
        /// Raids are open but none is planned for the character.
        case raidsOpen
        /// Signed out, or no character: one sentence and `Open the app`.
        case empty
    }

    let kind: Kind
    /// The small grey line on top: the guild's name, `First raids`, `Raids open now`.
    let label: String
    /// The countdown (`33d 8h`), or `Open`.
    let big: String
    /// The zone, the first raids' day, or the open raids' names.
    let name: String
    /// The time with its zone and the answer naming its character (E7), or the context line.
    let lines: [String]
    /// The answer drawn selected on the answering widget; nil until answered.
    let selected: String?
    /// The four answers can be pressed (open, not ended, verified).
    let canAnswer: Bool
    /// Every answer opens the spec ask instead (no spec saved yet).
    let asksSpec: Bool
    /// The answer whose button says `Not saved`, for ten seconds after a failed press.
    let notSaved: String?

    static let answers = ["in", "late", "bench", "out"]

    init(_ snapshot: WidgetSnapshot?, now: Date, zone: TimeZone) {
        let figures = WidgetFigures(snapshot, now: now, zone: zone)
        let clock = WidgetClock(zone: zone)
        guard let widget = snapshot?.widget, widget.character != nil else {
            self.init(empty: figures.state == .signedOut ? Copy.widgetSignedOut : Copy.widgetNoCharacter)
            return
        }
        if let raid = widget.nextRaid, widget.guild != nil {
            self.init(raid: raid, widget: widget, figures: figures, clock: clock, now: now)
            return
        }
        let context = widget.guild == nil ? Copy.widgetFindGuild : Copy.widgetNoRaidYet
        // `Barrow Deeps, Hyjal Summit, Onyxia's Lair` (the API's first_raids_zones).
        let zones = (widget.firstRaidsZones ?? []).joined(separator: ", ")
        if let opens = Self.firstRaids(widget.firstRaidsAt), opens > now {
            let count = Countdown(from: now, to: opens)
            self.init(
                kind: .firstRaids, label: Copy.widgetFirstRaids, big: count.compact, name: WidgetClock(zone: .gmt).dayDate(opens),
                lines: zones.isEmpty ? [] : [zones]
            )
        } else if zones.isEmpty {
            self.init(kind: .raidsOpen, label: Copy.widgetRaidsOpen, big: Copy.widgetOpen, name: context, lines: [])
        } else {
            self.init(kind: .raidsOpen, label: Copy.widgetRaidsOpen, big: Copy.widgetOpen, name: zones, lines: [context])
        }
    }

    private init(raid: WidgetNextRaid, widget: Widget, figures: WidgetFigures, clock: WidgetClock, now: Date) {
        let started = now >= raid.startsAt
        let big: String = switch raid.status {
        case "cancelled": Copy.widgetCancelled
        default: started ? Copy.widgetNow : Countdown(from: now, to: raid.startsAt).compact
        }
        var lines = [clock.dayDateTimeZone(raid.startsAt)]
        if raid.status == "locked" { lines.append(Copy.widgetLocked) }
        lines.append(Self.answerLine(state: figures.checked, name: widget.character?.name ?? "", role: raid.my.role))
        if let stale = figures.stale { lines.append(stale) }
        self.init(
            kind: .raid, label: raid.guildName, big: big, name: raid.zoneName, lines: lines, selected: figures.checked,
            canAnswer: figures.buttons, asksSpec: figures.tagKind == .chooseSpec, notSaved: figures.notSaved
        )
    }

    private init(empty sentence: String) {
        self.init(kind: .empty, label: Copy.appName, big: "", name: sentence, lines: [Copy.widgetOpenApp])
    }

    private init(
        kind: Kind, label: String, big: String, name: String, lines: [String],
        selected: String? = nil, canAnswer: Bool = false, asksSpec: Bool = false, notSaved: String? = nil
    ) {
        self.kind = kind
        self.label = label
        self.big = big
        self.name = name
        self.lines = lines
        self.selected = selected
        self.canAnswer = canAnswer
        self.asksSpec = asksSpec
        self.notSaved = notSaved
    }

    /// `In, as Thandrel, Damage`; `Not answered, as Thandrel, Damage` (E7).
    static func answerLine(state: String?, name: String, role: String) -> String {
        Copy.widgetAs(answer: WidgetFigures.stateWord(state) ?? Copy.widgetNotAnswered, name: name, role: roleWord(role))
    }

    /// The flows spec's role words: Tank, Healer, Damage.
    static func roleWord(_ role: String) -> String {
        switch role {
        case "tank": Copy.raidRoleTank
        case "healer": Copy.raidRoleHealer
        default: Copy.widgetRoleDamage
        }
    }

    /// The first raids open at the start of their day in UTC, as the site counts to them and names
    /// the day (`Wed 9 Dec` everywhere, not the evening before in the Americas).
    static func firstRaids(_ day: String?) -> Date? {
        guard let day else { return nil }
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_US_POSIX")
        formatter.timeZone = .gmt
        formatter.dateFormat = "yyyy-MM-dd"
        return formatter.date(from: day)
    }

    /// One spoken label for the whole glance.
    var spoken: String {
        Copy.widgetA11yNext(label: label, big: big, name: name, lines: lines.joined(separator: ". "))
    }
}

/// The upgrade side of the widgets: the next upgrade's gain, never a total (flows appendix C), and
/// the item with its icon and quality colour.
struct UpgradeGlance: Equatable, Sendable {
    struct Item: Equatable, Sendable {
        let name: String
        let quality: Int
        let icon: String
    }

    /// `+14.5 DPS`, or nil when there is no upgrade to show.
    let gain: String?
    let item: Item?
    /// Where it comes from (`Crafted: Blacksmithing`), the API's `source_line`.
    let source: String?
    /// The words that stand in when there is no upgrade: the server's sentence, `Open the app`.
    let sentence: String?

    init(_ snapshot: WidgetSnapshot?) {
        let figures = UpgradeFigures(snapshot)
        guard let row = figures.rows.first else {
            gain = nil
            item = nil
            source = nil
            sentence = figures.sentence
            return
        }
        gain = row.gain
        item = Item(name: row.name, quality: row.quality, icon: row.icon)
        let line = snapshot?.widget.upgrades?.rows.first?.pick.sourceLine?.trimmingCharacters(in: .whitespaces)
        source = line?.isEmpty == false ? line : nil
        sentence = nil
    }

    var spoken: String {
        guard let gain, let item else { return [Copy.widgetNextUpgrade, sentence].compactMap { $0 }.joined(separator: ", ") }
        return [Copy.widgetNextUpgrade, item.name, source, gain].compactMap { $0 }.joined(separator: ", ")
    }
}

extension Countdown {
    /// The widgets' big figure: `33d 8h`, `3h 40m`, `39m`.
    var compact: String {
        if days > 0 { return "\(Copy.widgetShortDays(n: days)) \(Copy.widgetShortHours(n: hours))" }
        if hours > 0 { return "\(Copy.widgetShortHours(n: hours)) \(Copy.widgetShortMinutes(n: minutes))" }
        return Copy.widgetShortMinutes(n: minutes)
    }
}
