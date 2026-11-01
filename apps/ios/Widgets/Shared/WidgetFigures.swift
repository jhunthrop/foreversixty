import Foundation
import Models

/// Which of the raid widget's states the snapshot is in (spec §5.6, board F14.2).
enum WidgetRaidState: Equatable, Sendable {
    case signedOut
    case noCharacter
    case noGuild
    case noRaid
    case raid
}

/// The foot tag's kind, which picks its colour: kill for In, gold for an ask, wipe for Cancelled.
enum WidgetTagKind: String, Equatable, Sendable {
    case `in`, late, bench, out, notSignedUp, chooseSpec, locked, cancelled
}

/// How far away a moment is, in the one rounding every surface uses: minutes left is
/// ceil(seconds / 60); a day or more is days and whole hours, less than a day hours and minutes.
struct Countdown: Equatable, Sendable {
    let days: Int
    let hours: Int
    let minutes: Int

    init(from now: Date, to target: Date) {
        let totalMinutes = Int((target.timeIntervalSince(now) / 60).rounded(.up))
        days = totalMinutes / 1440
        hours = (totalMinutes % 1440) / 60
        minutes = totalMinutes % 60
    }

    var totalMinutes: Int { days * 1440 + hours * 60 + minutes }

    /// `in 2 days 2 h`, `in 3 h 40 min`, `in 39 min`.
    var phrase: String {
        switch (days, hours, minutes) {
        case (1, 0, _): Copy.widgetInDays(days: 1)
        case (1, _, _): Copy.widgetInDayHours(hours: hours)
        case let (d, 0, _) where d > 1: Copy.widgetInDays(days: d)
        case let (d, h, _) where d > 1: Copy.widgetInDaysHours(days: d, hours: h)
        case let (_, h, 0) where h > 0: Copy.widgetInHours(hours: h)
        case let (_, h, m) where h > 0: Copy.widgetInHoursMinutes(hours: h, minutes: m)
        default: Copy.widgetInMinutes(minutes: minutes)
        }
    }

    /// The leading unit alone: `2d`, `3h`, `39m` (the second circular lock-screen widget).
    var short: String {
        if days > 0 { return Copy.widgetShortDays(n: days) }
        if hours > 0 { return Copy.widgetShortHours(n: hours) }
        return Copy.widgetShortMinutes(n: minutes)
    }

    /// The spoken form: `in 2 days 2 hours`.
    var spoken: String {
        var parts: [String] = []
        if days > 0 { parts.append(Copy.widgetA11yDays(n: days)) }
        if hours > 0 { parts.append(Copy.widgetA11yHours(n: hours)) }
        if days == 0, minutes > 0 { parts.append(Copy.widgetA11yMinutes(n: minutes)) }
        return Copy.widgetA11yIn(span: parts.joined(separator: " "))
    }
}

/// Every word and figure the widgets, lock-screen widgets and the Live Activity show, derived
/// once from ONE snapshot and ONE clock. No surface formats a raid figure on its own, so the
/// lock screen, the home screen and the Live Activity can never disagree (the wow-player SHOULD
/// on F14). The golden vectors in apps/shared/widgets/figures.json hold iOS and Android to the
/// same words.
struct WidgetFigures: Equatable, Sendable {
    let state: WidgetRaidState
    let zone: String?
    let tag: String?
    let tagKind: WidgetTagKind?
    let countdown: String?
    let short: String?
    let time: String?
    let timeShort: String?
    let inline: String
    let you: String?
    let signedUp: String?
    let signedUpCount: Int?
    let size: Int?
    let stale: String?
    let buttons: Bool
    /// The answer drawn checked on the medium widget (the pending one while a press saves).
    let checked: String?
    /// The answer whose button says `Not saved`, while it does.
    let notSaved: String?
    let liveActive: Bool
    let liveTitle: String?
    let liveUntil: String?
    /// When the Live Activity counts to: invite_at, else starts_at.
    let liveTarget: Date?
    let accessibility: String

    static let staleAfterHours = 12
    static let liveLead: TimeInterval = 60 * 60
    static let liveTail: TimeInterval = 15 * 60

