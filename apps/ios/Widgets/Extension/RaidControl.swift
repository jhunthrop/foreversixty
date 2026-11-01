import AppIntents
import SwiftUI
import WidgetKit

/// Lane F15: the Control Center control `In for the next raid` (spec §5.7 Controls). It answers
/// In through the same sign-up intent, without asking; the press feedback is the system's.
struct RaidInControl: ControlWidget {
    static let kind = "gg.foreversixty.control.raid-in"

    var body: some ControlWidgetConfiguration {
        StaticControlConfiguration(kind: Self.kind) {
            ControlWidgetButton(action: SignUpForRaidIntent(state: .in)) {
                Label(Copy.intentControlTitle, systemImage: "checkmark.circle")
            }
        }
        .displayName(LocalizedStringResource(stringLiteral: Copy.intentControlTitle))
    }
}
