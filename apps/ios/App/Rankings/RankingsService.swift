import Foundation
import Models
import Networking

/// The calls F11 makes. Live goes through `APIClient`; tests pass a fake.
protocol RankingsService: Sendable {
    func rankings(_ query: RankingsQuery) async throws(APIError) -> RankingsPage
    func encounters() async throws(APIError) -> RankedEncounters
    func phases() async throws(APIError) -> Phases
    func character(_ ref: CharacterRef) async throws(APIError) -> CharacterProfile
    func watching() async throws(APIError) -> WatchingList
    func watch(_ ref: CharacterRef) async throws(APIError)
    func unwatch(_ ref: CharacterRef) async throws(APIError)
}

struct LiveRankingsService: RankingsService {
    let client: APIClient

    func rankings(_ query: RankingsQuery) async throws(APIError) -> RankingsPage {
        try await client.send(.getRankings(
            encounter: query.encounter, metric: query.metric, spec: query.spec, class: query.classSlug, phase: query.phase,
            region: query.region, ruleset: query.ruleset, faction: query.faction, since: query.since,
            distinct: query.distinct, character: query.character, page: query.page
        ))
    }

    func encounters() async throws(APIError) -> RankedEncounters { try await client.send(.getRankingEncounters) }

    func phases() async throws(APIError) -> Phases { try await client.send(.getPhases) }

    func character(_ ref: CharacterRef) async throws(APIError) -> CharacterProfile {
        try await client.send(.getCharacter(region: ref.region, ruleset: ref.ruleset, name: ref.slug))
    }

    func watching() async throws(APIError) -> WatchingList { try await client.send(.getWatching) }

    func watch(_ ref: CharacterRef) async throws(APIError) {
        _ = try await client.send(.watchCharacter(region: ref.region, ruleset: ref.ruleset, name: ref.slug))
    }

    func unwatch(_ ref: CharacterRef) async throws(APIError) {
        _ = try await client.send(.unwatchCharacter(region: ref.region, ruleset: ref.ruleset, name: ref.slug))
    }
}
