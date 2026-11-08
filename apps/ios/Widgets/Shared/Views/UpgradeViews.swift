import DesignSystem
import DesignTokens
import SwiftUI
import WidgetKit

/// The upgrade side of the widgets (mocks ios-widgets-*): `Upgrades`, the next upgrade's gain in
/// gold and `next upgrade`; with room, the item in its icon and quality colour. Never a total
/// (flows appendix C: upgrades overlap, the sum is not a gain).
struct UpgradeGlanceColumn: View {
    let glance: UpgradeGlance
    let icons: WidgetIcons
    /// The medium widget names the item; the small one has room for the gain alone.
    var showsItem = false

    var body: some View {
        VStack(alignment: .leading, spacing: WidgetMetrics.gap) {
            Text(Copy.widgetUpgradesLabel).type(WidgetType.label).foregroundStyle(FSColor.muted).lineLimit(1)
            if let gain = glance.gain {
                Text(gain).type(WidgetType.big).foregroundStyle(FSColor.gold).lineLimit(1).minimumScaleFactor(0.6)
                Text(Copy.widgetNextUpgradeLine).type(WidgetType.sub).foregroundStyle(FSColor.muted).lineLimit(1)
                if showsItem, let item = glance.item {
                    WidgetItemLine(item: item, icons: icons)
                    if let source = glance.source {
                        Text(source).type(WidgetType.sub).foregroundStyle(FSColor.muted).lineLimit(1).minimumScaleFactor(0.8)
                    }
                }
            } else {
                Text(glance.sentence ?? "").type(WidgetType.sub).foregroundStyle(FSColor.text).lineLimit(4)
            }
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(glance.spoken)
    }
}

/// Next upgrade, small.
struct UpgradeSmallView: View {
    let glance: UpgradeGlance
    let icons: WidgetIcons

    var body: some View {
        UpgradeGlanceColumn(glance: glance, icons: icons).widgetTextCap()
    }
}

/// Raid and upgrade, medium (the mocks' Medium): the raid on the left, a hairline, the next
/// upgrade on the right.
struct GlanceMediumView: View {
    let raid: RaidGlance
    let upgrade: UpgradeGlance
    let icons: WidgetIcons

    var body: some View {
        HStack(alignment: .top, spacing: 16) {
            RaidGlanceColumn(glance: raid)
            WidgetDivider()
            UpgradeGlanceColumn(glance: upgrade, icons: icons, showsItem: true)
        }
        .widgetTextCap()
    }
}

/// `+14.5 DPS`, `next upgrade` and the item's name; monochrome.
struct UpgradeRectangularView: View {
    let glance: UpgradeGlance

    var body: some View {
        VStack(alignment: .leading, spacing: 1) {
            Text(glance.gain ?? Copy.widgetNextUpgrade).font(.headline.monospacedDigit()).lineLimit(1).widgetAccentable()
            if let item = glance.item {
                Text(Copy.widgetNextUpgradeLine).font(.caption).lineLimit(1)
                Text(item.name).font(.caption).lineLimit(1)
            } else {
                Text(glance.sentence ?? "").font(.caption).lineLimit(2)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(glance.spoken)
    }
}
