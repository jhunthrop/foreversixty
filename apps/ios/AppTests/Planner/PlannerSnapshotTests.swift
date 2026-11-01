@testable import ForeverSixty
import DesignTokens
import Models
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// The planner's boards (spec F13, board f13-planner-ios) at iPhone 17 size, at the default text
/// size and at AX5, with the repo's own icons. Recorded on the shared iOS 26.5 simulator like the
/// Talents snapshots. The Reset alert is a system alert and is not drawn by a view snapshot.
@MainActor
@Suite("Planner snapshots")
struct PlannerSnapshotTests {
    enum TextSize: String, CaseIterable {
        case standard, ax5

        var category: UIContentSizeCategory {
            self == .standard ? .large : .accessibilityExtraExtraExtraLarge
        }
    }

    private static let iPhone17 = ViewImageConfig(
        safeArea: UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0),
        size: CGSize(width: 402, height: 874),
        traits: UITraitCollection(userInterfaceStyle: .dark)
    )

    private let file: TalentsFile
    private let index: PlannerIndex
    private let races: [RaceInfo]

    init() throws {
        file = try PlannerTestData.shaman()
        index = PlannerIndex(file)
        races = PlannerNames.legalRaces(classID: file.classId, races: try PlannerTestData.races(), combos: try PlannerTestData.combos())
    }

    private func snap(_ view: some View, size: TextSize = .standard, named name: String) {
        let traits = UITraitCollection { mutable in
            mutable.userInterfaceStyle = .dark
            mutable.preferredContentSizeCategory = size.category
        }
        assertSnapshot(
            of: view.environment(\.itemIconSource, RepoItemIconSource()).preferredColorScheme(.dark),
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: Self.iPhone17), traits: traits),
            named: size.rawValue,
            testName: name
        )
    }

    private func editor(_ state: PlannerState, tree: Int = 0) -> some View {
        let model = PlannerEditorModel(
            launch: PlannerLaunch(classSlug: "shaman", raceSlug: "orc", level: 60),
            dependencies: .test(source: FakePlannerSource()), state: state, races: races
        )
        return PlannerEditorView(model: model, tree: .constant(tree), build: file.build)
    }

    private func id(_ name: String) -> Int { PlannerTestData.talentID(file, name) }

    /// 49 of the Elemental guide build's 51 points, Elemental Fury in the strip: 2 left.
    private var twoLeft: PlannerState {
        var state = PlannerState(index: index, order: Array(PlannerTestData.elementalOrder(index).prefix(49)), raceSlug: "orc")
        state.focus(id("Elemental Fury"))
        return state
    }

    @Test("F13.1 with points left", arguments: TextSize.allCases)
    func plan(size: TextSize) {
        snap(editor(twoLeft), size: size, named: "planner-plan")
    }

    @Test("F13.1 at zero left")
    func zero() {
        var state = PlannerState(index: index, order: PlannerTestData.elementalOrder(index), raceSlug: "orc")
        state.focus(id("Elemental Fury"))
        snap(editor(state), named: "planner-zero")
    }

    @Test("F13.2 a tier refusal")
    func refusal() throws {
        var state = PlannerState(index: index, order: Array(PlannerTestData.elementalOrder(index).prefix(14)), raceSlug: "orc")
        let deep = try #require(index.trees[0].talents.first { $0.tier == 4 })
        #expect(state.add(deep.id) == .refused)
        snap(editor(state), named: "planner-refusal")
    }

    @Test("F13.3 the share confirm with a title")
    func share() {
        snap(
            PlannerShareSheet(title: .constant("Elemental raid"), signedIn: true, state: .idle, proceed: {}, cancel: {}),
            named: "planner-share"
        )
    }

    private func summary(_ id: String, order: [Int], title: String?, views: Int64, race: (String, String) = ("Orc", "orc")) -> PlannerBuildSummary {
        PlannerBuildSummary(
            build: PlannerTestData.build(id: id, order: order, title: title, views: views), classSlug: "shaman", className: "Shaman",
            raceName: race.0, raceSlug: race.1, totals: PlannerRules.pointsPerTree(index, order: order)
        )
    }

    private func viewer(_ phase: SharedBuildPhase) -> some View {
        NavigationStack {
            SharedBuildView(phase: phase, tree: .constant(0), timeZone: .gmt)
                .navigationBarTitleDisplayMode(.inline)
                .toolbarBackground(FSColor.bg, for: .navigationBar)
                .toolbarBackground(.visible, for: .navigationBar)
        }
    }

    @Test("F13.5 a shared build (Compare with mine is not offered on iOS)")
    func shared() {
        let content = SharedBuildContent(summary: summary("k7q2m9xa", order: PlannerTestData.elementalOrder(index), title: "Elemental raid", views: 12), file: file)
        snap(viewer(.loaded(content)), named: "shared-build")
    }

    @Test("F13.5 an unknown id and offline with nothing saved")
    func sharedStates() {
        snap(viewer(.notFound), named: "shared-notfound")
        snap(viewer(.offline), named: "shared-offline")
    }

    private func builds(_ phase: YourBuildsPhase, removed: RemovedBuild? = nil) -> some View {
        NavigationStack {
            YourBuildsView(
                phase: phase, removed: removed,
                newBuild: PlannerNewBuild(classSlug: "shaman", className: "Shaman", raceSlug: "orc", level: 60), timeZone: .gmt
            )
            .navigationBarTitleDisplayMode(.inline)
            .toolbarBackground(FSColor.bg, for: .navigationBar)
            .toolbarBackground(.visible, for: .navigationBar)
        }
    }

    private var rows: [PlannerBuildSummary] {
        let elemental = PlannerTestData.elementalOrder(index)
        return [
            summary("a", order: elemental, title: "Elemental raid", views: 12),
            summary("b", order: Array(elemental.prefix(40)), title: nil, views: 3),
            summary("c", order: Array(elemental.prefix(20)), title: nil, views: 0, race: ("Troll", "troll")),
        ]
    }

    @Test("F13.6 Your builds", arguments: TextSize.allCases)
    func yourBuilds(size: TextSize) {
        snap(builds(.loaded(rows)), size: size, named: "your-builds")
    }

    @Test("F13.6 the Removed snackbar")
    func removed() {
        let rows = rows
        snap(builds(.loaded(Array(rows.dropFirst())), removed: RemovedBuild(summary: rows[0], position: 0)), named: "your-builds-removed")
    }

    @Test("F13.6 empty")
    func empty() {
        snap(builds(.loaded([])), named: "your-builds-empty")
    }
}
