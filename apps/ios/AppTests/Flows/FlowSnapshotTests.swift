@testable import ForeverSixty
import DesignSystem
import Models
import Networking
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// Each step of the flows (Android's screenshots/flows), on iPhone 17 at the default size, recorded
/// like RaidSnapshotTests (on CI, `SNAPSHOT_TESTING_RECORD`): `raid-signup-NN-step`,
/// `home-signup-NN-step`, `raid-link-NN-step`. The flows' assertions are in the Flow suites; these
/// walk the same steps only to capture them.
@MainActor
@Suite("Flow snapshots")
struct FlowSnapshotTests {
    static let phone = CGSize(width: 402, height: 874)
    static let safeArea = UIEdgeInsets(top: 62, left: 0, bottom: 34, right: 0)

    private func snap(_ view: some View, _ name: String) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = .large
        }
        let config = ViewImageConfig(safeArea: Self.safeArea, size: Self.phone, traits: traits)
        let sized = view
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.raidIconBuild, HomeFlowTestData.build)
            .environment(\.dynamicTypeSize, .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(
                precision: SnapshotTolerance.precision, perceptualPrecision: SnapshotTolerance.perceptualPrecision,
                layout: .device(config: config), traits: traits
            ),
            testName: name
        )
    }

    private func tab(_ flow: RaidTabFlow, _ step: String) {
        snap(NavigationStack { RaidRootView(content: flow.root) }, "raid-signup-\(step)")
    }

    private func home(_ flow: HomeCardFlow, _ step: String) {
        let content = HomeFlowContent(head: HomeFlowPresenter(now: RaidTime.now, clock: RaidTime.clock).head(flow.model.character), phase: .loaded(flow.blocks))
        snap(HomeFlowView(content: content, build: HomeFlowTestData.build), "home-signup-\(step)")
    }

    private func page(_ controller: SignupController, _ step: String) {
        snap(
            NavigationStack {
                RaidPageContent(
                    controller: controller, clock: RaidTime.clock, now: { RaidTime.now }, currentKey: nil,
                    offlineSince: nil, guildName: "OLYMPUS XXVII", openNight: { _ in }
                )
            },
            "raid-link-\(step)"
        )
    }

    @Test func raidSignup() async throws {
        let flow = try RaidTabFlow()
        await flow.load()
        tab(flow, "01-in")
        await flow.answer(.late)
        tab(flow, "02-late")
        await flow.late(30)
        tab(flow, "03-late-30")
        await flow.answer(.bench)
        tab(flow, "04-bench")
        await flow.answer(.out)
        tab(flow, "05-out")
    }

    @Test func raidSignupFailure() async throws {
        let flow = try RaidTabFlow()
        await flow.load()
        let gate = FlowGate()
        flow.service.gate = gate
        flow.service.failWith = .unexpectedStatus(500)
        let write = flow.startAnswer(.late)
        _ = await flowSettle { flow.service.inFlight == 1 }
        tab(flow, "06-late-saving")
        gate.open()
        await write.value
        tab(flow, "07-late-failed")
        flow.service.gate = nil
        flow.service.failWith = nil
        flow.service.offline = true
        await flow.load()
        tab(flow, "08-offline")
    }

    /// Brand B1: the answers as one segmented group, at the default size and stacked at AX5.
    @Test func segmented() {
        let options = SignupState.allCases.map { FSSegmented.Option($0, $0.label) }
        let group = FSSegmented("Answers", options: options, selection: SignupState.late) { _ in }.padding(16)
        snap(group, "segmented-answers")
        assertSnapshot(
            of: group.environment(\.dynamicTypeSize, .accessibility5).preferredColorScheme(.dark),
            as: .image(
                precision: SnapshotTolerance.precision, perceptualPrecision: SnapshotTolerance.perceptualPrecision,
                layout: .fixed(width: Self.phone.width, height: 600)
            ),
            testName: "segmented-answers-ax5"
        )
    }

    @Test func homeSignup() async throws {
        let flow = try HomeCardFlow()
        await flow.load()
        home(flow, "01-in")
        await flow.model.answer(.late)
        home(flow, "02-late")
        await flow.model.late(30)
        home(flow, "03-late-30")
        let gate = FlowGate()
        flow.service.gate = gate
        flow.service.failWith = .unexpectedStatus(500)
        let write = Task { await flow.model.answer(.bench) }
        _ = await flowSettle { flow.service.inFlight == 1 }
        home(flow, "04-bench-saving")
        gate.open()
        await write.value
        home(flow, "05-bench-failed")
    }

    private func signupDeps(_ service: EchoRaidService) -> SignupController.Dependencies {
        SignupController.Dependencies(
            service: service, haptics: RecordingHaptics(), permission: FakePermission(undetermined: false),
            askStore: MemoryAskStore(), now: { RaidTime.now }, hold: nil
        )
    }

    @Test func raidLink() async throws {
        let flow = try RaidTabFlow()
        await flow.load()
        snap(NavigationStack { RaidRootView(content: flow.root) }, "raid-link-01-tab")

        let late = RaidPageModel(
            guild: FlowPlayer.guild, raid: FlowPlayer.raid, deps: signupDeps(flow.service), cache: NoRaidCache(),
            openAction: RaidOpenAction(state: .late)
        )
        await late.load()
        if case let .loaded(controller) = late.state { page(controller, "02-raid-late-saved") }

        let service = try EchoRaidService()
        service.raidOverride = try RaidFixture.raid("guild.raid.ask-spec")
        let ask = RaidPageModel(
            guild: FlowPlayer.guild, raid: FlowPlayer.raid, deps: signupDeps(service), cache: NoRaidCache(),
            openAction: RaidOpenAction(state: .signedIn)
        )
        await ask.load()
        if case let .loaded(controller) = ask.state { page(controller, "03-spec-ask") }
    }
}
