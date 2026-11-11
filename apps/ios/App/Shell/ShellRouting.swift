import Routing
import SwiftUI

/// Where links, pushes and intents land (flows spec F4, F6): each screen pushed under its tab.
extension RootView {
    /// Sim, report, rankings, player and guide links land under Home (lanes F5, F10, F11, F9),
    /// list links under Gear (lane F9), raid links under Guild (lane F4), the account's under Home.
    func openLink(_ link: DeepLink) {
        // F6: the site's own pages open the app's screen for them, under its tab.
        if let app = siteLink(link) {
            openLink(app)
            return
        }
        if case .tiersPage = link {
            guard environment.isAvailable(.guides) else { return }
            shell.push(.tiers, on: .home)
        } else if case let .guildRaidPage(region, ruleset, guildName, raid) = link {
            openGuildRaid(region: region, ruleset: ruleset, slug: guildName, raid: raid, url: link.url)
        } else if PlannerRouter.handles(link) {
            // Lane F13: with the switch off, a pasted /b/ link opens the site's page instead.
            if !plannerRouter.open(link), case .sharedBuild = link { openURL(link.url) }
        } else if SimsRouter.handles(link) {
            guard environment.isAvailable(.sims) else { return }
            shell.selection = .home
            simsRouter.open(link)
        } else if ReportsRouter.handles(link) {
            guard environment.isAvailable(.reports) else { return }
            shell.selection = .home
            reportsRouter.open(link)
        } else if RankingsRouter.handles(link) {
            guard environment.isAvailable(.rankings) else { return }
            shell.selection = .home
            rankingsRouter.open(link)
        } else if GuidesRouter.handles(link) {
            guard environment.isAvailable(.guides) else { return }
            guidesRouter.open(link)
            shell.push(.guides, on: .home)
        } else if BisRouter.handles(link) {
            guard environment.isAvailable(.bis) else { return }
            shell.selection = .gear
            bisRouter.open(link)
        } else if let account = Self.accountDestination(link) {
            shell.push(account, on: .home)
        } else {
            openRaids(link)
        }
    }

    /// A site page (F6) as the app link for the same screen; nil for an app link or a page the
    /// app routes by itself.
    func siteLink(_ link: DeepLink) -> DeepLink? {
        switch link {
        case let .simPage(id): .sim(id: id)
        case let .reportPage(id): .report(id: id)
        case let .characterPage(region, ruleset, name): .player(region: region, ruleset: ruleset, name: name)
        case let .bisPage(className, spec): .bis(spec: "\(className)-\(spec)", band: BisRouter.raidLinkBand)
        case let .guidePage(className, spec): .guide(className: className, spec: spec)
        case let .guidesClassPage(className):
            guidesRouter.openClass(className) ? .guides : nil
        default: nil
        }
    }

    /// A site raid page names its guild by region, ruleset and name; the account's own guild with
    /// that name opens the raid under Raid, any other the site in the browser.
    private func openGuildRaid(region: String, ruleset: String, slug: String, raid: Int64, url: URL) {
        let guild = environment.session.phase.me?.guilds.first {
            $0.region.lowercased() == region.lowercased() && $0.ruleset.lowercased() == ruleset.lowercased()
                && GuildSlug.of($0.name) == slug.lowercased()
        }
        guard let guild else {
            openURL(url)
            return
        }
        openRaids(.raid(guild: guild.id, raid: raid))
    }

    /// The account screens a link names (`/app/you/notifications`, `/app/you/plan`).
    static func accountDestination(_ link: DeepLink) -> MenuDestination? {
        switch link {
        case .notifications: .notifications
        case .plan: .plan
        default: nil
        }
    }

    /// Lane F4: a raid, raid-night or guild link selects the Guild tab and pushes its screen (spec
    /// 2026-11-10-guild-tab.md: a raid push opens the raid inside the Guild tab).
    func openRaids(_ link: DeepLink) {
        guard environment.isAvailable(.raids), RaidsRouter.handles(link) else { return }
        shell.selection = .guild
        raidsRouter.open(link)
    }

    /// Lane F15: an intent's, a Spotlight tap's or a notification's destination, opened as the
    /// in-app link would be.
    func openIntent(_ destination: IntentDestination) {
        switch destination {
        case let .link(link): openLink(link)
        case let .tab(tab): shell.selection = tab
        case let .raid(route):
            shell.selection = .guild
            if environment.isAvailable(.raids) { raidsRouter.open(route) }
        case let .character(key): openCharacter(key)
        case .rankingsBoard:
            guard environment.isAvailable(.rankings) else { return }
            shell.selection = .home
            rankingsRouter.player = nil
            rankingsRouter.encounter = nil
            rankingsRouter.showsBoard = true
        }
    }

    /// A character opens F11.4 when `rankings` is on; else Home on it (your own characters only).
    private func openCharacter(_ key: String) {
        let parts = key.split(separator: "/").map(String.init)
        if environment.isAvailable(.rankings), parts.count == 3 {
            openLink(.player(region: parts[0], ruleset: parts[1], name: parts[2]))
            return
        }
        if environment.session.phase.me?.characters.contains(where: { $0.key == key }) == true {
            environment.session.pick(key)
        }
        shell.selection = .home
    }
}

/// A guild's name as the site's paths write it (api/internal/character Slug): lower case, trimmed,
/// spaces as hyphens.
enum GuildSlug {
    static func of(_ name: String) -> String {
        name.trimmingCharacters(in: .whitespaces).lowercased().replacingOccurrences(of: " ", with: "-")
    }
}
