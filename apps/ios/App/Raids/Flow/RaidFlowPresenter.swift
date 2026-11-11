import Auth
import DesignSystem
import Foundation
import Models

/// Turns the raid calls into what the Raid tab and its screens draw (flows spec 4.2, F2, F3).
/// Pure, so every mocked state is a test of content values.
struct RaidFlowPresenter {
    let now: Date
    var clock = RaidClock()

    static let maxLevel = 60
    /// A role row names everyone up to this many; past it, the first `namedWhenLong` and `and N more`.
    static let namesInFull = 5
    static let namedWhenLong = 3

    private var home: HomeFlowPresenter { HomeFlowPresenter(now: now, clock: clock) }

    // MARK: Words

    /// `Mirethil, Selka and Aurelith`.
    static func list(_ names: [String]) -> String {
        guard names.count > 1, let last = names.last else { return names.first ?? "" }
        return Copy.raidListAnd(first: names.dropLast().joined(separator: ", "), last: last)
    }

    /// `Holy Priest`, from the spec's name and the class slug.
    static func kind(spec: String?, className: String?) -> String {
        let classWord = className.map(HomePresenter.words)
        // Some rows name the class in the spec already (`Holy Paladin`).
        if let spec, let classWord, spec.lowercased().hasSuffix(classWord.lowercased()) { return spec }
        return [spec, classWord].compactMap(\.self).filter { !$0.isEmpty }.joined(separator: " ")
    }

    func when(_ start: Date, zone: TimeZone?) -> String {
        Copy.homeRaidWhen(date: clock.shortDate(start), time: clock.guildClock(start, guildZone: zone ?? clock.zone))
    }

    // MARK: The root

    /// Under 60: `26 levels to go`.
    static func leveler(level: Int) -> RaidRootContent {
        .leveler(toGo: Copy.raidLevelerToGo(n: max(1, maxLevel - level)))
    }

    /// No guild at 60: the guilds recruiting the spec, with the article the spec's name takes.
    static func recruiting(_ guilds: [RecruitingGuild], kind: String) -> RecruitingContent {
        let vowel = kind.first.map { "AEIOU".contains($0) } ?? false
        return RecruitingContent(
            heading: vowel ? Copy.raidRecruitHeadingVowel(kind: kind) : Copy.raidRecruitHeading(kind: kind),
            rows: guilds.map { guild in
                RecruitingRow(
                    guildID: guild.guildId, name: guild.name,
                    line: Copy.raidRecruitLine(faction: guild.faction.capitalized, schedule: guild.schedule, needs: needsWord(guild.needs))
                )
            }
        )
    }

    static func needsWord(_ needs: String) -> String {
        switch needs {
        case "tank": Copy.homeRoleTanks
        case "healer": Copy.homeRoleHealers
        default: Copy.homeRoleDamage
        }
    }

    /// The guild's pages; a figure the app does not have leaves its row's line out.
    func pages(
        guild: String, members: Int?, atCap: Int?, planned: [RaidView], badge: RaidBadge?, lastNight: RaidNightRow? = nil
    ) -> GuildPagesContent {
        let roster = members.map { n in
            let words = Copy.raidPagesMembers(n: n)
            return atCap.map { Copy.raidPagesMembersAtCap(members: words, n: $0) } ?? words
        }
        let upcoming = planned.filter { $0.startsAt > now }.sorted { $0.startsAt < $1.startsAt }
        let calendar = upcoming.isEmpty ? nil : upcoming.map { raid in
            Copy.raidPagesCalendarItem(zone: raid.title, weekday: String(clock.weekday(raid.startsAt).prefix(3)))
        }.joined(separator: ", ")
        return GuildPagesContent(guild: guild, rows: [
            GuildPageRow(page: .roster, sub: roster),
            GuildPageRow(page: .calendar, sub: calendar),
            GuildPageRow(page: .loot, sub: lootLine(badge: badge, lastNight: lastNight)),
            GuildPageRow(page: .sims, sub: badge?.simsRank.map { Copy.raidPagesSimsRank(position: $0.position, of: $0.of) }),
        ])
    }

