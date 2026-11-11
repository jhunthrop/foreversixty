@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// The Raid tab's round-2 facts (design/reviews/2026-11-10-mobile-backfill/round-2-decisions.md):
/// one tally model (S6), names in class colour with a Late's minutes (G5), the Damage split, the
/// head's invite and size, the other answers' counts, the cover (G4), the officers' readiness (G1).
@MainActor
@Suite("Raid tab facts")
struct RaidTabFactsTests {
    private let presenter = RaidFlowPresenter(now: RaidTime.now, clock: RaidTime.clock)

    private func person(_ name: String, _ className: String = "priest", late: Int? = nil) -> (String, String?, String, Int?) {
        (name, className, late == nil ? "in" : "late", late)
    }

    @Test("S6: `Healers 5 of 8 · 3 short` in wipe tone, `Tanks 2 of 2` met, `Healers 6` without a target")
    func tally() {
        let five = (1 ... 5).map { person("H\($0)") }
        let short = RaidFlowPresenter.roleLine("Healers", people: five, want: 8)
        #expect(short.tally == Copy.raidTabTallyShort(tally: Copy.raidTabOfWant(have: 5, want: 8), n: 3))
        #expect(short.tally.hasSuffix("3 short"))
        #expect(short.tone == .short)
        #expect(short.compact == "5/8")
        let met = RaidFlowPresenter.roleLine("Tanks", people: [person("A"), person("B")], want: 2)
        #expect(met.tally == Copy.raidTabOfWant(have: 2, want: 2))
        #expect(met.tone == .met)
        let open = RaidFlowPresenter.roleLine("Healers", people: (1 ... 6).map { person("H\($0)") }, want: nil)
        #expect(open.tally == "6")
        #expect(open.tone == .plain)
        #expect(open.compact == nil)
        #expect(!short.names.contains("short"), "the gap is the tally's, not the names'")
    }

    @Test("names keep their class and a Late's minutes; past five, three and `and N more`")
    func names() {
        let line = RaidFlowPresenter.roleLine("Healers", people: [person("Mirethil"), person("Brightwen", "paladin", late: 15)], want: 8)
        #expect(line.people.map(\.className) == ["priest", "paladin"])
        #expect(line.people.last?.text == Copy.raidTabLateNote(name: "Brightwen", n: 15))
        #expect(line.more == 0)
        let long = RaidFlowPresenter.roleLine("Damage", people: (1 ... 11).map { person("D\($0)", "rogue") }, want: 30)
        #expect(long.people.map(\.name) == ["D1", "D2", "D3"])
        #expect(long.more == 8)
    }

    @Test("Who is coming: the Damage line splits melee and ranged; a Late reads its minutes before the raid (G5)")
    func coming() throws {
        let raid = try EchoRaidService.echo(
            try RaidFixture.raid(), key: FlowPlayer.key, SignupRequest(state: "late", role: nil, note: "", spec: nil, lateMinutes: 30)
        )
        let lines = presenter.coming(raid)
        let damage = try #require(lines.first { $0.role == Copy.homeRaidDamage })
        let dps = raid.signups.filter { SignupState(rawValue: $0.state)?.isComing == true && RaidFlowPresenter.roleGroup($0.role) == "dps" }
        let melee = dps.filter { $0.group == RaidFlowPresenter.meleeGroup }.count
        #expect(damage.split == Copy.raidTabDamageSplit(melee: melee, ranged: dps.count - melee))
        #expect(damage.people.first?.text == Copy.raidTabLateNote(name: "Gorvash", n: 30), "yours first, with its minutes")
        #expect(damage.people.first?.className == "warrior")
    }

    @Test("row 1: one count model, melee + ranged = Damage, and Home's card counts what the Raid tab does")
    func countModel() throws {
        for name in ["guild.raid", "guild.raid.officer", "guild.raid.two-characters"] {
            let raid = try RaidFixture.load(RaidView.self, name)
            let counts = RaidRoleCounts.of(raid)
            #expect(counts.melee + counts.ranged == counts.damage, "\(name)")
            let coming = raid.signups.filter { SignupState(rawValue: $0.state)?.isComing == true }
            #expect(counts.tanks + counts.healers + counts.damage == coming.count, "\(name): every In and Late counted once")
            let raidTab = presenter.coming(raid)
            let home = HomeFlowPresenter(now: RaidTime.now, clock: RaidTime.clock).tally(raid)
            #expect(home.map(\.role) == raidTab.map(\.role), "\(name)")
            #expect(home.map(\.have) == raidTab.map(\.have), "\(name): Home equals Raid")
            let damage = try #require(raidTab.first { $0.role == Copy.homeRaidDamage })
            #expect(damage.have == counts.damage)
            #expect(damage.split == Copy.raidTabDamageSplit(melee: counts.melee, ranged: counts.ranged))
        }
        let raid = try RaidFixture.raid()
        #expect(RaidRoleCounts.of(raid) == RaidRoleCounts(tanks: 3, healers: 6, melee: 7, ranged: 8), "the Late melee is coming")
    }

