import DesignSystem
import SwiftUI

/// A chip or field's pill: the 1 pt line is the outer capsule and the fill an inset one, so the
/// line is built from fills only (a capsule stroke drew stray straight edges past its ends).
extension View {
    func guildPill(fill: Color, line: Color) -> some View {
        background {
            ZStack {
                Capsule().fill(line)
                Capsule().inset(by: 1).fill(FSColor.bg)
                Capsule().inset(by: 1).fill(fill)
            }
        }
        .contentShape(.capsule)
    }
}

/// The v1 switch drawn in the app's own shapes, so its knob shows in every capture: gold track
/// with the knob right when on, the line colour with the knob left when off. VoiceOver reads it
/// as the system switch.
struct GuildSwitchStyle: ToggleStyle {
    static let track = CGSize(width: 51, height: 31)
    static let knob: CGFloat = 27

    func makeBody(configuration: Configuration) -> some View {
        HStack(spacing: 12) {
            configuration.label
            Spacer(minLength: 8)
            ZStack(alignment: configuration.isOn ? .trailing : .leading) {
                Capsule().fill(configuration.isOn ? FSColor.gold : FSColor.line)
                Circle().fill(Color.white).frame(width: Self.knob, height: Self.knob).padding(2)
                    .shadow(color: .black.opacity(0.25), radius: 2, y: 1)
            }
            .frame(width: Self.track.width, height: Self.track.height)
            .animation(.easeInOut(duration: 0.15), value: configuration.isOn)
        }
        .contentShape(.rect)
        .onTapGesture { configuration.isOn.toggle() }
        .accessibilityRepresentation {
            Toggle(isOn: configuration.$isOn) { configuration.label }
        }
    }
}