    init(_ snapshot: WidgetSnapshot?, now: Date, zone timeZone: TimeZone) {
        let clock = WidgetClock(zone: timeZone)
        guard let snapshot else {
            self.init(empty: .signedOut, inline: Copy.widgetSignedOut)
            return
        }
        let widget = snapshot.widget
        let stale = Self.stale(fetchedAt: snapshot.fetchedAt, now: now)
        if widget.character == nil {
            self.init(empty: .noCharacter, inline: Copy.widgetNoCharacter, stale: stale)
        } else if widget.guild == nil {
            self.init(empty: .noGuild, inline: Copy.widgetNoGuild, stale: stale)
        } else if let raid = widget.nextRaid {
            self.init(raid: raid, snapshot: snapshot, stale: stale, now: now, clock: clock)
        } else {
            self.init(empty: .noRaid, inline: Copy.widgetNoRaid, stale: stale)
        }
    }

    private init(empty state: WidgetRaidState, inline: String, stale: String? = nil) {
        self.state = state
        zone = nil; tag = nil; tagKind = nil; countdown = nil; short = nil; time = nil; timeShort = nil
        self.inline = inline
        you = nil; signedUp = nil; signedUpCount = nil; size = nil
        self.stale = stale
        buttons = false; checked = nil; notSaved = nil
        liveActive = false; liveTitle = nil; liveUntil = nil; liveTarget = nil
        accessibility = inline
    }

    private init(raid: WidgetNextRaid, snapshot: WidgetSnapshot, stale: String?, now: Date, clock: WidgetClock) {
        let mine = Self.effectiveState(raid, snapshot: snapshot)
        let personal = Self.personalKind(state: mine, spec: raid.my.spec)
        let kind: WidgetTagKind = switch raid.status {
        case "cancelled": .cancelled
        case "locked": .locked
        default: personal
        }
        let role = Self.roleWord(raid.my.role)
        state = .raid
        zone = raid.zoneName
        tagKind = kind
        tag = Self.tagWord(kind, role: role)
        let cancelled = raid.status == "cancelled"
        let count = cancelled || now >= raid.startsAt ? nil : Countdown(from: now, to: raid.startsAt)
        countdown = cancelled ? nil : (count?.phrase ?? Copy.widgetNow)
        short = count?.short
        time = clock.dayTime(raid.startsAt)
        timeShort = clock.dayTimeShort(raid.startsAt)
        inline = Copy.widgetInline(zone: raid.zoneName, time: clock.dayTimeShort(raid.startsAt), answer: Self.inlineWord(kind, role: role))
        you = Copy.widgetYou(answer: Self.youWord(personal, role: role))
        signedUpCount = raid.signedUp
        size = raid.size
        signedUp = raid.size.map { Copy.widgetOfSize(n: raid.signedUp, size: $0) } ?? "\(raid.signedUp)"
        self.stale = stale
        buttons = raid.canSignUp && raid.status == "open"
        checked = mine
        notSaved = snapshot.notSavedUntil.map { now < $0 } == true ? snapshot.notSavedState : nil
        let target = raid.inviteAt ?? raid.startsAt
        liveTarget = target
        let inWindow = now >= target.addingTimeInterval(-Self.liveLead) && now < raid.startsAt.addingTimeInterval(Self.liveTail)
        liveActive = !cancelled && inWindow && (mine == "in" || mine == "late")
        liveTitle = liveActive ? Self.liveTitle(zone: raid.zoneName, target: target, invites: raid.inviteAt != nil, now: now) : nil
        liveUntil = raid.inviteAt.map { Copy.liveInvites(time: clock.time($0)) } ?? Copy.liveStarts(time: clock.time(raid.startsAt))
        accessibility = Self.spoken(raid, kind: personal, role: role, count: count, clock: clock)
    }

    /// The pending press wins over the snapshot's answer until the server answers.
    private static func effectiveState(_ raid: WidgetNextRaid, snapshot: WidgetSnapshot) -> String? {
        if let pending = snapshot.pending, pending.raidID == raid.raidId { return pending.state }
        return raid.my.state
    }

