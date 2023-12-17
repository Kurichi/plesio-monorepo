import 'package:flutter/material.dart';
import 'package:kiikuten/domain/entity/item.dart';
import 'package:kiikuten/presentation/page/home/component/item_component.dart';

class ItemContainer extends StatelessWidget {
  const ItemContainer({
    super.key,
    required this.items,
  });

  final List<Item> items;

  @override
  Widget build(BuildContext context) {
    List<Widget> buildItems() {
      if (items.any((item) => item.id == 'loading')) {
        return [const CircularProgressIndicator()];
      } else if (items.any((item) => item.id == 'error')) {
        return [
          Text(
            items.first.description,
            style: const TextStyle(color: Colors.red),
          ),
        ];
      } else {
        return items.map((item) => ItemComponent(item: item)).toList();
      }
    }

    return Container(
      margin: const EdgeInsets.all(32),
      width: double.infinity,
      height: 100,
      decoration: BoxDecoration(
        color: Colors.white,
        border: Border.all(
          color: Theme.of(context).colorScheme.outline,
          width: 8,
        ),
        boxShadow: const [
          BoxShadow(
            color: Color(0x20000000),
            offset: Offset(8, 8),
          ),
        ],
      ),
      child: SingleChildScrollView(
        padding: const EdgeInsets.all(8),
        scrollDirection: Axis.horizontal,
        child: Row(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: buildItems(),
        ),
      ),
    );
  }
}
