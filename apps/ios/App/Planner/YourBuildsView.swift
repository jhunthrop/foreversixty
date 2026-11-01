import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// The `Plan a new build` row: the current character's class and level.
struct PlannerNewBuild: Equatable, Sendable {
    let classSlug: String
    let className: String
    let raceSlug: String?
    let level: Int
}

/// F13.6 Your builds (board F13.6): rows to the viewer, swipe to remove with the Undo snackbar,
/// the footer, and `Plan a new build`. Pure over its inputs so the snapshots draw every state.
struct YourBuildsView: View {
    let phase: YourBuildsPhase
    let removed: RemovedBuild?
    var error: String?
    var newBuild: PlannerNewBuild?
    var timeZone: TimeZone = .current
    var open: (PlannerBuildSummary) -> Void = { _ in }
    var remove: (PlannerBuildSummary) -> Void = { _ in }
    var undo: () -> Void = {}
    var plan: (PlannerNewBuild) -> Void = { _ in }
    var retry: () -> Void = {}

    var body: some View {
        List {
            Text(Copy.planBuildsTitle).font(ScaledType.title).foregroundStyle(FSColor.strong)
                .textCase(.uppercase)
                .accessibilityAddTraits(.isHeader)
                .plainRow(EdgeInsets(top: 8, leading: 16, bottom: 12, trailing: 16))
            content
            if let newBuild {
                MicroLabel(Copy.planBuildsNewHeader)
                    .accessibilityAddTraits(.isHeader)
                    .plainRow(EdgeInsets(top: 22, leading: 16, bottom: 8, trailing: 16))
                newRow(newBuild).plainRow(EdgeInsets(top: 0, leading: 16, bottom: 0, trailing: 16))
            }
        }
        .listStyle(.plain)
        .scrollContentBackground(.hidden)
        .background(FSColor.bg)
        .safeAreaInset(edge: .bottom) { snackbar }
    }

    @ViewBuilder
    private var content: some View {
        switch phase {
        case .loading:
            ForEach(0..<3, id: \.self) { _ in
                RoundedRectangle(cornerRadius: FSRadius.control).fill(FSColor.cardTop).frame(height: 64)
                    .plainRow(EdgeInsets(top: 4, leading: 16, bottom: 4, trailing: 16))
            }
            .accessibilityHidden(true)
        case .failed:
            VStack(alignment: .leading, spacing: 12) {
                Text(Copy.planBuildsFailed).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise", action: retry)
            }
            .plainRow(EdgeInsets(top: 0, leading: 16, bottom: 0, trailing: 16))
        case let .loaded(rows) where rows.isEmpty:
            Text(Copy.planBuildsEmpty).font(ScaledType.body).foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
                .plainRow(EdgeInsets(top: 0, leading: 16, bottom: 0, trailing: 16))
        case let .loaded(rows):
            ForEach(rows) { row in
                YourBuildRow(summary: row, timeZone: timeZone) { open(row) }
                    .swipeActions(edge: .trailing, allowsFullSwipe: true) {
                        Button(role: .destructive) { remove(row) } label: { Text(Copy.planBuildsRemove) }
                    }
                    .accessibilityAction(named: Copy.planBuildsRemove) { remove(row) }
                    .plainRow(EdgeInsets(top: 0, leading: 16, bottom: 0, trailing: 16), separator: true)
            }
            footer
        }
    }

