import Foundation

/// `/data/active-build.json`: the client build whose static data the site serves now.
public struct ActiveBuild: Decodable, Sendable, Equatable {
    public let build: String
}

/// One entry of `/data/{build}/classes.json`.
public struct ClassInfo: Decodable, Sendable, Equatable, Identifiable {
    public let id: Int
    public let name: String
    public let slug: String
    /// The class colour, `#RRGGBB`.
    public let color: String
}
