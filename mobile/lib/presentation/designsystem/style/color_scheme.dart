import 'dart:math';
import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/style/color.dart';

class KiikutenThemeData {
  final ColorScheme colorScheme;

  late final Color _primary;
  late final Color _background;
  late final Color leaf;
  late final Color trunk;

  KiikutenThemeData({
    required this.colorScheme,
    required Brightness brightness,
  }) {
    _primary = colorScheme.primary;
    _background = colorScheme.background;
    leaf = _randomizeColor(KiikutenColor.green);
    trunk = _randomizeColor(KiikutenColor.brown);
  }

  Color _randomizeColor(Color baseColor) {
    const range = 8;
    Random random = Random();
    int r = baseColor.red + random.nextInt(range * 2) - range;
    int g = baseColor.green + random.nextInt(range * 2) - range;
    int b = baseColor.blue + random.nextInt(range * 2) - range;
    return Color.fromRGBO(
      _clampColorValue(r),
      _clampColorValue(g),
      _clampColorValue(b),
      1,
    );
  }

  int _clampColorValue(int value) {
    return max(0, min(255, value));
  }

  Color surfaceColorByAlpha(double alpha, {Color? foreground}) {
    final clr = foreground ?? _primary;
    return Color.alphaBlend(clr.withOpacity(alpha), _background);
  }

  Color surfaceColorByElevation(double elevation) {
    if (elevation <= 0.0) return _background;
    final alpha = ((4.5 * log(elevation + 1)) + 2) / 100;
    return surfaceColorByAlpha(alpha);
  }
}

extension KiikutenThemeDataExtension on ThemeData {
  KiikutenThemeData get kiikutenColorScheme =>
      KiikutenThemeData(colorScheme: colorScheme, brightness: brightness);
}
