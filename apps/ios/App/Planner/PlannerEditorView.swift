import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// The editor's buttons, so the view stays pure over the model and the snapshots draw it whole.
struct PlannerEditorActions {
    var back: () -> Void = {}
    var openRace: () -> Void = {}
    var openLevel: () -> Void = {}
    var reset: () -> Void = {}
    var share: () -> Void = {}
    var retry: () -> Void = {}
}

/// F13.1 Plan a build: the bar, the chips, the tabs, the tree, the stepper strip and the action
/// bar (spec F13.1, board F13.1). A full-screen task: no tab bar.
struct PlannerEditorView: View {
    let model: PlannerEditorModel
    @Binding var tree: Int
    let build: String
    var actions = PlannerEditorActions()

    @Environment(\.dynamicTypeSize) private var typeSize
    private var isLarge: Bool { typeSize >= .accessibility1 }

    var body: some View {
        VStack(spacing: 0) {
            PlannerBar(backTitle: backTitle, canUndo: model.state?.canUndo ?? false, back: actions.back, undo: model.undo)
            switch model.phase {
            case .loading:
                TalentsSkeleton().frame(maxHeight: .infinity, alignment: .top)
            case .failed:
                failed
            case .ready:
                if let state = model.state { ready(state) }
            }
        }
        .background(FSColor.bg)
    }

    private var backTitle: String { model.launch.origin == .guide ? Copy.commonCancel : Copy.talentsTitle }

    private var failed: some View {
        VStack(alignment: .leading, spacing: 12) {
            Text(Copy.talentsFailed).font(ScaledType.body).foregroundStyle(FSColor.wipe)
            SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise", action: actions.retry)
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
    }

    /// Below AX1 the strip and the action bar are pinned over the bottom (the scroll's inset lets
    /// the last tier clear them); at the accessibility sizes the strip grows, so it scrolls above
    /// the tree and only Reset and Share stay pinned (spec §9).
    @ViewBuilder
    private func ready(_ state: PlannerState) -> some View {
        if isLarge {
            ScrollView {
                VStack(spacing: 0) {
                    chips(state)
                    strip(state).padding(.top, 12)
                    tree(state)
                }
            }
            .safeAreaInset(edge: .bottom, spacing: 0) {
                PlannerActionBar(reset: actions.reset, share: actions.share, isLarge: true).background(FSColor.bg)
            }
        } else {
            ScrollView {
                VStack(spacing: 0) {
                    chips(state)
                    tree(state)
                }
            }
            .safeAreaInset(edge: .bottom, spacing: 0) {
                VStack(spacing: 0) {
                    strip(state)
                    PlannerActionBar(reset: actions.reset, share: actions.share, isLarge: false)
                }
                .background(FSColor.bg)
            }
        }
    }

    private func chips(_ state: PlannerState) -> some View {
        PlannerChips(
            classSlug: state.index.file.classSlug, raceName: model.raceName, level: state.level, left: state.left,
            isLarge: isLarge, openRace: actions.openRace, openLevel: actions.openLevel
        )
        .padding(EdgeInsets(top: 8, leading: 16, bottom: 0, trailing: 16))
    }

    private func strip(_ state: PlannerState) -> some View {
        PlannerStrip(state: state, build: build, isLarge: isLarge, add: model.addFocused, remove: model.removeFocused)
    }

    @ViewBuilder
    private func tree(_ state: PlannerState) -> some View {
        TalentCompareBar(showsSegments: false, tabs: tabs(state), shown: .constant(.yours), tree: $tree, isLarge: isLarge)
        TalentTreeView(
            cells: cells(state), connectors: connectors(state), build: build,
            selectedID: nil, open: model.add, remove: model.remove
        )
        .padding(.vertical, 16)
    }

    private var treeIndex: Int { tree }

    private func tabs(_ state: PlannerState) -> [TalentTabModel] {
        let totals = state.totals
        return state.index.trees.enumerated().map { index, tree in
            TalentTabModel(index: index, name: tree.name, points: totals[index], differing: 0)
        }
    }

    private func cells(_ state: PlannerState) -> [TalentCellModel] {
        guard state.index.trees.indices.contains(treeIndex) else { return [] }
        let ranks = state.ranks
        let left = state.left
        return state.index.trees[treeIndex].talents.map { talent in
            TalentCellModel(
                talent: talent, rank: ranks[talent.id] ?? 0,
                state: PlannerRules.cellState(talent, index: state.index, order: state.order, ranks: ranks, left: left),
                delta: nil, guideRank: nil
            )
        }
    }

    private func connectors(_ state: PlannerState) -> [TalentConnector] {
        guard state.index.trees.indices.contains(treeIndex) else { return [] }
        return TalentRules.connectors(state.index.trees[treeIndex], ranks: state.ranks)
    }
}

/// The bar: back (`Talents`, or `Cancel` from a guide) and the 48 pt Undo.
struct PlannerBar: View {
    let backTitle: String
    let canUndo: Bool
    let back: () -> Void
    let undo: () -> Void

