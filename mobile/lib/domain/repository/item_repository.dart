import 'package:kiikuten/domain/entity/item.dart';

abstract class ItemRepository {
  Future<Item> getItem(String itemId);
  Future<void> useItem(String itemId);
}
