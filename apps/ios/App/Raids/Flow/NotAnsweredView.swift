import DesignSystem
import SwiftUI

/// What Not answered hands back.
struct NotAnsweredActions {
    var select: (NotAnsweredFilter) -> Void = { _ in }
    var mentionRole: () -> Void = {}
    var mentionAll: () -> Void = {}
    var mark: (NotAnsweredRow, SignupState) -> Void = { _, _ in }
    var undo: (NotAnsweredRow) -> Void = { _ in }
    var message: (URL) -> Void = { _ in }
    var editOnSite: () -> Void = {}
}

/// Keep the raid full (flows spec F3, mock ios-raid-unanswered): the raid page filtered to Not
/// answered, the short role first; the mention that names it with its preview; per row Mark In
/// (it confirms on the row with Undo), Bench, Discord DM; the bench that can close the gap.
struct NotAnsweredView: View {
    let content: NotAnsweredContent
    let build: String
    /// What the last mention did (`Mentioned 2 raiders in Discord.`), or why it did not.
    var status: String?
    var mentioning = false
    var actions = NotAnsweredActions()

    /// `.preview`: 3 pt rule, 10 by 12 inside, 12 above.
    static let previewRule: CGFloat = 3
    static let previewTop: CGFloat = 12
    static let siteTop: CGFloat = 16

    var body: some View {
        RaidFlowScroll {
            RaidFlowHead(title: content.title, facts: content.facts) {
                ItemIconView(icon: content.icon, quality: nil, build: build, size: FSSize.crest56).parityItem(.image)
            }
            chips
            if let mention = content.mention { mentionBlock(mention) }
            list
            if let benchTitle = content.benchTitle { bench(benchTitle) }
            if content.actionable {
                let link = Text(Copy.raidNaSiteLink).foregroundStyle(FSColor.gold).font(.custom(FSFont.bodySemibold, size: 13))
                Text("\(Text(Copy.raidNaSite).foregroundStyle(FSColor.muted)) \(link)")
                    .type(FSType.small)
                    .onTapGesture(perform: actions.editOnSite)
                    .accessibilityAddTraits(.isLink)
                    .padding(.top, Self.siteTop)
            }
        }
    }

    private var chips: some View {
        ScrollView(.horizontal, showsIndicators: false) {
            HStack(spacing: FSSpace.space2) {
                ForEach(content.chips) { chip in
                    // The status filters are choice chips, never primary buttons (one primary a screen).
                    Button(chip.title) { actions.select(chip.filter) }
                        .buttonStyle(MinuteChipStyle(on: chip.filter == content.filter, height: FSButtonStyle.smallHeight, fillsWidth: false))
                        // Each chip keeps its whole label; the row scrolls sideways instead of clipping.
                        .fixedSize()
                        .accessibilityAddTraits(chip.filter == content.filter ? .isSelected : [])
                }
            }
            .padding(.top, FSSpace.space1)
            .padding(.bottom, FSSpace.space2)
        }
        .parityRegion(.controls)
    }

    private func mentionBlock(_ mention: MentionContent) -> some View {
        VStack(alignment: .leading, spacing: 0) {
            VStack(spacing: FSSpace.space3) {
                if let role = mention.role { FSButton(role, kind: .primary, action: actions.mentionRole) }
                // The role's mention is the one primary; Mention all is always the outline.
                FSButton(mention.all, kind: .quiet, action: actions.mentionAll)
            }
            .disabled(mentioning)
            .padding(.top, RaidFlowSpacing.buttons)
            if let status { SmallText(status).padding(.top, RaidFlowSpacing.line) }
            Text(mention.preview)
                .type(TypeRole(face: FSFont.body, size: 14, lineHeight: 20, scalesWith: .subheadline))
                .foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
                .padding(.vertical, 10)
                .padding(.horizontal, 12)
                .frame(maxWidth: .infinity, alignment: .leading)
                .overlay(alignment: .leading) { Rectangle().fill(FSColor.line).frame(width: Self.previewRule) }
                .padding(.top, Self.previewTop)
        }
    }

