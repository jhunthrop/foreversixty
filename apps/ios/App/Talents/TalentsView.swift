import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// The Talents screen's content under the bar, pure over its inputs so the snapshot tests draw
/// every state (spec F6 §5, §6).
struct TalentsView: View {
    let phase: TalentsPhase
    let build: String
    let offlineStamp: String?
    @Binding var shown: TalentBuildKind
    @Binding var tree: Int
    var selectedID: Int?
    var openTalent: (Int) -> Void = { _ in }
    var retry: () -> Void = {}
    var openPlanner: () -> Void = {}
    /// Lane F13: with `planner` on, `Plan a build` opens the app's planner (not the site), and
    /// signed in, `Your builds` is pushed.
    var planner: TalentPlannerEntry?

    @Environment(\.dynamicTypeSize) private var typeSize

    private var isLarge: Bool { typeSize >= .accessibility1 }
    /// Where a tapped cell lands: under the pinned bar and clear of a medium sheet.
    private static let sheetClearAnchor = UnitPoint(x: 0.5, y: 0.3)
    /// Room under the footer, so the tab bar never covers it and a bottom-row cell can scroll up
    /// above an open sheet (spec F6 §5.1 item 9).
    private static let bottomSpace: CGFloat = 130
    /// While the sheet is open, room enough for the last tier to scroll above a medium sheet.
    private static let sheetRoom: CGFloat = 480

    private func tier(of id: Int) -> Int? {
        guard case let .loaded(presenter, _) = phase else { return nil }
        return presenter.trees.flatMap(\.talents).first { $0.id == id }?.tier
    }

    var body: some View {
        ScrollViewReader { proxy in
            ScrollView {
                LazyVStack(alignment: .leading, spacing: 0, pinnedViews: isLarge ? [] : [.sectionHeaders]) {
                    if let offlineStamp { OfflineStrip(text: Copy.offlineStrip(time: offlineStamp)) }
                    content
                }
                .padding(.bottom, selectedID == nil ? Self.bottomSpace : Self.sheetRoom)
            }
            .onChange(of: selectedID) { _, id in
                // The tapped cell stays above the sheet (spec F6 §7, point 4).
                guard let id, let tier = tier(of: id) else { return }
                withAnimation { proxy.scrollTo(TalentTreeView.anchor(tier: tier), anchor: Self.sheetClearAnchor) }
            }
        }
        .background(FSColor.bg)
    }

    @ViewBuilder
    private var content: some View {
        switch phase {
        case .loading:
            TalentsSkeleton()
            footer
        case .failed:
            VStack(alignment: .leading, spacing: 12) {
                Text(Copy.talentsFailed).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                    .accessibilityAddTraits(.isStaticText)
                SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise", action: retry)
            }
            .padding(EdgeInsets(top: 8, leading: 16, bottom: 0, trailing: 16))
            footer
        case let .loaded(presenter, _):
            loaded(presenter)
        }
    }

    @ViewBuilder
    private func loaded(_ presenter: TalentsPresenter) -> some View {
        TalentStatusBlock(status: presenter.status, notes: presenter.notes, retry: retry)
            .padding(EdgeInsets(top: 8, leading: 16, bottom: 0, trailing: 16))
        if presenter.hasTree {
            Section {
                TalentTreeView(
                    cells: presenter.cells(shown, tree: tree),
                    connectors: presenter.connectors(shown, tree: tree),
                    build: build,
                    selectedID: selectedID,
                    open: openTalent
                )
                .padding(.top, 16)
                Text(presenter.legend(shown)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
                    .padding(EdgeInsets(top: 8, leading: 16, bottom: 0, trailing: 16))
            } header: {
                TalentCompareBar(
                    showsSegments: presenter.showsSegments,
                    tabs: presenter.tabs(shown),
                    shown: $shown,
                    tree: $tree,
                    isLarge: isLarge
                )
            }
        }
        footer
    }

    private var footer: some View {
        VStack(alignment: .leading, spacing: 4) {
            Text(Copy.talentsFooter).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
            if let planner {
                footerLink(Copy.planPlanBuild, systemImage: "chevron.right", action: planner.plan)
                if let yourBuilds = planner.yourBuilds {
                    footerLink(Copy.planYourBuilds, systemImage: "chevron.right", action: yourBuilds)
                }
            } else {
                footerLink(Copy.talentsPlanLink, systemImage: "arrow.up.right.square", action: openPlanner)
                    .accessibilityAddTraits(.isLink)
            }
        }
        .padding(EdgeInsets(top: 16, leading: 16, bottom: 0, trailing: 16))
    }
}

extension TalentsView {
    fileprivate func footerLink(_ title: String, systemImage: String, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            HStack(alignment: .firstTextBaseline, spacing: 6) {
                Text(title).multilineTextAlignment(.leading)
                Image(systemName: systemImage).font(.system(size: 14)).accessibilityHidden(true)
            }
            .font(ScaledType.bodyStrong)
            .foregroundStyle(FSColor.gold)
            .frame(minHeight: 48, alignment: .leading)
        }
        .buttonStyle(.plain)
    }
}

