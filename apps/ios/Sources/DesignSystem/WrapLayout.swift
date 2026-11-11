import SwiftUI

/// Lays its children left to right and wraps onto the next line, so a chip is never cut
/// (spec F3.1: the chips wrap).
public struct WrapLayout: Layout {
    public var spacing: CGFloat
    /// The gap between wrapped lines; the item gap when not given.
    public var lineSpacing: CGFloat
    /// True centres each item on its line's middle (a line of words where one carries an icon), else
    /// items sit on the line's top.
    public var centresLines: Bool

    public init(spacing: CGFloat = 8, lineSpacing: CGFloat? = nil, centresLines: Bool = false) {
        self.spacing = spacing
        self.lineSpacing = lineSpacing ?? spacing
        self.centresLines = centresLines
    }

    public func sizeThatFits(proposal: ProposedViewSize, subviews: Subviews, cache _: inout ()) -> CGSize {
        let rows = arrange(width: proposal.width ?? .infinity, subviews: subviews)
        let width = rows.map(\.width).max() ?? 0
        let height = rows.map(\.height).reduce(0, +) + lineSpacing * CGFloat(max(rows.count - 1, 0))
        return CGSize(width: proposal.width ?? width, height: height)
    }

    public func placeSubviews(in bounds: CGRect, proposal _: ProposedViewSize, subviews: Subviews, cache _: inout ()) {
        var y = bounds.minY
        for row in arrange(width: bounds.width, subviews: subviews) {
            var x = bounds.minX
            for index in row.indices {
                let size = subviews[index].sizeThatFits(.init(width: bounds.width, height: nil))
                let top = centresLines ? y + (row.height - size.height) / 2 : y
                subviews[index].place(at: CGPoint(x: x, y: top), proposal: .init(width: size.width, height: size.height))
                x += size.width + spacing
            }
            y += row.height + lineSpacing
        }
    }

    private struct Row {
        var indices: [Int] = []
        var width: CGFloat = 0
        var height: CGFloat = 0
    }

    private func arrange(width: CGFloat, subviews: Subviews) -> [Row] {
        var rows: [Row] = []
        var current = Row()
        for index in subviews.indices {
            let size = subviews[index].sizeThatFits(.init(width: width, height: nil))
            let needed = current.indices.isEmpty ? size.width : current.width + spacing + size.width
            if needed > width, !current.indices.isEmpty {
                rows.append(current)
                current = Row()
            }
            current.width = current.indices.isEmpty ? min(size.width, width) : current.width + spacing + size.width
            current.height = max(current.height, size.height)
            current.indices.append(index)
        }
        if !current.indices.isEmpty { rows.append(current) }
        return rows
    }
}

/// A column as wide as its widest child (within the offer), every child laid at that width: the
/// web's shrink-to-fit block whose children fill it (the page head's title and meta line).
public struct MatchedWidthStack: Layout {
    public var spacing: CGFloat

    public init(spacing: CGFloat = 0) {
        self.spacing = spacing
    }

    public func sizeThatFits(proposal: ProposedViewSize, subviews: Subviews, cache _: inout ()) -> CGSize {
        let width = columnWidth(proposal.width, subviews)
        let heights = subviews.map { $0.sizeThatFits(.init(width: width, height: nil)).height }
        return CGSize(width: width, height: heights.reduce(0, +) + spacing * CGFloat(max(subviews.count - 1, 0)))
    }

    public func placeSubviews(in bounds: CGRect, proposal _: ProposedViewSize, subviews: Subviews, cache _: inout ()) {
        var y = bounds.minY
        for subview in subviews {
            let height = subview.sizeThatFits(.init(width: bounds.width, height: nil)).height
            subview.place(at: CGPoint(x: bounds.minX, y: y), proposal: .init(width: bounds.width, height: height))
            y += height + spacing
        }
    }

    /// The widest child's natural width: a maxWidth-infinity frame reports its content's ideal
    /// width when offered nil, so each child is measured at its ideal and capped at the offer.
    private func columnWidth(_ offer: CGFloat?, _ subviews: Subviews) -> CGFloat {
        let natural = subviews.map { $0.sizeThatFits(.unspecified).width }.max() ?? 0
        guard let offer else { return natural }
        return min(natural, offer)
    }
}
