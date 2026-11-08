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
}

/// Composes the Raid tab's root (flows spec 4.2) from the tab's state and its extras.
struct RaidRootComposer {
    let presenter: RaidFlowPresenter
    /// The roster's figures for the Roster row, when the guild's roster is loaded.
    var members: Int?
    var atCap: Int?
    var answer: SignupState?
    var lateOpen = false

    func content(_ state: RaidsTabState, extras: RaidRootExtras) -> RaidRootContent {
        switch state {
        case .loading: return .loading
        case .failed: return .failed
        case .signedOut: return .signedOut
        case let .noGuild(character):
            if let level = character?.level, level < RaidFlowPresenter.maxLevel { return RaidFlowPresenter.leveler(level: level) }
            let kind = RaidFlowPresenter.kind(spec: character?.spec, className: character?.className)
            return .noGuild(RaidFlowPresenter.recruiting(extras.recruiting?.rows ?? [], kind: kind))
        case let .waiting(_, character):
            if let level = character.level, level < RaidFlowPresenter.maxLevel { return RaidFlowPresenter.leveler(level: level) }
            return .waiting
        case let .loaded(content):
            if let level = content.character.level, level < RaidFlowPresenter.maxLevel { return RaidFlowPresenter.leveler(level: level) }
            return .guild(guild(content, extras: extras))
        }
    }

    func inputs(_ content: RaidsTabContent, extras: RaidRootExtras) -> HomeFlowInputs {
        var inputs = HomeFlowInputs(character: HomeCharacter(content.character))
        inputs.raid = HomeFlowModel.current(content.planned, now: presenter.now)
        inputs.badge = extras.badge
        inputs.applicants = extras.applicants
        inputs.ready = extras.ready
        inputs.answer = answer
        inputs.lateOpen = lateOpen
        return inputs
    }

    private func guild(_ content: RaidsTabContent, extras: RaidRootExtras) -> RaidGuildContent {
        let inputs = inputs(content, extras: extras)
        let live = inputs.raid.map { $0.startsAt <= presenter.now } ?? false
        let night = live ? extras.night.flatMap { presenter.nightRoot(inputs, night: $0) } : nil
        let stamp = { (date: Date) in presenter.clock.stripTime(date, now: presenter.now) }
        return RaidGuildContent(
            night: night,
            next: night == nil ? presenter.nextRaid(inputs) : nil,
            ready: night == nil ? nil : presenter.ready(extras.night?.ready ?? extras.ready),
            needsYou: presenter.needsYou(inputs, raid: inputs.raid),
            pages: presenter.pages(
                guild: content.head.name, members: members ?? content.head.members, atCap: atCap, planned: content.planned,
                badge: extras.badge
            ),
            offlineStamp: content.offlineSince.map(stamp),
            staleStamp: content.staleSince.map(stamp)
        )
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
            extras = RaidRootExtras(recruiting: recruiting)
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
        var extras = RaidRootExtras(badge: await badge, ready: await ready, night: await night)
        if let badge = extras.badge, badge.officer, badge.applicants > 0 {
            extras.applicants = await applicants(guild: guild)
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
