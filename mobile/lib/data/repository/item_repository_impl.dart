import 'package:kiikuten/domain/entity/item.dart';
import 'package:kiikuten/domain/repository/item_repository.dart';

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
