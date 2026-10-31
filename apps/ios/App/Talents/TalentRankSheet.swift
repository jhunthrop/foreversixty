import DesignSystem
import DesignTokens
import SwiftUI

/// The rank sheet (spec F6 §7): the game's tooltip content in a system sheet, so it never clips.
/// At default text it opens at its content's height (220 pt to half the screen), then large; from
/// AX1 a single full-height detent, scrolling inside (spec F6 §7, point 3).
struct TalentRankSheet: View {
    let model: TalentSheetModel
    let build: String

    @Environment(\.dynamicTypeSize) private var typeSize
    @State private var contentHeight: CGFloat = Self.minHeight

    private static let minHeight: CGFloat = 220
    private static let maxHeight: CGFloat = 437
    /// The grabber and the sheet's own top inset above the content.
    private static let chrome: CGFloat = 24

    private var detents: Set<PresentationDetent> {
        guard typeSize < .accessibility1 else { return [.large] }
        let height = min(Self.maxHeight, max(Self.minHeight, contentHeight + Self.chrome))
        return [.height(height), .large]
    }

    var body: some View {
        TalentRankContent(model: model, build: build) { contentHeight = $0 }
            .presentationDetents(detents)
            .presentationDragIndicator(.visible)
            .presentationBackground(FSColor.raised)
    }
}

/// The sheet's body, pure so the snapshot tests draw it directly.
struct TalentRankContent: View {
    let model: TalentSheetModel
    let build: String
    var onHeight: (CGFloat) -> Void = { _ in }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                header
                panel
                if let guideLine = model.guideLine {
                    Text(guideLine)
                        .font(ScaledType.bodyStrong)
                        .foregroundStyle(FSColor.strong)
                        .frame(maxWidth: .infinity, alignment: .leading)
                        .padding(EdgeInsets(top: 10, leading: 12, bottom: 10, trailing: 12))
                        .overlay(
                            RoundedRectangle(cornerRadius: FSRadius.control)
                                .stroke(model.isOver ? FSColor.wipe : FSColor.gold, style: StrokeStyle(lineWidth: 1.5, dash: [4, 3]))
                        )
                }
            }
            .padding(.horizontal, 20)
            .padding(.top, 24)
            .padding(.bottom, 16)
            .onGeometryChange(for: CGFloat.self) { $0.size.height } action: { onHeight($0) }
        }
        .background(FSColor.raised)
    }

    private var header: some View {
        HStack(alignment: .top, spacing: 12) {
            TalentIconTile(talent: model.talent, state: model.state == .locked ? .open : model.state, build: build, size: 56)
            VStack(alignment: .leading, spacing: 2) {
                Text(model.talent.name)
                    .font(ScaledType.name)
                    .foregroundStyle(FSColor.strong)
                    .fixedSize(horizontal: false, vertical: true)
                    .accessibilityAddTraits(.isHeader)
                Text(model.rankLine).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted)
            }
        }
    }

    private var panel: some View {
        VStack(alignment: .leading, spacing: 12) {
            if let current = model.current {
                Text(current).font(ScaledType.body).foregroundStyle(FSColor.text)
                    .fixedSize(horizontal: false, vertical: true)
            }
            if let next = model.next {
                VStack(alignment: .leading, spacing: 4) {
                    MicroLabel(Copy.talentsNextRank)
                    Text(next).font(ScaledType.body).foregroundStyle(FSColor.muted)
                        .fixedSize(horizontal: false, vertical: true)
                }
            }
            ForEach(model.requirements, id: \.self) { line in
                Text(line).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                    .fixedSize(horizontal: false, vertical: true)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(EdgeInsets(top: 12, leading: 14, bottom: 12, trailing: 14))
        .background(GameColors.tooltipPanel, in: .rect(cornerRadius: FSRadius.control))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(GameColors.iconLine, lineWidth: 1))
    }
}
