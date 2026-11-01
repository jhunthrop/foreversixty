import DesignSystem
import SwiftUI

/// A roster row's spec icon (spec F16.1): the talent-tab icon on the shared icon loader, and the
/// class crest while the spec is unknown or the icon cannot load. Never a blank tile.
struct SpecOrCrest: View {
    let icon: String?
    let className: String?
    let size: CGFloat
    @Environment(\.raidIconBuild) private var build
    @Environment(\.itemIconSource) private var source
    @State private var loaded: UIImage?

    var body: some View {
        Group {
            if let image = shown {
                let shape = RoundedRectangle(cornerRadius: FSRadius.pill)
                Image(uiImage: image).resizable().scaledToFill()
                    .frame(width: size, height: size)
                    .clipShape(shape)
                    .overlay(shape.stroke(GameColors.iconLine, lineWidth: 1))
            } else {
                ClassCrest(className: className, size: size)
            }
        }
        .accessibilityHidden(true)
        .task(id: "\(build ?? "")/\(icon ?? "")") {
            guard let icon, let build, shown == nil else { return }
            loaded = await source.load(icon: icon, build: build)
        }
    }

    private var shown: UIImage? {
        loaded ?? icon.flatMap { icon in build.flatMap { source.cached(icon: icon, build: $0) } }
    }
}

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
