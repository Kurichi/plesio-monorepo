import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:kiikuten/data/datasource/item_datasource.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'item_datasource_provider.g.dart';

@riverpod
ItemDataSource itemDataSource(ItemDataSourceRef ref) {
  return ItemDataSource(dotenv.env['BASE_URL'] ?? '');
}
