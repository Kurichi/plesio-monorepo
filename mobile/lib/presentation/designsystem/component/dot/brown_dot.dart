import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/style/color_scheme.dart';

class BrownDot extends StatelessWidget {
  const BrownDot({
    super.key,
    this.opacity = 1,
  });

  final double opacity;

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 8,
      height: 8,
      color: Theme.of(context).kiikutenColorScheme.trunk.withOpacity(opacity),
    );
  }
}