    private static func personalKind(state: String?, spec: String?) -> WidgetTagKind {
        if spec == nil, state != "out" { return .chooseSpec }
        switch state {
        case "in": return .in
        case "late": return .late
        case "bench": return .bench
        case "out": return .out
        default: return .notSignedUp
        }
    }

    static func stateWord(_ state: String?) -> String? {
        switch state {
        case "in": Copy.raidStateIn
        case "late": Copy.raidStateLate
        case "bench": Copy.raidStateBench
        case "out": Copy.raidStateOut
        default: nil
        }
    }

    static func roleWord(_ role: String) -> String {
        switch role {
        case "tank": Copy.raidRoleTank
        case "healer": Copy.raidRoleHealer
        default: Copy.raidRoleDps
        }
    }

    /// `In · DPS`; Late, Bench and Out stand alone (spec §5.6).
    private static func tagWord(_ kind: WidgetTagKind, role: String) -> String {
        switch kind {
        case .in: Copy.widgetAnswer(state: Copy.raidStateIn, role: role)
        case .late: Copy.raidStateLate
        case .bench: Copy.raidStateBench
        case .out: Copy.raidStateOut
        case .notSignedUp: Copy.widgetNotSignedUp
        case .chooseSpec: Copy.widgetChooseSpec
        case .locked: Copy.widgetLocked
        case .cancelled: Copy.widgetCancelled
        }
    }

    /// The inline widget names the answer without the role: `Onyxia's Lair · Thu 8 pm · In`.
    private static func inlineWord(_ kind: WidgetTagKind, role: String) -> String {
        kind == .in ? Copy.raidStateIn : tagWord(kind, role: role)
    }

    private static func stale(fetchedAt: Date, now: Date) -> String? {
        let hours = Int(now.timeIntervalSince(fetchedAt) / 3600)
        return hours >= staleAfterHours ? Copy.widgetUpdated(n: hours) : nil
    }

    /// `Invites for {zone} in 39 minutes` with an invite time; `{zone} starts in 39 minutes` when
    /// the countdown runs to the pull, so the title never names invites a raid does not have.
    private static func liveTitle(zone: String, target: Date, invites: Bool, now: Date) -> String {
        let minutes = Countdown(from: now, to: target).totalMinutes
        switch (invites, minutes) {
        case (true, ...0): return Copy.liveNotifTitleNow(zone: zone)
        case (true, 1): return Copy.liveNotifTitleOne(zone: zone)
        case (true, _): return Copy.liveNotifTitle(zone: zone, n: minutes)
        case (false, ...0): return Copy.livePullTitleNow(zone: zone)
        case (false, 1): return Copy.livePullTitleOne(zone: zone)
        default: return Copy.livePullTitle(zone: zone, n: minutes)
        }
    }

    /// The `You:` line names the role with every answer: `You: Late · DPS` (figures.json).
    private static func youWord(_ kind: WidgetTagKind, role: String) -> String {
        switch kind {
        case .in, .late, .bench, .out:
            Copy.widgetAnswer(state: stateWord(kind.rawValue) ?? kind.rawValue, role: role)
        default:
            tagWord(kind, role: role)
        }
    }

    /// `Next raid, Onyxia's Lair, Thursday 8 pm, in 2 days 2 hours. You are In as DPS. 24 of 40 signed up.`
    private static func spoken(_ raid: WidgetNextRaid, kind: WidgetTagKind, role: String, count: Countdown?, clock: WidgetClock) -> String {
        let when = count?.spoken ?? (raid.status == "cancelled" ? Copy.widgetCancelled : Copy.widgetNow)
        var parts = [Copy.widgetA11yRaid(zone: raid.zoneName, when: clock.spokenDayTime(raid.startsAt), countdown: when)]
        if let word = stateWord(raid.my.state) {
            parts.append(Copy.widgetA11yYou(state: word, role: role))
        } else {
            parts.append(tagWord(kind, role: role) + ".")
        }
        if let size = raid.size { parts.append(Copy.widgetA11ySignedUp(n: raid.signedUp, size: size)) }
        return parts.joined(separator: " ")
    }
}

/// The phone's clock words: the widget shows the phone's time only (spec §5.6 Time).
struct WidgetClock: Sendable {
    let zone: TimeZone

