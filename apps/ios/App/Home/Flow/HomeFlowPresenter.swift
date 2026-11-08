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
    /// The answer just sent, shown before the raid reloads.
    var answer: SignupState?
    var lateOpen = false
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
            facts.append(Copy.homeHeadLevel(n: level, kind: kind))
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
                source: "\(upgrade.slotLabel) · \(upgrade.pick.sourceLine)",
                gain: upgrade.modelled ? upgrade.gain.map { HomeGain(figure: ItemNumbers.signedTenths($0), unit: unit) } : nil,
                wornID: upgrade.worn?.id,
                sourceLine: upgrade.pick.sourceLine
            )
        }
        let note: String? = rows.isEmpty ? (home.state == "no_gear" ? Copy.homeAnswerNoGear : Copy.homeAnswerAllMatch(band: home.bandLabel ?? "")) : nil
        return HomeUpgradesContent(rows: Array(rows), bestRun: bestRun.map { bestRunLine($0, unit: unit) }, note: note)
    }

    func bestRunLine(_ run: BestRun, unit: String) -> String {
        let gain = Copy.homeUpgradesGain(figure: ItemNumbers.signedTenths(run.gain), unit: unit)
        guard let chance = run.chance else { return Copy.homeBestRun(place: run.place, item: run.item, gain: gain) }
        return Copy.homeBestRunChance(place: run.place, item: run.item, gain: gain, chance: "\(Int(chance.rounded()))%")
    }

    // MARK: The next raid

    func nextRaid(_ inputs: HomeFlowInputs) -> NextRaidContent? {
        guard let raid = inputs.raid else { return nil }
        let viewer = raid.viewer.characters.first { $0.characterKey == inputs.character.key } ?? raid.viewer.characters.first
        let who = Copy.homeRaidAs(name: viewer?.name ?? inputs.character.name, role: Self.roleWord(viewer?.role))
        let selected = inputs.answer ?? viewer?.state.flatMap(SignupState.init(rawValue:))
        let zone = TimeZone(identifier: raid.guildTimezone) ?? clock.zone
        let time = Copy.homeRaidWhen(date: clock.shortDate(raid.startsAt), time: clock.guildClock(raid.startsAt, guildZone: zone))
        let counts = Copy.homeRaidCounts(n: raid.counts.signedIn, late: raid.counts.late)
        let answer = selected.map { Copy.homeRaidAnswered(state: $0.label, who: who) } ?? Copy.homeRaidNotAnswered(who: who)
        return NextRaidContent(
            guildID: inputs.character.guildID ?? 0, raidID: raid.id, characterKey: viewer?.characterKey ?? inputs.character.key,
            eyebrow: Copy.homeRaidTonight(guild: inputs.character.guild ?? "", zone: raid.zoneName),
            clock: clockParts(to: raid.startsAt),
            when: time,
            line: HomeLine(
                title: Copy.homeRaidLineTitle(zone: raid.zoneName, when: time),
                sub: selected == nil ? Copy.homeRaidLineOpen(counts: counts, who: who) : answer
            ),
            icon: RaidIcons.zone(raid.zone),
            answer: answer,
            selected: selected,
            lateOpen: inputs.lateOpen && selected == .late,
            tally: tally(raid),
            ready: ready(inputs.ready),
            officer: officer(inputs, raid: raid)
        )
    }

    /// `02 hours`, `48 min` until the pull; none once it has started.
    func clockParts(to start: Date) -> [FSClockPart] {
        let seconds = Int(start.timeIntervalSince(now))
        guard seconds > 0 else { return [] }
        let hours = seconds / 3600
        let minutes = (seconds % 3600) / 60
        return [
            FSClockPart(String(format: "%02d", hours), Copy.homeRaidHours),
            FSClockPart(String(format: "%02d", minutes), Copy.homeRaidMinutes),
        ]
    }

    /// Who is coming against the comp target: tanks, healers, damage.
    func tally(_ raid: RaidView) -> [FSTally.Item] {
        guard let target = raid.compTarget else { return [] }
        return [
            FSTally.Item(label: Copy.raidGroupTanks, have: raid.counts.tanks, want: target.tanks),
            FSTally.Item(label: Copy.raidGroupHealers, have: raid.counts.healers, want: target.healers),
            FSTally.Item(label: Copy.homeRaidDamage, have: raid.counts.dps, want: target.melee + target.ranged),
        ]
    }

    func ready(_ check: ReadyCheck?) -> HomeLine? {
        guard let check, !check.open.isEmpty else { return nil }
        let items = check.open.map(\.line).joined(separator: "; ")
        let line = check.exportedAt.map { Copy.homeReadyLine(items: items, age: RelativeAge(from: $0, to: now).phrase) } ?? items
        return HomeLine(title: Copy.homeReadyTitle(n: check.open.count), sub: line)
    }

    /// The officer's open items (F3): who has not answered with the short role, and the applicants.
    func officer(_ inputs: HomeFlowInputs, raid: RaidView) -> OfficerContent? {
        guard let badge = inputs.badge, badge.officer else { return nil }
        let unanswered: HomeLine? = badge.unanswered > 0 ? unansweredLine(badge.unanswered, raid: raid) : nil
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
        let counts = Copy.homeRaidCounts(n: raid.counts.signedIn, late: raid.counts.late)
        let viewer = raid.viewer.characters.first { $0.characterKey == next.characterKey }
        let who = Copy.homeRaidAs(name: viewer?.name ?? inputs.character.name, role: Self.roleWord(viewer?.role))
        let state = next.selected?.label ?? Copy.raidChipNotSigned
        return RaidLiveContent(
            eyebrow: Copy.homeRaidInProgress(guild: inputs.character.guild ?? "", zone: raid.zoneName),
            line: Copy.homeRaidPulled(span: RaidClock.span(max(0, now.timeIntervalSince(raid.startsAt))), counts: counts, state: state, who: who)
        )
    }

    // MARK: Last night

    func lastNight(_ night: RaidNightSummary) -> LastNightContent {
        let loot = night.loot.map { item in
            LastNightLoot(
                look: HomeItemLook(id: item.itemId, name: item.name, quality: item.quality, icon: item.icon, itemLevel: nil),
                line: Copy.homeLastNightLoot(boss: item.boss, slot: item.slot.lowercased())
            )
        }
        let report: HomeLine? = night.bests.first.map { best in
            let rank = ReportsCopy.ordinal(best.percentile)
            var parts = [night.newBests > 0
                ? Copy.homeLastNightPercentile(rank: rank, boss: best.boss)
                : Copy.nightPercentile(ordinal: rank)]
            if let notes = night.coachNotes, !notes.isEmpty { parts.append(Copy.homeLastNightCoach(n: notes.count)) }
            return HomeLine(title: Copy.homeLastNightReport(zone: night.zoneName), sub: parts.joined(separator: " · "))
        }
        let simAgain = night.loot.first.map {
            HomeLine(title: Copy.homeLastNightSimAgain(item: $0.name), sub: Copy.homeLastNightSimAgainLine)
        }
        return LastNightContent(loot: loot, report: report, reportIcon: RaidIcons.zoneName(night.zoneName), reportID: night.reportId, simAgain: simAgain)
    }

    // MARK: New for you, Find a guild

    func newForYou(_ page: ActivityPage, query: HomeActivityQuery) -> NewForYouContent {
        let rows = query.rows(page.rows).map { row in
            let day = ActivityDay.of(row.occurredAt, now: now, calendar: calendar)
            let when = ActivitySections.when(row.occurredAt, day: day, now: now, calendar: calendar, locale: Locale(identifier: "en_US"))
            return NewForYouRow(row: row, line: Copy.homeNewForYouLine(body: row.body, when: when))
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
        guard let target = raid.compTarget else { return nil }
        let gaps = [
            RaidShortRole(role: "tank", missing: target.tanks - raid.counts.tanks),
            RaidShortRole(role: "healer", missing: target.healers - raid.counts.healers),
            RaidShortRole(role: "dps", missing: target.melee + target.ranged - raid.counts.dps),
        ]
        return gaps.first { $0.missing > 0 }
    }
}
