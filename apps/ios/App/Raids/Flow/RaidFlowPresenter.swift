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
    func pages(guild: String, members: Int?, atCap: Int?, planned: [RaidView], badge: RaidBadge?) -> GuildPagesContent {
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
            GuildPageRow(page: .loot, sub: badge?.lastRaidLoot.map { Copy.raidPagesLootLast(n: $0.count) }),
            GuildPageRow(page: .sims, sub: badge?.simsRank.map { Copy.raidPagesSimsRank(position: $0.position, of: $0.of) }),
        ])
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
            let names = Self.names(group.people.map { ($0.name, $0.state, $0.lateMinutes) })
            let gap = group.want.map { $0 - group.people.count }.flatMap { $0 > 0 ? $0 : nil }
            return RoleLineContent(
                role: group.label, have: group.people.count, want: group.want,
                names: gap.map { Copy.raidTabShort(names: names, n: $0) } ?? names
            )
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
        let short = RaidShortRole.of(raid)
        let coming = raid.signups
            .filter { SignupState(rawValue: $0.state)?.isComing == true }
            .sorted { $0.own && !$1.own }
        let groups: [(wire: String, label: String, want: Int?)] = [
            ("tank", Copy.raidGroupTanks, target?.tanks),
            ("healer", Copy.raidGroupHealers, target?.healers),
            ("dps", Copy.homeRaidDamage, target.map { $0.melee + $0.ranged }),
        ]
        return groups.compactMap { group in
            let people = coming.filter { Self.roleGroup($0.role) == group.wire }
            guard !people.isEmpty || group.want != nil else { return nil }
            let names = Self.names(people.map { ($0.name, $0.state, $0.lateMinutes) })
            let gap = short.flatMap { $0.role == group.wire ? $0.missing : nil }
            return RoleLineContent(
                role: group.label, have: people.count, want: group.want,
                names: gap.map { Copy.raidTabShort(names: names, n: $0) } ?? names
            )
        }
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

    static func roleGroup(_ role: String) -> String {
        switch role {
        case "tank", "healer": role
        default: "dps"
        }
    }

    /// The Ready check (F2): every item, the export's age; nil without an export.
    func ready(_ check: ReadyCheck?) -> ReadyListContent? {
        guard let check, !check.items.isEmpty else { return nil }
        let rows = check.items.map { item in
            ReadyRowContent(
                key: item.key, mark: ReadyRowContent.Mark(rawValue: item.state) ?? .warn, title: item.title, line: item.line,
                fix: item.state == ReadyCheck.okState ? nil : item.whereToGet?.label,
                link: item.whereToGet?.link.flatMap(URL.init(string:))
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
