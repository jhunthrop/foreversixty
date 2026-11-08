import DesignSystem
import DesignTokens
import SwiftUI
import UIKit
import WidgetKit

/// The home-screen widget container (mobile redesign R13, the flat system): the `raised` surface
/// with a `line` hairline. The system clips it to its own corner radius; previews and snapshots
/// clip at 22 (`radius.widget`).
struct WidgetCard: View {
    var body: some View {
        ZStack {
            FSColor.raised
            RoundedRectangle(cornerRadius: WidgetMetrics.radius).strokeBorder(FSColor.line, lineWidth: FSStroke.hairline)
        }
    }
}

enum WidgetMetrics {
    static let radius: CGFloat = 22
    static let padding: CGFloat = 14
    static let gap: CGFloat = 6
    static let small = CGSize(width: 170, height: 170)
    static let medium = CGSize(width: 364, height: 170)
    static let rectangular = CGSize(width: 172, height: 76)
    static let circular = CGSize(width: 76, height: 76)
    static let inline = CGSize(width: 260, height: 24)
}

/// The widgets' line boxes (the mocks' `.wd .k`, `.big`, `.nm`, `.s`): the token faces at the
/// sizes a 170 pt widget holds.
enum WidgetType {
    static var label: TypeRole { TypeRole(face: FSFont.bodyBold, size: 12, lineHeight: 16, scalesWith: .caption) }
    static var big: TypeRole { TypeRole(face: FSFont.bodyBold, size: 30, lineHeight: 34, scalesWith: .title2, tabular: true) }
    static var name: TypeRole { TypeRole(face: FSFont.bodyBold, size: 15, lineHeight: 20, scalesWith: .body) }
    static var sub: TypeRole { RowType.sub }
    static var item: TypeRole { TypeRole(face: FSFont.bodyBold, size: 13, lineHeight: 18, scalesWith: .footnote) }
}

extension View {
    /// Widgets follow the system's text size up to the largest non-accessibility size, as
    /// WidgetKit itself does on the home screen; past it the words would cut.
    func widgetTextCap() -> some View {
        dynamicTypeSize(...DynamicTypeSize.xxxLarge)
    }

    /// The widget's background in WidgetKit; ignored outside a widget, where WidgetFrame draws it.
    func widgetCardBackground() -> some View {
        containerBackground(for: .widget) { WidgetCard() }
    }
}

/// A widget drawn outside WidgetKit (snapshot tests, the parity gallery): its system size, the
/// card behind, clipped like the home screen draws it.
struct WidgetFrame<Content: View>: View {
    let size: CGSize
    var card = true
    @ViewBuilder let content: () -> Content

    var body: some View {
        content()
            .padding(card ? WidgetMetrics.padding : 0)
            .frame(width: size.width, height: size.height, alignment: .topLeading)
            .background { if card { WidgetCard() } }
            .clipShape(.rect(cornerRadius: card ? WidgetMetrics.radius : 0))
    }
}

/// Where widget icons come from: the files the app leaves in the App Group (the extension cannot
/// fetch while it draws). Snapshots inject the repo's own icons.
struct WidgetIcons: Sendable {
    let image: @Sendable (String) -> UIImage?

    static let shared = WidgetIcons { icon in
        UIImage(contentsOfFile: WidgetSnapshotStore.shared().iconURL(icon).path)
    }
}

/// A game icon in its quality border (an item) or the plain line (an instance); the bordered tile
/// draws even without the art.
struct WidgetItemIcon: View {
    let icon: String
    var quality: Int?
    let size: CGFloat
    let icons: WidgetIcons
    @Environment(\.widgetRenderingMode) private var renderingMode

    var body: some View {
        // Tinted mode has no quality colours: the border joins the accent group so it stays legible.
        let tinted = renderingMode == .accented
        let color = tinted ? Color.white.opacity(0.8) : (quality.map { ItemQuality(value: $0).border } ?? FSColor.line)
        ZStack {
            RoundedRectangle(cornerRadius: 4).fill(FSColor.hover)
            if let image = icons.image(icon) {
                // Tinted home screens: the art stays visible, desaturated, instead of a flat tint.
                Image(uiImage: image).resizable().widgetAccentedRenderingMode(.desaturated).scaledToFill()
            }
        }
        .frame(width: size, height: size)
        .clipShape(.rect(cornerRadius: 4))
        .overlay(RoundedRectangle(cornerRadius: 4).stroke(color, lineWidth: 1).widgetAccentable(tinted))
        .accessibilityHidden(true)
        .parityItem(.image)
    }
}

/// The `lineSoft` hairline between a medium widget's two halves.
struct WidgetDivider: View {
    var body: some View {
        Rectangle().fill(FSColor.lineSoft).frame(width: FSStroke.hairline).frame(maxHeight: .infinity)
    }
}

/// The item a widget names: its icon in the quality border and its name in the quality colour.
struct WidgetItemLine: View {
    let item: UpgradeGlance.Item
    let icons: WidgetIcons

    var body: some View {
        HStack(spacing: WidgetMetrics.gap) {
            WidgetItemIcon(icon: item.icon, quality: item.quality, size: 20, icons: icons)
            Text(item.name)
                .type(WidgetType.item)
                .foregroundStyle(ItemQuality(value: item.quality).text)
                .lineLimit(1)
                .minimumScaleFactor(0.8)
        }
    }
}