    var body: some View {
        HStack {
            Button(action: back) {
                HStack(spacing: 6) {
                    Image(systemName: "chevron.left").font(.system(size: 15, weight: .semibold)).accessibilityHidden(true)
                    Text(backTitle).font(ScaledType.bodyStrong)
                }
                .foregroundStyle(FSColor.strong)
                .padding(.horizontal, 14)
                .frame(minHeight: 44)
                .background(FSColor.raised, in: .rect(cornerRadius: 22))
                .overlay(RoundedRectangle(cornerRadius: 22).strokeBorder(FSColor.line, lineWidth: 1))
            }
            .buttonStyle(.plain)
            Spacer()
            Button(action: undo) {
                Image(systemName: "arrow.uturn.backward")
                    .font(.system(size: 17, weight: .semibold))
                    .foregroundStyle(canUndo ? FSColor.strong : FSColor.muted)
                    .frame(width: 48, height: 48)
                    .background(FSColor.raised, in: .circle)
                    .overlay(Circle().strokeBorder(FSColor.line, lineWidth: 1))
            }
            .buttonStyle(.plain)
            .disabled(!canUndo)
            .accessibilityLabel(Copy.planUndo)
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 6)
    }
}

/// The chips row: the class crest 30, Race and Level, and the points left (spec F13.1).
struct PlannerChips: View {
    let classSlug: String
    let raceName: String?
    let level: Int
    let left: Int
    let isLarge: Bool
    let openRace: () -> Void
    let openLevel: () -> Void

    var body: some View {
        if isLarge {
            VStack(alignment: .leading, spacing: 8) {
                HStack(spacing: 10) { crest; Spacer(minLength: 8); points }
                if raceName != nil { raceChip }
                levelChip
            }
        } else {
            HStack(spacing: 8) {
                crest
                if raceName != nil { raceChip }
                levelChip
                Spacer(minLength: 8)
                points
            }
        }
    }

    private var crest: some View {
        ClassCrest(className: classSlug, size: 30)
            .accessibilityHidden(false)
            .accessibilityLabel(PlannerCopy.className(classSlug, names: .empty))
    }

    private var raceChip: some View {
        PlannerChip(title: raceName ?? "", accessibilityLabel: Copy.planRaceSheet, action: openRace)
    }

    private var levelChip: some View {
        PlannerChip(title: Copy.planLevel(n: level), accessibilityLabel: Copy.planLevelSheet, action: openLevel)
    }

    private var points: some View {
        Text(Copy.planLeft(n: left))
            .font(ScaledType.monoDay)
            .foregroundStyle(left == 0 ? FSColor.kill : FSColor.gold)
            .accessibilityLabel(Copy.planA11yLeft(n: left))
    }
}

/// A chip that opens a sheet: the value and a chevron, 44 tall.
struct PlannerChip: View {
    let title: String
    let accessibilityLabel: String
    let action: () -> Void

