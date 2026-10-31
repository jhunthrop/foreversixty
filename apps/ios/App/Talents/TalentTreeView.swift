import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// The tree's fixed geometry (spec F6 §5.2): 48 pt cells, a 70 pt tier pitch, a column pitch of
/// at most 96 that fits a 360 pt screen, and a region always reserved at seven tiers.
enum TalentGeometry {
    static let cell: CGFloat = 48
    static let icon: CGFloat = 44
    static let border: CGFloat = 2
    static let tierPitch: CGFloat = 70
    static let maxColumnPitch: CGFloat = 96
    static let columns = 4
    static let reservedTiers = 7
    static let margin: CGFloat = 16
    /// The badge hangs 6 below the last tier's cells.
    static let badgeHang: CGFloat = 6
    static let arrowhead: CGFloat = 7
    static let lineWidth: CGFloat = 2
    static let haloWidth: CGFloat = 5

    static var height: CGFloat { CGFloat(reservedTiers - 1) * tierPitch + cell + badgeHang + 2 }

    static func columnPitch(width: CGFloat) -> CGFloat {
        min(maxColumnPitch, ((width - 2 * margin - cell) / CGFloat(columns - 1)).rounded(.down))
    }

    /// The top-left corner of a cell, the grid centred in `width`.
    static func origin(_ cell: TalentCell, width: CGFloat) -> CGPoint {
        let pitch = columnPitch(width: width)
        let left = (width - (Self.cell + CGFloat(columns - 1) * pitch)) / 2
        return CGPoint(x: left + CGFloat(cell.column) * pitch, y: CGFloat(cell.tier) * tierPitch)
    }

    /// A prerequisite line from cell edge to cell edge (`connectors.ts` `connectorPath`), and the
    /// point and direction its arrowhead ends on.
    static func path(_ line: TalentConnector, width: CGFloat) -> (path: Path, tip: CGPoint, direction: CGVector) {
        let half = cell / 2
        let from = origin(line.from, width: width)
        let to = origin(line.to, width: width)
        let start: CGPoint
        let tip: CGPoint
        let direction: CGVector
        var path = Path()
        switch line.shape {
        case .down:
            start = CGPoint(x: from.x + half, y: from.y + cell)
            tip = CGPoint(x: to.x + half, y: to.y)
            direction = CGVector(dx: 0, dy: 1)
            path.move(to: start)
            path.addLine(to: tip)
        case .across:
            let side: CGFloat = line.to.column > line.from.column ? 1 : -1
            start = CGPoint(x: from.x + half + side * half, y: from.y + half)
            tip = CGPoint(x: to.x + half - side * half, y: to.y + half)
            direction = CGVector(dx: side, dy: 0)
            path.move(to: start)
            path.addLine(to: tip)
        case .elbow:
            let side: CGFloat = line.to.column > line.from.column ? 1 : -1
            start = CGPoint(x: from.x + half, y: from.y + cell)
            tip = CGPoint(x: to.x + half - side * half, y: to.y + half)
            direction = CGVector(dx: side, dy: 0)
            path.move(to: start)
            path.addLine(to: CGPoint(x: start.x, y: tip.y))
            path.addLine(to: tip)
        }
        return (path, tip, direction)
    }

    static func arrowhead(tip: CGPoint, direction: CGVector) -> Path {
        let size = arrowhead
        let back = CGPoint(x: tip.x - direction.dx * size, y: tip.y - direction.dy * size)
        let normal = CGVector(dx: -direction.dy, dy: direction.dx)
        var path = Path()
        path.move(to: tip)
        path.addLine(to: CGPoint(x: back.x + normal.dx * size / 2, y: back.y + normal.dy * size / 2))
        path.addLine(to: CGPoint(x: back.x - normal.dx * size / 2, y: back.y - normal.dy * size / 2))
        path.closeSubpath()
        return path
    }
}

/// One tree: the prerequisite lines under the cells (spec F6 §5.2).
struct TalentTreeView: View {
    let cells: [TalentCellModel]
    let connectors: [TalentConnector]
    let build: String
    let selectedID: Int?
    let open: (Int) -> Void

    var body: some View {
        GeometryReader { proxy in
            let width = proxy.size.width
            ZStack(alignment: .topLeading) {
                lines(width: width)
                ForEach(cells) { cell in
                    let origin = TalentGeometry.origin(TalentCell(tier: cell.talent.tier, column: cell.talent.column), width: width)
                    TalentCellView(cell: cell, build: build, isSelected: selectedID == cell.id) { open(cell.id) }
                        .id(cell.id)
                        .offset(x: origin.x, y: origin.y)
                }
            }
        }
        .frame(height: TalentGeometry.height)
    }

