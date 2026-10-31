@testable import ForeverSixty
import DesignSystem
import Models
import Networking
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// Every F5 screen and state on iPhone 17 (402 x 874), at the default size and at AX5 (spec 15).
/// Recorded on the shared iOS 26.5 simulator; the PNGs double as the review captures in
/// design/reviews/2026-10-31-mobile-sims.
@MainActor
@Suite("Sims snapshots")
struct SimsSnapshotTests {
    static let phone = CGSize(width: 402, height: 874)
    static let safeArea = UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0)
    private let specs = SpecCatalog.bundled()
    private let now = SimsTestData.now

    private func snap(_ view: some View, _ name: String, height: CGFloat = 874, large: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let config = ViewImageConfig(safeArea: Self.safeArea, size: CGSize(width: Self.phone.width, height: height), traits: traits)
        let sized = view
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    private func list(
        _ phase: SimsPhase, entitled: Bool = true, progress: [String: SimProgress] = [:],
        rerun: [String: RerunStatus] = [:], pageFailed: Bool = false, plans: Bool = false
    ) -> some View {
        pushed(back: Copy.tabHome) {
            ScrollView {
                SimsListView(
                    phase: phase, entitled: entitled, progress: progress, rerun: rerun, pageFailed: pageFailed,
                    mayLinkToPlans: plans, specs: specs, build: SimsTestData.build, now: now
                )
            }
            .background(FSColor.bg)
            .navigationTitle(Copy.simsTitle)
            .navigationBarTitleDisplayMode(.inline)
        }
    }

    /// The screen as pushed, so the back control shows (root titled `back`).
    private func pushed(back: String, @ViewBuilder _ content: @escaping () -> some View) -> some View {
        NavigationStack(path: .constant([1])) {
            Color.clear.navigationTitle(back).navigationDestination(for: Int.self) { _ in content() }
        }
    }

    private func result(_ phase: SimResultPhase, looks: [Int: SimGearLook] = [:]) -> some View {
        pushed(back: Copy.simsTitle) {
            ScrollView {
                SimResultView(phase: phase, looks: looks, build: SimsTestData.build, now: now)
            }
            .background(FSColor.bg)
            .navigationTitle(resultTitle(phase))
            .navigationBarTitleDisplayMode(.inline)
        }
    }

    private func resultTitle(_ phase: SimResultPhase) -> String {
        if case let .loaded(model, _) = phase { return model.title }
        return Copy.simsTitle
    }

    private var bulkProgress: [String: SimProgress] {
        ["q8wd2c9vh4tz": SimProgress(state: "running", iterationsDone: 0, stage: 1, combosDone: 40, combosTotal: 64)]
    }

    // MARK: F5.1 Sims

    @Test func sims_premium() throws {
        snap(list(.loaded(try SimsTestData.content()), progress: bulkProgress), "sims-premium-ios", height: 1100)
    }

    @Test func sims_premium_large() throws {
        snap(list(.loaded(try SimsTestData.content()), progress: bulkProgress), "sims-premium-large-ios", height: 2600, large: true)
    }

    @Test func sims_free() throws {
        let mine = [SimsTestData.browserRow]
        snap(list(.loaded(try SimsTestData.content(server: [], mine: mine)), entitled: false), "sims-free-ios")
    }

    @Test func sims_lapsed() throws {
        snap(list(.loaded(try SimsTestData.content()), entitled: false, rerun: ["b7nf5x1pa3le": .lapsed], plans: true), "sims-lapsed-ios", height: 1100)
    }

    @Test func sims_runAgainStates() throws {
        let failed = try SimsTestData.row("b7nf5x1pa3le")
        let rows = [RerunStatus.pending, .started, .failed, .lapsed].enumerated().map { index, _ in
            SimRow(
                simId: "f\(index)", spec: failed.spec, kind: failed.kind, state: "error", createdAt: failed.createdAt,
                title: failed.title, lane: "server", finishedAt: failed.finishedAt, failureReason: failed.failureReason, seen: true
            )
        }
        let rerun: [String: RerunStatus] = ["f0": .pending, "f1": .started, "f2": .failed, "f3": .lapsed]
        let content = SimsContent(serverRows: rows, mine: [SimsTestData.browserRow], total: 1, page: 1, offlineSince: nil)
        snap(list(.loaded(content), rerun: rerun), "sims-rerun-states-ios", height: 1300)
    }

    @Test func sims_loading() {
        snap(list(.loading), "sims-loading-ios")
    }

    @Test func sims_empty() throws {
        snap(list(.loaded(try SimsTestData.content(server: [], mine: []))), "sims-empty-ios")
    }

    @Test func sims_error() {
        snap(list(.failed), "sims-error-ios")
    }

    @Test func sims_offline() throws {
        snap(list(.loaded(try SimsTestData.content(offline: RaidTime.date("2026-12-15T04:40:00Z")))), "sims-offline-ios", height: 1100)
    }

    @Test func sims_pageFailed() throws {
        var content = try SimsTestData.content()
        content = SimsContent(serverRows: content.serverRows, mine: content.mine, total: 200, page: 1, offlineSince: nil)
        snap(list(.loaded(content), pageFailed: true), "sims-page-failed-ios", height: 1200)
    }

    // MARK: F5.3 A result

    @Test func result_plainRun() throws {
        let model = SimResultModel(id: "km3xibdrsiyr", result: try SimsTestData.result(), row: try SimsTestData.row("km3xibdrsiyr"), specs: specs)
        snap(result(.loaded(model, offlineSince: nil), looks: try SimsTestData.looks()), "result-run-ios")
    }

    @Test func result_plainRun_large() throws {
        let model = SimResultModel(id: "km3xibdrsiyr", result: try SimsTestData.result(), row: try SimsTestData.row("km3xibdrsiyr"), specs: specs)
        snap(result(.loaded(model, offlineSince: nil), looks: try SimsTestData.looks()), "result-run-large-ios", height: 1800, large: true)
    }

    @Test func result_topGear() throws {
        let model = SimResultModel(id: "g5tq8n2rk1wa", result: try SimsTestData.result(topGear: true, untitled: true), row: SimsTestData.topGearRow, specs: specs)
        snap(result(.loaded(model, offlineSince: nil), looks: try SimsTestData.looks()), "result-gear-ios")
    }

    @Test func result_offline() throws {
        let model = SimResultModel(id: "km3xibdrsiyr", result: try SimsTestData.result(), row: try SimsTestData.row("km3xibdrsiyr"), specs: specs)
        snap(result(.loaded(model, offlineSince: RaidTime.date("2026-12-15T04:40:00Z")), looks: try SimsTestData.looks()), "result-offline-ios")
    }

    @Test func result_loading() {
        snap(result(.loading), "result-loading-ios")
    }

    @Test func result_failed() {
        snap(result(.failed), "result-failed-ios")
    }
}
