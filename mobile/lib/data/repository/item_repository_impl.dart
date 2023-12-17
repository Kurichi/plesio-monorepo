import 'package:kiikuten/data/datasource/item_datasource.dart';
import 'package:kiikuten/domain/entity/item.dart';
import 'package:kiikuten/domain/repository/item_repository.dart';

class ItemRepositoryImpl implements ItemRepository {
  final ItemDataSource _itemDataSource;

  ItemRepositoryImpl(this._itemDataSource);

  @override
  Future<List<Item>> getItems() async {
    final itemModels = await _itemDataSource.getItems();
    return itemModels.map((model) => model.toEntity()).toList();
  }

  @override
  Future<void> useItem(String itemId) async {
    await _itemDataSource.useItem(itemId);
  }
}
