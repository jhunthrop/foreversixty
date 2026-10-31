import DesignSystem
import DesignTokens
import SwiftUI
import UIKit
import WidgetKit

/// The home-screen widget container (spec §4): the card gradient and a 17% gold border. The
/// system clips it to its own corner radius; previews and snapshots clip at 22.
struct WidgetCard: View {
    var body: some View {
        ZStack {
            LinearGradient(colors: [FSColor.cardTop, FSColor.raised], startPoint: .top, endPoint: .bottom)
            RoundedRectangle(cornerRadius: WidgetMetrics.radius).strokeBorder(FSColor.lineCard, lineWidth: 1)
        }
    }
}

enum WidgetMetrics {
    static let radius: CGFloat = 22
    static let padding: CGFloat = 14
    static let small = CGSize(width: 170, height: 170)
    static let medium = CGSize(width: 364, height: 170)
    static let rectangular = CGSize(width: 172, height: 76)
    static let circular = CGSize(width: 76, height: 76)
    static let inline = CGSize(width: 260, height: 24)
    static let buttonHeight: CGFloat = 44
}

extension View {
    /// Widgets follow the system's text size up to the largest non-accessibility size, as
    /// WidgetKit itself does on the home screen (spec §9: widgets follow the system); past it the
    /// words would cut, and §9 allows a cut only on a widget's item name.
    func widgetTextCap() -> some View {
        dynamicTypeSize(...DynamicTypeSize.xxxLarge)
    }

    /// The widget's background in WidgetKit; ignored outside a widget, where WidgetFrame draws it.
    func widgetCardBackground() -> some View {
        containerBackground(for: .widget) { WidgetCard() }
    }
}

/// A widget drawn outside WidgetKit (snapshot tests, review captures): its system size, the
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

/// The seal and the gold micro label a widget opens with: `Next raid`.
struct WidgetHeader: View {
    let title: String
    var color: Color = FSColor.gold

    var body: some View {
        HStack(spacing: 6) {
            if let seal = BrandImages.seal {
                Image(uiImage: seal).resizable().scaledToFit().frame(width: 16, height: 16).accessibilityHidden(true)
            }
            MicroLabel(title, color: color).lineLimit(1)
        }
    }
}

/// The foot tag in its kind's colour (spec §5.6): kill for In, gold for an ask, wipe for Cancelled.
struct WidgetTag: View {
    let text: String
    let kind: WidgetTagKind

    var body: some View {
        TagView(text: text, color: Self.color(kind))
    }

    static func color(_ kind: WidgetTagKind) -> Color {
        switch kind {
        case .in: FSColor.kill
        case .notSignedUp, .chooseSpec: FSColor.gold
        case .cancelled: FSColor.wipe
        case .late, .bench, .out, .locked: FSColor.muted
        }
    }
}

/// Where widget item icons come from: the files the app leaves in the App Group (the extension
/// cannot fetch while it draws). Snapshots inject the repo's own icons.
struct WidgetIcons: Sendable {
    let image: @Sendable (String) -> UIImage?

    static let shared = WidgetIcons { icon in
        UIImage(contentsOfFile: WidgetSnapshotStore.shared().iconURL(icon).path)
    }
}

/// An item icon in its quality border; the bordered tile draws even without the art (spec §11).
struct WidgetItemIcon: View {
    let icon: String
    let quality: Int
    let size: CGFloat
    let icons: WidgetIcons
    @Environment(\.widgetRenderingMode) private var renderingMode

    var body: some View {
        // Tinted mode has no quality colours: the border joins the accent group so it stays legible.
        let tinted = renderingMode == .accented
        let color = tinted ? Color.white.opacity(0.8) : ItemQuality(value: quality).border
        ZStack {
            RoundedRectangle(cornerRadius: 6).fill(FSColor.cardTop)
            if let image = icons.image(icon) {
                // Tinted home screens: the art stays visible, desaturated, instead of a flat tint.
                Image(uiImage: image).resizable().widgetAccentedRenderingMode(.desaturated).scaledToFill()
            }
        }
        .frame(width: size, height: size)
        .clipShape(.rect(cornerRadius: 6))
        .overlay(RoundedRectangle(cornerRadius: 6).stroke(color, lineWidth: 1.5).widgetAccentable(tinted))
        .accessibilityHidden(true)
    }
}

/// The LX seal alone: the Island's compact and minimal mark.
struct SealMark: View {
    var body: some View {
        if let seal = BrandImages.seal {
            Image(uiImage: seal).resizable().scaledToFit().frame(width: 20, height: 20)
        }
    }
}
