import ActivityKit
import DesignSystem
import DesignTokens
import Routing
import SwiftUI
import WidgetKit

/// The extension (flows spec step 7): the Next raid, Next upgrade and Raid and upgrade widgets on
/// the home screen and lock screen, the raid and server-sim Live Activities, and the Control.
@main
struct ForeverSixtyWidgets: WidgetBundle {
    var body: some SwiftUI.Widget {
        NextRaidWidget()
        NextRaidCountdownWidget()
        NextUpgradeWidget()
        RaidAndUpgradeWidget()
        RaidLiveActivity()
        SimLiveActivity()
        RaidControl()
    }
}

/// One timeline entry: the snapshot as of `date`, read through the glances at that moment.
struct SnapshotEntry: TimelineEntry {
    let date: Date
    let snapshot: WidgetSnapshot?
    let characterKey: String?

    var raid: RaidGlance { RaidGlance(snapshot, now: date, zone: .current) }
    var upgrade: UpgradeGlance { UpgradeGlance(snapshot) }
    var figures: WidgetFigures { WidgetFigures(snapshot, now: date, zone: .current) }
}

/// Reads the App Group snapshot for the configured character; when a character other than the
/// current one has no file yet, asks the API once (allowed while a timeline is built).
struct SnapshotProvider: AppIntentTimelineProvider {
    func placeholder(in context: Context) -> SnapshotEntry {
        SnapshotEntry(date: .now, snapshot: WidgetSnapshotStore.shared().load(), characterKey: nil)
    }

    func snapshot(for configuration: CharacterConfigurationIntent, in context: Context) async -> SnapshotEntry {
        SnapshotEntry(date: .now, snapshot: await read(configuration.characterKey), characterKey: configuration.characterKey)
    }

    func timeline(for configuration: CharacterConfigurationIntent, in context: Context) async -> Timeline<SnapshotEntry> {
        let key = configuration.characterKey
        let snapshot = await read(key)
        let entries = WidgetTimeline.dates(now: .now, snapshot: snapshot).map {
            SnapshotEntry(date: $0, snapshot: snapshot, characterKey: key)
        }
        return Timeline(entries: entries, policy: .atEnd)
    }

    private func read(_ key: String?) async -> WidgetSnapshot? {
        let store = WidgetSnapshotStore.shared()
        if let found = store.snapshot(for: key) { return found }
        guard let key, store.load() != nil, let calls = LiveWidgetCalls.fromBundle(),
              let widget = try? await calls.snapshot(character: key) else { return store.snapshot(for: key) }
        let fresh = WidgetSnapshot(widget: widget, fetchedAt: .now)
        try? store.save(fresh, character: key)
        return fresh
    }
}

/// The deep link a raid widget opens: the raid page, else the app.
private func raidURL(_ snapshot: WidgetSnapshot?) -> URL? {
    guard let raid = snapshot?.widget.nextRaid else { return nil }
    return DeepLink.raid(guild: raid.guildId, raid: raid.raidId).url
}

/// The deep link an upgrade widget opens: the character's page (no item route exists in
/// apps/shared/links.json yet).
private func characterURL(_ snapshot: WidgetSnapshot?) -> URL? {
    guard let parts = snapshot?.widget.character?.key.split(separator: "/").map(String.init), parts.count == 3 else { return nil }
    return DeepLink.character(region: parts[0], ruleset: parts[1], name: parts[2]).url
}

/// Next raid (flows F1): small, the glance; medium, the glance and the four answers.
struct NextRaidWidget: SwiftUI.Widget {
    var body: some WidgetConfiguration {
        AppIntentConfiguration(kind: "gg.foreversixty.next-raid", intent: CharacterConfigurationIntent.self, provider: SnapshotProvider()) { entry in
            NextRaidEntryView(entry: entry)
        }
        .configurationDisplayName(Copy.widgetNextRaid)
        .description(Copy.widgetGalleryRaid)
        .supportedFamilies([.systemSmall, .systemMedium, .accessoryInline, .accessoryRectangular, .accessoryCircular])
    }
}

struct NextRaidEntryView: View {
    @Environment(\.widgetFamily) private var family
    let entry: SnapshotEntry

    var body: some View {
        Group {
            switch family {
            case .systemMedium:
                RaidAnswerMediumView(glance: entry.raid, raid: entry.snapshot?.widget.nextRaid, characterKey: entry.characterKey)
            case .accessoryInline: RaidInlineView(glance: entry.raid)
            case .accessoryRectangular: RaidRectangularView(glance: entry.raid)
            case .accessoryCircular: RaidSignedUpCircularView(figures: entry.figures)
            default: NextRaidSmallView(glance: entry.raid)
            }
        }
        .widgetCardBackground()
        .widgetURL(raidURL(entry.snapshot))
    }
}