/// The Talents screen's planner rows (lane F13): Plan a build, and Your builds when signed in.
struct TalentPlannerEntry {
    let plan: () -> Void
    let yourBuilds: (() -> Void)?
}

/// The status block, the screen's one primary element, then the secondary notes (spec F6 §5.1).
struct TalentStatusBlock: View {
    let status: TalentStatusLine
    let notes: [String]
    var retry: () -> Void = {}

    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        VStack(alignment: .leading, spacing: 4) {
            Group {
                if typeSize >= .accessibility1 {
                    VStack(alignment: .leading, spacing: 6) { tag; sentence }
                } else {
                    HStack(alignment: .firstTextBaseline, spacing: 10) { tag; sentence }
                }
            }
            .frame(minHeight: 28, alignment: .leading)
            .accessibilityElement(children: .ignore)
            .accessibilityLabel(status.accessibilityLabel)
            if status.isError {
                SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise", action: retry).padding(.top, 4)
            }
            ForEach(notes, id: \.self) { note in
                Text(note).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
    }

    @ViewBuilder
    private var tag: some View {
        if let text = status.tag {
            TagView(text: text, color: color)
        }
    }

    private var sentence: some View {
        Text(status.sentence)
            .font(ScaledType.body)
            .foregroundStyle(status.isError ? FSColor.wipe : FSColor.strong)
            .fixedSize(horizontal: false, vertical: true)
    }

    private var color: Color {
        switch status.tagKind {
        case .matches: FSColor.kill
        case .differs: FSColor.gold
        case .muted: FSColor.muted
        }
    }
}

/// The two segments and the three tree tabs, pinned under the bar below AX1 (spec F6 §5.1 item 5).
struct TalentCompareBar: View {
    let showsSegments: Bool
    let tabs: [TalentTabModel]
    @Binding var shown: TalentBuildKind
    @Binding var tree: Int
    let isLarge: Bool

    var body: some View {
        VStack(spacing: 8) {
            if showsSegments { segments }
            if isLarge {
                stackedTabs
            } else {
                ViewThatFits(in: .horizontal) {
                    rowTabs
                    stackedTabs
                }
            }
        }
        .padding(EdgeInsets(top: 12, leading: 16, bottom: 8, trailing: 16))
        .background(FSColor.bg)
        .overlay(alignment: .bottom) { if !isLarge { FSColor.lineSoft.frame(height: 1) } }
        .sensoryFeedback(.selection, trigger: shown)
        .sensoryFeedback(.selection, trigger: tree)
    }

