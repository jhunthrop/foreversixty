import SwiftUI
import UIKit

// The flat system's primitives (mobile redesign spec 3.4), named after the web's `ui/*` so a grep
// finds all three. Each reads tokens only and reports its parity region or item.

/// `ui/Button` (round 5, web lib/ui/button.ts): 44 pt, 13/16 700 in sentence case, the 6 pt control
/// radius; primary their teal with the teal edge, secondary their blue, quiet the rule outline (the
/// action beside a primary, as Android's), ghost the gold text, danger the wipe outline. One primary
/// per screen.
public struct FSButtonStyle: ButtonStyle {
    public enum Kind: Sendable { case primary, secondary, quiet, danger, ghost }

    public static let padX: CGFloat = 18
    /// The side padding of a button in a row of equal buttons.
    public static let padXTight: CGFloat = 4
    private static let labelSize: CGFloat = 13
    private static let labelTracking: CGFloat = 0.01

    private let kind: Kind
    private let fullWidth: Bool
    private let small: Bool
    private let padX: CGFloat
    @Environment(\.isEnabled) private var isEnabled

    /// `.btn.sm`: a chip in a row of filters, 36 high inside its 44 pt target.
    public static let smallHeight: CGFloat = 36

    /// `padX`: the label's side padding; a row of equal buttons (the raid's four answers) sets 4 so a
    /// word like `Bench` keeps one line at 393 pt.
    public init(_ kind: Kind = .secondary, fullWidth: Bool = true, small: Bool = false, padX: CGFloat = Self.padX) {
        self.kind = kind
        self.fullWidth = fullWidth
        self.small = small
        self.padX = padX
    }

    public func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .font(.custom(FSFont.bodyBold, size: Self.labelSize, relativeTo: .subheadline))
            .tracking(Self.labelTracking * Self.labelSize)
            .multilineTextAlignment(.center)
            .foregroundStyle(foreground.opacity(isEnabled ? 1 : 0.5))
            .padding(.horizontal, padX)
            .frame(maxWidth: fullWidth ? .infinity : nil, minHeight: small ? Self.smallHeight : FSTarget.min)
            .background(background(pressed: configuration.isPressed), in: .rect(cornerRadius: Theme.Radius.control))
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.control).strokeBorder(border, lineWidth: FSStroke.hairline))
            .contentShape(.rect)
            .parityItem(.button)
    }

    private var foreground: Color {
        switch kind {
        case .primary, .secondary: FSColor.strong
        case .quiet: FSColor.cream
        case .danger: FSColor.wipe
        case .ghost: FSColor.gold
        }
    }

    private func background(pressed: Bool) -> Color {
        switch kind {
        case .primary: pressed ? FSColor.tealEdge : FSColor.teal
        case .secondary: pressed ? FSColor.blueEdge : FSColor.blue
        case .quiet, .danger, .ghost: pressed ? FSColor.pressed : .clear
        }
    }

    private var border: Color {
        switch kind {
        case .primary: FSColor.tealEdge
        case .secondary: FSColor.blueEdge
        // The cream frame line: the rule's slate disappears on the painted ground.
        case .quiet: FSColor.frameLine
        case .danger: FSColor.wipe.opacity(0.4)
        case .ghost: .clear
        }
    }
}

/// `ui/Button` as a view.
public struct FSButton: View {
    private let title: String
    private let kind: FSButtonStyle.Kind
    private let fullWidth: Bool
    private let action: () -> Void

    public init(_ title: String, kind: FSButtonStyle.Kind = .secondary, fullWidth: Bool = true, action: @escaping () -> Void) {
        self.title = title
        self.kind = kind
        self.fullWidth = fullWidth
        self.action = action
    }

    public var body: some View {
        Button(title, action: action).buttonStyle(FSButtonStyle(kind, fullWidth: fullWidth))
    }
}

