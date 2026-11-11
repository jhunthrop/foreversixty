import Auth
import DesignSystem
import Foundation
import Models

extension ReadyCheck {
    static let okState = "ok"

    /// What is not ready before the pull (flows spec F2): every item not ok, in the server's order.
    var open: [ReadyItem] { items.filter { $0.state != Self.okState } }
}

/// A pending guild applicant (F3), from `GET /v1/guilds/{id}/applications`.
struct ApplicantLine: Equatable, Sendable {
    /// The application's id (the applicant screen opens it).
    let id: Int64
    let name: String
    /// `Holy Priest`.
    let kind: String
    let level: Int
    let appliedAt: Date
}

/// Everything Home's blocks are drawn from, as loaded.
struct HomeFlowInputs: Equatable {
    var character: HomeCharacter
    var home: CharacterHome?
    var raid: RaidView?
    var night: RaidNightSummary?
    var activity: ActivityPage?
    var badge: RaidBadge?
    var applicants: [ApplicantLine] = []
    var ready: ReadyCheck?
    var bestRun: BestRun?
    var firstRaids: Date?
    /// The raid answer's state (F1): the answer just sent, Late's minutes, the status line, offline.
    var answering = RaidAnswering()

    /// The answer just sent, shown before the raid reloads.
    var answer: SignupState? {
        get { answering.pending }
        set { answering.pending = newValue }
    }

    var lateOpen: Bool {
        get { answering.lateOpen }
        set { answering.lateOpen = newValue }
    }
}

/// Turns the server's blocks and the loaded data into what Home draws (flows spec 4.1). A block
/// whose data is missing is left out rather than drawn empty.
struct HomeFlowPresenter {
    let now: Date
    var clock = RaidClock()
    var calendar: Calendar = .current

    static let maxLevel = 60
    static let bandStep = 10

    // MARK: Head

    func head(_ character: HomeCharacter) -> HomeHeadContent {
        var facts: [String] = []
        if let level = character.level {
            let kind = [character.specName, character.className.map(HomePresenter.words)].compactMap(\.self).joined(separator: " ")
            facts.append(Copy.homeHeadLevel(n: level, who: kind))
            if level < Self.maxLevel {
                let band = min(Self.maxLevel, (level / Self.bandStep + 1) * Self.bandStep)
                facts.append(Copy.homeHeadToBand(n: band - level, band: band))
            } else if let guild = character.guild {
                facts.append(guild)
            }
        }
        return HomeHeadContent(name: character.name, className: character.className, facts: facts)
    }

    // MARK: Blocks

    func blocks(_ blocks: [HomeBlock], _ inputs: HomeFlowInputs) -> [HomeBlockContent] {
        let bestRun = HomeBlock.showsBestRun(blocks)
        return blocks.compactMap { block in
            switch block {
            case .head: nil
            case let .upgrades(limit): inputs.home.map { .upgrades(upgrades($0, limit: limit, bestRun: bestRun ? inputs.bestRun : nil)) }
            case let .nextRaid(card): nextRaid(inputs).map { card ? .raidCard($0) : .raidLine($0) }
            case .raidLive: raidLive(inputs).map { .raidLive($0) }
            case .raidSummary: inputs.night.map { .lastNight(lastNight($0)) }
            case let .activity(query): inputs.activity.map { .newForYou(newForYou($0, query: query)) }
            case .guildCTA: .findGuild(findGuild(inputs.firstRaids))
            case let .notice(text, _): .notice(text)
            }
        }
    }

    // MARK: Next upgrades

    func upgrades(_ home: CharacterHome, limit: Int, bestRun: BestRun?) -> HomeUpgradesContent {
        let unit = home.unitWord ?? ""
        let rows = HomePresenter.presentationOrder(home.upgrades).prefix(limit).map { upgrade in
            HomeUpgradeItem(
                slot: upgrade.slot,
                look: HomeItemLook(
                    id: upgrade.pick.id, name: upgrade.pick.name, quality: upgrade.pick.quality,
                    icon: upgrade.pick.icon, itemLevel: upgrade.pick.itemLevel
                ),
                source: Self.sourceLine(upgrade),
                gain: upgrade.modelled ? upgrade.gain.map { HomeGain(figure: ReportsCopy.signedHouseFigure($0), unit: unit) } : nil,
                wornID: upgrade.worn?.id,
                sourceLine: upgrade.pick.sourceLine
            )
        }
        let note: String? = rows.isEmpty ? Self.emptyNote(home) : nil
        var content = HomeUpgradesContent(rows: Array(rows), bestRun: bestRun.map { bestRunRow($0, home: home) }, note: note)
        content.meta = upgradesMeta(home)
        content.stamp = syncedStamp(home)
        content.stampStale = home.syncedAt.map { SyncAge.isStale($0, now: now) } ?? false
        return content
    }

