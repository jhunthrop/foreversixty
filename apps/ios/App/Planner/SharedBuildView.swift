import DesignSystem
import DesignTokens
import Models
import Routing
import SwiftUI

/// F13.5 A shared build: the head, Edit a copy, and the trees drawn read only from the build's
/// own `tree_version` (spec F13.5, board F13.5). Pure over its phase so snapshots draw each state.
struct SharedBuildView: View {
    let phase: SharedBuildPhase
    @Binding var tree: Int
    var timeZone: TimeZone = .current
    var editCopy: (PlannerBuildSummary) -> Void = { _ in }
    var retry: () -> Void = {}

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                Text(Copy.planSharedTitle).font(ScaledType.title).foregroundStyle(FSColor.strong)
                    .textCase(.uppercase)
                    .accessibilityAddTraits(.isHeader)
                    .padding(EdgeInsets(top: 8, leading: 16, bottom: 12, trailing: 16))
                content
            }
            .padding(.bottom, 40)
        }
        .background(FSColor.bg)
    }

    @ViewBuilder
    private var content: some View {
        switch phase {
        case .loading:
            TalentsSkeleton()
        case .notFound:
            message(Copy.planSharedNotFound, retry: false)
        case .offline:
            message(Copy.planSharedOffline, retry: true)
        case .failed:
            message(Copy.planSharedFailed, retry: true)
        case let .loaded(content):
            loaded(content)
        }
    }

    private func message(_ text: String, retry showsRetry: Bool) -> some View {
        VStack(alignment: .leading, spacing: 12) {
            Text(text).font(ScaledType.body).foregroundStyle(showsRetry ? FSColor.wipe : FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
            if showsRetry { SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise", action: retry) }
        }
        .padding(.horizontal, 16)
    }

    @ViewBuilder
    private func loaded(_ content: SharedBuildContent) -> some View {
        let summary = content.summary
        VStack(alignment: .leading, spacing: 14) {
            HStack(alignment: .center, spacing: 12) {
                ClassCrest(className: summary.classSlug, size: 44)
                VStack(alignment: .leading, spacing: 3) {
                    Text(summary.head).font(ScaledType.name).foregroundStyle(FSColor.strong)
                        .fixedSize(horizontal: false, vertical: true)
                    Text(summary.viewerMeta(timeZone: timeZone)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                        .fixedSize(horizontal: false, vertical: true)
                }
            }
            .accessibilityElement(children: .combine)
            Button(Copy.planEdit) { editCopy(summary) }
                .buttonStyle(BrandButtonStyle(.primary, height: 48))
                .disabled(summary.classSlug.isEmpty)
        }
        .padding(.horizontal, 16)
        if let file = content.file {
            SharedBuildTrees(file: file, order: summary.build.pointOrder, tree: $tree)
        }
    }
}

/// The build's trees, read only: F6's drawing with the build's ranks.
struct SharedBuildTrees: View {
    let file: TalentsFile
    let order: [Int]
    @Binding var tree: Int

    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let index = PlannerIndex(file)
        let ranks = PlannerRules.ranks(order)
        let totals = PlannerRules.pointsPerTree(index, order: order)
        let tabs = index.trees.enumerated().map { TalentTabModel(index: $0.offset, name: $0.element.name, points: totals[$0.offset], differing: 0) }
        VStack(spacing: 0) {
            TalentCompareBar(showsSegments: false, tabs: tabs, shown: .constant(.yours), tree: $tree, isLarge: typeSize >= .accessibility1)
            if index.trees.indices.contains(tree) {
                let shown = index.trees[tree]
                TalentTreeView(
                    cells: shown.talents.map { talent in
                        TalentCellModel(
                            talent: talent, rank: ranks[talent.id] ?? 0,
                            state: TalentRules.cellState(talent, in: shown, ranks: ranks), delta: nil, guideRank: nil
                        )
                    },
                    connectors: TalentRules.connectors(shown, ranks: ranks), build: file.build, selectedID: nil, open: { _ in }
                )
                .padding(.top, 16)
            }
        }
        .padding(.top, 14)
    }
}

/// The viewer as pushed: the model, the share button, Edit a copy through the router.
struct SharedBuildScreen: View {
    @State private var model: SharedBuildModel
    @State private var tree = 0
    @State private var settled = false
    let router: PlannerRouter

    init(id: String, dependencies: PlannerDependencies, router: PlannerRouter) {
        _model = State(initialValue: SharedBuildModel(id: id, source: dependencies.source, cache: dependencies.cache))
        self.router = router
    }

    var body: some View {
        SharedBuildView(phase: model.phase, tree: $tree, editCopy: router.editCopy, retry: { Task { await model.load() } })
            .navigationTitle("")
            .navigationBarTitleDisplayMode(.inline)
            .toolbarBackground(FSColor.bg, for: .navigationBar)
            .toolbar {
                ToolbarItem(placement: .topBarTrailing) {
                    ShareLink(item: DeepLink.sharedBuild(id: model.id).url) {
                        Image(systemName: "square.and.arrow.up")
                    }
                    .accessibilityLabel(Copy.planShare)
                }
            }
            .task {
                if case .loading = model.phase { await model.load() }
                settle()
            }
    }

    private func settle() {
        guard !settled, case let .loaded(content) = model.phase else { return }
        settled = true
        tree = TalentRules.openingTree(content.summary.totals)
    }
}
