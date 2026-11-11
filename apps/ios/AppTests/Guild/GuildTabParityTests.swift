@testable import ForeverSixty
import DesignSystem
import Foundation
import SwiftUI
import Testing

/// The Guild tab's four states (spec 2026-11-10-guild-tab.md section 2) against their mocks
/// (design/mocks/ws_app_guild.py): at 60, officer, below 60, no guild (at 60 and below).
@MainActor
@Suite("Guild tab parity")
struct GuildTabParityTests {
    static let device = CGSize(width: 393, height: 852)
    typealias T = GuildTabTestData

    private func root(_ content: RaidRootContent, _ screen: String, testName: String = #function) async throws {
        let view = RaidRootView(content: content)
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.raidIconBuild, HomeFlowTestData.build)
        try await ParityCheck.assert(view, screen: screen, size: Self.device, testName: testName)
    }

    @Test("ios-guild-tab") func atCap() async throws { try await root(T.atCap(), "guild-tab") }
    @Test("ios-guild-tab-officer") func officer() async throws { try await root(T.atCap(officer: true), "guild-tab-officer") }
    @Test("ios-guild-tab-leveler") func leveler() async throws { try await root(T.leveler, "guild-tab-leveler") }
    @Test("ios-guild-tab-none") func noGuild() async throws { try await root(T.noGuild(leveler: false), "guild-tab-none") }
    @Test("ios-guild-tab-none-leveler") func noGuildLeveler() async throws {
        try await root(T.noGuild(leveler: true), "guild-tab-none-leveler")
    }
}
