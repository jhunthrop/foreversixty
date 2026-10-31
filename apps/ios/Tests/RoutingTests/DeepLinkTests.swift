import Foundation
import Routing
import Testing

/// apps/shared/links.json, read in place: every route's example opens its screen and round-trips.
private struct LinkTable: Decodable {
    struct Route: Decodable {
        let path: String
        let screen: String
        let example: String
    }

    let host: String
    let routes: [Route]

    static func load() throws -> LinkTable {
        let file = URL(fileURLWithPath: #filePath)
            .deletingLastPathComponent() // RoutingTests
            .deletingLastPathComponent() // Tests
            .deletingLastPathComponent() // apps/ios
            .deletingLastPathComponent() // apps
            .appending(path: "shared/links.json")
        return try JSONDecoder().decode(LinkTable.self, from: Data(contentsOf: file))
    }
}

@Suite("Deep links")
struct DeepLinkTests {
    @Test("every route's example opens its screen and builds the same link back")
    func everyRoute() throws {
        let table = try LinkTable.load()
        #expect(table.host == DeepLink.host)
        for route in table.routes {
            let url = try #require(URL(string: "https://\(table.host)\(route.example)"))
            let link = try #require(DeepLink(url: url), "\(route.path) did not parse \(route.example)")
            #expect(link.screen == route.screen, "\(route.path)")
            #expect(DeepLink(url: link.url) == link, "\(route.path) round trip")
        }
    }

    @Test("each route reads its parameters")
    func parameters() throws {
        let raid = DeepLink(url: URL(string: "https://foreversixty.gg/app/raid/2/41?spec=ask&state=in")!)
        #expect(raid == .raid(guild: 2, raid: 41, spec: "ask", state: "in"))
        let character = DeepLink(url: URL(string: "https://foreversixty.gg/app/character/eu/normal/j%C3%A4ger")!)
        #expect(character == .character(region: "eu", ruleset: "normal", name: "jäger"))
        #expect(DeepLink(url: URL(string: "https://foreversixty.gg/app/raid/2/41/night")!) == .raidNight(guild: 2, raid: 41))
        #expect(DeepLink(url: URL(string: "https://foreversixty.gg/app/guild/7")!) == .guild(guild: 7))
        #expect(DeepLink(url: URL(string: "https://foreversixty.gg/app/bis/warrior-fury/40")!) == .bis(spec: "warrior-fury", band: "40"))
        #expect(DeepLink(url: URL(string: "https://foreversixty.gg/app/sim/abc")!) == .sim(id: "abc"))
        #expect(DeepLink(url: URL(string: "https://foreversixty.gg/app/you/notifications")!) == .notifications)
        #expect(DeepLink(url: URL(string: "https://foreversixty.gg/app/you/plan")!) == .plan)
        #expect(DeepLink(url: URL(string: "https://foreversixty.gg/app/auth?code=c&state=s")!) == .signIn(code: "c", state: "s"))
    }

    @Test("anything else is not an app link", arguments: [
        "https://foreversixty.gg/raid/2/41",
        "https://evil.example/app/raid/2/41",
        "http://foreversixty.gg/app/raid/2/41",
        "https://foreversixty.gg/app/raid/two/41",
        "https://foreversixty.gg/app/raid/2",
        "https://foreversixty.gg/app/you",
    ])
    func rejects(link: String) throws {
        #expect(DeepLink(url: try #require(URL(string: link))) == nil)
    }

    @Test("an unknown query word is dropped, never trusted")
    func unknownWord() {
        let link = DeepLink(url: URL(string: "https://foreversixty.gg/app/raid/2/41?state=maybe")!)
        #expect(link == .raid(guild: 2, raid: 41))
    }
}
