import 'package:kiikuten/domain/repository/item_repository.dart';
import 'package:kiikuten/provider/repository/item_repository_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'use_item_usecase.g.dart';

@riverpod
UseItemUseCase getItemsUseCase(GetItemsUseCaseRef ref) {
  final itemRepository = ref.watch(itemRepositoryProvider);
  return UseItemUseCase(itemRepository);
}

class UseItemUseCase {
  final ItemRepository _itemRepository;

  UseItemUseCase(this._itemRepository);

  Future<void> execute(String itemId) async {
    await _itemRepository.useItem(itemId);
  }
}
