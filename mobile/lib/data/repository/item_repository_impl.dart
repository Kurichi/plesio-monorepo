import 'package:kiikuten/data/datasource/item_datasource.dart';
import 'package:kiikuten/domain/entity/item.dart';
import 'package:kiikuten/domain/repository/item_repository.dart';

class ItemRepositoryImpl implements ItemRepository {
  final ItemDataSource _itemDataSource;

  ItemRepositoryImpl(this._itemDataSource);

  @override
  Future<Item> getItem(String itemId) async {
    final itemModel = await _itemDataSource.getItem(itemId);
    return itemModel.toEntity();
  }

  @override
  Future<void> useItem(String itemId) async {
    final item = await getItem(itemId);
    item.use();
  }
}