/// `ui/PageHeader`: the Marcellus title (32/38) once per screen, in the class colour or the cream,
/// an optional lead (a crest) and one wrapping meta line. Reports `head`, `head-title` and
/// `head-meta`.
public struct PageHeader<Lead: View, Meta: View>: View {
    private let title: String
    private let titleColor: Color
    private let alignment: HorizontalAlignment
    private let balanced: Bool
    private let lead: Lead
    private let meta: Meta
    private let hasMeta: Bool

    static var padTop: CGFloat { FSSpace.space4 }
    static var padBottom: CGFloat { 14 }

    public init(
        title: String,
        titleColor: Color = Theme.Palette.title,
        alignment: HorizontalAlignment = .leading,
        balanced: Bool = false,
        @ViewBuilder lead: () -> Lead,
        @ViewBuilder meta: () -> Meta
    ) {
        self.title = title
        self.titleColor = titleColor
        self.alignment = alignment
        self.balanced = balanced
        self.lead = lead()
        self.meta = meta()
        hasMeta = Meta.self != EmptyView.self
    }

    /// The title on one line when it fits; balanced (the web's `text-wrap: balance`) over two lines
    /// when asked and it does not, so a long guild name breaks evenly on every engine.
    @ViewBuilder
    private var titleText: some View {
        if balanced {
            ViewThatFits(in: .horizontal) {
                Text(title).fixedSize()
                Text(BalancedLines.two(title, role: FSType.title)).fixedSize(horizontal: false, vertical: true)
            }
        } else {
            Text(title).fixedSize(horizontal: false, vertical: true)
        }
    }

    public var body: some View {
        HStack(spacing: FSSpace.space4) {
            lead
            // The web's shrink-to-fit block: as wide as the title or the meta line, whichever is
            // wider, and the meta line fills it.
            MatchedWidthStack(spacing: 6) {
                titleText
                    .type(FSType.title)
                    .foregroundStyle(titleColor)
                    .multilineTextAlignment(alignment == .center ? .center : .leading)
                    .frame(maxWidth: .infinity, alignment: Alignment(horizontal: alignment, vertical: .center))
                    .accessibilityLabel(title)
                    .accessibilityAddTraits(.isHeader)
                if hasMeta {
                    meta.frame(maxWidth: .infinity, alignment: Alignment(horizontal: alignment, vertical: .center))
                        .parityRegion(.headMeta)
                }
            }
            .frame(maxWidth: .infinity, alignment: Alignment(horizontal: alignment, vertical: .center))
        }
        .parityRegion(.headTitle)
        .padding(.top, Self.padTop)
        .padding(.bottom, Self.padBottom)
        .frame(maxWidth: .infinity, alignment: .leading)
        .parityRegion(.head)
    }
}

public extension PageHeader where Lead == EmptyView {
    init(title: String, titleColor: Color = Theme.Palette.title, alignment: HorizontalAlignment = .leading, @ViewBuilder meta: () -> Meta) {
        self.init(title: title, titleColor: titleColor, alignment: alignment, lead: { EmptyView() }, meta: meta)
    }
}

public extension PageHeader where Lead == EmptyView, Meta == EmptyView {
    init(title: String, titleColor: Color = Theme.Palette.title, alignment: HorizontalAlignment = .leading, balanced: Bool = false) {
        self.init(
            title: title, titleColor: titleColor, alignment: alignment, balanced: balanced, lead: { EmptyView() }, meta: { EmptyView() }
        )
    }
}

/// The secondary ink of the surface a line sits on (round 5): the panel's `muted` by default; the
/// painted ground (`screenGround()`) provides `mutedGround`, AA there, and a panel restores `muted`.
/// Android: `LocalSubInk`.
private struct SubInkKey: EnvironmentKey {
    static let defaultValue = FSColor.muted
}

public extension EnvironmentValues {
    var subInk: Color {
        get { self[SubInkKey.self] }
        set { self[SubInkKey.self] = newValue }
    }
}

