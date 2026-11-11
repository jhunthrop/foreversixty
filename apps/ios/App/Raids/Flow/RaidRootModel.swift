import Foundation
import Models
import Networking
import Observation

/// What the Raid tab's root reads beyond `/v1/me` and the guild's raids.
struct RaidRootExtras: Equatable {
    var badge: RaidBadge?
    var applicants: [ApplicantLine] = []
    var ready: ReadyCheck?
    /// The live raid's night (who is in by role, the Ready check).
    var night: RaidNightSummary?
    var recruiting: RecruitingGuilds?
    /// An officer's setup checklist (the Guild tab's Setup row).
    var setup: GuildSetup?
    /// When the player applied to each recruiting guild still deciding, by guild id (spec 2.4).
    var applied: [Int64: Date] = [:]
    /// The newest loot night (the Guild tab's Loot section).
    var lootNights: LootNights?
}

/// Composes the Raid tab's root (flows spec 4.2) from the tab's state and its extras.
struct RaidRootComposer {
    let presenter: RaidFlowPresenter
    /// The roster's figures for the Roster row, when the guild's roster is loaded.
    var members: Int?
    var atCap: Int?
    /// G1: the guild roster's `logged_recently` by character key, when the roster is loaded.
    var synced: [String: Bool]?
    /// The head's answer (F1): pending, Late's minutes, the status line, offline.
    var answering = RaidAnswering()
    /// The guild home the Guild tab's sections read (the roster store's), when it is loaded.
    var home: GuildHome?
    /// The first raids' date (the bundled key dates), for Progression before the first pull.
    var firstRaids: Date?

    func content(_ state: RaidsTabState, extras: RaidRootExtras) -> RaidRootContent {
        switch state {
        case .loading: return .loading
        case .failed: return .failed
        case .signedOut: return .signedOut
        case let .noGuild(character):
            let kind = RaidFlowPresenter.kind(spec: character?.spec, className: character?.className)
            var recruiting = RaidFlowPresenter.recruiting(extras.recruiting?.rows ?? [], kind: kind)
            recruiting.rows = recruiting.rows.map { row in
                var row = row
                row.applied = extras.applied[row.guildID].map { Copy.guildTabApplied(day: String(presenter.clock.weekday($0).prefix(3))) }
                return row
            }
            recruiting.raidsOpen = raidsOpen(character?.level)
            return .noGuild(recruiting)
        case .waiting:
            return .waiting
        case let .loaded(content):
            // A guild member under 60 signs up too (coordinator ruling 2026-11-10): one quiet line, never a block.
            return .guild(guild(content, extras: extras))
        }
    }

    func inputs(_ content: RaidsTabContent, extras: RaidRootExtras) -> HomeFlowInputs {
        var inputs = HomeFlowInputs(character: HomeCharacter(content.character))
        inputs.raid = HomeFlowModel.current(content.planned, now: presenter.now)
        inputs.badge = extras.badge
        inputs.applicants = extras.applicants
        inputs.ready = extras.ready
        inputs.answering = answering
        if content.offlineSince != nil { inputs.answering.offline = true }
        return inputs
    }

    private func raid(_ content: RaidsTabContent, extras: RaidRootExtras) -> RaidGuildContent {
        let inputs = inputs(content, extras: extras)
        let live = inputs.raid.map { $0.startsAt <= presenter.now } ?? false
        let night = live ? extras.night.flatMap { presenter.nightRoot(inputs, night: $0) } : nil
        let stamp = { (date: Date) in presenter.clock.stripTime(date, now: presenter.now) }
        let upcoming = night == nil ? inputs.raid : nil
        var next = night == nil ? presenter.nextRaid(inputs) : nil
        next?.invites = upcoming.flatMap(presenter.invites)
        let officer = extras.badge?.officer == true
        return RaidGuildContent(
            night: night,
            next: next,
            coming: night == nil ? inputs.raid.map(presenter.coming) ?? [] : [],
            ready: presenter.ready(extras.night?.ready ?? extras.ready),
            needsYou: presenter.needsYou(inputs, raid: inputs.raid),
            pages: presenter.pages(
                guild: content.head.name, members: members ?? content.head.members, atCap: atCap, planned: content.planned,
                badge: extras.badge, lastNight: content.nights.first
            ),
            offlineStamp: content.offlineSince.map(stamp),
            staleStamp: content.staleSince.map(stamp),
            notes: upcoming.flatMap(RaidFlowPresenter.notes),
            answers: upcoming.flatMap(presenter.answers),
            cover: upcoming.flatMap(presenter.cover),
            notSynced: officer ? upcoming.flatMap { presenter.notSynced($0, synced: synced) } : nil
        )
    }

