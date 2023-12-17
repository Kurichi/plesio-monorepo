import 'package:kiikuten/domain/entity/item.dart';
import 'package:kiikuten/domain/repository/item_repository.dart';

class GetItemsUseCase {
  final ItemRepository _itemRepository;

  GetItemsUseCase(this._itemRepository);

  Future<List<Item>> execute() async {
    return await _itemRepository.getItems();
  }
}