/// Small secondary text (`t-small`, 13/20), in the surface's secondary ink unless given a colour.
public struct SmallText: View {
    private let text: String
    private let color: Color?
    @Environment(\.subInk) private var subInk

    public init(_ text: String, color: Color? = nil) {
        self.text = text
        self.color = color
    }

    public var body: some View {
        Text(text).type(FSType.small).foregroundStyle(color ?? subInk).fixedSize(horizontal: false, vertical: true)
    }
}

/// `ui/Section`: a Marcellus title (18/24, cream) over a `lineSoft` hairline, no box; meta and an
/// action at the right. Reports `section-<slug of the title>`.
public struct FSSection<Content: View, Trailing: View>: View {
    private let title: String
    private let meta: String?
    private let trailing: Trailing
    private let content: Content

    /// `meta`: the small muted words beside the title (`6 open`, `4 of 31, healers first`).
    public init(_ title: String, meta: String? = nil, @ViewBuilder trailing: () -> Trailing, @ViewBuilder content: () -> Content) {
        self.title = title
        self.meta = meta
        self.trailing = trailing()
        self.content = content()
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            HStack(alignment: .firstTextBaseline, spacing: FSSpace.space3) {
                Text(title).type(FSType.section).foregroundStyle(Theme.Palette.title).accessibilityAddTraits(.isHeader)
                if let meta { SmallText(meta) }
                Spacer(minLength: 0)
                trailing
            }
            .padding(.bottom, 10)
            // The hairline takes its own point, as the web's border-bottom does.
            .padding(.bottom, FSStroke.hairline)
            .overlay(alignment: .bottom) { Hairline() }
            content
        }
        .parityRegion(.section(title))
    }
}

public extension FSSection where Trailing == EmptyView {
    init(_ title: String, meta: String? = nil, @ViewBuilder content: () -> Content) {
        self.init(title, meta: meta, trailing: { EmptyView() }, content: content)
    }
}

/// The `lineSoft` hairline under a row or a section head.
public struct Hairline: View {
    private let color: Color

    public init(_ color: Color = FSColor.lineSoft) {
        self.color = color
    }

    public var body: some View {
        Rectangle().fill(color).frame(height: FSStroke.hairline).accessibilityHidden(true)
    }
}

/// `ui/Row`: lead, body (title and sub), trail; 56 min (44 slim), a hairline under. Selected rows
/// take the `selected` ground and a 3 pt gold bar. `inset` pads the sides (the web's `.rows` rows
/// pad 16; a row in an open list pads only when selected).
public struct FSRow<Lead: View, Trail: View>: View {
    public enum Density: Sendable { case regular, slim }
    /// Where the trail sits: at the right, or under the lead and body (a tag on a phone row).
    public enum TrailPlacement: Sendable { case trailing, below }

    private let title: String
    private let titleColor: Color
    private let sub: String?
    private let detail: String?
    private let selected: Bool
    private let density: Density
    private let inset: CGFloat
    private let trailPlacement: TrailPlacement
    private let lead: Lead
    private let trail: Trail
    @Environment(\.subInk) private var subInk