    /// Loot history's line: `Last: Onyxia's Lair, 1 kill, 2 wipes`, then the last raid's pieces.
    func lootLine(badge: RaidBadge?, lastNight: RaidNightRow?) -> String? {
        let night = lastNight.map { Copy.raidPagesLastNight(zone: $0.zone, kills: Copy.raidPagesKills(n: $0.killCount), wipes: Copy.raidPagesWipes(n: $0.wipeCount)) }
        let parts = [night, badge?.lastRaidLoot.map { Copy.raidPagesLootLast(n: $0.count) }].compactMap(\.self)
        return parts.isEmpty ? nil : parts.joined(separator: Self.separator)
    }

    /// The officer's Needs you (F3): the badge's open count over the unanswered and the applicants.
    func needsYou(_ inputs: HomeFlowInputs, raid: RaidView?) -> NeedsYouContent? {
        guard let badge = inputs.badge, badge.officer, badge.count > 0 else { return nil }
        let items = home.officer(inputs, raid: raid)
        guard let items else { return nil }
        return NeedsYouContent(meta: Copy.raidNeedsYouOpen(n: badge.count), items: items)
    }

    /// The next raid's card on the Raid tab: Home's card without the officer's items (Needs you
    /// holds them).
    func nextRaid(_ inputs: HomeFlowInputs) -> NextRaidContent? {
        var cardInputs = inputs
        cardInputs.badge = nil
        return home.nextRaid(cardInputs)
    }

    /// Raid night on the root: the live card and who is in, tanks and healers named.
    func nightRoot(_ inputs: HomeFlowInputs, night: RaidNightSummary) -> RaidNightRootContent? {
        guard let live = home.raidLive(inputs) else { return nil }
        let target = inputs.raid?.compTarget
        let groups: [(label: String, people: [RaidNightPerson], want: Int?)] = [
            (Copy.raidGroupTanks, night.coming.tanks, target?.tanks),
            (Copy.raidGroupHealers, night.coming.healers, target?.healers),
            (Copy.homeRaidDamage, night.coming.damage, target.map { $0.melee + $0.ranged }),
        ]
        let roles = groups.map { group in
            Self.roleLine(group.label, people: group.people.map { ($0.name, $0.className, $0.state, $0.lateMinutes) }, want: group.want)
        }
        return RaidNightRootContent(live: RaidLiveContent(eyebrow: live.eyebrow, line: live.line), roles: roles)
    }

    /// `Tanks 2 of 2` over the names; a late raider's minutes on the root, the gap on the raid page.
    func roleLine(_ label: String, _ people: [RaidNightPerson], want: Int?, lateShown: Bool, short: Bool) -> HomeLine {
        let names = people.map { person -> String in
            guard lateShown, person.state == SignupState.late.rawValue else { return person.name }
            return person.lateMinutes.map { Copy.raidNightLateName(name: person.name, n: $0) } ?? Copy.raidNightLateNameOnly(name: person.name)
        }
        var line = names.count > Self.namesInFull
            ? Copy.raidNightMore(names: names.prefix(Self.namedWhenLong).joined(separator: ", "), n: names.count - Self.namedWhenLong)
            : names.joined(separator: ", ")
        if short, let want, want > people.count { line = Copy.raidNightShort(names: line, n: want - people.count) }
        let title = want.map { Copy.raidNightRole(role: label, n: people.count, want: $0) } ?? label
        return HomeLine(title: title, sub: line)
    }

    /// Who is coming (spec 4.2): tanks, healers and damage with the names of those In or Late (yours
    /// first, a Late with its minutes), the count against the comp, and the short role's gap.
    func coming(_ raid: RaidView) -> [RoleLineContent] {
        let target = raid.compTarget
        let coming = raid.signups
            .filter { SignupState(rawValue: $0.state)?.isComing == true }
            .sorted { $0.own && !$1.own }
        let groups: [(wire: String, label: String, want: Int?)] = [
            ("tank", Copy.raidGroupTanks, target?.tanks),
            ("healer", Copy.raidGroupHealers, target?.healers),
            ("dps", Copy.homeRaidDamage, target.map { $0.melee + $0.ranged }),
        ]
        let counts = RaidRoleCounts.of(raid)
        return groups.compactMap { group in
            let people = coming.filter { Self.roleGroup($0.role) == group.wire }
            guard !people.isEmpty || group.want != nil else { return nil }
            var line = Self.roleLine(group.label, people: people.map { ($0.name, $0.className, $0.state, $0.lateMinutes) }, want: group.want)
            // The split is the count model's, so `Damage 15` always reads `7 melee, 8 ranged` (row 1).
            if group.wire == RaidRoleCounts.dps, counts.damage > 0 {
                line.split = Copy.raidTabDamageSplit(melee: counts.melee, ranged: counts.ranged)
            }
            return line
        }
    }