    static let partSeparator = " · "

    /// `Head · Blacksmithing · replaces Lionheart Helm · 4%`: the worn item when the API names it,
    /// the drop chance when the source has one.
    static func sourceLine(_ upgrade: HomeUpgrade) -> String {
        let worn = upgrade.worn.flatMap { $0.known ? $0.name : nil }.map { Copy.homeUpgradesReplacesWorn(item: $0) }
        let chance = upgrade.pick.dropChance.map(ItemNumbers.percent)
        return [upgrade.slotLabel, upgrade.pick.sourceLine, worn, chance].compactMap(\.self).joined(separator: partSeparator)
    }

    /// `13 upgrades · +157.6 DPS`; the count alone without a modelled total.
    func upgradesMeta(_ home: CharacterHome) -> String? {
        guard home.upgradeCount > 0 else { return nil }
        let count = Copy.homeAnswerUpgrades(n: home.upgradeCount)
        guard let total = home.totalGain, total > 0 else { return count }
        let gain = Copy.homeUpgradesGain(figure: ReportsCopy.signedHouseFigure(total), unit: home.unitWord ?? "")
        return Copy.homeUpgradesMeta(upgrades: count, gain: gain)
    }

    /// `Synced 2 h ago from the addon`; `Synced 2 h ago` from anywhere else.
    func syncedStamp(_ home: CharacterHome) -> String? {
        guard let synced = home.syncedAt else { return nil }
        let age = HomeTime(now: now, calendar: calendar).age(since: synced)
        return home.syncedFrom == Self.addonSource ? Copy.homeAnswerSyncedAddon(age: age) : Copy.homeAnswerSynced(age: age)
    }

    static let addonSource = "addon"

    /// Best run as a row: the dungeon over the drop and its chance, the drop's icon when the answer
    /// names that item, the gain.
    func bestRunRow(_ run: BestRun, home: CharacterHome) -> HomeBestRunRow {
        let pick = run.itemID.flatMap { id in home.upgrades.first { $0.pick.id == id }?.pick }
        let line = [Copy.homeBestRunRow(item: run.item), run.chance.map(ItemNumbers.percent)].compactMap(\.self)
            .joined(separator: Self.partSeparator)
        return HomeBestRunRow(
            place: run.place, line: line, icon: pick?.icon,
            gain: HomeGain(figure: ReportsCopy.signedHouseFigure(run.gain), unit: home.unitWord ?? "")
        )
    }

    /// The line an answer with no rows shows, by the server's state: only `ok` means every slot matches.
    static func emptyNote(_ home: CharacterHome) -> String {
        switch home.state {
        case "no_gear": Copy.homeAnswerNoGear
        case "no_spec": Copy.homeAnswerPickSpec
        case "under_20": Copy.homeAnswerUnder20
        default: Copy.homeAnswerAllMatch(band: home.bandLabel ?? "")
        }
    }


    // MARK: The next raid

