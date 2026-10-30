@testable import ForeverSixty
import DesignTokens
import Models
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// F2's states on the iPhone 17 (402 x 874 pt) at the default size and at AX5, from the contract
/// fixtures with the repo's own icons. The same list is drawn on Android (`HomeScreenshotTest.kt`).
/// Re-record with `SNAPSHOT_TESTING_RECORD=all`; the PNGs are also the review captures
/// (design/reviews/2026-10-31-mobile-home).
@MainActor
@Suite("Home snapshots")
struct HomeSnapshotTests {
    private static let iPhone17 = ViewImageConfig(
        safeArea: UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0),
        size: CGSize(width: 402, height: 874),
        traits: UITraitCollection(userInterfaceStyle: .dark)
    )

    private static let iPhone17Tall = ViewImageConfig(
        safeArea: UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0),
        size: CGSize(width: 402, height: 2100),
        traits: UITraitCollection(userInterfaceStyle: .dark)
    )

    private func snap(_ view: some View, large: Bool = false, tall: Bool = false, named name: String) {
        let framed = view
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: framed,
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: tall ? Self.iPhone17Tall : Self.iPhone17)),
            testName: name
        )
    }

    private func home(_ phase: HomePhase, offline: String? = nil) -> some View {
        HomeCaptureFrame {
            HomeAnswerView(
                phase: phase,
                presenter: HomeFixtures.presenter(),
                build: "1.60.1.70658",
                offlineStamp: offline
            )
        }
    }

    @Test("F2.1 the answer, default and AX5")
    func answer() throws {
        let phase = HomePhase.loaded(try HomeFixtures.home(), offlineSince: nil)
        snap(home(phase), named: "home-default")
        snap(home(phase), large: true, named: "home-default-xl")
    }

    @Test("F2.1 the answer, full length")
    func answerFull() throws {
        snap(home(.loaded(try HomeFixtures.home(), offlineSince: nil)), tall: true, named: "home-full")
    }

    @Test("F2.2 loading")
    func loading() { snap(home(.loading), named: "home-loading") }

    @Test("F2.3 offline with the stamp")
    func offline() throws {
        let phase = HomePhase.loaded(try HomeFixtures.home(), offlineSince: HomeFixtures.now)
        snap(home(phase, offline: "Thu 9:41 pm"), named: "home-offline")
    }

    @Test("F2.4 no gear export")
    func noGear() throws {
        snap(home(.loaded(try HomeFixtures.home("character.home.no-gear"), offlineSince: nil)), named: "home-nogear")
    }

    @Test("the empty states: under 20, no spec, no list, every slot matches", arguments: ["under_20", "no_spec", "no_list"])
    func emptyStates(state: String) throws {
        snap(home(.loaded(try HomeFixtures.home(state: state), offlineSince: nil)), named: "home-\(state.replacingOccurrences(of: "_", with: ""))")
    }

    @Test("every slot matches")
    func allMatch() throws {
        snap(home(.loaded(try HomeFixtures.home("character.home.empty"), offlineSince: nil)), named: "home-allmatch")
    }

    @Test("F2.5 error, with the cached Next rows")
    func error() throws {
        snap(home(.failed(cached: try HomeFixtures.home())), named: "home-error")
    }

    @Test("F2.6 the tooltip at the medium and large detents, default and AX5")
    func tooltip() throws {
        let request = ItemTooltipRequest.forHome(
            HomeItemLook(id: 252_605, name: "Black Dragonscale Helm", quality: 4, icon: nil, itemLevel: 61),
            home: try HomeFixtures.home(), character: HomeFixtures.velaren, sourceLine: "Crafted: Leatherworking", wornID: 279_253
        )
        let model = ItemTooltipModel.build(request, sources: try HomeFixtures.tooltipSources())
        let behind = home(.loaded(try HomeFixtures.home(), offlineSince: nil))
        snap(SheetFrame(height: 470, behind: behind) { ItemTooltipContent(request: request, phase: .loaded(model)) }, named: "tooltip-medium")
        snap(SheetFrame(height: 800, behind: behind) { ItemTooltipContent(request: request, phase: .loaded(model)) }, named: "tooltip-large")
        snap(SheetFrame(height: 800, behind: behind) { ItemTooltipContent(request: request, phase: .loaded(model)) }, large: true, named: "tooltip-large-xl")
    }

    @Test("F2.6 a vendor item with a standing, and a weapon")
    func tooltipVariants() throws {
        let sources = try HomeFixtures.tooltipSources()
        for (id, name) in [(276_202, "tooltip-vendor"), (250_619, "tooltip-weapon")] {
            let item = try #require(sources.items[id])
            let request = ItemTooltipRequest.forHome(
                HomeItemLook(id: id, name: item.name, quality: item.quality, icon: item.icon, itemLevel: item.itemLevel),
                home: try HomeFixtures.home(), character: HomeFixtures.velaren
            )
            let model = ItemTooltipModel.build(request, sources: sources)
            snap(SheetFrame(height: 800, behind: Color.clear) { ItemTooltipContent(request: request, phase: .loaded(model)) }, named: name)
        }
    }

    @Test("F2.7 Your gear, all 17 slots")
    func gear() throws {
        let home = try HomeFixtures.home()
        let sources = try HomeFixtures.tooltipSources()
        let probe = ItemTooltipRequest.forHome(HomeItemLook(id: 0, name: nil, quality: nil, icon: nil, itemLevel: nil), home: home, character: HomeFixtures.velaren)
        let blocker = GearModel.offHandBlocker(home: home, gear: [:], items: sources.items, band: ItemTooltipModel.band(for: probe, in: sources.bis))
        let rows = GearModel.rows(home: home, gear: [:], items: sources.items, offHandBlockedBy: blocker)
        let view = HomeCaptureFrame(title: Copy.gearTitle, large: true) {
            GearView(rows: rows, intro: Copy.gearIntro(band: "60 · raid"), build: home.dataBuild)
        }
        snap(view, named: "gear-default")
        snap(view, large: true, named: "gear-default-xl")
        snap(view, tall: true, named: "gear-full")
    }
}

/// The navigation chrome around a captured screen. The tab bar is left out: the iOS 26 glass
/// bar does not render in an off-screen snapshot.
struct HomeCaptureFrame<Content: View>: View {
    var title = "Velaren"
    var large = false
    @ViewBuilder let content: Content

    var body: some View {
        NavigationStack {
            if large {
                content
            } else {
                ScrollView { content }
                    .background(FSColor.bg)
                    .navigationTitle(title)
                    .navigationBarTitleDisplayMode(.inline)
            }
        }
        .tint(FSColor.gold)
    }
}

/// A sheet at a detent's height over the dimmed screen, for the tooltip captures.
struct SheetFrame<Behind: View, Sheet: View>: View {
    let height: CGFloat
    let behind: Behind
    @ViewBuilder let sheet: Sheet

    var body: some View {
        ZStack(alignment: .bottom) {
            behind
            Color.black.opacity(0.45).ignoresSafeArea()
            VStack(spacing: 0) {
                Capsule().fill(FSColor.muted.opacity(0.6)).frame(width: 36, height: 5).padding(.top, 6)
                sheet
            }
            .frame(height: height)
            .background(FSColor.raised)
            .clipShape(UnevenRoundedRectangle(topLeadingRadius: 36, topTrailingRadius: 36))
        }
        .ignoresSafeArea(edges: .bottom)
    }
}