    static let nameSeparator = ", "

    /// The signup's `group` for a raider who fights in melee.
    static let meleeGroup = RaidRoleCounts.meleeGroup

    /// One role line (S6): the tally, the names in class colour, a Late's minutes.
    static func roleLine(
        _ label: String, people: [(name: String, className: String?, state: String, lateMinutes: Int?)], want: Int?
    ) -> RoleLineContent {
        let persons = people.map { person in
            RolePerson(name: person.name, className: person.className, text: lateText(person.name, state: person.state, minutes: person.lateMinutes))
        }
        let long = persons.count > namesInFull
        return RoleLineContent(
            role: label, have: people.count, want: want,
            names: names(people.map { ($0.name, $0.state, $0.lateMinutes) }),
            people: long ? Array(persons.prefix(namedWhenLong)) : persons,
            more: long ? persons.count - namedWhenLong : 0
        )
    }

    private static func lateText(_ name: String, state: String, minutes: Int?) -> String {
        guard state == SignupState.late.rawValue, let minutes else { return name }
        return Copy.raidTabLateNote(name: name, n: minutes)
    }

    // MARK: The head's facts and Who is coming's foot

    /// `Invites 7:45 pm · 40`, when the raid names both.
    func invites(_ raid: RaidView) -> String? {
        guard let invite = raid.inviteAt, let size = raid.size else { return nil }
        let zone = TimeZone(identifier: raid.guildTimezone) ?? clock.zone
        return Copy.raidTabInviteSize(time: clock.guildClock(invite, guildZone: zone), n: size)
    }

    /// `Bench 1 · Out 4 · 14 not answered`; a count the raid leaves out is left out.
    func answers(_ raid: RaidView) -> String? {
        let unanswered = raid.notAnsweredCount
        let parts = [
            raid.counts.bench.map(Copy.raidTabCountBench),
            raid.counts.signedOut.map(Copy.raidTabCountOut),
            unanswered > 0 ? Copy.raidTabCountNotAnswered(n: unanswered) : nil,
        ].compactMap(\.self)
        return parts.isEmpty ? nil : parts.joined(separator: Self.separator)
    }

    /// G4 and row 4, one line a raid leader reads at a glance: `Missing` once, the gaps nobody brings
    /// first, then the short ones, by the data's short label (`Fort`, `WF`) else the name, at most
    /// `coverNames` and `and N more` (`Missing Fort · WF 1 of 2`). Nil when all is covered.
    func cover(_ raid: RaidView) -> RaidCoverContent? {
        let all = Self.coverGaps(raid)
        let gaps = all.filter { $0.have == nil } + all.filter { $0.have != nil }
        guard !gaps.isEmpty else { return nil }
        let shown = gaps.prefix(Self.coverNames)
        return RaidCoverContent(
            missing: shown.filter { $0.have == nil },
            short: shown.filter { $0.have != nil },
            more: gaps.count - shown.count
        )
    }

    /// The cover rows that are short or missing (In and Late against need), in the cover's order; a
    /// raid without cover rows falls back to the buffs no one brings. A gap takes the buff's short
    /// label when the composition has one for its key.
    static func coverGaps(_ raid: RaidView) -> [CoverGap] {
        let buffs = raid.composition?.buffs ?? []
        let labels = Dictionary(buffs.map { ($0.key, $0.label) }, uniquingKeysWith: { first, _ in first })
        let rows = raid.cover ?? []
        guard !rows.isEmpty else {
            return buffs.filter { !$0.present }.map { CoverGap(name: label($0.label, else: $0.name), icon: $0.icon) }
        }
        return rows.compactMap { row in
            let name = label(labels[row.key], else: row.name)
            return switch row.state {
            case coverMissing: CoverGap(name: name, icon: row.icon)
            case coverShort: CoverGap(name: name, icon: row.icon, have: row.providersIn + row.providersLate, need: row.need)
            default: nil
            }
        }
    }

    private static func label(_ label: String?, else name: String) -> String {
        guard let label, !label.isEmpty else { return name }
        return label
    }

    /// The most names the Covered line spells out before `and N more`.
    static let coverNames = 3
    static let list = ", "
    static let coverMissing = "missing"
    static let coverShort = "short"
    static let separator = " · "