    @Test("row 2: Home's compact tally carries the short count; without a target the plain count")
    func compactTally() {
        #expect(RoleLineContent(role: "Healers", have: 5, want: 8, names: "").compactTally == "5/8 · 3 short")
        #expect(RoleLineContent(role: "Tanks", have: 2, want: 2, names: "").compactTally == "2/2")
        #expect(RoleLineContent(role: "Healers", have: 6, want: nil, names: "").compactTally == "6")
        #expect(RoleLineContent(role: "Healers", have: 6, want: nil, names: "").meter == nil)
    }

    @Test("row 6: Needs you and the foot count the raid's own not-answered list")
    func unansweredAgree() throws {
        let raid = try RaidFixture.load(RaidView.self, "guild.raid.officer")
        let badge = RaidBadge(count: 6, youUnanswered: false, officer: true, unanswered: 5, applicants: 1)
        let me = try RaidFixture.load(Me.self, "me")
        var inputs = HomeFlowInputs(character: HomeCharacter(try #require(me.characters.first)))
        inputs.raid = raid
        inputs.badge = badge
        let officer = try #require(HomeFlowPresenter(now: RaidTime.now, clock: RaidTime.clock).officer(inputs, raid: raid))
        let title = try #require(officer.unanswered?.title)
        #expect(title.hasPrefix("\(raid.notAnsweredCount) "), "not the badge's \(badge.unanswered)")
        #expect(try #require(presenter.answers(raid)).hasSuffix(Copy.raidTabCountNotAnswered(n: raid.notAnsweredCount)))
    }

    @Test("the head's facts: `Invites 6:45 pm · 40`; the foot: bench, out and not answered")
    func facts() throws {
        let raid = try RaidFixture.raid()
        let invites = try #require(presenter.invites(raid))
        #expect(invites.hasPrefix("Invites "))
        #expect(invites.hasSuffix(" · 40-man"))
        #expect(invites.contains("6:45"))
        let answers = try #require(presenter.answers(raid))
        let unanswered = raid.notSignedRaiders?.count ?? raid.notSigned.count
        var parts = [raid.counts.bench.map(Copy.raidTabCountBench), raid.counts.signedOut.map(Copy.raidTabCountOut)].compactMap(\.self)
        if unanswered > 0 { parts.append(Copy.raidTabCountNotAnswered(n: unanswered)) }
        #expect(answers == parts.joined(separator: " · "))
        #expect(RaidFlowPresenter.notes(raid) == raid.notes.trimmingCharacters(in: .whitespacesAndNewlines))
    }

    @Test("G4: the cover names what is missing or short, and says nothing when all is covered")
    func cover() throws {
        let raid = try RaidFixture.raid()
        let cover = try #require(presenter.cover(raid))
        #expect(cover.line == "Missing Kings · Might · Wisdom and 2 more", "the short labels, one separator")
        #expect(cover.missing.map(\.icon) == ["spell_magic_magearmor", "spell_holy_fistofjustice", "spell_holy_sealofwisdom"])
        let covered = try edited(raid) { json in
            json["cover"] = (json["cover"] as? [[String: Any]] ?? []).map { var row = $0; row["state"] = "ok"; return row }
        }
        #expect(presenter.cover(covered) == nil)
    }

    @Test("G4: one line, missing first, three names, one Missing, then and N more")
    func coverLine() throws {
        let raid = try RaidFixture.raid()
        func with(_ states: [(String, String, Int)]) throws -> RaidView {
            try edited(raid) { json in
                let template = (json["cover"] as? [[String: Any]] ?? [])[0]
                json["cover"] = states.enumerated().map { index, entry in
                    var row = template
                    row["key"] = "gap\(index)"
                    row["name"] = entry.0
                    row["state"] = entry.1
                    row["providers_in"] = entry.2
                    row["providers_late"] = 0
                    row["need"] = 2
                    return row
                }
            }
        }
        let windfury = ("Windfury Totem", "short", 1)
        let fortitude = ("Power Word: Fortitude", "missing", 0)
        let gaps = (1 ... 12).map { ("Gap \($0)", "missing", 0) }
        let one = try with([windfury, fortitude])
        let many = try with([windfury] + gaps)
        let two = try with([windfury] + gaps.prefix(2))
        #expect(presenter.cover(many)?.line == "Missing Gap 1 · Gap 2 · Gap 3 and 10 more")
        // These rows' keys (`gap0`...) have no composition label, so they read by name.
        #expect(presenter.cover(two)?.line == "Missing Gap 1 · Gap 2 · Windfury Totem 1 of 2", "missing first")
        #expect(presenter.cover(one)?.line == "Missing Power Word: Fortitude · Windfury Totem 1 of 2")
    }

    @Test("row 4: a gap takes the composition's short label for its key, else its name")
    func coverLabels() throws {
        let raid = try edited(try RaidFixture.raid()) { json in
            json["cover"] = (json["cover"] as? [[String: Any]] ?? []).map { row in
                var row = row
                if row["key"] as? String == "fortitude" { row["state"] = "missing" }
                return row
            }
        }
        let cover = try #require(presenter.cover(raid))
        #expect(cover.missing.first?.name == "Fort")
        #expect(cover.missing.dropFirst().first?.name == "Kings")
    }

    @Test("G4: without cover rows the line falls back to the buffs no one brings")
    func coverFallback() throws {
        let raid = try edited(try RaidFixture.raid()) { json in json["cover"] = NSNull() }
        let missing = (raid.composition?.buffs ?? []).filter { !$0.present }.map(\.label)
        #expect(presenter.cover(raid)?.line == (missing.isEmpty ? nil : Copy.raidTabCoverMissing(names: missing.joined(separator: " · "))))
    }

    @Test("round 4, s5: Not synced this week reads the addon's logged_at against a week, not logged_recently")
    func syncedThisWeek() throws {
        let roster = try GuildFixture.home().roster
        let synced = presenter.syncedThisWeek(roster)
        let stale = Set(["draknar", "skarn", "zekra", "lightningbolltz"].map { "us/pvp/\($0)" })
        #expect(synced.filter { !$0.value }.keys.sorted() == stale.sorted(), "logged_at eight days before the fixture clock")
        #expect(roster.filter { stale.contains($0.characterKey) }.allSatisfy { !$0.loggedRecently })
        let raid = try RaidFixture.raid("guild.raid.officer")
        let line = try #require(presenter.notSynced(raid, synced: synced))
        #expect(Set(line.keys) == stale)
    }

    @Test("G1: officers see how many signed In have not synced; only for them, only when some have not")
    func notSynced() async throws {
        let raid = try RaidFixture.raid()
        #expect(presenter.notSynced(raid, synced: nil) == nil)
        #expect(presenter.notSynced(raid, synced: [FlowPlayer.key: true]) == nil)
        let gorvash = try #require(raid.signups.first { $0.characterKey == FlowPlayer.key })
        let one = NotSyncedContent(line: Copy.raidTabNotSyncedNames(names: gorvash.name), keys: [FlowPlayer.key])
        #expect(presenter.notSynced(raid, synced: [FlowPlayer.key: false]) == one)
        let coming = raid.signups.filter { SignupState(rawValue: $0.state)?.isComing == true }
        let five = Dictionary(coming.prefix(5).map { ($0.characterKey, false) }) { first, _ in first }
        let named = try #require(presenter.notSynced(raid, synced: five))
        #expect(named.keys.count == 5)
        #expect(named.line.hasSuffix(" and 2 more"), "three names, then and N more")
        let four = Dictionary(coming.prefix(4).map { ($0.characterKey, false) }) { first, _ in first }
        #expect(try #require(presenter.notSynced(raid, synced: four)).line.components(separatedBy: ", ").count == 4)

        let tab = RaidsTabModel(service: try EchoRaidService(), cache: NoRaidCache(), now: { RaidTime.now })
        await tab.load()
        guard case let .loaded(loaded) = tab.state else { Issue.record("not loaded"); return }
        // Not synced is a 60's line (spec 2026-11-10-guild-tab.md 2.1): the fixture's officer is 58, so the root
        // is read with them at 60, and at 58 the line stays off the tab for officers too.
        let sixty = RaidsTabState.loaded(loaded.at(level: RaidFlowPresenter.maxLevel))
        let badge = { (officer: Bool) in RaidBadge(count: 0, youUnanswered: false, officer: officer, unanswered: 0, applicants: 0) }
        let synced = [FlowPlayer.key: false]
        func guild(_ officer: Bool, _ state: RaidsTabState = sixty) -> RaidGuildContent? {
            let root = RaidRootComposer(presenter: presenter, synced: synced).content(state, extras: RaidRootExtras(badge: badge(officer)))
            if case let .guild(guild) = root { return guild }
            return nil
        }
        #expect(guild(true)?.notSynced?.keys == [FlowPlayer.key])
        #expect(guild(false)?.notSynced == nil)
        #expect(guild(true, tab.state)?.notSynced == nil, "under 60 the line stays on the raid page")
        #expect(guild(false)?.notes == "Bring **flasks**")
        #expect(guild(false)?.next?.invites != nil)
    }

    @Test("a sign-up's replaced raid keeps the stale strip's time")
    func replacingKeepsStale() throws {
        let raid = try RaidFixture.raid()
        let me = try RaidFixture.load(Me.self, "me")
        let character = try #require(me.characters.first)
        var content = RaidsTabContent(
            head: GuildHead(character: try #require(character.guild), guild: nil), character: character,
            planned: [raid], nights: [], waiting: nil, offlineSince: nil
        )
        content.staleSince = RaidTime.now
        #expect(content.replacing(raid).staleSince == RaidTime.now)
    }

    private func edited(_ raid: RaidView, _ change: (inout [String: Any]) -> Void) throws -> RaidView {
        var json = try #require(JSONSerialization.jsonObject(with: WireFormat.encoder().encode(raid)) as? [String: Any])
        change(&json)
        return try WireFormat.decoder().decode(RaidView.self, from: JSONSerialization.data(withJSONObject: json))
    }
}
