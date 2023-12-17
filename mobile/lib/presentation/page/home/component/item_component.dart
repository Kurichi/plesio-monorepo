import 'package:flutter/material.dart';
import 'package:kiikuten/domain/entity/item.dart';

class ItemComponent extends StatelessWidget {
  const ItemComponent({
    super.key,
    required this.item,
  });

  final Item item;

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
