import SwiftUI

// The round-5 page and its panels (spec 2026-11-09-directions.md Mobile; design/mocks/ws_app_flows.py
// HOME_R5_CSS, RAID_R5_CSS, GEAR_R5_CSS, ACTIVITY_R5_CSS): the tab roots Home, Raid, Gear and Activity
// are built from these, as Android's FlowBandPage, FlowSection(panel =) and FlowRow.

/// A tab root's page in the round-5 look: the painted ground, an optional head on the band art (Home's
/// hero, Activity's chips), then the blocks 16 under it, 14 at the sides and 20 apart. Parity regions
/// are measured from the top of the head, as the mocks' are from the navigation bar.
public struct BandPage<Head: View, Notice: View, Content: View>: View {
    static var side: CGFloat { 14 }
    static var bottom: CGFloat { 22 }

    private let art: BandArt.Art
    private let head: Head
    private let notice: Notice
    private let content: Content

    public init(
        art: BandArt.Art = .band,
        @ViewBuilder head: () -> Head,
        @ViewBuilder notice: () -> Notice = { EmptyView() },
        @ViewBuilder content: () -> Content
    ) {
        self.art = art
        self.head = head()
        self.notice = notice()
        self.content = content()
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            head
                .frame(maxWidth: .infinity, alignment: .leading)
                // The art runs up under the navigation bar and the status bar, the head's own height below them.
                .background(alignment: .bottom) { BandArt(art).padding(.top, -Self.underBars) }
            notice
            VStack(alignment: .leading, spacing: FSSpace.space5) { content }
                .padding(.horizontal, Self.side)
                .padding(.top, FSSpace.space4)
                .padding(.bottom, Self.bottom)
        }
        .parityContent()
    }

    /// How far the art reaches above the head: past any navigation bar, status bar and island.
    static var underBars: CGFloat { 240 }
}

public extension BandPage where Notice == EmptyView {
    init(art: BandArt.Art = .band, @ViewBuilder head: () -> Head, @ViewBuilder content: () -> Content) {
        self.init(art: art, head: head, notice: { EmptyView() }, content: content)
    }
}

/// A round-5 section over the ground: the cream display title at 21/26 with its gold action (or a muted
/// meta beside it), 8 under it, then the content in a panel of that frame. Reports `section-<slug>`.
public struct PanelSection<Content: View>: View {
    private let title: String
    private let meta: String?
    private let action: (title: String, run: () -> Void)?
    private let frame: PanelFrame
    private let content: Content

    public init(
        _ title: String,
        meta: String? = nil,
        action: (title: String, run: () -> Void)? = nil,
        frame: PanelFrame = .none,
        @ViewBuilder content: () -> Content
    ) {
        self.title = title
        self.meta = meta
        self.action = action
        self.frame = frame
        self.content = content()
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: FSSpace.space2) {
            HStack(alignment: .center, spacing: FSSpace.space3) {
                PanelTitle(title)
                if let meta { Text(meta).type(FSType.small).foregroundStyle(FSColor.muted) }
                Spacer(minLength: 0)
                if let action { PanelAction(action.title, action: action.run) }
            }
            PanelBlock(frame: frame) { content }
        }
        .parityRegion(.section(title))
    }
}

/// A section's title on the ground: Marcellus 21/26, cream, with the art's shadow.
public struct PanelTitle: View {
    private let text: String

    public init(_ text: String) {
        self.text = text
    }

    public var body: some View {
        Text(text)
            .type(TypeRole(face: FSFont.title, size: 21, lineHeight: 26, scalesWith: .title3))
            .foregroundStyle(FSColor.cream)
            .shadow(color: .black.opacity(PanelShadow.opacity), radius: PanelShadow.radius, y: PanelShadow.drop)
            .accessibilityAddTraits(.isHeader)
    }
}

/// The display face over the art carries the mock's drop shadow (`0 1px 2px` black at .45).
public enum PanelShadow {
    public static let opacity = 0.45
    public static let radius: CGFloat = 1
    public static let drop: CGFloat = 1
}

/// A section's gold action at the right (`Gear ›`, `All activity (3) ›`, `Mark all read`), padded to the target.
public struct PanelAction: View {
    private let title: String
    private let action: () -> Void

    public init(_ title: String, action: @escaping () -> Void) {
        self.title = title
        self.action = action
    }

    public var body: some View {
        Button(title, action: action)
            .font(.custom(FSFont.bodySemibold, size: 14, relativeTo: .subheadline))
            .foregroundStyle(FSColor.gold)
            .frame(minHeight: FSTarget.min)
            .contentShape(.rect)
            // The mock's line is 24; the target grows past it without moving the title.
            .padding(.vertical, -(FSTarget.min - 24) / 2)
    }
}

/// The round-5 panel a block's rows sit in: the brown-black panel with its frame, 16 at the sides and 8
/// above and below, so a row's own padding finishes the inset.
public struct PanelBlock<Content: View>: View {
    private let frame: PanelFrame
    private let art: PanelArt?
    private let padding: EdgeInsets
    private let content: Content

