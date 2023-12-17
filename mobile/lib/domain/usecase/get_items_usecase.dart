import 'package:kiikuten/domain/entity/item.dart';
import 'package:kiikuten/domain/repository/item_repository.dart';
import 'package:kiikuten/provider/repository/item_repository_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'get_items_usecase.g.dart';

@riverpod
GetItemsUseCase getItemsUseCase(GetItemsUseCaseRef ref) {
  final itemRepository = ref.watch(itemRepositoryProvider);
  return GetItemsUseCase(itemRepository);
}

class GetItemsUseCase {
  final ItemRepository _itemRepository;

  GetItemsUseCase(this._itemRepository);

  Future<List<Item>> execute() async {
    return await _itemRepository.getItems();
  }
}
