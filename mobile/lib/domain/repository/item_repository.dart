import 'package:kiikuten/domain/entity/item.dart';

abstract class ItemRepository {
  Future<List<Item>> getItems();
  Future<void> useItem(String itemId);
}