    public init(
        title: String,
        titleColor: Color = FSColor.strong,
        sub: String? = nil,
        detail: String? = nil,
        selected: Bool = false,
        density: Density = .regular,
        inset: CGFloat = FSSpace.rowPadX,
        trailPlacement: TrailPlacement = .trailing,
        @ViewBuilder lead: () -> Lead,
        @ViewBuilder trail: () -> Trail
    ) {
        self.title = title
        self.titleColor = titleColor
        self.sub = sub
        self.detail = detail
        self.selected = selected
        self.density = density
        self.inset = inset
        self.trailPlacement = trailPlacement
        self.lead = lead()
        self.trail = trail()
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: FSSpace.space2) {
            HStack(spacing: FSSpace.space3) {
                lead
                VStack(alignment: .leading, spacing: 0) {
                    Text(title).type(FSType.item).foregroundStyle(titleColor)
                    if let sub {
                        Text(sub).type(RowType.sub).foregroundStyle(subInk)
                    }
                    if let detail {
                        Text(detail).type(RowType.detail).foregroundStyle(subInk).padding(.top, 2)
                    }
                }
                .frame(maxWidth: .infinity, alignment: .leading)
                if trailPlacement == .trailing { trail }
            }
            if trailPlacement == .below { trail }
        }
        .padding(.horizontal, inset)
        .padding(.vertical, density == .slim ? 6 : FSSpace.rowPadY)
        // The web's border-box: the minimum includes the hairline, which takes its own point.
        .frame(minHeight: (density == .slim ? FSSize.rowSlim : FSSize.row) - FSStroke.hairline)
        .padding(.bottom, FSStroke.hairline)
        .background(selected ? FSColor.selected : .clear)
        .overlay(alignment: .leading) {
            if selected { Rectangle().fill(FSColor.gold).frame(width: FSStroke.selection) }
        }
        .overlay(alignment: .bottom) { Hairline() }
        .contentShape(.rect)
        .accessibilityElement(children: .combine)
        .accessibilityAddTraits(selected ? .isSelected : [])
        .parityItem(.row)
    }
}

public extension FSRow where Trail == EmptyView {
    init(
        title: String, titleColor: Color = FSColor.strong, sub: String? = nil, detail: String? = nil, selected: Bool = false,
        density: Density = .regular, inset: CGFloat = FSSpace.rowPadX, @ViewBuilder lead: () -> Lead
    ) {
        self.init(
            title: title, titleColor: titleColor, sub: sub, detail: detail, selected: selected, density: density, inset: inset,
            trailPlacement: .trailing, lead: lead, trail: { EmptyView() }
        )
    }
}

/// `ui/Tag`: 20 high, 3 radius, a 1 pt border in its tone, 12/16 600, sentence case. Not tappable.
public struct FSTag: View {
    public enum Tone: Sendable { case gold, kill, wipe, muted }

    private let text: String
    private let tone: Tone

    public init(_ text: String, tone: Tone = .muted) {
        self.text = text
        self.tone = tone
    }

    public var body: some View {
        Text(text)
            .font(.custom(FSFont.bodySemibold, size: 12, relativeTo: .caption))
            .foregroundStyle(color)
            .padding(.horizontal, 6)
            .frame(minHeight: 20)
            .overlay(RoundedRectangle(cornerRadius: FSRadius.pill).strokeBorder(border, lineWidth: FSStroke.hairline))
            .fixedSize()
    }

    private var color: Color {
        switch tone {
        case .gold: FSColor.gold
        case .kill: FSColor.kill
        case .wipe: FSColor.wipe
        case .muted: FSColor.muted
        }
    }

    private var border: Color {
        switch tone {
        case .gold: FSColor.gold
        case .kill: FSColor.kill.opacity(0.4)
        case .wipe: FSColor.wipe.opacity(0.4)
        case .muted: FSColor.line
        }
    }
}

/// The web's `text-wrap: balance` for two lines: the break between words that makes the longer
/// line shortest, measured in the role's face with its tracking (widths scale together, so the
/// base size decides the break at every text size).
public enum BalancedLines {
    public static func two(_ text: String, role: FSTextRole) -> String {
        let words = text.split(separator: " ").map(String.init)
        guard words.count > 1, let font = UIFont(name: role.face, size: role.size) else { return text }
        let kern = role.tracking * role.size
        func width(_ line: [String]) -> CGFloat {
            NSAttributedString(string: line.joined(separator: " "), attributes: [.font: font, .kern: kern]).size().width
        }
        let best = (1 ..< words.count).min { a, b in
            max(width(Array(words[..<a])), width(Array(words[a...]))) < max(width(Array(words[..<b])), width(Array(words[b...])))
        } ?? 1
        return words[..<best].joined(separator: " ") + "\n" + words[best...].joined(separator: " ")
    }
}
