import Foundation
import Models
import Networking

/// The sim calls F5 makes, all on today's API. Live goes through `APIClient`; tests pass a fake.
protocol SimsService: Sendable {
    func list(page: Int, lane: String?) async throws(APIError) -> SimList
    func result(id: String) async throws(APIError) -> SimResult
    func progress(id: String) async throws(APIError) -> SimProgress
    func rerun(id: String) async throws(APIError) -> SimStarted
    func markSeen(id: String) async throws(APIError)
}

struct LiveSimsService: SimsService {
    let client: APIClient

    func list(page: Int, lane: String?) async throws(APIError) -> SimList {
        try await client.send(.listMySims(page: page, lane: lane))
    }

    func result(id: String) async throws(APIError) -> SimResult { try await client.send(.getSim(id: id)) }
    func progress(id: String) async throws(APIError) -> SimProgress { try await client.send(.getSimProgress(id: id)) }
    func rerun(id: String) async throws(APIError) -> SimStarted { try await client.send(.rerunSim(id: id)) }
    func markSeen(id: String) async throws(APIError) { _ = try await client.send(.markSimSeen(id: id)) }
}

/// The server lane's name on the list query and on a row.
enum SimLane {
    static let server = "server"
}