    /// Whether the addon synced each roster character this week (round 4, s5): `logged_at` is the
    /// addon's last upload, read against the one stale rule; `logged_recently` is only the last 24 h.
    func syncedThisWeek(_ roster: [GuildRosterRow]) -> [String: Bool] {
        Dictionary(roster.map { ($0.characterKey, !SyncAge.isStale($0.loggedAt, now: now)) }) { first, _ in first }
    }

    /// G1 and row 5, for officers: the signed In or Late the roster has not seen sync this week, by
    /// name (`Not synced this week: Kaelra, Thalrin, Brakka and 1 more`); the tap opens the roster
    /// filtered to them.
    func notSynced(_ raid: RaidView, synced: [String: Bool]?) -> NotSyncedContent? {
        guard let synced else { return nil }
        let stale = raid.signups.filter { signup in
            SignupState(rawValue: signup.state)?.isComing == true && synced[signup.characterKey] == false
        }
        guard !stale.isEmpty else { return nil }
        let names = stale.map(\.name)
        let shown = names.count > Self.namedWhenLong + 1
            ? Copy.raidTabAndMore(names: names.prefix(Self.namedWhenLong).joined(separator: Self.list), n: names.count - Self.namedWhenLong)
            : names.joined(separator: Self.list)
        return NotSyncedContent(line: Copy.raidTabNotSyncedNames(names: shown), keys: stale.map(\.characterKey))
    }

    /// The officers' notes, trimmed; nil when there are none.
    static func notes(_ raid: RaidView) -> String? {
        let text = raid.notes.trimmingCharacters(in: .whitespacesAndNewlines)
        return text.isEmpty ? nil : text
    }

    /// Up to five names in full; more read `Thandrel, Kaelra, Thalrin and 8 more`.
    static func names(_ people: [(name: String, state: String, lateMinutes: Int?)]) -> String {
        let words = people.map { person in
            guard person.state == SignupState.late.rawValue, let minutes = person.lateMinutes else { return person.name }
            return Copy.raidTabLateNote(name: person.name, n: minutes)
        }
        guard words.count > namesInFull else { return words.joined(separator: ", ") }
        return Copy.raidTabAndMore(names: words.prefix(namedWhenLong).joined(separator: ", "), n: words.count - namedWhenLong)
    }

    static func roleGroup(_ role: String) -> String { RaidRoleCounts.group(role) }

    /// The Ready check (F2): every item, the export's age; nil without an export.
    func ready(_ check: ReadyCheck?) -> ReadyListContent? {
        guard let check, !check.items.isEmpty else { return nil }
        let rows = check.items.map { item in
            ReadyRowContent(
                key: item.key, mark: ReadyRowContent.Mark(rawValue: item.state) ?? .warn, title: item.title, line: item.line,
                fix: item.state == ReadyCheck.okState ? nil : item.whereToGet?.label,
                link: item.whereToGet?.link.flatMap(URL.init(string:)),
                icon: item.icon.flatMap { $0.isEmpty ? nil : $0 }
            )
        }
        return ReadyListContent(rows: rows, stamp: check.exportedAt.map { Copy.raidReadyStamp(age: RelativeAge(from: $0, to: now).phrase) })
    }

    // MARK: The raid page on raid night

    func nightPage(_ night: RaidNightSummary, raid: RaidView?, lateOpen: Bool, answer: SignupState?) -> RaidNightPageContent {
        let zone = raid.flatMap { TimeZone(identifier: $0.guildTimezone) }
        let time = when(night.startsAt, zone: zone)
        let selected = answer ?? night.you.state.flatMap(SignupState.init(rawValue:))
        let target = raid?.compTarget
        return RaidNightPageContent(
            title: night.zoneName,
            icon: night.zoneIcon ?? RaidIcons.generic,
            facts: [night.guildName, time],
            eyebrow: Copy.homeRaidTonight(guild: night.guildName, raid: night.zoneName),
            clock: home.clockParts(to: night.startsAt),
            aside: Copy.raidNightAside(when: time, counts: Copy.homeRaidCounts(in: night.coming.inCount, late: night.coming.late)),
            tag: selected.map { Copy.raidNightYouAre(state: $0.label) } ?? Copy.raidChipNotSigned,
            answered: selected != nil,
            selected: selected,
            lateOpen: lateOpen && selected == .late,
            roles: [
                roleLine(Copy.raidGroupTanks, night.coming.tanks, want: target?.tanks, lateShown: false, short: true),
                roleLine(Copy.raidGroupHealers, night.coming.healers, want: target?.healers, lateShown: false, short: true),
                roleLine(Copy.homeRaidDamage, night.coming.damage, want: target.map { $0.melee + $0.ranged }, lateShown: false, short: false),
            ],
            ready: ready(night.ready)
        )
    }

