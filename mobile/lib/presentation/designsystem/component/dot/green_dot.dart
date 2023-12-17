import 'package:flutter/material.dart';
import 'package:kiikuten/presentation/designsystem/style/color_scheme.dart';

class GreenDot extends StatelessWidget {
  const GreenDot({
    super.key,
    this.opacity = 1,
  });

  final double opacity;

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 8,
      height: 8,
      color: Theme.of(context).kiikutenColorScheme.leaf.withOpacity(opacity),
    );
  }
}
