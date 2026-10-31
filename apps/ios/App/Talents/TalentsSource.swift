import Foundation
import Models
import Networking

/// What the Talents screen reads: the active build, the class's talent file and the spec's BiS
/// file (its bands carry the guide builds), each with its offline copy. A protocol so the model
/// is tested without a network.
protocol TalentsSource: Sendable {
    func activeBuild() async -> String
    func talents(build: String, classSlug: String) async throws(APIError) -> CachedFile<TalentsFile>
    func bis(build: String, spec: String) async throws(APIError) -> CachedFile<BisFile>
}

/// The live source: the CDN through the offline cache, the build as the BiS tab reads it.
struct LiveTalentsSource: TalentsSource {
    let files: CachedStaticDataClient
    let bisSource: LiveBisDataSource

    func activeBuild() async -> String { await bisSource.activeBuild() }

    func talents(build: String, classSlug: String) async throws(APIError) -> CachedFile<TalentsFile> {
        try await files.load(.getTalents(build: build, className: classSlug))
    }

    func bis(build: String, spec: String) async throws(APIError) -> CachedFile<BisFile> {
        try await bisSource.bis(build: build, spec: spec)
    }
}
