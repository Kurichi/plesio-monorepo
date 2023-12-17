import 'package:kiikuten/domain/repository/item_repository.dart';

class UseItemUseCase {
  final ItemRepository _itemRepository;

  UseItemUseCase(this._itemRepository);

  Future<void> execute(String itemId) async {
    await _itemRepository.useItem(itemId);
  }
}
