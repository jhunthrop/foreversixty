import SwiftUI
import UIKit

/// The site's uppercase section label (`micro`, .14em).
public struct MicroLabel: View {
    private let text: String
    private let color: Color

    public init(_ text: String, color: Color = FSColor.muted) {
        self.text = text
        self.color = color
    }

    public var body: some View {
        Text(text.uppercased())
            .font(ScaledType.micro)
            .tracking(1.5)
            .foregroundStyle(color)
            .accessibilityLabel(text)
    }
}

/// A tag (`.pill`): 22 pt, 11 pt 700 uppercase, tinted ground and border, an optional glyph first.
public struct TagView: View {
    private let text: String
    private let color: Color
    private let systemImage: String?

    public init(text: String, color: Color, systemImage: String? = nil) {
        self.text = text
        self.color = color
        self.systemImage = systemImage
    }

    public var body: some View {
        HStack(spacing: 4) {
            if let systemImage {
                Image(systemName: systemImage).font(.system(size: 9, weight: .bold)).accessibilityHidden(true)
            }
            Text(text.uppercased()).font(ScaledType.micro).tracking(1)
        }
        .foregroundStyle(color)
        .padding(.horizontal, FSSpace.space2)
        .frame(minHeight: 22)
        .background(color.opacity(0.12), in: .rect(cornerRadius: FSRadius.legacyPill))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.legacyPill).stroke(color.opacity(0.35)))
        .fixedSize()
    }
}

/// The offline strip (§5.0): 36 pt, `#11151f`, wifi-off glyph, then the line
/// (`offline.strip` with its stamp).
public struct OfflineStrip: View {
    private let text: String

    public init(text: String) {
        self.text = text
    }

    public var body: some View {
        HStack(spacing: FSSpace.space2) {
            Image(systemName: "wifi.slash").font(.system(size: 14)).accessibilityHidden(true)
            Text(text).font(ScaledType.meta)
            Spacer(minLength: 0)
        }
        .foregroundStyle(FSColor.nav)
        .padding(.horizontal, FSSpace.space4)
        .frame(minHeight: 36)
        .background(GameColors.offlineGround)
        .overlay(alignment: .bottom) { FSColor.line.frame(height: 1) }
        .accessibilityElement(children: .combine)
        .onAppear { AccessibilityNotification.Announcement(text).post() }
    }
}

/// The round-5 art (apps/shared/brand/art, tools/art/build_art.py), bundled as `art/`: the ground
/// tile, the band behind the bars, Home's hero and the raid card's cave. Background only, never
/// captioned (owner, 2026-11-09). Loaded once; a missing file draws the flat token colour instead.
public enum BrandArt {
    public static let ground = load("ground")
    public static let band = load("band")
    public static let hero = load("hero")
    public static let cave = load("cave")
    /// The LX seal v4 (web/public/seal-400.png, bundled at the root): the medallion without a portrait.
    public static let seal = load("seal-400", type: "png", folder: nil)

    private static func load(_ name: String, type: String = "jpg", folder: String? = "art") -> UIImage? {
        Bundle.main.url(forResource: name, withExtension: type, subdirectory: folder)
            .flatMap { UIImage(contentsOfFile: $0.path) }
    }
}

/// The screen ground (spec Mobile, round 5): the painted brown tile at 1:1 (1600 pt, the web's
/// `GROUND_TILE_PX`) over the earth colour, edge to edge under the bars.
public struct ScreenGround: View {
    public init() {}

    public var body: some View {
        ZStack {
            Theme.Palette.ground
            if let tile = BrandArt.ground {
                Image(uiImage: tile).resizable(resizingMode: .tile)
            }
        }
        .ignoresSafeArea()
        .accessibilityHidden(true)
    }
}

/// The painted art behind the bars, under its veil: the blue band (navy into teal cloud) on every
/// tab but Home, the camp-road hero on Home. Stretched to cover, never tiled (spec round 5).
public struct BandArt: View {
    public enum Art: Sendable { case band, hero }

    private let art: Art

    public init(_ art: Art = .band) {
        self.art = art
    }

    public var body: some View {
        ZStack {
            art == .band ? Theme.Palette.band : FSColor.night
            if let image = art == .band ? BrandArt.band : BrandArt.hero {
                Image(uiImage: image).resizable().scaledToFill()
            }
            art == .band ? FSColor.veilBand : FSColor.veilNight
        }
        .clipped()
        .accessibilityHidden(true)
    }
}

/// A panel's frame (spec round 5, "The frame"): none, the plain 1 pt cream line inset 6 pt, or the
/// lead panel's ornate frame (the line plus the owner's corner filigree). Two or three framed
/// panels a screen, one of them ornate.
public enum PanelFrame: Sendable { case none, plain, ornate }

