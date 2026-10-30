import Foundation
import Models
import Networking
import TestSupport
import Testing

/// The models generated from api/openapi.yaml decode every success fixture of their routes.
/// The Go contract test holds the fixtures to the spec; this holds the generated Swift to both.
@Suite("Generated models decode the contract fixtures")
struct GeneratedModelsTests {
    private func data<T: Decodable & Sendable>(_: T.Type, _ file: String) throws -> T {
        let envelope = try WireFormat.decoder().decode(Envelope<T>.self, from: FixtureLocation.data("\(file).json"))
        return try #require(envelope.data)
    }

    private func bare<T: Decodable>(_: T.Type, _ file: String) throws -> T {
        try WireFormat.decoder().decode(T.self, from: FixtureLocation.data("\(file).json"))
    }

    @Test(arguments: ["me", "me.premium", "me.empty", "me.patch"])
    func me(file: String) throws { _ = try data(Me.self, file) }

    @Test(arguments: ["character.card", "character.card.empty"])
    func characterCard(file: String) throws { _ = try data(CharacterProfile.self, file) }

    @Test(arguments: ["character.home", "character.home.empty", "character.home.no-gear"])
    func characterHome(file: String) throws {
        let home = try data(CharacterHome.self, file)
        #expect(home.upgradeCount == home.upgrades.count)
    }

    @Test(arguments: ["character.search", "character.search.empty"])
    func search(file: String) throws { _ = try data(CharacterSearch.self, file) }

    @Test(arguments: ["character.sim-input.addon", "character.sim-input.fight"])
    func simInput(file: String) throws { _ = try data(SimInput.self, file) }

    @Test func sims() throws {
        _ = try data(SimList.self, "sims.list")
        _ = try data(SimList.self, "sims.list.empty")
        _ = try data(SimResult.self, "sims.result")
        _ = try data(SimProgress.self, "sims.progress")
        _ = try data(SimProgress.self, "sims.progress.queued")
        _ = try data(SimStarted.self, "sims.rerun")
    }

    @Test func guilds() throws {
        _ = try data(GuildHome.self, "guild.home")
        _ = try data(GuildRaids.self, "guild.raids")
        _ = try data(GuildRaids.self, "guild.raids.public")
        _ = try data(RaidView.self, "guild.raid")
        _ = try data(RaidView.self, "guild.raid.done")
        _ = try data(RaidView.self, "guild.raid.patch")
        let moved = try data(SignupAnswer.self, "guild.signup.moved")
        #expect(moved.moved != nil)
        _ = try data(SignupAnswer.self, "guild.signup")
        _ = try data(GuildLoot.self, "guild.loot")
    }

    @Test func reportsAndBuilds() throws {
        _ = try data(ReportList.self, "reports.list")
        _ = try data(ReportList.self, "reports.list.empty")
        _ = try data(Report.self, "reports.get")
        _ = try data(BuildList.self, "builds.list")
        _ = try data(Build.self, "builds.get")
        _ = try data(BuildSaved.self, "builds.save")
    }

    @Test func app() throws {
        _ = try data(AppSession.self, "app.signin.exchange")
        _ = try data(SignedOut.self, "app.logout")
        _ = try data(NotificationSettings.self, "app.notifications")
        _ = try data(PushTokenRegistered.self, "app.push-tokens.register")
        _ = try data(AppConfig.self, "app.config")
    }

    @Test func staticFiles() throws {
        _ = try bare(ActiveBuild.self, "cdn.active-build")
        _ = try bare([KeyDate].self, "cdn.dates")
        _ = try bare([ClassInfo].self, "cdn.classes")
        _ = try bare(BisFile.self, "cdn.bis.warrior-fury")
        _ = try bare(ItemsFile.self, "cdn.items.warrior")
        _ = try bare(TalentsFile.self, "cdn.talents.warrior")
    }

    @Test("a request body encodes in the API's snake_case")
    func encodesBody() throws {
        let body = AppExchangeRequest(code: "c", codeVerifier: "v", deviceName: "Phone")
        let json = try #require(String(data: WireFormat.encoder().encode(body), encoding: .utf8))
        #expect(json == #"{"code":"c","code_verifier":"v","device_name":"Phone"}"#)
    }
}
