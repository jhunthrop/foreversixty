@testable import ForeverSixty
import Foundation
import Models
import Networking
import Routing
import Testing

/// A raid link or push opening the Raids stack (flows spec F1, F4), as Android's
/// RaidDeepLinkFlowTest: the URL parsed as RootView.onOpenURL does, RaidsRouter's stack, the raid
/// page (RaidPageModel) with the press applied or the spec asked, back to the tab root; and a raid
/// push's payload resolving to those links.
@MainActor
@Suite("Raid link flow")
struct RaidDeepLinkFlowTests {
    private func deps(_ service: EchoRaidService) -> SignupController.Dependencies {
        SignupController.Dependencies(
            service: service, haptics: RecordingHaptics(), permission: FakePermission(undetermined: false),
            askStore: MemoryAskStore(), now: { RaidTime.now }, hold: nil
        )
    }

    /// The link as RootView opens it: its URL parsed back, then the Raids router.
    private func open(_ link: DeepLink, on router: RaidsRouter) throws {
        let parsed = try #require(DeepLink(url: link.url), "not an app link: \(link.url)")
        #expect(parsed == link)
        #expect(RaidsRouter.handles(parsed))
        router.open(parsed)
    }

    /// The raid page the top `.raid` route builds, loaded.
    private func page(_ route: RaidsRoute?, service: EchoRaidService) async throws -> SignupController {
        guard case let .raid(guild, raid, _, open)? = route else {
            Issue.record("not a raid page: \(String(describing: route))")
            throw CancellationError()
        }
        let model = RaidPageModel(guild: guild, raid: raid, deps: deps(service), cache: NoRaidCache(), openAction: open)
        await model.load()
        guard case let .loaded(controller) = model.state else {
            Issue.record("the raid page did not load")
            throw CancellationError()
        }
        return controller
    }

    @Test("a state=late link opens the raid page with Late pressed and saved; back returns to the tab root")
    func lateLink() async throws {
        let service = try EchoRaidService()
        let router = RaidsRouter()
        try open(.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, state: "late"), on: router)
        #expect(router.path == [.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, roster: false, open: RaidOpenAction(state: .late))])

        let controller = try await page(router.path.last, service: service)
        #expect(service.requests.map(\.key) == [FlowPlayer.key])
        #expect(service.sent.map(\.state) == ["late"])
        #expect(controller.block(FlowPlayer.key).shown == .late)
        #expect(controller.block(FlowPlayer.key).status == .saved(Copy.commonSaved))

        router.path.removeLast()
        #expect(router.path.isEmpty)
    }

    @Test("a spec=ask link opens the raid page asking the spec; the pick saves In with it")
    func specAsk() async throws {
        let service = try EchoRaidService()
        service.raidOverride = try RaidFixture.raid("guild.raid.ask-spec")
        let router = RaidsRouter()
        try open(.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, spec: "ask"), on: router)
        #expect(router.path == [.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, roster: false, open: RaidOpenAction(state: .signedIn))])

        let controller = try await page(router.path.last, service: service)
        #expect(controller.specAsk == .ask(characterKey: FlowPlayer.askKey, state: .signedIn))
        #expect(controller.block(FlowPlayer.askKey).pending == .signedIn)
        #expect(service.requests.isEmpty)

        let option = try #require(controller.character(FlowPlayer.askKey)?.specOptions?.first)
        await controller.pickSpec(option)
        #expect(service.requests.map(\.key) == [FlowPlayer.askKey])
        #expect(service.sent.first?.state == "in")
        #expect(service.sent.first?.spec == option.spec)
        #expect(controller.block(FlowPlayer.askKey).shown == .signedIn)
        #expect(controller.specAsk == nil)
    }

    @Test("a raid night link pushes the raid then the night; back pops each, then the tab root")
    func nightLink() throws {
        let router = RaidsRouter()
        try open(.raidNight(guild: FlowPlayer.guild, raid: FlowPlayer.raid), on: router)
        #expect(router.path == [.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, roster: false, open: nil), .night(raid: FlowPlayer.raid)])
        router.path.removeLast()
        #expect(router.path == [.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, roster: false, open: nil)])
        router.path.removeLast()
        #expect(router.path.isEmpty)
        // A second link starts a fresh stack: back still lands on the root.
        try open(.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid), on: router)
        try open(.raid(guild: FlowPlayer.guild, raid: 32), on: router)
        #expect(router.path.count == 1)
    }

    @Test("a view=unanswered link opens Not answered, as the raid.short push does")
    func unansweredLink() throws {
        let router = RaidsRouter()
        try open(.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, view: "unanswered"), on: router)
        #expect(router.path == [.notAnswered(guild: FlowPlayer.guild, raid: FlowPlayer.raid)])
    }

    // MARK: Pushes

    private func push(link: String?, kind: String = "raid.posted") throws -> [AnyHashable: Any] {
        var userInfo = try #require(JSONSerialization.jsonObject(with: RaidFixture.data("push.raid-posted.apns")) as? [AnyHashable: Any])
        userInfo["kind"] = kind
        userInfo["link"] = link
        var data = try #require(userInfo["data"] as? [String: Any])
        data["raid_id"] = String(FlowPlayer.raid)
        userInfo["data"] = data
        return userInfo
    }

    @Test("a raid push opens its raid page by the payload's ids, whatever site page its link names")
    func pushOpensRaid() throws {
        let userInfo = try push(link: "https://foreversixty.gg/guild/us/pvp/olympus-xxvii/raids/7")
        #expect(PushLanding.destination(userInfo: userInfo) == .link(.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid)))
        let payload = try #require(RaidPushPayload(userInfo: userInfo))
        #expect(payload.guild == FlowPlayer.guild && payload.raid == FlowPlayer.raid)
        #expect(payload.deepLink(askingSpec: true) == .raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, spec: "ask"))
        let router = RaidsRouter()
        guard case let .link(link)? = PushLanding.destination(userInfo: userInfo) else { Issue.record("no link"); return }
        router.open(link)
        #expect(router.path == [.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, roster: false, open: nil)])
    }

    @Test("a raid push whose app link says state= or spec= carries the press to the raid page")
    func pushCarriesPress() throws {
        let late = DeepLink.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, state: "late")
        #expect(PushLanding.destination(userInfo: try push(link: late.url.absoluteString)) == .link(late))
        let ask = DeepLink.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid, spec: "ask")
        #expect(PushLanding.destination(userInfo: try push(link: ask.url.absoluteString)) == .link(ask))
        // The ids win over another raid's link.
        let other = DeepLink.raid(guild: FlowPlayer.guild, raid: 99, state: "out")
        #expect(PushLanding.destination(userInfo: try push(link: other.url.absoluteString))
            == .link(.raid(guild: FlowPlayer.guild, raid: FlowPlayer.raid)))
        let payload = try #require(RaidPushPayload(userInfo: try push(link: late.url.absoluteString)))
        #expect(payload.deepLink() == late)
    }
}