    func nextRaid(_ inputs: HomeFlowInputs) -> NextRaidContent? {
        guard let raid = inputs.raid else { return nil }
        let viewer = raid.viewer.characters.first { $0.characterKey == inputs.character.key } ?? raid.viewer.characters.first
        let named = Self.namedCharacter(raid, key: inputs.character.key)
        let name = named?.name ?? ""
        let role = Self.roleWord(named?.role)
        let selected = inputs.answer ?? viewer?.state.flatMap(SignupState.init(rawValue:))
        let zone = TimeZone(identifier: raid.guildTimezone) ?? clock.zone
        let clockTime = clock.guildClock(raid.startsAt, guildZone: zone).replacingOccurrences(of: " ", with: Self.noBreak)
        let time = Copy.homeRaidWhen(date: clock.shortDate(raid.startsAt), time: clockTime)
        let counts = Copy.homeRaidCounts(in: raid.counts.signedIn, late: raid.counts.late)
        // AS5: once answered the Segmented is the answer; the line names only the character (E7), or is empty.
        let answer: String = switch (selected, named) {
        case (_?, nil): ""
        case (_?, _?): Copy.raidTabAs(name: name, role: role)
        case (nil, nil): Copy.homeRaidNotAnsweredAlone
        case (nil, _?): Copy.homeRaidNotAnswered(name: name, role: role)
        }
        let open = named == nil ? Copy.homeRaidLineSubAlone(counts: counts) : Copy.homeRaidLineSub(counts: counts, name: name, role: role)
        // AS5: once answered the Segmented is the answer; the line keeps the counts and who it is for.
        let answeredSub = ([counts] + (named.map { [Copy.raidTabAs(name: $0.name, role: role)] } ?? []))
            .joined(separator: Self.partSeparator)
        let guild = inputs.character.guild ?? ""
        let saved = raid.signups.first { $0.own && $0.characterKey == viewer?.characterKey }?.lateMinutes
        let minutes = selected == .late ? (inputs.answer == .late ? inputs.answering.pendingMinutes ?? saved : saved) : nil
        var content = NextRaidContent(
            guildID: inputs.character.guildID ?? 0, raidID: raid.id, characterKey: viewer?.characterKey ?? inputs.character.key,
            zone: raid.zoneName,
            place: Copy.homeRaidPlace(guild: guild, when: time),
            clock: clockParts(to: raid.startsAt),
            clockLabel: clockLabel(to: raid.startsAt),
            line: HomeLine(title: Copy.homeRaidLineTitle(raid: raid.zoneName, when: time), sub: selected == nil ? open : answeredSub),
            icon: raid.zoneIcon ?? RaidIcons.of(zone: raid.zone),
            answer: answer,
            selected: selected,
            lateOpen: inputs.lateOpen && selected == .late,
            tally: tally(raid),
            ready: ready(inputs.ready),
            officer: officer(inputs, raid: raid)
        )
        content.headLine = headLine(to: raid.startsAt, zone: zone, when: time)
        content.who = named.map { Copy.raidTabAs(name: $0.name, role: role) }
        content.lateMinutes = minutes
        content.status = inputs.answering.line
        content.busy = inputs.answering.busy
        content.answersEnabled = !inputs.answering.offline
        return content
    }

    /// `In`, or `Late · 30 min` when a Late answer said how late (F1).
    static func answerWord(_ state: SignupState, lateMinutes: Int?) -> String {
        guard state == .late, let lateMinutes else { return state.label }
        return Copy.raidAnswerLateMinutes(n: lateMinutes)
    }

    /// Within a day: `8:00 pm CST · pull in 2 h 48 min`; further out the date and time.
    func headLine(to start: Date, zone: TimeZone, when: String) -> String {
        let seconds = max(0, Int(start.timeIntervalSince(now)))
        guard seconds < Self.day else { return when }
        let hours = seconds / 3600
        let minutes = (seconds % 3600) / 60
        let until = hours > 0 ? Copy.raidTabUntilHours(h: hours, m: minutes) : Copy.raidTabUntilMinutes(m: minutes)
        return Copy.raidTabPullIn(time: clock.guildClock(start, guildZone: zone), until: until)
    }

    static let day = 86_400
    /// Keeps `7:00 pm CST` together when a raid's line wraps.
    static let noBreak = "\u{00A0}"

    /// E7 (owner, 2026-11-09): the character an answer is for, named only when the account has more
    /// than one character in the raid's guild; nil when the answer alone says it.
    static func namedCharacter(_ raid: RaidView, key: String) -> RaidViewerCharacter? {
        let mine = raid.viewer.characters
        guard mine.count > 1 else { return nil }
        return mine.first { $0.characterKey == key } ?? mine.first
    }

    /// The clock read aloud as one timer (`2 hours, 48 minutes`).
    func clockLabel(to start: Date) -> String {
        let seconds = max(0, Int(start.timeIntervalSince(now)))
        return Copy.homeRaidClock(h: seconds / 3600, m: (seconds % 3600) / 60)
    }