    // MARK: The applicant

    /// `cloak and boots` from the API's slot keys.
    static func slotWord(_ slot: String) -> String {
        switch slot {
        case "back": Copy.raidApplicantSlotBack
        case "feet": Copy.raidApplicantSlotFeet
        case "hands": Copy.raidApplicantSlotHands
        case "wrist": Copy.raidApplicantSlotWrist
        case "chest": Copy.raidApplicantSlotChest
        case "legs": Copy.raidApplicantSlotLegs
        case "shoulder", "shoulders": Copy.raidApplicantSlotShoulder
        case "head": Copy.raidApplicantSlotHead
        case "main_hand", "mainhand", "weapon", "two_hand": Copy.raidApplicantSlotWeapon
        default: slot.replacingOccurrences(of: "_", with: " ")
        }
    }

    /// The applicant's sheet; the head names faction and race (`Alliance Dwarf`) when the API does.
    func applicant(_ application: GuildOfficerApplication, guild: String, accepted: Bool) -> ApplicantContent {
        let character = application.character
        let kind = Self.kind(spec: character.specName, className: character.className)
        let facts = [
            character.level.map { Copy.raidApplicantKind(kind: kind, level: $0) } ?? kind,
            character.origin,
            Copy.raidApplicantApplied(age: RelativeAge(from: application.createdAt, to: now).phrase),
        ].compactMap(\.self).filter { !$0.isEmpty }
        let sheet = application.sheet
        var figures: [ApplicantFigure] = []
        if let score = sheet?.score { figures.append(ApplicantFigure(label: Copy.raidApplicantScore, value: "\(score)", unit: nil)) }
        if let sim = sheet?.sim {
            figures.append(ApplicantFigure(label: Copy.raidApplicantSim, value: "\(Int(sim.value.rounded()))", unit: Codes.unit(sim.unit)))
        }
        if let level = character.itemLevel { figures.append(ApplicantFigure(label: Copy.raidApplicantItemLevel, value: "\(level)", unit: nil)) }
        return ApplicantContent(
            name: character.name, className: character.className, facts: facts, figures: figures,
            ready: applicantReady(sheet, kind: kind),
            history: (sheet?.history ?? []).map { row in
                HomeLine(
                    title: row.guildName,
                    sub: Copy.raidApplicantHistoryLine(raids: Copy.raidApplicantRaids(n: row.raids), date: clock.dayMonth(row.lastRaidAt))
                )
            },
            note: application.message.isEmpty ? nil : application.message,
            accepted: accepted ? Copy.raidApplicantAccepted(guild: guild, name: character.name) : nil,
            invite: Copy.raidApplicantGinvite(name: character.name),
            after: Copy.raidApplicantAfter(name: character.name)
        )
    }

    private func applicantReady(_ sheet: ApplicantSheet?, kind: String) -> [HomeLine] {
        guard let sheet else { return [] }
        var lines: [HomeLine] = []
        if let attune = sheet.attunement {
            let title = switch attune.state {
            case "ok": Copy.raidApplicantAttuned(zone: attune.zoneName)
            case "miss": Copy.raidApplicantNotAttuned(zone: attune.zoneName)
            default: Copy.raidApplicantAttuneUnknown(zone: attune.zoneName)
            }
            lines.append(HomeLine(title: title, sub: attune.item?.name ?? ""))
        }
        if let enchants = sheet.enchants {
            let line = Copy.raidApplicantEnchantsLine(n: enchants.enchanted, slots: enchants.slots)
            let sub = enchants.missing.isEmpty
                ? line
                : Copy.raidApplicantEnchantsMissing(line: line, missing: Self.list(enchants.missing.map(Self.slotWord)))
            lines.append(HomeLine(title: Copy.raidApplicantEnchants, sub: sub))
        }
        if let gap = sheet.bisGap {
            lines.append(HomeLine(
                title: Copy.raidApplicantGap,
                sub: Copy.raidApplicantGapLine(pieces: Copy.raidApplicantPieces(n: gap.pieces), list: gap.list)
            ))
        }
        return lines
    }
}
