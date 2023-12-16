import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:kiikuten/domain/entity/item.dart';

part 'item_model.freezed.dart';

part 'item_model.g.dart';

@freezed
class ItemModel with _$ItemModel {
  const ItemModel._();

  const factory ItemModel({
    required String id,
    required String name,
    required String description,
    required int growthEffect,
  }) = _ItemModel;

  factory ItemModel.fromJson(Map<String, dynamic> json) =>
      _$ItemModelFromJson(json);

  Item toEntity() {
    return Item(
      id: id,
      name: name,
      description: description,
      growthEffect: growthEffect,
    );
  }
}
