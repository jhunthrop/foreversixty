import ActivityKit
import DesignSystem
import DesignTokens
import Routing
import SwiftUI
import WidgetKit

/// Lane F14's extension: the Next raid and Next upgrade widgets on the home screen and lock
/// screen, and the raid and server-sim Live Activities.
@main
struct ForeverSixtyWidgets: WidgetBundle {
    var body: some SwiftUI.Widget {
        NextRaidWidget()
        NextRaidCountdownWidget()
        NextUpgradeWidget()
        RaidLiveActivity()
        SimLiveActivity()
    }
}

/// One timeline entry: the snapshot as of `date`, read through WidgetFigures at that moment.
struct SnapshotEntry: TimelineEntry {
    let date: Date
    let snapshot: WidgetSnapshot?
    let characterKey: String?
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

private func lastRaidLine(_ snapshot: WidgetSnapshot?) -> String? {
    guard let last = snapshot?.widget.lastRaid else { return nil }
    return Copy.widgetLastRaid(zone: last.zoneName, date: WidgetClock(zone: .current).dayMonth(last.startsAt))
}

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
        let figures = WidgetFigures(entry.snapshot, now: entry.date, zone: .current)
        Group {
            switch family {
            case .systemMedium:
                NextRaidMediumView(figures: figures, lastRaid: lastRaidLine(entry.snapshot), raid: entry.snapshot?.widget.nextRaid, characterKey: entry.characterKey)
            case .accessoryInline: RaidInlineView(figures: figures)
            case .accessoryRectangular: RaidRectangularView(figures: figures)
            case .accessoryCircular: RaidSignedUpCircularView(figures: figures)
            default: NextRaidSmallView(figures: figures, lastRaid: lastRaidLine(entry.snapshot))
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
            RaidCountdownCircularView(figures: WidgetFigures(entry.snapshot, now: entry.date, zone: .current))
                .widgetCardBackground()
                .widgetURL(raidURL(entry.snapshot))
        }
        .configurationDisplayName(Copy.widgetNextRaid)
        .description(Copy.widgetGalleryRaid)
        .supportedFamilies([.accessoryCircular])
    }
}

struct NextUpgradeWidget: SwiftUI.Widget {
    var body: some WidgetConfiguration {
        AppIntentConfiguration(kind: "gg.foreversixty.next-upgrade", intent: CharacterConfigurationIntent.self, provider: SnapshotProvider()) { entry in
            NextUpgradeEntryView(entry: entry)
        }
        .configurationDisplayName(Copy.widgetNextUpgrade)
        .description(Copy.widgetGalleryUpgrade)
        .supportedFamilies([.systemSmall, .systemMedium, .accessoryRectangular])
    }
}

struct NextUpgradeEntryView: View {
    @Environment(\.widgetFamily) private var family
    let entry: SnapshotEntry

    var body: some View {
        let figures = UpgradeFigures(entry.snapshot)
        let character = entry.snapshot?.widget.character
        Group {
            switch family {
            case .systemMedium: UpgradeMediumView(figures: figures, className: character?.className, icons: .shared)
            case .accessoryRectangular: UpgradeRectangularView(figures: figures)
            default: UpgradeSmallView(figures: figures, className: character?.className, icons: .shared)
            }
        }
        .widgetCardBackground()
        // No item route exists in apps/shared/links.json yet, so a tap opens Home on the character.
        .widgetURL(character.flatMap { homeURL($0.key) })
    }

    private func homeURL(_ key: String) -> URL? {
        let parts = key.split(separator: "/").map(String.init)
        guard parts.count == 3 else { return nil }
        return DeepLink.character(region: parts[0], ruleset: parts[1], name: parts[2]).url
    }
}

struct RaidLiveActivity: SwiftUI.Widget {
    var body: some WidgetConfiguration {
        ActivityConfiguration(for: RaidActivityAttributes.self) { context in
            RaidLiveLockView(attributes: context.attributes, state: context.state)
                .activityBackgroundTint(FSColor.raised)
                .widgetURL(DeepLink.raid(guild: Int64(context.attributes.guildID) ?? 0, raid: Int64(context.attributes.raidID) ?? 0).url)
        } dynamicIsland: { context in
            DynamicIsland {
                DynamicIslandExpandedRegion(.bottom) {
                    RaidIslandExpandedView(attributes: context.attributes, state: context.state)
                }
            } compactLeading: {
                SealMark()
            } compactTrailing: {
                LiveCountdown(target: context.state.targetDate, fixedNow: nil, font: .system(.footnote, design: .monospaced, weight: .semibold))
                    .frame(maxWidth: 44)
            } minimal: {
                SealMark()
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
