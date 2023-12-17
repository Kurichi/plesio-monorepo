import 'package:flutter/material.dart';

class Item extends StatelessWidget {
  const Item({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 80,
      height: 80,
      decoration: BoxDecoration(
        color: Colors.blue,
        border: Border.all(
          color: Theme.of(context).colorScheme.outline,
          width: 8,
        ),
      ),
      child: Icon(
        Icons.park,
        color: Theme.of(context).colorScheme.primary.withOpacity(0.08),
      ),
    );
  }
}