/// The art a lead panel may sit on: the cave behind the raid card, under its veil.
public enum PanelArt: Sendable { case cave }

/// The round-5 panel (`.ui-surface`): their card brown-black, a black hairline inside the warm
/// outer ring, the 10 pt radius; optionally framed, optionally on the cave art.
struct Panel: ViewModifier {
    static let frameInset: CGFloat = 6
    static let cornerInset: CGFloat = 4
    static let cornerSize: CGFloat = 44
    static let cornerOpacity = 0.8

    let frame: PanelFrame
    let art: PanelArt?

    func body(content: Content) -> some View {
        let shape = RoundedRectangle(cornerRadius: Theme.Radius.panel)
        content
            .background { ground.clipShape(shape) }
            .overlay { shape.strokeBorder(Color.black, lineWidth: FSStroke.hairline) }
            .overlay {
                // The warm ring outside the black hairline (`box-shadow: 0 0 0 1px`).
                RoundedRectangle(cornerRadius: Theme.Radius.panel + FSStroke.hairline)
                    .strokeBorder(FSColor.panelEdge, lineWidth: FSStroke.hairline)
                    .padding(-FSStroke.hairline)
            }
            .overlay {
                if frame != .none {
                    RoundedRectangle(cornerRadius: Theme.Radius.panel - Self.cornerInset)
                        .strokeBorder(FSColor.frameLine, lineWidth: FSStroke.hairline)
                        .padding(Self.frameInset)
                }
            }
            .overlay {
                if frame == .ornate { FrameCorners() }
            }
    }

    private var ground: some View {
        ZStack {
            FSColor.raised
            if art == .cave, let cave = BrandArt.cave {
                Image(uiImage: cave).resizable().scaledToFill()
                FSColor.veilCave
            }
        }
    }
}

/// The lead panel's four corner ornaments: the owner's filigree (design/assets/ornaments, the
/// `frame-corner` vector in Brand.xcassets), cream at 80%, mirrored per corner. Decorative only.
public struct FrameCorners: View {
    public init() {}

    public var body: some View {
        ZStack {
            corner(.topLeading, x: 1, y: 1)
            corner(.topTrailing, x: -1, y: 1)
            corner(.bottomTrailing, x: -1, y: -1)
            corner(.bottomLeading, x: 1, y: -1)
        }
        .padding(Panel.cornerInset)
        .allowsHitTesting(false)
        .accessibilityHidden(true)
    }

    private func corner(_ alignment: Alignment, x: CGFloat, y: CGFloat) -> some View {
        Image("frame-corner", bundle: .main)
            .renderingMode(.template)
            .resizable()
            .frame(width: Panel.cornerSize, height: Panel.cornerSize)
            .scaleEffect(x: x, y: y)
            .foregroundStyle(FSColor.cream.opacity(Panel.cornerOpacity))
            .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: alignment)
    }
}

/// The band ground (`.surfaceBand`): the painted band under its veil, with the black bottom edge
/// when `bottomLine` (the mock's `.band`). Home's hero bands (with the line) and the signed-out
/// screens' full-bleed ground (without).
struct SurfaceBand: ViewModifier {
    let bottomLine: Bool

    func body(content: Content) -> some View {
        content
            .background { BandArt() }
            .overlay(alignment: .bottom) {
                if bottomLine { Color.black.frame(height: FSStroke.hairline) }
            }
    }
}

public extension View {
    /// A round-5 panel, optionally framed and on the cave art.
    func panel(frame: PanelFrame = .none, art: PanelArt? = nil) -> some View { modifier(Panel(frame: frame, art: art)) }
    /// A featured card: the plain-framed panel.
    func primaryCard() -> some View { panel(frame: .plain) }
    /// A list card: the plain panel.
    func quietCard() -> some View { panel() }
    func surfaceBand(bottomLine: Bool = true) -> some View { modifier(SurfaceBand(bottomLine: bottomLine)) }
    /// The screen's painted ground behind the content.
    func screenGround() -> some View { background { ScreenGround() } }
}

/// A list row's press: the 10% gold wash while held (spec §8).
public struct RowPressStyle: ButtonStyle {
    public init() {}

    public func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .background(configuration.isPressed ? GameColors.pressedWash : .clear)
    }
}

public extension View {
    /// A chip or pill's ground: the capsule filled with `fill` and its 1 pt border inside the edge.
    func pillSurface(fill: Color, line: Color) -> some View {
        background { Capsule().fill(fill) }
            .overlay { Capsule().strokeBorder(line, lineWidth: 1) }
            .contentShape(.capsule)
    }
}
