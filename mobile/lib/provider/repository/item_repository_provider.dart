import 'package:kiikuten/data/repository/item_repository_impl.dart';
import 'package:kiikuten/domain/repository/item_repository.dart';
import 'package:kiikuten/provider/datasource/item_datasource_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'item_repository_provider.g.dart';

@riverpod
ItemRepository itemRepository(ItemRepositoryRef ref) {
  final itemDataSource = ref.watch(itemDataSourceProvider);
  return ItemRepositoryImpl(itemDataSource);
}