    private func format(_ pattern: String, _ date: Date) -> String {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_US_POSIX")
        formatter.timeZone = zone
        formatter.amSymbol = "am"
        formatter.pmSymbol = "pm"
        formatter.dateFormat = pattern
        return formatter.string(from: date)
    }

    private func onTheHour(_ date: Date) -> Bool {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = zone
        return calendar.component(.minute, from: date) == 0
    }

    /// `Thu 8:00 pm`
    func dayTime(_ date: Date) -> String { format("EEE h:mm a", date) }
    /// `Thu 8 pm`, minutes kept when not on the hour.
    func dayTimeShort(_ date: Date) -> String { onTheHour(date) ? format("EEE h a", date) : dayTime(date) }
    /// `8:00 pm`
    func time(_ date: Date) -> String { format("h:mm a", date) }
    /// `Thursday 8 pm`
    func spokenDayTime(_ date: Date) -> String { onTheHour(date) ? format("EEEE h a", date) : format("EEEE h:mm a", date) }
    /// `10 Dec`
    func dayMonth(_ date: Date) -> String { format("d MMM", date) }
    /// `24 Sep 2026`
    func fullDate(_ date: Date) -> String { format("d MMM yyyy", date) }
    /// `Thursday` (lane F15's spoken answers).
    func weekday(_ date: Date) -> String { format("EEEE", date) }
    /// `Thu`
    func shortWeekday(_ date: Date) -> String { format("EEE", date) }
    /// `8 pm`, minutes kept when not on the hour (`8:30 pm`).
    func hourTime(_ date: Date) -> String { onTheHour(date) ? format("h a", date) : time(date) }
}

/// The upgrade widgets' words, from the same snapshot.
struct UpgradeFigures: Equatable, Sendable {
    struct Row: Equatable, Sendable {
        let slot: String
        let name: String
        let quality: Int
        let icon: String
        let gain: String
        let itemID: Int
    }

    let header: String?
    let rows: [Row]
    let footer: String?
    let rectTitle: String?
    /// The words that stand in for rows: the server's sentence, or `Open the app`.
    let sentence: String?
    let opensApp: Bool

    init(_ snapshot: WidgetSnapshot?) {
        guard let widget = snapshot?.widget, let character = widget.character else {
            self.init(sentence: Copy.widgetSignedOut, opensApp: true)
            return
        }
        guard let upgrades = widget.upgrades else {
            self.init(sentence: Copy.widgetOpenApp, opensApp: true)
            return
        }
        guard upgrades.state == "ok", !upgrades.rows.isEmpty else {
            self.init(sentence: upgrades.sentence ?? Copy.widgetNoGear, opensApp: false)
            return
        }
        let unit = upgrades.unit.uppercased()
        let rows = upgrades.rows.map { row in
            Row(
                slot: row.slotLabel, name: row.pick.name, quality: row.pick.quality, icon: row.pick.icon,
                gain: row.gain.map { Copy.widgetGain(gain: String(format: "%.1f", $0), unit: unit) } ?? Copy.widgetNoGain,
                itemID: row.pick.id
            )
        }
        let count = Copy.homeAnswerUpgrades(n: upgrades.count)
        self.init(
            header: Copy.widgetUpgrades(name: character.name), rows: rows,
            footer: upgrades.totalGain.map { Copy.widgetUpgradesFooter(upgrades: count, total: String(format: "%.1f", $0), unit: unit) } ?? count,
            rectTitle: upgrades.rows.first.map { "\(Copy.widgetNextUpgrade) · \($0.slotLabel)" }, sentence: nil, opensApp: false
        )
    }

    private init(header: String?, rows: [Row], footer: String?, rectTitle: String?, sentence: String?, opensApp: Bool) {
        self.header = header
        self.rows = rows
        self.footer = footer
        self.rectTitle = rectTitle
        self.sentence = sentence
        self.opensApp = opensApp
    }

    private init(sentence: String, opensApp: Bool) {
        self.init(header: nil, rows: [], footer: nil, rectTitle: nil, sentence: sentence, opensApp: opensApp)
    }
}