    /// `02 hours`, `48 min` until the pull; none once it has started.
    func clockParts(to start: Date) -> [FSClockPart] {
        let seconds = Int(start.timeIntervalSince(now))
        guard seconds > 0 else { return [] }
        let hours = seconds / 3600
        let minutes = (seconds % 3600) / 60
        return [
            FSClockPart(String(format: "%02d", hours), Copy.homeRaidHours),
            FSClockPart(String(format: "%02d", minutes), Copy.homeRaidMin),
        ]
    }

    /// Who is coming by role from the one count model (rows 1 and 2): against the comp target when the
    /// raid has one (`Healers 5/8 · 3 short`), the plain count without (`Healers 6`).
    func tally(_ raid: RaidView) -> [RoleLineContent] {
        let counts = RaidRoleCounts.of(raid)
        let target = raid.compTarget
        let roles: [(wire: String, label: String, want: Int?)] = [
            (RaidRoleCounts.tank, Copy.raidGroupTanks, target?.tanks),
            (RaidRoleCounts.healer, Copy.raidGroupHealers, target?.healers),
            (RaidRoleCounts.dps, Copy.homeRaidDamage, target.map { $0.melee + $0.ranged }),
        ]
        return roles.map { RoleLineContent(role: $0.label, have: counts.count($0.wire), want: $0.want, names: "") }
    }

    func ready(_ check: ReadyCheck?) -> HomeLine? {
        guard let check, !check.open.isEmpty else { return nil }
        let items = check.open.map(\.line).joined(separator: "; ")
        let line = check.exportedAt.map { Copy.homeReadyLine(items: items, age: RelativeAge(from: $0, to: now).phrase) } ?? items
        return HomeLine(title: Copy.homeReadyTitle(n: check.open.count), sub: line)
    }

    /// The officer's open items (F3): who has not answered with the short role, and the applicants.
    func officer(_ inputs: HomeFlowInputs, raid: RaidView?) -> OfficerContent? {
        guard let badge = inputs.badge, badge.officer else { return nil }
        // Row 6: the raid's own not-answered list, the population Who is coming's foot counts.
        let unanswered: HomeLine? = if let raid, raid.notAnsweredCount > 0 { unansweredLine(raid.notAnsweredCount, raid: raid) } else { nil }
        let applicants: HomeLine? = badge.applicants > 0 ? HomeLine(
            title: Copy.homeOfficerApplicants(n: badge.applicants),
            sub: inputs.applicants.first.map { applicant in
                Copy.homeOfficerApplicant(
                    name: applicant.name, kind: applicant.kind, level: applicant.level,
                    age: RelativeAge(from: applicant.appliedAt, to: now).phrase
                )
            } ?? ""
        ) : nil
        guard unanswered != nil || applicants != nil else { return nil }
        return OfficerContent(unanswered: unanswered, applicants: applicants)
    }

    private func unansweredLine(_ count: Int, raid: RaidView) -> HomeLine {
        guard let short = RaidShortRole.of(raid) else { return HomeLine(title: Copy.homeOfficerUnanswered(n: count), sub: "") }
        let ofRole = raid.notSigned.filter { $0.role == short.role }.count
        let title = ofRole > 0
            ? Copy.homeOfficerUnansweredRole(n: count, m: ofRole, role: short.lowerWord)
            : Copy.homeOfficerUnanswered(n: count)
        return HomeLine(title: title, sub: Copy.homeOfficerShort(role: short.word, n: short.missing))
    }

    static func roleWord(_ role: String?) -> String {
        switch role {
        case "tank": Copy.raidRoleTank
        case "healer": Copy.raidRoleHealer
        default: Copy.homeRaidDamage
        }
    }

    // MARK: Raid live

    func raidLive(_ inputs: HomeFlowInputs) -> RaidLiveContent? {
        guard let raid = inputs.raid, let next = nextRaid(inputs) else { return nil }
        let counts = Copy.homeRaidCounts(in: raid.counts.signedIn, late: raid.counts.late)
        let age = RelativeAge(from: raid.startsAt, to: now).phrase
        let named = Self.namedCharacter(raid, key: inputs.character.key)
        let line: String = switch (next.selected, named) {
        case (nil, _): Copy.homeLiveLineNoAnswer(age: age, counts: counts)
        case let (state?, nil): Copy.homeLiveLineAlone(age: age, counts: counts, answer: state.label)
        case let (state?, who?): Copy.homeLiveLine(age: age, counts: counts, answer: state.label, name: who.name, role: Self.roleWord(who.role))
        }
        return RaidLiveContent(eyebrow: Copy.homeRaidLive(guild: inputs.character.guild ?? "", raid: raid.zoneName), line: line)
    }

