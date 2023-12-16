import 'package:kiikuten/domain/entity/item.dart';

abstract class ItemRepository {
  Future<Item> getItem(String itemId);
  Future<void> useItem(String itemId);
}

class ItemRepositoryImpl implements ItemRepository {
  @override
  Future<Item> getItem(String itemId) async {
    return Item(
      id: '1',
      name: 'ミツバチの群れ',
      description: 'ミツバチの群れを呼び寄せる',
      growthEffect: 10,
    );
  }

  @override
  Future<void> useItem(String itemId) async {
    final item = await getItem(itemId);
    item.use();
  }
}