/// The second circular lock-screen widget: days (or hours, or minutes) to the raid.
struct NextRaidCountdownWidget: SwiftUI.Widget {
    var body: some WidgetConfiguration {
        AppIntentConfiguration(kind: "gg.foreversixty.next-raid-countdown", intent: CharacterConfigurationIntent.self, provider: SnapshotProvider()) { entry in
            RaidCountdownCircularView(figures: entry.figures)
                .widgetCardBackground()
                .widgetURL(raidURL(entry.snapshot))
        }
        .configurationDisplayName(Copy.widgetNextRaid)
        .description(Copy.widgetGalleryRaid)
        .supportedFamilies([.accessoryCircular])
    }
}

/// Next upgrade (flows F5): the next upgrade's gain, never a total.
struct NextUpgradeWidget: SwiftUI.Widget {
    var body: some WidgetConfiguration {
        AppIntentConfiguration(kind: "gg.foreversixty.next-upgrade", intent: CharacterConfigurationIntent.self, provider: SnapshotProvider()) { entry in
            NextUpgradeEntryView(entry: entry)
        }
        .configurationDisplayName(Copy.widgetNextUpgrade)
        .description(Copy.widgetGalleryUpgrade)
        .supportedFamilies([.systemSmall, .accessoryRectangular])
    }
}

struct NextUpgradeEntryView: View {
    @Environment(\.widgetFamily) private var family
    let entry: SnapshotEntry

    var body: some View {
        Group {
            switch family {
            case .accessoryRectangular: UpgradeRectangularView(glance: entry.upgrade)
            default: UpgradeSmallView(glance: entry.upgrade, icons: .shared)
            }
        }
        .widgetCardBackground()
        .widgetURL(characterURL(entry.snapshot))
    }
}

/// Raid and upgrade, medium (the mocks' Medium): the next raid beside the next upgrade.
struct RaidAndUpgradeWidget: SwiftUI.Widget {
    var body: some WidgetConfiguration {
        AppIntentConfiguration(kind: "gg.foreversixty.raid-and-upgrade", intent: CharacterConfigurationIntent.self, provider: SnapshotProvider()) { entry in
            GlanceMediumView(raid: entry.raid, upgrade: entry.upgrade, icons: .shared)
                .widgetCardBackground()
                .widgetURL(raidURL(entry.snapshot) ?? characterURL(entry.snapshot))
        }
        .configurationDisplayName(Copy.widgetGlanceName)
        .description(Copy.widgetGalleryGlance)
        .supportedFamilies([.systemMedium])
    }
}

/// The raid night Live Activity (flows F2): the lock screen, the compact `Ony 42:00` and the
/// expanded Island with the four answers.
struct RaidLiveActivity: SwiftUI.Widget {
    var body: some WidgetConfiguration {
        ActivityConfiguration(for: RaidActivityAttributes.self) { context in
            RaidLiveLockView(attributes: context.attributes, state: context.state, icons: .shared)
                .activityBackgroundTint(FSColor.raised)
                .widgetURL(DeepLink.raid(guild: Int64(context.attributes.guildID) ?? 0, raid: Int64(context.attributes.raidID) ?? 0).url)
        } dynamicIsland: { context in
            DynamicIsland {
                DynamicIslandExpandedRegion(.bottom) {
                    RaidIslandExpandedView(attributes: context.attributes, state: context.state, icons: .shared)
                }
            } compactLeading: {
                RaidIslandCompactLeading(attributes: context.attributes, state: context.state)
            } compactTrailing: {
                LiveCountdown(target: context.state.startDate, fixedNow: nil, role: LiveType.island)
                    .frame(maxWidth: 48)
            } minimal: {
                RaidIslandCompactLeading(attributes: context.attributes, state: context.state)
            }
        }
    }
}

struct SimLiveActivity: SwiftUI.Widget {
    var body: some WidgetConfiguration {
        ActivityConfiguration(for: SimActivityAttributes.self) { context in
            SimLiveLockView(attributes: context.attributes, state: context.state)
                .activityBackgroundTint(FSColor.raised)
                .widgetURL(DeepLink.sim(id: context.attributes.simID).url)
        } dynamicIsland: { context in
            DynamicIsland {
                DynamicIslandExpandedRegion(.bottom) {
                    SimLiveLockView(attributes: context.attributes, state: context.state)
                }
            } compactLeading: {
                Image(systemName: "flask").foregroundStyle(FSColor.gold)
            } compactTrailing: {
                Text(SimLiveFigures(title: context.attributes.title, state: context.state).status)
                    .font(.caption2).lineLimit(1)
            } minimal: {
                Image(systemName: "flask").foregroundStyle(FSColor.gold)
            }
        }
    }
}
