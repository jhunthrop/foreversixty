#if DEBUG
import Foundation
import Networking
import Routing
import UserNotifications

/// Debug-only launch arguments that drive the raid screens in fixture mode for the review
/// captures (ForeverSixtyUITests/RaidCaptureTests, BackfillCaptureTests): open a raid link, answer
/// the raid view and the badge from a named fixture state, and post a fixture push as a local
/// notification. Never compiled into a release.
enum RaidDebugLaunch {
    private static var arguments: [String] { ProcessInfo.processInfo.arguments }

    private static func value(_ flag: String) -> String? {
        guard let index = arguments.firstIndex(of: flag), arguments.indices.contains(index + 1) else { return nil }
        return arguments[index + 1]
    }

    /// `-FSOpenLink https://foreversixty.gg/app/raid/2/31`
    static var link: DeepLink? { value("-FSOpenLink").flatMap(URL.init(string:)).flatMap(DeepLink.init(url:)) }

    /// `-FSFailOfficerPush` refuses every officer push action (Mention, Accept, Decline), so the
    /// failure line replaces the push in the review captures (round 6, row 3).
    static var failsOfficerPush: Bool { arguments.contains("-FSFailOfficerPush") }

    /// `-FSRaidFixture officer` answers the raid view from `guild.raid.officer` and the badge from
    /// `guild.raid-badge.officer`: the officer's Raid tab.
    static var raidVariant: String? { value("-FSRaidFixture") }

    /// The routes a raid variant names, where the fixtures have that state.
    private static let variantStems = ["guild.raid", "guild.raid-badge"]

    private static var folder: URL? { Bundle.main.url(forResource: FixtureTransport.bundledFolder, withExtension: nil) }

    /// The bundled fixtures with the named raid state on every route that has it.
    static func client(baseURL: URL) -> APIClient? {
        guard let variant = raidVariant, let folder else { return nil }
        let stems = variantStems.filter { FileManager.default.fileExists(atPath: folder.appending(path: "\($0).\(variant).json").path) }
        let variants = Dictionary(uniqueKeysWithValues: stems.map { ($0, FixtureVariant.named(variant)) })
        return APIClient(baseURL: baseURL, transport: FixtureTransport(directory: folder, variants: variants))
    }

    /// A raid service over the bundled fixtures with the named raid state.
    static func service(baseURL: URL) -> (any RaidService)? {
        client(baseURL: baseURL).map(LiveRaidService.init(client:))
    }

    /// The Raid tab's flow calls (the badge says officer) over the same fixtures.
    static func flow(baseURL: URL) -> (any RaidFlowService)? {
        client(baseURL: baseURL).map(LiveRaidFlowService.init(client:))
    }

    /// `-FSPostRaidPush` posts push.raid-posted.apns on the fixture app's raid (31); `-FSPostPush
    /// push.raid-short` (or `push.guild-applicant`) posts that fixture as the server sends it, its
    /// actions and data in the user info, so the officer actions run as from APNs.
    static func postPushIfAsked() async {
        guard arguments.contains("-FSPostRaidPush") || value("-FSPostPush") != nil else { return }
        // A walk's earlier pushes would stack with this one, and a long press opens the stack, not
        // the actions.
        UNUserNotificationCenter.current().removeAllDeliveredNotifications()
        if arguments.contains("-FSPostRaidPush") {
            await post("push.raid-posted", raid: fixtureRaid)
        }
        if let name = value("-FSPostPush") {
            await post(name, raid: nil)
        }
    }

    /// The raid the fixture app lists.
    private static let fixtureRaid = "31"

    private static func post(_ name: String, raid: String?) async {
        guard let folder,
              let data = try? Data(contentsOf: folder.appending(path: "\(name).apns.json")),
              let payload = try? JSONSerialization.jsonObject(with: data) as? [String: Any],
              let aps = payload["aps"] as? [String: Any],
              let alert = aps["alert"] as? [String: Any]
        else { return }
        let center = UNUserNotificationCenter.current()
        guard (try? await center.requestAuthorization(options: [.alert, .sound])) == true else { return }
        let content = UNMutableNotificationContent()
        content.title = alert["title"] as? String ?? ""
        content.body = alert["body"] as? String ?? ""
        content.categoryIdentifier = aps["category"] as? String ?? RaidPushCategory.signup
        content.threadIdentifier = aps["thread-id"] as? String ?? ""
        var userInfo = payload.filter { $0.key != "aps" }
        if let raid {
            var fields = payload["data"] as? [String: String] ?? [:]
            fields["raid_id"] = raid
            userInfo["data"] = fields
        }
        content.userInfo = userInfo
        let request = UNNotificationRequest(identifier: name, content: content, trigger: UNTimeIntervalNotificationTrigger(timeInterval: 2, repeats: false))
        try? await center.add(request)
    }
}
#endif