    var body: some View {
        Button(action: action) {
            HStack(spacing: 6) {
                Text(title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                Image(systemName: "chevron.down").font(.system(size: 11, weight: .semibold)).foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
            .padding(.horizontal, 12)
            .frame(minHeight: 44)
            .background(FSColor.raised, in: .rect(cornerRadius: 22))
            .overlay(RoundedRectangle(cornerRadius: 22).strokeBorder(FSColor.line, lineWidth: 1))
        }
        .buttonStyle(.plain)
        .accessibilityLabel(accessibilityLabel)
        .accessibilityValue(title)
    }
}

/// The stepper strip (spec F13.1, §3 component table): the strip's talent icon 44, its name and
/// `{rank}/{max}`, its rank text or the refusal or the loaded note, and 48 pt minus and plus.
struct PlannerStrip: View {
    let state: PlannerState
    let build: String
    var isLarge = false
    let add: () -> Void
    let remove: () -> Void

    var body: some View {
        if let talent = state.focus {
            let rank = state.rank(of: talent.id)
            let tile = TalentIconTile(
                talent: talent,
                state: PlannerRules.cellState(talent, index: state.index, order: state.order, ranks: state.ranks, left: state.left),
                build: build, size: 44
            )
            let text = VStack(alignment: .leading, spacing: 3) {
                ViewThatFits(in: .horizontal) {
                    HStack(alignment: .firstTextBaseline, spacing: 8) { name(talent); rankLabel(talent, rank: rank) }
                    VStack(alignment: .leading, spacing: 2) { name(talent); rankLabel(talent, rank: rank) }
                }
                line(talent, rank: rank)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            .accessibilityElement(children: .combine)
            let minus = stepper(systemImage: "minus", label: Copy.planRemovePoint, enabled: rank > 0, action: remove)
            let plus = stepper(systemImage: "plus", label: Copy.planAddPoint, enabled: rank < talent.maxRank, action: add)
            Group {
                if isLarge {
                    VStack(alignment: .leading, spacing: 10) {
                        tile
                        text
                        HStack(spacing: 12) { minus; plus }
                    }
                } else {
                    HStack(alignment: .center, spacing: 12) { tile; text; minus; plus }
                }
            }
            .padding(.horizontal, 16)
            .padding(.vertical, 10)
            .frame(minHeight: 96)
            .background(FSColor.raised)
            .overlay(alignment: .top) { FSColor.lineSoft.frame(height: 1) }
        }
    }

    private func name(_ talent: Talent) -> some View {
        Text(talent.name).font(ScaledType.name).foregroundStyle(FSColor.strong).fixedSize(horizontal: false, vertical: true)
    }

    private func rankLabel(_ talent: Talent, rank: Int) -> some View {
        Text(Copy.planRank(n: rank, max: talent.maxRank)).font(ScaledType.mono).foregroundStyle(FSColor.muted)
    }

    @ViewBuilder
    private func line(_ talent: Talent, rank: Int) -> some View {
        switch state.line {
        case let .refusal(reason):
            Text(reason).font(ScaledType.meta).foregroundStyle(FSColor.wipe).fixedSize(horizontal: false, vertical: true)
        case let .note(note):
            Text(note).font(ScaledType.meta).foregroundStyle(FSColor.nav).fixedSize(horizontal: false, vertical: true)
        case nil:
            if let text = Self.rankText(talent, rank: rank) {
                Text(text).font(ScaledType.meta).foregroundStyle(FSColor.nav).fixedSize(horizontal: false, vertical: true)
            }
        }
    }

    /// The client's text for the rank held, or for the first rank before any point.
    static func rankText(_ talent: Talent, rank: Int) -> String? {
        let position = max(rank, 1) - 1
        return talent.ranks.indices.contains(position) ? talent.ranks[position].description : nil
    }

    private func stepper(systemImage: String, label: String, enabled: Bool, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            Image(systemName: systemImage)
                .font(.system(size: 18, weight: .semibold))
                .foregroundStyle(enabled ? FSColor.strong : FSColor.muted)
                .frame(width: 48, height: 48)
                .background(GameColors.controlFill, in: .circle)
                .overlay(Circle().strokeBorder(FSColor.lineCard, lineWidth: 1))
        }
        .buttonStyle(.plain)
        .accessibilityLabel(label)
    }
}

/// Reset (secondary) and Share (primary, the screen's one gold button), 48 pt.
struct PlannerActionBar: View {
    let reset: () -> Void
    let share: () -> Void
    let isLarge: Bool

    var body: some View {
        // At the accessibility sizes the two share the width so Reset never breaks mid-word.
        HStack(spacing: 10) {
            Button(Copy.planReset, action: reset)
                .buttonStyle(BrandButtonStyle(.quiet, height: 48))
                .frame(maxWidth: isLarge ? .infinity : 120)
            Button(action: share) {
                Label(Copy.planShare, systemImage: "square.and.arrow.up")
            }
            .buttonStyle(BrandButtonStyle(.primary, height: 48))
        }
        .padding(EdgeInsets(top: 10, leading: 16, bottom: 8, trailing: 16))
        // Pinned, so it stops growing at AX1: both labels stay whole and the tree keeps the screen.
        .dynamicTypeSize(...DynamicTypeSize.accessibility1)
    }
}
