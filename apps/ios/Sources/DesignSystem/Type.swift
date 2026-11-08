import SwiftUI
import UIKit

/// A type role as a view modifier (mobile redesign spec 3.2): the face at its size scaling with its
/// Dynamic Type style, the line box through `lineSpacing` and half-leading padding (the line minus
/// the face's natural line), its tracking, and tabular figures where the role has them.
public struct TypeRole: ViewModifier {
    private let face: String
    private let size: CGFloat
    private let lineHeight: CGFloat
    private let tracking: CGFloat
    private let tabular: Bool
    private let style: Font.TextStyle
    @ScaledMetric private var scale: CGFloat = 1

    public init(_ role: FSTextRole) {
        self.init(
            face: role.face, size: role.size, lineHeight: role.lineHeight, tracking: role.tracking,
            scalesWith: role.scalesWith, tabular: role.tabular
        )
    }

    /// A line box the tokens have no role for (a row's sub line, 13/18).
    public init(face: String, size: CGFloat, lineHeight: CGFloat, tracking: CGFloat = 0, scalesWith: Font.TextStyle, tabular: Bool = false) {
        self.face = face
        self.size = size
        self.lineHeight = lineHeight
        self.tracking = tracking
        self.tabular = tabular
        style = scalesWith
        _scale = ScaledMetric(wrappedValue: 1, relativeTo: scalesWith)
    }

    public func body(content: Content) -> some View {
        let font = Font.custom(face, size: size, relativeTo: style)
        let leading = max(0, (lineHeight - naturalLine) * scale)
        content
            .font(tabular ? font.monospacedDigit() : font)
            .tracking(tracking * size)
            .lineSpacing(leading)
            .padding(.vertical, leading / 2)
    }

    /// The face's natural line height at the size.
    private var naturalLine: CGFloat {
        UIFont(name: face, size: size)?.lineHeight ?? size * 1.2
    }
}

/// Line boxes the rows use that the tokens carry no role for (`.row .sub`, `.act-row .when`).
public enum RowType {
    /// `.row .sub`: 13/18, muted.
    public static var sub: TypeRole { TypeRole(face: FSFont.body, size: 13, lineHeight: 18, scalesWith: .footnote) }
    /// `.act-row .when`: 12/16.
    public static var detail: TypeRole { TypeRole(face: FSFont.body, size: 12, lineHeight: 16, scalesWith: .caption) }
}

public extension View {
    /// Sets the text in a type role (`FSType`).
    func type(_ role: FSTextRole) -> some View { modifier(TypeRole(role)) }

    /// Sets the text in a line box of its own.
    func type(_ role: TypeRole) -> some View { modifier(role) }
}
