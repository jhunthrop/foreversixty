#if DEBUG
import Foundation
import Networking
import Routing
import UserNotifications

/// Debug-only launch arguments that drive the raid screens in fixture mode for the review
/// captures (ForeverSixtyUITests/RaidCaptureTests): open a raid link, answer the raid view from a
/// named fixture state, and post the fixture raid notification. Never compiled into a release.
enum RaidDebugLaunch {
    private static var arguments: [String] { ProcessInfo.processInfo.arguments }

    private static func value(_ flag: String) -> String? {
        guard let index = arguments.firstIndex(of: flag), arguments.indices.contains(index + 1) else { return nil }
        return arguments[index + 1]
    }

    /// `-FSOpenLink https://foreversixty.gg/app/raid/2/31`
    static var link: DeepLink? { value("-FSOpenLink").flatMap(URL.init(string:)).flatMap(DeepLink.init(url:)) }

    /// `-FSRaidFixture officer` answers the raid view from `guild.raid.officer`.
    static var raidVariant: String? { value("-FSRaidFixture") }

    /// A raid service over the bundled fixtures with the named raid state.
    static func service(baseURL: URL) -> (any RaidService)? {
        guard let variant = raidVariant,
              let folder = Bundle.main.url(forResource: FixtureTransport.bundledFolder, withExtension: nil)
        else { return nil }
        let transport = FixtureTransport(directory: folder, variants: ["guild.raid": .named(variant)])
        return LiveRaidService(client: APIClient(baseURL: baseURL, transport: transport))
    }

    /// `-FSPostRaidPush` posts push.raid-posted.apns as a local notification with its category.
    static func postRaidPushIfAsked() async {
        guard arguments.contains("-FSPostRaidPush"),
              let folder = Bundle.main.url(forResource: FixtureTransport.bundledFolder, withExtension: nil),
              let data = try? Data(contentsOf: folder.appending(path: "push.raid-posted.apns.json")),
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
        var fields = payload["data"] as? [String: String] ?? [:]
        fields["raid_id"] = "31"
        content.userInfo = ["data": fields, "link": payload["link"] as? String ?? ""]
        let request = UNNotificationRequest(identifier: "raid.posted-31", content: content, trigger: UNTimeIntervalNotificationTrigger(timeInterval: 2, repeats: false))
        try? await center.add(request)
    }
}
#endif
