import AppIntents
import SwiftUI
import WidgetKit

/// The Control Center control (flows F1, E1): it shows the next raid and your answer as your
/// character, and a press answers In through the same sign-up intent, without asking.
struct RaidControl: ControlWidget {
    static let kind = "gg.foreversixty.control.raid-in"

    var body: some ControlWidgetConfiguration {
        StaticControlConfiguration(kind: Self.kind, provider: RaidControlProvider()) { value in
            ControlWidgetButton(action: SignUpForRaidIntent(state: .in)) {
                Label {
                    Text(value.title)
                    Text(value.subtitle)
                } icon: {
                    Image(systemName: value.answered ? "checkmark.circle.fill" : "checkmark.circle")
                }
            }
        }
        .displayName(LocalizedStringResource(stringLiteral: Copy.intentControlTitle))
    }
}

/// What the control draws: the zone (or `In for the next raid`) over the answer line.
struct RaidControlValue: Equatable, Sendable {
    let title: String
    let subtitle: String
    let answered: Bool

    init(_ glance: RaidGlance) {
        guard glance.kind == .raid else {
            title = Copy.intentControlTitle
            subtitle = glance.name
            answered = false
            return
        }
        title = glance.name
        subtitle = glance.lines.dropFirst().first ?? glance.lines.first ?? ""
        answered = glance.selected == "in"
    }
}

struct RaidControlProvider: ControlValueProvider {
    var previewValue: RaidControlValue {
        RaidControlValue(RaidGlance(nil, now: .now, zone: .current))
    }

    func currentValue() async throws -> RaidControlValue {
        RaidControlValue(RaidGlance(WidgetSnapshotStore.shared().load(), now: .now, zone: .current))
    }
}
