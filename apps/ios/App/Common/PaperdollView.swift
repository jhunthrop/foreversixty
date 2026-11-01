import DesignSystem
import SwiftUI

/// The one Paperdoll (tenet 10; design/specs/2026-10-31-paperdoll.md phone frame): two mirrored
/// columns with the sockets on the outer edges and the names toward the centre in quality
/// colour, an `ilvl` micro line, the gold dot on a slot whose BiS list has better, and the
/// weapons in a row under the columns. At the accessibility sizes it is one column of 17 rows.
/// The character page (F11.5) and Home's Your gear (F2.7) both draw it.
struct PaperdollView: View {
    let frame: PaperdollFrame
    let build: String
    /// A tap on a socket holding an item: the F2.6 tooltip sheet.
    var open: (PaperdollSlot) -> Void = { _ in }

    @Environment(\.dynamicTypeSize) private var typeSize
    @State private var width: CGFloat = PaperdollLayout.phoneWidth

    private var socket: CGFloat { width < PaperdollLayout.narrowContainer ? PaperdollLayout.narrowSocket : PaperdollLayout.socket }

    var body: some View {
        Group {
            if typeSize.isAccessibilitySize {
                VStack(alignment: .leading, spacing: PaperdollLayout.rowGap) {
                    ForEach(frame.all) { cell($0, side: .leading) }
                }
            } else {
                columns
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .onGeometryChange(for: CGFloat.self) { $0.size.width } action: { width = $0 }
        .accessibilityElement(children: .contain)
    }

    private var columns: some View {
        VStack(spacing: PaperdollLayout.weaponGap) {
            HStack(alignment: .top, spacing: PaperdollLayout.columnGap) {
                VStack(alignment: .leading, spacing: PaperdollLayout.rowGap) {
                    ForEach(frame.left) { cell($0, side: .leading) }
                }
                .frame(maxWidth: .infinity, alignment: .leading)
                VStack(alignment: .trailing, spacing: PaperdollLayout.rowGap) {
                    ForEach(frame.right) { cell($0, side: .trailing) }
                }
                .frame(maxWidth: .infinity, alignment: .trailing)
            }
            HStack(alignment: .top, spacing: PaperdollLayout.rowGap) {
                ForEach(frame.weapons) { cell($0, side: .center).frame(maxWidth: .infinity, alignment: .top) }
            }
            .padding(.top, PaperdollLayout.weaponGap)
            .overlay(alignment: .top) { FSColor.lineSoft.frame(height: 1) }
        }
    }

    @ViewBuilder
    private func cell(_ slot: PaperdollSlot, side: HorizontalAlignment) -> some View {
        let content = PaperdollCell(slot: slot, build: build, socket: socket, side: side)
        if slot.look != nil {
            Button { open(slot) } label: { content }
                .buttonStyle(RowPressStyle())
                .accessibilityHint(Copy.itemsA11yHint)
        } else {
            content
        }
    }
}

/// One socket and its words: the socket on the outer edge (right in the right column).
struct PaperdollCell: View {
    let slot: PaperdollSlot
    let build: String
    let socket: CGFloat
    /// Leading and trailing: the mirrored columns; centre: the weapons row, socket above the name.
    let side: HorizontalAlignment

    private var trailing: Bool { side == .trailing }

    var body: some View {
        Group {
            if side == .center {
                VStack(spacing: PaperdollLayout.rowGap) { icon; words }
            } else {
                HStack(alignment: .center, spacing: PaperdollLayout.socketGap) {
                    if trailing { words; icon } else { icon; words }
                }
            }
        }
        .frame(minHeight: PaperdollLayout.cellHeight)
        .contentShape(Rectangle())
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(PaperdollModel.a11y(slot))
        .accessibilityAddTraits(slot.look != nil ? .isButton : [])
    }

    private var icon: some View {
        let look = slot.look
        return ItemIconView(icon: look?.icon, quality: look?.quality, build: build, size: socket, dashed: look == nil)
            .opacity(look == nil ? PaperdollLayout.emptyOpacity : 1)
            .overlay(alignment: .topTrailing) {
                if slot.better {
                    Circle().fill(FSColor.gold)
                        .frame(width: PaperdollLayout.dot, height: PaperdollLayout.dot)
                        .overlay(Circle().stroke(FSColor.raised, lineWidth: 2))
                        .offset(x: PaperdollLayout.dotOffset, y: -PaperdollLayout.dotOffset)
                }
            }
    }

    private var words: some View {
        VStack(alignment: side, spacing: 1) {
            // Never clamped: a name or the two-hand words wrap in full (paperdoll spec §3.3).
            Text(name).font(ScaledType.socketName).foregroundStyle(nameColor)
                .fixedSize(horizontal: false, vertical: true)
            // An unknown item's name already carries its slot (`Wrist · Item 288887`).
            if PaperdollModel.showsSlotLine(slot) {
                Text(micro).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
        }
        .multilineTextAlignment(side == .center ? .center : trailing ? .trailing : .leading)
        .frame(maxWidth: .infinity, alignment: Alignment(horizontal: side, vertical: .center))
    }

    private var name: String { PaperdollModel.name(slot) }

    private var nameColor: Color {
        if let look = slot.look { return ItemQuality(value: look.quality).text }
        return FSColor.muted
    }

    /// `Head · ilvl 61`, or the slot word alone.
    private var micro: String {
        guard let level = slot.look?.itemLevel else { return slot.label }
        return "\(slot.label) · \(Copy.itemsIlvl(n: level))"
    }
}

enum PaperdollLayout {
    static let socket: CGFloat = 40
    static let narrowSocket: CGFloat = 36
    /// Under this container width the sockets are 36 (a 375/360 pt phone).
    static let narrowContainer: CGFloat = 340
    static let phoneWidth: CGFloat = 354
    static let cellHeight: CGFloat = 50
    static let socketGap: CGFloat = 8
    static let columnGap: CGFloat = 8
    static let rowGap: CGFloat = 4
    static let weaponGap: CGFloat = 8
    static let dot: CGFloat = 10
    static let dotOffset: CGFloat = 3
    static let emptyOpacity = 0.6
}
