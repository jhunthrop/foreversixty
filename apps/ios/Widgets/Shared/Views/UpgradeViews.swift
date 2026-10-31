import DesignSystem
import DesignTokens
import SwiftUI
import WidgetKit

/// Next upgrade, small (spec §5.6): the crest, the item icon in its quality border, the name in
/// quality colour and the gain with its unit.
struct UpgradeSmallView: View {
    let figures: UpgradeFigures
    let className: String?
    let icons: WidgetIcons

    var body: some View {
        VStack(alignment: .leading, spacing: 6) {
            UpgradeHeader(title: Copy.widgetNextUpgrade, className: className)
            if let row = figures.rows.first {
                WidgetItemIcon(icon: row.icon, quality: row.quality, size: 44, icons: icons)
                Spacer(minLength: 0)
                Text(row.name)
                    .font(ScaledType.bodyStrong)
                    .foregroundStyle(ItemQuality(value: row.quality).text)
                    .lineLimit(2)
                    .minimumScaleFactor(0.8)
                Text(row.gain).font(ScaledType.mono).foregroundStyle(FSColor.strong)
            } else {
                UpgradeSentence(figures: figures)
            }
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(UpgradeSpoken.label(figures))
        .widgetTextCap()
    }
}

/// Next upgrade, medium: `Upgrades · {name}`, three rows, and the footer with the whole total.
struct UpgradeMediumView: View {
    let figures: UpgradeFigures
    let className: String?
    let icons: WidgetIcons

    var body: some View {
        VStack(alignment: .leading, spacing: 5) {
            UpgradeHeader(title: figures.header ?? Copy.widgetNextUpgrade, className: className)
            if figures.rows.isEmpty {
                UpgradeSentence(figures: figures)
                Spacer(minLength: 0)
            } else {
                ForEach(figures.rows, id: \.slot) { row in
                    HStack(spacing: 8) {
                        WidgetItemIcon(icon: row.icon, quality: row.quality, size: 28, icons: icons)
                        Text(row.name)
                            .font(ScaledType.bodyStrong)
                            .foregroundStyle(ItemQuality(value: row.quality).text)
                            .lineLimit(1)
                        Spacer(minLength: 4)
                        Text(row.gain).font(ScaledType.mono).foregroundStyle(FSColor.strong).lineLimit(1).fixedSize()
                    }
                    .accessibilityElement(children: .ignore)
                    .accessibilityLabel(Copy.widgetA11yUpgrade(slot: row.slot, item: row.name, gain: row.gain))
                }
                Spacer(minLength: 0)
                if let footer = figures.footer {
                    Text(footer).font(ScaledType.meta).foregroundStyle(FSColor.muted).lineLimit(1).minimumScaleFactor(0.8)
                }
            }
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .widgetTextCap()
    }
}

/// `Next upgrade · Waist`, the item name and the gain; monochrome.
struct UpgradeRectangularView: View {
    let figures: UpgradeFigures

    var body: some View {
        VStack(alignment: .leading, spacing: 1) {
            Text(figures.rectTitle ?? Copy.widgetNextUpgrade).font(.headline).lineLimit(1).minimumScaleFactor(0.8).widgetAccentable()
            if let row = figures.rows.first {
                Text(row.name).font(.caption).lineLimit(1)
                Text(row.gain).font(.caption.monospacedDigit())
            } else {
                Text(figures.sentence ?? "").font(.caption).lineLimit(2)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(UpgradeSpoken.label(figures))
    }
}

struct UpgradeHeader: View {
    let title: String
    let className: String?

    var body: some View {
        HStack(spacing: 6) {
            ClassCrest(className: className, size: 16).accessibilityHidden(true)
            MicroLabel(title, color: FSColor.gold).lineLimit(1)
        }
    }
}

struct UpgradeSentence: View {
    let figures: UpgradeFigures

    var body: some View {
        Text(figures.sentence ?? "")
            .font(ScaledType.body)
            .foregroundStyle(figures.opensApp ? FSColor.gold : FSColor.text)
            .lineLimit(3)
    }
}

enum UpgradeSpoken {
    static func label(_ figures: UpgradeFigures) -> String {
        guard let row = figures.rows.first else { return [Copy.widgetNextUpgrade, figures.sentence].compactMap { $0 }.joined(separator: ", ") }
        return Copy.widgetA11yUpgrade(slot: row.slot, item: row.name, gain: row.gain)
    }
}
