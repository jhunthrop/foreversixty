import Foundation
import Models
import Networking

/// The guild extras' calls (spec F16 Data): the guild home for the roster, the raid nights for
/// history, and the calendar link. Live goes through `APIClient`; tests pass a fake.
protocol GuildExtrasService: Sendable {
    func home(guild: Int64) async throws(APIError) -> GuildHome
    func history(guild: Int64, cursor: String?) async throws(APIError) -> GuildRaids
    func calendarFeed() async throws(APIError) -> CalendarFeed
    func resetCalendarFeed() async throws(APIError) -> CalendarFeed
}

/// History reads 20 nights a page (spec F16.3).
enum HistoryPaging {
    static let pageSize = 20
}

struct LiveGuildExtrasService: GuildExtrasService {
    let client: APIClient

    func home(guild: Int64) async throws(APIError) -> GuildHome {
        try await client.send(.getGuildHome(id: guild))
    }

    func history(guild: Int64, cursor: String?) async throws(APIError) -> GuildRaids {
        try await client.send(.listGuildRaids(id: guild, limit: HistoryPaging.pageSize, cursor: cursor))
    }

    func calendarFeed() async throws(APIError) -> CalendarFeed { try await client.send(.getCalendarFeed) }

    func resetCalendarFeed() async throws(APIError) -> CalendarFeed { try await client.send(.resetCalendarFeed) }
}

/// Fixture mode: the roster fixture for the guild home, and the two history pages by cursor
/// (the fixture table keys on the path, so the cursor picks the page here).
struct FixtureGuildExtrasService: GuildExtrasService {
    let directory: URL
    let baseURL: URL

    private func client(_ variants: [String: FixtureVariant]) -> APIClient {
        APIClient(baseURL: baseURL, transport: FixtureTransport(directory: directory, variants: variants))
    }

    func home(guild: Int64) async throws(APIError) -> GuildHome {
        try await client(["guild.home": .named("roster")]).send(.getGuildHome(id: guild))
    }

    func history(guild: Int64, cursor: String?) async throws(APIError) -> GuildRaids {
        let page = cursor == nil ? "history" : "history-2"
        return try await client(["guild.raids": .named(page)])
            .send(.listGuildRaids(id: guild, limit: HistoryPaging.pageSize, cursor: cursor))
    }

    func calendarFeed() async throws(APIError) -> CalendarFeed { try await client([:]).send(.getCalendarFeed) }

    func resetCalendarFeed() async throws(APIError) -> CalendarFeed { try await client([:]).send(.resetCalendarFeed) }
}