    private var segments: some View {
        let layout = isLarge ? AnyLayout(VStackLayout(spacing: 0)) : AnyLayout(HStackLayout(spacing: 0))
        return layout {
            segment(.guide, Copy.talentsGuide)
            segment(.yours, Copy.talentsYours)
        }
        .background(FSColor.raised)
        .clipShape(.rect(cornerRadius: FSRadius.control))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line, lineWidth: 1))
    }

    private func segment(_ kind: TalentBuildKind, _ title: String) -> some View {
        let checked = shown == kind
        return Button { shown = kind } label: {
            Text(title)
                .font(ScaledType.segment)
                .foregroundStyle(checked ? FSColor.strong : FSColor.muted)
                .frame(maxWidth: .infinity, minHeight: 44)
                .background(checked ? FSColor.gold.opacity(0.12) : .clear)
                .contentShape(Rectangle())
        }
        .buttonStyle(.plain)
        .accessibilityAddTraits(checked ? .isSelected : [])
    }

    private var rowTabs: some View {
        HStack(spacing: Self.tabGutter) {
            ForEach(tabs) { tab in
                Button { tree = tab.index } label: {
                    tabLabel(tab)
                        .lineLimit(1)
                        .fixedSize()
                        .frame(maxWidth: .infinity, minHeight: 44)
                        .overlay(alignment: .bottom) {
                            if tree == tab.index { FSColor.gold.frame(height: 2) }
                        }
                        .contentShape(Rectangle())
                }
                .buttonStyle(.plain)
                .tabAccessibility(tab, selected: tree == tab.index)
            }
        }
        .overlay(alignment: .bottom) { FSColor.line.frame(height: 1) }
    }

    /// The least space between two row tabs; with less, the tabs stack (F6 ux re-check).
    private static let tabGutter: CGFloat = 12
    /// A stacked tab is 56 tall at large text, 44 when it stacks only for a narrow screen.
    private var stackedRowHeight: CGFloat { isLarge ? 56 : 44 }

    private var stackedTabs: some View {
        VStack(spacing: 0) {
            ForEach(tabs) { tab in
                if tab.index > 0 { FSColor.lineSoft.frame(height: 1) }
                Button { tree = tab.index } label: {
                    HStack {
                        Text(tab.name).font(ScaledType.bodyStrong)
                            .foregroundStyle(tree == tab.index ? FSColor.strong : FSColor.nav)
                        Spacer(minLength: 8)
                        Text("\(tab.points)").font(ScaledType.mono)
                            .foregroundStyle(tree == tab.index ? FSColor.strong : FSColor.nav)
                        if tab.differing > 0 { DiffDot() }
                    }
                    .padding(.horizontal, 14)
                    .frame(maxWidth: .infinity, minHeight: stackedRowHeight)
                    .background(tree == tab.index ? FSColor.gold.opacity(0.12) : .clear)
                    .overlay(alignment: .leading) { if tree == tab.index { FSColor.gold.frame(width: 3) } }
                    .contentShape(Rectangle())
                }
                .buttonStyle(.plain)
                .tabAccessibility(tab, selected: tree == tab.index)
            }
        }
        .background(FSColor.raised)
        .clipShape(.rect(cornerRadius: FSRadius.control))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line, lineWidth: 1))
    }

    private func tabLabel(_ tab: TalentTabModel) -> some View {
        let selected = tree == tab.index
        return HStack(spacing: 6) {
            Text(tab.name).font(selected ? ScaledType.bodyStrong : ScaledType.body)
            Text("\(tab.points)").font(ScaledType.mono)
            if tab.differing > 0 { DiffDot() }
        }
        .foregroundStyle(selected ? FSColor.strong : FSColor.nav)
    }
}

/// The 8 pt wipe dot on a tab whose tree holds a dashed cell.
private struct DiffDot: View {
    var body: some View {
        Circle().fill(FSColor.wipe).frame(width: 8, height: 8).accessibilityHidden(true)
    }
}

private extension View {
    func tabAccessibility(_ tab: TalentTabModel, selected: Bool) -> some View {
        accessibilityElement(children: .ignore)
            .accessibilityLabel(tab.name)
            .accessibilityValue(tab.accessibilityValue)
            .accessibilityAddTraits(selected ? [.isSelected, .isButton] : .isButton)
    }
}

/// First load with nothing saved: skeletons at the real sizes, never a spinner (spec F6 §6).
struct TalentsSkeleton: View {
    private static let pattern = [
        TalentCell(tier: 0, column: 1), TalentCell(tier: 0, column: 2), TalentCell(tier: 1, column: 0),
        TalentCell(tier: 1, column: 1), TalentCell(tier: 1, column: 2), TalentCell(tier: 1, column: 3),
        TalentCell(tier: 2, column: 0), TalentCell(tier: 2, column: 1), TalentCell(tier: 3, column: 0),
        TalentCell(tier: 3, column: 1), TalentCell(tier: 3, column: 3), TalentCell(tier: 4, column: 1),
        TalentCell(tier: 5, column: 2), TalentCell(tier: 6, column: 1),
    ]

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            bar(72, 22)
            bar(220, 14)
            RoundedRectangle(cornerRadius: FSRadius.control).fill(FSColor.cardTop).frame(height: 44).padding(.top, 12)
            HStack { bar(100, 14); Spacer(); bar(100, 14); Spacer(); bar(100, 14) }.frame(height: 44)
            GeometryReader { proxy in
                ForEach(Array(Self.pattern.enumerated()), id: \.offset) { _, cell in
                    let origin = TalentGeometry.origin(cell, width: proxy.size.width + 32)
                    RoundedRectangle(cornerRadius: FSRadius.control).fill(FSColor.cardTop)
                        .frame(width: TalentGeometry.cell, height: TalentGeometry.cell)
                        .offset(x: origin.x - 16, y: origin.y)
                }
            }
            .frame(height: TalentGeometry.height)
            .padding(.top, 16)
            Text(Copy.talentsLoading).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
        .padding(EdgeInsets(top: 8, leading: 16, bottom: 0, trailing: 16))
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.talentsLoading)
    }

    private func bar(_ width: CGFloat, _ height: CGFloat) -> some View {
        RoundedRectangle(cornerRadius: 4).fill(FSColor.cardTop).frame(width: width, height: height)
    }
}