    private func lines(width: CGFloat) -> some View {
        Canvas { context, _ in
            for line in connectors {
                let (path, tip, direction) = TalentGeometry.path(line, width: width)
                let color = line.met ? FSColor.gold : FSColor.line
                let head = TalentGeometry.arrowhead(tip: tip, direction: direction)
                context.stroke(path, with: .color(FSColor.bg), style: StrokeStyle(lineWidth: TalentGeometry.haloWidth, lineCap: .round))
                context.stroke(head, with: .color(FSColor.bg), style: StrokeStyle(lineWidth: TalentGeometry.haloWidth - 2, lineJoin: .round))
                context.stroke(path, with: .color(color), style: StrokeStyle(lineWidth: TalentGeometry.lineWidth, lineCap: .round))
                context.fill(head, with: .color(color))
            }
        }
        .accessibilityHidden(true)
    }
}

/// One talent: the icon on its tile in the state border, the rank badge, and on Your build the
/// dashed ring and the signed chip (spec F6 §5.2).
struct TalentCellView: View {
    let cell: TalentCellModel
    let build: String
    let isSelected: Bool
    let open: () -> Void

    var body: some View {
        Button(action: open) {
            TalentIconTile(talent: cell.talent, state: cell.state, build: build)
                .overlay { ring }
                .overlay(alignment: .bottomTrailing) { badge.offset(x: 6, y: 6) }
                .overlay(alignment: .topLeading) { chip.offset(x: -8, y: -8) }
                .overlay {
                    if isSelected {
                        RoundedRectangle(cornerRadius: FSRadius.control + 9)
                            .stroke(FSColor.gold, lineWidth: 2)
                            .padding(-9)
                    }
                }
        }
        .buttonStyle(CellPressStyle())
        .accessibilityLabel(cell.accessibilityLabel)
        .accessibilityHint(Copy.talentsA11yHint)
    }

    private var deltaColor: Color { (cell.delta ?? 0) > 0 ? FSColor.wipe : FSColor.gold }

    @ViewBuilder
    private var ring: some View {
        if cell.delta != nil {
            RoundedRectangle(cornerRadius: 12)
                .stroke(deltaColor, style: StrokeStyle(lineWidth: 2, dash: [4, 3]))
                .padding(-3)
        }
    }

    private var badge: some View {
        TalentBadge(text: cell.badge, color: TalentColors.badgeBorder(cell.state), textColor: TalentColors.badgeText(cell.state))
    }

    @ViewBuilder
    private var chip: some View {
        if let chip = cell.chip {
            TalentBadge(text: chip, color: deltaColor, textColor: deltaColor)
        }
    }
}

/// The small mono capsule used for the rank badge and the signed chip; its text caps at 14.
struct TalentBadge: View {
    let text: String
    let color: Color
    let textColor: Color

    var body: some View {
        Text(text)
            .font(.system(size: 11, weight: .semibold, design: .monospaced))
            .dynamicTypeSize(...DynamicTypeSize.xxxLarge)
            .foregroundStyle(textColor)
            .padding(.horizontal, 5)
            .padding(.vertical, 1)
            .background(FSColor.bg, in: .rect(cornerRadius: 6))
            .overlay(RoundedRectangle(cornerRadius: 6).stroke(color, lineWidth: 1))
            .fixedSize()
            .accessibilityHidden(true)
    }
}

/// The 48 pt tile: the icon at 44 on `#131824` inside a 2 pt state border; a locked face at 40%;
/// a missing icon keeps the tile with the name's first two letters (spec F6 §5.2).
struct TalentIconTile: View {
    let talent: Talent
    let state: TalentCellState
    let build: String
    var size: CGFloat = TalentGeometry.cell

    var body: some View {
        let inner = size - 2 * TalentGeometry.border
        GameIconImage(icon: talent.icon, build: build, fallback: String(talent.name.prefix(2)))
            .id("\(build)/\(talent.icon)")
            .frame(width: inner, height: inner)
            .clipShape(.rect(cornerRadius: FSRadius.pill))
            .opacity(state == .locked ? 0.4 : 1)
            .frame(width: size, height: size)
            .background(FSColor.cardTop, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(TalentColors.border(state), lineWidth: TalentGeometry.border))
            .accessibilityHidden(true)
    }
}

/// The state colours (spec F6 §5.2's table).
enum TalentColors {
    static func border(_ state: TalentCellState) -> Color {
        switch state {
        case .maxed, .filled: FSColor.kill
        case .open, .locked: FSColor.line
        }
    }

    static func badgeBorder(_ state: TalentCellState) -> Color { border(state) }

    static func badgeText(_ state: TalentCellState) -> Color {
        switch state {
        case .maxed: FSColor.kill
        case .filled: FSColor.text
        case .open, .locked: FSColor.muted
        }
    }
}

/// §8: a pressed cell scales to 0.96 under the gold wash.
private struct CellPressStyle: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .overlay {
                if configuration.isPressed {
                    RoundedRectangle(cornerRadius: FSRadius.control).fill(GameColors.pressedWash)
                }
            }
            .scaleEffect(configuration.isPressed ? 0.96 : 1)
    }
}