    private var footer: some View {
        VStack(alignment: .leading, spacing: 6) {
            if let error {
                Text(error).font(ScaledType.meta).foregroundStyle(FSColor.wipe).fixedSize(horizontal: false, vertical: true)
            }
            Text(Copy.planBuildsFooter).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
        .plainRow(EdgeInsets(top: 12, leading: 16, bottom: 0, trailing: 16))
    }

    private func newRow(_ build: PlannerNewBuild) -> some View {
        Button { plan(build) } label: {
            HStack(spacing: 12) {
                Image(systemName: "plus").font(.system(size: 16, weight: .semibold)).foregroundStyle(FSColor.gold)
                    .frame(width: 36, height: 36).accessibilityHidden(true)
                Text(Copy.planBuildsNewRow(className: build.className, n: build.level))
                    .font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    .frame(maxWidth: .infinity, alignment: .leading)
                Image(systemName: "chevron.right").foregroundStyle(FSColor.muted).accessibilityHidden(true)
            }
            .padding(.horizontal, 14)
            .frame(minHeight: 60)
            .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line, lineWidth: 1))
            .contentShape(.rect)
        }
        .buttonStyle(.plain)
    }

    @ViewBuilder
    private var snackbar: some View {
        if removed != nil {
            HStack {
                Text(Copy.planRemoved).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Spacer()
                Button(Copy.planRemovedUndo, action: undo)
                    .font(ScaledType.bodyStrong).foregroundStyle(FSColor.gold)
                    .frame(minHeight: 44)
            }
            .padding(.horizontal, 16)
            .frame(minHeight: 52)
            .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line, lineWidth: 1))
            .padding(.horizontal, 16)
            .padding(.bottom, 8)
        }
    }
}

/// One saved build: crest 36, the title, its meta, `{n} views` in mono, the chevron.
struct YourBuildRow: View {
    let summary: PlannerBuildSummary
    let timeZone: TimeZone
    let open: () -> Void

    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let isLarge = typeSize >= .accessibility1
        Button(action: open) {
            HStack(alignment: isLarge ? .top : .center, spacing: 12) {
                ClassCrest(className: summary.classSlug, size: 36)
                VStack(alignment: .leading, spacing: 3) {
                    Text(summary.rowTitle).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                        .fixedSize(horizontal: false, vertical: true)
                    Text(summary.rowMeta(timeZone: timeZone)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                        .fixedSize(horizontal: false, vertical: true)
                    if isLarge { views }
                }
                .frame(maxWidth: .infinity, alignment: .leading)
                if !isLarge { views }
                Image(systemName: "chevron.right").foregroundStyle(FSColor.muted).accessibilityHidden(true)
            }
            .padding(.vertical, 12)
            .frame(minHeight: 64)
            .contentShape(.rect)
        }
        .buttonStyle(.plain)
        .accessibilityElement(children: .combine)
    }

    private var views: some View {
        Text(Copy.planBuildsViews(n: Int(summary.build.views))).font(ScaledType.monoMeta).foregroundStyle(FSColor.nav)
    }
}

private extension View {
    func plainRow(_ insets: EdgeInsets, separator: Bool = false) -> some View {
        listRowInsets(insets)
            .listRowBackground(FSColor.bg)
            .listRowSeparator(separator ? .visible : .hidden)
            .listRowSeparatorTint(FSColor.lineSoft)
    }
}

/// Your builds as pushed: loads, removes with the 5 s snackbar, and opens rows and the editor.
struct YourBuildsScreen: View {
    @State private var model: YourBuildsModel
    @State private var shown: String?
    let dependencies: PlannerDependencies
    let router: PlannerRouter
    let newBuild: PlannerNewBuild?

    init(dependencies: PlannerDependencies, router: PlannerRouter, newBuild: PlannerNewBuild?) {
        _model = State(initialValue: YourBuildsModel(source: dependencies.source))
        self.dependencies = dependencies
        self.router = router
        self.newBuild = newBuild
    }

    var body: some View {
        YourBuildsView(
            phase: model.phase, removed: model.removed, error: model.error, newBuild: newBuild,
            open: { shown = $0.id },
            remove: { row in Task { await model.remove(row) } },
            undo: { Task { await model.undo() } },
            plan: { build in
                router.openEditor(PlannerLaunch(classSlug: build.classSlug, raceSlug: build.raceSlug, level: build.level))
            },
            retry: { Task { await model.load() } }
        )
        .navigationTitle("")
        .navigationBarTitleDisplayMode(.inline)
        .toolbarBackground(FSColor.bg, for: .navigationBar)
        .navigationDestination(item: $shown) { id in
            SharedBuildScreen(id: id, dependencies: dependencies, router: router)
        }
        .task { await model.load() }
        .task(id: model.removed?.summary.id) {
            guard model.removed != nil else { return }
            try? await Task.sleep(for: .seconds(YourBuildsModel.snackbarSeconds))
            if !Task.isCancelled { model.dismissSnackbar() }
        }
    }
}