    private var list: some View {
        FSSection(content.listTitle, meta: content.listMeta) {
            if content.rows.isEmpty { SmallText(Copy.raidNaEmpty).padding(.top, RaidFlowSpacing.line) }
            VStack(spacing: 0) {
                ForEach(content.rows) { row in raiderRow(row, actions: content.actionable ? .all : .none) }
            }
        }
        .padding(.top, RaidFlowSpacing.section)
    }

    private func bench(_ title: String) -> some View {
        FSSection(title) {
            VStack(alignment: .leading, spacing: 0) {
                ForEach(content.bench) { row in raiderRow(row, actions: .markIn) }
                if let note = content.benchNote { SmallText(note).padding(.top, RaidFlowSpacing.line) }
            }
        }
        .padding(.top, RaidFlowSpacing.section)
    }

    private enum RowActions { case all, markIn, none }

    private func raiderRow(_ row: NotAnsweredRow, actions rowActions: RowActions) -> some View {
        FSRow(title: row.name, titleColor: ClassColor.of(row.className), sub: row.sub, inset: 0, trailPlacement: .below) {
            ClassCrest(className: row.className, size: FSSize.icon40, flat: true)
        } trail: {
            if let marked = row.marked {
                HStack(spacing: 0) {
                    SmallText(Copy.raidNaMarked(state: marked.label) + " · ")
                    Button(Copy.raidNaUndo) { actions.undo(row) }
                        .font(.custom(FSFont.bodySemibold, size: 13, relativeTo: .footnote))
                        .foregroundStyle(FSColor.gold)
                }
            } else {
                switch rowActions {
                case .all:
                    EqualWidthRow(spacing: FSSpace.space2) {
                        Button(Copy.raidNaMarkIn) { actions.mark(row, .signedIn) }
                        Button(SignupState.bench.label) { actions.mark(row, .bench) }
                        if let dm = row.dm { Button(Copy.raidNaDiscordDM) { actions.message(dm) } }
                    }
                    .buttonStyle(FSButtonStyle(.quiet, fullWidth: true, small: true))
                case .markIn:
                    Button(Copy.raidNaMarkIn) { actions.mark(row, .signedIn) }
                        .buttonStyle(FSButtonStyle(.quiet, fullWidth: false, small: true))
                case .none:
                    EmptyView()
                }
            }
        }
    }
}

/// Buttons in equal columns as wide as the widest label (`.btns` on a row's foot).
struct EqualWidthRow: Layout {
    var spacing: CGFloat

    func sizeThatFits(proposal: ProposedViewSize, subviews: Subviews, cache: inout ()) -> CGSize {
        let (width, height) = cell(subviews)
        let total = width * CGFloat(subviews.count) + spacing * CGFloat(max(0, subviews.count - 1))
        return CGSize(width: min(total, proposal.width ?? total), height: height)
    }

    func placeSubviews(in bounds: CGRect, proposal: ProposedViewSize, subviews: Subviews, cache: inout ()) {
        let (width, height) = cell(subviews)
        let fitted = subviews.isEmpty ? width : min(width, (bounds.width - spacing * CGFloat(subviews.count - 1)) / CGFloat(subviews.count))
        for (index, subview) in subviews.enumerated() {
            let x = bounds.minX + CGFloat(index) * (fitted + spacing)
            subview.place(at: CGPoint(x: x, y: bounds.minY), proposal: ProposedViewSize(width: fitted, height: height))
        }
    }

    private func cell(_ subviews: Subviews) -> (CGFloat, CGFloat) {
        let sizes = subviews.map { $0.sizeThatFits(.unspecified) }
        return (sizes.map(\.width).max() ?? 0, sizes.map(\.height).max() ?? 0)
    }
}