    public init(
        frame: PanelFrame = .none,
        art: PanelArt? = nil,
        padding: EdgeInsets = EdgeInsets(top: FSSpace.space2, leading: FSSpace.space4, bottom: FSSpace.space2, trailing: FSSpace.space4),
        @ViewBuilder content: () -> Content
    ) {
        self.frame = frame
        self.art = art
        self.padding = padding
        self.content = content()
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: 0) { content }
            .padding(padding)
            .frame(maxWidth: .infinity, alignment: .leading)
            .panel(frame: frame, art: art)
    }
}

/// One row on a panel (the mock's `.row`): an optional 40 lead, the title (15/20 bold), its line (13/18)
/// and an optional third line (Activity's time), the trail under the text; 10 above and below (6 and 44
/// tall when slim); the unread gold dot in the gutter (`unread` is its word to VoiceOver); a `line` rule under it unless it is the last.
public struct PanelRow<Lead: View, Trail: View>: View {
    static var dot: CGFloat { 6 }
    static var dotInset: CGFloat { 10 }

    @Environment(\.panelRowTrailCell) private var trailCell

    private let title: String
    private let titleColor: Color
    private let sub: Text?
    private let meta: String?
    private let slim: Bool
    /// The unread dot's word to VoiceOver; nil for a read row.
    private let unread: String?
    private let last: Bool
    /// iOS's disclosure chevron at the right of a row that pushes a screen.
    private let chevron: Bool
    private let lead: Lead
    private let trail: Trail

    public init(
        _ title: String,
        titleColor: Color = FSColor.strong,
        sub: Text? = nil,
        meta: String? = nil,
        slim: Bool = false,
        unread: String? = nil,
        last: Bool = false,
        chevron: Bool = false,
        @ViewBuilder lead: () -> Lead,
        @ViewBuilder trail: () -> Trail
    ) {
        self.title = title
        self.titleColor = titleColor
        self.sub = sub
        self.meta = meta
        self.slim = slim
        self.unread = unread
        self.last = last
        self.chevron = chevron
        self.lead = lead()
        self.trail = trail()
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: FSSpace.space2) {
            HStack(spacing: FSSpace.space3) {
                lead
                VStack(alignment: .leading, spacing: 0) {
                    Text(title).type(FSType.item).foregroundStyle(titleColor)
                    if let sub { sub.type(RowType.sub).foregroundStyle(FSColor.muted) }
                    if let meta { Text(meta).type(RowType.sub).foregroundStyle(FSColor.muted) }
                }
                .frame(maxWidth: .infinity, alignment: .leading)
                .padding(.trailing, trailCell && !chevron ? FSSpace.space3 : 0)
                if chevron {
                    Image(systemName: "chevron.right")
                        .font(.system(size: 13, weight: .semibold))
                        .foregroundStyle(FSColor.muted)
                        .accessibilityHidden(true)
                }
            }
            .overlay(alignment: .leading) {
                if unread != nil {
                    Circle().fill(FSColor.gold).frame(width: Self.dot, height: Self.dot)
                        .offset(x: -(Self.dotInset + Self.dot / 2))
                        .accessibilityHidden(true)
                }
            }
            trail
        }
        .padding(.vertical, slim ? 6 : FSSpace.rowPadY)
        .frame(minHeight: slim ? FSSize.rowSlim : nil)
        .frame(maxWidth: .infinity, alignment: .leading)
        .overlay(alignment: .bottom) { if !last { Hairline(FSColor.line) } }
        .contentShape(.rect)
        .accessibilityElement(children: .combine)
        .accessibilityValue(unread ?? "")
        .parityItem(.row)
    }
}

public extension PanelRow where Trail == EmptyView {
    init(
        _ title: String, titleColor: Color = FSColor.strong, sub: Text? = nil, meta: String? = nil, slim: Bool = false,
        unread: String? = nil, last: Bool = false, @ViewBuilder lead: () -> Lead
    ) {
        self.init(
            title, titleColor: titleColor, sub: sub, meta: meta, slim: slim, unread: unread, last: last, lead: lead, trail: { EmptyView() }
        )
    }
}

public extension PanelRow where Lead == EmptyView, Trail == EmptyView {
    init(
        _ title: String, titleColor: Color = FSColor.strong, sub: Text? = nil, meta: String? = nil, slim: Bool = false, last: Bool = false,
        chevron: Bool = false
    ) {
        self.init(
            title, titleColor: titleColor, sub: sub, meta: meta, slim: slim, last: last, chevron: chevron,
            lead: { EmptyView() }, trail: { EmptyView() }
        )
    }
}

public extension Text {
    /// A row's line ending in its gold link (`None in your bags · Auction house`).
    static func line(_ sub: String?, link: String, size: CGFloat = 13) -> Text {
        let tail = Text(link).font(.custom(FSFont.bodySemibold, size: size, relativeTo: .footnote)).foregroundColor(FSColor.gold)
        guard let sub, !sub.isEmpty else { return tail }
        return Text("\(sub) · \(tail)")
    }
}

public extension EnvironmentValues {
    /// Rows keep the web Row's empty trail cell (its 12 pt gap after the text): Home's rows, built on
    /// ws_prims `Row`, wrap where the mock's do.
    @Entry var panelRowTrailCell = false
}