    private func guild(_ content: RaidsTabContent, extras: RaidRootExtras) -> RaidGuildContent {
        var guild = raid(content, extras: extras)
        let level = content.character.level
        let raid = HomeFlowModel.current(content.planned, now: presenter.now)
        guild.raidsOpen = raidsOpen(level)
        guild.jobs = RaidAssignmentWords.jobs(raid)
        let officer = extras.badge?.officer == true || home?.viewer.officerTools == true
        guild.assignments = officer && raid?.assignments != nil ? RaidAssignmentWords.encounters(raid) : nil
        var inputs = GuildTabInputs(home: home, setup: extras.setup, applicants: extras.applicants, badge: extras.badge)
        inputs.level = level
        inputs.characterKey = content.character.key
        inputs.planned = content.planned
        inputs.firstRaids = firstRaids
        inputs.lootNights = extras.lootNights
        guild.tab = GuildTabPresenter(now: presenter.now, clock: presenter.clock).content(inputs, guildName: content.head.name)
        if let tab = guild.tab { guild.pages = Self.pages(guild.pages, tab: tab) }
        return guild
    }

    /// The pages the Guild tab's sections do not already open (spec 2.1, 2.2): Roster has its
    /// section, Loot history too at 60, Guild sims is a 60's page.
    static func pages(_ pages: GuildPagesContent, tab: GuildTabContent) -> GuildPagesContent {
        let leveler = tab.spread != nil
        let rows = pages.rows.filter { row in
            switch row.page {
            case .roster: false
            case .loot: leveler
            case .sims: !leveler
            case .calendar: true
            }
        }
        return GuildPagesContent(guild: Copy.guildTabPages, rows: rows)
    }

    /// `Raids open at 60 · you're 36` under 60; nil at 60.
    private func raidsOpen(_ level: Int?) -> String? {
        guard let level, level < RaidFlowPresenter.maxLevel else { return nil }
        return Copy.guildTabRaidsOpen(level: RaidFlowPresenter.maxLevel, you: level)
    }
}

/// Loads the Raid tab's extras for the state the tab model reached: the officer's badge and
/// applicants, the next raid's Ready check, the live raid's night; with no guild, the recruiting
/// guilds. A call that fails leaves its part out, never the tab.
@MainActor
@Observable
final class RaidRootLoader {
    private(set) var extras = RaidRootExtras()
    private let service: any RaidFlowService
    private let raids: any RaidService
    private let now: @Sendable () -> Date

    init(service: any RaidFlowService, raids: any RaidService, now: @escaping @Sendable () -> Date) {
        self.service = service
        self.raids = raids
        self.now = now
    }

    func load(for state: RaidsTabState) async {
        switch state {
        case let .noGuild(character?):
            guard let spec = character.specKey, let faction = character.faction else { extras = RaidRootExtras(); return }
            let recruiting = try? await service.recruiting(spec: spec, faction: faction, region: character.region, ruleset: character.ruleset)
            extras = RaidRootExtras(recruiting: recruiting, applied: await applied(recruiting?.rows ?? []))
        case let .loaded(content):
            extras = await guildExtras(content)
        default:
            extras = RaidRootExtras()
        }
    }

    private func guildExtras(_ content: RaidsTabContent) async -> RaidRootExtras {
        let guild = content.head.id
        let raid = HomeFlowModel.current(content.planned, now: now())
        let liveID = raid.flatMap { $0.startsAt <= now() ? $0.id : nil }
        async let badge = try? service.badge(guild: guild)
        async let ready = readyCheck(character: content.character.key, raid: raid?.id)
        async let night = liveNight(liveID)
        async let loot = try? service.lootNights(guild: guild)
        var extras = RaidRootExtras(badge: await badge, ready: await ready, night: await night)
        extras.lootNights = await loot
        if let badge = extras.badge, badge.officer {
            if badge.applicants > 0 { extras.applicants = await applicants(guild: guild) }
            extras.setup = try? await service.setup(guild: guild)
        }
        return extras
    }

    private func readyCheck(character: String, raid: Int64?) async -> ReadyCheck? {
        guard let raid else { return nil }
        return await service.ready(character: character, raid: raid)
    }

    private func liveNight(_ raid: Int64?) async -> RaidNightSummary? {
        guard let raid else { return nil }
        return try? await raids.night(raid: raid)
    }

    /// When the player applied to each recruiting guild that is still deciding (spec 2.4).
    private func applied(_ guilds: [RecruitingGuild]) async -> [Int64: Date] {
        await withTaskGroup(of: (Int64, Date?).self) { group in
            for guild in guilds {
                group.addTask { [service] in
                    let rows = (try? await service.myApplications(guild: guild.guildId))?.rows ?? []
                    return (guild.guildId, rows.filter { $0.status == LiveApplicantsSource.pending }.map(\.createdAt).max())
                }
            }
            var out: [Int64: Date] = [:]
            for await (guild, at) in group { if let at { out[guild] = at } }
            return out
        }
    }

    private func applicants(guild: Int64) async -> [ApplicantLine] {
        guard let page = try? await service.applications(guild: guild) else { return [] }
        return ApplicantLine.pending(page)
    }
}

extension ApplicantLine {
    /// The pending applications, newest first.
    static func pending(_ page: GuildOfficerApplications) -> [ApplicantLine] {
        page.rows.filter { $0.status == LiveApplicantsSource.pending }.sorted { $0.createdAt > $1.createdAt }.map { row in
            ApplicantLine(
                id: row.id,
                name: row.character.name,
                kind: RaidFlowPresenter.kind(spec: row.character.specName, className: row.character.className),
                level: row.character.level ?? 0,
                appliedAt: row.createdAt
            )
        }
    }
}