    // MARK: Last night

    func lastNight(_ night: RaidNightSummary) -> LastNightContent {
        let loot = night.loot.map { item in
            LastNightLoot(
                look: HomeItemLook(id: item.itemId, name: item.name, quality: item.quality, icon: item.icon, itemLevel: nil),
                line: lootLine(item)
            )
        }
        // The report row needs the report; its line names the best percentile when the night has one.
        let report: HomeLine? = night.reportId.map { _ in
            var parts: [String] = []
            if let best = night.bests.first, let percentile = best.percentile {
                let rank = ReportsCopy.ordinal(percentile)
                parts.append((night.newBests ?? 0) > 0
                    ? Copy.homeLastNightPercentile(rank: rank, boss: best.boss)
                    : Copy.nightPercentile(ordinal: rank))
            }
            if let notes = night.coachNotes, !notes.isEmpty { parts.append(Copy.homeLastNightCoach(n: notes.count)) }
            return HomeLine(title: Copy.homeLastNightReport(zone: night.zoneName), sub: parts.joined(separator: " · "))
        }
        let simAgain = night.loot.first.map {
            HomeLine(title: Copy.homeLastNightSimAgain(item: $0.name), sub: Copy.homeLastNightSimAgainLine)
        }
        return LastNightContent(loot: loot, report: report, reportIcon: night.zoneIcon ?? RaidIcons.generic, reportID: night.reportId, simAgain: simAgain)
    }

    /// `You got it from Onyxia · main hand, +12.6 DPS over your worn`; the slot alone without a gain.
    func lootLine(_ item: RaidNightLoot) -> String {
        let slot = item.slot.lowercased()
        guard let gain = item.gain, gain > 0 else { return Copy.homeLastNightLoot(boss: item.boss, slot: slot) }
        let figure = Copy.homeUpgradesGain(figure: ReportsCopy.signedHouseFigure(gain), unit: item.unit ?? "")
        return Copy.homeLastNightLootGain(boss: item.boss, slot: slot, gain: figure)
    }

    // MARK: New for you, Find a guild

    func newForYou(_ page: ActivityPage, query: HomeActivityQuery) -> NewForYouContent {
        let rows = query.rows(page.rows).map { row in
            let day = ActivityDay.of(row.occurredAt, now: now, calendar: calendar)
            let when = ActivitySections.when(row.occurredAt, day: day, now: now, calendar: calendar, locale: Locale(identifier: "en_US"))
            // Row 3: Activity's bucket (`Sat`) on its own meta line, never `· 3 d ago` after the body.
            return NewForYouRow(row: row, line: row.body, when: when)
        }
        return NewForYouContent(rows: rows, unread: page.unread)
    }

    func findGuild(_ firstRaids: Date?) -> String {
        guard let firstRaids, firstRaids > now else { return Copy.homeFindGuildLine }
        return Copy.homeFindGuildFirstRaids(date: clock.shortDate(firstRaids))
    }
}

/// The role furthest under its comp target (F3: `healers short by 3`).
struct RaidShortRole: Equatable {
    let role: String
    let missing: Int

    var word: String {
        switch role {
        case "tank": Copy.raidGroupTanks
        case "healer": Copy.raidGroupHealers
        default: Copy.homeRaidDamage
        }
    }

    var lowerWord: String {
        switch role {
        case "tank": Copy.homeRoleTanks
        case "healer": Copy.homeRoleHealers
        default: Copy.homeRoleDamage
        }
    }

    /// Tanks and healers first: a raid short of damage is rarely held up by it.
    static func of(_ raid: RaidView) -> RaidShortRole? {
        guard let target = raid.compTarget, let short = RaidRoleCounts.of(raid).short(target) else { return nil }
        return RaidShortRole(role: short.role, missing: short.missing)
    }
}
